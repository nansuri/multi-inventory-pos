<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed, watch } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useRoute } from 'vue-router';
import { MapPin, ChevronRight, Check } from 'lucide-vue-next';

const props = defineProps<{
  isCollapsed?: boolean
}>();

const authStore = useAuthStore();
const route = useRoute();
const isOpen = ref(false);
const dropdownRef = ref<HTMLElement | null>(null);

// Only allow "All Branches" on the Order History page
const isHistoryPage = computed(() => route.path === '/reports/orders');

// Safety check: if we leave history page while All Branches is selected, switch to first branch
watch(() => route.path, (newPath) => {
  if (newPath !== '/reports/orders' && authStore.selectedBranchID === null && authStore.branches.length > 0) {
    const firstBranch = authStore.branches[0];
    if (firstBranch) {
      authStore.setBranch(firstBranch.id);
    }
  }
});

const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

const selectBranch = (branchID: number | null) => {
  if (branchID === null) {
    authStore.selectedBranchID = null;
    localStorage.removeItem('selectedBranchID');
    window.location.reload();
  } else {
    authStore.setBranch(branchID);
  }
  isOpen.value = false;
};

const handleClickOutside = (event: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
    isOpen.value = false;
  }
};

onMounted(() => {
  window.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  window.removeEventListener('click', handleClickOutside);
});
</script>

<template>
  <div v-if="authStore.branches.length > 0" class="relative px-2 mb-4" ref="dropdownRef">
    <!-- Sidebar Trigger -->
    <button 
      @click.stop="toggleDropdown"
      :class="[
        'w-full flex items-center transition-all duration-300 rounded-xl group relative overflow-hidden',
        isCollapsed ? 'justify-center p-2.5' : 'px-3 py-2.5 bg-slate-50 dark:bg-slate-800/50 hover:bg-indigo-50 dark:hover:bg-indigo-900/20 border border-slate-100 dark:border-slate-800'
      ]"
    >
      <div :class="[
        'rounded-lg text-white shadow-lg flex items-center justify-center shrink-0 transition-all duration-300',
        isCollapsed ? 'w-8 h-8 bg-indigo-600' : 'w-7 h-7 bg-indigo-600'
      ]">
        <MapPin :class="isCollapsed ? 'w-4 h-4' : 'w-3.5 h-3.5'" />
      </div>
      
      <div v-if="!isCollapsed" class="ml-3 flex flex-col items-start overflow-hidden">
        <span class="text-[8px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest leading-none mb-1">
          {{ authStore.selectedBranchID === null ? 'Organization' : 'Current Branch' }}
        </span>
        <span class="text-xs font-black text-slate-700 dark:text-slate-200 tracking-tight leading-none truncate w-full pr-4 text-left">
          {{ authStore.selectedBranchID === null ? 'All Branches' : (authStore.currentBranch?.name || 'Select Branch') }}
        </span>
      </div>
      
      <ChevronRight 
        v-if="!isCollapsed" 
        :class="['absolute right-2 w-3 h-3 text-slate-300 transition-transform duration-300', isOpen ? 'rotate-90' : '']" 
      />

      <!-- Tooltip for collapsed mode -->
      <div 
        v-if="isCollapsed"
        class="absolute left-14 bg-indigo-600 text-white text-[9px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 pointer-events-none transition-opacity whitespace-nowrap z-50 shadow-xl font-black uppercase tracking-widest"
      >
        {{ authStore.selectedBranchID === null ? 'All Branches' : authStore.currentBranch?.name }}
      </div>
    </button>

    <!-- Dropdown -->
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform scale-95 opacity-0 translate-x-4"
      enter-to-class="transform scale-100 opacity-100 translate-x-0"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="transform scale-100 opacity-100 translate-x-0"
      leave-to-class="transform scale-95 opacity-0 translate-x-4"
    >
      <div 
        v-if="isOpen"
        :class="[
          'absolute mt-1 w-64 bg-white dark:bg-slate-900 rounded-2xl shadow-[0_20px_50px_rgba(0,0,0,0.3)] border border-slate-100 dark:border-slate-800 overflow-hidden z-[100] p-2 space-y-1',
          isCollapsed ? 'left-14 top-0' : 'left-2 right-2'
        ]"
      >
        <div class="px-3 py-2 mb-1">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Switch Location</p>
        </div>

        <button 
          v-if="authStore.user?.is_owner && isHistoryPage"
          @click="selectBranch(null)"
          class="w-full flex items-center justify-between px-4 py-3 rounded-xl transition-all mb-1 text-left"
          :class="authStore.selectedBranchID === null 
            ? 'bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400' 
            : 'text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800 hover:text-slate-900 dark:hover:text-slate-100'"
        >
          <div class="flex flex-col">
            <span class="text-sm font-bold">All Branches</span>
            <span class="text-[10px] opacity-60">Aggregated Organization View</span>
          </div>
          <Check v-if="authStore.selectedBranchID === null" class="w-4 h-4" />
        </button>
        
        <div v-if="authStore.user?.is_owner && isHistoryPage" class="h-px bg-slate-100 dark:bg-slate-800 mx-2 my-1"></div>
        
        <button 
          v-for="branch in authStore.branches" 
          :key="branch.id"
          @click="selectBranch(branch.id)"
          class="w-full flex items-center justify-between px-4 py-3 rounded-xl transition-all text-left"
          :class="authStore.selectedBranchID === branch.id 
            ? 'bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400' 
            : 'text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800 hover:text-slate-900 dark:hover:text-slate-100'"
        >
          <div class="flex flex-col min-w-0">
            <span class="text-sm font-bold truncate">{{ branch.name }}</span>
            <span v-if="branch.address" class="text-[10px] opacity-60 truncate">{{ branch.address }}</span>
          </div>
          <Check v-if="authStore.selectedBranchID === branch.id" class="w-4 h-4" />
        </button>
      </div>
    </Transition>
  </div>
</template>
