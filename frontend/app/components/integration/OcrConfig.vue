<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <!-- Security Info Banner -->
      <div class="p-6 bg-amber-50 dark:bg-amber-900/20 rounded-2xl border border-amber-100 dark:border-amber-800/30 flex items-start gap-4 text-left">
        <div class="p-3 bg-white dark:bg-slate-800 rounded-xl shadow-sm">
          <LucideScanLine class="w-6 h-6 text-amber-600" />
        </div>
        <div class="space-y-1">
          <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">OCR Service Security</h4>
          <p class="text-xs text-slate-500 leading-relaxed italic">Konfigurasi batasan host dan kredensial untuk akses ke engine OCR (Tesseract/Cloud Service).</p>
        </div>
      </div>

      <!-- Main Config -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Allowed Client Host</label>
          <input 
            type="text" 
            v-model="config.client_host" 
            placeholder="e.g. 192.168.1.100 atau *" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-amber-500/10 focus:border-amber-500 transition-all"
          />
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">OCR Driver</label>
          <select 
            v-model="modelValue.driver" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-amber-500/10 focus:border-amber-500 transition-all"
          >
            <option value="tesseract">Tesseract (Local)</option>
            <option value="google_vision">Google Vision API</option>
            <option value="custom">Custom OCR Service</option>
          </select>
        </div>
      </div>

      <!-- Auth Config -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Auth Username</label>
          <input 
            type="text" 
            v-model="config.user" 
            placeholder="Username akses service"
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-amber-500/10 focus:border-amber-500 transition-all"
          />
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Auth Password</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.password" 
              placeholder="••••••••"
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-amber-500/10 focus:border-amber-500 transition-all pr-14"
            />
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
import { LucideScanLine, LucideEye, LucideEyeOff } from 'lucide-vue-next'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  }
})

const showPassword = ref(false)

// Ensure config object exists
const config = computed({
  get: () => props.modelValue.config || {},
  set: (val) => {
    props.modelValue.config = val
  }
})
</script>
