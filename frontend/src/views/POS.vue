<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { ShoppingCart, CheckCircle, Users, Plus, Minus, UtensilsCrossed, Trash2, CreditCard, ChefHat } from 'lucide-vue-next';

const tables = ref<any[]>([]);
const products = ref<any[]>([]);
const selectedTable = ref<any>(null);
const cart = ref<any[]>([]);
const loading = ref(true);
const configStore = useConfigStore();
const isAddTableModalOpen = ref(false);

const newTable = ref({
  table_number: '',
  capacity: 2
});

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

const addTable = async () => {
  try {
    await api.post('/api/tables', newTable.value);
    isAddTableModalOpen.value = false;
    newTable.value = { table_number: '', capacity: 2 };
    await fetchInitialData();
  } catch (error) {
    console.error('Failed to add table', error);
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
  
  try {
    const orderData = {
      table_id: selectedTable.value.id,
      total_price: totalPrice(),
      items: cart.value.map(item => ({
        product_id: item.product_id,
        quantity: item.quantity,
        price: item.price
      }))
    };
    
    const response = await api.post('/api/orders', orderData);
    const orderId = response.data.id;
    
    // Automatically complete for this demo version to show inventory deduction
    await api.post(`/api/orders/${orderId}/complete`);
    
    alert('Order placed and payment processed!');
    cart.value = [];
    selectedTable.value = null;
    await fetchInitialData();
  } catch (error: any) {
    alert('Failed to place order: ' + (error.response?.data?.error || error.message));
  }
};

onMounted(fetchInitialData);
</script>

<template>
  <DashboardLayout>
    <div class="flex flex-col lg:flex-row gap-8 h-[calc(100vh-160px)] -mt-2">
      <!-- Main Selection Area -->
      <div class="flex-1 space-y-10 overflow-auto pr-2 custom-scrollbar">
        <!-- Tables Selection -->
        <section>
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-black text-slate-800 tracking-tight flex items-center gap-3">
              <div class="p-2 bg-indigo-100 text-indigo-600 rounded-lg">
                <Users class="w-5 h-5" />
              </div>
              {{ $t('pos.floorPlan') }}
            </h3>
            <span class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ tables.length }} {{ $t('pos.tablesAvailable', 'Tables Available') }}</span>
          </div>
          
          <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 xl:grid-cols-5 gap-4">
            <button 
              v-for="table in tables" 
              :key="table.id"
              @click="selectedTable = table"
              :class="[
                'p-6 rounded-[2rem] border-2 transition-all duration-300 relative group flex flex-col items-center justify-center gap-1',
                table.status === 'occupied' 
                  ? 'bg-slate-50 border-slate-100 cursor-not-allowed opacity-60' 
                  : selectedTable?.id === table.id 
                    ? 'bg-indigo-600 border-indigo-600 text-white shadow-xl shadow-indigo-200' 
                    : 'bg-white border-slate-100 hover:border-indigo-300 hover:shadow-lg'
              ]"
              :disabled="table.status === 'occupied'"
            >
              <span class="text-sm font-black uppercase tracking-tighter" :class="selectedTable?.id === table.id ? 'text-indigo-100' : 'text-slate-400'">{{ $t('common.table', 'Table') }}</span>
              <span class="text-2xl font-black">{{ table.table_number }}</span>
              <div class="mt-2 px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-wider" 
                :class="[
                  table.status === 'occupied' ? 'bg-orange-100 text-orange-600' : 
                  selectedTable?.id === table.id ? 'bg-indigo-500 text-white' : 'bg-green-50 text-green-600'
                ]"
              >
                {{ table.status }}
              </div>
            </button>
            
            <button 
              @click="isAddTableModalOpen = true"
              class="p-6 rounded-[2rem] border-4 border-dashed border-slate-100 flex flex-col items-center justify-center text-slate-300 hover:text-indigo-500 hover:border-indigo-200 transition-all group"
            >
              <Plus class="w-8 h-8 group-hover:scale-110 transition-transform" />
              <span class="text-[10px] font-black uppercase tracking-widest mt-2">{{ $t('pos.newTable') }}</span>
            </button>
          </div>
        </section>

        <!-- Product Selection -->
        <section v-if="selectedTable" class="animate-in slide-in-from-bottom-4 duration-500">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-black text-slate-800 tracking-tight flex items-center gap-3">
              <div class="p-2 bg-orange-100 text-orange-600 rounded-lg">
                <UtensilsCrossed class="w-5 h-5" />
              </div>
              {{ $t('pos.availableMenu') }}
            </h3>
            <div class="bg-indigo-600 text-white text-[10px] font-black px-3 py-1 rounded-full uppercase">{{ $t('pos.orderingFor', 'Ordering for') }} Table {{ selectedTable.table_number }}</div>
          </div>

          <div class="grid grid-cols-2 sm:grid-cols-3 xl:grid-cols-4 gap-6">
            <button 
              v-for="product in products" 
              :key="product.id"
              @click="addToCart(product)"
              class="bg-white p-6 rounded-[2.5rem] border border-slate-100 shadow-sm hover:shadow-xl transition-all duration-300 text-left group flex flex-col relative overflow-hidden"
              :disabled="product.stock <= 0"
              :class="product.stock <= 0 ? 'opacity-50 grayscale' : ''"
            >
              <div class="flex justify-between items-start mb-4">
                <div class="p-3 bg-slate-50 rounded-2xl group-hover:bg-indigo-600 group-hover:text-white transition-colors">
                  <ChefHat class="w-6 h-6" />
                </div>
                <span 
                  :class="[
                    'text-[10px] px-2 py-1 rounded-lg font-black uppercase tracking-wider border',
                    product.stock <= 0 ? 'bg-red-50 text-red-600 border-red-100' : 'bg-green-50 text-green-600 border-green-100'
                  ]"
                >
                  {{ product.stock > 0 ? product.stock + ' ' + $t('pos.ready') : $t('pos.out') }}
                </span>
              </div>
              
              <h4 class="font-black text-slate-800 text-lg leading-tight mb-4">{{ product.name }}</h4>
              
              <div class="mt-auto flex items-center justify-between">
                <span class="text-xl font-black text-indigo-600">{{ configStore.formatCurrency(product.price) }}</span>
                <div class="w-8 h-8 rounded-full bg-indigo-50 text-indigo-600 flex items-center justify-center group-hover:bg-indigo-600 group-hover:text-white transition-all scale-0 group-hover:scale-100">
                  <Plus class="w-5 h-5 font-black" />
                </div>
              </div>
            </button>
          </div>
        </section>
        
        <div v-else class="h-64 flex flex-col items-center justify-center text-slate-300 border-4 border-dashed border-slate-100 rounded-[3rem]">
          <Users class="w-16 h-16 mb-4 opacity-20" />
          <p class="font-black text-lg opacity-40 uppercase tracking-widest">{{ $t('pos.selectTable') }}</p>
        </div>
      </div>

      <!-- Right Panel: Current Order (Receipt Style) -->
      <div class="w-full lg:w-[400px] bg-white rounded-[3rem] shadow-2xl shadow-slate-200 border border-slate-100 flex flex-col overflow-hidden relative">
        <!-- Header -->
        <div class="p-8 bg-slate-900 text-white relative">
          <div class="flex justify-between items-center mb-6">
            <h3 class="font-black text-xl flex items-center gap-3">
              <ShoppingCart class="w-6 h-6 text-indigo-400" />
              {{ $t('pos.yourOrder') }}
            </h3>
            <button @click="clearCart" class="text-slate-400 hover:text-white transition-colors p-2">
              <Trash2 class="w-5 h-5" />
            </button>
          </div>
          
          <div v-if="selectedTable" class="bg-indigo-600/30 border border-indigo-500/30 rounded-2xl p-4 flex items-center justify-between">
            <div>
              <p class="text-[10px] font-black uppercase text-indigo-300 tracking-widest">{{ $t('pos.selectedTable', 'Selected Table') }}</p>
              <p class="text-2xl font-black">T-{{ selectedTable.table_number }}</p>
            </div>
            <Users class="w-8 h-8 text-indigo-400 opacity-50" />
          </div>
        </div>

        <!-- Receipt Items -->
        <div class="flex-1 overflow-auto p-8 space-y-6 custom-scrollbar">
          <div v-for="(item, index) in cart" :key="index" class="flex items-start gap-4 animate-in slide-in-from-right-4 duration-300">
            <div class="flex flex-col items-center gap-1 bg-slate-50 rounded-xl p-1 border border-slate-100">
              <button @click="addToCart({id: item.product_id, name: item.name, price: item.price})" class="p-1 hover:text-indigo-600 transition-colors">
                <Plus class="w-3 h-3" />
              </button>
              <span class="text-sm font-black text-slate-800 w-6 text-center">{{ item.quantity }}</span>
              <button @click="removeFromCart(index)" class="p-1 hover:text-red-600 transition-colors">
                <Minus class="w-3 h-3" />
              </button>
            </div>
            
            <div class="flex-1 min-w-0">
              <p class="font-bold text-slate-800 truncate">{{ item.name }}</p>
              <p class="text-[10px] font-black text-slate-400 uppercase">{{ configStore.formatCurrency(item.price) }} unit</p>
            </div>
            
            <p class="font-black text-slate-800">{{ configStore.formatCurrency(item.price * item.quantity) }}</p>
          </div>

          <div v-if="cart.length === 0" class="h-full flex flex-col items-center justify-center text-slate-200 py-20">
            <ShoppingCart class="w-20 h-20 mb-4 opacity-10" />
            <p class="font-black uppercase tracking-widest text-sm opacity-30 text-slate-400">{{ $t('pos.cartEmpty') }}</p>
          </div>
        </div>

        <!-- Footer / Checkout -->
        <div class="p-8 border-t-2 border-dashed border-slate-100 bg-slate-50/50">
          <div class="space-y-3 mb-8">
            <div class="flex justify-between items-center text-slate-400 font-bold text-xs uppercase tracking-widest">
              <span>{{ $t('pos.subtotal') }}</span>
              <span>{{ configStore.formatCurrency(totalPrice()) }}</span>
            </div>
            <div class="flex justify-between items-center text-slate-400 font-bold text-xs uppercase tracking-widest">
              <span>{{ $t('pos.tax') }} (0%)</span>
              <span>{{ configStore.formatCurrency(0) }}</span>
            </div>
            <div class="pt-3 border-t border-slate-200 flex justify-between items-center">
              <span class="text-slate-800 font-black text-lg">{{ $t('pos.totalAmount') }}</span>
              <span class="text-3xl font-black text-indigo-600">{{ configStore.formatCurrency(totalPrice()) }}</span>
            </div>
          </div>
          
          <button 
            @click="placeOrder"
            :disabled="!selectedTable || cart.length === 0"
            class="w-full bg-slate-900 text-white py-5 rounded-[1.5rem] font-black flex items-center justify-center gap-3 hover:bg-indigo-600 disabled:opacity-20 disabled:cursor-not-allowed transition-all shadow-xl shadow-slate-200 group"
          >
            <CreditCard class="w-6 h-6" />
            {{ $t('pos.processPayment') }}
            <CheckCircle class="w-5 h-5 text-green-400 group-hover:text-white transition-colors" />
          </button>
        </div>
      </div>
    </div>

    <!-- Add Table Modal -->
    <div v-if="isAddTableModalOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center p-4 z-50">
      <div class="bg-white rounded-[2.5rem] w-full max-w-sm p-10 shadow-2xl">
        <h3 class="text-2xl font-black text-slate-800 mb-2">{{ $t('pos.newTable') }}</h3>
        <p class="text-sm font-medium text-slate-400 mb-8">Define a new physical spot in your floor plan.</p>
        
        <form @submit.prevent="addTable" class="space-y-6">
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('pos.tableIdentifier', 'Table Identifier') }}</label>
            <input v-model="newTable.table_number" type="text" required placeholder="e.g. 05, A2, Patio" class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none" />
          </div>
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('pos.seatingCapacity', 'Seating Capacity') }}</label>
            <input v-model.number="newTable.capacity" type="number" min="1" required class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none" />
          </div>
          <div class="flex flex-col gap-2 pt-4">
            <button type="submit" class="w-full py-5 bg-indigo-600 text-white rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all">{{ $t('common.save') }}</button>
            <button type="button" @click="isAddTableModalOpen = false" class="w-full py-4 text-slate-400 font-bold hover:text-slate-600 transition-colors">{{ $t('common.cancel') }}</button>
          </div>
        </form>
      </div>
    </div>
  </DashboardLayout>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-200);
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: var(--color-indigo-200);
}

.scale-in-center {
  animation: scale-in-center 0.3s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}

@keyframes scale-in-center {
  0% { transform: scale(0.9); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}
</style>
