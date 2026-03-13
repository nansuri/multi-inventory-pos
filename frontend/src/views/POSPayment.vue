<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useAuthStore } from '../stores/auth';
import { useI18n } from 'vue-i18n';
import { CreditCard, CheckCircle, Users, ShoppingCart, Clock, User, ArrowRight } from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
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

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, fetchOrders);
</script>

<template>
  <div class="flex flex-col lg:flex-row gap-6 h-[calc(100vh-140px)] -mt-1">
    <!-- Orders List -->
    <div class="flex-1 space-y-4 overflow-auto pr-2 custom-scrollbar">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <div class="p-1.5 bg-indigo-100 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 rounded-lg">
            <CreditCard class="w-4 h-4" />
          </div>
          {{ t('pos.activeBills') }}
        </h3>
        <span class="text-[10px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ orders.length }} {{ t('pos.pendingPayments') }}</span>
      </div>

      <div v-if="orders.length > 0" class="grid grid-cols-1 md:grid-cols-2 gap-3">
        <button 
          v-for="order in orders" 
          :key="order.id"
          @click="selectedOrder = order"
          :class="[
            'p-4 rounded-[1.5rem] border-2 text-left transition-all duration-300 group relative overflow-hidden',
            selectedOrder?.id === order.id 
              ? 'bg-indigo-600 border-indigo-600 text-white shadow-lg shadow-indigo-200 dark:shadow-none' 
              : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 hover:border-indigo-300 dark:hover:border-indigo-500'
          ]"
        >
          <div class="flex justify-between items-start mb-3">
            <div :class="[
              'p-2 rounded-xl transition-colors',
              selectedOrder?.id === order.id ? 'bg-indigo-500' : 'bg-slate-50 dark:bg-slate-800'
            ]">
              <Users class="w-5 h-5" :class="selectedOrder?.id === order.id ? 'text-white' : 'text-indigo-600 dark:text-indigo-400'" />
            </div>
            <div class="text-right">
              <p class="text-[8px] font-black uppercase opacity-60">{{ t('common.table') }}</p>
              <p class="text-lg font-black leading-none">T-{{ order.table?.table_number || 'N/A' }}</p>
            </div>
          </div>

          <div class="space-y-0.5 mb-4">
            <div class="flex items-center gap-1.5">
              <User class="w-3 h-3 opacity-60" />
              <span class="font-bold text-xs truncate">{{ order.customer_name }}</span>
            </div>
            <div class="flex items-center gap-1.5 text-[9px] font-medium opacity-60">
              <Clock class="w-2.5 h-2.5" />
              {{ formatDate(order.created_at) }}
            </div>
          </div>

          <div class="flex items-center justify-between mt-auto">
            <span class="text-base font-black">{{ configStore.formatCurrency(order.total_price) }}</span>
            <div :class="[
              'w-6 h-6 rounded-full flex items-center justify-center transition-all',
              selectedOrder?.id === order.id ? 'bg-white text-indigo-600' : 'bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 opacity-0 group-hover:opacity-100'
            ]">
              <ArrowRight class="w-3.5 h-3.5" />
            </div>
          </div>
        </button>
      </div>

      <div v-else class="h-48 flex flex-col items-center justify-center text-slate-300 dark:text-slate-800 border-2 border-dashed border-slate-200 dark:border-slate-900 rounded-[2rem]">
        <CreditCard class="w-12 h-12 mb-3 opacity-20" />
        <p class="font-black text-sm opacity-40 uppercase tracking-widest text-center px-6">{{ t('pos.noBills') }}<br><span class="text-[10px] font-medium">{{ t('pos.allClear') }}</span></p>
      </div>
    </div>

    <!-- Payment Summary Panel -->
    <div class="w-full lg:w-[340px] bg-white dark:bg-slate-900 rounded-[2rem] shadow-xl border border-slate-100 dark:border-slate-800 flex flex-col overflow-hidden">
      <div v-if="selectedOrder" class="flex flex-col h-full animate-in slide-in-from-right-4 duration-300">
        <div class="p-6 bg-slate-900 dark:bg-slate-950 text-white">
          <h3 class="font-black text-lg flex items-center gap-2 mb-6">
            <CreditCard class="w-5 h-5 text-indigo-400" />
            {{ t('pos.billDetails') }}
          </h3>
          <div class="grid grid-cols-2 gap-3 text-slate-900">
            <div class="bg-indigo-600/30 border border-indigo-500/30 rounded-xl p-3">
              <p class="text-[8px] font-black uppercase text-indigo-300 tracking-widest mb-0.5">{{ t('common.customer') }}</p>
              <p class="text-sm font-black truncate text-white">{{ selectedOrder.customer_name }}</p>
            </div>
            <div class="bg-slate-800 dark:bg-slate-900 border border-slate-700 dark:border-slate-800 rounded-xl p-3">
              <p class="text-[8px] font-black uppercase text-slate-500 dark:text-slate-600 tracking-widest mb-0.5">{{ t('common.table') }}</p>
              <p class="text-sm font-black text-white">T-{{ selectedOrder.table?.table_number }}</p>
            </div>
          </div>
        </div>

        <div class="flex-1 overflow-auto p-6 space-y-4 custom-scrollbar dark:bg-slate-900">
          <div v-for="item in selectedOrder.items" :key="item.id" class="flex justify-between items-center">
            <div>
              <p class="font-bold text-slate-800 dark:text-slate-100 text-xs">{{ item.product?.name }}</p>
              <p class="text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase">{{ configStore.formatCurrency(item.price) }} x {{ item.quantity }}</p>
            </div>
            <p class="font-black text-slate-800 dark:text-slate-100 text-xs">{{ configStore.formatCurrency(item.price * item.quantity) }}</p>
          </div>
        </div>

        <div class="p-6 border-t border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900">
          <div class="flex justify-between items-center mb-6">
            <span class="text-slate-800 dark:text-slate-400 font-black text-xs uppercase tracking-widest">{{ t('pos.totalAmount') }}</span>
            <span class="text-2xl font-black text-indigo-600 dark:text-indigo-400">{{ configStore.formatCurrency(selectedOrder.total_price) }}</span>
          </div>
          <button @click="processPayment" :disabled="actionLoading" class="w-full bg-slate-900 dark:bg-indigo-600 text-white py-4 rounded-xl font-black text-sm uppercase tracking-widest flex items-center justify-center gap-2 hover:bg-green-600 transition-all shadow-lg group">
            <span v-if="actionLoading" class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
            <template v-else>
              <CheckCircle class="w-5 h-5 text-green-400 group-hover:text-white transition-colors" />
              {{ t('pos.processPayment') }}
            </template>
          </button>
        </div>
      </div>

      <div v-else class="h-full flex flex-col items-center justify-center text-slate-200 dark:text-slate-800 py-16 p-8 text-center">
        <div class="w-16 h-16 bg-slate-50 dark:bg-slate-800 rounded-2xl flex items-center justify-center mb-4 border border-slate-100 dark:border-slate-700">
          <ShoppingCart class="w-8 h-8 text-slate-200 dark:text-slate-700" />
        </div>
        <p class="font-black uppercase tracking-widest text-[9px] text-slate-400 dark:text-slate-600 px-4">{{ t('pos.selectBillDesc') }}</p>
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
