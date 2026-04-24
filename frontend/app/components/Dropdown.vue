<template>
  <div class="relative" v-on-click-outside="() => $emit('update:modelValue', false)">
    <div @click="$emit('update:modelValue', !modelValue)">
      <slot name="trigger" />
    </div>

    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform scale-95 opacity-0 -translate-y-2"
      enter-to-class="transform scale-100 opacity-100 translate-y-0"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="transform scale-100 opacity-100 translate-y-0"
      leave-to-class="transform scale-95 opacity-0 -translate-y-2"
    >
      <div 
        v-if="modelValue"
        class="absolute right-0 mt-3 w-64 bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border border-slate-200 dark:border-slate-800 py-2 z-50 overflow-hidden"
      >
        <slot />
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { vOnClickOutside } from '@vueuse/components'

defineProps({
  modelValue: Boolean
})

defineEmits(['update:modelValue'])
</script>
