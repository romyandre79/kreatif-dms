<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <div class="space-y-3 text-left">
        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">AI Provider</label>
        <div class="grid grid-cols-2 gap-4">
          <button 
            @click="modelValue.driver = 'gemini'; config.model = 'gemini-1.5-flash'" 
            :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${modelValue.driver === 'gemini' ? 'bg-blue-50 border-blue-200 text-blue-600 dark:bg-blue-900/20 dark:border-blue-800' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-400'}`"
          >
            <LucideCpu class="w-4 h-4" /> Google Gemini
          </button>
          <button 
            @click="modelValue.driver = 'openai'; config.model = 'gpt-4o'" 
            :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${modelValue.driver === 'openai' ? 'bg-blue-50 border-blue-200 text-blue-600 dark:bg-blue-900/20 dark:border-blue-800' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-400'}`"
          >
            <LucideCpu class="w-4 h-4" /> OpenAI
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">API Key</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.token" 
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all pr-14" 
              placeholder="Enter API Key"
            >
            <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
              <LucideEye v-if="!showPassword" class="w-4 h-4" />
              <LucideEyeOff v-else class="w-4 h-4" />
            </button>
          </div>
        </div>
        <div class="space-y-2">
          <div class="flex items-center justify-between">
            <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Model Name</label>
            <button 
              @click="$emit('fetch-models')" 
              :disabled="fetching" 
              class="text-[9px] font-black text-blue-500 uppercase tracking-widest hover:text-blue-600 disabled:opacity-50 flex items-center gap-1 transition-colors"
            >
              <LucideRefreshCw class="w-3 h-3" :class="{ 'animate-spin': fetching }" />
              Refresh
            </button>
          </div>
          <select 
            v-model="config.model" 
            class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
          >
            <option v-for="m in availableModels" :key="m.id" :value="m.id">
              {{ m.name }}
            </option>
          </select>
        </div>
      </div>

      <div class="space-y-2 text-left">
        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">System Instruction / Prompt</label>
        <textarea 
          v-model="config.system_prompt" 
          rows="4" 
          class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-medium leading-relaxed resize-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" 
          placeholder="Default system prompt for AI responses..."
        ></textarea>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { LucideCpu, LucideEye, LucideEyeOff, LucideRefreshCw } from 'lucide-vue-next'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  },
  aiModels: {
    type: Object,
    default: () => ({})
  },
  fetching: {
    type: Boolean,
    default: false
  }
})

defineEmits(['fetch-models'])

const showPassword = ref(false)

const config = computed({
  get: () => props.modelValue.config || {},
  set: (val) => {
    props.modelValue.config = val
  }
})

const availableModels = computed(() => {
  return props.aiModels[props.modelValue.driver] || []
})
</script>
