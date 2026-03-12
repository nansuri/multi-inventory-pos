<script setup lang="ts">
import { LayoutList, LayoutGrid } from 'lucide-vue-next';

defineProps<{
  mode: 'list' | 'grid'
}>();

const emit = defineEmits(['update:mode']);
</script>

<template>
  <div class="flex items-center bg-white border border-slate-200 p-1 rounded-xl shadow-sm">
    <button 
      @click="emit('update:mode', 'list')"
      class="p-2 rounded-lg transition-all flex items-center gap-2"
      :class="mode === 'list' ? 'bg-indigo-600 text-white shadow-md' : 'text-slate-400 hover:text-slate-600'"
    >
      <LayoutList class="w-4 h-4" />
      <span v-if="mode === 'list'" class="text-[10px] font-black uppercase tracking-widest pr-1">List</span>
    </button>
    <button 
      @click="emit('update:mode', 'grid')"
      class="p-2 rounded-lg transition-all flex items-center gap-2"
      :class="mode === 'grid' ? 'bg-indigo-600 text-white shadow-md' : 'text-slate-400 hover:text-slate-600'"
    >
      <LayoutGrid class="w-4 h-4" />
      <span v-if="mode === 'grid'" class="text-[10px] font-black uppercase tracking-widest pr-1">Grid</span>
    </button>
  </div>
</template>
