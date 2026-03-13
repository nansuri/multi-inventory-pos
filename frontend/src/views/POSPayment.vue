<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useI18n } from 'vue-i18n';
import { CreditCard, CheckCircle, Users, ShoppingCart, Clock, User, ArrowRight } from 'lucide-vue-next';

const { t } = useI18n();
const orders = ref<any[]>([]);
const selectedOrder = ref<any>(null);
const loading = ref(true);
const actionLoading = ref(false);
const configStore = useConfigStore();

const alertConfig = ref({
  show: false,
  title: '',
  message: '',
  type: 'info' as 'info' | 'success' | 'warning' | 'danger',
  confirmText: 'OK'
});

const showAlert = (title: string, message: string, type: any = 'info', confirmText = 'OK') => {
  alertConfig.value = { show: true, title, message, type, confirmText };
};

const fetchOrders = async () => {
  try {
    const response = await api.get('/api/orders');
    orders.value = response.data.filter((o: any) => o.status === 'pending');
  } catch (error) {
    console.error('Failed to fetch orders', error);
  } finally {
    loading.value = false;
  }
};

const processPayment = async () => {
  if (!selectedOrder.value) return;
  
  actionLoading.value = true;
  try {
    await api.post(`/api/orders/${selectedOrder.value.id}/complete`);
    showAlert(t('common.success'), t('pos.successPayment'), "success");
    selectedOrder.value = null;
    await fetchOrders();
  } catch (error: any) {
    showAlert(t('common.error'), error.response?.data?.error || error.message, "danger");
  } finally {
    actionLoading.value = false;
  }
};

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr);
  return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
};

onMounted(fetchOrders);
</script>

<template>
  <DashboardLayout>
    <div class="flex flex-col lg:flex-row gap-8 h-[calc(100vh-160px)] -mt-2">
      <!-- Orders List -->
      <div class="flex-1 space-y-6 overflow-auto pr-2 custom-scrollbar">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <div class="p-2 bg-indigo-100 text-indigo-600 rounded-lg">
              <CreditCard class="w-5 h-5" />
            </div>
            {{ t('pos.activeBills') }}
          </h3>
          <span class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ orders.length }} {{ t('pos.pendingPayments') }}</span>
        </div>

        <div v-if="orders.length > 0" class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <button 
            v-for="order in orders" 
            :key="order.id"
            @click="selectedOrder = order"
            :class="[
              'p-6 rounded-[2.5rem] border-2 text-left transition-all duration-300 group relative overflow-hidden',
              selectedOrder?.id === order.id 
                ? 'bg-indigo-600 border-indigo-600 text-white shadow-xl shadow-indigo-200' 
                : 'bg-white border-slate-100 hover:border-indigo-300'
            ]"
          >
            <div class="flex justify-between items-start mb-4">
              <div :class="[
                'p-3 rounded-2xl transition-colors',
                selectedOrder?.id === order.id ? 'bg-indigo-500' : 'bg-slate-50'
              ]">
                <Users class="w-6 h-6" :class="selectedOrder?.id === order.id ? 'text-white' : 'text-indigo-600'" />
              </div>
              <div class="text-right">
                <p class="text-[10px] font-black uppercase opacity-60">{{ t('common.table') }}</p>
                <p class="text-xl font-black">{{ order.table?.table_number || 'N/A' }}</p>
              </div>
            </div>

            <div class="space-y-1 mb-6">
              <div class="flex items-center gap-2">
                <User class="w-3 h-3 opacity-60" />
                <span class="font-bold text-sm">{{ order.customer_name }}</span>
              </div>
              <div class="flex items-center gap-2 text-[10px] font-medium opacity-60">
                <Clock class="w-3 h-3" />
                {{ formatDate(order.created_at) }}
              </div>
            </div>

            <div class="flex items-center justify-between mt-auto">
              <span class="text-lg font-black">{{ configStore.formatCurrency(order.total_price) }}</span>
              <div :class="[
                'w-8 h-8 rounded-full flex items-center justify-center transition-all',
                selectedOrder?.id === order.id ? 'bg-white text-indigo-600' : 'bg-indigo-50 text-indigo-600 opacity-0 group-hover:opacity-100'
              ]">
                <ArrowRight class="w-4 h-4" />
              </div>
            </div>
          </button>
        </div>

        <div v-else class="h-64 flex flex-col items-center justify-center text-slate-300 border-4 border-dashed border-slate-100 rounded-[3rem]">
          <CreditCard class="w-16 h-16 mb-4 opacity-20" />
          <p class="font-black text-lg opacity-40 uppercase tracking-widest text-center px-10">{{ t('pos.noBills') }}<br><span class="text-xs font-medium">{{ t('pos.allClear') }}</span></p>
        </div>
      </div>

      <!-- Payment Summary Panel -->
      <div class="w-full lg:w-[400px] bg-white rounded-[3rem] shadow-2xl shadow-slate-200 border border-slate-100 flex flex-col overflow-hidden">
        <div v-if="selectedOrder" class="flex flex-col h-full animate-in slide-in-from-right-4 duration-300">
          <div class="p-8 bg-slate-900 text-white">
            <h3 class="font-black text-xl flex items-center gap-3 mb-8">
              <CreditCard class="w-6 h-6 text-indigo-400" />
              {{ t('pos.billDetails') }}
            </h3>
            <div class="grid grid-cols-2 gap-4">
              <div class="bg-indigo-600/30 border border-indigo-500/30 rounded-2xl p-4">
                <p class="text-[10px] font-black uppercase text-indigo-300 tracking-widest mb-1">{{ t('common.customer') }}</p>
                <p class="text-lg font-black truncate">{{ selectedOrder.customer_name }}</p>
              </div>
              <div class="bg-slate-800 border border-slate-700 rounded-2xl p-4">
                <p class="text-[10px] font-black uppercase text-slate-500 tracking-widest mb-1">{{ t('common.table') }}</p>
                <p class="text-lg font-black">{{ selectedOrder.table?.table_number }}</p>
              </div>
            </div>
          </div>

          <div class="flex-1 overflow-auto p-8 space-y-6 custom-scrollbar">
            <div v-for="item in selectedOrder.items" :key="item.id" class="flex justify-between items-center">
              <div>
                <p class="font-bold text-slate-800 text-sm">{{ item.product?.name }}</p>
                <p class="text-[10px] font-black text-slate-400 uppercase">{{ configStore.formatCurrency(item.price) }} x {{ item.quantity }}</p>
              </div>
              <p class="font-black text-slate-800 text-sm">{{ configStore.formatCurrency(item.price * item.quantity) }}</p>
            </div>
          </div>

          <div class="p-8 border-t-2 border-dashed border-slate-100 bg-slate-50/50">
            <div class="flex justify-between items-center mb-8">
              <span class="text-slate-800 font-black text-lg">{{ t('pos.totalAmount') }}</span>
              <span class="text-3xl font-black text-indigo-600">{{ configStore.formatCurrency(selectedOrder.total_price) }}</span>
            </div>
            <button @click="processPayment" :disabled="actionLoading" class="w-full bg-slate-900 text-white py-5 rounded-[1.5rem] font-black flex items-center justify-center gap-3 hover:bg-green-600 transition-all shadow-xl shadow-slate-200 group">
              <span v-if="actionLoading" class="w-5 h-5 border-4 border-white border-t-transparent rounded-full animate-spin"></span>
              <template v-else>
                <CheckCircle class="w-6 h-6 text-green-400 group-hover:text-white transition-colors" />
                {{ t('pos.processPayment') }}
              </template>
            </button>
          </div>
        </div>

        <div v-else class="h-full flex flex-col items-center justify-center text-slate-200 py-20 p-10 text-center">
          <div class="w-20 h-20 bg-slate-50 rounded-[2rem] flex items-center justify-center mb-6">
            <ShoppingCart class="w-10 h-10 text-slate-200" />
          </div>
          <p class="font-black uppercase tracking-widest text-[10px] text-slate-400">{{ t('pos.selectBillDesc') }}</p>
        </div>
      </div>
    </div>

    <ConfirmModal 
      :show="alertConfig.show"
      :title="alertConfig.title"
      :message="alertConfig.message"
      :type="alertConfig.type"
      :confirmText="alertConfig.confirmText"
      :showCancel="false"
      @confirm="alertConfig.show = false"
    />
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
</style>
