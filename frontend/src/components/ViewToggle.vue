<script setup lang="ts">
import { LayoutGrid, List } from 'lucide-vue-next';

defineProps<{
  mode: 'list' | 'grid';
}>();

defineEmits(['update:mode']);
</script>

<template>
  <div class="flex items-center bg-white dark:bg-slate-900 p-1 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm h-[60px]">
    <button 
      @click="$emit('update:mode', 'list')"
      :class="[
        'flex items-center gap-2 px-5 py-2.5 rounded-xl transition-all font-black text-xs uppercase tracking-widest',
        mode === 'list' 
          ? 'bg-slate-900 dark:bg-indigo-600 text-white shadow-lg' 
          : 'text-slate-400 dark:text-slate-500 hover:text-slate-600 dark:hover:text-slate-300'
      ]"
    >
      <List class="w-4 h-4" />
      <span class="hidden sm:inline">List</span>
    </button>
    <button 
      @click="$emit('update:mode', 'grid')"
      :class="[
        'flex items-center gap-2 px-5 py-2.5 rounded-xl transition-all font-black text-xs uppercase tracking-widest',
        mode === 'grid' 
          ? 'bg-slate-900 dark:bg-indigo-600 text-white shadow-lg' 
          : 'text-slate-400 dark:text-slate-500 hover:text-slate-600 dark:hover:text-slate-300'
      ]"
    >
      <LayoutGrid class="w-4 h-4" />
      <span class="hidden sm:inline">Grid</span>
    </button>
  </div>
</template>
