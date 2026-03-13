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
/* Base styles if any */
</style>
