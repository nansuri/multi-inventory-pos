<script setup lang="ts">
import { onMounted } from 'vue';
import { useConfigStore } from './stores/config';

const configStore = useConfigStore();

onMounted(() => {
  configStore.initTheme();
});
</script>

<template>
  <router-view />
</template>

<style>
/* Page Transitions */
.page-fade-enter-active,
.page-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.page-fade-enter-from {
  opacity: 0;
  transform: translateY(8px);
}

.page-fade-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}

/* Base styles if any */
</style>
