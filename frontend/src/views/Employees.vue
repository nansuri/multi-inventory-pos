<script setup lang="ts">
import { ref, onMounted } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import api from '../api/axios';
import { Users, UserPlus, Trash2, Mail, Key, MapPin, Edit2, CheckCircle2, XCircle } from 'lucide-vue-next';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const employees = ref<any[]>([]);
const roles = ref<any[]>([]);
const branches = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isDeleteModalOpen = ref(false);
const selectedEmployee = ref<any>(null);

const employeeForm = ref({
  id: 0,
  username: '',
  password: '',
  role_id: null as number | null,
  branch_id: null as number | null,
  is_active: true
});

const fetchEmployees = async () => {
  loading.value = true;
  try {
    const [empRes, rolesRes, branchesRes] = await Promise.all([
      api.get('/api/employees'),
      api.get('/api/roles'),
      api.get('/api/branches')
    ]);
    employees.value = empRes.data;
    roles.value = rolesRes.data;
    branches.value = branchesRes.data;
  } catch (error) {
    console.error('Failed to fetch data', error);
  } finally {
    loading.value = false;
  }
};

const createEmployee = async () => {
  if (!employeeForm.value.role_id) return;
  
  actionLoading.value = true;
  try {
    await api.post('/api/employees', employeeForm.value);
    isAddModalOpen.value = false;
    resetForm();
    await fetchEmployees();
  } catch (error: any) {
    console.error('Failed to create employee', error);
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (emp: any) => {
  selectedEmployee.value = emp;
  employeeForm.value = {
    id: emp.id,
    username: emp.username,
    password: '', // Keep empty unless changing
    role_id: emp.role_id,
    branch_id: emp.branch_id,
    is_active: emp.is_active
  };
  isEditModalOpen.value = true;
};

const updateEmployee = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/employees/${employeeForm.value.id}`, employeeForm.value);
    isEditModalOpen.value = false;
    resetForm();
    await fetchEmployees();
  } catch (error) {
    console.error('Failed to update employee', error);
  } finally {
    actionLoading.value = false;
  }
};

const toggleStatus = async (emp: any) => {
  try {
    await api.patch(`/api/employees/${emp.id}/toggle`);
    await fetchEmployees();
  } catch (error) {
    console.error('Failed to toggle status', error);
  }
};

const deleteEmployee = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/employees/${selectedEmployee.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchEmployees();
  } catch (error) {
    console.error('Failed to delete employee', error);
  } finally {
    actionLoading.value = false;
  }
};

const resetForm = () => {
  employeeForm.value = { id: 0, username: '', password: '', role_id: null, branch_id: null, is_active: true };
};

onMounted(fetchEmployees);
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <Users class="w-6 h-6 text-indigo-600" />
          {{ t('employees.title') }}
        </h1>
        <p class="text-slate-500 dark:text-slate-400 font-medium text-xs mt-0.5">{{ t('employees.desc') }}</p>
      </div>
      <button @click="isAddModalOpen = true" class="bg-indigo-600 text-white px-6 py-2.5 rounded-xl text-sm font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2">
        <UserPlus class="w-4 h-4" />
        {{ t('employees.addEmployee') }}
      </button>
    </div>

    <!-- Employee Table -->
    <div class="bg-white dark:bg-slate-900 rounded-3xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('employees.username') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('employees.assignedRole') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('common.branches') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('common.status') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <tr v-for="emp in employees" :key="emp.id" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors">
            <td class="px-6 py-3">
              <div class="flex items-center gap-3">
                <div class="w-8 h-8 rounded-lg bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 flex items-center justify-center font-black text-xs">
                  {{ emp.username.charAt(0).toUpperCase() }}
                </div>
                <div class="flex flex-col">
                  <span class="font-bold text-slate-800 dark:text-slate-100 text-sm">{{ emp.username }}</span>
                  <span v-if="emp.is_owner" class="text-[8px] font-black text-indigo-500 uppercase tracking-tighter">System Owner</span>
                </div>
              </div>
            </td>
            <td class="px-6 py-3">
              <span class="px-2 py-0.5 bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 text-[9px] font-black uppercase rounded-md border border-slate-200 dark:border-slate-700">
                {{ emp.role?.name || 'No Role' }}
              </span>
            </td>
            <td class="px-6 py-3">
              <span v-if="emp.branch" class="text-[10px] font-bold text-slate-600 dark:text-slate-400 flex items-center gap-1.5">
                <MapPin class="w-3 h-3 text-slate-400" />
                {{ emp.branch.name }}
              </span>
              <span v-else-if="emp.is_owner" class="text-[10px] font-bold text-slate-400 italic">All Branches</span>
              <span v-else class="text-[10px] font-bold text-red-400 italic">Unassigned</span>
            </td>
            <td class="px-6 py-3 text-center">
              <button 
                v-if="!emp.is_owner"
                @click="toggleStatus(emp)"
                :class="['inline-flex items-center gap-1.5 px-2 py-0.5 rounded-full text-[9px] font-black uppercase transition-all', 
                  emp.is_active ? 'bg-green-50 text-green-600 dark:bg-green-900/20' : 'bg-red-50 text-red-600 dark:bg-red-900/20']"
              >
                <component :is="emp.is_active ? CheckCircle2 : XCircle" class="w-3 h-3" />
                {{ emp.is_active ? t('common.active') : t('common.inactive') }}
              </button>
              <span v-else class="text-[9px] font-black text-green-600 uppercase tracking-widest">Active</span>
            </td>
            <td class="px-6 py-3 text-right">
              <div class="flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                <button v-if="!emp.is_owner" @click="openEditModal(emp)" class="p-1.5 text-slate-400 hover:text-blue-600 transition-all">
                  <Edit2 class="w-4 h-4" />
                </button>
                <button v-if="!emp.is_owner" @click="selectedEmployee = emp; isDeleteModalOpen = true" class="p-1.5 text-slate-400 hover:text-red-600 transition-all">
                  <Trash2 class="w-4 h-4" />
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modals -->
    <FormModal 
      :show="isAddModalOpen || isEditModalOpen" 
      :title="isEditModalOpen ? 'Edit Employee' : t('employees.newEmployee')" 
      :icon="UserPlus" 
      :loading="actionLoading" 
      @submit="isEditModalOpen ? updateEmployee() : createEmployee()" 
      @cancel="isAddModalOpen = false; isEditModalOpen = false; resetForm()"
    >
      <div class="space-y-4">
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('employees.username') }}</label>
          <div class="relative">
            <input v-model="employeeForm.username" type="text" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-10 dark:text-slate-100 text-sm" />
            <Mail class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 dark:text-slate-600" />
          </div>
        </div>
        
        <div class="space-y-1.5">
          <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ isEditModalOpen ? 'New Password (Optional)' : t('common.password') }}</label>
          <div class="relative">
            <input v-model="employeeForm.password" type="password" :required="!isEditModalOpen" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-10 dark:text-slate-100 text-sm" />
            <Key class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 dark:text-slate-600" />
          </div>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1.5">
            <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('employees.assignedRole') }}</label>
            <select v-model="employeeForm.role_id" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none appearance-none dark:text-slate-100 text-sm">
              <option :value="null" disabled>Select Role</option>
              <option v-for="role in roles" :key="role.id" :value="role.id">{{ role.name }}</option>
            </select>
          </div>
          <div class="space-y-1.5">
            <label class="block text-[9px] font-black text-slate-400 uppercase tracking-widest ml-1">Branch</label>
            <select v-model="employeeForm.branch_id" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none appearance-none dark:text-slate-100 text-sm">
              <option :value="null">All Branches</option>
              <option v-for="branch in branches" :key="branch.id" :value="branch.id">{{ branch.name }}</option>
            </select>
          </div>
        </div>

        <div v-if="isEditModalOpen" class="flex items-center gap-2 pt-2">
          <input type="checkbox" v-model="employeeForm.is_active" id="edit_active" class="w-4 h-4 rounded border-slate-300 text-indigo-600" />
          <label for="edit_active" class="text-xs font-bold text-slate-600 dark:text-slate-400">Account is active</label>
        </div>
      </div>
    </FormModal>

    <ConfirmModal 
      :show="isDeleteModalOpen" 
      type="danger" 
      :title="t('common.confirm')" 
      :message="`${t('common.confirm')} delete ${selectedEmployee?.username}?`" 
      :loading="actionLoading" 
      @confirm="deleteEmployee" 
      @cancel="isDeleteModalOpen = false" 
    />
  </div>
</template>
