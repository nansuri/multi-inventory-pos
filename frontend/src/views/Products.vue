<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import ViewToggle from '../components/ViewToggle.vue';
import api from '../api/axios';
import { ChefHat, History, Package, AlertCircle, Trash2, ArrowRight } from 'lucide-vue-next';
import { useRouter } from 'vue-router';
import { useConfigStore } from '../stores/config';

const products = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);
const configStore = useConfigStore();
const router = useRouter();
const viewMode = ref<'list' | 'grid'>('list');

const isPrepareModalOpen = ref(false);
const isResetModalOpen = ref(false);
const selectedProduct = ref<any>(null);
const preparePax = ref(1);

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
  preparePax.value = 1;
  isPrepareModalOpen.value = true;
};

const submitPrepare = async () => {
  actionLoading.value = true;
  try {
    await api.post(`/api/products/${selectedProduct.value.id}/prepare`, {
      pax: preparePax.value
    });
    isPrepareModalOpen.value = false;
    await fetchProducts();
    showAlert("Batch Prepared", `${preparePax.value} pax of ${selectedProduct.value.name} has been added to stock.`, "success");
  } catch (error: any) {
    showAlert("Preparation Failed", error.response?.data?.error || 'Failed to record preparation', "danger");
  } finally {
    actionLoading.value = false;
  }
};

const resetStock = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/products/${selectedProduct.value.id}`, {
      ...selectedProduct.value,
      stock: 0
    });
    isResetModalOpen.value = false;
    await fetchProducts();
    showAlert("Stock Reset", `Available stock for ${selectedProduct.value.name} is now 0.`, "info");
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
        <div class="flex flex-1 gap-4 items-center">
          <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <ChefHat class="w-8 h-8 text-indigo-600" />
            {{ $t('production.hall') }}
          </h1>
          <ViewToggle v-model:mode="viewMode" />
        </div>
        <button 
          @click="router.push('/preparations')"
          class="flex items-center gap-2 bg-white border border-slate-200 text-slate-700 px-6 py-3 rounded-2xl font-bold hover:bg-slate-50 transition-all shadow-sm group"
        >
          <History class="w-5 h-5 text-slate-400 group-hover:text-slate-600 transition-colors" />
          <span>{{ $t('production.viewHistory') }}</span>
        </button>
      </div>

      <!-- List View -->
      <div v-if="viewMode === 'list'" class="bg-white rounded-3xl shadow-sm border border-slate-100 overflow-hidden animate-in fade-in duration-300">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 border-b border-slate-100">
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">Product Name</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">Available Stock</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">Selling Price</th>
              <th class="px-8 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="product in products" :key="product.id" class="group hover:bg-slate-50/50 transition-colors">
              <td class="px-8 py-5">
                <div class="flex items-center gap-3">
                  <div class="p-2 bg-orange-50 text-orange-600 rounded-lg"><ChefHat class="w-4 h-4" /></div>
                  <span class="font-bold text-slate-800">{{ product.name }}</span>
                </div>
              </td>
              <td class="px-8 py-5 text-right font-black" :class="product.stock <= 5 ? 'text-orange-600' : 'text-slate-800'">
                {{ product.stock }} <span class="text-[10px] text-slate-400 uppercase">Pax</span>
              </td>
              <td class="px-8 py-5 text-right font-black text-slate-700">{{ configStore.formatCurrency(product.price) }}</td>
              <td class="px-8 py-5 text-right">
                <div class="flex justify-end gap-2">
                  <button @click="openPrepareModal(product)" class="bg-indigo-600 text-white px-4 py-2 rounded-xl text-xs font-black hover:bg-indigo-700 transition-all flex items-center gap-2 shadow-lg shadow-indigo-100"><ChefHat class="w-3 h-3" />Cook</button>
                  <button @click="selectedProduct = product; isResetModalOpen = true" class="p-2 text-slate-400 hover:text-red-600 hover:bg-red-50 rounded-xl transition-all" title="Reset Stock"><Trash2 class="w-4 h-4" /></button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Grid View -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8 animate-in fade-in duration-300">
        <div v-for="product in products" :key="product.id" class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 flex flex-col group hover:shadow-xl transition-all duration-300 overflow-hidden">
          <div class="h-2 w-full" :class="product.stock <= 5 ? 'bg-orange-400' : 'bg-green-400'"></div>
          <div class="p-8 pt-6 flex flex-col flex-1">
            <div class="flex justify-between items-start mb-6">
              <div class="p-4 bg-orange-50 rounded-2xl text-orange-600 group-hover:bg-orange-600 group-hover:text-white transition-all"><ChefHat class="w-6 h-6" /></div>
              <div class="text-right">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('production.availableStock') }}</p>
                <div class="flex items-center justify-end gap-1"><span class="text-3xl font-black text-slate-800 tracking-tighter">{{ product.stock }}</span><span class="text-xs font-bold text-slate-400 uppercase">Pax</span></div>
              </div>
            </div>
            <h3 class="text-xl font-black text-slate-800 mb-1 group-hover:text-indigo-600 transition-colors">{{ product.name }}</h3>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-8">{{ $t('recipes.sellingPrice') }}: {{ configStore.formatCurrency(product.price) }}</p>
            <div class="mt-auto space-y-4">
              <div v-if="product.stock <= 5" class="flex items-center gap-2 text-orange-600 bg-orange-50 p-3 rounded-xl border border-orange-100"><AlertCircle class="w-4 h-4" /><span class="text-[10px] font-black uppercase tracking-wider">Low Cooked Stock</span></div>
              <div class="flex gap-2">
                <button @click="openPrepareModal(product)" class="flex-[3] py-4 bg-slate-900 text-white rounded-2xl font-black hover:bg-orange-600 transition-all flex items-center justify-center gap-3 shadow-xl shadow-slate-200"><ChefHat class="w-5 h-5" />{{ $t('production.cookNow') }}</button>
                <button @click="selectedProduct = product; isResetModalOpen = true" class="flex-1 py-4 bg-slate-50 text-slate-400 rounded-2xl flex items-center justify-center hover:bg-red-50 hover:text-red-600 transition-all" title="Reset Stock"><Trash2 class="w-5 h-5" /></button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="products.length === 0 && !loading" class="col-span-full py-32 text-center bg-white rounded-[3rem] border-4 border-dashed border-slate-100">
        <div class="max-w-xs mx-auto space-y-4 opacity-40">
          <Package class="w-20 h-20 mx-auto text-slate-400" />
          <p class="text-xl font-black text-slate-500 uppercase">{{ $t('common.noData') }}</p>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <FormModal :show="isPrepareModalOpen" :title="$t('production.cookingPrep')" :subtitle="selectedProduct?.name" :icon="ChefHat" iconClass="bg-orange-50 text-orange-600" :loading="actionLoading" maxWidth="max-w-sm" submitText="Record & Add Stock" @submit="submitPrepare" @cancel="isPrepareModalOpen = false">
      <div class="space-y-8 py-4">
        <p class="text-sm text-slate-500 leading-relaxed text-center">Recording this batch will automatically deduct all required <span class="font-bold text-slate-800 italic">raw ingredients</span> from your inventory.</p>
        <div class="space-y-3 text-center">
          <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('production.portionsToCook') }}</label>
          <div class="flex items-center justify-center gap-4">
            <button type="button" @click="preparePax = Math.max(1, preparePax - 1)" class="w-12 h-12 rounded-xl border-2 border-slate-100 flex items-center justify-center text-slate-400 hover:border-orange-500 hover:text-orange-600 transition-all">-</button>
            <input v-model.number="preparePax" type="number" min="1" required class="w-24 text-center text-4xl font-black text-slate-800 bg-transparent border-none outline-none focus:ring-0" />
            <button type="button" @click="preparePax++" class="w-12 h-12 rounded-xl border-2 border-slate-100 flex items-center justify-center text-slate-400 hover:border-orange-500 hover:text-orange-600 transition-all">+</button>
          </div>
        </div>
      </div>
    </FormModal>

    <ConfirmModal :show="isResetModalOpen" type="warning" title="Reset Cooked Stock" :message="`Are you sure you want to reset the available stock for ${selectedProduct?.name} to 0? This will NOT refund your raw ingredients.`" confirmText="Reset to Zero" :loading="actionLoading" @confirm="resetStock" @cancel="isResetModalOpen = false" />

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

<style scoped>
/* Hide number input spinners */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
input[type=number] {
  -moz-appearance: textfield;
}
</style>
