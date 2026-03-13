<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import ViewToggle from '../components/ViewToggle.vue';
import api from '../api/axios';
import { Package, Plus, Search, Trash2, Edit2, AlertCircle, Barcode, Activity } from 'lucide-vue-next';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const items = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);
const searchQuery = ref('');
const viewMode = ref<'list' | 'grid'>('grid');

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isStockModalOpen = ref(false);
const isDeleteModalOpen = ref(false);

const selectedItem = ref<any>(null);
const itemForm = ref({
  id: 0,
  name: '',
  barcode: '',
  cost_price: 0,
  stock: 0,
  unit: 'GR',
  min_stock_alert: 0
});

const stockAdjustment = ref({
  quantity: 0,
  type: 'add'
});

const alertConfig = ref({
  show: false,
  title: '',
  message: '',
  type: 'info' as 'info' | 'success' | 'warning' | 'danger'
});

const showAlert = (title: string, message: string, type: any = 'info') => {
  alertConfig.value = { show: true, title, message, type };
};

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

const createItem = async () => {
  actionLoading.value = true;
  try {
    await api.post('/api/inventory/items', itemForm.value);
    isAddModalOpen.value = false;
    itemForm.value = { id: 0, name: '', barcode: '', cost_price: 0, stock: 0, unit: 'GR', min_stock_alert: 0 };
    await fetchItems();
    showAlert(t('common.success'), t('inventory.successAdd'), "success");
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || "Action failed", "danger");
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (item: any) => {
  itemForm.value = { ...item };
  isEditModalOpen.value = true;
};

const updateItem = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/inventory/items/${itemForm.value.id}`, itemForm.value);
    isEditModalOpen.value = false;
    await fetchItems();
    showAlert(t('common.success'), t('inventory.successUpdate'), "success");
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || "Update failed", "danger");
  } finally {
    actionLoading.value = false;
  }
};

const openStockModal = (item: any) => {
  selectedItem.value = item;
  stockAdjustment.value = { quantity: 0, type: 'add' };
  isStockModalOpen.value = true;
};

const updateStock = async () => {
  actionLoading.value = true;
  try {
    const finalQty = stockAdjustment.value.type === 'add' ? stockAdjustment.value.quantity : -stockAdjustment.value.quantity;
    await api.post(`/api/inventory/items/${selectedItem.value.id}/stock`, {
      adjustment: finalQty
    });
    isStockModalOpen.value = false;
    await fetchItems();
    showAlert(t('common.success'), t('inventory.successStock'), "success");
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || "Adjustment failed", "danger");
  } finally {
    actionLoading.value = false;
  }
};

const deleteItem = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/inventory/items/${selectedItem.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchItems();
    showAlert(t('common.success'), t('inventory.successDelete'), "success");
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || "Delete failed", "danger");
  } finally {
    actionLoading.value = false;
  }
};

onMounted(fetchItems);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8 animate-in fade-in duration-500">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div class="flex flex-1 gap-4 items-center w-full md:w-auto">
          <h1 class="text-2xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-3">
            <Package class="w-8 h-8 text-indigo-600" />
            {{ t('common.inventory') }}
          </h1>
          <ViewToggle v-model:mode="viewMode" />
        </div>
        <div class="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
          <div class="relative flex-1 sm:w-64">
            <input v-model="searchQuery" type="text" :placeholder="t('common.search')" class="w-full pl-12 pr-4 py-4 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl font-bold shadow-sm focus:border-indigo-600 outline-none transition-all dark:text-slate-100" />
            <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300 dark:text-slate-600" />
          </div>
          <button @click="isAddModalOpen = true" class="bg-indigo-600 text-white px-8 py-4 rounded-2xl font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2">
            <Plus class="w-5 h-5" />
            {{ t('inventory.addIngredient') }}
          </button>
        </div>
      </div>

      <!-- List View -->
      <div v-if="viewMode === 'list'" class="bg-white dark:bg-slate-900 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('inventory.ingredientName') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('inventory.stockLevel') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('inventory.minAlert') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="item in items" :key="item.id" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors">
              <td class="px-8 py-6">
                <div class="flex flex-col">
                  <span class="font-bold text-slate-800 dark:text-slate-100">{{ item.name }}</span>
                  <span class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ item.barcode }}</span>
                </div>
              </td>
              <td class="px-8 py-6 text-right">
                <div class="flex flex-col items-end">
                  <span class="text-lg font-black" :class="item.stock <= item.min_stock_alert ? 'text-red-600' : 'text-slate-700 dark:text-slate-300'">{{ item.stock }} <span class="text-[10px] text-slate-400 dark:text-slate-500">{{ item.unit }}</span></span>
                  <span v-if="item.stock <= item.min_stock_alert" class="text-[9px] font-black uppercase text-red-500 bg-red-50 dark:bg-red-900/30 px-2 py-0.5 rounded-full">{{ t('inventory.lowStock') }}</span>
                </div>
              </td>
              <td class="px-8 py-6 text-right font-black text-slate-400 dark:text-slate-500">{{ item.min_stock_alert }} {{ item.unit }}</td>
              <td class="px-8 py-6 text-right">
                <div class="flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button @click="openStockModal(item)" class="p-2 text-indigo-600 hover:bg-indigo-50 dark:hover:bg-indigo-900/30 rounded-xl transition-all"><Activity class="w-5 h-5" /></button>
                  <button @click="openEditModal(item)" class="p-2 text-slate-400 hover:text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/30 rounded-xl transition-all"><Edit2 class="w-5 h-5" /></button>
                  <button @click="selectedItem = item; isDeleteModalOpen = true" class="p-2 text-slate-400 hover:text-red-600 hover:bg-red-50 dark:hover:bg-red-900/30 rounded-xl transition-all"><Trash2 class="w-5 h-5" /></button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Grid View -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <div v-for="item in items" :key="item.id" class="bg-white dark:bg-slate-900 p-6 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all duration-300 flex flex-col group relative overflow-hidden">
          <div class="flex justify-between items-start mb-6">
            <div class="p-4 bg-indigo-50 dark:bg-indigo-900/30 rounded-2xl text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
              <Package class="w-6 h-6" />
            </div>
            <div class="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity relative z-20">
              <button @click="openEditModal(item)" class="p-2 text-slate-400 dark:text-slate-500 hover:text-blue-600 dark:hover:text-blue-400 transition-all"><Edit2 class="w-4 h-4" /></button>
              <button @click="selectedItem = item; isDeleteModalOpen = true" class="p-2 text-slate-400 dark:text-slate-500 hover:text-red-600 dark:hover:text-red-400 transition-all"><Trash2 class="w-4 h-4" /></button>
            </div>
          </div>

          <div class="mb-6">
            <h3 class="text-xl font-black text-slate-800 dark:text-slate-100 group-hover:text-indigo-600 transition-colors">{{ item.name }}</h3>
            <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ item.barcode }}</p>
          </div>

          <div class="mt-auto space-y-4">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">{{ t('inventory.stockLevel') }}</p>
                <div class="flex items-center gap-2">
                  <span class="text-2xl font-black" :class="item.stock <= item.min_stock_alert ? 'text-red-600' : 'text-slate-800 dark:text-slate-100'">{{ item.stock }}</span>
                  <span class="text-xs font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ item.unit }}</span>
                </div>
              </div>
              <div v-if="item.stock <= item.min_stock_alert" class="bg-red-50 dark:bg-red-900/30 p-2 rounded-xl text-red-600 dark:text-red-400">
                <AlertCircle class="w-5 h-5" />
              </div>
            </div>

            <button @click="openStockModal(item)" class="w-full py-4 bg-slate-50 dark:bg-slate-800 text-slate-700 dark:text-slate-200 rounded-2xl font-black hover:bg-indigo-600 hover:text-white transition-all flex items-center justify-center gap-2">
              <Activity class="w-4 h-4" />
              {{ t('inventory.updateStock') }}
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="items.length === 0 && !loading" class="py-32 text-center bg-white dark:bg-slate-900 rounded-[3rem] border-4 border-dashed border-slate-100 dark:border-slate-800">
        <div class="max-w-xs mx-auto space-y-4 opacity-40 dark:opacity-20">
          <Package class="w-20 h-20 mx-auto text-slate-400 dark:text-slate-600" />
          <p class="text-xl font-black text-slate-500 dark:text-slate-400 uppercase">{{ t('common.noData') }}</p>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <FormModal :show="isAddModalOpen || isEditModalOpen" :title="isEditModalOpen ? t('inventory.itemDetails') : t('inventory.newIngredient')" :icon="Package" :loading="actionLoading" @submit="isEditModalOpen ? updateItem() : createItem()" @cancel="isAddModalOpen = false; isEditModalOpen = false">
      <div class="space-y-6">
        <div class="space-y-2"><label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('inventory.ingredientName') }}</label><input v-model="itemForm.name" type="text" required class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100" /></div>
        <div class="grid grid-cols-2 gap-4">
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('inventory.barcodeSku') }}</label>
            <div class="relative">
              <input v-model="itemForm.barcode" type="text" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-12 dark:text-slate-100" />
              <Barcode class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300 dark:text-slate-600" />
            </div>
          </div>
          <div class="space-y-2"><label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('common.unit') }}</label><select v-model="itemForm.unit" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 outline-none appearance-none dark:text-slate-100"><option>KG</option><option>GR</option><option>Piece</option><option>Litre</option><option>ML</option></select></div>
        </div>
        <div class="grid grid-cols-3 gap-4">
          <div class="space-y-2"><label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('inventory.cost') }}</label><input v-model.number="itemForm.cost_price" type="number" step="0.01" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100" /></div>
          <div v-if="!isEditModalOpen" class="space-y-2"><label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('inventory.initialStock') }}</label><input v-model.number="itemForm.stock" type="number" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100" /></div>
          <div class="space-y-2" :class="isEditModalOpen ? 'col-span-2' : ''"><label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('inventory.minAlert') }}</label><input v-model.number="itemForm.min_stock_alert" type="number" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100" /></div>
        </div>
      </div>
    </FormModal>

    <FormModal :show="isStockModalOpen" :title="t('inventory.updateStock')" :subtitle="selectedItem?.name" :icon="Activity" :loading="actionLoading" @submit="updateStock" @cancel="isStockModalOpen = false">
      <div class="space-y-8">
        <div class="flex items-center justify-center gap-4">
          <button type="button" @click="stockAdjustment.type = 'add'" :class="['flex-1 py-4 rounded-2xl font-black transition-all border-2', stockAdjustment.type === 'add' ? 'bg-indigo-600 border-indigo-600 text-white' : 'bg-slate-50 dark:bg-slate-800 border-transparent text-slate-400 dark:text-slate-500']">Restock (+)</button>
          <button type="button" @click="stockAdjustment.type = 'sub'" :class="['flex-1 py-4 rounded-2xl font-black transition-all border-2', stockAdjustment.type === 'sub' ? 'bg-orange-600 border-orange-600 text-white' : 'bg-slate-50 dark:bg-slate-800 border-transparent text-slate-400 dark:text-slate-500']">Deduct (-)</button>
        </div>
        <div class="space-y-2">
          <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('inventory.adjustmentQty') }} ({{ selectedItem?.unit }})</label>
          <input v-model.number="stockAdjustment.quantity" type="number" step="0.1" class="w-full px-5 py-8 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-[2rem] font-black text-4xl text-center focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100" />
          <p class="text-center text-xs font-medium text-slate-400 dark:text-slate-500 mt-2">{{ t('inventory.stockNote') }}</p>
        </div>
      </div>
    </FormModal>

    <ConfirmModal :show="isDeleteModalOpen" type="danger" :title="t('common.confirm')" :message="`${t('common.confirm')} delete ${selectedItem?.name}?`" :loading="actionLoading" @confirm="deleteItem" @cancel="isDeleteModalOpen = false" />

    <ConfirmModal :show="alertConfig.show" :title="alertConfig.title" :message="alertConfig.message" :type="alertConfig.type" :showCancel="false" @confirm="alertConfig.show = false" />
  </DashboardLayout>
</template>
