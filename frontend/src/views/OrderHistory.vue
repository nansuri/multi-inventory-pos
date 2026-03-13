<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useAuthStore } from '../stores/auth';
import { useI18n } from 'vue-i18n';
import { LineChart, Calendar, DollarSign, Utensils, ChevronDown, ChevronRight, User } from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
const logs = ref<any[]>([]);
const totalRevenue = ref(0);
const orderCount = ref(0);
const loading = ref(true);
const period = ref('day');
const configStore = useConfigStore();
const expandedOrders = ref<Record<number, boolean>>({});

const fetchHistory = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/reports/orders', {
      params: { period: period.value }
    });
    logs.value = response.data.orders || [];
    totalRevenue.value = response.data.total_revenue || 0;
    orderCount.value = response.data.order_count || 0;
  } catch (error) {
    console.error('Failed to fetch history', error);
  } finally {
    loading.value = false;
  }
};

const toggleExpand = (id: number) => {
  expandedOrders.value[id] = !expandedOrders.value[id];
};

watch(period, fetchHistory);
onMounted(fetchHistory);

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, fetchHistory);
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <LineChart class="w-6 h-6 text-indigo-600" />
          {{ t('reports.title') }}
        </h1>
        <p class="text-slate-500 dark:text-slate-400 font-medium text-xs mt-0.5">{{ t('reports.desc') }}</p>
      </div>

      <div class="flex items-center bg-white dark:bg-slate-900 p-1 rounded-xl shadow-sm border border-slate-200 dark:border-slate-800">
        <button 
          v-for="p in ['day', 'week', 'month']" 
          :key="p"
          @click="period = p"
          :class="['px-4 py-1.5 rounded-lg text-[9px] font-black uppercase tracking-widest transition-all', 
            period === p ? 'bg-indigo-600 text-white shadow-md' : 'text-slate-400 hover:text-slate-600 dark:hover:text-slate-200']"
        >
          {{ t(`common.${p}`) }}
        </button>
      </div>
    </div>

    <!-- Stats -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div class="bg-white dark:bg-slate-900 p-6 rounded-[2rem] shadow-sm border border-slate-200 dark:border-slate-800 relative overflow-hidden group">
        <div class="absolute right-[-10px] top-[-10px] opacity-[0.03] group-hover:scale-110 transition-transform duration-700">
          <DollarSign class="w-32 h-32" />
        </div>
        <p class="text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">{{ t('reports.totalRevenue') }}</p>
        <h3 class="text-2xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tight">{{ configStore.formatCurrency(totalRevenue) }}</h3>
      </div>

      <div class="bg-indigo-600 p-6 rounded-[2rem] shadow-lg shadow-indigo-100 dark:shadow-none text-white relative overflow-hidden group">
        <div class="absolute right-[-10px] top-[-10px] opacity-10 group-hover:scale-110 transition-transform duration-700">
          <Utensils class="w-32 h-32" />
        </div>
        <p class="text-[9px] font-black text-indigo-200 uppercase tracking-widest mb-1">{{ t('reports.ordersCompleted') }}</p>
        <h3 class="text-2xl font-[1000] tracking-tight">{{ orderCount }}</h3>
      </div>
    </div>

    <!-- History Table -->
    <div class="bg-white dark:bg-slate-900 rounded-3xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden text-sm">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('reports.transaction') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('reports.customer') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('reports.amount') }}</th>
            <th class="w-16"></th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <template v-for="order in logs" :key="order.id">
            <tr 
              @click="toggleExpand(order.id)"
              class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors cursor-pointer"
            >
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-9 h-9 rounded-xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-slate-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
                    <Calendar class="w-4 h-4" />
                  </div>
                  <div class="flex flex-col">
                    <span class="font-black text-slate-800 dark:text-slate-100 uppercase text-[10px] tracking-wider">#{{ order.id }}</span>
                    <span class="text-[8px] font-bold text-slate-400 dark:text-slate-500 uppercase leading-tight">{{ new Date(order.created_at).toLocaleTimeString([], {hour:'2-digit', minute:'2-digit'}) }}</span>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex flex-col">
                  <div class="flex items-center gap-1.5 text-xs">
                    <User class="w-3 h-3 text-indigo-500" />
                    <span class="font-bold text-slate-700 dark:text-slate-300">{{ order.customer_name || 'Walk-in' }}</span>
                  </div>
                  <span v-if="order.table" class="text-[8px] font-black text-slate-400 uppercase tracking-widest ml-4">Table {{ order.table.table_number }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-right">
                <span class="text-sm font-black text-slate-800 dark:text-slate-100">{{ configStore.formatCurrency(order.total_price) }}</span>
              </td>
              <td class="px-4 py-4 text-center">
                <ChevronDown :class="['w-4 h-4 text-slate-300 transition-transform duration-300', expandedOrders[order.id] ? 'rotate-180' : '']" />
              </td>
            </tr>
            <!-- Expanded Items -->
            <tr v-if="expandedOrders[order.id]" class="bg-slate-50/30 dark:bg-slate-800/20">
              <td colspan="4" class="px-6 py-4">
                <div class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 p-4 shadow-inner">
                  <h4 class="text-[8px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-[0.2em] mb-3 flex items-center gap-1.5">
                    <ChevronRight class="w-2.5 h-2.5 text-indigo-500" />
                    {{ t('reports.orderItems') }}
                  </h4>
                  <div class="space-y-2">
                    <div v-for="item in order.items" :key="item.id" class="flex justify-between items-center pb-2 border-b border-slate-50 dark:border-slate-800 last:border-0 last:pb-0">
                      <div class="flex items-center gap-2">
                        <span class="text-[9px] font-black text-indigo-600 dark:text-indigo-400 w-6">{{ item.quantity }}x</span>
                        <span class="font-bold text-slate-700 dark:text-slate-200 text-xs">{{ item.product?.name }}</span>
                      </div>
                      <span class="font-black text-slate-400 text-[10px]">{{ configStore.formatCurrency(item.price * item.quantity) }}</span>

                    </div>
                  </div>
                </div>
              </td>
            </tr>
          </template>
          <tr v-if="logs.length === 0 && !loading">
            <td colspan="4" class="px-6 py-16 text-center text-slate-400 italic">
              {{ t('common.noData') }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
