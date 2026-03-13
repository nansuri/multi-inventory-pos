<script setup lang="ts">
import { ref, onMounted } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import api from '../api/axios';
import { Shield, Plus, Trash2, Edit2, ShieldCheck, Check } from 'lucide-vue-next';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const roles = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isDeleteModalOpen = ref(false);
const selectedRole = ref<any>(null);

const roleForm = ref({
  id: 0,
  name: '',
  permissions: [] as string[]
});

const availablePermissions = [
  { id: 'dashboard', name: 'Dashboard Access' },
  { id: 'inventory', name: 'Manage Inventory' },
  { id: 'recipes', name: 'Manage Recipes' },
  { id: 'production', name: 'Manage Production' },
  { id: 'production_log', name: 'View Production Logs' },
  { id: 'pos_order', name: 'Create Orders' },
  { id: 'pos_payment', name: 'Process Payments' },
  { id: 'order_history', name: 'View Order History' },
  { id: 'employees', name: 'Manage Employees' },
  { id: 'branches', name: 'Manage Branches' },
  { id: 'settings', name: 'System Settings' }
];

const fetchRoles = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/roles');
    roles.value = response.data;
  } catch (error) {
    console.error('Failed to fetch roles', error);
  } finally {
    loading.value = false;
  }
};

const createRole = async () => {
  actionLoading.value = true;
  try {
    await api.post('/api/roles', roleForm.value);
    isAddModalOpen.value = false;
    resetForm();
    await fetchRoles();
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (role: any) => {
  roleForm.value = { 
    id: role.id, 
    name: role.name, 
    permissions: Array.isArray(role.permissions) ? [...role.permissions] : [] 
  };
  isEditModalOpen.value = true;
};

const updateRole = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/roles/${roleForm.value.id}`, roleForm.value);
    isEditModalOpen.value = false;
    resetForm();
    await fetchRoles();
  } finally {
    actionLoading.value = false;
  }
};

const deleteRole = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/roles/${selectedRole.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchRoles();
  } finally {
    actionLoading.value = false;
  }
};

const togglePermission = (permId: string) => {
  const index = roleForm.value.permissions.indexOf(permId);
  if (index === -1) {
    roleForm.value.permissions.push(permId);
  } else {
    roleForm.value.permissions.splice(index, 1);
  }
};

const resetForm = () => {
  roleForm.value = { id: 0, name: '', permissions: [] };
};

onMounted(fetchRoles);
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <Shield class="w-6 h-6 text-indigo-600" />
          {{ t('roles.title') }}
        </h1>
        <p class="text-slate-500 dark:text-slate-400 font-medium text-xs mt-0.5">{{ t('roles.desc') }}</p>
      </div>
      <button @click="isAddModalOpen = true" class="bg-indigo-600 text-white px-6 py-2.5 rounded-xl text-sm font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2">
        <Plus class="w-4 h-4" />
        {{ t('roles.createRole') }}
      </button>
    </div>

    <!-- Role List -->
    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="i in 3" :key="i" class="h-48 bg-slate-100 dark:bg-slate-800 animate-pulse rounded-2xl"></div>
    </div>

    <div v-else-if="roles.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <div v-for="role in roles" :key="role.id" class="bg-white dark:bg-slate-900 p-6 rounded-[2rem] shadow-sm border border-slate-200 dark:border-slate-800 hover:shadow-lg transition-all duration-300 flex flex-col group relative overflow-hidden">
        <div class="flex justify-between items-start mb-4 relative z-10">
          <div class="p-3 bg-indigo-50 dark:bg-indigo-900/30 rounded-xl text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
            <ShieldCheck class="w-5 h-5" />
          </div>
          <div class="flex gap-0.5 opacity-0 group-hover:opacity-100 transition-opacity">
            <button @click="openEditModal(role)" class="p-1.5 text-slate-400 hover:text-blue-600 transition-all"><Edit2 class="w-4 h-4" /></button>
            <button @click="selectedRole = role; isDeleteModalOpen = true" class="p-1.5 text-slate-400 hover:text-red-600 transition-all"><Trash2 class="w-4 h-4" /></button>
          </div>
        </div>

        <div class="mb-4 relative z-10">
          <h3 class="text-base font-black text-slate-800 dark:text-slate-100 leading-tight">{{ role.name }}</h3>
          <p class="text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mt-1">{{ (role.permissions || []).length }} permissions assigned</p>
        </div>

        <div class="flex flex-wrap gap-1 mt-auto relative z-10">
          <span v-for="perm in (role.permissions || []).slice(0, 3)" :key="perm" class="px-1.5 py-0.5 bg-slate-50 dark:bg-slate-800 text-slate-500 dark:text-slate-400 text-[8px] font-black uppercase tracking-wider rounded">
            {{ perm.replace('_', ' ') }}
          </span>
          <span v-if="role.permissions?.length > 3" class="px-1.5 py-0.5 text-slate-400 text-[8px] font-bold italic">
            +{{ role.permissions.length - 3 }} more
          </span>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <FormModal 
      :show="isAddModalOpen || isEditModalOpen" 
      :title="isEditModalOpen ? 'Edit Role' : t('roles.createRole')" 
      :icon="Shield" 
      :loading="actionLoading" 
      maxWidth="max-w-2xl"
      @submit="isEditModalOpen ? updateRole() : createRole()" 
      @cancel="isAddModalOpen = false; isEditModalOpen = false; resetForm()"
    >
      <div class="space-y-6">
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('roles.roleName') }}</label>
          <input v-model="roleForm.name" type="text" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100 text-sm" placeholder="e.g. Store Manager" />
        </div>
        
        <div class="space-y-3">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('roles.pagePermissions') }}</label>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-2 max-h-[40vh] overflow-y-auto pr-2 custom-scrollbar">
            <button 
              v-for="perm in availablePermissions" 
              :key="perm.id"
              type="button"
              @click="togglePermission(perm.id)"
              :class="[
                'flex items-center justify-between p-3 rounded-xl border-2 transition-all text-left',
                roleForm.permissions.includes(perm.id) 
                  ? 'bg-indigo-50 dark:bg-indigo-900/30 border-indigo-200 dark:border-indigo-800 text-indigo-700 dark:text-indigo-300' 
                  : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-500 dark:text-slate-400 hover:border-slate-200 dark:hover:border-slate-700'
              ]"
            >
              <span class="text-[11px] font-black uppercase tracking-tight">{{ perm.name }}</span>
              <div v-if="roleForm.permissions.includes(perm.id)" class="w-4 h-4 bg-indigo-600 rounded-full flex items-center justify-center text-white">
                <Check class="w-2.5 h-2.5" />
              </div>
            </button>
          </div>
        </div>
      </div>
    </FormModal>

    <ConfirmModal 
      :show="isDeleteModalOpen" 
      type="danger" 
      :title="t('common.confirm')" 
      :message="`Delete role ${selectedRole?.name}? This may affect users assigned to this role.`" 
      :loading="actionLoading" 
      @confirm="deleteRole" 
      @cancel="isDeleteModalOpen = false" 
    />
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-200);
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-800);
}
</style>
