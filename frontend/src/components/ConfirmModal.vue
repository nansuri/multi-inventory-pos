<script setup lang="ts">
import BaseModal from './BaseModal.vue';
import { AlertTriangle, Trash2, Info, HelpCircle, CheckCircle } from 'lucide-vue-next';

interface Props {
  show: boolean;
  title: string;
  message: string;
  type?: 'warning' | 'danger' | 'info' | 'question' | 'success';
  confirmText?: string;
  cancelText?: string;
  showCancel?: boolean;
  loading?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  type: 'info',
  confirmText: 'Confirm',
  cancelText: 'Cancel',
  showCancel: true,
  loading: false
});

const emit = defineEmits(['confirm', 'cancel']);

const getIcon = () => {
  switch (props.type) {
    case 'danger': return Trash2;
    case 'warning': return AlertTriangle;
    case 'question': return HelpCircle;
    case 'success': return CheckCircle;
    default: return Info;
  }
};

const getIconClass = () => {
  switch (props.type) {
    case 'danger': return 'bg-red-50 text-red-600';
    case 'warning': return 'bg-orange-50 text-orange-600';
    case 'question': return 'bg-indigo-50 text-indigo-600';
    case 'success': return 'bg-green-50 text-green-600';
    default: return 'bg-blue-50 text-blue-600';
  }
};

const getButtonClass = () => {
  switch (props.type) {
    case 'danger': return 'bg-red-600 hover:bg-red-700 shadow-red-100 shadow-lg';
    case 'warning': return 'bg-orange-600 hover:bg-orange-700 shadow-orange-100 shadow-lg';
    case 'success': return 'bg-green-600 hover:bg-green-700 shadow-green-100 shadow-lg';
    default: return 'bg-indigo-600 hover:bg-indigo-700 shadow-indigo-100 shadow-lg';
  }
};
</script>

<template>
  <BaseModal :show="show" maxWidth="max-w-sm" @close="emit('cancel')">
    <div class="p-10 text-center">
      <div :class="['w-20 h-20 rounded-[2rem] flex items-center justify-center mx-auto mb-6 transition-all duration-500', getIconClass()]">
        <component :is="getIcon()" class="w-10 h-10" />
      </div>
      
      <h3 class="text-2xl font-black text-slate-800 mb-2 leading-tight tracking-tight">{{ title }}</h3>
      <p class="text-sm font-medium text-slate-400 mb-10 leading-relaxed text-balance">{{ message }}</p>
      
      <div class="flex flex-col gap-2">
        <button 
          @click="emit('confirm')" 
          :disabled="loading"
          :class="['w-full py-5 text-white rounded-[1.5rem] font-black transition-all flex items-center justify-center gap-2', getButtonClass(), loading ? 'opacity-50 cursor-not-allowed' : '']"
        >
          <span v-if="loading" class="w-5 h-5 border-4 border-white border-t-transparent rounded-full animate-spin"></span>
          {{ confirmText }}
        </button>
        <button 
          v-if="showCancel"
          @click="emit('cancel')" 
          class="w-full py-4 text-slate-400 font-bold hover:text-slate-600 hover:bg-slate-50 rounded-2xl transition-all"
        >
          {{ cancelText }}
        </button>
      </div>
    </div>
  </BaseModal>
</template>
