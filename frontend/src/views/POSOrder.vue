<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useAuthStore } from '../stores/auth';
import { ShoppingCart, CheckCircle, Users, Plus, Minus, UtensilsCrossed, Trash2, User, ChefHat } from 'lucide-vue-next';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const authStore = useAuthStore();
const tables = ref<any[]>([]);
const products = ref<any[]>([]);
const selectedTable = ref<any>(null);
const cart = ref<any[]>([]);
const customerName = ref('');
const loading = ref(true);
const actionLoading = ref(false);
const configStore = useConfigStore();
const router = useRouter();

const alertConfig = ref({
  show: false,
  title: '',
  message: '',
  type: 'info' as 'info' | 'success' | 'warning' | 'danger',
  confirmText: 'OK',
  callback: null as (() => void) | null
});

const showAlert = (title: string, message: string, type: any = 'info', confirmText = 'OK', callback: any = null) => {
  alertConfig.value = { show: true, title, message, type, confirmText, callback };
};

const handleAlertConfirm = () => {
  const cb = alertConfig.value.callback;
  alertConfig.value.show = false;
  if (cb) cb();
};

const fetchInitialData = async () => {
  try {
    const [tablesRes, productsRes] = await Promise.all([
      api.get('/api/tables'),
      api.get('/api/products')
    ]);
    tables.value = tablesRes.data;
    products.value = productsRes.data;
  } catch (error) {
    console.error('Failed to fetch POS data', error);
  } finally {
    loading.value = false;
  }
};

const addToCart = (product: any) => {
  const existing = cart.value.find(item => item.product_id === product.id);
  if (existing) {
    existing.quantity++;
  } else {
    cart.value.push({
      product_id: product.id,
      name: product.name,
      price: product.price,
      quantity: 1
    });
  }
};

const removeFromCart = (index: number) => {
  if (cart.value[index].quantity > 1) {
    cart.value[index].quantity--;
  } else {
    cart.value.splice(index, 1);
  }
};

const clearCart = () => {
  cart.value = [];
};

const totalPrice = () => cart.value.reduce((sum, item) => sum + (item.price * item.quantity), 0);

const placeOrder = async () => {
  if (!selectedTable.value || cart.value.length === 0) return;
  if (!customerName.value) {
    showAlert(t('common.error'), t('pos.customerBooker') + " required.", "warning");
    return;
  }
  
  actionLoading.value = true;
  try {
    const orderData = {
      table_id: selectedTable.value.id,
      customer_name: customerName.value,
      total_price: totalPrice(),
      items: cart.value.map(item => ({
        product_id: item.product_id,
        quantity: item.quantity,
        price: item.price
      }))
    };
    
    await api.post('/api/orders', orderData);
    
    showAlert(
      t('common.success'), 
      `${t('pos.orderPlaced')} ${customerName.value} (T-${selectedTable.value.table_number}).`, 
      "success",
      t('common.confirm'),
      () => {
        router.push('/pos/payment');
      }
    );

    cart.value = [];
    selectedTable.value = null;
    customerName.value = '';
    await fetchInitialData();
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || error.message, "danger");
  } finally {
    actionLoading.value = false;
  }
};

onMounted(fetchInitialData);

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, fetchInitialData);
</script>

<template>
  <div class="flex flex-col lg:flex-row gap-6 h-[calc(100vh-140px)] -mt-1">
    <!-- Main Selection Area -->
    <div class="flex-1 space-y-8 overflow-auto pr-2 custom-scrollbar">
      <!-- Tables Selection -->
      <section>
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
            <div class="p-1.5 bg-indigo-100 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 rounded-lg">
              <Users class="w-4 h-4" />
            </div>
            {{ t('pos.floorPlan') }}
          </h3>
        </div>
        
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 xl:grid-cols-6 gap-3">
          <button 
            v-for="table in tables" 
            :key="table.id"
            @click="selectedTable = table"
            :class="[
              'p-4 rounded-2xl border-2 transition-all duration-300 relative group flex flex-col items-center justify-center gap-0.5',
              table.status === 'occupied' 
                ? 'bg-slate-50 dark:bg-slate-900 border-slate-100 dark:border-slate-800 cursor-not-allowed opacity-60' 
                : selectedTable?.id === table.id 
                  ? 'bg-indigo-600 border-indigo-600 text-white shadow-lg shadow-indigo-200 dark:shadow-none' 
                  : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 hover:border-indigo-300 dark:hover:border-indigo-500 hover:shadow-md'
            ]"
            :disabled="table.status === 'occupied'"
          >
            <span class="text-[9px] font-black uppercase tracking-tighter transition-colors" :class="selectedTable?.id === table.id ? 'text-indigo-100' : 'text-slate-400 dark:text-slate-500'">{{ t('common.table') }}</span>
            <span class="text-xl font-black transition-colors" :class="selectedTable?.id === table.id ? 'text-white' : 'text-slate-800 dark:text-slate-100'">{{ table.table_number }}</span>
            <div class="mt-1 px-2 py-0.5 rounded-full text-[8px] font-black uppercase tracking-wider" 
              :class="[
                table.status === 'occupied' ? 'bg-orange-100 dark:bg-orange-950/30 text-orange-600 dark:text-orange-400' : 
                selectedTable?.id === table.id ? 'bg-indigo-500 text-white' : 'bg-green-50 dark:bg-green-950/30 text-green-600 dark:text-green-400'
              ]"
            >
              {{ table.status }}
            </div>
          </button>
        </div>
      </section>

      <!-- Product Selection -->
      <section v-if="selectedTable" class="animate-in slide-in-from-bottom-4 duration-500 pb-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
            <div class="p-1.5 bg-orange-100 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400 rounded-lg">
              <UtensilsCrossed class="w-4 h-4" />
            </div>
            {{ t('pos.availableMenu') }}
          </h3>
        </div>

        <div class="grid grid-cols-2 sm:grid-cols-3 xl:grid-cols-4 gap-4">
          <button 
            v-for="product in products" 
            :key="product.id"
            @click="addToCart(product)"
            class="bg-white dark:bg-slate-900 p-4 rounded-[1.5rem] border border-slate-100 dark:border-slate-800 shadow-sm hover:shadow-md transition-all duration-300 text-left group flex flex-col relative overflow-hidden"
            :disabled="product.stock <= 0"
            :class="product.stock <= 0 ? 'opacity-50 grayscale' : ''"
          >
            <div class="flex justify-between items-start mb-3">
              <div class="p-2 bg-slate-50 dark:bg-slate-800 rounded-xl group-hover:bg-indigo-600 group-hover:text-white transition-colors text-slate-400">
                <ChefHat class="w-4 h-4 group-hover:text-white" />
              </div>
              <span 
                :class="[
                  'text-[8px] px-1.5 py-0.5 rounded-md font-black uppercase tracking-wider border',
                  product.stock <= 0 ? 'bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400 border-red-100' : 'bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400 border-green-100'
                ]"
              >
                {{ product.stock > 0 ? product.stock + ' ' + t('pos.ready') : t('pos.out') }}
              </span>
            </div>
            
            <h4 class="font-bold text-slate-800 dark:text-slate-100 text-sm leading-tight mb-2 truncate">{{ product.name }}</h4>
            
            <div class="mt-auto flex items-center justify-between">
              <span class="text-sm font-black text-indigo-600 dark:text-indigo-400">{{ configStore.formatCurrency(product.price) }}</span>
              <div class="w-6 h-6 rounded-full bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 flex items-center justify-center group-hover:bg-indigo-600 group-hover:text-white transition-all scale-0 group-hover:scale-100">
                <Plus class="w-4 h-4 font-black" />
              </div>
            </div>
          </button>
        </div>
      </section>
      
      <div v-else class="h-48 flex flex-col items-center justify-center text-slate-300 dark:text-slate-800 border-2 border-dashed border-slate-200 dark:border-slate-900 rounded-[2rem]">
        <Users class="w-12 h-12 mb-3 opacity-20" />
        <p class="font-black text-sm opacity-40 uppercase tracking-widest">{{ t('pos.selectTable') }}</p>
      </div>
    </div>

    <!-- Right Panel: Order Setup -->
    <div class="w-full lg:w-[340px] bg-white dark:bg-slate-900 rounded-[2rem] shadow-xl border border-slate-100 dark:border-slate-800 flex flex-col overflow-hidden relative">
      <div class="p-6 bg-slate-900 dark:bg-slate-950 text-white">
        <div class="flex justify-between items-center mb-6">
          <h3 class="font-black text-lg flex items-center gap-2">
            <ShoppingCart class="w-5 h-5 text-indigo-400" />
            {{ t('pos.yourOrder') }}
          </h3>
          <button @click="clearCart" class="text-slate-400 hover:text-white transition-colors p-1.5">
            <Trash2 class="w-4 h-4" />
          </button>
        </div>

        <div class="space-y-4">
          <div class="space-y-1">
            <label class="text-[8px] font-black uppercase text-indigo-300 tracking-widest ml-1">{{ t('pos.customerBooker') }}</label>
            <div class="relative text-slate-900">
              <input 
                v-model="customerName" 
                type="text" 
                placeholder="e.g. John Doe"
                class="w-full bg-white/10 border border-white/10 rounded-xl px-10 py-3 text-sm font-bold text-white focus:bg-white/20 focus:border-indigo-400 outline-none transition-all"
              />
              <User class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-indigo-400" />
            </div>
          </div>

          <div v-if="selectedTable" class="bg-indigo-600/30 border border-indigo-500/30 rounded-xl p-3 flex items-center justify-between">
            <div>
              <p class="text-[8px] font-black uppercase text-indigo-300 tracking-widest">{{ t('common.table') }}</p>
              <p class="text-xl font-black">T-{{ selectedTable.table_number }}</p>
            </div>
            <Users class="w-6 h-6 text-indigo-400 opacity-50" />
          </div>
        </div>
      </div>

      <!-- Cart Items -->
      <div class="flex-1 overflow-auto p-6 space-y-4 custom-scrollbar dark:bg-slate-900">
        <div v-for="(item, index) in cart" :key="index" class="flex items-start gap-3">
          <div class="flex flex-col items-center gap-0.5 bg-slate-50 dark:bg-slate-800 rounded-lg p-0.5 border border-slate-100 dark:border-slate-700">
            <button @click="addToCart({id: item.product_id, name: item.name, price: item.price})" class="p-0.5 hover:text-indigo-600 dark:text-slate-400 transition-colors">
              <Plus class="w-2.5 h-2.5" />
            </button>
            <span class="text-xs font-black text-slate-800 dark:text-slate-100 w-5 text-center">{{ item.quantity }}</span>
            <button @click="removeFromCart(index)" class="p-0.5 hover:text-red-600 dark:text-slate-400 transition-colors">
              <Minus class="w-2.5 h-2.5" />
            </button>
          </div>
          <div class="flex-1 min-w-0">
            <p class="font-bold text-slate-800 dark:text-slate-100 truncate text-xs">{{ item.name }}</p>
            <p class="text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase">{{ configStore.formatCurrency(item.price) }}</p>
          </div>
          <p class="font-black text-slate-800 dark:text-slate-100 text-xs">{{ configStore.formatCurrency(item.price * item.quantity) }}</p>
        </div>

        <div v-if="cart.length === 0" class="h-full flex flex-col items-center justify-center text-slate-200 dark:text-slate-800 py-10">
          <ShoppingCart class="w-12 h-12 mb-3 opacity-10" />
          <p class="font-black uppercase tracking-widest text-[8px] opacity-30 text-slate-400 dark:text-slate-600">{{ t('pos.cartEmpty') }}</p>
        </div>
      </div>

      <div class="p-6 border-t border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900">
        <div class="flex justify-between items-center mb-4">
          <span class="text-slate-800 dark:text-slate-400 font-black text-sm uppercase tracking-widest">{{ t('common.total') }}</span>
          <span class="text-2xl font-black text-indigo-600 dark:text-indigo-400">{{ configStore.formatCurrency(totalPrice()) }}</span>
        </div>
        
        <button 
          @click="placeOrder"
          :disabled="!selectedTable || cart.length === 0 || actionLoading"
          class="w-full bg-slate-900 dark:bg-indigo-600 text-white py-4 rounded-xl font-black text-sm uppercase tracking-widest flex items-center justify-center gap-2 hover:bg-indigo-600 dark:hover:bg-indigo-700 disabled:opacity-20 disabled:cursor-not-allowed transition-all shadow-lg group"
        >
          <span v-if="actionLoading" class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
          <CheckCircle v-else class="w-5 h-5 text-green-400 group-hover:text-white transition-colors" />
          {{ t('common.add') }} {{ t('common.pos') }}
        </button>
      </div>
    </div>

    <!-- Universal Alert Modal -->
    <ConfirmModal 
      :show="alertConfig.show"
      :title="alertConfig.title"
      :message="alertConfig.message"
      :type="alertConfig.type"
      :confirmText="alertConfig.confirmText"
      :showCancel="false"
      @confirm="handleAlertConfirm"
    />
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-200);
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-800);
}
</style>
