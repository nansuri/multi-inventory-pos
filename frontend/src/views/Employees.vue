<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import api from '../api/axios';
import { useI18n } from 'vue-i18n';
import { Users, UserPlus, Trash2, Mail, Shield, Key } from 'lucide-vue-next';

const { t } = useI18n();
const employees = ref<any[]>([]);
const roles = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);

const isAddModalOpen = ref(false);
const isDeleteModalOpen = ref(false);
const selectedEmployee = ref<any>(null);

const employeeForm = ref({
  username: '',
  password: '',
  role_id: null as number | null
});

const fetchEmployees = async () => {
  loading.value = true;
  try {
    const [empRes, rolesRes] = await Promise.all([
      api.get('/api/employees'),
      api.get('/api/roles')
    ]);
    employees.value = empRes.data;
    roles.value = rolesRes.data;
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
    employeeForm.value = { username: '', password: '', role_id: null };
    await fetchEmployees();
  } catch (error: any) {
    console.error('Failed to create employee', error);
  } finally {
    actionLoading.value = false;
  }
};

const deleteEmployee = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/employees/${selectedEmployee.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchEmployees();
  } finally {
    actionLoading.value = false;
  }
};

onMounted(fetchEmployees);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8 animate-in fade-in duration-500">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-2xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-3">
            <Users class="w-8 h-8 text-indigo-600" />
            {{ t('employees.title') }}
          </h1>
          <p class="text-slate-500 dark:text-slate-400 font-medium mt-1">{{ t('employees.desc') }}</p>
        </div>
        <div class="flex gap-3">
          <router-link to="/roles" class="bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-700 dark:text-slate-200 px-6 py-4 rounded-2xl font-bold hover:bg-slate-50 dark:hover:bg-slate-800 transition-all shadow-sm flex items-center gap-2">
            <Shield class="w-5 h-5 text-slate-400 dark:text-slate-500" />
            {{ t('employees.manageRoles') }}
          </router-link>
          <button 
            @click="isAddModalOpen = true"
            class="bg-indigo-600 text-white px-8 py-4 rounded-2xl font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/30 hover:bg-indigo-700 transition-all flex items-center gap-2"
          >
            <UserPlus class="w-5 h-5" />
            {{ t('employees.addEmployee') }}
          </button>
        </div>
      </div>

      <div class="bg-white dark:bg-slate-900 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('employees.username') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('employees.assignedRole') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="emp in employees" :key="emp.id" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors">
              <td class="px-8 py-6">
                <div class="flex items-center gap-3">
                  <div class="w-10 h-10 rounded-xl bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 flex items-center justify-center font-black">
                    {{ emp.username.charAt(0).toUpperCase() }}
                  </div>
                  <span class="font-bold text-slate-800 dark:text-slate-100">{{ emp.username }}</span>
                </div>
              </td>
              <td class="px-8 py-6">
                <span class="px-3 py-1 bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 text-[10px] font-black uppercase rounded-lg border border-indigo-100 dark:border-indigo-900/50">
                  {{ emp.role?.name || 'No Role' }}
                </span>
              </td>
              <td class="px-8 py-6 text-right">
                <button @click="selectedEmployee = emp; isDeleteModalOpen = true" class="p-2 text-slate-400 dark:text-slate-500 hover:text-red-600 dark:hover:text-red-400 hover:bg-red-50 dark:hover:bg-red-950/30 rounded-xl transition-all">
                  <Trash2 class="w-5 h-5" />
                </button>
              </td>
            </tr>
            <tr v-if="employees.length === 0 && !loading">
              <td colspan="3" class="px-8 py-20 text-center text-slate-400 dark:text-slate-600 italic font-medium">{{ t('common.noData') }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modals -->
    <FormModal 
      :show="isAddModalOpen" 
      :title="t('employees.createStaff')" 
      :subtitle="t('employees.newEmployee')"
      :icon="UserPlus" 
      :loading="actionLoading"
      @submit="createEmployee" 
      @cancel="isAddModalOpen = false"
    >
      <div class="space-y-6">
        <div class="space-y-2">
          <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('employees.username') }}</label>
          <div class="relative">
            <input v-model="employeeForm.username" type="text" required class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-12 dark:text-slate-100" />
            <Mail class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300 dark:text-slate-600" />
          </div>
        </div>
        <div class="space-y-2">
          <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('employees.tempPassword') }}</label>
          <div class="relative">
            <input v-model="employeeForm.password" type="password" required class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-12 dark:text-slate-100" />
            <Key class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300 dark:text-slate-600" />
          </div>
        </div>
        <div class="space-y-2">
          <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('employees.assignRole') }}</label>
          <select v-model="employeeForm.role_id" required class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none appearance-none dark:text-slate-100">
            <option :value="null" disabled>{{ t('employees.assignRole') }}...</option>
            <option v-for="role in roles" :key="role.id" :value="role.id">{{ role.name }}</option>
          </select>
        </div>
      </div>
    </FormModal>

    <ConfirmModal 
      :show="isDeleteModalOpen"
      type="danger"
      :title="t('employees.removeEmployee')"
      :message="`${t('common.confirm')} delete ${selectedEmployee?.username}?`"
      :loading="actionLoading"
      @confirm="deleteEmployee"
      @cancel="isDeleteModalOpen = false"
    />
  </DashboardLayout>
</template>
