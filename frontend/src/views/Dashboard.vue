<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { 
  Package, 
  AlertTriangle, 
  DollarSign, 
  Utensils, 
  Grid,
  TrendingUp,
  ArrowRight
} from 'lucide-vue-next';

const summary = ref<any>(null);
const loading = ref(true);
const configStore = useConfigStore();

const fetchSummary = async () => {
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
</script>

<template>
  <DashboardLayout>
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-xl h-12 w-12 border-4 border-indigo-600 border-t-transparent"></div>
    </div>

    <div v-else-if="summary" class="space-y-10 animate-in fade-in duration-500">
      <!-- Welcome Section -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 class="text-3xl font-black text-slate-800 tracking-tight">{{ $t('dashboard.overview') }}</h1>
          <p class="text-slate-500 font-medium">{{ $t('dashboard.overviewDesc', "Here's what's happening with your restaurant today.") }}</p>
        </div>
        <div class="flex items-center gap-2 bg-white px-4 py-2 rounded-2xl shadow-sm border border-slate-100">
          <span class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></span>
          <span class="text-sm font-bold text-slate-600 uppercase tracking-widest">{{ $t('dashboard.liveUpdates') }}</span>
        </div>
      </div>

      <!-- Stats Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-6">
        <!-- Revenue Card -->
        <div class="bg-white p-6 rounded-3xl shadow-sm border border-slate-100 card-hover flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:scale-110 transition-transform">
            <DollarSign class="w-24 h-24 text-indigo-600" />
          </div>
          <div class="flex items-center justify-between mb-4">
            <div class="p-3 bg-green-50 rounded-2xl text-green-600">
              <DollarSign class="w-6 h-6" />
            </div>
            <TrendingUp class="w-4 h-4 text-green-500" />
          </div>
          <div>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">{{ $t('dashboard.todayRevenue') }}</p>
            <p class="text-3xl font-black text-slate-800">{{ configStore.formatCurrency(summary.today_revenue) }}</p>
          </div>
        </div>

        <!-- Orders Card -->
        <div class="bg-white p-6 rounded-3xl shadow-sm border border-slate-100 card-hover flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:scale-110 transition-transform">
            <Utensils class="w-24 h-24 text-indigo-600" />
          </div>
          <div class="flex items-center justify-between mb-4">
            <div class="p-3 bg-indigo-50 rounded-2xl text-indigo-600">
              <Utensils class="w-6 h-6" />
            </div>
          </div>
          <div>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">{{ $t('dashboard.todayOrders') }}</p>
            <p class="text-3xl font-black text-slate-800">{{ summary.today_orders }}</p>
          </div>
        </div>

        <!-- Low Stock Card -->
        <div class="bg-white p-6 rounded-3xl shadow-sm border border-slate-100 card-hover flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:scale-110 transition-transform">
            <AlertTriangle class="w-24 h-24 text-indigo-600" />
          </div>
          <div class="flex items-center justify-between mb-4">
            <div class="p-3 bg-red-50 rounded-2xl text-red-600">
              <AlertTriangle class="w-6 h-6" />
            </div>
            <span v-if="summary.low_stock_items > 0" class="bg-red-100 text-red-600 text-[10px] font-black px-2 py-1 rounded-full uppercase">Action Required</span>
          </div>
          <div>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">{{ $t('dashboard.lowStockAlerts') }}</p>
            <p class="text-3xl font-black text-slate-800">{{ summary.low_stock_items }}</p>
          </div>
        </div>

        <!-- Inventory Items Card -->
        <div class="bg-white p-6 rounded-3xl shadow-sm border border-slate-100 card-hover flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:scale-110 transition-transform">
            <Package class="w-24 h-24 text-indigo-600" />
          </div>
          <div class="flex items-center justify-between mb-4">
            <div class="p-3 bg-slate-100 rounded-2xl text-slate-600">
              <Package class="w-6 h-6" />
            </div>
          </div>
          <div>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">{{ $t('dashboard.totalIngredients') }}</p>
            <p class="text-3xl font-black text-slate-800">{{ summary.total_items }}</p>
          </div>
        </div>

        <!-- Tables Card -->
        <div class="bg-white p-6 rounded-3xl shadow-sm border border-slate-100 card-hover flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:scale-110 transition-transform">
            <Grid class="w-24 h-24 text-indigo-600" />
          </div>
          <div class="flex items-center justify-between mb-4">
            <div class="p-3 bg-yellow-50 rounded-2xl text-yellow-600">
              <Grid class="w-6 h-6" />
            </div>
          </div>
          <div>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">{{ $t('dashboard.tableOccupancy') }}</p>
            <p class="text-3xl font-black text-slate-800">{{ summary.occupied_tables }}<span class="text-lg text-slate-300 mx-1">/</span>{{ summary.total_tables }}</p>
          </div>
        </div>
      </div>

      <!-- Secondary Info Section -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <div class="bg-white rounded-3xl shadow-sm border border-slate-100 p-8">
          <div class="flex justify-between items-center mb-8">
            <h3 class="text-xl font-black text-slate-800 tracking-tight">{{ $t('dashboard.quickActions') }}</h3>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <router-link to="/pos" class="flex flex-col items-center gap-3 p-6 bg-indigo-600 text-white rounded-3xl hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100 group">
              <ShoppingCart class="w-8 h-8 group-hover:scale-110 transition-transform" />
              <span class="font-bold">{{ $t('dashboard.newOrder') }}</span>
            </router-link>
            <router-link to="/inventory" class="flex flex-col items-center gap-3 p-6 bg-white border-2 border-slate-100 text-slate-700 rounded-3xl hover:border-indigo-600 hover:text-indigo-600 transition-all group">
              <Package class="w-8 h-8 group-hover:scale-110 transition-transform" />
              <span class="font-bold">{{ $t('dashboard.addInventory') }}</span>
            </router-link>
          </div>
        </div>

        <div class="bg-indigo-900 rounded-3xl shadow-xl p-8 text-white relative overflow-hidden">
          <div class="absolute -right-10 -bottom-10 opacity-10">
            <ChefHat class="w-64 h-64" />
          </div>
          <div class="relative z-10">
            <h3 class="text-2xl font-black mb-2">{{ $t('dashboard.productionReady') }}</h3>
            <p class="text-indigo-200 mb-8 max-w-xs">{{ $t('dashboard.productionDesc') }}</p>
            <router-link to="/products" class="inline-flex items-center gap-2 bg-white text-indigo-900 px-6 py-3 rounded-2xl font-black hover:bg-indigo-50 transition-colors">
              {{ $t('dashboard.goToProduction') }}
              <ArrowRight class="w-5 h-5" />
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>
