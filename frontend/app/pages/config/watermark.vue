<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
          {{ $t('admin.watermark.title') || 'Watermark Configuration' }}
        </h1>
        <p class="text-sm text-slate-500 mt-1 uppercase font-bold tracking-widest text-[10px]">
          Manage security overlays for document previews
        </p>
      </div>
      <button @click="saveSettings" 
              :disabled="saving"
              class="px-8 py-4 bg-primary-500 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-primary-500/20 hover:bg-primary-600 transition-all flex items-center gap-3 disabled:opacity-50">
        <LucideSave v-if="!saving" class="w-4 h-4" />
        <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
        {{ saving ? 'Saving...' : 'Save Configuration' }}
      </button>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Settings -->
      <div class="lg:col-span-5 space-y-8">
        <div class="glass p-8 rounded-lg space-y-8 border border-white/10 shadow-sm bg-white/50 dark:bg-slate-900/50 backdrop-blur-xl">
          <div class="flex items-center gap-4 border-b border-slate-100 dark:border-slate-800 pb-6">
            <div class="w-10 h-10 rounded-xl bg-primary-500/10 flex items-center justify-center">
              <LucideSettings2 class="w-5 h-5 text-primary-500" />
            </div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">General Settings</h3>
          </div>

          <div class="space-y-6">
            <!-- Type Switcher -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Watermark Type</label>
              <div class="grid grid-cols-2 gap-4">
                <button v-for="t in ['text', 'image']" :key="t"
                        @click="settings.type = t"
                        :class="settings.type === t ? 'bg-primary-500 text-white shadow-lg shadow-primary-500/20' : 'bg-slate-100 dark:bg-slate-800 text-slate-500'"
                        class="py-3 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all">
                  {{ t }}
                </button>
              </div>
            </div>

            <!-- Text Content -->
            <div v-if="settings.type === 'text'" class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Watermark Text</label>
              <input v-model="settings.text" 
                     type="text"
                     placeholder="e.g. CONFIDENTIAL - {user} - {date}"
                     class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-medium focus:ring-2 focus:ring-primary-500 transition-all" />
              <p class="text-[9px] text-slate-400 italic">Available variables: {user}, {date}, {ip}</p>
            </div>

            <!-- Image Upload -->
            <div v-if="settings.type === 'image'" class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Watermark Image (Logo)</label>
              <div class="border-2 border-dashed border-slate-200 dark:border-slate-800 rounded-2xl p-6 text-center hover:border-primary-500 transition-all cursor-pointer group">
                <LucideUploadCloud class="w-8 h-8 text-slate-300 group-hover:text-primary-500 mx-auto mb-2" />
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Click to upload PNG</p>
              </div>
            </div>

            <!-- Opacity Slider -->
            <div class="space-y-3">
              <div class="flex justify-between items-center">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Opacity ({{ (settings.opacity * 100).toFixed(0) }}%)</label>
              </div>
              <input v-model.number="settings.opacity" type="range" min="0" max="1" step="0.1" class="w-full h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full appearance-none cursor-pointer accent-primary-500" />
            </div>

            <!-- Position -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Positioning Style</label>
              <div class="grid grid-cols-3 gap-2">
                <button v-for="pos in ['center', 'diagonal', 'tiled']" :key="pos"
                        @click="settings.position = pos"
                        :class="settings.position === pos ? 'bg-primary-500/10 text-primary-500 border-primary-500/50' : 'bg-transparent text-slate-400 border-slate-100 dark:border-slate-800'"
                        class="py-2 border-2 rounded-lg text-[9px] font-black uppercase tracking-widest transition-all">
                  {{ pos }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Live Preview -->
      <div class="lg:col-span-7">
        <div class="glass rounded-lg overflow-hidden border border-white/10 shadow-2xl h-full flex flex-col bg-white dark:bg-slate-900">
          <div class="p-6 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Live Visual Preview</h3>
            <div class="flex gap-2">
              <div class="w-3 h-3 rounded-full bg-red-400"></div>
              <div class="w-3 h-3 rounded-full bg-amber-400"></div>
              <div class="w-3 h-3 rounded-full bg-green-400"></div>
            </div>
          </div>
          
          <div class="flex-1 bg-slate-100 dark:bg-[#05080F] p-10 flex items-center justify-center relative overflow-hidden group">
            <!-- Mock Document -->
            <div class="w-[300px] h-[400px] bg-white dark:bg-slate-800 shadow-2xl rounded-sm p-8 space-y-4 relative overflow-hidden select-none transition-transform duration-700 group-hover:scale-[1.02]">
              <div class="h-4 w-3/4 bg-slate-100 dark:bg-slate-700 rounded-full mb-8"></div>
              <div v-for="i in 12" :key="i" class="h-2 bg-slate-50 dark:bg-slate-700/50 rounded-full" :style="{ width: Math.random() * 50 + 50 + '%' }"></div>
              
              <!-- Real-time Watermark Overlay -->
              <div class="absolute inset-0 pointer-events-none flex items-center justify-center" 
                   :style="{ opacity: settings.opacity }">
                <!-- Tiled Style -->
                <div v-if="settings.position === 'tiled'" class="grid grid-cols-3 gap-20 -rotate-12 scale-150">
                  <div v-for="i in 9" :key="i" class="text-[8px] font-black uppercase whitespace-nowrap text-slate-400 dark:text-slate-500">
                    {{ previewText }}
                  </div>
                </div>
                <!-- Diagonal Style -->
                <div v-else-if="settings.position === 'diagonal'" 
                     class="-rotate-45 text-2xl font-black uppercase whitespace-nowrap text-slate-400 dark:text-slate-500 tracking-tighter">
                  {{ previewText }}
                </div>
                <!-- Center Style -->
                <div v-else class="text-xl font-black uppercase text-slate-400 dark:text-slate-500 tracking-tighter">
                  {{ previewText }}
                </div>
              </div>
            </div>

            <!-- Zoom Indicator -->
            <div class="absolute bottom-6 right-6 px-4 py-2 bg-black/50 backdrop-blur-md rounded-full text-[10px] font-black text-white uppercase tracking-widest border border-white/10">
              Simulation Mode
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  LucideSave, LucideLoader2, LucideSettings2, LucideType, 
  LucideImage, LucideUploadCloud, LucideLayers 
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const config = useRuntimeConfig()
const saving = ref(false)

const settings = ref({
  type: 'text',
  text: 'CONFIDENTIAL - {user} - {date}',
  image_url: '',
  opacity: 0.3,
  position: 'diagonal',
  font_size: 24,
  color: '#FFFFFF'
})

const previewText = computed(() => {
  let t = settings.value.text
  t = t.replace('{user}', 'ADMIN')
  t = t.replace('{date}', new Date().toLocaleDateString())
  return t
})

const fetchSettings = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/settings/watermark`)
    if (res && res.data) {
      settings.value = { ...settings.value, ...res.data }
    }
  } catch (err) {
    console.error('Failed to fetch watermark settings')
  }
}

const saveSettings = async () => {
  saving.value = true
  try {
    await $api(`${config.public.apiBase}/master/settings/watermark`, {
      method: 'POST',
      body: settings.value
    })
    // Show success toast (implement if needed)
  } catch (err) {
    alert('Failed to save settings: ' + err.message)
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  fetchSettings()
})

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
}
.dark .glass {
  background: rgba(15, 23, 42, 0.7);
}
</style>
