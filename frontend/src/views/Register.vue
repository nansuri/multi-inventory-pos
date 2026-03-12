<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter } from 'vue-router';
import { Package, ArrowLeft, CheckCircle2, Store } from 'lucide-vue-next';

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
  <div class="min-h-screen flex items-center justify-center bg-slate-50 p-6 font-sans">
    <!-- Language Switcher on Register -->
    <div class="absolute top-8 right-8 flex items-center bg-white p-1 rounded-xl shadow-sm border border-slate-100">
      <button @click="configStore.setLanguage('en')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'en' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-slate-600'">EN</button>
      <button @click="configStore.setLanguage('id')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'id' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-slate-600'">ID</button>
      <button @click="configStore.setLanguage('ja')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'ja' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-slate-600'">JA</button>
    </div>

    <div class="bg-white rounded-[3rem] shadow-2xl shadow-slate-200 w-full max-w-lg overflow-hidden border border-slate-100 animate-in slide-in-from-bottom-8 duration-700">
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
            <h2 class="text-3xl font-black leading-tight mb-4">Start your journey.</h2>
          </div>
        </div>

        <!-- Form side -->
        <div class="flex-1 p-10 md:p-12">
          <div class="mb-10">
            <h2 class="text-3xl font-black text-slate-800 tracking-tight">Invent Story</h2>
            <p class="text-slate-400 font-medium mt-2">Every Story start with small story</p>
          </div>
          
          <div v-if="authStore.error" class="bg-red-50 text-red-600 p-4 rounded-2xl mb-8 text-sm font-bold border border-red-100">
            {{ authStore.error }}
          </div>

          <form @submit.prevent="register" class="space-y-6">
            <div class="space-y-2">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] ml-1">{{ $t('settings.tenantName') }}</label>
              <div class="relative">
                <input 
                  v-model="tenantName" 
                  type="text" 
                  required
                  placeholder="e.g. Gusto Italian"
                  class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none pl-14"
                />
                <Store class="absolute left-6 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300" />
              </div>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-2">
                <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] ml-1">Admin Username</label>
                <input 
                  v-model="username" 
                  type="text" 
                  required
                  class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none"
                />
              </div>
              <div class="space-y-2">
                <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] ml-1">Secure Password</label>
                <input 
                  v-model="password" 
                  type="password" 
                  required
                  placeholder="••••••••"
                  class="w-full px-6 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none"
                />
              </div>
            </div>
            
            <div class="pt-4 flex flex-col gap-4">
              <button 
                type="submit" 
                :disabled="authStore.loading"
                class="w-full bg-slate-900 text-white py-5 rounded-[1.5rem] font-black shadow-xl shadow-slate-200 hover:bg-indigo-600 disabled:opacity-50 transition-all flex items-center justify-center gap-3 group"
              >
                {{ authStore.loading ? $t('common.loading') : 'Finalize Registration' }}
                <CheckCircle2 class="w-5 h-5 text-green-400 group-hover:text-white transition-colors" />
              </button>
              
              <button 
                type="button"
                @click="router.push('/login')"
                class="flex items-center justify-center gap-2 text-sm font-bold text-slate-400 hover:text-slate-600 transition-colors"
              >
                <ArrowLeft class="w-4 h-4" />
                {{ $t('common.back') }} to Login
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
