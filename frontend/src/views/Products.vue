<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import ViewToggle from '../components/ViewToggle.vue';
import api from '../api/axios';
import { ChefHat, History, AlertCircle } from 'lucide-vue-next';
import { useRouter } from 'vue-router';
import { useConfigStore } from '../stores/config';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const products = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);
const configStore = useConfigStore();
useRouter();
const viewMode = ref<'list' | 'grid'>('grid');

const isPrepareModalOpen = ref(false);
const selectedProduct = ref<any>(null);
const prepareQty = ref(1);

const decrementQty = () => {
  if (prepareQty.value > 1) {
    prepareQty.value--;
  }
};

const alertConfig = ref({
  show: false,
  title: '',
  message: '',
  type: 'info' as 'info' | 'success' | 'warning' | 'danger'
});

const showAlert = (title: string, message: string, type: any = 'info') => {
  alertConfig.value = { show: true, title, message, type };
};

const fetchProducts = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/products');
    products.value = response.data;
  } catch (error) {
    console.error('Failed to fetch products', error);
  } finally {
    loading.value = false;
  }
};

const openPrepareModal = (product: any) => {
  selectedProduct.value = product;
  prepareQty.value = 1;
  isPrepareModalOpen.value = true;
};

const prepareProduct = async () => {
  actionLoading.value = true;
  try {
    await api.post(`/api/products/${selectedProduct.value.id}/prepare`, {
      pax: prepareQty.value
    });
    isPrepareModalOpen.value = false;
    await fetchProducts();
    showAlert(t('common.success'), `${prepareQty.value} ${t('common.pax')} ${selectedProduct.value.name} ${t('production.successCook')}`, "success");
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || "Failed to prepare product.", "danger");
  } finally {
    actionLoading.value = false;
  }
};

onMounted(fetchProducts);
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8 animate-in fade-in duration-500">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-2xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-3">
            <ChefHat class="w-8 h-8 text-indigo-600" />
            {{ t('production.hall') }}
          </h1>
          <p class="text-slate-500 dark:text-slate-400 font-medium mt-1">{{ t('production.cookRecipes') }}</p>
        </div>
        <div class="flex gap-3">
          <ViewToggle v-model:mode="viewMode" />
          <router-link to="/preparations" class="bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-700 dark:text-slate-200 px-6 py-4 rounded-2xl font-bold hover:bg-slate-50 dark:hover:bg-slate-800 transition-all shadow-sm flex items-center gap-2">
            <History class="w-5 h-5 text-slate-400 dark:text-slate-500" />
            {{ t('production.viewHistory') }}
          </router-link>
        </div>
      </div>

      <!-- List View -->
      <div v-if="viewMode === 'list'" class="bg-white dark:bg-slate-900 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('recipes.dishName') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('production.availableStock') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.price') }}</th>
              <th class="px-8 py-6 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="product in products" :key="product.id" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors">
              <td class="px-8 py-6 font-bold text-slate-800 dark:text-slate-100">{{ product.name }}</td>
              <td class="px-8 py-6 text-right font-black" :class="product.stock <= 5 ? 'text-red-600' : 'text-slate-700 dark:text-slate-300'">
                {{ product.stock }} <span class="text-[10px] text-slate-400 dark:text-slate-500 uppercase">{{ t('common.pax') }}</span>
              </td>
              <td class="px-8 py-6 text-right font-black text-slate-700 dark:text-slate-300">{{ configStore.formatCurrency(product.price) }}</td>
              <td class="px-8 py-6 text-right">
                <button @click="openPrepareModal(product)" class="bg-indigo-600 text-white px-4 py-2 rounded-xl text-xs font-black hover:bg-indigo-700 transition-all flex items-center gap-2 shadow-lg shadow-indigo-100 dark:shadow-indigo-900/30 ml-auto">
                  <ChefHat class="w-3 h-3" />
                  {{ t('production.cookNow') }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Grid View -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <div v-for="product in products" :key="product.id" class="bg-white dark:bg-slate-900 p-6 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-800 hover:shadow-xl transition-all duration-300 flex flex-col group relative overflow-hidden">
          <div class="flex justify-between items-start mb-6">
            <div class="p-4 bg-indigo-50 dark:bg-indigo-900/30 rounded-2xl text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
              <ChefHat class="w-6 h-6" />
            </div>
            <div class="text-right">
              <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest mb-1">{{ t('production.availableStock') }}</p>
              <div class="flex items-center justify-end gap-1"><span class="text-3xl font-black text-slate-800 dark:text-slate-100 tracking-tighter">{{ product.stock }}</span><span class="text-[10px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('common.pax') }}</span></div>
            </div>
          </div>

          <h3 class="text-xl font-black text-slate-800 dark:text-slate-100 mb-6 group-hover:text-indigo-600 transition-colors">{{ product.name }}</h3>
          
          <div class="mt-auto space-y-4">
            <div v-if="product.stock <= 5" class="flex items-center gap-2 text-orange-600 bg-orange-50 dark:bg-orange-950/30 p-3 rounded-xl border border-orange-100 dark:border-orange-900/50">
              <AlertCircle class="w-4 h-4" />
              <span class="text-[10px] font-black uppercase tracking-wider">{{ t('production.lowCookedStock') }}</span>
            </div>
            <button @click="openPrepareModal(product)" class="w-full py-4 bg-slate-900 dark:bg-indigo-600 text-white rounded-2xl font-black hover:bg-indigo-600 dark:hover:bg-indigo-700 transition-all flex items-center justify-center gap-2 shadow-xl shadow-slate-200 dark:shadow-none">
              <ChefHat class="w-5 h-5" />
              {{ t('production.cookNow') }}
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="products.length === 0 && !loading" class="py-32 text-center bg-white dark:bg-slate-900 rounded-[3rem] border-4 border-dashed border-slate-100 dark:border-slate-800">
        <div class="max-w-xs mx-auto space-y-4 opacity-40 dark:opacity-20">
          <ChefHat class="w-20 h-20 mx-auto text-slate-400 dark:text-slate-600" />
          <p class="text-xl font-black text-slate-500 dark:text-slate-400 uppercase">{{ t('common.noData') }}</p>
        </div>
      </div>
    </div>

    <!-- Prepare Modal -->
    <FormModal 
      :show="isPrepareModalOpen" 
      :title="t('production.cookingPrep')" 
      :subtitle="selectedProduct?.name"
      :icon="ChefHat" 
      :loading="actionLoading"
      maxWidth="max-w-md"
      @submit="prepareProduct" 
      @cancel="isPrepareModalOpen = false"
    >
      <div class="space-y-6">
        <p class="text-sm text-slate-500 dark:text-slate-400 leading-relaxed text-center font-medium">
          {{ t('production.recordDesc') }}
        </p>
        <div class="space-y-4">
          <label class="block text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('production.portionsToCook') }}</label>
          <div class="flex items-center justify-center gap-6">
            <button type="button" @click="decrementQty" class="w-14 h-14 rounded-2xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-xl font-black hover:bg-slate-200 dark:hover:bg-slate-700 dark:text-slate-100 transition-all">-</button>
            <span class="text-6xl font-black text-slate-800 dark:text-slate-100 w-24 text-center tracking-tighter">{{ prepareQty }}</span>
            <button type="button" @click="prepareQty++" class="w-14 h-14 rounded-2xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-xl font-black hover:bg-slate-200 dark:hover:bg-slate-700 dark:text-slate-100 transition-all">+</button>
          </div>
        </div>
      </div>
    </FormModal>

    <!-- Alert Modal -->
    <ConfirmModal 
      :show="alertConfig.show"
      :title="alertConfig.title"
      :message="alertConfig.message"
      :type="alertConfig.type"
      :showCancel="false"
      @confirm="alertConfig.show = false"
    />
  </DashboardLayout>
</template>
