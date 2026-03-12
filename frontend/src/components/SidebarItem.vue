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
    class="flex items-center rounded-2xl transition-all duration-200 group relative mb-1 outline-none"
    :class="[
      isCollapsed ? 'justify-center p-3' : 'px-4 py-3',
      isActive 
        ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-100' 
        : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'
    ]"
  >
    <component 
      :is="icon" 
      :class="[
        'w-5 h-5 flex-shrink-0 transition-colors',
        isActive ? 'text-white' : 'group-hover:text-slate-900',
        isSubItem && !isActive ? 'text-slate-400' : ''
      ]" 
    />
    
    <span 
      v-if="!isCollapsed" 
      class="ml-3 font-bold text-sm whitespace-nowrap overflow-hidden transition-all"
    >
      {{ name }}
    </span>

    <!-- Tooltip for collapsed mode -->
    <div 
      v-if="isCollapsed"
      class="absolute left-14 bg-slate-800 text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 pointer-events-none transition-opacity whitespace-nowrap z-50 shadow-xl font-black uppercase tracking-widest"
    >
      {{ name }}
    </div>
  </router-link>
</template>
