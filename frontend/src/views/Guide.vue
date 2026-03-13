<script setup lang="ts">
import DashboardLayout from '../layouts/DashboardLayout.vue';
import { useI18n } from 'vue-i18n';
import { 
  Package, 
  BookOpen, 
  ChefHat, 
  Utensils, 
  CreditCard, 
  LineChart,
  ArrowRight,
  ShieldCheck,
  CheckCircle2
} from 'lucide-vue-next';
import { computed } from 'vue';

const { t } = useI18n();

const steps = computed(() => [
  {
    title: t('guide.step1Title', '1. Inventory Setup'),
    icon: Package,
    color: 'text-blue-600',
    bg: 'bg-blue-50',
    desc: t('guide.step1Desc', 'Begin by adding your raw materials (e.g., Rice, Meat, Oil). Define their units (KG, GR, Litre) and set minimum stock alerts to avoid running out.'),
    usage: t('guide.step1Usage', 'Go to Kitchen & Storage > Inventory.')
  },
  {
    title: t('guide.step2Title', '2. Define Recipes'),
    icon: BookOpen,
    color: 'text-indigo-600',
    bg: 'bg-indigo-50',
    desc: t('guide.step2Desc', 'Create "Blueprints" for your dishes. Link your raw ingredients to a menu item and specify exactly how much of each is needed for 1 portion.'),
    usage: t('guide.step2Usage', 'Go to Kitchen & Storage > Recipe Management.')
  },
  {
    title: t('guide.step3Title', '3. Production (Cooking)'),
    icon: ChefHat,
    color: 'text-orange-600',
    bg: 'bg-orange-50',
    desc: t('guide.step3Desc', 'When the Chef cooks a batch, record it here. The system automatically deducts raw materials from Inventory and increases the "Ready to Sell" stock.'),
    usage: t('guide.step3Usage', 'Go to Kitchen & Storage > Production Hall.')
  },
  {
    title: t('guide.step4Title', '4. Take Orders'),
    icon: Utensils,
    color: 'text-pink-600',
    bg: 'bg-pink-50',
    desc: t('guide.step4Desc', 'Waiters select a table and add cooked products to the cart. This marks the table as "Occupied" and prepares the bill.'),
    usage: t('guide.step4Usage', 'Go to Sales & Floor > New Order.')
  },
  {
    title: t('guide.step5Title', '5. Process Payments'),
    icon: CreditCard,
    color: 'text-green-600',
    bg: 'bg-green-50',
    desc: t('guide.step5Desc', 'When the customer is ready to leave, process the payment. This finalizes the order, deducts the cooked product stock, and frees the table.'),
    usage: t('guide.step5Usage', 'Go to Sales & Floor > Payments.')
  },
  {
    title: t('guide.step6Title', '6. Review & Analytics'),
    icon: LineChart,
    color: 'text-purple-600',
    bg: 'bg-purple-50',
    desc: t('guide.step6Desc', 'Monitor daily revenue, order trends, and production history. Use these insights to optimize your stock and menu pricing.'),
    usage: t('guide.step6Usage', 'Check Dashboard or Administration > Order History.')
  }
]);
</script>

<template>
  <DashboardLayout>
    <div class="max-w-5xl mx-auto space-y-12 animate-in fade-in duration-700">
      <!-- Header -->
      <div class="text-center space-y-4">
        <h1 class="text-4xl font-black text-slate-800 tracking-tight">{{ t('guide.title') }}</h1>
        <p class="text-lg text-slate-500 max-w-2xl mx-auto font-medium">{{ t('guide.desc') }}</p>
      </div>

      <!-- Business Flow Diagram -->
      <div class="bg-white p-10 rounded-[3rem] shadow-sm border border-slate-100 relative overflow-hidden">
        <div class="absolute top-0 right-0 p-10 opacity-5">
          <ShieldCheck class="w-40 h-40 text-indigo-600" />
        </div>
        
        <h3 class="text-xl font-black text-slate-800 mb-10 flex items-center gap-2">
          <div class="w-2 h-8 bg-indigo-600 rounded-full"></div>
          {{ t('guide.coreFlow') }}
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 items-center">
          <div class="flex flex-col items-center text-center gap-3">
            <div class="w-16 h-16 rounded-2xl bg-slate-900 text-white flex items-center justify-center shadow-xl">
              <Package class="w-8 h-8" />
            </div>
            <span class="font-black text-xs uppercase tracking-widest text-slate-400">{{ t('common.inventory') }}</span>
          </div>
          <div class="hidden md:flex justify-center text-slate-200">
            <ArrowRight class="w-8 h-8" />
          </div>
          <div class="flex flex-col items-center text-center gap-3">
            <div class="w-16 h-16 rounded-2xl bg-orange-600 text-white flex items-center justify-center shadow-xl">
              <ChefHat class="w-8 h-8" />
            </div>
            <span class="font-black text-xs uppercase tracking-widest text-slate-400">{{ t('common.production') }}</span>
          </div>
          <div class="hidden md:flex justify-center text-slate-200">
            <ArrowRight class="w-8 h-8" />
          </div>
          <div class="flex flex-col items-center text-center gap-3">
            <div class="w-16 h-16 rounded-2xl bg-indigo-600 text-white flex items-center justify-center shadow-xl">
              <ShoppingCart class="w-8 h-8" />
            </div>
            <span class="font-black text-xs uppercase tracking-widest text-slate-400">Sales (POS)</span>
          </div>
          <div class="hidden md:flex justify-center text-slate-200">
            <ArrowRight class="w-8 h-8" />
          </div>
          <div class="flex flex-col items-center text-center gap-3">
            <div class="w-16 h-16 rounded-2xl bg-green-600 text-white flex items-center justify-center shadow-xl">
              <LineChart class="w-8 h-8" />
            </div>
            <span class="font-black text-xs uppercase tracking-widest text-slate-400">{{ t('common.reports') }}</span>
          </div>
        </div>
      </div>

      <!-- Step by Step Guide -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div 
          v-for="(step, index) in steps" 
          :key="index"
          class="bg-white p-8 rounded-[2.5rem] shadow-sm border border-slate-100 hover:shadow-xl transition-all duration-300 flex flex-col group"
        >
          <div class="flex items-start justify-between mb-6">
            <div :class="['p-4 rounded-2xl transition-all group-hover:scale-110', step.bg, step.color]">
              <component :is="step.icon" class="w-8 h-8" />
            </div>
            <span class="text-4xl font-black text-slate-50 italic opacity-50 group-hover:text-indigo-50 transition-colors">0{{ index + 1 }}</span>
          </div>
          
          <h4 class="text-xl font-black text-slate-800 mb-3">{{ step.title }}</h4>
          <p class="text-sm text-slate-500 font-medium leading-relaxed flex-1">{{ step.desc }}</p>
          
          <div class="mt-8 pt-6 border-t border-slate-50 flex items-center gap-2 text-indigo-600">
            <CheckCircle2 class="w-4 h-4" />
            <span class="text-[10px] font-black uppercase tracking-widest">{{ step.usage }}</span>
          </div>
        </div>
      </div>

      <!-- Security & Multi-tenancy Note -->
      <div class="bg-slate-900 rounded-[3rem] p-12 text-white relative overflow-hidden">
        <div class="relative z-10 grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
          <div class="space-y-6">
            <h3 class="text-3xl font-black tracking-tight">{{ t('guide.securityTitle') }}</h3>
            <p class="text-slate-400 font-medium leading-relaxed">
              {{ t('guide.securityDesc') }}
            </p>
            <router-link to="/employees" class="inline-flex items-center gap-2 bg-indigo-600 text-white px-8 py-4 rounded-2xl font-black hover:bg-indigo-700 transition-all shadow-xl shadow-indigo-900/20">
              {{ t('guide.manageAccess') }}
              <ArrowRight class="w-5 h-5" />
            </router-link>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div class="p-6 bg-white/5 rounded-3xl border border-white/10">
              <p class="text-2xl font-black mb-1">100%</p>
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">Data Isolation</p>
            </div>
            <div class="p-6 bg-white/5 rounded-3xl border border-white/10">
              <p class="text-2xl font-black mb-1">RBAC</p>
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">Role Control</p>
            </div>
            <div class="p-6 bg-white/5 rounded-3xl border border-white/10">
              <p class="text-2xl font-black mb-1">Live</p>
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">Stock Tracking</p>
            </div>
            <div class="p-6 bg-white/5 rounded-3xl border border-white/10">
              <p class="text-2xl font-black mb-1">Fast</p>
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">Payment Flow</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>
