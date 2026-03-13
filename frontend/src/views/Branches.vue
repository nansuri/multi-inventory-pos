<script setup lang="ts">
import { ref, onMounted } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import api from '../api/axios';
import { MapPin, Plus, Trash2, Edit2, Phone, Globe } from 'lucide-vue-next';
import { useI18n } from 'vue-i18n';
import { useAuthStore } from '../stores/auth';

const { t } = useI18n();
const authStore = useAuthStore();
const branches = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isDeleteModalOpen = ref(false);
const selectedBranch = ref<any>(null);

const branchForm = ref({
  id: 0,
  name: '',
  address: '',
  phone: '',
  is_active: true
});

const fetchBranches = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/branches');
    branches.value = response.data;
    // Also update auth store branches list
    authStore.branches = response.data;
  } catch (error) {
    console.error('Failed to fetch branches', error);
  } finally {
    loading.value = false;
  }
};

const createBranch = async () => {
  actionLoading.value = true;
  try {
    await api.post('/api/branches', branchForm.value);
    isAddModalOpen.value = false;
    resetForm();
    await fetchBranches();
  } catch (error) {
    console.error('Failed to create branch', error);
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (branch: any) => {
  branchForm.value = { ...branch };
  isEditModalOpen.value = true;
};

const updateBranch = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/branches/${branchForm.value.id}`, branchForm.value);
    isEditModalOpen.value = false;
    resetForm();
    await fetchBranches();
  } catch (error) {
    console.error('Failed to update branch', error);
  } finally {
    actionLoading.value = false;
  }
};

const deleteBranch = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/branches/${selectedBranch.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchBranches();
  } catch (error) {
    console.error('Failed to delete branch', error);
  } finally {
    actionLoading.value = false;
  }
};

const resetForm = () => {
  branchForm.value = { id: 0, name: '', address: '', phone: '', is_active: true };
};

onMounted(fetchBranches);
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <MapPin class="w-6 h-6 text-indigo-600" />
          {{ t('common.branches') }}
        </h1>
        <p class="text-slate-500 dark:text-slate-400 font-medium text-xs mt-0.5">Manage your restaurant locations and branches.</p>
      </div>
      <button @click="isAddModalOpen = true" class="bg-indigo-600 text-white px-6 py-2.5 rounded-xl text-sm font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2">
        <Plus class="w-4 h-4" />
        Add Branch
      </button>
    </div>

    <!-- Branch Grid -->
    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <div v-for="i in 4" :key="i" class="h-40 bg-slate-100 dark:bg-slate-800 animate-pulse rounded-2xl"></div>
    </div>

    <div v-else-if="branches.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <div v-for="branch in branches" :key="branch.id" class="bg-white dark:bg-slate-900 p-6 rounded-[2rem] shadow-sm border border-slate-200 dark:border-slate-800 hover:shadow-lg transition-all duration-300 flex flex-col group relative overflow-hidden">
        <div class="flex justify-between items-start mb-4">
          <div class="p-3 bg-indigo-50 dark:bg-indigo-900/30 rounded-xl text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
            <MapPin class="w-5 h-5" />
          </div>
          <div class="flex gap-0.5 opacity-0 group-hover:opacity-100 transition-opacity">
            <button @click="openEditModal(branch)" class="p-1.5 text-slate-400 hover:text-blue-600 transition-all"><Edit2 class="w-4 h-4" /></button>
            <button @click="selectedBranch = branch; isDeleteModalOpen = true" class="p-1.5 text-slate-400 hover:text-red-600 transition-all"><Trash2 class="w-4 h-4" /></button>
          </div>
        </div>

        <div class="space-y-3">
          <div>
            <h3 class="text-base font-black text-slate-800 dark:text-slate-100 leading-tight">{{ branch.name }}</h3>
            <div v-if="authStore.selectedBranchID === branch.id" class="inline-flex items-center gap-1 mt-1 px-2 py-0.5 bg-green-50 dark:bg-green-900/30 text-green-600 dark:text-green-400 text-[8px] font-black uppercase tracking-widest rounded-full">
              Current Branch
            </div>
          </div>

          <div class="space-y-1.5">
            <div class="flex items-start gap-2 text-slate-500 dark:text-slate-400">
              <Globe class="w-3.5 h-3.5 mt-0.5 flex-shrink-0" />
              <p class="text-[11px] font-medium leading-relaxed line-clamp-2">{{ branch.address || 'No address provided' }}</p>
            </div>
            <div class="flex items-center gap-2 text-slate-500 dark:text-slate-400">
              <Phone class="w-3.5 h-3.5 flex-shrink-0" />
              <p class="text-[11px] font-medium">{{ branch.phone || 'No phone provided' }}</p>
            </div>
          </div>
        </div>

        <div class="mt-6 pt-4 border-t border-slate-50 dark:border-slate-800 flex justify-between items-center">
          <span :class="['text-[9px] font-black uppercase tracking-widest px-2 py-0.5 rounded-full', branch.is_active ? 'bg-green-50 text-green-600 dark:bg-green-900/20' : 'bg-slate-100 text-slate-400 dark:bg-slate-800']">
            {{ branch.is_active ? 'Active' : 'Inactive' }}
          </span>
          <button 
            v-if="authStore.selectedBranchID !== branch.id"
            @click="authStore.setBranch(branch.id)" 
            class="text-[9px] font-black uppercase tracking-widest text-indigo-600 dark:text-indigo-400 hover:underline"
          >
            Switch
          </button>
        </div>
      </div>
    </div>

    <div v-else class="py-20 text-center bg-white dark:bg-slate-900 rounded-[2.5rem] border-2 border-dashed border-slate-200 dark:border-slate-800">
      <div class="max-w-xs mx-auto space-y-3 opacity-40">
        <MapPin class="w-16 h-16 mx-auto text-slate-400" />
        <p class="text-sm font-black text-slate-500 uppercase tracking-widest">No branches found</p>
        <button @click="isAddModalOpen = true" class="text-indigo-600 font-bold hover:underline text-xs">Add your first branch</button>
      </div>
    </div>

    <!-- Modals -->
    <FormModal 
      :show="isAddModalOpen || isEditModalOpen" 
      :title="isEditModalOpen ? 'Edit Branch' : 'New Branch'" 
      :icon="MapPin" 
      :loading="actionLoading" 
      @submit="isEditModalOpen ? updateBranch() : createBranch()" 
      @cancel="isAddModalOpen = false; isEditModalOpen = false; resetForm()"
    >
      <div class="space-y-4">
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">Branch Name</label>
          <input v-model="branchForm.name" type="text" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100 text-sm" placeholder="Main Street Outlet" />
        </div>
        
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">Phone Number</label>
          <div class="relative">
            <input v-model="branchForm.phone" type="text" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-10 dark:text-slate-100 text-sm" placeholder="+1 234 567 890" />
            <Phone class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 dark:text-slate-600" />
          </div>
        </div>

        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">Full Address</label>
          <textarea v-model="branchForm.address" rows="3" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100 text-sm" placeholder="Street name, City, Zip code"></textarea>
        </div>

        <div class="flex items-center gap-3 px-1">
          <input type="checkbox" v-model="branchForm.is_active" id="is_active" class="w-4 h-4 rounded border-slate-300 text-indigo-600 focus:ring-indigo-600" />
          <label for="is_active" class="text-xs font-bold text-slate-600 dark:text-slate-400 cursor-pointer">This branch is active and operational</label>
        </div>
      </div>
    </FormModal>

    <ConfirmModal 
      :show="isDeleteModalOpen" 
      type="danger" 
      title="Delete Branch" 
      :message="`Are you sure you want to delete ${selectedBranch?.name}? This action cannot be undone.`" 
      :loading="actionLoading" 
      @confirm="deleteBranch" 
      @cancel="isDeleteModalOpen = false" 
    />
  </div>
</template>
