<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';

interface Props {
  name: string;
  path: string;
  icon: any;
  isCollapsed: boolean;
  isSubItem?: boolean;
}

const props = defineProps<Props>();
const route = useRoute();

const isActive = computed(() => route.path === props.path);
</script>

<template>
  <router-link 
    :to="path"
    class="flex items-center rounded-xl transition-all duration-200 group relative mb-0.5 outline-none"
    :class="[
      isCollapsed ? 'justify-center p-2.5' : 'px-3 py-2',
      isActive 
        ? 'bg-indigo-600 text-white shadow-md shadow-indigo-100 dark:shadow-indigo-900/40' 
        : 'text-slate-500 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800 hover:text-slate-900 dark:hover:text-slate-100'
    ]"
  >
    <component 
      :is="icon" 
      :class="[
        'w-4.5 h-4.5 flex-shrink-0 transition-colors',
        isActive ? 'text-white' : 'group-hover:text-slate-900 dark:group-hover:text-slate-100',
        isSubItem && !isActive ? 'text-slate-400 dark:text-slate-500' : ''
      ]" 
    />
    
    <span 
      v-if="!isCollapsed" 
      class="ml-2.5 font-bold text-xs whitespace-nowrap overflow-hidden transition-all"
    >
      {{ name }}
    </span>

    <!-- Tooltip for collapsed mode -->
    <div 
      v-if="isCollapsed"
      class="absolute left-12 bg-slate-800 dark:bg-slate-700 text-white text-[9px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 pointer-events-none transition-opacity whitespace-nowrap z-50 shadow-xl font-black uppercase tracking-widest"
    >
      {{ name }}
    </div>
  </router-link>
</template>
