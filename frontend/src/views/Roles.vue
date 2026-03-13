<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import api from '../api/axios';
import { useI18n } from 'vue-i18n';
import { Users, ShieldCheck, Plus, Trash2, Edit2, Lock, Check } from 'lucide-vue-next';

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

const permissionModules = [
  { key: 'dashboard', name: 'Dashboard' },
  { key: 'inventory', name: 'Inventory' },
  { key: 'recipes', name: 'Recipes' },
  { key: 'production', name: 'Production Hall' },
  { key: 'production_log', name: 'Production Log' },
  { key: 'pos_order', name: 'New Order' },
  { key: 'pos_payment', name: 'Payments' },
  { key: 'order_history', name: 'Order History' },
  { key: 'employees', name: 'Employee Management' },
  { key: 'settings', name: 'Settings' }
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
    await fetchRoles();
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (role: any) => {
  roleForm.value = { 
    id: role.id, 
    name: role.name, 
    permissions: [...(role.permissions || [])] 
  };
  isEditModalOpen.value = true;
};

const updateRole = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/roles/${roleForm.value.id}`, roleForm.value);
    isEditModalOpen.value = false;
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

const togglePermission = (key: string) => {
  const index = roleForm.value.permissions.indexOf(key);
  if (index > -1) {
    roleForm.value.permissions.splice(index, 1);
  } else {
    roleForm.value.permissions.push(key);
  }
};

onMounted(fetchRoles);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8 animate-in fade-in duration-500">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <ShieldCheck class="w-8 h-8 text-indigo-600" />
            {{ t('roles.title') }}
          </h1>
          <p class="text-slate-500 font-medium mt-1">{{ t('roles.desc') }}</p>
        </div>
        <button 
          @click="isAddModalOpen = true; roleForm = { id: 0, name: '', permissions: [] }"
          class="bg-indigo-600 text-white px-8 py-4 rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all flex items-center gap-2"
        >
          <Plus class="w-5 h-5" />
          {{ t('roles.createRole') }}
        </button>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div v-for="role in roles" :key="role.id" class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 p-8 flex flex-col group hover:shadow-xl transition-all duration-300 overflow-hidden relative">
          <div class="absolute top-0 right-0 p-10 opacity-5 group-hover:scale-110 transition-transform">
            <Lock class="w-32 h-32 text-indigo-600" />
          </div>

          <div class="flex justify-between items-start mb-8">
            <div class="p-4 bg-indigo-50 rounded-2xl text-indigo-600">
              <Users class="w-6 h-6" />
            </div>
            <div class="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
              <button @click="openEditModal(role)" class="p-2 text-slate-400 hover:text-blue-600 transition-all"><Edit2 class="w-4 h-4" /></button>
              <button @click="selectedRole = role; isDeleteModalOpen = true" class="p-2 text-slate-400 hover:text-red-600 transition-all"><Trash2 class="w-4 h-4" /></button>
            </div>
          </div>

          <h3 class="text-xl font-black text-slate-800 mb-6">{{ role.name }}</h3>
          
          <div class="flex flex-wrap gap-2 mb-8 relative z-10">
            <span v-for="perm in role.permissions" :key="perm" class="px-2 py-1 bg-slate-100 text-slate-500 text-[9px] font-black uppercase rounded-lg border border-slate-50">
              {{ perm.replace('_', ' ') }}
            </span>
            <span v-if="!role.permissions?.length" class="text-xs text-slate-400 italic">No permissions assigned.</span>
          </div>

          <button @click="openEditModal(role)" class="mt-auto w-full py-4 bg-slate-50 text-slate-700 rounded-2xl font-black hover:bg-indigo-600 hover:text-white transition-all">
            {{ t('roles.managePermissions') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <FormModal 
      :show="isAddModalOpen || isEditModalOpen" 
      :title="isEditModalOpen ? t('common.edit') + ' Role' : t('roles.createRole')" 
      :icon="ShieldCheck" 
      :loading="actionLoading"
      maxWidth="max-w-2xl"
      @submit="isEditModalOpen ? updateRole() : createRole()" 
      @cancel="isAddModalOpen = false; isEditModalOpen = false"
    >
      <div class="space-y-8">
        <div class="space-y-2">
          <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('roles.roleName') }}</label>
          <input v-model="roleForm.name" type="text" required placeholder="e.g. Head Chef, Junior Waiter" class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none" />
        </div>

        <div class="space-y-4">
          <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('roles.pagePermissions') }}</label>
          <div class="grid grid-cols-2 gap-3">
            <button 
              v-for="mod in permissionModules" 
              :key="mod.key"
              type="button"
              @click="togglePermission(mod.key)"
              :class="[
                'flex items-center justify-between px-4 py-3 rounded-xl border-2 transition-all text-sm font-bold',
                roleForm.permissions.includes(mod.key) 
                  ? 'bg-indigo-50 border-indigo-600 text-indigo-700' 
                  : 'bg-white border-slate-100 text-slate-400 hover:border-slate-200'
              ]"
            >
              {{ mod.name }}
              <div v-if="roleForm.permissions.includes(mod.key)" class="w-4 h-4 bg-indigo-600 rounded-full flex items-center justify-center">
                <Check class="w-2.5 h-2.5 text-white" />
              </div>
            </button>
          </div>
        </div>
      </div>
    </FormModal>

    <ConfirmModal 
      :show="isDeleteModalOpen"
      type="danger"
      title="Delete Role?"
      :message="`Are you sure you want to delete the role '${selectedRole?.name}'? Users assigned to this role will lose access.`"
      :loading="actionLoading"
      @confirm="deleteRole"
      @cancel="isDeleteModalOpen = false"
    />
  </DashboardLayout>
</template>
