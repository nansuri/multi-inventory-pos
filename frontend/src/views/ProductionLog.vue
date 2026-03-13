<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useI18n } from 'vue-i18n';
import { History, Calendar, ChefHat, ChevronDown, ChevronRight, PackageOpen } from 'lucide-vue-next';

const { t } = useI18n();
const logs = ref<any[]>([]);
const loading = ref(true);
const expandedLogs = ref<Record<number, boolean>>({});

const fetchLogs = async () => {
  try {
    const response = await api.get('/api/preparations');
    logs.value = response.data;
  } catch (error) {
    console.error('Failed to fetch production logs', error);
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
  expandedLogs.value[id] = !expandedLogs.value[id];
};

onMounted(fetchLogs);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8 animate-in fade-in duration-500">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <History class="w-8 h-8 text-indigo-600" />
            {{ t('productionLog.title') }}
          </h1>
          <p class="text-slate-500 font-medium mt-1">{{ t('productionLog.desc') }}</p>
        </div>
      </div>

      <!-- Production Log Table -->
      <div class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 border-b border-slate-100">
              <th class="w-10 px-4 py-6"></th>
              <th class="px-4 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ t('productionLog.timestamp') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ t('productionLog.product') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] text-center">{{ t('productionLog.output') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] text-right">{{ t('productionLog.consumption') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <template v-for="log in logs" :key="log.id">
              <tr 
                class="group hover:bg-slate-50/30 transition-colors cursor-pointer"
                @click="toggleExpand(log.id)"
              >
                <td class="px-4 py-6 text-center">
                  <component 
                    :is="expandedLogs[log.id] ? ChevronDown : ChevronRight" 
                    class="w-4 h-4 text-slate-300 group-hover:text-indigo-50 transition-colors"
                  />
                </td>
                <td class="px-4 py-6">
                  <div class="flex items-center gap-3">
                    <div class="p-2 bg-slate-100 rounded-lg text-slate-400">
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
                  <div class="inline-flex items-center gap-2 px-4 py-2 rounded-2xl bg-indigo-50 text-indigo-700 border border-indigo-100">
                    <PackageOpen class="w-4 h-4" />
                    <span class="text-[10px] font-black uppercase tracking-widest">{{ log.ingredients?.length || 0 }} {{ t('productionLog.itemsUsed') }}</span>
                  </div>
                </td>
              </tr>
              
              <!-- Expandable Ingredient List -->
              <tr v-if="expandedLogs[log.id]" class="bg-slate-50/50 animate-in slide-in-from-top-2 duration-200">
                <td colspan="5" class="px-20 py-6">
                  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    <div 
                      v-for="ing in log.ingredients" 
                      :key="ing.id"
                      class="flex items-center justify-between p-4 bg-white rounded-2xl border border-slate-100 shadow-sm"
                    >
                      <div class="flex flex-col">
                        <span class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">{{ t('common.inventory') }}</span>
                        <span class="font-black text-slate-800">{{ ing.ingredient_name }}</span>
                      </div>
                      <div class="text-right">
                        <span class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1 block">{{ t('productionLog.consumption') }}</span>
                        <span class="font-black text-indigo-600">{{ ing.quantity }} <span class="text-[10px] text-slate-400">{{ ing.unit }}</span></span>
                      </div>
                    </div>
                  </div>
                </td>
              </tr>
            </template>

            <tr v-if="logs.length === 0 && !loading">
              <td colspan="5" class="px-8 py-32 text-center">
                <div class="flex flex-col items-center gap-4 opacity-20">
                  <History class="w-20 h-20 text-slate-400" />
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
