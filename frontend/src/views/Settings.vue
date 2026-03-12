<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import { useConfigStore } from '../stores/config';
import { Settings, Globe, Store, Save, CheckCircle2 } from 'lucide-vue-next';

const configStore = useConfigStore();
const successMessage = ref('');

const tenantName = ref('');
const currency = ref('');
const language = ref('');

const alertConfig = ref({
  show: false,
  title: '',
  message: '',
  type: 'info' as 'info' | 'success' | 'warning' | 'danger'
});

const showAlert = (title: string, message: string, type: any = 'info') => {
  alertConfig.value = { show: true, title, message, type };
};

onMounted(async () => {
  await configStore.fetchTenantSettings();
  tenantName.value = configStore.tenantName;
  currency.value = configStore.currency;
  language.value = configStore.language;
});

const saveSettings = async () => {
  try {
    await configStore.updateTenantSettings(tenantName.value, currency.value);
    configStore.setLanguage(language.value);
    successMessage.value = "Settings saved successfully!";
    setTimeout(() => successMessage.value = '', 3000);
  } catch (error) {
    showAlert("Update Failed", "We couldn't save your tenant settings. Please try again.", "danger");
  }
};
</script>

<template>
  <DashboardLayout>
    <div class="max-w-4xl mx-auto space-y-8 animate-in fade-in duration-500">
      <div>
        <h1 class="text-3xl font-black text-slate-800 tracking-tight flex items-center gap-3">
          <Settings class="w-10 h-10 text-indigo-600" />
          {{ $t('settings.title') }}
        </h1>
      </div>

      <div v-if="successMessage" class="bg-green-50 border border-green-100 text-green-700 p-4 rounded-2xl flex items-center gap-3 font-bold animate-in slide-in-from-top-2">
        <CheckCircle2 class="w-5 h-5" />
        {{ successMessage }}
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <!-- Sidebar Info -->
        <div class="space-y-6">
          <div class="bg-white p-6 rounded-[2rem] shadow-sm border border-slate-100">
            <div class="w-12 h-12 bg-indigo-50 rounded-2xl flex items-center justify-center text-indigo-600 mb-4">
              <Globe class="w-6 h-6" />
            </div>
            <h3 class="font-black text-slate-800 mb-2">Localization</h3>
            <p class="text-xs text-slate-400 font-medium leading-relaxed">Adjust your language and currency to match your local market requirements.</p>
          </div>
          
          <div class="bg-indigo-900 p-6 rounded-[2rem] shadow-xl text-white">
            <div class="w-12 h-12 bg-white/10 rounded-2xl flex items-center justify-center text-white mb-4">
              <Store class="w-6 h-6" />
            </div>
            <h3 class="font-black mb-2">Tenant Profile</h3>
            <p class="text-xs text-indigo-200 font-medium leading-relaxed">This name will appear on your dashboard and internal reports.</p>
          </div>
        </div>

        <!-- Form Area -->
        <div class="md:col-span-2">
          <div class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 p-8 md:p-10">
            <form @submit.prevent="saveSettings" class="space-y-8">
              <div class="space-y-2">
                <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('settings.tenantName') }}</label>
                <input v-model="tenantName" type="text" required class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none" />
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('settings.language') }}</label>
                  <select v-model="language" class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none appearance-none">
                    <option value="en">English (US)</option>
                    <option value="id">Bahasa Indonesia</option>
                    <option value="ja">日本語 (Japanese)</option>
                  </select>
                </div>

                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('settings.currency') }}</label>
                  <select v-model="currency" class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none appearance-none">
                    <option value="USD">USD ($)</option>
                    <option value="IDR">IDR (Rp)</option>
                    <option value="JPY">JPY (¥)</option>
                  </select>
                </div>
              </div>

              <div class="pt-6">
                <button type="submit" class="w-full py-5 bg-indigo-600 text-white rounded-3xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all flex items-center justify-center gap-3">
                  <Save class="w-6 h-6" />
                  {{ $t('settings.saveSettings') }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Alert Modal -->
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
