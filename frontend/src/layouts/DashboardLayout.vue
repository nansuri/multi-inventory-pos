<script setup lang="ts">
import { useAuthStore } from '../stores/auth';
import { useConfigStore } from '../stores/config';
import { useRouter, useRoute } from 'vue-router';
import { useI18n } from 'vue-i18n';
import SidebarItem from '../components/SidebarItem.vue';
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
  HelpCircle
} from 'lucide-vue-next';
import { ref, computed, onMounted } from 'vue';

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
      name: 'System Guide', 
      path: '/guide', 
      icon: HelpCircle,
      permission: 'dashboard'
    },
    {
      type: 'group',
      id: 'kitchen',
      name: 'Kitchen & Storage',
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
      name: 'Sales & Floor',
      icon: Utensils,
      children: [
        { name: 'New Order', path: '/pos/order', icon: Utensils, permission: 'pos_order' },
        { name: 'Payments', path: '/pos/payment', icon: CreditCard, permission: 'pos_payment' },
      ].filter(c => hasPermission(c.permission))
    },
    {
      type: 'group',
      id: 'admin',
      name: 'Administration',
      icon: Users,
      children: [
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
  }
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
        isSidebarCollapsed ? 'w-20' : 'w-72'
      ]"
    >
      <!-- Logo Area -->
      <div class="h-16 flex items-center px-6 border-b border-slate-100 mb-4">
        <div class="flex items-center gap-3 overflow-hidden">
          <div class="bg-indigo-600 p-2 rounded-xl text-white flex-shrink-0 shadow-lg shadow-indigo-100">
            <Package class="w-6 h-6" />
          </div>
          <span v-if="!isSidebarCollapsed" class="font-black text-xl tracking-tight text-slate-800 truncate">
            {{ configStore.tenantName || 'Invent' }}<span class="text-indigo-600" v-if="!configStore.tenantName"> Story</span>
          </span>
        </div>
      </div>

      <!-- Navigation -->
      <nav class="flex-1 px-3 space-y-1 overflow-y-auto custom-scrollbar pb-10">
        <template v-for="nav in navItems" :key="nav.name">
          <!-- Single Top-Level Item -->
          <SidebarItem 
            v-if="nav.type === 'item' && nav.path"
            :name="nav.name"
            :path="nav.path"
            :icon="nav.icon"
            :isCollapsed="isSidebarCollapsed"
          />

          <!-- Collapsible Group -->
          <div v-else-if="nav.type === 'group'" class="space-y-1 mb-2 pt-2">
            <!-- Group Header (Only visible if not collapsed) -->
            <div 
              v-if="!isSidebarCollapsed"
              @click="toggleGroup(nav.id!)"
              class="flex items-center justify-between px-4 py-2 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] cursor-pointer hover:text-slate-600 transition-colors group"
            >
              <div class="flex items-center gap-2">
                <component :is="nav.icon" class="w-3.5 h-3.5" />
                <span>{{ nav.name }}</span>
              </div>
              <ChevronDown :class="['w-3 h-3 transition-transform duration-300', openGroups[nav.id!] ? '' : '-rotate-90']" />
            </div>
            
            <!-- Children Items -->
            <div v-if="isSidebarCollapsed || openGroups[nav.id!]" class="space-y-1">
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
      <div class="p-3 border-t border-slate-100 space-y-1">
        <button 
          @click="logout"
          class="flex items-center w-full px-4 py-3 text-slate-500 hover:bg-red-50 hover:text-red-600 transition-all rounded-2xl group relative"
        >
          <LogOut class="w-5 h-5 flex-shrink-0" />
          <span v-if="!isSidebarCollapsed" class="ml-3 font-bold text-sm">{{ $t('common.logout') }}</span>
          <div v-if="isSidebarCollapsed" class="absolute left-14 bg-red-600 text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 pointer-events-none transition-opacity whitespace-nowrap z-50 shadow-xl font-black uppercase tracking-widest">Logout</div>
        </button>
        
        <button 
          @click="isSidebarCollapsed = !isSidebarCollapsed"
          class="flex items-center w-full px-4 py-2 text-slate-400 hover:bg-slate-50 hover:text-slate-600 transition-all rounded-2xl hidden md:flex"
        >
          <ChevronLeft :class="['w-5 h-5 transition-transform duration-300', isSidebarCollapsed ? 'rotate-180' : '']" />
          <span v-if="!isSidebarCollapsed" class="ml-3 font-black text-[10px] uppercase tracking-widest">Collapse</span>
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col min-w-0 h-screen">
      <header class="h-16 bg-white/80 backdrop-blur-md border-b border-slate-200 px-8 flex justify-between items-center sticky top-0 z-20">
        <h2 class="text-lg font-black text-slate-800 tracking-tight uppercase tracking-widest text-xs">{{ currentRouteName }}</h2>
        
        <div class="flex items-center gap-6">
          <!-- Language Switcher -->
          <div class="flex items-center bg-slate-100 p-1 rounded-xl">
            <button v-for="lang in languages" :key="lang.code" @click="configStore.setLanguage(lang.code)" class="px-3 py-1 rounded-lg text-[10px] font-black transition-all" :class="configStore.language === lang.code ? 'bg-white text-indigo-600 shadow-sm' : 'text-slate-400 hover:text-slate-600'">{{ lang.name }}</button>
          </div>

          <div class="h-8 w-px bg-slate-200"></div>
          
          <div class="flex items-center gap-3 group cursor-pointer" @click="router.push('/settings')">
            <div class="text-right hidden sm:block">
              <p class="text-sm font-black text-slate-800 leading-tight">{{ authStore.user?.username }}</p>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ authStore.user?.is_owner ? 'Owner' : authStore.user?.role?.name || 'Staff' }}</p>
            </div>
            <div class="w-10 h-10 rounded-2xl bg-indigo-100 text-indigo-600 flex items-center justify-center border-2 border-indigo-50 group-hover:bg-indigo-600 group-hover:text-white transition-all shadow-sm">
              <User class="w-5 h-5 font-bold" />
            </div>
          </div>
        </div>
      </header>

      <main class="flex-1 p-8 overflow-y-auto custom-scrollbar bg-slate-50/50">
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
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: var(--color-indigo-200);
}
</style>
