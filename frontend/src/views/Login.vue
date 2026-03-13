<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { Package, ArrowRight, ShieldCheck, Lock } from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
const configStore = useConfigStore();
const router = useRouter();

const username = ref('');
const password = ref('');

const login = async () => {
  try {
    await authStore.login({ username: username.value, password: password.value });
    router.push('/');
  } catch (error) {
    // Error handled by store
  }
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-50 p-6 font-sans">
    <!-- Language Switcher on Login -->
    <div class="absolute top-8 right-8 flex items-center bg-white p-1 rounded-xl shadow-sm border border-slate-100">
      <button @click="configStore.setLanguage('en')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'en' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-slate-600'">EN</button>
      <button @click="configStore.setLanguage('id')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'id' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-slate-600'">ID</button>
      <button @click="configStore.setLanguage('ja')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'ja' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-slate-600'">JA</button>
    </div>

    <div class="bg-white rounded-[3rem] shadow-2xl shadow-slate-200 w-full max-w-md overflow-hidden border border-slate-100 animate-in zoom-in-95 duration-500">
      <div class="p-10 pt-12">
        <div class="flex flex-col items-center mb-10 text-center">
          <div class="bg-indigo-600 p-4 rounded-3xl text-white shadow-xl shadow-indigo-200 mb-6">
            <Package class="w-10 h-10" />
          </div>
          <h2 class="text-3xl font-black text-slate-800 tracking-tight">Invent Story</h2>
          <p class="text-slate-400 font-medium mt-2">{{ t('common.tagline') }}</p>
        </div>
        
        <div v-if="authStore.error" class="bg-red-50 text-red-600 p-4 rounded-2xl mb-8 text-sm font-bold border border-red-100 flex items-center gap-2">
          <ShieldCheck class="w-4 h-4 flex-shrink-0" />
          {{ authStore.error }}
        </div>

        <form @submit.prevent="login" class="space-y-6">
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] ml-1">{{ t('common.username') }}</label>
            <input 
              v-model="username" 
              type="text" 
              required
              class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none"
            />
          </div>
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] ml-1">{{ t('common.password') }}</label>
            <div class="relative">
              <input 
                v-model="password" 
                type="password" 
                required
                class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none"
              />
              <Lock class="absolute right-6 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300" />
            </div>
          </div>
          
          <button 
            type="submit" 
            :disabled="authStore.loading"
            class="w-full bg-indigo-600 text-white py-5 rounded-[1.5rem] font-black shadow-xl shadow-indigo-100 hover:bg-indigo-700 disabled:opacity-50 transition-all flex items-center justify-center gap-3 group"
          >
            {{ authStore.loading ? t('common.loading') : t('common.login') }}
            <ArrowRight class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
          </button>
        </form>

        <div class="mt-12 text-center">
          <router-link 
            to="/register" 
            class="inline-block px-8 py-3 bg-white border-2 border-slate-100 text-indigo-600 rounded-2xl font-black hover:border-indigo-600 hover:bg-indigo-50 transition-all"
          >
            {{ t('common.registerNewTenant') }}
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
