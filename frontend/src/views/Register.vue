<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { Package, ArrowLeft, CheckCircle2, Store, Sun, Moon } from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
const configStore = useConfigStore();
const router = useRouter();

const username = ref('');
const password = ref('');
const tenantName = ref('');

const register = async () => {
  try {
    await authStore.register({ 
      username: username.value, 
      password: password.value,
      tenant_name: tenantName.value
    });
    router.push('/');
  } catch (error) {
    // Error handled by store
  }
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-50 dark:bg-slate-950 p-6 font-sans transition-colors duration-300">
    <!-- Top Actions -->
    <div class="absolute top-8 right-8 flex items-center gap-3">
      <!-- Theme Toggle -->
      <button 
        @click="configStore.toggleTheme"
        class="p-2.5 rounded-xl bg-white dark:bg-slate-900 text-slate-500 dark:text-slate-400 shadow-sm border border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-slate-800 transition-all"
      >
        <Sun v-if="configStore.theme === 'dark'" class="w-5 h-5" />
        <Moon v-else class="w-5 h-5" />
      </button>

      <!-- Language Switcher -->
      <div class="flex items-center bg-white dark:bg-slate-900 p-1 rounded-xl shadow-sm border border-slate-100 dark:border-slate-800">
        <button @click="configStore.setLanguage('en')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'en' ? 'bg-indigo-600 text-white' : 'text-slate-400 dark:text-slate-500 hover:text-slate-600 dark:hover:text-slate-300'">EN</button>
        <button @click="configStore.setLanguage('id')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'id' ? 'bg-indigo-600 text-white' : 'text-slate-400 dark:text-slate-500 hover:text-slate-600 dark:hover:text-slate-300'">ID</button>
        <button @click="configStore.setLanguage('ja')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'ja' ? 'bg-indigo-600 text-white' : 'text-slate-400 dark:text-slate-500 hover:text-slate-600 dark:hover:text-slate-300'">JA</button>
      </div>
    </div>

    <div class="bg-white dark:bg-slate-900 rounded-[3rem] shadow-2xl shadow-slate-200 dark:shadow-none w-full max-w-lg overflow-hidden border border-slate-100 dark:border-slate-800 animate-in slide-in-from-bottom-8 duration-700">
      <div class="flex flex-col md:flex-row h-full">
        <!-- Decoration side -->
        <div class="hidden md:flex md:w-1/3 bg-indigo-600 p-10 flex-col justify-between text-white relative overflow-hidden">
          <div class="absolute -right-10 top-20 opacity-10">
            <Package class="w-64 h-64 rotate-12" />
          </div>
          
          <div class="relative z-10">
            <div class="bg-white/20 p-3 rounded-2xl w-fit mb-8">
              <Package class="w-8 h-8" />
            </div>
            <h2 class="text-3xl font-black leading-tight mb-4">Invent Story</h2>
            <p class="text-indigo-100 text-sm font-medium">{{ t('common.tagline') }}</p>
          </div>
          
          <div class="relative z-10 space-y-4">
            <div class="flex items-center gap-3">
              <CheckCircle2 class="w-5 h-5 text-indigo-300" />
              <span class="text-xs font-bold uppercase tracking-widest text-indigo-100">Multi-Tenancy</span>
            </div>
            <div class="flex items-center gap-3">
              <CheckCircle2 class="w-5 h-5 text-indigo-300" />
              <span class="text-xs font-bold uppercase tracking-widest text-indigo-100">Stock Control</span>
            </div>
          </div>
        </div>

        <!-- Form side -->
        <div class="flex-1 p-10 md:p-12">
          <div class="mb-10">
            <h2 class="text-3xl font-black text-slate-800 dark:text-slate-100 tracking-tight">Invent Story</h2>
            <p class="text-slate-400 dark:text-slate-500 font-medium mt-2">{{ t('common.tagline') }}</p>
          </div>
          
          <div v-if="authStore.error" class="bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400 p-4 rounded-2xl mb-8 text-sm font-bold border border-red-100 dark:border-red-900/50">
            {{ authStore.error }}
          </div>

          <form @submit.prevent="register" class="space-y-6">
            <div class="space-y-2">
              <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-[0.2em] ml-1">{{ t('common.restaurantName') }}</label>
              <div class="relative">
                <input 
                  v-model="tenantName" 
                  type="text" 
                  required
                  placeholder="e.g. Gusto Italian"
                  class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none pl-14 dark:text-slate-100"
                />
                <Store class="absolute left-6 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300 dark:text-slate-600" />
              </div>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-2">
                <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-[0.2em] ml-1">{{ t('common.username') }}</label>
                <input 
                  v-model="username" 
                  type="text" 
                  required
                  placeholder="admin"
                  class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100"
                />
              </div>
              <div class="space-y-2">
                <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-[0.2em] ml-1">{{ t('common.password') }}</label>
                <input 
                  v-model="password" 
                  type="password" 
                  required
                  placeholder="••••••••"
                  class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border-2 border-transparent rounded-2xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100"
                />
              </div>
            </div>
            
            <div class="pt-4 flex flex-col gap-4">
              <button 
                type="submit" 
                :disabled="authStore.loading"
                class="w-full bg-slate-900 dark:bg-indigo-600 text-white py-5 rounded-[1.5rem] font-black shadow-xl shadow-slate-200 dark:shadow-none hover:bg-indigo-600 dark:hover:bg-indigo-700 disabled:opacity-50 transition-all flex items-center justify-center gap-3 group"
              >
                {{ authStore.loading ? t('common.loading') : t('common.register') }}
                <CheckCircle2 class="w-5 h-5 text-green-400 group-hover:text-white transition-colors" />
              </button>
              
              <button 
                type="button"
                @click="router.push('/login')"
                class="flex items-center justify-center gap-2 text-sm font-bold text-slate-400 dark:text-slate-500 hover:text-slate-600 dark:hover:text-slate-300 transition-colors"
              >
                <ArrowLeft class="w-4 h-4" />
                {{ t('common.back') }} {{ t('common.login') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
