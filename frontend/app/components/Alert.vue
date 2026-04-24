<template>
  <Transition
    enter-active-class="transition duration-300 ease-out"
    enter-from-class="transform -translate-y-4 opacity-0"
    enter-to-class="transform translate-y-0 opacity-100"
    leave-active-class="transition duration-200 ease-in"
    leave-from-class="transform translate-y-0 opacity-100"
    leave-to-class="transform -translate-y-4 opacity-0"
  >
    <div 
      v-if="modelValue"
      :class="[
        'flex items-center gap-3 p-4 rounded-xl shadow-lg border relative',
        type === 'error' ? 'bg-red-500 border-red-600 text-white' : 'bg-green-500 border-green-600 text-white'
      ]"
    >
      <div class="flex-shrink-0">
        <LucideAlertCircle v-if="type === 'error'" class="w-5 h-5" />
        <LucideCheckCircle v-else class="w-5 h-5" />
      </div>
      
      <div class="flex-1 text-sm font-medium">
        {{ message }}
      </div>

      <button 
        @click="$emit('update:modelValue', false)"
        class="flex-shrink-0 p-1 hover:bg-white/20 rounded-lg transition-colors"
      >
        <LucideX class="w-4 h-4" />
      </button>
    </div>
  </Transition>
</template>

<script setup>
import { LucideAlertCircle, LucideCheckCircle, LucideX } from 'lucide-vue-next'

defineProps({
  modelValue: Boolean,
  message: String,
  type: {
    type: String,
    default: 'error'
  }
})

defineEmits(['update:modelValue'])
</script>
