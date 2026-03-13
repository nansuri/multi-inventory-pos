<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useI18n } from 'vue-i18n';
import { Settings, Globe, DollarSign, Store, Save, CheckCircle } from 'lucide-vue-next';

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
    await api.put('/api/tenant/settings', settings.value);
    
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
  <DashboardLayout>
    <div class="max-w-4xl mx-auto space-y-8 animate-in fade-in duration-500">
      <div>
        <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
          <Settings class="w-8 h-8 text-indigo-600" />
          {{ t('settings.title') }}
        </h1>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <!-- Sidebar Info -->
        <div class="space-y-8">
          <div class="p-6 bg-white rounded-3xl border border-slate-100 shadow-sm">
            <div class="w-12 h-12 bg-indigo-50 rounded-2xl flex items-center justify-center text-indigo-600 mb-4">
              <Globe class="w-6 h-6" />
            </div>
            <h3 class="font-black text-slate-800 mb-2">{{ t('settings.localization') }}</h3>
            <p class="text-xs text-slate-400 font-medium leading-relaxed">{{ t('settings.localizationDesc') }}</p>
          </div>

          <div class="p-6 bg-indigo-600 rounded-3xl shadow-xl shadow-indigo-100 text-white">
            <div class="w-12 h-12 bg-white/20 rounded-2xl flex items-center justify-center text-white mb-4">
              <Store class="w-6 h-6" />
            </div>
            <h3 class="font-black mb-2">{{ t('settings.tenantProfile') }}</h3>
            <p class="text-xs text-indigo-200 font-medium leading-relaxed">{{ t('settings.tenantProfileDesc') }}</p>
          </div>
        </div>

        <!-- Form Area -->
        <div class="md:col-span-2 space-y-6">
          <div class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 p-10">
            <form @submit.prevent="saveSettings" class="space-y-8">
              <div class="space-y-6">
                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('settings.language') }}</label>
                  <div class="grid grid-cols-1 gap-3">
                    <select v-model="settings.language" class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none appearance-none">
                      <option value="en">English (US)</option>
                      <option value="id">Bahasa Indonesia</option>
                      <option value="ja">日本語 (Japanese)</option>
                    </select>
                  </div>
                </div>

                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('settings.currency') }}</label>
                  <select v-model="settings.currency" class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none appearance-none">
                    <option value="USD">USD ($)</option>
                    <option value="IDR">IDR (Rp)</option>
                    <option value="JPY">JPY (¥)</option>
                  </select>
                </div>

                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ t('settings.tenantName') }}</label>
                  <div class="relative">
                    <input v-model="settings.tenant_name" type="text" required class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none pl-12" />
                    <Store class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300" />
                  </div>
                </div>
              </div>

              <div class="pt-4">
                <button 
                  type="submit" 
                  :disabled="loading"
                  class="w-full py-5 bg-slate-900 text-white rounded-[1.5rem] font-black flex items-center justify-center gap-3 hover:bg-indigo-600 transition-all shadow-xl shadow-slate-200 group disabled:opacity-50"
                >
                  <Save v-if="!loading" class="w-5 h-5 text-indigo-400 group-hover:text-white transition-colors" />
                  <span v-else class="w-5 h-5 border-4 border-white/30 border-t-white rounded-full animate-spin"></span>
                  {{ t('settings.saveSettings') }}
                </button>
                
                <div v-if="success" class="mt-4 flex items-center justify-center gap-2 text-green-600 font-bold animate-in fade-in slide-in-from-top-2">
                  <CheckCircle class="w-5 h-5" />
                  Settings updated successfully!
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>
