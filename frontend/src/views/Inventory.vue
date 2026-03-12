<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import ViewToggle from '../components/ViewToggle.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { Plus, Search, Scan, ArrowUpDown, AlertTriangle, Package, Edit2, Trash2, Power, PowerOff } from 'lucide-vue-next';

const items = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);
const configStore = useConfigStore();
const viewMode = ref<'list' | 'grid'>('list');

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isStockModalOpen = ref(false);
const isDeleteModalOpen = ref(false);

const selectedItem = ref<any>(null);
const stockAdjustment = ref(0);

const alertConfig = ref({
  show: false,
  title: '',
  message: '',
  type: 'info' as 'info' | 'success' | 'warning' | 'danger'
});

const showAlert = (title: string, message: string, type: any = 'info') => {
  alertConfig.value = { show: true, title, message, type };
};

const newItem = ref({
  name: '',
  barcode: '',
  sku: '',
  price: 0,
  current_stock: 0,
  min_stock_alert: 0,
  unit: 'KG'
});

const editItemData = ref({
  id: 0,
  name: '',
  barcode: '',
  sku: '',
  price: 0,
  min_stock_alert: 0,
  unit: 'KG'
});

const fetchItems = async () => {
  loading.value = true;
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
  actionLoading.value = true;
  try {
    await api.post('/api/inventory/items', newItem.value);
    isAddModalOpen.value = false;
    newItem.value = { name: '', barcode: '', sku: '', price: 0, current_stock: 0, min_stock_alert: 0, unit: 'KG' };
    await fetchItems();
    showAlert("Success", "New ingredient has been added.", "success");
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (item: any) => {
  editItemData.value = { 
    id: item.id,
    name: item.name, 
    barcode: item.barcode, 
    sku: item.sku, 
    price: item.price, 
    min_stock_alert: item.min_stock_alert, 
    unit: item.unit 
  };
  isEditModalOpen.value = true;
};

const updateItem = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/inventory/items/${editItemData.value.id}`, editItemData.value);
    isEditModalOpen.value = false;
    await fetchItems();
    showAlert("Success", "Ingredient details updated.", "success");
  } finally {
    actionLoading.value = false;
  }
};

const openStockModal = (item: any) => {
  selectedItem.value = item;
  stockAdjustment.value = 0;
  isStockModalOpen.value = true;
};

const updateStock = async () => {
  actionLoading.value = true;
  try {
    await api.patch(`/api/inventory/items/${selectedItem.value.id}/stock`, {
      quantity: stockAdjustment.value
    });
    isStockModalOpen.value = false;
    await fetchItems();
    showAlert("Stock Updated", "Ingredient inventory has been adjusted.", "success");
  } catch (error: any) {
    showAlert("Adjustment Failed", error.response?.data?.error || 'Failed to update stock', "danger");
  } finally {
    actionLoading.value = false;
  }
};

const toggleActive = async (item: any) => {
  try {
    await api.patch(`/api/inventory/items/${item.id}/toggle`);
    await fetchItems();
  } catch (error) {
    console.error('Failed to toggle status', error);
  }
};

const deleteItem = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/inventory/items/${selectedItem.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchItems();
    showAlert("Deleted", "Ingredient has been removed from inventory.", "info");
  } finally {
    actionLoading.value = false;
  }
};

onMounted(fetchItems);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-6">
      <!-- Action Bar -->
      <div class="flex flex-col md:flex-row justify-between gap-4">
        <div class="flex flex-1 gap-3 max-w-2xl">
          <div class="relative flex-1">
            <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
            <input 
              type="text" 
              :placeholder="$t('common.search')" 
              class="w-full pl-12 pr-4 py-3 bg-white border border-slate-200 rounded-2xl shadow-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all outline-none"
            />
          </div>
          <ViewToggle v-model:mode="viewMode" />
        </div>
        <div class="flex gap-3">
          <button class="flex items-center gap-2 bg-white border border-slate-200 text-slate-700 px-6 py-3 rounded-2xl font-bold hover:bg-slate-50 transition-all shadow-sm">
            <Scan class="w-5 h-5" />
            <span class="hidden sm:inline">{{ $t('inventory.scan', 'Scan') }}</span>
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

      <!-- List View -->
      <div v-if="viewMode === 'list'" class="bg-white rounded-3xl shadow-sm border border-slate-100 overflow-hidden animate-in fade-in duration-300">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 border-b border-slate-100">
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('inventory.itemDetails') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('common.price') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('inventory.stockLevel') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('common.status') }}</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="item in items" :key="item.id" class="group transition-colors" :class="item.is_active ? 'hover:bg-slate-50/50' : 'bg-slate-50/30 opacity-60'">
              <td class="px-8 py-5">
                <div class="flex flex-col">
                  <span class="font-bold text-slate-800">{{ item.name }}</span>
                  <span class="text-[10px] text-slate-400 font-medium uppercase tracking-tighter">{{ item.barcode || item.sku || 'No Identifier' }}</span>
                </div>
              </td>
              <td class="px-8 py-5 text-right font-black text-slate-700">{{ configStore.formatCurrency(item.price) }}</td>
              <td class="px-8 py-5 text-right">
                <div class="flex flex-col items-end">
                  <span class="font-black text-lg" :class="item.current_stock <= item.min_stock_alert ? 'text-red-600' : 'text-slate-800'">{{ item.current_stock }}</span>
                  <span class="text-[10px] font-bold text-slate-400 uppercase">{{ item.unit }}</span>
                </div>
              </td>
              <td class="px-8 py-5">
                <div class="flex flex-col gap-1">
                  <div v-if="item.current_stock <= item.min_stock_alert" class="inline-flex items-center gap-1.5 px-2 py-0.5 rounded-full bg-red-50 text-red-600 text-[9px] font-black uppercase tracking-wider border border-red-100 w-fit">
                    <AlertTriangle class="w-2.5 h-2.5" />
                    {{ $t('inventory.lowStock') }}
                  </div>
                  <div class="inline-flex items-center gap-1.5 px-2 py-0.5 rounded-full text-[9px] font-black uppercase tracking-wider border w-fit" :class="item.is_active ? 'bg-green-50 text-green-600 border-green-100' : 'bg-slate-100 text-slate-400 border-slate-200'">
                    <div class="w-1.5 h-1.5 rounded-full" :class="item.is_active ? 'bg-green-500' : 'bg-slate-400'"></div>
                    {{ item.is_active ? $t('common.active') : $t('common.inactive') }}
                  </div>
                </div>
              </td>
              <td class="px-8 py-5 text-right">
                <div class="flex justify-end gap-1">
                  <button @click="openStockModal(item)" class="p-2 text-slate-400 hover:text-indigo-600 hover:bg-indigo-50 rounded-xl transition-all" :title="$t('inventory.updateStock')"><ArrowUpDown class="w-4 h-4" /></button>
                  <button @click="openEditModal(item)" class="p-2 text-slate-400 hover:text-blue-600 hover:bg-blue-50 rounded-xl transition-all" :title="$t('common.edit')"><Edit2 class="w-4 h-4" /></button>
                  <button @click="toggleActive(item)" class="p-2 transition-all rounded-xl" :class="item.is_active ? 'text-slate-400 hover:text-orange-600 hover:bg-orange-50' : 'text-orange-600 bg-orange-50 hover:bg-orange-100'"><component :is="item.is_active ? PowerOff : Power" class="w-4 h-4" /></button>
                  <button @click="selectedItem = item; isDeleteModalOpen = true" class="p-2 text-slate-400 hover:text-red-600 hover:bg-red-50 rounded-xl transition-all"><Trash2 class="w-4 h-4" /></button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Grid View -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 animate-in fade-in duration-300">
        <div v-for="item in items" :key="item.id" class="bg-white rounded-[2rem] shadow-sm border border-slate-100 p-6 flex flex-col group transition-all" :class="item.is_active ? 'hover:shadow-xl hover:border-indigo-100' : 'opacity-60 grayscale bg-slate-50'">
          <div class="flex justify-between items-start mb-6">
            <div class="p-3 bg-slate-50 rounded-2xl group-hover:bg-indigo-600 group-hover:text-white transition-colors">
              <Package class="w-6 h-6" />
            </div>
            <div class="flex gap-1">
              <button @click="openEditModal(item)" class="p-2 text-slate-400 hover:text-blue-600 transition-colors"><Edit2 class="w-4 h-4" /></button>
              <button @click="toggleActive(item)" class="p-2 transition-all rounded-lg" :class="item.is_active ? 'text-slate-400 hover:text-orange-600' : 'text-orange-600'"><component :is="item.is_active ? PowerOff : Power" class="w-4 h-4" /></button>
            </div>
          </div>

          <div class="mb-6">
            <h4 class="font-black text-slate-800 text-lg leading-tight mb-1 truncate">{{ item.name }}</h4>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ item.barcode || 'No Identifier' }}</p>
          </div>

          <div class="mt-auto pt-6 border-t border-slate-50 flex items-center justify-between">
            <div>
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('inventory.stockLevel') }}</p>
              <p class="font-black text-xl" :class="item.current_stock <= item.min_stock_alert ? 'text-red-600' : 'text-slate-800'">
                {{ item.current_stock }} <span class="text-xs font-bold text-slate-400 uppercase">{{ item.unit }}</span>
              </p>
            </div>
            <button @click="openStockModal(item)" class="w-10 h-10 rounded-xl bg-indigo-50 text-indigo-600 flex items-center justify-center hover:bg-indigo-600 hover:text-white transition-all shadow-sm">
              <ArrowUpDown class="w-5 h-5" />
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="items.length === 0 && !loading" class="py-32 text-center bg-white rounded-[3rem] border-4 border-dashed border-slate-100">
        <div class="max-w-xs mx-auto space-y-4 opacity-40">
          <Package class="w-20 h-20 mx-auto text-slate-400" />
          <p class="text-xl font-black text-slate-500 uppercase">{{ $t('common.noData') }}</p>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <FormModal :show="isAddModalOpen" :title="$t('inventory.newIngredient')" :icon="Package" :loading="actionLoading" @submit="addItem" @cancel="isAddModalOpen = false">
      <div class="space-y-5">
        <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.ingredientName') }}</label><input v-model="newItem.name" type="text" required class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none" /></div>
        <div class="grid grid-cols-2 gap-4">
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.barcodeSku') }}</label><input v-model="newItem.barcode" type="text" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none" /></div>
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('common.unit') }}</label><select v-model="newItem.unit" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 outline-none appearance-none font-bold"><option>KG</option><option>GR</option><option>Piece</option><option>Litre</option><option>ML</option></select></div>
        </div>
        <div class="grid grid-cols-3 gap-4">
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.cost') }}</label><input v-model.number="newItem.price" type="number" step="0.01" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" /></div>
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.initialStock') }}</label><input v-model.number="newItem.current_stock" type="number" step="0.1" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" /></div>
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.minAlert') }}</label><input v-model.number="newItem.min_stock_alert" type="number" step="0.1" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" /></div>
        </div>
      </div>
    </FormModal>

    <FormModal :show="isEditModalOpen" :title="$t('common.edit') + ' ' + $t('inventory.ingredientName')" :icon="Edit2" :loading="actionLoading" @submit="updateItem" @cancel="isEditModalOpen = false">
      <div class="space-y-5">
        <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.ingredientName') }}</label><input v-model="editItemData.name" type="text" required class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" /></div>
        <div class="grid grid-cols-2 gap-4">
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.barcodeSku') }}</label><input v-model="editItemData.barcode" type="text" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none" /></div>
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('common.unit') }}</label><select v-model="editItemData.unit" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 outline-none font-bold"><option>KG</option><option>GR</option><option>Piece</option><option>Litre</option><option>ML</option></select></div>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.cost') }}</label><input v-model.number="editItemData.price" type="number" step="0.01" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" /></div>
          <div class="space-y-1"><label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.minAlert') }}</label><input v-model.number="editItemData.min_stock_alert" type="number" step="0.1" class="w-full px-4 py-3 bg-slate-50 border-2 border-transparent rounded-xl focus:bg-white focus:border-indigo-600 transition-all outline-none font-bold" /></div>
        </div>
      </div>
    </FormModal>

    <FormModal :show="isStockModalOpen" :title="$t('inventory.updateStock')" subtitle="Stock Adjustment" :icon="ArrowUpDown" :loading="actionLoading" maxWidth="max-w-sm" @submit="updateStock" @cancel="isStockModalOpen = false">
      <div class="space-y-6">
        <div class="bg-slate-50 p-4 rounded-2xl flex justify-between items-center border border-slate-100"><span class="text-[10px] font-black text-slate-400 uppercase">{{ $t('common.current') }}</span><span class="font-black text-slate-800">{{ selectedItem?.current_stock }} {{ selectedItem?.unit }}</span></div>
        <input v-model.number="stockAdjustment" type="number" step="0.1" required autofocus class="w-full px-4 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-black text-xl text-center focus:bg-white focus:border-indigo-600 transition-all outline-none" />
        <p class="text-[10px] text-center font-bold text-slate-400">{{ $t('inventory.stockNote') }}</p>
      </div>
    </FormModal>

    <ConfirmModal :show="isDeleteModalOpen" type="danger" :title="$t('common.delete') + ' ' + $t('common.confirm')" :message="`Are you sure you want to delete ${selectedItem?.name}? This action cannot be undone.`" :loading="actionLoading" @confirm="deleteItem" @cancel="isDeleteModalOpen = false" />

    <!-- Universal Alert Modal -->
    <ConfirmModal 
      :show="alertConfig.show"
      :title="alertConfig.title"
      :message="alertConfig.message"
      :type="alertConfig.type"
      :showCancel="false"
      @confirm="alertConfig.show = false"
    />
  </DashboardLayout>
</template>
