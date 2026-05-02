<template>
  <div class="space-y-8 animate-in fade-in slide-in-from-bottom-4 duration-700">
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 bg-white dark:bg-[#0D121F] p-8 rounded-[2.5rem] border border-slate-200/60 dark:border-slate-800/40 shadow-sm relative overflow-hidden">
      <div class="absolute top-0 right-0 p-8 opacity-10">
        <LucideScanLine class="w-32 h-32 text-primary-500" />
      </div>
      <div class="relative z-10">
        <div class="flex items-center gap-3 mb-2">
          <div class="p-2 bg-primary-500/10 rounded-xl">
            <LucideScanLine class="w-6 h-6 text-primary-500" />
          </div>
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
            OCR Intelligence <span class="text-primary-500">Playground</span>
          </h1>
        </div>
        <p class="text-slate-500 dark:text-slate-400 text-sm font-medium">
          Test and monitor the high-performance document extraction engine in real-time.
        </p>
      </div>
      
      <div class="flex items-center gap-4 relative z-10">
        <button @click="triggerUpload" :disabled="isUploading"
                class="flex items-center gap-3 px-8 py-4 bg-primary-500 text-white rounded-2xl font-black text-xs uppercase tracking-widest hover:bg-primary-600 hover:shadow-xl hover:shadow-primary-500/20 active:scale-95 transition-all disabled:opacity-50">
          <LucideUploadCloud v-if="!isUploading" class="w-4 h-4" />
          <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
          {{ isUploading ? 'Extracting...' : 'Upload Test File' }}
        </button>
        <input type="file" ref="fileInput" @change="handleUpload" class="hidden" accept=".pdf,.jpg,.png,.jpeg">
      </div>
    </div>

    <!-- Stats Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div v-for="stat in stats" :key="stat.label" 
           class="bg-white dark:bg-[#0D121F] p-6 rounded-[2rem] border border-slate-200/60 dark:border-slate-800/40 shadow-sm group hover:border-primary-500/30 transition-all duration-500">
        <div class="flex items-center gap-4">
          <div :class="['w-12 h-12 rounded-2xl flex items-center justify-center transition-transform group-hover:scale-110', stat.bg]">
            <component :is="stat.icon" :class="['w-5 h-5', stat.color]" />
          </div>
          <div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-0.5">{{ stat.label }}</p>
            <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ stat.value }}</h3>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Recent History (Table) -->
      <div class="lg:col-span-8 bg-white dark:bg-[#0D121F] rounded-[2.5rem] border border-slate-200/60 dark:border-slate-800/40 shadow-sm flex flex-col overflow-hidden">
        <div class="p-8 border-b border-slate-100 dark:border-slate-800/60 bg-slate-50/30 dark:bg-slate-900/20 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <LucideHistory class="w-5 h-5 text-primary-500" />
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Processing History</h3>
          </div>
          <button @click="fetchHistory" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
            <LucideRefreshCw class="w-4 h-4 text-slate-400" />
          </button>
        </div>
        
        <div class="flex-1 overflow-x-auto custom-scrollbar">
          <table class="w-full text-left border-collapse">
            <thead class="text-[10px] font-black text-slate-400 uppercase tracking-widest bg-slate-50/50 dark:bg-slate-900/40">
              <tr>
                <th class="py-5 px-8">File Name</th>
                <th class="py-5 px-6">Size</th>
                <th class="py-5 px-6">Duration</th>
                <th class="py-5 px-6">Accuracy</th>
                <th class="py-5 px-6">Status</th>
                <th class="py-5 px-8 text-right">Action</th>
              </tr>
            </thead>
            <tbody class="text-sm">
              <tr v-if="history.length === 0" class="border-t border-slate-100 dark:border-slate-800">
                <td colspan="5" class="py-20 text-center text-slate-400 font-bold uppercase tracking-widest text-[10px]">No processing history yet</td>
              </tr>
              <tr v-for="item in history" :key="item.id" 
                  class="border-t border-slate-100 dark:border-slate-800 hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-colors group">
                <td class="py-5 px-8 font-bold text-slate-700 dark:text-slate-300">{{ item.filename }}</td>
                <td class="py-5 px-6 text-slate-500 text-xs">{{ item.size }}</td>
                <td class="py-5 px-6 font-black text-xs text-primary-500">{{ item.duration.toFixed(2) }}s</td>
                <td class="py-5 px-6">
                  <span :class="['font-black text-xs', 
                             item.accuracy >= 0.8 ? 'text-green-500' : (item.accuracy >= 0.5 ? 'text-orange-500' : 'text-red-500')]">
                    {{ (item.accuracy * 100).toFixed(1) }}%
                  </span>
                </td>
                <td class="py-5 px-6">
                  <span :class="['px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest', 
                             item.status === 'Success' ? 'bg-green-500/10 text-green-500' : 'bg-red-500/10 text-red-500']">
                    {{ item.status }}
                  </span>
                </td>
                <td class="py-5 px-8 text-right">
                  <button v-if="item.status === 'Success'" @click="openVisualizer(item.id)"
                          class="p-2 bg-primary-500/10 text-primary-500 hover:bg-primary-500 hover:text-white rounded-xl transition-all scale-0 group-hover:scale-100">
                    <LucideEye class="w-4 h-4" />
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Live Console -->
      <div class="lg:col-span-4 bg-[#0A0F1C] rounded-[2.5rem] border border-slate-800/50 shadow-2xl flex flex-col overflow-hidden h-[600px] ring-1 ring-white/5">
        <div class="p-6 border-b border-slate-800/50 bg-[#0D121F] flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Engine Live Console</h3>
          </div>
          <button @click="logs = []" class="text-[10px] font-bold text-slate-500 hover:text-white transition-colors">Clear</button>
        </div>
        <div class="flex-1 p-6 overflow-y-auto custom-scrollbar font-mono text-[11px] leading-relaxed" ref="logContainer">
          <div v-for="(log, i) in logs" :key="i" class="mb-1 text-slate-400 break-all">
            <span class="text-primary-500/50 mr-2">[{{ new Date().toLocaleTimeString() }}]</span> {{ log }}
          </div>
          <div v-if="logs.length === 0" class="h-full flex items-center justify-center text-slate-600 italic">
            Waiting for engine pulse...
          </div>
        </div>
      </div>
    </div>

    <!-- Visualizer Modal -->
    <Transition name="scale">
      <div v-if="showVisualizer" class="fixed inset-0 z-[200] flex items-center justify-center p-4 lg:p-12">
        <div class="absolute inset-0 bg-[#020617]/95 backdrop-blur-xl" @click="closeVisualizer"></div>
        
        <div class="relative w-full h-full glass rounded-[3rem] border border-white/10 shadow-2xl flex flex-col overflow-hidden">
            <!-- Modal Header -->
            <div class="p-8 border-b border-white/5 bg-white/5 flex items-center justify-between shrink-0">
              <div class="flex items-center gap-6">
                <div>
                  <h2 class="text-2xl font-black text-white uppercase tracking-tight leading-none mb-2">{{ activeJob?.filename }}</h2>
                  <p class="text-slate-400 text-xs font-bold uppercase tracking-widest">
                    {{ activeJob?.timestamp }} | Status: <span class="text-green-500">{{ activeJob?.status }}</span>
                  </p>
                </div>
                <div v-if="activeJob?.ai_analysis && activeJob.ai_analysis !== 'null'" class="h-10 w-px bg-white/10 hidden md:block"></div>
                <div v-if="activeJob?.ai_analysis && activeJob.ai_analysis !== 'null'" class="hidden md:block">
                  <span class="px-4 py-2 bg-primary-500 text-white text-[10px] font-black rounded-xl uppercase tracking-widest shadow-lg shadow-primary-500/20">
                    {{ parseAI(activeJob.ai_analysis).doc_type || 'Document' }}
                  </span>
                </div>
              </div>
              <button @click="closeVisualizer" class="w-12 h-12 bg-white/10 hover:bg-red-500/20 text-white hover:text-red-500 rounded-2xl flex items-center justify-center transition-all">
                <LucideX class="w-6 h-6" />
              </button>
            </div>
            
            <!-- AI Insights Panel (if available) -->
            <div v-if="activeJob?.ai_analysis && activeJob.ai_analysis !== 'null' && parseAI(activeJob.ai_analysis).summary" class="px-8 py-4 bg-primary-500/5 border-b border-white/5 flex flex-wrap items-center gap-6 overflow-x-auto custom-scrollbar no-scrollbar">
              <div class="flex items-center gap-3">
                <LucideHistory class="w-4 h-4 text-primary-500" />
                <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">History:</span>
              </div>
            </div>

            <!-- AI Cleaned Text Section -->
            <div v-if="parseAI(activeJob?.ai_analysis).cleaned_text" class="px-8 py-6 bg-slate-900/40 border-b border-white/5">
              <div class="flex items-center gap-3 mb-4">
                <div class="p-1.5 bg-green-500/10 rounded-lg"><LucideCheckCircle2 class="w-4 h-4 text-green-500" /></div>
                <h4 class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-400">AI Intelligent Text Recovery</h4>
              </div>
              <div class="p-6 bg-black/40 rounded-[1.5rem] border border-white/5 text-sm text-slate-300 leading-relaxed font-medium whitespace-pre-wrap max-h-48 overflow-y-auto custom-scrollbar">
                {{ parseAI(activeJob.ai_analysis).cleaned_text }}
              </div>
            </div>
              <p class="text-xs font-medium text-slate-300 italic">"{{ parseAI(activeJob.ai_analysis).summary }}"</p>
              <div class="flex gap-2">
                <div v-for="(val, key) in parseAI(activeJob.ai_analysis).entities" :key="key" 
                     class="px-3 py-1 bg-white/5 rounded-lg border border-white/10 flex items-center gap-2">
                  <span class="text-[9px] font-black text-slate-500 uppercase">{{ key }}:</span>
                  <span class="text-[10px] font-bold text-primary-400">{{ val }}</span>
                </div>
              </div>
            </div>
          
          <!-- Modal Content -->
          <div class="flex-1 grid grid-cols-1 lg:grid-cols-12 gap-8 p-8 overflow-hidden">
            <!-- Left: Image View (Taking 8 columns) -->
            <div class="lg:col-span-8 bg-[#020617] rounded-[2rem] relative overflow-auto custom-scrollbar flex items-start justify-center p-8 group">
              <div class="relative">
                <img v-if="activeJob" :src="`${ocrUrl}${activeJob.preview_path || activeJob.file_path}`" 
                     ref="visualizerImage"
                     @load="onImageLoad"
                     class="max-w-none shadow-2xl rounded-sm ring-1 ring-white/10" 
                     alt="Document">
                <div class="absolute inset-0 z-10 overflow-hidden pointer-events-none">
                  <div v-for="(word, i) in words" :key="i"
                       :style="getBoxStyle(word.box)"
                       @mouseenter="onHoverBox(i)"
                       @mouseleave="onLeaveBox(i)"
                       @click="onBoxClick(i)"
                       :class="['absolute border-2 border-primary-500/40 bg-primary-500/10 hover:bg-primary-500/40 hover:border-primary-400 cursor-pointer transition-all pointer-events-auto',
                                { 'ring-4 ring-white z-20 scale-[1.05]': hoveredIndex === i }]"
                       :data-index="i">
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Right: Text List (Taking 4 columns) -->
            <div class="lg:col-span-4 bg-white/5 rounded-[2rem] flex flex-col overflow-hidden border border-white/5">
              <div class="p-6 border-b border-white/5 bg-white/5 flex items-center justify-between">
                <h4 class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-400">Extracted Intelligence</h4>
                <button @click="toggleHighlightAll" 
                        :class="['px-4 py-1.5 text-[10px] font-black rounded-lg transition-all', 
                                highlightAll ? 'bg-primary-500 text-white' : 'bg-primary-500/10 text-primary-500 border border-primary-500/20']">
                  {{ highlightAll ? 'Hide All' : 'Highlight All' }}
                </button>
              </div>
              <div class="flex-1 p-6 overflow-y-auto custom-scrollbar space-y-3 scroll-smooth" ref="textContainer">
                <div v-for="(word, i) in words" :key="i"
                     :id="`text-item-${i}`"
                     @mouseenter="onHoverText(i)"
                     @mouseleave="onLeaveText(i)"
                     class="group p-4 rounded-xl border border-white/5 transition-all cursor-pointer"
                     :class="{ 'bg-primary-500/20 border-primary-500 scale-[1.02]': hoveredIndex === i, 'hover:bg-white/5': hoveredIndex !== i }">
                  <div class="flex items-center justify-between mb-1">
                    <span class="text-[9px] font-black text-slate-500 uppercase">Page {{ word.page }} | Conf: {{ (word.confidence * 100).toFixed(1) }}%</span>
                  </div>
                  <p class="text-sm font-medium text-slate-300 group-hover:text-white transition-colors">{{ word.text }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { 
  LucideScanLine, LucideUploadCloud, LucideLoader2, LucideHistory, 
  LucideRefreshCw, LucideEye, LucideActivity, LucideClock, LucideCheckCircle2, 
  LucideX, LucideMonitor, LucideZap
} from 'lucide-vue-next'

const config = useRuntimeConfig()
const ocrUrl = config.public.ocrUrl || 'http://localhost:8000'
const fileInput = ref(null)
const logContainer = ref(null)
const textContainer = ref(null)
const isUploading = ref(false)
const history = ref([])
const logs = ref([])
const statsData = ref({ total: 0, avg_time: 0, success_rate: 100 })
const showVisualizer = ref(false)
const activeJob = ref(null)
const words = ref([])
const hoveredIndex = ref(null)
const highlightAll = ref(true)
const ws = ref(null)

const stats = computed(() => [
  { label: 'Total Process', value: statsData.value.total, icon: LucideActivity, bg: 'bg-primary-500/10', color: 'text-primary-500' },
  { label: 'Avg Time', value: `${statsData.value.avg_time.toFixed(2)}s`, icon: LucideClock, bg: 'bg-sky-500/10', color: 'text-sky-500' },
  { label: 'Success Rate', value: `${statsData.value.success_rate.toFixed(1)}%`, icon: LucideCheckCircle2, bg: 'bg-green-500/10', color: 'text-green-500' },
  { label: 'Service Status', value: 'ONLINE', icon: LucideMonitor, bg: 'bg-purple-500/10', color: 'text-purple-500' },
])

const triggerUpload = () => fileInput.value.click()

const handleUpload = async (e) => {
  const file = e.target.files[0]
  if (!file) return

  isUploading.value = true
  addLog(`[UI] Uploading ${file.name} to OCR Service...`)

  const formData = new FormData()
  formData.append('file', file)

  try {
    const response = await fetch(`${ocrUrl}/ocr/process`, {
      method: 'POST',
      body: formData,
      // Simple auth for playground
      headers: {
        'Authorization': 'Basic ' + btoa('admin:admin123')
      }
    })

    if (!response.ok) throw new Error(`HTTP Error: ${response.status}`)
    
    const result = await response.json()
    addLog(`[UI] Extraction successful for ${file.name}`, 'text-green-400')
    fetchHistory()
    
    // Auto-open visualizer if it's a small file
    if (result.words && result.words.length > 0) {
      setTimeout(() => openVisualizer(null, result), 500)
    }
  } catch (err) {
    addLog(`[UI] Error: ${err.message}`, 'text-red-500')
    console.error(err)
  } finally {
    isUploading.value = false
    e.target.value = ''
  }
}

const fetchHistory = async () => {
  try {
    const res = await fetch(`${ocrUrl}/`, {
      headers: { 'Authorization': 'Basic ' + btoa('admin:admin123') }
    })
    const html = await res.text()
    
    // Simple way to get data since the dashboard returns HTML with data embedded
    // We'll fetch from our API endpoint instead for history
    const apiRes = await fetch(`${ocrUrl}/ocr/stats`, {
      headers: { 'Authorization': 'Basic ' + btoa('admin:admin123') }
    }).catch(() => null)
    
    if (apiRes && apiRes.ok) {
       const data = await apiRes.json()
       history.value = data.history || []
       statsData.value = data.stats || { total: 0, avg_time: 0, success_rate: 100 }
    } else {
       // Fallback: If stats endpoint not ready, just reload
       window.location.reload()
    }
  } catch (err) {
    console.error('Failed to fetch history:', err)
  }
}

const addLog = (msg) => {
  logs.value.push(msg)
  nextTick(() => {
    if (logContainer.value) {
      logContainer.value.scrollTop = logContainer.value.scrollHeight
    }
  })
}

const setupWebSocket = () => {
  const wsUrl = ocrUrl.replace('http', 'ws') + '/ws/logs'
  ws.value = new WebSocket(wsUrl)
  
  ws.value.onmessage = (event) => {
    addLog(event.data)
  }
  
  ws.value.onclose = () => {
    setTimeout(setupWebSocket, 5000) // Reconnect
  }
}

const openVisualizer = async (id, preloadedData = null) => {
  if (preloadedData) {
    activeJob.value = preloadedData
    words.value = preloadedData.words || []
  } else {
    try {
      const res = await fetch(`${ocrUrl}/ocr/result/${id}`, {
        headers: { 'Authorization': 'Basic ' + btoa('admin:admin123') }
      })
      if (!res.ok) throw new Error('Result not found')
      const data = await res.json()
      activeJob.value = data
      words.value = JSON.parse(data.words_json)
    } catch (err) {
      alert('Could not load visualizer data')
      return
    }
  }
  showVisualizer.value = true
  highlightAll.value = true
}

const closeVisualizer = () => {
  showVisualizer.value = false
  activeJob.value = null
  words.value = []
}

const parseAI = (data) => {
  if (!data) return {}
  if (typeof data === 'string') {
    try { return JSON.parse(data) } catch { return { summary: data } }
  }
  return data
}

const getBoxStyle = (box) => ({
  left: `${box.x}px`,
  top: `${box.y}px`,
  width: `${box.w}px`,
  height: `${box.h}px`,
  display: highlightAll.value || hoveredIndex.value === words.value.indexOf(box) ? 'block' : 'none'
})

const onHoverBox = (index) => {
  hoveredIndex.value = index
  const textItem = document.getElementById(`text-item-${index}`)
  if (textItem) {
    textItem.scrollIntoView({ behavior: 'smooth', block: 'center' })
  }
}

const onLeaveBox = () => {
  hoveredIndex.value = null
}

const onHoverText = (index) => {
  hoveredIndex.value = index
}

const onLeaveText = () => {
  hoveredIndex.value = null
}

const onBoxClick = (index) => {
  // Toggle individual highlight?
}

const toggleHighlightAll = () => {
  highlightAll.value = !highlightAll.value
}

onMounted(() => {
  fetchHistory()
  setupWebSocket()
})

onUnmounted(() => {
  if (ws.value) ws.value.close()
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 5px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(148, 163, 184, 0.2);
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(148, 163, 184, 0.4);
}

.glass {
  background: rgba(15, 23, 42, 0.8);
  backdrop-filter: blur(20px);
}

.scale-enter-active, .scale-leave-active {
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.scale-enter-from, .scale-leave-to {
  transform: scale(0.9);
  opacity: 0;
}
</style>
