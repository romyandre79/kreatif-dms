<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Bucket Name</label>
          <input 
            v-model="config.bucket" 
            type="text" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="e.g. krdms-vault"
          >
        </div>
        <div class="flex items-center gap-3 pt-8">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Use SSL (HTTPS)</p>
          <button 
            @click="config.use_ssl = !config.use_ssl" 
            :class="`w-12 h-6 rounded-full transition-all relative ${config.use_ssl ? 'bg-primary-500' : 'bg-slate-300 dark:bg-slate-700'}`"
          >
            <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${config.use_ssl ? 'left-7' : 'left-1'}`"></div>
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Access Key</label>
          <input 
            type="text" 
            v-model="config.access_key" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="Enter Access Key"
          >
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Secret Key</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.secret_key" 
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all pr-14" 
              placeholder="Enter Secret Key"
            >
            <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
              <LucideEye v-if="!showPassword" class="w-4 h-4" />
              <LucideEyeOff v-else class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { LucideEye, LucideEyeOff } from 'lucide-vue-next'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  }
})

const showPassword = ref(false)

const config = computed({
  get: () => props.modelValue.config || {},
  set: (val) => {
    props.modelValue.config = val
  }
})
</script>
