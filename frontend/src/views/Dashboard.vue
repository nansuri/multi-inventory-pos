<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useAuthStore } from '../stores/auth';
import { useRouter } from 'vue-router';
import { 
  Package, 
  AlertTriangle, 
  DollarSign, 
  Grid,
  TrendingUp,
  ArrowRight,
  ShoppingCart
} from 'lucide-vue-next';

const summary = ref<any>(null);
const loading = ref(true);
const configStore = useConfigStore();
const authStore = useAuthStore();
const router = useRouter();

const fetchSummary = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/dashboard/summary');
    summary.value = response.data;
  } catch (error) {
    console.error('Failed to fetch summary', error);
  } finally {
    loading.value = false;
  }
};

onMounted(fetchSummary);

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, fetchSummary);
</script>

<template>
  <div class="space-y-10 animate-in fade-in duration-700">
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="w-12 h-12 border-4 border-indigo-600/30 border-t-indigo-600 rounded-full animate-spin"></div>
    </div>

    <div v-else-if="summary" class="space-y-10">
      <!-- Welcome Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-3xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tighter">
            System Overview
          </h1>
          <p class="text-slate-500 dark:text-slate-400 font-medium mt-1">Here's what's happening with your restaurant today.</p>
        </div>
        <div class="flex items-center gap-2 px-4 py-2 bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 rounded-2xl border border-indigo-100 dark:border-indigo-900/50">
          <div class="w-2 h-2 rounded-full bg-indigo-600 animate-pulse"></div>
          <span class="text-[10px] font-black uppercase tracking-widest">Live Updates</span>
        </div>
      </div>

      <!-- Stats Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- Revenue -->
        <div class="bg-white dark:bg-slate-900 p-8 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all group">
          <div class="flex justify-between items-start mb-6">
            <div class="p-4 bg-emerald-50 dark:bg-emerald-900/30 rounded-2xl text-emerald-600 dark:text-emerald-400 group-hover:bg-emerald-600 group-hover:text-white transition-all">
              <DollarSign class="w-6 h-6" />
            </div>
            <span class="text-[10px] font-black text-emerald-600 dark:text-emerald-400 bg-emerald-50 dark:bg-emerald-900/30 px-2 py-1 rounded-lg">+12%</span>
          </div>
          <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">Today's Revenue</p>
          <h3 class="text-2xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tight">{{ configStore.formatCurrency(summary.today_revenue) }}</h3>
        </div>

        <!-- Orders -->
        <div class="bg-white dark:bg-slate-900 p-8 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all group">
          <div class="flex justify-between items-start mb-6">
            <div class="p-4 bg-indigo-50 dark:bg-indigo-900/30 rounded-2xl text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
              <ShoppingCart class="w-6 h-6" />
            </div>
          </div>
          <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">Orders Today</p>
          <h3 class="text-2xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tight">{{ summary.today_orders }}</h3>
        </div>

        <!-- Low Stock -->
        <div class="bg-white dark:bg-slate-900 p-8 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all group">
          <div class="flex justify-between items-start mb-6">
            <div class="p-4 bg-orange-50 dark:bg-orange-900/30 rounded-2xl text-orange-600 dark:text-orange-400 group-hover:bg-orange-600 group-hover:text-white transition-all">
              <AlertTriangle class="w-6 h-6" />
            </div>
            <span v-if="summary.low_stock_items > 0" class="flex h-3 w-3">
              <span class="animate-ping absolute inline-flex h-3 w-3 rounded-full bg-orange-400 opacity-75"></span>
              <span class="relative inline-flex rounded-full h-3 w-3 bg-orange-500"></span>
            </span>
          </div>
          <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">Low Stock Alerts</p>
          <h3 class="text-2xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tight">{{ summary.low_stock_items }}</h3>
        </div>

        <!-- Ingredients -->
        <div class="bg-white dark:bg-slate-900 p-8 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all group">
          <div class="flex justify-between items-start mb-6">
            <div class="p-4 bg-blue-50 dark:bg-blue-900/30 rounded-2xl text-blue-600 dark:text-blue-400 group-hover:bg-blue-600 group-hover:text-white transition-all">
              <Package class="w-6 h-6" />
            </div>
          </div>
          <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">Total Ingredients</p>
          <h3 class="text-2xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tight">{{ summary.total_items }}</h3>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Tables/Occupancy -->
        <div class="lg:col-span-2 bg-white dark:bg-slate-900 p-10 rounded-[3rem] shadow-sm border border-slate-100 dark:border-slate-800">
          <div class="flex justify-between items-center mb-10">
            <div>
              <h3 class="text-xl font-black text-slate-800 dark:text-slate-100 flex items-center gap-3">
                <Grid class="w-6 h-6 text-indigo-600" />
                Table Occupancy
              </h3>
            </div>
            <router-link to="/pos/order" class="text-xs font-black text-indigo-600 hover:underline uppercase tracking-widest">Manage Tables</router-link>
          </div>
          
          <div class="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-5 gap-4">
            <div v-for="table in summary.tables" :key="table.id" 
              :class="['p-4 rounded-2xl border-2 flex flex-col items-center justify-center gap-2 transition-all', 
                table.status === 'occupied' ? 'bg-indigo-50 dark:bg-indigo-900/20 border-indigo-100 dark:border-indigo-800 text-indigo-600 dark:text-indigo-400' : 'bg-slate-50 dark:bg-slate-800/50 border-transparent text-slate-400']">
              <span class="text-lg font-black tracking-tighter">{{ table.table_number }}</span>
              <span class="text-[8px] font-black uppercase tracking-widest">{{ table.status }}</span>
            </div>
          </div>
        </div>

        <!-- Quick Actions -->
        <div class="bg-indigo-600 p-10 rounded-[3rem] shadow-2xl shadow-indigo-200 dark:shadow-none text-white flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute -right-10 -bottom-10 opacity-10 group-hover:scale-110 transition-transform duration-700">
            <TrendingUp class="w-64 h-64" />
          </div>
          
          <div class="relative z-10">
            <h3 class="text-2xl font-black tracking-tight mb-2">Production Ready?</h3>
            <p class="text-indigo-100 text-sm font-medium leading-relaxed">Check your recipes and start cooking to fill up your available stock.</p>
          </div>

          <div class="relative z-10 space-y-3 mt-10">
            <button @click="router.push('/pos/order')" class="w-full py-4 bg-white text-indigo-600 rounded-2xl font-black shadow-xl hover:bg-indigo-50 transition-all flex items-center justify-center gap-3">
              New Order
              <ArrowRight class="w-4 h-4" />
            </button>
            <button @click="router.push('/products')" class="w-full py-4 bg-indigo-500 text-white border border-indigo-400 rounded-2xl font-black hover:bg-indigo-400 transition-all">
              Go to Production Hall
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
