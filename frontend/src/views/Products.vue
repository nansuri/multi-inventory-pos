<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { ChefHat, History, ArrowRight, Package, AlertCircle } from 'lucide-vue-next';
import { useRouter } from 'vue-router';
import { useConfigStore } from '../stores/config';

const products = ref<any[]>([]);
const loading = ref(true);
const configStore = useConfigStore();
const isPrepareModalOpen = ref(false);
const selectedProduct = ref<any>(null);
const preparePax = ref(1);
const router = useRouter();

const fetchProducts = async () => {
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
  try {
    await api.post(`/api/products/${selectedProduct.value.id}/prepare`, {
      pax: preparePax.value
    });
    isPrepareModalOpen.value = false;
    await fetchProducts();
    alert('Preparation recorded. Stock updated.');
  } catch (error: any) {
    alert(error.response?.data?.error || 'Failed to record preparation');
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
          <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <ChefHat class="w-8 h-8 text-indigo-600" />
            {{ $t('production.hall') }}
          </h1>
          <p class="text-slate-500 font-medium mt-1">{{ $t('production.cookRecipes') }}</p>
        </div>
        <button 
          @click="router.push('/preparations')"
          class="flex items-center gap-2 bg-white border border-slate-200 text-slate-700 px-6 py-3 rounded-2xl font-bold hover:bg-slate-50 transition-all shadow-sm group"
        >
          <History class="w-5 h-5 text-slate-400 group-hover:text-slate-600 transition-colors" />
          <span>{{ $t('production.viewHistory') }}</span>
        </button>
      </div>

      <!-- Production Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
        <div 
          v-for="product in products" 
          :key="product.id" 
          class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 flex flex-col group hover:shadow-xl transition-all duration-300 overflow-hidden"
        >
          <!-- Status Bar -->
          <div 
            class="h-2 w-full" 
            :class="product.stock <= 5 ? 'bg-orange-400' : 'bg-green-400'"
          ></div>

          <div class="p-8 pt-6 flex flex-col flex-1">
            <div class="flex justify-between items-start mb-6">
              <div class="p-4 bg-orange-50 rounded-2xl text-orange-600 group-hover:bg-orange-600 group-hover:text-white transition-all">
                <ChefHat class="w-6 h-6" />
              </div>
              <div class="text-right">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('production.availableStock') }}</p>
                <div class="flex items-center justify-end gap-1">
                  <span class="text-3xl font-black text-slate-800 tracking-tighter">{{ product.stock }}</span>
                  <span class="text-xs font-bold text-slate-400 uppercase">Pax</span>
                </div>
              </div>
            </div>
            
            <h3 class="text-xl font-black text-slate-800 mb-1 group-hover:text-indigo-600 transition-colors">{{ product.name }}</h3>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-8">{{ $t('recipes.sellingPrice') }}: {{ configStore.formatCurrency(product.price) }}</p>

            <div class="mt-auto space-y-4">
              <!-- Stock Warning -->
              <div v-if="product.stock <= 5" class="flex items-center gap-2 text-orange-600 bg-orange-50 p-3 rounded-xl border border-orange-100">
                <AlertCircle class="w-4 h-4" />
                <span class="text-[10px] font-black uppercase tracking-wider">Low Cooked Stock</span>
              </div>

              <button 
                @click="openPrepareModal(product)" 
                class="w-full py-4 bg-slate-900 text-white rounded-2xl font-black hover:bg-orange-600 transition-all flex items-center justify-center gap-3 shadow-xl shadow-slate-200 group/btn"
              >
                <ChefHat class="w-5 h-5 group-hover/btn:scale-110 transition-transform" />
                {{ $t('production.cookNow') }}
              </button>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-if="products.length === 0 && !loading" class="col-span-full py-32 text-center bg-white rounded-[3rem] border-4 border-dashed border-slate-100">
          <div class="max-w-xs mx-auto space-y-4 opacity-40">
            <Package class="w-20 h-20 mx-auto text-slate-400" />
            <p class="text-xl font-black text-slate-500 uppercase">{{ $t('common.noData') }}</p>
            <p class="text-sm font-medium text-slate-400">You need to create recipes before you can start production.</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Prepare Product Modal -->
    <div v-if="isPrepareModalOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center p-4 z-50 animate-in fade-in duration-200">
      <div class="bg-white rounded-[2.5rem] w-full max-w-sm p-10 shadow-2xl scale-in-center">
        <div class="flex items-center gap-4 mb-8">
          <div class="p-4 bg-orange-100 text-orange-600 rounded-2xl">
            <ChefHat class="w-8 h-8" />
          </div>
          <div>
            <h3 class="text-2xl font-black text-slate-800 leading-tight">{{ $t('production.cookingPrep') }}</h3>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">{{ selectedProduct?.name }}</p>
          </div>
        </div>
        
        <p class="text-sm text-slate-500 mb-10 leading-relaxed">
          Recording this batch will automatically deduct all required <span class="font-bold text-slate-800 italic">raw ingredients</span> from your inventory.
        </p>
        
        <form @submit.prevent="submitPrepare" class="space-y-8">
          <div class="space-y-3 text-center">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('production.portionsToCook') }}</label>
            <div class="flex items-center justify-center gap-4">
              <button type="button" @click="preparePax = Math.max(1, preparePax - 1)" class="w-12 h-12 rounded-xl border-2 border-slate-100 flex items-center justify-center text-slate-400 hover:border-orange-500 hover:text-orange-600 transition-all">-</button>
              <input v-model.number="preparePax" type="number" min="1" required autofocus class="w-24 text-center text-4xl font-black text-slate-800 bg-transparent border-none outline-none focus:ring-0" />
              <button type="button" @click="preparePax++" class="w-12 h-12 rounded-xl border-2 border-slate-100 flex items-center justify-center text-slate-400 hover:border-orange-500 hover:text-orange-600 transition-all">+</button>
            </div>
          </div>

          <div class="flex flex-col gap-2 pt-4">
            <button type="submit" class="w-full py-5 bg-orange-600 text-white rounded-3xl font-black shadow-lg shadow-orange-100 hover:bg-orange-700 transition-all">
              {{ $t('production.recordAdd') }}
            </button>
            <button type="button" @click="isPrepareModalOpen = false" class="w-full py-4 text-slate-400 font-bold hover:text-slate-600 transition-colors">
              {{ $t('common.cancel') }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </DashboardLayout>
</template>

<style scoped>
.scale-in-center {
  animation: scale-in-center 0.3s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}

@keyframes scale-in-center {
  0% { transform: scale(0.9); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}

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
