<script setup lang="ts">
import BaseModal from './BaseModal.vue';
import { AlertCircle, CheckCircle2, Info } from 'lucide-vue-next';

interface Props {
  show: boolean;
  title: string;
  message: string;
  type?: 'info' | 'success' | 'warning' | 'danger';
  confirmText?: string;
  cancelText?: string;
  loading?: boolean;
  showCancel?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  type: 'info',
  confirmText: 'Confirm',
  cancelText: 'Cancel',
  loading: false,
  showCancel: true
});

defineEmits(['confirm', 'cancel']);

const getIcon = () => {
  switch (props.type) {
    case 'success': return CheckCircle2;
    case 'warning': return AlertCircle;
    case 'danger': return AlertCircle;
    default: return Info;
  }
};

const getTypeStyles = () => {
  switch (props.type) {
    case 'success': return 'bg-green-600 hover:bg-green-700 shadow-green-100 dark:shadow-green-900/20';
    case 'warning': return 'bg-orange-500 hover:bg-orange-600 shadow-orange-100 dark:shadow-orange-900/20';
    case 'danger': return 'bg-red-600 hover:bg-red-700 shadow-red-100 dark:shadow-red-900/20';
    default: return 'bg-indigo-600 hover:bg-indigo-700 shadow-indigo-100 dark:shadow-indigo-900/20';
  }
};
</script>

<template>
  <BaseModal :show="show" :title="title" :icon="getIcon()" @close="$emit('cancel')">
    <div class="space-y-10">
      <p class="text-lg font-medium text-slate-600 dark:text-slate-400 leading-relaxed">
        {{ message }}
      </p>

      <div class="flex flex-col sm:flex-row gap-3">
        <button 
          @click="$emit('confirm')"
          :disabled="loading"
          class="flex-1 py-5 rounded-2xl font-black text-white transition-all shadow-xl disabled:opacity-50 flex items-center justify-center gap-2"
          :class="getTypeStyles()"
        >
          <span v-if="loading" class="w-5 h-5 border-4 border-white/30 border-t-white rounded-full animate-spin"></span>
          {{ confirmText }}
        </button>
        <button 
          v-if="showCancel"
          @click="$emit('cancel')"
          :disabled="loading"
          class="flex-1 py-5 rounded-2xl font-black text-slate-400 dark:text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800 transition-all"
        >
          {{ cancelText }}
        </button>
      </div>
    </div>
  </BaseModal>
</template>
