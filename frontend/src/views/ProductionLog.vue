<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import api from '../api/axios';
import { useI18n } from 'vue-i18n';
import { useAuthStore } from '../stores/auth';
import { History, Calendar, ChefHat, ChevronDown, ChevronRight, PackageOpen } from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
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

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, fetchLogs);
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <History class="w-6 h-6 text-indigo-600" />
          {{ t('productionLog.title') }}
        </h1>
        <p class="text-slate-500 dark:text-slate-400 font-medium text-xs mt-0.5">{{ t('productionLog.desc') }}</p>
      </div>
    </div>

    <!-- Production Log Table -->
    <div class="bg-white dark:bg-slate-900 rounded-3xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden text-sm">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
            <th class="w-8 px-2 py-4"></th>
            <th class="px-2 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('productionLog.timestamp') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('productionLog.product') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('productionLog.output') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('productionLog.consumption') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <template v-for="log in logs" :key="log.id">
            <tr 
              class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors cursor-pointer"
              @click="toggleExpand(log.id)"
            >
              <td class="px-2 py-4 text-center">
                <component 
                  :is="expandedLogs[log.id] ? ChevronDown : ChevronRight" 
                  class="w-3.5 h-3.5 text-slate-300 dark:text-slate-600 transition-colors"
                />
              </td>
              <td class="px-2 py-4 whitespace-nowrap">
                <div class="flex items-center gap-2">
                  <div class="p-1.5 bg-slate-100 dark:bg-slate-800 rounded-lg text-slate-400">
                    <Calendar class="w-3.5 h-3.5" />
                  </div>
                  <div class="flex flex-col">
                    <span class="font-bold text-slate-700 dark:text-slate-200 text-xs">{{ formatDate(log.created_at).date }}</span>
                    <span class="text-[8px] font-black text-slate-400 dark:text-slate-500 uppercase leading-none">{{ formatDate(log.created_at).time }}</span>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 bg-orange-50 dark:bg-orange-950/30 rounded-lg flex items-center justify-center text-orange-600">
                    <ChefHat class="w-4 h-4" />
                  </div>
                  <span class="font-black text-slate-800 dark:text-slate-100 tracking-tight text-xs">{{ log.product?.name }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-center">
                <span class="text-sm font-black text-slate-700 dark:text-slate-200">{{ log.pax }} <span class="text-[8px] font-bold text-slate-400 uppercase tracking-widest ml-0.5">{{ t('common.pax') }}</span></span>
              </td>
              <td class="px-6 py-4 text-right">
                <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-xl bg-indigo-50 dark:bg-indigo-900/30 text-indigo-700 dark:text-indigo-400 border border-indigo-100 dark:border-indigo-900/50">
                  <PackageOpen class="w-3.5 h-3.5" />
                  <span class="text-[8px] font-black uppercase tracking-widest">{{ log.ingredients?.length || 0 }} items</span>
                </div>
              </td>
            </tr>
            
            <!-- Expandable Ingredient List -->
            <tr v-if="expandedLogs[log.id]" class="bg-slate-50/50 dark:bg-slate-800/20 animate-in slide-in-from-top-2 duration-200">
              <td colspan="5" class="px-12 py-4">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3">
                  <div 
                    v-for="ing in log.ingredients" 
                    :key="ing.id"
                    class="flex items-center justify-between p-3 bg-white dark:bg-slate-900 rounded-xl border border-slate-100 dark:border-slate-800 shadow-sm"
                  >
                    <div class="flex flex-col">
                      <span class="text-[8px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-0.5">{{ t('common.inventory') }}</span>
                      <span class="font-black text-slate-800 dark:text-slate-100 text-[11px]">{{ ing.ingredient_name }}</span>
                    </div>
                    <div class="text-right">
                      <span class="text-[8px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-0.5 block">{{ t('productionLog.consumption') }}</span>
                      <span class="font-black text-indigo-600 dark:text-indigo-400 text-[11px]">{{ ing.quantity }}{{ ing.unit }}</span>
                    </div>
                  </div>
                </div>
              </td>
            </tr>
          </template>

          <tr v-if="logs.length === 0 && !loading">
            <td colspan="5" class="px-6 py-20 text-center text-slate-400 italic">
              {{ t('common.noData') }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
