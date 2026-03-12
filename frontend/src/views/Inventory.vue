<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { Plus, Search, Scan, MoreHorizontal, ArrowUpDown, AlertTriangle, Package } from 'lucide-vue-next';

const items = ref<any[]>([]);
const loading = ref(true);
const configStore = useConfigStore();
const isAddModalOpen = ref(false);
const isStockModalOpen = ref(false);
const selectedItem = ref<any>(null);
const stockAdjustment = ref(0);

const newItem = ref({
  name: '',
  barcode: '',
  sku: '',
  price: 0,
  current_stock: 0,
  min_stock_alert: 0,
  unit: 'KG'
});

const fetchItems = async () => {
  try {
    const response = await api.get('/api/inventory/items');
    items.value = response.data;
  } catch (error) {
    console.error('Failed to fetch items', error);
  } finally {
    loading.value = false;
  }
};

const addItem = async () => {
  try {
    await api.post('/api/inventory/items', newItem.value);
    isAddModalOpen.value = false;
    newItem.value = { name: '', barcode: '', sku: '', price: 0, current_stock: 0, min_stock_alert: 0, unit: 'KG' };
    await fetchItems();
  } catch (error) {
    console.error('Failed to add item', error);
  }
};

const openStockModal = (item: any) => {
  selectedItem.value = item;
  stockAdjustment.value = 0;
  isStockModalOpen.value = true;
};

const updateStock = async () => {
  try {
    await api.patch(`/api/inventory/items/${selectedItem.value.id}/stock`, {
      quantity: stockAdjustment.value
    });
    isStockModalOpen.value = false;
    await fetchItems();
  } catch (error: any) {
    alert(error.response?.data?.error || 'Failed to update stock');
  }
};

onMounted(fetchItems);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-6">
      <!-- Action Bar -->
      <div class="flex flex-col md:flex-row justify-between gap-4">
        <div class="relative flex-1 max-w-md">
          <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
          <input 
            type="text" 
            :placeholder="$t('common.search')" 
            class="w-full pl-12 pr-4 py-3 bg-white border border-slate-200 rounded-2xl shadow-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all outline-none"
          />
        </div>
        <div class="flex gap-3">
          <button class="flex items-center gap-2 bg-white border border-slate-200 text-slate-700 px-6 py-3 rounded-2xl font-bold hover:bg-slate-50 transition-all shadow-sm">
            <Scan class="w-5 h-5" />
            <span>{{ $t('inventory.scan', 'Scan') }}</span>
          </button>
          <button 
            @click="isAddModalOpen = true"
            class="flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-2xl font-bold hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100"
          >
            <Plus class="w-5 h-5" />
            <span>{{ $t('inventory.addIngredient') }}</span>
          </button>
        </div>
      </div>

      <!-- Inventory Table -->
      <div class="bg-white rounded-3xl shadow-sm border border-slate-100 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 border-b border-slate-100">
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <div class="flex items-center gap-2 cursor-pointer hover:text-slate-600">
                  {{ $t('inventory.itemDetails') }}
                  <ArrowUpDown class="w-3 h-3" />
                </div>
              </th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('common.price') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('inventory.stockLevel') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('common.status') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="item in items" :key="item.id" class="group hover:bg-slate-50/50 transition-colors">
              <td class="px-8 py-5">
                <div class="flex flex-col">
                  <span class="font-bold text-slate-800">{{ item.name }}</span>
                  <span class="text-[10px] text-slate-400 font-medium uppercase tracking-tighter">{{ item.barcode || item.sku || 'No Identifier' }}</span>
                </div>
              </td>
              <td class="px-8 py-5 text-right font-black text-slate-700">
                {{ configStore.formatCurrency(item.price) }}
              </td>
              <td class="px-8 py-5 text-right">
                <div class="flex flex-col items-end">
                  <span 
                    class="font-black text-lg" 
                    :class="item.current_stock <= item.min_stock_alert ? 'text-red-600' : 'text-slate-800'"
                  >
                    {{ item.current_stock }}
                  </span>
                  <span class="text-[10px] font-bold text-slate-400 uppercase">{{ item.unit }}</span>
                </div>
              </td>
              <td class="px-8 py-5">
                <div 
                  v-if="item.current_stock <= item.min_stock_alert"
                  class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-red-50 text-red-600 text-[10px] font-black uppercase tracking-wider border border-red-100"
                >
                  <AlertTriangle class="w-3 h-3" />
                  {{ $t('inventory.lowStock') }}
                </div>
                <div 
                  v-else
                  class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-green-50 text-green-600 text-[10px] font-black uppercase tracking-wider border border-green-100"
                >
                  <div class="w-1.5 h-1.5 bg-green-500 rounded-full"></div>
                  {{ $t('inventory.inStock') }}
                </div>
              </td>
              <td class="px-8 py-5 text-right">
                <div class="flex justify-end gap-2">
                  <button 
                    @click="openStockModal(item)"
                    class="p-2 text-indigo-600 hover:bg-indigo-50 rounded-xl transition-colors font-bold text-xs"
                  >
                    {{ $t('inventory.updateStock') }}
                  </button>
                  <button class="p-2 text-slate-400 hover:bg-slate-100 rounded-xl transition-colors">
                    <MoreHorizontal class="w-5 h-5" />
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="items.length === 0 && !loading">
              <td colspan="5" class="px-8 py-20 text-center">
                <div class="flex flex-col items-center gap-3 opacity-40">
                  <Package class="w-16 h-16" />
                  <p class="font-bold text-slate-500 italic">{{ $t('common.noData') }}</p>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Update Stock Modal -->
    <div v-if="isStockModalOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center p-4 z-50 animate-in fade-in duration-200">
      <div class="bg-white rounded-[2rem] w-full max-w-sm p-8 shadow-2xl scale-in-center">
        <h3 class="text-xl font-black text-slate-800 mb-1">{{ $t('inventory.updateStock') }}</h3>
        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-6">{{ selectedItem?.name }}</p>
        
        <div class="bg-slate-50 p-4 rounded-2xl mb-6 flex justify-between items-center border border-slate-100">
          <span class="text-xs font-bold text-slate-500 uppercase">{{ $t('common.current', 'Current') }}</span>
          <span class="font-black text-slate-800">{{ selectedItem?.current_stock }} {{ selectedItem?.unit }}</span>
        </div>
        
        <form @submit.prevent="updateStock" class="space-y-6">
          <div>
            <label class="block text-xs font-black text-slate-500 uppercase tracking-widest mb-2">{{ $t('inventory.adjustmentQty') }}</label>
            <input 
              v-model.number="stockAdjustment" 
              type="number" 
              step="0.1" 
              required 
              autofocus 
              class="w-full px-4 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-black text-xl text-center focus:bg-white focus:border-indigo-600 transition-all outline-none" 
            />
            <p class="text-[10px] text-center font-bold text-slate-400 mt-2">{{ $t('inventory.stockNote') }}</p>
          </div>
          
          <div class="flex flex-col gap-2">
            <button type="submit" class="w-full py-4 bg-indigo-600 text-white rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all">
              {{ $t('common.confirm') }}
            </button>
            <button type="button" @click="isStockModalOpen = false" class="w-full py-3 text-slate-400 font-bold hover:text-slate-600 transition-colors">
              {{ $t('common.cancel') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Add Item Modal -->
    <div v-if="isAddModalOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center p-4 z-50 animate-in fade-in duration-200">
      <div class="bg-white rounded-[2rem] w-full max-w-lg p-8 shadow-2xl scale-in-center">
        <h3 class="text-2xl font-black text-slate-800 mb-6">{{ $t('inventory.newIngredient') }}</h3>
        
        <form @submit.prevent="addItem" class="space-y-5">
          <div class="space-y-1">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('inventory.ingredientName') }}</label>
            <input v-model="newItem.name" type="text" required class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none" />
          </div>
          
          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('inventory.barcodeSku') }}</label>
              <input v-model="newItem.barcode" type="text" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none" />
            </div>
            <div class="space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('common.unit') }}</label>
              <select v-model="newItem.unit" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none appearance-none font-bold text-slate-700">
                <option>KG</option><option>GR</option><option>Piece</option><option>Litre</option><option>ML</option>
              </select>
            </div>
          </div>
          
          <div class="grid grid-cols-3 gap-4">
            <div class="space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('inventory.cost') }}</label>
              <input v-model.number="newItem.price" type="number" step="0.01" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" />
            </div>
            <div class="space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('inventory.initialStock') }}</label>
              <input v-model.number="newItem.current_stock" type="number" step="0.1" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" />
            </div>
            <div class="space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('inventory.minAlert') }}</label>
              <input v-model.number="newItem.min_stock_alert" type="number" step="0.1" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" />
            </div>
          </div>
          
          <div class="flex gap-3 mt-8">
            <button type="button" @click="isAddModalOpen = false" class="flex-1 py-4 text-slate-400 font-bold hover:text-slate-600 transition-colors">{{ $t('common.cancel') }}</button>
            <button type="submit" class="flex-[2] py-4 bg-indigo-600 text-white rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all">{{ $t('common.save') }}</button>
          </div>
        </form>
      </div>
    </div>
  </DashboardLayout>
</template>
