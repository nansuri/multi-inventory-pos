<script setup lang="ts">
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter, useRoute } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { 
  LayoutDashboard, 
  Package, 
  ShoppingCart, 
  LogOut,
  History,
  BookOpen,
  ChefHat,
  ChevronLeft,
  User,
  Settings
} from 'lucide-vue-next';
import { ref, computed, onMounted } from 'vue';

const { t } = useI18n();
const authStore = useAuthStore();
const configStore = useConfigStore();
const router = useRouter();
const route = useRoute();
const isSidebarCollapsed = ref(false);

const logout = () => {
  authStore.logout();
  router.push('/login');
};

const navItems = computed(() => [
  { name: t('common.dashboard'), path: '/', icon: LayoutDashboard },
  { name: t('common.inventory'), path: '/inventory', icon: Package },
  { name: t('common.recipes'), path: '/recipes', icon: BookOpen },
  { name: t('common.production'), path: '/products', icon: ChefHat },
  { name: t('common.pos'), path: '/pos', icon: ShoppingCart },
  { name: t('common.preparations'), path: '/preparations', icon: History },
  { name: t('common.settings'), path: '/settings', icon: Settings },
]);

const currentRouteName = computed(() => route.name as string);

onMounted(() => {
  configStore.fetchTenantSettings();
});

const languages = [
  { code: 'en', name: 'EN' },
  { code: 'id', name: 'ID' },
  { code: 'ja', name: 'JA' }
];
</script>

<template>
  <div class="min-h-screen bg-slate-50 flex font-sans text-slate-900">
    <!-- Sidebar -->
    <aside 
      :class="[
        'bg-white border-r border-slate-200 transition-all duration-300 flex flex-col sticky top-0 h-screen z-30',
        isSidebarCollapsed ? 'w-20' : 'w-64'
      ]"
    >
      <div class="h-16 flex items-center px-6 border-b border-slate-100 mb-4">
        <div class="flex items-center gap-3 overflow-hidden">
          <div class="bg-indigo-600 p-2 rounded-xl text-white flex-shrink-0">
            <Package class="w-6 h-6" />
          </div>
          <span v-if="!isSidebarCollapsed" class="font-black text-xl tracking-tight text-slate-800 truncate">
            {{ configStore.tenantName || 'Invent' }}<span class="text-indigo-600" v-if="!configStore.tenantName"> Story</span>
          </span>
        </div>
      </div>

      <nav class="flex-1 px-3 space-y-1">
        <router-link 
          v-for="item in navItems" 
          :key="item.path"
          :to="item.path"
          class="flex items-center px-3 py-2.5 rounded-xl transition-all duration-200 group relative"
          :class="[
            $route.path === item.path 
              ? 'bg-indigo-50 text-indigo-700 shadow-sm shadow-indigo-100' 
              : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'
          ]"
        >
          <component 
            :is="item.icon" 
            :class="[
              'w-5 h-5 flex-shrink-0 transition-colors',
              $route.path === item.path ? 'text-indigo-600' : 'group-hover:text-slate-900'
            ]" 
          />
          <span 
            v-if="!isSidebarCollapsed" 
            class="ml-3 font-semibold text-sm whitespace-nowrap transition-opacity duration-300"
          >
            {{ item.name }}
          </span>
        </router-link>
      </nav>

      <div class="p-3 border-t border-slate-100 space-y-1">
        <button 
          @click="logout"
          class="flex items-center w-full px-3 py-2.5 text-slate-500 hover:bg-red-50 hover:text-red-600 transition-all rounded-xl group relative"
        >
          <LogOut class="w-5 h-5 flex-shrink-0" />
          <span v-if="!isSidebarCollapsed" class="ml-3 font-semibold text-sm">{{ $t('common.logout') }}</span>
        </button>
        
        <button 
          @click="isSidebarCollapsed = !isSidebarCollapsed"
          class="flex items-center w-full px-3 py-2.5 text-slate-400 hover:bg-slate-50 hover:text-slate-600 transition-all rounded-xl hidden md:flex"
        >
          <ChevronLeft 
            :class="['w-5 h-5 transition-transform duration-300', isSidebarCollapsed ? 'rotate-180' : '']" 
          />
          <span v-if="!isSidebarCollapsed" class="ml-3 font-medium text-xs uppercase tracking-widest">Collapse</span>
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col min-w-0">
      <header class="h-16 bg-white/80 backdrop-blur-md border-b border-slate-200 px-8 flex justify-between items-center sticky top-0 z-20">
        <h2 class="text-lg font-bold text-slate-800 capitalize tracking-tight">
          {{ currentRouteName }}
        </h2>
        
        <div class="flex items-center gap-6">
          <!-- Language Switcher -->
          <div class="flex items-center bg-slate-100 p-1 rounded-xl">
            <button 
              v-for="lang in languages" 
              :key="lang.code"
              @click="configStore.setLanguage(lang.code)"
              class="px-3 py-1 rounded-lg text-[10px] font-black transition-all"
              :class="configStore.language === lang.code ? 'bg-white text-indigo-600 shadow-sm' : 'text-slate-400 hover:text-slate-600'"
            >
              {{ lang.name }}
            </button>
          </div>

          <div class="h-8 w-px bg-slate-200"></div>
          
          <div class="flex items-center gap-3 group cursor-pointer" @click="router.push('/settings')">
            <div class="text-right hidden sm:block">
              <p class="text-sm font-bold text-slate-800 leading-tight">{{ authStore.user?.username }}</p>
              <p class="text-[10px] font-medium text-slate-400 uppercase tracking-widest">{{ authStore.user?.role }}</p>
            </div>
            <div class="w-10 h-10 rounded-xl bg-indigo-100 text-indigo-600 flex items-center justify-center border-2 border-indigo-50 group-hover:bg-indigo-600 group-hover:text-white transition-all shadow-sm">
              <User class="w-5 h-5 font-bold" />
            </div>
          </div>
        </div>
      </header>

      <main class="p-8 pb-20">
        <slot />
      </main>
    </div>
  </div>
</template>
