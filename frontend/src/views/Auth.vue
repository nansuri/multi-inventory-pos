<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter, useRoute } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { 
  Package, 
  ArrowRight, 
  ShieldCheck, 
  Lock, 
  Sun, 
  Moon,
  Store,
  Mail
} from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
const configStore = useConfigStore();
const router = useRouter();
const route = useRoute();

// State to toggle between login and register
const mode = ref(route.path === '/register' ? 'register' : 'login');

// Form data
const form = ref({
  username: '',
  password: '',
  tenant_name: ''
});

const switchMode = (newMode: 'login' | 'register') => {
  mode.value = newMode;
  // Update URL without reloading
  window.history.pushState({}, '', `/${newMode}`);
  authStore.error = null;
};

const handleSubmit = async () => {
  try {
    if (mode.value === 'login') {
      await authStore.login({ username: form.value.username, password: form.value.password });
    } else {
      await authStore.register({ 
        username: form.value.username, 
        password: form.value.password,
        tenant_name: form.value.tenant_name
      });
    }
    router.push('/');
  } catch (error) {
    // Error handled by store
  }
};

onMounted(() => {
  // Clear any old errors
  authStore.error = null;
});
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-6 font-sans relative overflow-hidden bg-[#0a0a0c]">
    <!-- Premium Immersive Background -->
    <div class="absolute inset-0 z-0">
      <!-- Vignette -->
      <div class="absolute inset-0 bg-gradient-to-b from-transparent via-black/20 to-black z-10"></div>
      <div class="absolute inset-0 shadow-[inset_0_0_150px_rgba(0,0,0,0.8)] z-10"></div>
      
      <!-- Subtle Tech Grid -->
      <div class="absolute inset-0 opacity-[0.03] z-0" style="background-image: radial-gradient(#fff 1px, transparent 1px); background-size: 40px 40px;"></div>
      
      <!-- Steam-like Flowing Nebula -->
      <div class="absolute top-[-20%] left-[-10%] w-[70%] h-[70%] bg-indigo-600/10 rounded-full blur-[140px] animate-pulse"></div>
      <div class="absolute bottom-[-20%] right-[-10%] w-[70%] h-[70%] bg-blue-600/10 rounded-full blur-[140px] animate-pulse" style="animation-delay: 3s"></div>
    </div>

    <!-- Top Actions -->
    <div class="absolute top-8 right-8 flex items-center gap-3 z-20">
      <button 
        @click="configStore.toggleTheme"
        class="p-2.5 rounded-xl bg-white/5 backdrop-blur-md text-slate-400 border border-white/5 hover:bg-white/10 hover:text-white transition-all"
      >
        <Sun v-if="configStore.theme === 'dark'" class="w-5 h-5" />
        <Moon v-else class="w-5 h-5" />
      </button>

      <div class="flex items-center bg-white/5 backdrop-blur-md p-1 rounded-xl border border-white/5">
        <button @click="configStore.setLanguage('en')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'en' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-white'">EN</button>
        <button @click="configStore.setLanguage('id')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'id' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-white'">ID</button>
        <button @click="configStore.setLanguage('ja')" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === 'ja' ? 'bg-indigo-600 text-white' : 'text-slate-400 hover:text-white'">JA</button>
      </div>
    </div>

    <div class="relative z-10 w-full max-w-[480px]">
      <!-- Logo Area -->
      <div class="flex flex-col items-center mb-10 text-center animate-in fade-in slide-in-from-top-4 duration-1000">
        <div class="bg-indigo-600 p-4 rounded-[1.5rem] text-white shadow-2xl shadow-indigo-500/20 mb-6">
          <Package class="w-10 h-10" />
        </div>
        <h1 class="text-3xl font-[1000] text-white tracking-tighter uppercase">Invent Story</h1>
      </div>

      <!-- Main Card -->
      <div class="bg-[#16161a]/80 backdrop-blur-2xl rounded-[2.5rem] shadow-[0_32px_64px_-16px_rgba(0,0,0,0.6)] border border-white/5 overflow-hidden">
        <div class="p-10 md:p-12">
          
          <Transition name="form-fade" mode="out-in">
            <div :key="mode">
              <!-- Header Section inside Transition -->
              <div class="mb-10">
                <h2 class="text-2xl font-black text-white tracking-tight">
                  {{ mode === 'login' ? t('common.welcomeBack') : 'Create Account' }}
                </h2>
                <p class="text-slate-500 font-medium text-sm mt-1">
                  {{ mode === 'login' ? t('common.tagline') : 'Start managing your restaurant today' }}
                </p>
              </div>

              <!-- Error Message -->
              <div v-if="authStore.error" class="bg-red-500/10 text-red-400 p-4 rounded-2xl mb-8 text-xs font-bold border border-red-500/20 flex items-center gap-3">
                <ShieldCheck class="w-4 h-4 shrink-0" />
                {{ authStore.error }}
              </div>

              <form @submit.prevent="handleSubmit" class="space-y-6">
                <!-- Register Only Fields -->
                <div v-if="mode === 'register'" class="space-y-2 animate-in fade-in slide-in-from-left-2 duration-300">
                  <label class="block text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] ml-1">{{ t('common.restaurantName') }}</label>
                  <div class="relative">
                    <input v-model="form.tenant_name" type="text" required placeholder="e.g. Gusto Italian"
                      class="w-full px-6 py-4 bg-white/[0.03] border border-white/10 rounded-2xl font-bold text-white placeholder:text-slate-700 focus:bg-white/[0.07] focus:border-indigo-500/50 transition-all outline-none pl-12" />
                    <Store class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-700" />
                  </div>
                </div>

                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] ml-1">{{ t('common.username') }}</label>
                  <div class="relative">
                    <input v-model="form.username" type="text" required placeholder="Username"
                      class="w-full px-6 py-4 bg-white/[0.03] border border-white/10 rounded-2xl font-bold text-white placeholder:text-slate-700 focus:bg-white/[0.07] focus:border-indigo-500/50 transition-all outline-none pl-12" />
                    <Mail class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-700" />
                  </div>
                </div>

                <div class="space-y-2">
                  <label class="block text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] ml-1">{{ t('common.password') }}</label>
                  <div class="relative">
                    <input v-model="form.password" type="password" required placeholder="••••••••"
                      class="w-full px-6 py-4 bg-white/[0.03] border border-white/10 rounded-2xl font-bold text-white placeholder:text-slate-700 focus:bg-white/[0.07] focus:border-indigo-500/50 transition-all outline-none pl-12" />
                    <Lock class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-700" />
                  </div>
                </div>

                <div class="pt-4">
                  <button type="submit" :disabled="authStore.loading"
                    class="w-full bg-indigo-600 text-white py-5 rounded-2xl font-black shadow-2xl shadow-indigo-500/20 hover:bg-indigo-500 hover:-translate-y-0.5 active:translate-y-0 disabled:opacity-50 transition-all flex items-center justify-center gap-3 group">
                    <span v-if="authStore.loading" class="w-5 h-5 border-4 border-white/30 border-t-white rounded-full animate-spin"></span>
                    <template v-else>
                      {{ mode === 'login' ? t('common.login') : t('common.register') }}
                      <ArrowRight class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
                    </template>
                  </button>
                </div>
              </form>

              <!-- Switch Mode Footer -->
              <div class="mt-10 text-center">
                <button 
                  @click="switchMode(mode === 'login' ? 'register' : 'login')"
                  class="text-sm font-bold text-slate-500 hover:text-indigo-400 transition-colors"
                >
                  {{ mode === 'login' ? t('common.registerNewTenant') : t('common.alreadyHaveAccount') + ' ' + t('common.login') }}
                </button>
              </div>
            </div>
          </Transition>

        </div>
      </div>

      <!-- System Status / Info -->
      <div class="mt-10 flex justify-center gap-8 text-[10px] font-black text-slate-600 uppercase tracking-[0.2em] animate-in fade-in slide-in-from-bottom-2 duration-1000">
        <div class="flex items-center gap-2">
          <div class="w-1.5 h-1.5 rounded-full bg-green-500 shadow-[0_0_8px_rgba(34,197,94,0.6)]"></div>
          System Live
        </div>
        <div>v1.0.4 - Enterprise</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.form-fade-enter-active,
.form-fade-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.form-fade-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.form-fade-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}

/* Glass effect override for dark theme */
.glass {
  background: rgba(22, 22, 26, 0.7);
  border: 1px solid rgba(255, 255, 255, 0.05);
}
</style>
