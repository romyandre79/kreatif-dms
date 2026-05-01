<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">WhatsApp Provider</label>
          <select 
            v-model="modelValue.driver" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
          >
            <option value="fonnte">Fonnte.com</option>
            <option value="wablas">Wablas.com</option>
            <option value="generic">Generic Webhook / API</option>
          </select>
        </div>
      </div>

      <div v-if="modelValue.driver === 'fonnte' || modelValue.driver === 'wablas'" class="space-y-6" v-motion-slide-visible-top>
        <div class="p-6 bg-blue-50 dark:bg-blue-900/20 rounded-2xl border border-blue-100 dark:border-blue-800/30 flex items-start gap-4 text-left">
          <div class="p-3 bg-white dark:bg-slate-800 rounded-xl shadow-sm">
            <LucideMessageSquare class="w-6 h-6 text-blue-600" />
          </div>
          <div class="space-y-1">
            <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ modelValue.driver === 'fonnte' ? 'Fonnte' : 'Wablas' }} Integration</h4>
            <p class="text-xs text-slate-500 leading-relaxed italic">Masukkan API Token Anda untuk mengaktifkan notifikasi WhatsApp otomatis via gateway {{ modelValue.driver === 'fonnte' ? 'Fonnte' : 'Wablas' }}.</p>
          </div>
        </div>

        <div class="space-y-2 text-left">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">API Token / API Key</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.token" 
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all pr-14" 
              placeholder="Enter your API Token"
            >
            <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
              <LucideEye v-if="!showPassword" class="w-4 h-4" />
              <LucideEyeOff v-else class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>

      <div v-if="modelValue.driver === 'generic'" class="space-y-6 text-left" v-motion-slide-visible-top>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Custom API Endpoint</label>
          <input 
            type="text" 
            v-model="modelValue.endpoint" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="https://api.yourprovider.com/send"
          >
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Auth Header (Optional)</label>
          <input 
            type="text" 
            v-model="config.auth_header" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
            placeholder="e.g. Bearer your_token"
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { LucideMessageSquare, LucideEye, LucideEyeOff } from 'lucide-vue-next'

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
