<script setup lang="ts">
import { ref, onMounted } from 'vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useI18n } from 'vue-i18n';
import { Settings, Globe, Store, Save, CheckCircle } from 'lucide-vue-next';

const { t } = useI18n();
const configStore = useConfigStore();
const loading = ref(false);
const success = ref(false);

const settings = ref({
  language: configStore.language,
  currency: configStore.currency,
  tenant_name: configStore.tenantName
});

const saveSettings = async () => {
  loading.value = true;
  try {
    // Fixed: Using PATCH /api/tenant to match backend route
    await api.patch('/api/tenant', settings.value);
    
    // Update local store immediately
    configStore.setLanguage(settings.value.language);
    configStore.currency = settings.value.currency;
    configStore.tenantName = settings.value.tenant_name;
    
    success.value = true;
    setTimeout(() => success.value = false, 3000);
  } catch (error) {
    console.error('Failed to save settings', error);
  } finally {
    loading.value = false;
  }
};

onMounted(async () => {
  await configStore.fetchTenantSettings();
  settings.value = {
    language: configStore.language,
    currency: configStore.currency,
    tenant_name: configStore.tenantName
  };
});
</script>

<template>
  <div class="max-w-4xl mx-auto space-y-6 animate-in fade-in duration-500">
    <div>
      <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
        <Settings class="w-6 h-6 text-indigo-600" />
        {{ t('settings.title') }}
      </h1>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <!-- Sidebar Info -->
      <div class="space-y-4">
        <div class="p-5 bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-sm">
          <div class="w-10 h-10 bg-indigo-50 dark:bg-indigo-900/30 rounded-xl flex items-center justify-center text-indigo-600 dark:text-indigo-400 mb-3">
            <Globe class="w-5 h-5" />
          </div>
          <h3 class="font-black text-sm text-slate-800 dark:text-slate-100 mb-1.5">{{ t('settings.localization') }}</h3>
          <p class="text-[10px] text-slate-400 dark:text-slate-500 font-medium leading-relaxed">{{ t('settings.localizationDesc') }}</p>
        </div>

        <div class="p-5 bg-indigo-600 rounded-2xl shadow-lg shadow-indigo-100 dark:shadow-none text-white">
          <div class="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center text-white mb-3">
            <Store class="w-5 h-5" />
          </div>
          <h3 class="font-black text-sm mb-1.5">{{ t('settings.tenantProfile') }}</h3>
          <p class="text-[10px] text-indigo-100 dark:text-indigo-200 font-medium leading-relaxed">{{ t('settings.tenantProfileDesc') }}</p>
        </div>
      </div>

      <!-- Form Area -->
      <div class="md:col-span-2 space-y-4">
        <div class="bg-white dark:bg-slate-900 rounded-[2rem] shadow-sm border border-slate-200 dark:border-slate-800 p-8">
          <form @submit.prevent="saveSettings" class="space-y-6">
            <div class="space-y-4">
              <div class="space-y-1.5">
                <label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('settings.language') }}</label>
                <select v-model="settings.language" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none appearance-none dark:text-slate-100 text-sm">
                  <option value="en">English (US)</option>
                  <option value="id">Bahasa Indonesia</option>
                  <option value="ja">日本語 (Japanese)</option>
                </select>
              </div>

              <div class="space-y-1.5">
                <label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('settings.currency') }}</label>
                <select v-model="settings.currency" class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none appearance-none dark:text-slate-100 text-sm">
                  <option value="USD">USD ($)</option>
                  <option value="IDR">IDR (Rp)</option>
                  <option value="JPY">JPY (¥)</option>
                </select>
              </div>

              <div class="space-y-1.5">
                <label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('settings.tenantName') }}</label>
                <div class="relative">
                  <input v-model="settings.tenant_name" type="text" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-10 dark:text-slate-100 text-sm" />
                  <Store class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 dark:text-slate-600" />
                </div>
              </div>
            </div>

            <div class="pt-2">
              <button 
                type="submit" 
                :disabled="loading"
                class="w-full py-4 bg-slate-900 dark:bg-indigo-600 text-white rounded-xl font-black text-xs uppercase tracking-widest flex items-center justify-center gap-2 hover:bg-indigo-600 dark:hover:bg-indigo-700 transition-all shadow-lg group disabled:opacity-50"
              >
                <Save v-if="!loading" class="w-4 h-4 text-indigo-400 dark:text-indigo-200 group-hover:text-white transition-colors" />
                <span v-else class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
                {{ t('settings.saveSettings') }}
              </button>
              
              <div v-if="success" class="mt-3 flex items-center justify-center gap-1.5 text-green-600 dark:text-green-400 font-bold text-[10px] uppercase tracking-widest animate-in fade-in slide-in-from-top-1">
                <CheckCircle class="w-3.5 h-3.5" />
                Settings updated successfully!
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
