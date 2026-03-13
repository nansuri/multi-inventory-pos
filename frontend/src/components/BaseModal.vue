<script setup lang="ts">
import { X } from 'lucide-vue-next';

defineProps<{
  show: boolean;
  title: string;
  subtitle?: string;
  icon?: any;
  maxWidth?: string;
}>();

defineEmits(['close']);
</script>

<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div v-if="show" class="fixed inset-0 z-50 overflow-y-auto bg-slate-900/40 dark:bg-slate-950/60 backdrop-blur-sm">
        <div class="flex min-h-full items-center justify-center p-4 text-center">
          <Transition
            enter-active-class="transition duration-300 ease-out"
            enter-from-class="opacity-0 scale-95"
            enter-to-class="opacity-100 scale-100"
            leave-active-class="transition duration-200 ease-in"
            leave-from-class="opacity-100 scale-100"
            leave-to-class="opacity-0 scale-95"
          >
            <div 
              class="w-full transform overflow-hidden rounded-[2rem] bg-white dark:bg-slate-900 p-8 text-left align-middle shadow-2xl transition-all border border-slate-100 dark:border-slate-800"
              :class="maxWidth || 'max-w-xl'"
            >
              <div class="flex items-start justify-between mb-6">
                <div class="flex items-center gap-4">
                  <div v-if="icon" class="p-3 bg-indigo-50 dark:bg-indigo-900/30 rounded-xl text-indigo-600 dark:text-indigo-400">
                    <component :is="icon" class="w-5 h-5" />
                  </div>
                  <div>
                    <h3 class="text-xl font-black text-slate-800 dark:text-slate-100 leading-none mb-1">
                      {{ title }}
                    </h3>
                    <p v-if="subtitle" class="text-xs font-medium text-slate-400 dark:text-slate-500">
                      {{ subtitle }}
                    </p>
                  </div>
                </div>
                <button 
                  @click="$emit('close')"
                  class="p-1.5 text-slate-300 dark:text-slate-600 hover:text-slate-600 dark:hover:text-slate-300 hover:bg-slate-50 dark:hover:bg-slate-800 rounded-lg transition-all"
                >
                  <X class="w-5 h-5" />
                </button>
              </div>

              <slot />
            </div>
          </Transition>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
