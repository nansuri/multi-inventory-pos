<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { 
  Package, 
  BookOpen, 
  ChefHat, 
  Utensils, 
  CreditCard, 
  LineChart,
  ArrowRight
} from 'lucide-vue-next';

const { t } = useI18n();

const steps = [
  {
    id: 1,
    title: t('common.inventory'),
    desc: t('guide.step1Desc'),
    usage: t('guide.step1Usage'),
    icon: Package,
    color: 'indigo'
  },
  {
    id: 2,
    title: t('common.recipes'),
    desc: t('guide.step2Desc'),
    usage: t('guide.step2Usage'),
    icon: BookOpen,
    color: 'blue'
  },
  {
    id: 3,
    title: t('common.production'),
    desc: t('guide.step3Desc'),
    usage: t('guide.step3Usage'),
    icon: ChefHat,
    color: 'orange'
  },
  {
    id: 4,
    title: t('dashboard.newOrder'),
    desc: t('guide.step4Desc'),
    usage: t('guide.step4Usage'),
    icon: Utensils,
    color: 'green'
  },
  {
    id: 5,
    title: t('pos.processPayment'),
    desc: t('guide.step5Desc'),
    usage: t('guide.step5Usage'),
    icon: CreditCard,
    color: 'emerald'
  },
  {
    id: 6,
    title: t('common.reports'),
    desc: t('guide.step6Desc'),
    usage: t('guide.step6Usage'),
    icon: LineChart,
    color: 'rose'
  }
];
</script>

<template>
  <div class="max-w-5xl mx-auto space-y-10 animate-in fade-in duration-700 pb-10">
    <!-- Header -->
    <div class="text-center space-y-4">
      <div class="inline-flex items-center gap-2 px-4 py-2 bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 rounded-2xl border border-indigo-100 dark:border-indigo-900/50">
        <BookOpen class="w-4 h-4" />
        <span class="text-[10px] font-black uppercase tracking-[0.2em]">{{ t('common.guide') }}</span>
      </div>
      <h1 class="text-4xl font-[1000] text-slate-800 dark:text-slate-100 tracking-tighter">{{ t('guide.title') }}</h1>
      <p class="text-slate-500 dark:text-slate-400 font-medium max-w-2xl mx-auto text-sm">{{ t('guide.desc') }}</p>
    </div>

    <!-- Core Flow -->
    <div class="relative">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 relative z-10">
        <div v-for="step in steps" :key="step.id" class="bg-white dark:bg-slate-900 p-6 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all group">
          <div class="flex justify-between items-start mb-6">
            <div :class="[
              'p-3.5 rounded-2xl text-white shadow-lg transition-all group-hover:scale-110 duration-500',
              step.color === 'indigo' ? 'bg-indigo-600 shadow-indigo-200 dark:shadow-none' :
              step.color === 'blue' ? 'bg-blue-600 shadow-blue-200 dark:shadow-none' :
              step.color === 'orange' ? 'bg-orange-500 shadow-orange-200 dark:shadow-none' :
              step.color === 'green' ? 'bg-green-600 shadow-green-200 dark:shadow-none' :
              step.color === 'emerald' ? 'bg-emerald-600 shadow-emerald-200 dark:shadow-none' : 'bg-rose-600 shadow-rose-200 dark:shadow-none'
            ]">
              <component :is="step.icon" class="w-5 h-5" />
            </div>
            <span class="text-3xl font-[1000] text-slate-100 dark:text-slate-800 select-none">{{ step.id }}</span>
          </div>

          <h3 class="text-lg font-black text-slate-800 dark:text-slate-100 mb-2">{{ step.title }}</h3>
          <p class="text-slate-500 dark:text-slate-400 text-xs leading-relaxed mb-4 font-medium">{{ step.desc }}</p>
          
          <div class="pt-4 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between">
            <p class="text-[9px] font-black text-indigo-600 dark:text-indigo-400 uppercase tracking-widest">{{ step.usage }}</p>
            <ArrowRight class="w-3.5 h-3.5 text-slate-300 dark:text-slate-700 group-hover:translate-x-1 transition-transform" />
          </div>
        </div>
      </div>
    </div>

    <!-- Security Info -->
    <div class="bg-slate-900 rounded-[2.5rem] p-8 md:p-12 text-white relative overflow-hidden">
      <div class="absolute right-[-5%] top-[-10%] w-[40%] h-[60%] bg-indigo-600/20 rounded-full blur-[100px]"></div>
      
      <div class="relative z-10 flex flex-col lg:flex-row items-center gap-10">
        <div class="flex-1 space-y-4">
          <h2 class="text-2xl font-[1000] tracking-tighter">{{ t('guide.securityTitle') }}</h2>
          <p class="text-slate-400 leading-relaxed font-medium text-sm">{{ t('guide.securityDesc') }}</p>
          <div class="grid grid-cols-2 gap-3">
            <div class="p-3 bg-white/5 rounded-xl border border-white/10">
              <p class="text-indigo-400 font-black uppercase tracking-widest text-[8px] mb-0.5">{{ t('guide.manageAccess') }}</p>
              <p class="text-[10px] font-bold text-slate-300">{{ t('guide.roleControl') }}</p>
            </div>
            <div class="p-3 bg-white/5 rounded-xl border border-white/10">
              <p class="text-indigo-400 font-black uppercase tracking-widest text-[8px] mb-0.5">{{ t('guide.dataIsolation') }}</p>
              <p class="text-[10px] font-bold text-slate-300">{{ t('guide.stockTracking') }}</p>
            </div>
          </div>
        </div>
        
        <div class="w-full lg:w-1/3">
          <div class="bg-white/5 backdrop-blur-md rounded-[2rem] p-6 border border-white/10 space-y-4">
            <div class="flex items-center gap-3">
              <div class="w-2 h-2 rounded-full bg-red-500"></div>
              <div class="w-2 h-2 rounded-full bg-yellow-500"></div>
              <div class="w-2 h-2 rounded-full bg-green-500"></div>
            </div>
            <div class="space-y-3">
              <div class="h-1.5 w-full bg-white/10 rounded-full"></div>
              <div class="h-1.5 w-3/4 bg-white/10 rounded-full"></div>
              <div class="h-1.5 w-1/2 bg-white/10 rounded-full"></div>
              <div class="pt-2">
                <p class="text-[9px] font-bold text-slate-500 dark:text-indigo-300 uppercase tracking-widest">{{ t('guide.paymentFlow') }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
