<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Auth Method</label>
          <div class="grid grid-cols-2 gap-4">
            <button 
              @click="modelValue.driver = 'basic'" 
              :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all ${modelValue.driver === 'basic' ? 'bg-blue-50 border-blue-200 text-blue-600 dark:bg-blue-900/20 dark:border-blue-800' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-400'}`"
            >
              Basic Auth
            </button>
            <button 
              @click="modelValue.driver = 'apikey'" 
              :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all ${modelValue.driver === 'apikey' ? 'bg-blue-50 border-blue-200 text-blue-600 dark:bg-blue-900/20 dark:border-blue-800' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-400'}`"
            >
              API Key
            </button>
          </div>
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Index Name</label>
          <input 
            type="text" 
            v-model="config.index_name" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="e.g. documents"
          >
        </div>
      </div>

      <div v-if="modelValue.driver === 'basic'" class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left" v-motion-slide-visible-top>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Username</label>
          <input 
            type="text" 
            v-model="config.username" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="elastic"
          >
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Password</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.password" 
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all pr-14"
            >
            <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
              <LucideEye v-if="!showPassword" class="w-4 h-4" />
              <LucideEyeOff v-else class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>

      <div v-if="modelValue.driver === 'apikey'" class="space-y-2 text-left" v-motion-slide-visible-top>
        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Elastic API Key</label>
        <input 
          type="text" 
          v-model="config.api_key" 
          class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
          placeholder="Enter Base64 API Key"
        >
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
