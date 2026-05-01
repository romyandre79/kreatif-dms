<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Sender Email</label>
          <input 
            type="email" 
            v-model="config.from_email" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="noreply@company.com"
          >
        </div>
        <div class="flex items-center gap-3 pt-8">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Enable Authentication</p>
          <button 
            @click="config.auth = !config.auth" 
            :class="`w-12 h-6 rounded-full transition-all relative ${config.auth ? 'bg-primary-500' : 'bg-slate-300 dark:bg-slate-700'}`"
          >
            <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${config.auth ? 'left-7' : 'left-1'}`"></div>
          </button>
        </div>
      </div>

      <div v-if="config.auth" class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left" v-motion-slide-visible-top>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">SMTP Username</label>
          <input 
            type="text" 
            v-model="config.user" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="Enter username"
          >
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">SMTP Password</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.pass" 
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all pr-14" 
              placeholder="Enter password"
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
