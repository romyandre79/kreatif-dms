<template>
  <div class="space-y-8 animate-in fade-in slide-in-from-bottom-4 duration-500">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <!-- Driver Selection -->
      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Scanner Driver / Protocol</label>
        <select v-model="localModel.driver" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-blue-500/10 focus:border-[#1E3A5F] transition-all">
          <option v-if="localModel.service_type === 'SCANNER_LOCAL'" value="wia">WIA (Windows Image Acquisition)</option>
          <option v-if="localModel.service_type === 'SCANNER_LOCAL'" value="twain">TWAIN (Legacy)</option>
          <option v-if="localModel.service_type === 'SCANNER_NETWORK'" value="ftp">FTP (Network Pull)</option>
          <option v-if="localModel.service_type === 'SCANNER_NETWORK'" value="smb">SMB / Shared Folder</option>
          <option v-if="localModel.service_type === 'SCANNER_NETWORK'" value="api">Direct API (REST)</option>
        </select>
      </div>

      <!-- Device Identifier -->
      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">
          {{ localModel.service_type === 'SCANNER_LOCAL' ? 'Local Device ID' : 'IP Address / Hostname' }}
        </label>
        <input v-model="localModel.endpoint" type="text" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-blue-500/10 focus:border-[#1E3A5F] transition-all" 
               :placeholder="localModel.service_type === 'SCANNER_LOCAL' ? 'e.g. {6BDD1FC6-810F-11D0-BEC7-08002BE2092F}\\0001' : 'e.g. 192.168.1.50'">
      </div>
    </div>

    <!-- Scanner Specific Config -->
    <div class="p-8 bg-white/50 dark:bg-slate-900/50 rounded-[2rem] border border-slate-100 dark:border-slate-800 space-y-6">
      <div class="flex items-center gap-3 mb-2">
        <LucideSettings class="w-4 h-4 text-primary-500" />
        <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Hardware Parameters</h4>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-500 uppercase tracking-widest">Default Resolution (DPI)</label>
          <select v-model="localModel.config.resolution" class="w-full px-4 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold">
            <option :value="150">150 DPI</option>
            <option :value="300">300 DPI</option>
            <option :value="600">600 DPI</option>
          </select>
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-500 uppercase tracking-widest">Color Mode</label>
          <select v-model="localModel.config.color_mode" class="w-full px-4 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold">
            <option value="color">Full Color</option>
            <option value="grayscale">Grayscale</option>
            <option value="bw">Black & White</option>
          </select>
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-500 uppercase tracking-widest">File Format</label>
          <select v-model="localModel.config.format" class="w-full px-4 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold">
            <option value="jpeg">JPEG (Compressed)</option>
            <option value="png">PNG (Lossless)</option>
            <option value="pdf">PDF (Document)</option>
          </select>
        </div>
      </div>

      <div v-if="localModel.service_type === 'SCANNER_NETWORK'" class="pt-6 border-t border-slate-100 dark:border-slate-800 grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-500 uppercase tracking-widest">Auth Username</label>
          <input v-model="localModel.config.username" type="text" class="w-full px-4 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold" placeholder="Optional">
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-500 uppercase tracking-widest">Auth Password</label>
          <input v-model="localModel.config.password" type="password" class="w-full px-4 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold" placeholder="Optional">
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { LucideSettings } from 'lucide-vue-next'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['update:modelValue'])

// Local reactive state to avoid direct prop mutation
const localModel = ref(JSON.parse(JSON.stringify(props.modelValue)))

// Update local state when prop changes (e.g. when opening a different node)
watch(() => props.modelValue, (newVal) => {
  localModel.value = JSON.parse(JSON.stringify(newVal))
}, { deep: true })

// Emit changes to parent
watch(localModel, (newVal) => {
  emit('update:modelValue', newVal)
}, { deep: true })
</script>
