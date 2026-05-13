<template>
  <div class="fixed top-8 right-8 z-[200] flex flex-col gap-3 w-80 pointer-events-none">
    <TransitionGroup 
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="translate-x-full opacity-0 scale-95"
      enter-to-class="translate-x-0 opacity-100 scale-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="translate-x-0 opacity-100 scale-100"
      leave-to-class="translate-x-full opacity-0 scale-95"
    >
      <div 
        v-for="toast in toastStore.toasts" 
        :key="toast.id"
        class="pointer-events-auto bg-white dark:bg-[#0D121F] border rounded-2xl p-4 shadow-2xl flex items-start gap-4 ring-1 ring-black/5"
        :class="[
          toast.type === 'success' ? 'border-green-100 dark:border-green-900/30' : 
          toast.type === 'error' ? 'border-red-100 dark:border-red-900/30' :
          'border-slate-100 dark:border-slate-800'
        ]"
      >
        <div :class="[
          'w-10 h-10 rounded-xl flex items-center justify-center shrink-0 shadow-sm',
          toast.type === 'success' ? 'bg-green-50 text-green-500' : 
          toast.type === 'error' ? 'bg-red-50 text-red-500' :
          'bg-slate-50 text-slate-500'
        ]">
          <LucideCheckCircle2 v-if="toast.type === 'success'" class="w-5 h-5" />
          <LucideAlertCircle v-else-if="toast.type === 'error'" class="w-5 h-5" />
          <LucideInfo v-else class="w-5 h-5" />
        </div>
        
        <div class="flex-1 pt-1">
          <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ toast.type }}</p>
          <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed mt-0.5">{{ toast.message }}</p>
        </div>

        <button 
          @click="toastStore.removeToast(toast.id)"
          class="p-1 hover:bg-slate-50 dark:hover:bg-slate-800 rounded-lg text-slate-300 transition-colors"
        >
          <LucideX class="w-4 h-4" />
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup>
import { useToastStore } from '~/stores/toast'
import { 
  LucideCheckCircle2, 
  LucideAlertCircle, 
  LucideInfo, 
  LucideX 
} from 'lucide-vue-next'

const toastStore = useToastStore()
</script>
