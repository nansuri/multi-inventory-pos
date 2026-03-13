<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import api from '../api/axios';
import { Grid, Plus, Trash2, Edit2, Users } from 'lucide-vue-next';
import { useI18n } from 'vue-i18n';
import { useAuthStore } from '../stores/auth';

const { t } = useI18n();
const authStore = useAuthStore();
const tables = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isDeleteModalOpen = ref(false);
const selectedTable = ref<any>(null);

const tableForm = ref({
  id: 0,
  table_number: '',
  capacity: 2,
  status: 'available'
});

const fetchTables = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/tables');
    tables.value = response.data;
  } catch (error) {
    console.error('Failed to fetch tables', error);
  } finally {
    loading.value = false;
  }
};

const createTable = async () => {
  actionLoading.value = true;
  try {
    await api.post('/api/tables', tableForm.value);
    isAddModalOpen.value = false;
    resetForm();
    await fetchTables();
  } catch (error) {
    console.error('Failed to create table', error);
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (table: any) => {
  tableForm.value = { ...table };
  isEditModalOpen.value = true;
};

const updateTable = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/tables/${tableForm.value.id}`, tableForm.value);
    isEditModalOpen.value = false;
    resetForm();
    await fetchTables();
  } catch (error) {
    console.error('Failed to update table', error);
  } finally {
    actionLoading.value = false;
  }
};

const deleteTable = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/tables/${selectedTable.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchTables();
  } catch (error) {
    console.error('Failed to delete table', error);
  } finally {
    actionLoading.value = false;
  }
};

const resetForm = () => {
  tableForm.value = { id: 0, table_number: '', capacity: 2, status: 'available' };
};

onMounted(fetchTables);

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, fetchTables);
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <Grid class="w-6 h-6 text-indigo-600" />
          Table Management
        </h1>
        <p class="text-slate-500 dark:text-slate-400 font-medium text-xs mt-0.5">Configure your restaurant's floor plan and seating.</p>
      </div>
      <button @click="isAddModalOpen = true" class="bg-indigo-600 text-white px-6 py-2.5 rounded-xl text-sm font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2">
        <Plus class="w-4 h-4" />
        {{ t('pos.newTable') }}
      </button>
    </div>

    <!-- Tables Table -->
    <div class="bg-white dark:bg-slate-900 rounded-3xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden text-sm">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('pos.tableIdentifier') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('pos.seatingCapacity') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('common.status') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <tr v-for="table in tables" :key="table.id" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors">
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <div class="w-8 h-8 rounded-lg bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 flex items-center justify-center font-black text-xs">
                  {{ table.table_number }}
                </div>
                <span class="font-bold text-slate-800 dark:text-slate-100">{{ t('common.table') }} {{ table.table_number }}</span>
              </div>
            </td>
            <td class="px-6 py-4 text-center">
              <div class="flex items-center justify-center gap-1.5 text-slate-600 dark:text-slate-400">
                <Users class="w-3.5 h-3.5" />
                <span class="font-bold">{{ table.capacity }} pax</span>
              </div>
            </td>
            <td class="px-6 py-4 text-center">
              <span :class="['inline-flex items-center px-2.5 py-0.5 rounded-full text-[9px] font-black uppercase tracking-wider', 
                table.status === 'available' ? 'bg-green-50 text-green-600 dark:bg-green-900/20' : 'bg-orange-50 text-orange-600 dark:bg-orange-950/20']">
                {{ table.status }}
              </span>
            </td>
            <td class="px-6 py-4 text-right">
              <div class="flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                <button @click="openEditModal(table)" class="p-1.5 text-slate-400 hover:text-blue-600 transition-all">
                  <Edit2 class="w-4 h-4" />
                </button>
                <button @click="selectedTable = table; isDeleteModalOpen = true" class="p-1.5 text-slate-400 hover:text-red-600 transition-all">
                  <Trash2 class="w-4 h-4" />
                </button>
              </div>
            </td>
          </tr>
          <tr v-if="tables.length === 0 && !loading">
            <td colspan="4" class="px-6 py-20 text-center text-slate-400 italic">No tables configured for this branch.</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modals -->
    <FormModal 
      :show="isAddModalOpen || isEditModalOpen" 
      :title="isEditModalOpen ? 'Edit Table' : t('pos.newTable')" 
      :icon="Grid" 
      :loading="actionLoading" 
      @submit="isEditModalOpen ? updateTable() : createTable()" 
      @cancel="isAddModalOpen = false; isEditModalOpen = false; resetForm()"
    >
      <div class="space-y-4">
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('pos.tableIdentifier') }}</label>
          <input v-model="tableForm.table_number" type="text" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100 text-sm" placeholder="e.g. A1, 12, VIP-1" />
        </div>
        
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('pos.seatingCapacity') }}</label>
          <div class="relative">
            <input v-model.number="tableForm.capacity" type="number" min="1" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-10 dark:text-slate-100 text-sm" />
            <Users class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 dark:text-slate-600" />
          </div>
        </div>

        <div v-if="isEditModalOpen" class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('common.status') }}</label>
          <select v-model="tableForm.status" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none appearance-none dark:text-slate-100 text-sm">
            <option value="available">Available</option>
            <option value="occupied">Occupied</option>
          </select>
        </div>
      </div>
    </FormModal>

    <ConfirmModal 
      :show="isDeleteModalOpen" 
      type="danger" 
      title="Delete Table" 
      :message="`Are you sure you want to delete Table ${selectedTable?.table_number}? This may affect active orders.`" 
      :loading="actionLoading" 
      @confirm="deleteTable" 
      @cancel="isDeleteModalOpen = false" 
    />
  </div>
</template>
