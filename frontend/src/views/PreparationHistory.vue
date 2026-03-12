<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { History, Calendar, ChefHat, CheckCircle2 } from 'lucide-vue-next';

const preparations = ref<any[]>([]);
const loading = ref(true);

const fetchPreparations = async () => {
  try {
    const response = await api.get('/api/preparations');
    preparations.value = response.data;
  } catch (error) {
    console.error('Failed to fetch preparation logs', error);
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

onMounted(fetchPreparations);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8 animate-in fade-in duration-500">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <History class="w-8 h-8 text-indigo-600" />
            {{ $t('common.preparations') }}
          </h1>
          <p class="text-slate-500 font-medium mt-1">Historical log of all batch preparations and stock additions.</p>
        </div>
      </div>

      <!-- Preparation Table -->
      <div class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 border-b border-slate-100">
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Timestamp</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Dish Prepared</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] text-center">Output Volume</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] text-right">Inventory Impact</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="log in preparations" :key="log.id" class="group hover:bg-slate-50/30 transition-colors">
              <td class="px-8 py-6">
                <div class="flex items-center gap-3">
                  <div class="p-2 bg-slate-100 rounded-lg text-slate-400 group-hover:text-slate-600 transition-colors">
                    <Calendar class="w-4 h-4" />
                  </div>
                  <div class="flex flex-col">
                    <span class="font-bold text-slate-700 text-sm">{{ formatDate(log.created_at).date }}</span>
                    <span class="text-[10px] font-black text-slate-400 uppercase">{{ formatDate(log.created_at).time }}</span>
                  </div>
                </div>
              </td>
              <td class="px-8 py-6">
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 bg-orange-50 rounded-xl flex items-center justify-center text-orange-600 group-hover:bg-orange-600 group-hover:text-white transition-all">
                    <ChefHat class="w-5 h-5" />
                  </div>
                  <span class="font-black text-slate-800 tracking-tight">{{ log.product?.name }}</span>
                </div>
              </td>
              <td class="px-8 py-6 text-center">
                <span class="text-lg font-black text-slate-700">{{ log.pax }} <span class="text-xs font-bold text-slate-400 uppercase tracking-widest ml-1">Pax</span></span>
              </td>
              <td class="px-8 py-6 text-right">
                <div class="inline-flex items-center gap-2 px-4 py-2 rounded-2xl bg-green-50 text-green-700 border border-green-100">
                  <CheckCircle2 class="w-4 h-4" />
                  <span class="text-[10px] font-black uppercase tracking-widest">Raw Deducted</span>
                </div>
              </td>
            </tr>
            <tr v-if="preparations.length === 0 && !loading">
              <td colspan="4" class="px-8 py-32 text-center">
                <div class="flex flex-col items-center gap-4 opacity-20">
                  <History class="w-20 h-20 text-slate-400" />
                  <p class="text-xl font-black text-slate-500 uppercase tracking-[0.2em]">{{ $t('common.noData') }}</p>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </DashboardLayout>
</template>
