<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useI18n } from 'vue-i18n';
import { LineChart, Calendar, DollarSign, Utensils, ChevronDown, ChevronRight, User } from 'lucide-vue-next';

const { t } = useI18n();
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
    const response = await api.get(`/api/reports/orders?period=${period.value}`);
    logs.value = response.data.orders;
    totalRevenue.value = response.data.total_revenue;
    orderCount.value = response.data.order_count;
  } catch (error) {
    console.error('Failed to fetch order history', error);
  } finally {
    loading.value = false;
  }
};

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr);
  return {
    full: date.toLocaleString(),
    time: date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
    date: date.toLocaleDateString([], { month: 'short', day: 'numeric', year: 'numeric' })
  };
};

const toggleExpand = (id: number) => {
  expandedOrders.value[id] = !expandedOrders.value[id];
};

watch(period, fetchHistory);
onMounted(fetchHistory);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-10 animate-in fade-in duration-500">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
        <div>
          <h1 class="text-3xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <LineChart class="w-10 h-10 text-indigo-600" />
            {{ t('reports.title') }}
          </h1>
          <p class="text-slate-500 font-medium mt-1">{{ t('reports.desc') }}</p>
        </div>

        <div class="flex items-center bg-white p-1.5 rounded-2xl shadow-sm border border-slate-100">
          <button 
            v-for="p in ['day', 'week', 'month']" 
            :key="p"
            @click="period = p"
            class="px-6 py-2.5 rounded-xl text-xs font-black uppercase tracking-widest transition-all"
            :class="period === p ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-100' : 'text-slate-400 hover:text-slate-600'"
          >
            {{ t(`common.${p}`) }}
          </button>
        </div>
      </div>

      <!-- Summary Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white p-8 rounded-[2.5rem] shadow-sm border border-slate-100 flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-10 opacity-5">
            <DollarSign class="w-32 h-32 text-indigo-600" />
          </div>
          <div class="flex items-center justify-between mb-6">
            <div class="p-4 bg-green-50 rounded-2xl text-green-600 font-black">
              <DollarSign class="w-8 h-8" />
            </div>
          </div>
          <div>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-[0.2em] mb-2">{{ t('reports.totalRevenue') }} ({{ t(`common.${period}`) }})</p>
            <p class="text-5xl font-black text-slate-800">{{ configStore.formatCurrency(totalRevenue) }}</p>
          </div>
        </div>

        <div class="bg-indigo-900 p-8 rounded-[2.5rem] shadow-xl text-white flex flex-col justify-between relative overflow-hidden group">
          <div class="absolute top-0 right-0 p-10 opacity-10">
            <Utensils class="w-32 h-32" />
          </div>
          <div class="flex items-center justify-between mb-6">
            <div class="p-4 bg-white/10 rounded-2xl text-white font-black border border-white/10">
              <Utensils class="w-8 h-8" />
            </div>
          </div>
          <div>
            <p class="text-xs font-bold text-indigo-300 uppercase tracking-[0.2em] mb-2">{{ t('reports.ordersCompleted') }}</p>
            <p class="text-5xl font-black">{{ orderCount }}</p>
          </div>
        </div>
      </div>

      <!-- Detailed Table -->
      <div class="bg-white rounded-[3rem] shadow-sm border border-slate-100 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 border-b border-slate-100">
              <th class="w-10 px-4 py-6"></th>
              <th class="px-4 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ t('reports.transaction') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ t('reports.customer') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] text-center">{{ t('reports.table') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] text-right">{{ t('reports.amount') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <template v-for="order in logs" :key="order.id">
              <tr 
                class="group hover:bg-slate-50/30 transition-colors cursor-pointer"
                @click="toggleExpand(order.id)"
              >
                <td class="px-4 py-6 text-center">
                  <component :is="expandedOrders[order.id] ? ChevronDown : ChevronRight" class="w-4 h-4 text-slate-300 group-hover:text-indigo-50" />
                </td>
                <td class="px-4 py-6">
                  <div class="flex items-center gap-3">
                    <div class="p-2 bg-slate-100 rounded-lg text-slate-400">
                      <Calendar class="w-4 h-4" />
                    </div>
                    <div class="flex flex-col">
                      <span class="font-bold text-slate-700 text-sm">{{ formatDate(order.created_at).date }}</span>
                      <span class="text-[10px] font-black text-slate-400 uppercase">{{ formatDate(order.created_at).time }}</span>
                    </div>
                  </div>
                </td>
                <td class="px-8 py-6">
                  <div class="flex items-center gap-3">
                    <div class="w-8 h-8 rounded-full bg-slate-100 flex items-center justify-center text-slate-400">
                      <User class="w-4 h-4" />
                    </div>
                    <span class="font-black text-slate-800 tracking-tight">{{ order.customer_name }}</span>
                  </div>
                </td>
                <td class="px-8 py-6 text-center">
                  <span class="px-3 py-1 bg-slate-100 text-slate-500 text-[10px] font-black uppercase rounded-lg">T-{{ order.table?.table_number }}</span>
                </td>
                <td class="px-8 py-6 text-right">
                  <span class="font-black text-slate-800">{{ configStore.formatCurrency(order.total_price) }}</span>
                </td>
              </tr>

              <tr v-if="expandedOrders[order.id]" class="bg-slate-50/50 animate-in slide-in-from-top-2 duration-200">
                <td colspan="5" class="px-20 py-8">
                  <div class="bg-white rounded-3xl border border-slate-100 shadow-sm p-6 max-w-xl">
                    <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-4">{{ t('reports.orderItems') }}</p>
                    <div class="space-y-3">
                      <div v-for="item in order.items" :key="item.id" class="flex justify-between items-center text-sm">
                        <div class="flex items-center gap-3">
                          <span class="font-black text-indigo-600 bg-indigo-50 w-6 h-6 flex items-center justify-center rounded text-[10px]">{{ item.quantity }}x</span>
                          <span class="font-bold text-slate-700">{{ item.product?.name }}</span>
                        </div>
                        <span class="font-black text-slate-800">{{ configStore.formatCurrency(item.price * item.quantity) }}</span>
                      </div>
                    </div>
                  </div>
                </td>
              </tr>
            </template>

            <tr v-if="logs.length === 0 && !loading">
              <td colspan="5" class="px-8 py-32 text-center">
                <div class="flex flex-col items-center gap-4 opacity-20">
                  <LineChart class="w-20 h-20 text-slate-400" />
                  <p class="text-xl font-black text-slate-500 uppercase tracking-[0.2em]">{{ t('common.noData') }}</p>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </DashboardLayout>
</template>
