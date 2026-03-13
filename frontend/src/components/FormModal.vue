<script setup lang="ts">
import BaseModal from './BaseModal.vue';

defineProps<{
  show: boolean;
  title: string;
  subtitle?: string;
  icon?: any;
  loading?: boolean;
  maxWidth?: string;
}>();

defineEmits(['submit', 'cancel']);
</script>

<template>
  <BaseModal :show="show" :title="title" :subtitle="subtitle" :icon="icon" :maxWidth="maxWidth" @close="$emit('cancel')">
    <form @submit.prevent="$emit('submit')" class="space-y-10">
      <slot />

      <div class="flex flex-col sm:flex-row gap-3 pt-4">
        <button 
          type="submit" 
          :disabled="loading"
          class="flex-1 py-5 bg-indigo-600 hover:bg-indigo-700 text-white rounded-2xl font-black transition-all shadow-xl shadow-indigo-100 dark:shadow-indigo-900/20 disabled:opacity-50 flex items-center justify-center gap-2"
        >
          <span v-if="loading" class="w-5 h-5 border-4 border-white/30 border-t-white rounded-full animate-spin"></span>
          {{ $t('common.save') }}
        </button>
        <button 
          type="button" 
          @click="$emit('cancel')"
          class="flex-1 py-5 text-slate-400 dark:text-slate-500 font-bold hover:bg-slate-50 dark:hover:bg-slate-800 rounded-2xl transition-all"
        >
          {{ $t('common.cancel') }}
        </button>
      </div>
    </form>
  </BaseModal>
</template>
