<template>
  <Teleport to="body" v-if="modelValue">
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div v-if="modelValue" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
        <Transition
          enter-active-class="transition duration-300 ease-out"
          enter-from-class="transform scale-95 opacity-0"
          enter-to-class="transform scale-100 opacity-100"
          leave-active-class="transition duration-200 ease-in"
          leave-from-class="transform scale-100 opacity-100"
          leave-to-class="transform scale-95 opacity-0"
        >
          <div 
            v-if="modelValue"
            class="bg-white dark:bg-slate-900 rounded-2xl shadow-2xl w-full max-w-md overflow-hidden border border-slate-200 dark:border-slate-800"
            @click.stop
          >
            <!-- Header -->
            <div class="px-6 py-4 flex items-center justify-between border-b border-slate-100 dark:border-slate-800">
              <h3 class="text-xl font-bold text-slate-900 dark:text-white">{{ title }}</h3>
              <button @click="$emit('update:modelValue', false)" class="text-slate-400 hover:text-slate-600 transition-colors p-1 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800">
                <LucideX class="w-5 h-5" />
              </button>
            </div>

            <!-- Content -->
            <div class="p-6">
              <slot />
            </div>

            <!-- Footer -->
            <div v-if="$slots.footer" class="px-6 py-4 bg-slate-50 dark:bg-slate-800/50 flex justify-end gap-3">
              <slot name="footer" />
            </div>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { LucideX } from 'lucide-vue-next'

defineProps({
  modelValue: Boolean,
  title: String
})

defineEmits(['update:modelValue'])
</script>
