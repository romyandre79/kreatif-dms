<template>
  <div class="space-y-8 animate-in fade-in slide-in-from-bottom-4 duration-700">
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 bg-white dark:bg-[#0D121F] p-8 rounded-lg border border-slate-200/60 dark:border-slate-800/40 shadow-sm relative overflow-hidden">
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
        <!-- Department Selector -->
        <div class="flex flex-col gap-1.5 min-w-[240px]">
          <span class="text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-2 flex items-center gap-2">
            <LucideBuilding2 class="w-3 h-3" />
            Target Department
          </span>
          <select v-model="selectedDept" 
                  class="bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-800/60 rounded-xl px-4 py-3.5 text-xs font-black text-[#1E3A5F] dark:text-slate-300 outline-none focus:ring-2 focus:ring-primary-500/20 transition-all cursor-pointer hover:border-primary-500/30">
            <option value="" disabled>Select Department</option>
            <option v-for="dept in departments" :key="dept.id" :value="dept.id">
              {{ dept.name }}
            </option>
          </select>
        </div>

        <button @click="triggerUpload" :disabled="isUploading"
                class="flex items-center gap-3 px-8 py-4 bg-primary-500 text-white rounded-2xl font-black text-xs uppercase tracking-widest hover:bg-primary-600 hover:shadow-xl hover:shadow-primary-500/20 active:scale-95 transition-all disabled:opacity-50 mt-5">
          <LucideUploadCloud v-if="!isUploading" class="w-4 h-4" />
          <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
          {{ isUploading ? 'Extracting...' : 'Upload Test' }}
        </button>
        <input type="file" ref="fileInput" @change="handleUpload" class="hidden" accept=".pdf,.jpg,.png,.jpeg">
      </div>
    </div>

    <!-- Stats Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div v-for="stat in stats" :key="stat.label" 
           class="bg-white dark:bg-[#0D121F] p-6 rounded-lg border border-slate-200/60 dark:border-slate-800/40 shadow-sm group hover:border-primary-500/30 transition-all duration-500">
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
      <div class="lg:col-span-8 bg-white dark:bg-[#0D121F] rounded-lg border border-slate-200/60 dark:border-slate-800/40 shadow-sm flex flex-col overflow-hidden">
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
                <th class="py-5 px-6">Owner</th>
                <th class="py-5 px-6">Date</th>
                <th class="py-5 px-6">Size</th>
                <th class="py-5 px-6 text-center">Dur</th>
                <th class="py-5 px-6 text-center">Acc</th>
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
                <td class="py-5 px-6 text-slate-500 text-[10px] font-bold uppercase">{{ item.owner }}</td>
                <td class="py-5 px-6 text-slate-500 text-[10px]">{{ new Date(item.created_at).toLocaleDateString() }}</td>
                <td class="py-5 px-6 text-slate-500 text-xs">{{ item.size }}</td>
                <td class="py-5 px-6 text-center font-black text-xs text-primary-500">{{ item.duration }}s</td>
                <td class="py-5 px-6 text-center">
                  <span :class="['font-black text-xs', 
                             item.accuracy >= 0.8 ? 'text-green-500' : (item.accuracy >= 0.5 ? 'text-orange-500' : 'text-red-500')]">
                    {{ (item.accuracy * 100).toFixed(0) }}%
                  </span>
                </td>
                <td class="py-5 px-6">
                  <span :class="['px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest', 
                             (item.status && item.status.toLowerCase() === 'success') ? 'bg-green-500/10 text-green-500' : 
                             (item.status && item.status.toLowerCase() === 'processing') ? 'bg-sky-500/10 text-sky-500 animate-pulse' : 'bg-red-500/10 text-red-500']">
                    {{ item.status }}
                  </span>
                </td>
                <td class="py-5 px-8 text-right">
                  <button v-if="item.status && item.status.toLowerCase() === 'success'" @click="openVisualizer(item.id)"
                          class="p-2 bg-primary-500/10 text-primary-500 hover:bg-primary-500 hover:text-white rounded-xl transition-all">
                    <LucideEye class="w-4 h-4" />
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        
        <!-- Pagination Controls -->
        <div v-if="pagination.total_pages > 1" class="px-8 py-5 border-t border-slate-100 dark:border-slate-800 bg-slate-50/30 dark:bg-slate-900/20 flex items-center justify-between">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            Showing {{ history.length }} of {{ pagination.total_items }} results
          </p>
          <div class="flex items-center gap-2">
            <button @click="changePage(currentPage - 1)" :disabled="currentPage === 1" 
                    class="p-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg disabled:opacity-50 transition-all hover:border-primary-500">
              <LucideChevronLeft class="w-4 h-4 text-slate-400" />
            </button>
            <div class="flex items-center gap-1">
              <button v-for="p in pagination.total_pages" :key="p" @click="changePage(p)"
                      :class="['w-8 h-8 rounded-lg text-[10px] font-black transition-all', 
                              currentPage === p ? 'bg-primary-500 text-white shadow-lg shadow-primary-500/20' : 'bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-400 hover:border-primary-500']">
                {{ p }}
              </button>
            </div>
            <button @click="changePage(currentPage + 1)" :disabled="currentPage === pagination.total_pages"
                    class="p-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg disabled:opacity-50 transition-all hover:border-primary-500">
              <LucideChevronRight class="w-4 h-4 text-slate-400" />
            </button>
          </div>
        </div>
      </div>

      <!-- Live Console -->
      <div class="lg:col-span-4 bg-[#0A0F1C] rounded-lg border border-slate-800/50 shadow-2xl flex flex-col overflow-hidden h-[600px] ring-1 ring-white/5">
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
        
        <div class="relative w-full h-full glass rounded-lg border border-white/10 shadow-2xl flex flex-col overflow-hidden">
            <!-- Modal Header -->
            <div class="px-8 py-6 border-b border-white/5 bg-white/5 flex items-center justify-between shrink-0">
              <div class="flex items-center gap-6">
                <div>
                  <h2 class="text-xl font-black text-white uppercase tracking-tight leading-none mb-1">{{ activeJob?.filename }}</h2>
                  <p class="text-slate-500 text-[10px] font-bold uppercase tracking-widest">
                    {{ activeJob?.timestamp }} | <span class="text-green-500">{{ activeJob?.status }}</span>
                  </p>
                </div>
                <div v-if="activeJob?.ai_analysis && activeJob.ai_analysis !== 'null'" class="h-8 w-px bg-white/10 hidden md:block"></div>
                <div v-if="activeJob?.ai_analysis && activeJob.ai_analysis !== 'null'" class="hidden md:block">
                  <span class="px-3 py-1.5 bg-primary-500 text-white text-[9px] font-black rounded-lg uppercase tracking-widest shadow-lg shadow-primary-500/20">
                    {{ parseAI(activeJob?.ai_analysis).doc_type || 'Document' }}
                  </span>
                </div>
              </div>

              <!-- Page Controls -->
              <div v-if="activeJob?.preview_paths?.length > 1" class="flex items-center gap-4 bg-white/5 rounded-xl p-1 border border-white/10 ml-4">
                <button @click="prevPage" :disabled="currentPreviewIndex === 0" 
                        class="p-2 hover:bg-white/10 rounded-lg text-slate-400 hover:text-white disabled:opacity-30 transition-all">
                  <LucideChevronLeft class="w-4 h-4" />
                </button>
                <span class="text-[10px] font-black text-white uppercase tracking-widest px-2">
                  Page {{ currentPreviewIndex + 1 }} / {{ activeJob.preview_paths.length }}
                </span>
                <button @click="nextPage" :disabled="currentPreviewIndex === activeJob.preview_paths.length - 1"
                        class="p-2 hover:bg-white/10 rounded-lg text-slate-400 hover:text-white disabled:opacity-30 transition-all">
                  <LucideChevronRight class="w-4 h-4" />
                </button>
              </div>

              <div class="flex items-center gap-4 ml-auto">
                <!-- Zoom Controls -->
                <div class="flex items-center bg-white/5 rounded-xl p-1 border border-white/10 gap-1">
                  <button @click="zoomOut" class="p-2 hover:bg-white/10 rounded-lg text-slate-400 hover:text-white transition-all"><LucideZoomOut class="w-4 h-4" /></button>
                  <button @click="resetZoom" class="px-3 text-[10px] font-black text-slate-400 hover:text-white uppercase tracking-widest">{{ (scale * 100).toFixed(0) }}%</button>
                  <button @click="zoomIn" class="p-2 hover:bg-white/10 rounded-lg text-slate-400 hover:text-white transition-all"><LucideZoomIn class="w-4 h-4" /></button>
                </div>

                <button @click="closeVisualizer" class="w-10 h-10 bg-white/5 hover:bg-red-500/20 text-white hover:text-red-500 rounded-xl flex items-center justify-center transition-all">
                  <LucideX class="w-5 h-5" />
                </button>
              </div>
            </div>
          
          <!-- Modal Content -->
          <div class="flex-1 grid grid-cols-1 lg:grid-cols-12 gap-0 overflow-hidden">
            <!-- Left: Image View (Taking 8 columns) -->
            <div class="lg:col-span-8 bg-[#010409] relative overflow-hidden flex items-center justify-center group border-r border-white/5 select-none"
                 @mousedown="startDrag" 
                 @mousemove="onDrag" 
                 @mouseup="stopDrag" 
                 @mouseleave="stopDrag">
              
              <div class="relative transition-transform duration-200 ease-out" 
                   :style="{ 
                     transform: `scale(${scale}) translate(${translateX}px, ${translateY}px)`,
                     cursor: isDragging ? 'grabbing' : 'grab'
                   }">
                <img v-if="activeJob && currentPreviewBlobUrl" 
                     :src="currentPreviewBlobUrl" 
                     ref="visualizerImage"
                     @load="onImageLoad"
                     class="max-w-none shadow-[0_0_100px_rgba(0,0,0,0.5)] rounded-sm ring-1 ring-white/10 pointer-events-none" 
                     alt="Document">
                <div class="absolute inset-0 z-10 pointer-events-none">
                  <div v-for="(word, i) in filteredWords" :key="i"
                       :style="getBoxStyle(word.box, i)"
                       @mouseenter="onHoverBox(i)"
                       @mouseleave="onLeaveBox(i)"
                       @click="onBoxClick(i)"
                       :class="['absolute border-2 border-primary-500/40 bg-primary-500/10 hover:bg-primary-500/40 hover:border-primary-400 cursor-pointer transition-all pointer-events-auto',
                                { 'ring-4 ring-white z-20 scale-[1.05] shadow-2xl': hoveredIndex === i }]"
                       :data-index="i">
                  </div>
                </div>
              </div>

              <!-- Floating Zoom Hint -->
              <div class="absolute bottom-8 left-1/2 -translate-x-1/2 px-4 py-2 bg-black/60 backdrop-blur-md rounded-full border border-white/10 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] opacity-0 group-hover:opacity-100 transition-opacity">
                Click & Drag to move • Use buttons to zoom
              </div>
            </div>
            
            <!-- Right Sidebar: AI Intelligence & OCR List -->
            <div class="lg:col-span-4 bg-[#0d1117] flex flex-col overflow-hidden">
              <div class="flex-1 overflow-y-auto custom-scrollbar">
                
                <!-- AI Insight Sidebar Section -->
                <div v-if="activeJob?.ai_analysis && activeJob.ai_analysis !== 'null'" class="p-6 border-b border-white/5 space-y-6">
                  <!-- Summary -->
                  <div v-if="parseAI(activeJob?.ai_analysis).summary">
                    <div class="flex items-center gap-2 mb-3">
                      <LucideZap class="w-3.5 h-3.5 text-primary-500" />
                      <h4 class="text-[10px] font-black uppercase tracking-widest text-slate-400">AI Summary</h4>
                    </div>
                    <p class="text-[11px] font-bold text-slate-300 italic leading-relaxed bg-primary-500/5 p-4 rounded-xl border border-primary-500/10">
                      "{{ parseAI(activeJob?.ai_analysis).summary }}"
                    </p>
                  </div>

                  <!-- Entities -->
                  <div v-if="parseAI(activeJob?.ai_analysis).entities">
                    <div class="flex items-center gap-2 mb-3">
                      <LucideLayers class="w-3.5 h-3.5 text-sky-500" />
                      <h4 class="text-[10px] font-black uppercase tracking-widest text-slate-400">Extracted Entities</h4>
                    </div>
                    <div class="space-y-2">
                      <div v-for="(val, key) in parseAI(activeJob?.ai_analysis).entities" :key="key" 
                           class="group flex items-center justify-between p-3 bg-white/5 rounded-xl border border-white/5 hover:border-primary-500/30 transition-all">
                        <span class="text-[9px] font-black text-slate-500 uppercase tracking-tight">{{ key.replace('_', ' ') }}</span>
                        <span class="text-[11px] font-bold text-primary-400 group-hover:text-primary-300">{{ val }}</span>
                      </div>
                    </div>
                  </div>

                  <!-- Cleaned Text -->
                  <div v-if="parseAI(activeJob?.ai_analysis).cleaned_text">
                    <div class="flex items-center gap-2 mb-3">
                      <LucideCheckCircle2 class="w-3.5 h-3.5 text-green-500" />
                      <h4 class="text-[10px] font-black uppercase tracking-widest text-slate-400">Reconstructed Text</h4>
                    </div>
                    <div class="p-4 bg-black/40 rounded-xl border border-white/5 text-[11px] text-slate-300 leading-relaxed font-medium whitespace-pre-wrap ring-1 ring-inset ring-white/5">
                      {{ parseAI(activeJob?.ai_analysis).cleaned_text }}
                    </div>
                  </div>
                </div>

                <!-- Raw OCR Word List -->
                <div class="p-6 space-y-4">
                  <div class="flex items-center justify-between">
                    <h4 class="text-[10px] font-black uppercase tracking-widest text-slate-400">OCR Fragments</h4>
                    <button @click="toggleHighlightAll" 
                            :class="['px-3 py-1 text-[9px] font-black rounded-lg transition-all', 
                                    highlightAll ? 'bg-primary-500 text-white' : 'bg-primary-500/10 text-primary-500 border border-primary-500/20']">
                      {{ highlightAll ? 'Hide Overlay' : 'Show Overlay' }}
                    </button>
                  </div>
                  <div class="space-y-2">
                    <div v-for="(word, i) in filteredWords" :key="i"
                         :id="`text-item-${i}`"
                         @mouseenter="onHoverText(i)"
                         @mouseleave="onLeaveText(i)"
                         class="group p-3.5 rounded-xl border border-white/5 transition-all cursor-pointer relative overflow-hidden"
                         :class="{ 'bg-primary-500/20 border-primary-500/50 scale-[1.02] shadow-xl z-10': hoveredIndex === i, 'hover:bg-white/5': hoveredIndex !== i }">
                      <div class="flex items-center justify-between mb-1">
                        <span class="text-[8px] font-black text-slate-500 uppercase tracking-tighter">P{{ word.page }} • {{ (word.confidence * 100).toFixed(0) }}%</span>
                        <div v-if="hoveredIndex === i" class="w-1.5 h-1.5 rounded-full bg-primary-500 animate-ping"></div>
                      </div>
                      <p class="text-xs font-medium text-slate-300 group-hover:text-white transition-colors">{{ word.text }}</p>
                    </div>
                    <div v-if="filteredWords.length === 0" class="py-12 flex flex-col items-center justify-center text-slate-700 gap-4 opacity-50">
                       <LucideSearch class="w-12 h-12" />
                       <p class="text-[10px] font-black uppercase tracking-widest">No results for this page</p>
                    </div>
                  </div>
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
  LucideX, LucideMonitor, LucideZap, LucideZoomIn, LucideZoomOut, LucideLayers,
  LucideSearch, LucideChevronLeft, LucideChevronRight, LucideBuilding2
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
const scale = ref(1)
const translateX = ref(0)
const translateY = ref(0)
const isDragging = ref(false)
const lastMouseX = ref(0)
const lastMouseY = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const pagination = ref({ total_items: 0, total_pages: 0 })
const currentPreviewIndex = ref(0)
const departments = ref([])
const selectedDept = ref('')

const currentPreviewBlobUrl = ref('')

const fetchPreviewImage = async (url) => {
  if (!url) return
  try {
    console.log('[OCR] Fetching preview from:', url)
    const res = await fetch(url, {
      headers: {
        'Authorization': 'Basic ' + btoa('admin:admin123')
      }
    })
    if (res.ok) {
      const blob = await res.blob()
      console.log(`[OCR] Preview loaded: ${blob.size} bytes, type: ${blob.type}`)
      
      if (blob.size < 100) {
        const text = await blob.text()
        console.warn('[OCR] Warning: Preview blob is very small, might be an error message:', text)
      }

      if (currentPreviewBlobUrl.value) {
        URL.revokeObjectURL(currentPreviewBlobUrl.value)
      }
      currentPreviewBlobUrl.value = URL.createObjectURL(blob)
    } else {
      console.error(`[OCR] HTTP Error ${res.status}: ${res.statusText}`)
    }
  } catch (err) {
    console.error('[OCR] Failed to fetch preview image:', err)
  }
}

const currentPreviewUrl = computed(() => {
  if (!activeJob.value) return ''
  const paths = activeJob.value.preview_paths || []
  let path = ''
  if (paths.length > 0 && currentPreviewIndex.value < paths.length) {
    path = paths[currentPreviewIndex.value]
  } else {
    path = activeJob.value.preview_path || activeJob.value.file_path
  }
  // Encode URI to handle spaces in filenames
  return path ? `${ocrUrl}${encodeURI(path)}` : ''
})

watch(currentPreviewUrl, (newUrl) => {
  if (newUrl) fetchPreviewImage(newUrl)
}, { immediate: true })

const filteredWords = computed(() => {
  if (activeJob.value?.preview_paths?.length > 1) {
    return words.value.filter(w => w.page === currentPreviewIndex.value + 1)
  }
  return words.value
})

const nextPage = () => {
  if (currentPreviewIndex.value < (activeJob.value?.preview_paths?.length || 0) - 1) {
    currentPreviewIndex.value++
    resetZoom()
  }
}

const prevPage = () => {
  if (currentPreviewIndex.value > 0) {
    currentPreviewIndex.value--
    resetZoom()
  }
}

const stats = computed(() => [
  { label: 'Total Process', value: statsData.value.total, icon: LucideActivity, bg: 'bg-primary-500/10', color: 'text-primary-500' },
  { label: 'Avg Time', value: `${statsData.value.avg_time.toFixed(2)}s`, icon: LucideClock, bg: 'bg-sky-500/10', color: 'text-sky-500' },
  { label: 'Success Rate', value: `${statsData.value.success_rate.toFixed(1)}%`, icon: LucideCheckCircle2, bg: 'bg-green-500/10', color: 'text-green-500' },
  { label: 'Service Status', value: 'ONLINE', icon: LucideMonitor, bg: 'bg-purple-500/10', color: 'text-purple-500' },
])

const zoomIn = () => { scale.value = Math.min(scale.value + 0.25, 5) }
const zoomOut = () => { scale.value = Math.max(scale.value - 0.25, 0.5) }
const resetZoom = () => { scale.value = 1; translateX.value = 0; translateY.value = 0; }

const startDrag = (e) => {
  isDragging.value = true
  lastMouseX.value = e.clientX
  lastMouseY.value = e.clientY
}

const onDrag = (e) => {
  if (!isDragging.value) return
  const dx = e.clientX - lastMouseX.value
  const dy = e.clientY - lastMouseY.value
  translateX.value += dx / scale.value
  translateY.value += dy / scale.value
  lastMouseX.value = e.clientX
  lastMouseY.value = e.clientY
}

const stopDrag = () => { isDragging.value = false }

const onImageLoad = () => { console.log('[OCR] Image loaded into viewer') }

const triggerUpload = () => fileInput.value.click()

const handleUpload = async (e) => {
  const file = e.target.files[0]
  if (!file) return

  if (!selectedDept.value) {
    addLog(`[UI] Error: Please select a department first!`, 'text-red-500')
    alert('Please select a target department first.')
    return
  }

  isUploading.value = true
  addLog(`[UI] Uploading ${file.name} to Target Department...`)

  const formData = new FormData()
  formData.append('file', file)
  formData.append('title', file.name)
  formData.append('department_id', selectedDept.value)

  try {
    const apiBase = config.public.apiBase || 'http://localhost:8080/api/v1'
    const token = localStorage.getItem('kreatif_access_token') || localStorage.getItem('token')
    
    const response = await fetch(`${apiBase}/documents`, {
      method: 'POST',
      body: formData,
      headers: { 
        'Authorization': `Bearer ${token}`,
        'Accept': 'application/json'
      }
    })

    if (!response.ok) {
      const errData = await response.json().catch(() => ({}))
      throw new Error(errData.message || `HTTP Error: ${response.status}`)
    }
    
    const json = await response.json()
    addLog(`[UI] Upload successful. Document ID: ${json.data.id}`, 'text-green-400')
    addLog(`[UI] OCR Task enqueued. Please wait for processing...`, 'text-sky-400')
    
    // Refresh history to show the new "Processing" item
    fetchHistory()
    
  } catch (err) {
    addLog(`[UI] Upload Error: ${err.message}`, 'text-red-500')
    console.error(err)
  } finally {
    isUploading.value = false
    e.target.value = ''
  }
}

const fetchHistory = async () => {
  try {
    // Fetch from Main Go Backend instead of OCR Service
    const apiBase = config.public.apiBase || 'http://localhost:8080/api/v1'
    const token = localStorage.getItem('kreatif_access_token') || localStorage.getItem('token')
    
    const res = await fetch(`${apiBase}/documents/history?page=${currentPage.value}&page_size=${pageSize.value}`, {
      headers: { 
        'Authorization': `Bearer ${token}`,
        'Accept': 'application/json'
      }
    })
    
    if (res.ok) {
       const json = await res.json()
       const data = json.data || {}
       
       // Map combined DB fields to what our table expects
       history.value = (data.history || []).map(row => ({
         id: row.entity_id || row.id, // Prefer document ID for visualizer
         filename: row.filename,
         size: (row.file_size / 1024).toFixed(1) + ' KB',
         duration: (row.processing_time_ms / 1000).toFixed(2),
         accuracy: parseFloat(row.confidence_avg || 0),
         status: row.status === 'Success' || row.status === 'success' || row.status === 'active' ? 'Success' : row.status,
         owner: row.owner_name || 'System',
         created_at: row.created_at
       }))
       
       if (data.pagination) pagination.value = data.pagination
    }
    
    // Also fetch stats from OCR Service to keep the top cards updated
    const statsRes = await fetch(`${ocrUrl}/ocr/stats`, {
      headers: { 'Authorization': 'Basic ' + btoa('admin:admin123') }
    }).catch(() => null)
    if (statsRes && statsRes.ok) {
       const statsDataRaw = await statsRes.json()
       statsData.value = statsDataRaw.stats || statsData.value
    }
  } catch (err) {
    console.error('Failed to fetch history:', err)
  }
}

const fetchDepartments = async () => {
  try {
    const apiBase = config.public.apiBase || 'http://localhost:8080/api/v1'
    const token = localStorage.getItem('kreatif_access_token') || localStorage.getItem('token')
    
    addLog(`[UI] Fetching departments from ${apiBase}...`)
    
    const res = await fetch(`${apiBase}/master/departments`, {
      headers: { 
        'Authorization': `Bearer ${token}`,
        'Accept': 'application/json'
      }
    })
    
    if (res.ok) {
      const json = await res.json()
      departments.value = json.data || []
      
      if (departments.value.length === 0) {
        addLog(`[UI] Warning: No departments found in database.`, 'text-orange-400')
      } else {
        addLog(`[UI] Loaded ${departments.value.length} departments.`, 'text-green-400')
        // Auto-select first department
        if (!selectedDept.value) {
          selectedDept.value = departments.value[0].id
        }
      }
    } else {
      const errText = await res.text()
      addLog(`[UI] Error fetching departments: ${res.status} ${res.statusText}`, 'text-red-500')
      console.error('Dept Fetch Error:', errText)
    }
  } catch (err) {
    addLog(`[UI] Connection Error: ${err.message}`, 'text-red-500')
    console.error('Failed to fetch departments:', err)
  }
}

const changePage = (p) => {
  if (p < 1 || p > pagination.value.total_pages) return
  currentPage.value = p
  fetchHistory()
}

const addLog = (msg) => {
  logs.value.push(msg)
}

const setupWebSocket = () => {
  const wsUrl = ocrUrl.replace('http', 'ws') + '/ws/logs'
  ws.value = new WebSocket(wsUrl)
  ws.value.onmessage = (event) => { addLog(event.data) }
  ws.value.onclose = () => { setTimeout(setupWebSocket, 5000) }
}

const openVisualizer = async (id, preloadedData = null) => {
  resetZoom()
  if (preloadedData) {
    if (preloadedData.insight && !preloadedData.ai_analysis) preloadedData.ai_analysis = preloadedData.insight
    activeJob.value = preloadedData
    words.value = preloadedData.words || (typeof preloadedData.words_json === 'string' ? JSON.parse(preloadedData.words_json) : (preloadedData.words_json || []))
  } else {
    try {
      const apiBase = config.public.apiBase || 'http://localhost:8080/api/v1'
      const token = localStorage.getItem('kreatif_access_token') || localStorage.getItem('token')
      
      const res = await fetch(`${apiBase}/documents/${id}/ocr`, {
        headers: { 
          'Authorization': `Bearer ${token}`,
          'Accept': 'application/json'
        }
      })
      if (!res.ok) throw new Error(`HTTP Error: ${res.status}`)
      const json = await res.json()
      const data = json.data
      
      // Ensure preview_paths is an array
      if (typeof data.preview_paths === 'string') {
        try { data.preview_paths = JSON.parse(data.preview_paths) } catch (e) { data.preview_paths = [] }
      }
      
      activeJob.value = data
      currentPreviewIndex.value = 0
      words.value = typeof data.words_json === 'string' ? JSON.parse(data.words_json) : (data.words_json || [])
    } catch (err) {
      alert('Could not load visualizer data: ' + err.message)
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
  if (currentPreviewBlobUrl.value) {
    URL.revokeObjectURL(currentPreviewBlobUrl.value)
    currentPreviewBlobUrl.value = ''
  }
  resetZoom()
}

const parseAI = (data) => {
  if (!data || data === 'null') return {}
  if (typeof data === 'string') {
    try { return JSON.parse(data) || {} } catch (e) { return { summary: data } }
  }
  return data || {}
}

const getBoxStyle = (box, index) => ({
  left: `${box.x}px`,
  top: `${box.y}px`,
  width: `${box.w}px`,
  height: `${box.h}px`,
  display: highlightAll.value || hoveredIndex.value === index ? 'block' : 'none'
})

const onHoverBox = (index) => { hoveredIndex.value = index }
const onLeaveBox = () => { hoveredIndex.value = null }
const onHoverText = (index) => { hoveredIndex.value = index }
const onLeaveText = () => { hoveredIndex.value = null }

const onBoxClick = (index) => {
  const item = document.getElementById(`text-item-${index}`)
  if (item) item.scrollIntoView({ behavior: 'smooth', block: 'center' })
}

const toggleHighlightAll = () => { highlightAll.value = !highlightAll.value }

onMounted(() => {
  fetchHistory()
  fetchDepartments()
  setupWebSocket()
})

onUnmounted(() => {
  if (ws.value) ws.value.close()
})

watch(logs, () => {
  nextTick(() => {
    if (logContainer.value) logContainer.value.scrollTop = logContainer.value.scrollHeight
  })
}, { deep: true })
</script>

<style>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
  height: 6px;
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
