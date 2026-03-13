<script setup lang="ts">
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter, useRoute } from 'vue-router';
import { useI18n } from 'vue-i18n';
import SidebarItem from '../components/SidebarItem.vue';
import BranchSelector from '../components/BranchSelector.vue';
import { 
  LayoutDashboard, 
  Package, 
  LogOut,
  History,
  BookOpen,
  ChefHat,
  ChevronLeft,
  ChevronDown,
  User,
  Settings,
  CreditCard,
  Utensils,
  Users,
  LineChart,
  HelpCircle,
  Sun,
  Moon,
  MapPin,
  Grid
} from 'lucide-vue-next';
import { ref, computed, onMounted, watch } from 'vue';

interface NavChild {
  name: string;
  path: string;
  icon: any;
  permission: string;
}

interface NavItem {
  type: 'item' | 'group';
  id?: string;
  name: string;
  path?: string;
  icon: any;
  permission?: string;
  children?: NavChild[];
}

const { t } = useI18n();
const authStore = useAuthStore();
const configStore = useConfigStore();
const router = useRouter();
const route = useRoute();
const isSidebarCollapsed = ref(false);
const mainKey = ref(0);
const isSwitchingBranch = ref(false);

const logout = () => {
  authStore.logout();
  router.push('/login');
};

const openGroups = ref<Record<string, boolean>>({
  kitchen: true,
  sales: true,
  admin: true
});

const toggleGroup = (group: string) => {
  openGroups.value[group] = !openGroups.value[group];
};

const hasPermission = (permission: string) => {
  if (authStore.user?.is_owner) return true;
  return authStore.user?.role?.permissions?.includes(permission);
};

const navItems = computed<NavItem[]>(() => {
  const items: NavItem[] = [
    { 
      type: 'item',
      name: t('common.dashboard'), 
      path: '/', 
      icon: LayoutDashboard,
      permission: 'dashboard'
    },
    { 
      type: 'item',
      name: t('common.guide'), 
      path: '/guide', 
      icon: HelpCircle,
      permission: 'dashboard'
    },
    {
      type: 'group',
      id: 'kitchen',
      name: t('common.kitchenStorage'),
      icon: Package,
      children: [
        { name: t('common.inventory'), path: '/inventory', icon: Package, permission: 'inventory' },
        { name: t('common.recipes'), path: '/recipes', icon: BookOpen, permission: 'recipes' },
        { name: t('common.production'), path: '/products', icon: ChefHat, permission: 'production' },
        { name: t('common.productionLog'), path: '/preparations', icon: History, permission: 'production_log' },
      ].filter(c => hasPermission(c.permission))
    },
    {
      type: 'group',
      id: 'sales',
      name: t('common.salesFloor'),
      icon: Utensils,
      children: [
        { name: 'New Order', path: '/pos/order', icon: Utensils, permission: 'pos_order' },
        { name: 'Payments', path: '/pos/payment', icon: CreditCard, permission: 'pos_payment' },
      ].filter(c => hasPermission(c.permission))
    },
    {
      type: 'group',
      id: 'admin',
      name: t('common.administration'),
      icon: Users,
      children: [
        { name: t('common.branches'), path: '/branches', icon: MapPin, permission: 'branches' },
        { name: 'Tables', path: '/tables', icon: Grid, permission: 'pos_order' },
        { name: 'Order History', path: '/reports/orders', icon: LineChart, permission: 'order_history' },
        { name: 'Employees', path: '/employees', icon: Users, permission: 'employees' },
      ].filter(c => hasPermission(c.permission))
    },
    { 
      type: 'item',
      name: t('common.settings'), 
      path: '/settings', 
      icon: Settings,
      permission: 'settings'
    },
  ];

  return items.filter(item => {
    if (item.type === 'item') return hasPermission(item.permission!);
    return item.children && item.children.length > 0;
  });
});

const currentRouteName = computed(() => route.name as string);

onMounted(async () => {
  await configStore.fetchTenantSettings();
  if (authStore.isAuthenticated) {
    await authStore.fetchUser();
    await authStore.fetchBranches();
  }
});

// Force refresh components when branch changes with animation
watch(() => authStore.selectedBranchID, () => {
  isSwitchingBranch.value = true;
  setTimeout(() => {
    mainKey.value++;
    isSwitchingBranch.value = false;
  }, 400);
});

const languages = [
  { code: 'en', name: 'EN' },
  { code: 'id', name: 'ID' },
  { code: 'ja', name: 'JA' }
];
</script>

<template>
  <div class="min-h-screen bg-slate-50 dark:bg-slate-950 flex font-sans text-slate-900 dark:text-slate-100 transition-colors duration-300 overflow-hidden">
    <!-- Sidebar -->
    <aside 
      :class="[
        'bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 transition-all duration-300 flex flex-col sticky top-0 h-screen z-30',
        isSidebarCollapsed ? 'w-16' : 'w-64'
      ]"
    >
      <!-- Logo Area -->
      <div class="h-14 flex items-center px-4 border-b border-slate-100 dark:border-slate-800 mb-2">
        <div class="flex items-center gap-2 overflow-hidden">
          <div class="bg-indigo-600 p-1.5 rounded-lg text-white flex-shrink-0 shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20">
            <Package class="w-5 h-5" />
          </div>
          <span v-if="!isSidebarCollapsed" class="font-black text-lg tracking-tight text-slate-800 dark:text-slate-100 truncate">
            {{ configStore.tenantName || 'Invent' }}
          </span>
        </div>
      </div>

      <!-- Branch Switcher -->
      <BranchSelector :isCollapsed="isSidebarCollapsed" />

      <!-- Navigation -->
      <nav class="flex-1 px-2 space-y-0.5 overflow-y-auto custom-scrollbar pb-10">
        <template v-for="nav in navItems" :key="nav.name">
          <SidebarItem 
            v-if="nav.type === 'item' && nav.path"
            :name="nav.name"
            :path="nav.path"
            :icon="nav.icon"
            :isCollapsed="isSidebarCollapsed"
          />

          <div v-else-if="nav.type === 'group'" class="space-y-0.5 mb-1 pt-1">
            <div 
              v-if="!isSidebarCollapsed"
              @click="toggleGroup(nav.id!)"
              class="flex items-center justify-between px-3 py-1.5 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-[0.2em] cursor-pointer hover:text-slate-600 dark:hover:text-slate-300 transition-colors group"
            >
              <div class="flex items-center gap-2">
                <span>{{ nav.name }}</span>
              </div>
              <ChevronDown :class="['w-2.5 h-2.5 transition-transform duration-300', openGroups[nav.id!] ? '' : '-rotate-90']" />
            </div>
            
            <div v-if="isSidebarCollapsed || openGroups[nav.id!]" class="space-y-0.5">
              <SidebarItem 
                v-for="child in nav.children" 
                :key="child.path"
                :name="child.name"
                :path="child.path"
                :icon="child.icon"
                :isCollapsed="isSidebarCollapsed"
                isSubItem
              />
            </div>
          </div>
        </template>
      </nav>

      <!-- Bottom Actions -->
      <div class="p-2 border-t border-slate-100 dark:border-slate-800 space-y-0.5">
        <button 
          @click="logout"
          class="flex items-center w-full px-3 py-2 text-slate-500 dark:text-slate-400 hover:bg-red-50 dark:hover:bg-red-950/30 hover:text-red-600 transition-all rounded-xl group relative"
        >
          <LogOut class="w-4 h-4 flex-shrink-0" />
          <span v-if="!isSidebarCollapsed" class="ml-3 font-bold text-xs">{{ t('common.logout') }}</span>
        </button>
        
        <button 
          @click="isSidebarCollapsed = !isSidebarCollapsed"
          class="flex items-center w-full px-3 py-2 text-slate-400 dark:text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800 hover:text-slate-600 dark:hover:text-slate-300 transition-all rounded-xl hidden md:flex"
        >
          <ChevronLeft :class="['w-4 h-4 transition-transform duration-300', isSidebarCollapsed ? 'rotate-180' : '']" />
          <span v-if="!isSidebarCollapsed" class="ml-3 font-black text-[9px] uppercase tracking-widest">Collapse</span>
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col min-w-0 h-screen relative">
      <!-- Branch Switching Overlay -->
      <Transition
        enter-active-class="transition duration-300 ease-out"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition duration-300 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="isSwitchingBranch" class="absolute inset-0 bg-slate-950/40 backdrop-blur-md z-50 flex items-center justify-center">
          <div class="flex flex-col items-center gap-4">
            <div class="w-10 h-10 border-4 border-indigo-600/30 border-t-indigo-600 rounded-full animate-spin"></div>
            <p class="text-white font-black uppercase tracking-[0.3em] text-[9px]">Syncing Branch Data...</p>
          </div>
        </div>
      </Transition>

      <header class="h-14 bg-white/80 dark:bg-slate-900/80 backdrop-blur-md border-b border-slate-200 dark:border-slate-800 px-6 flex justify-between items-center sticky top-0 z-20">
        <div class="flex items-center gap-6">
          <h2 class="text-sm font-black text-slate-800 dark:text-slate-100 tracking-tight uppercase tracking-widest">{{ currentRouteName }}</h2>
        </div>
        
        <div class="flex items-center gap-4">
          <!-- Theme Toggle -->
          <button @click="configStore.toggleTheme" class="p-2 rounded-lg bg-slate-100 dark:bg-slate-800 text-slate-500 hover:bg-slate-200 dark:hover:bg-slate-700 transition-all">
            <Sun v-if="configStore.theme === 'dark'" class="w-4 h-4" />
            <Moon v-else class="w-4 h-4" />
          </button>

          <!-- Language Switcher -->
          <div class="flex items-center bg-slate-100 dark:bg-slate-800 p-0.5 rounded-lg">
            <button v-for="lang in languages" :key="lang.code" @click="configStore.setLanguage(lang.code)" class="px-2 py-1 rounded-md text-[9px] font-black transition-all" :class="configStore.language === lang.code ? 'bg-white dark:bg-slate-700 text-indigo-600 dark:text-indigo-400 shadow-sm' : 'text-slate-400 hover:text-slate-600 dark:hover:text-slate-300'">{{ lang.name }}</button>
          </div>

          <div class="h-6 w-px bg-slate-200 dark:bg-slate-800"></div>
          
          <div class="flex items-center gap-2.5 group cursor-pointer" @click="router.push('/settings')">
            <div class="text-right hidden sm:block">
              <p class="text-xs font-black text-slate-800 dark:text-slate-100 leading-none">{{ authStore.user?.username }}</p>
              <p class="text-[9px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest mt-0.5">{{ authStore.user?.is_owner ? 'Owner' : authStore.user?.role?.name || 'Staff' }}</p>
            </div>
            <div class="w-8 h-8 rounded-xl bg-indigo-100 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 flex items-center justify-center border border-indigo-50 dark:border-indigo-900/50 group-hover:bg-indigo-600 group-hover:text-white transition-all shadow-sm">
              <User class="w-4 h-4 font-bold" />
            </div>
          </div>
        </div>
      </header>

      <main :key="mainKey" class="flex-1 p-5 overflow-y-auto custom-scrollbar bg-slate-50/50 dark:bg-slate-950/50">
        <slot />
      </main>
    </div>
  </div>
</template>

<style scoped>
.router-link-active {
  position: relative;
}
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-200);
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-800);
}
</style>
