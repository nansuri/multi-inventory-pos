<script setup lang="ts">
import BaseModal from './BaseModal.vue';
import type { LucideIcon } from 'lucide-vue-next';

interface Props {
  show: boolean;
  title: string;
  subtitle?: string;
  icon?: LucideIcon;
  iconClass?: string;
  maxWidth?: string;
  submitText?: string;
  cancelText?: string;
  loading?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  maxWidth: 'max-w-lg',
  submitText: 'Save Changes',
  cancelText: 'Cancel',
  loading: false
});

const emit = defineEmits(['submit', 'cancel']);
</script>

<template>
  <BaseModal :show="show" :maxWidth="maxWidth" @close="emit('cancel')">
    <div class="p-10">
      <div class="flex items-center gap-4 mb-8">
        <div v-if="icon" :class="['p-4 rounded-2xl flex items-center justify-center', iconClass || 'bg-indigo-50 text-indigo-600']">
          <component :is="icon" class="w-8 h-8" />
        </div>
        <div>
          <h3 class="text-2xl font-black text-slate-800 leading-tight">{{ title }}</h3>
          <p v-if="subtitle" class="text-sm font-medium text-slate-400 uppercase tracking-widest mt-1">{{ subtitle }}</p>
        </div>
      </div>

      <form @submit.prevent="emit('submit')" class="space-y-6">
        <slot />

        <div class="flex gap-4 pt-6 border-t border-slate-100">
          <button type="button" @click="emit('cancel')" class="flex-1 py-4 text-slate-400 font-bold hover:text-slate-600 transition-colors">
            {{ cancelText }}
          </button>
          <button 
            type="submit" 
            :disabled="loading"
            class="flex-[2] py-4 bg-indigo-600 text-white rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2"
          >
            <span v-if="loading" class="w-5 h-5 border-4 border-white border-t-transparent rounded-full animate-spin"></span>
            {{ submitText }}
          </button>
        </div>
      </form>
    </div>
  </BaseModal>
</template>
