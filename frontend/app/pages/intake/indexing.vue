<template>
  <div class="h-screen bg-slate-50 flex flex-col overflow-hidden" v-motion-fade>
    <!-- Top Header -->
    <header class="bg-white border-b border-slate-200 px-8 py-4 flex items-center justify-between z-10 shadow-sm">
      <div class="flex items-center gap-4">
        <button @click="confirmExit" class="p-2 hover:bg-slate-100 rounded-xl transition-all">
          <LucideArrowLeft class="w-6 h-6 text-slate-400" />
        </button>
        <div class="space-y-0.5">
          <h1 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">Indexing Dokumen</h1>
          <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Manifest: <span class="text-orange-500 font-black">{{ manifestNo || 'MNF-2026-0312-007' }}</span> • Dokumen {{ currentDocIndex + 1 }} dari {{ totalDocs }}</p>
        </div>
      </div>
      <div class="flex items-center gap-6">
        <div class="flex items-center gap-2 px-4 py-2 bg-emerald-50 text-emerald-600 rounded-full border border-emerald-100">
          <LucideCheckCircle2 class="w-4 h-4" />
          <span class="text-xs font-black uppercase tracking-widest">Kondisi Fisik: BAIK</span>
        </div>
        <div class="w-px h-6 bg-slate-200"></div>
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-none text-right">
          Document Controller:<br/>
          <span class="text-[#1E3A5F]">Romy Andre</span>
        </p>
      </div>
    </header>

    <div class="flex-1 flex overflow-hidden">
      <!-- Left: Document Viewer -->
      <main class="flex-1 relative bg-slate-200 flex flex-col p-6 overflow-hidden">
        <div class="flex-1 glass-viewer rounded-3xl overflow-hidden shadow-2xl relative p-8 border border-white/20">
          <div ref="viewerContainer" @wheel.prevent="handleWheel" @mousedown="startDrag" @mousemove="onDrag" @mouseup="stopDrag" @mouseleave="stopDrag" class="w-full h-full relative overflow-hidden custom-scrollbar bg-white/50 backdrop-blur-sm rounded-2xl cursor-grab" :class="{ 'cursor-grabbing': isDragging }" style="touch-action: none;">
            <div class="min-w-full min-h-full flex items-center justify-center p-4">
              <img 
                v-if="!useIframe"
                :src="previewImageUrl" 
                @error="handleImageError"
                draggable="false"
                class="max-w-none shadow-2xl origin-center select-none"
                :style="{ 
                  transform: `translate(${panX}px, ${panY}px) scale(${zoom / 100}) rotate(${rotation}deg)`,
                  filter: isNegative ? 'invert(1)' : 'none',
                  transition: isDragging ? 'none' : 'transform 0.3s ease'
                }"
              />
              <iframe 
                v-else
                :src="previewPdfUrl" 
                class="w-full h-full min-h-[600px] border-none transition-all duration-300 rounded-2xl"
              ></iframe>
            </div>
          </div>

          <!-- Floating Toolbar -->
          <div class="absolute bottom-8 left-1/2 -translate-x-1/2 flex items-center gap-2 p-2 bg-white/90 backdrop-blur-xl border border-white rounded-2xl shadow-2xl z-20">
            <button @click="zoom = Math.max(10, zoom - 10)" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 transition-all"><LucideZoomOut class="w-5 h-5" /></button>
            <span class="text-xs font-black text-[#1E3A5F] w-12 text-center">{{ zoom }}%</span>
            <button @click="zoom = Math.min(500, zoom + 10)" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 transition-all"><LucideZoomIn class="w-5 h-5" /></button>
            <div class="w-px h-4 bg-slate-200 mx-2"></div>
            <button @click="rotation -= 90" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 transition-all"><LucideRotateCcw class="w-5 h-5" /></button>
            <button @click="rotation += 90" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 transition-all"><LucideRotateCw class="w-5 h-5" /></button>
            <div class="w-px h-4 bg-slate-200 mx-2"></div>
            <button @click="isNegative = !isNegative" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 transition-all"><LucideContrast class="w-5 h-5" /></button>
            <button @click="zoom = 100; rotation = 0; panX = 0; panY = 0" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 transition-all"><LucideMaximize2 class="w-5 h-5" /></button>
            <div class="w-px h-4 bg-slate-200 mx-2"></div>
            <div class="flex items-center gap-3">
              <button @click="prevDoc" :disabled="currentDocIndex <= 0" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 disabled:opacity-30"><LucideChevronLeft class="w-5 h-5" /></button>
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Dok <span class="text-[#1E3A5F]">{{ currentDocIndex + 1 }}</span> / {{ totalDocs }}</p>
              <button @click="nextDoc" :disabled="currentDocIndex >= totalDocs - 1" class="p-2 hover:bg-slate-100 rounded-lg text-slate-500 disabled:opacity-30"><LucideChevronRight class="w-5 h-5" /></button>
            </div>
          </div>
        </div>
      </main>

      <!-- Right: Indexing Form -->
      <aside class="w-[450px] bg-white border-l border-slate-200 flex flex-col shadow-2xl relative z-10 overflow-hidden">
        <div class="flex-1 overflow-y-auto p-10 space-y-12 custom-scrollbar pb-52">
          <!-- Section Title -->
          <div class="space-y-4">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-indigo-50 text-indigo-600 rounded-xl flex items-center justify-center">
                <LucideBraces class="w-5 h-5" />
              </div>
              <div>
                <h3 class="text-lg font-black text-[#1E3A5F] tracking-tight">Document Indexing</h3>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Lengkapi metadata berdasarkan hasil scan</p>
              </div>
            </div>
          </div>

          <!-- Section 01: Classification -->
          <div class="space-y-8">
            <div class="flex items-center gap-3">
              <span class="w-6 h-6 bg-slate-100 text-[#1E3A5F] text-[10px] font-black rounded flex items-center justify-center">01</span>
              <h4 class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">Classification</h4>
            </div>
            <div class="grid grid-cols-2 gap-6">
              <div class="space-y-2">
                <label class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest">Document Category</label>
                <select v-model="form.category" class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 text-xs font-bold focus:ring-2 focus:ring-blue-500/20 transition-all outline-none">
                  <option v-for="cat in documentCategories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
                </select>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest">Document Type</label>
                <select v-model="form.type" class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 text-xs font-bold focus:ring-2 focus:ring-blue-500/20 transition-all outline-none">
                  <option v-for="t in filteredTypes" :key="t.id" :value="t.id">{{ t.name }}</option>
                </select>
              </div>
            </div>
          </div>

          <!-- Section 02: Dynamic Metadata -->
          <div class="space-y-8">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <span class="w-6 h-6 bg-slate-100 text-[#1E3A5F] text-[10px] font-black rounded flex items-center justify-center">02</span>
                <h4 class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">Metadata Fields</h4>
              </div>
              <button @click="addMetadataField" class="text-[10px] font-black text-primary-600 uppercase tracking-widest hover:text-primary-700 transition-colors flex items-center gap-2">
                <LucidePlus class="w-3.5 h-3.5" />
                Add Field
              </button>
            </div>
            <div class="grid grid-cols-1 gap-6">
              <div v-for="(field, index) in dynamicFields" :key="index" class="space-y-2 group">
                <div class="flex items-center justify-between">
                  <label class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest">{{ formatLabel(field.key) }}</label>
                  <button @click="removeMetadataField(index)" class="opacity-0 group-hover:opacity-100 text-red-400 hover:text-red-600 transition-all">
                    <LucideX class="w-3.5 h-3.5" />
                  </button>
                </div>
                <div class="relative">
                  <input 
                    v-model="field.value"
                    type="text" 
                    class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3.5 text-xs font-bold text-[#1E3A5F] focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all outline-none shadow-sm"
                    :placeholder="'Enter ' + formatLabel(field.key)"
                  />
                  <div class="absolute right-4 top-1/2 -translate-y-1/2 flex items-center gap-2">
                    <LucideCheckCircle2 v-if="field.value" class="w-3.5 h-3.5 text-emerald-500" />
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Physical Condition -->
          <div class="p-6 bg-slate-50 rounded-[24px] flex items-center justify-between border border-slate-100 shadow-inner">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-white rounded-xl flex items-center justify-center text-slate-400 shadow-sm">
                <LucideArchive class="w-5 h-5" />
              </div>
              <p class="text-xs font-black text-[#1E3A5F] uppercase tracking-widest">Physical Condition</p>
            </div>
            <span class="px-4 py-1.5 bg-emerald-500 text-white text-[10px] font-black rounded-lg uppercase tracking-widest shadow-lg shadow-emerald-500/20">BAIK</span>
          </div>

          <!-- OCR Text Preview -->
          <div v-if="currentItem?.extracted_text" class="space-y-4 pt-8 border-t border-slate-100">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <div class="w-8 h-8 bg-orange-50 text-orange-600 rounded-lg flex items-center justify-center">
                  <LucideInfo class="w-4 h-4" />
                </div>
                <h4 class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest">Raw OCR Text</h4>
              </div>
              <button @click="copyOcrText" class="text-[10px] font-black text-indigo-600 hover:text-indigo-700 uppercase tracking-widest flex items-center gap-1 transition-colors">
                Copy
              </button>
            </div>
            <div class="p-6 bg-slate-50 rounded-[24px] border border-slate-100 max-h-[300px] overflow-y-auto custom-scrollbar">
              <pre class="text-[11px] text-[#1E3A5F] font-bold whitespace-pre-wrap leading-relaxed opacity-60">{{ currentItem.extracted_text }}</pre>
            </div>
          </div>
        </div>

        <!-- Sticky Footer -->
        <div class="absolute bottom-0 left-0 right-0 p-8 bg-white border-t border-slate-100 flex flex-col gap-4 shadow-[0_-20px_50px_rgba(0,0,0,0.05)]">
          <button @click="saveAndNext" :disabled="isSaving" class="w-full py-4 bg-[#1E3A5F] hover:bg-[#2A4B7C] disabled:bg-slate-300 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3">
            <LucideLoader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
            <LucideCheckCircle2 v-else class="w-4 h-4" />
            {{ isSaving ? 'Menyimpan...' : 'Simpan & Lanjut' }}
          </button>
          <div class="flex gap-4">
            <button @click="skipDoc" class="flex-1 py-4 border border-red-500/10 hover:bg-red-50 text-red-500 rounded-2xl text-[10px] font-black uppercase tracking-widest transition-all">
              Batal / Skip
            </button>
            <button @click="finishBatch" class="flex-1 py-4 border border-slate-200 hover:bg-slate-50 text-slate-400 rounded-2xl text-[10px] font-black uppercase tracking-widest transition-all flex items-center justify-center gap-2">
              <LucideListChecks class="w-4 h-4" />
              Selesai Batch
            </button>
          </div>
        </div>
      </aside>
    </div>

    <!-- Success Modal -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100" leave-active-class="transition duration-200 ease-in" leave-from-class="opacity-100" leave-to-class="opacity-0">
      <div v-if="showSuccessModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/60 backdrop-blur-md">
        <div class="bg-white w-full max-w-sm rounded-3xl shadow-2xl p-10 flex flex-col items-center text-center" v-motion-pop>
          <!-- Success Icon -->
          <div class="w-16 h-16 rounded-full bg-emerald-500 flex items-center justify-center mb-6 shadow-xl shadow-emerald-500/30">
            <LucideCheckCircle2 class="w-8 h-8 text-white" />
          </div>

          <!-- === View 1: Per-Document Success === -->
          <template v-if="isBatchComplete === false">
            <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight mb-3">Dokumen Berhasil Disimpan</h3>
            <p class="text-xs text-slate-500 font-medium leading-relaxed mb-8">
              Data manifest <span class="font-black text-[#1E3A5F]">{{ manifestNo }}</span> 
              dokumen ke-<span class="font-black text-[#1E3A5F]">{{ savedDocIndex + 1 }}</span> telah 
              berhasil diindeks dan masuk ke <span class="font-black text-[#1E3A5F]">Arsip Digital</span>.
            </p>
            <div class="w-full space-y-3">
              <button 
                @click="continueToNextDoc" 
                class="w-full py-4 bg-[#1E3A5F] hover:bg-[#2A4B7C] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3"
              >
                Lanjut Dokumen Berikutnya ({{ currentDocIndex + 1 }}/{{ totalDocs }})
                <LucideArrowRight class="w-4 h-4" />
              </button>
              <button 
                @click="showSuccessModal = false" 
                class="w-full py-4 border border-slate-200 hover:bg-slate-50 text-[#1E3A5F] rounded-2xl text-xs font-black uppercase tracking-widest transition-all flex items-center justify-center gap-3"
              >
                <LucideEye class="w-4 h-4" />
                Lihat Hasil Index
              </button>
              <button 
                @click="showSuccessModal = false" 
                class="w-full py-3 text-slate-400 hover:text-slate-600 text-[10px] font-black uppercase tracking-widest transition-colors"
              >
                Tutup Panel
              </button>
            </div>
          </template>

          <!-- === View 2: Batch Complete === -->
          <template v-else>
            <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight mb-3">Batch Digitalisasi Selesai</h3>
            <p class="text-xs text-slate-500 font-medium leading-relaxed mb-8">
              Semua dokumen dalam Manifest <span class="font-black text-[#1E3A5F]">{{ manifestNo }}</span> 
              telah berhasil didigitalisasi dan disimpan dengan aman.
            </p>

            <!-- Stats -->
            <div class="w-full border border-slate-100 rounded-2xl divide-y divide-slate-100 mb-8">
              <div class="flex items-center justify-between px-6 py-4">
                <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Total Dokumen</span>
                <span class="text-sm font-black text-[#1E3A5F]">{{ totalDocs }} Documents</span>
              </div>
              <div class="flex items-center justify-between px-6 py-4">
                <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Berhasil</span>
                <span class="text-sm font-black text-emerald-500 flex items-center gap-2">
                  <LucideCheckCircle2 class="w-3.5 h-3.5" />
                  {{ totalDocs }} Dokumen
                </span>
              </div>
              <div class="flex items-center justify-between px-6 py-4">
                <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Waktu Proses</span>
                <span class="text-sm font-black text-[#1E3A5F]">{{ processingTime }}</span>
              </div>
            </div>

            <div class="w-full space-y-3">
              <button 
                @click="finishBatch" 
                class="w-full py-4 bg-[#1E3A5F] hover:bg-[#2A4B7C] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3"
              >
                <LucideArrowLeft class="w-4 h-4" />
                Kembali ke Staging Area
              </button>
              <button 
                @click="navigateTo('/documents')" 
                class="w-full py-4 border border-slate-200 hover:bg-slate-50 text-[#1E3A5F] rounded-2xl text-xs font-black uppercase tracking-widest transition-all flex items-center justify-center gap-3"
              >
                <LucideArchive class="w-4 h-4" />
                Lihat Arsip Digital
              </button>
            </div>

            <p class="mt-6 text-[9px] font-bold text-slate-300 uppercase tracking-widest flex items-center gap-2">
              <LucideCheckCircle2 class="w-3 h-3" />
              Otomatis diklasifikasikan oleh OCR
            </p>
          </template>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { 
  LucideArrowLeft, LucideZoomIn, LucideZoomOut, LucideRotateCw, LucideRotateCcw,
  LucideContrast, LucideMaximize2, LucideChevronLeft, LucideChevronRight,
  LucideBraces, LucideCalendar, LucideArchive, LucideCheckCircle2, 
  LucideListChecks, LucideInfo, LucideLoader2, LucidePlus, LucideX,
  LucideArrowRight, LucideEye
} from 'lucide-vue-next'

const route = useRoute()
const { $api } = useApi()
const documentCategories = ref([])
const documentTypes = ref([])

const manifestNo = ref(route.query.manifest_no)
const manifestData = ref(null)
const currentDocIndex = ref(0)
const zoom = ref(100)
const rotation = ref(0)
const isNegative = ref(false)
const isSaving = ref(false)
const useIframe = ref(false)

const form = ref({
  category: '',
  type: '',
})

const dynamicFields = ref([])

const addMetadataField = () => {
  const key = prompt('Enter field name (e.g. Invoice Number):')
  if (key) {
    dynamicFields.value.push({ key: key.toLowerCase().replace(/ /g, '_'), value: '' })
  }
}

const removeMetadataField = (index) => {
  dynamicFields.value.splice(index, 1)
}

const formatLabel = (key) => {
  return key.split('_').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ')
}

// Filtered types based on selected category
const filteredTypes = computed(() => {
  if (!form.value.category) return documentTypes.value
  return documentTypes.value.filter(t => t.category_name === form.value.category || t.category_id === form.value.category)
})

const currentItem = computed(() => {
  if (!manifestData.value || !manifestData.value.items) return null
  return manifestData.value.items[currentDocIndex.value]
})

const totalDocs = computed(() => manifestData.value?.total_items || 0)

const copyOcrText = () => {
  if (currentItem.value?.extracted_text) {
    navigator.clipboard.writeText(currentItem.value.extracted_text)
    alert('OCR Text copied to clipboard!')
  }
}

onMounted(async () => {
  // 1. Fetch Master Data FIRST (so dropdowns are populated before selecting values)
  try {
    const [catsRes, typesRes] = await Promise.all([
      $api('/master/document-categories'),
      $api('/master/document-types')
    ])
    documentCategories.value = catsRes?.data || []
    documentTypes.value = typesRes?.data || []
    console.log('Categories loaded:', documentCategories.value.length, documentCategories.value)
    console.log('Types loaded:', documentTypes.value.length, documentTypes.value)
  } catch (e) {
    console.error('Failed to fetch master data', e)
  }

  // 2. Fetch Manifest Data AFTER master data is loaded
  if (manifestNo.value) {
    try {
      const res = await $api(`/intake/manifest/${manifestNo.value}`)
      manifestData.value = res.data
      
      // Initialize form with first item data
      if (manifestData.value.items && manifestData.value.items.length > 0) {
        loadItemToForm(manifestData.value.items[0])
      }
    } catch (e) {
      console.error('Failed to fetch manifest', e)
    }
  }
})

const previewImageUrl = ref(null)
const previewPdfUrl = ref(null)

const handleImageError = () => {
  console.log('Image failed to load, switching to iframe fallback')
  useIframe.value = true
}

const viewerContainer = ref(null)
const isDragging = ref(false)
const panX = ref(0)
const panY = ref(0)
const dragStart = ref({ x: 0, y: 0 })
const panStart = ref({ x: 0, y: 0 })

const handleWheel = (e) => {
  e.preventDefault()
  const delta = e.deltaY < 0 ? 10 : -10
  zoom.value = Math.max(10, Math.min(500, zoom.value + delta))
  
  // Reset pan when zooming back to 100% or below
  if (zoom.value <= 100) {
    panX.value = 0
    panY.value = 0
  }
}

const startDrag = (e) => {
  if (e.button !== 0) return // left click only
  isDragging.value = true
  dragStart.value = { x: e.clientX, y: e.clientY }
  panStart.value = { x: panX.value, y: panY.value }
}

const onDrag = (e) => {
  if (!isDragging.value) return
  panX.value = panStart.value.x + (e.clientX - dragStart.value.x)
  panY.value = panStart.value.y + (e.clientY - dragStart.value.y)
}

const stopDrag = () => {
  isDragging.value = false
}

const loadItemToForm = (item) => {
  if (!item) return
  
  console.log('Loading item to form:', JSON.stringify(item, null, 2))
  
  // Set type_id directly
  form.value.type = item.type_id || ''
  
  // Resolve category: either from item directly, or lookup from the type's category
  if (item.category_id) {
    form.value.category = item.category_id
  } else if (item.type_id && documentTypes.value.length > 0) {
    // Lookup category from the document type
    const docType = documentTypes.value.find(t => t.id === item.type_id)
    form.value.category = docType?.category_id || ''
  } else {
    form.value.category = ''
  }
  
  // Load all metadata keys into dynamic fields
  const meta = item.metadata || item.document_metadata || {}
  dynamicFields.value = Object.entries(meta).map(([key, value]) => ({
    key: key,
    value: typeof value === 'string' ? value : String(value || '')
  }))

  // Ensure basic fields are present if empty
  const basicKeys = ['document_number', 'document_date']
  basicKeys.forEach(k => {
    if (!dynamicFields.value.find(f => f.key === k)) {
      dynamicFields.value.push({ key: k, value: '' })
    }
  })

  // Resolve Preview URL
  if (item.document_id) {
    const config = useRuntimeConfig()
    const auth = useAuthStore()
    const apiBase = config.public.apiBase || 'http://localhost:8080/api/v1'
    previewImageUrl.value = `${apiBase}/documents/${item.document_id}/image?token=${auth.accessToken}`
    previewPdfUrl.value = `${apiBase}/documents/${item.document_id}/preview?token=${auth.accessToken}`
    useIframe.value = false
    // Reset viewer state for new document
    zoom.value = 100
    panX.value = 0
    panY.value = 0
    rotation.value = 0
  } else {
    previewImageUrl.value = null
    previewPdfUrl.value = null
  }
}

const showSuccessModal = ref(false)
const savedDocIndex = ref(0)
const isBatchComplete = ref(false)
const batchStartTime = ref(Date.now())
const processingTime = ref('')

const hasMoreDocs = computed(() => {
  if (!manifestData.value?.items) return false
  return currentDocIndex.value < manifestData.value.items.length - 1
})

const calcProcessingTime = () => {
  const elapsed = Math.floor((Date.now() - batchStartTime.value) / 1000)
  const mins = Math.floor(elapsed / 60)
  const secs = elapsed % 60
  processingTime.value = mins > 0 ? `${mins}m ${secs}s` : `${secs}s`
}

const saveAndNext = async () => {
  if (!currentItem.value) return
  
  isSaving.value = true
  try {
    // Construct metadata from dynamic fields
    const metadata = {}
    dynamicFields.value.forEach(f => {
      if (f.key) metadata[f.key] = f.value
    })

    await $api(`/intake/indexing/${currentItem.value.document_id}`, {
      method: 'PUT',
      body: {
        manifest_id: manifestData.value.id,
        category_id: form.value.category,
        type_id: form.value.type,
        metadata: metadata
      }
    })

    // Save which doc was just saved
    savedDocIndex.value = currentDocIndex.value

    // Check if this was the last doc
    if (hasMoreDocs.value) {
      isBatchComplete.value = false
      currentDocIndex.value++
      loadItemToForm(manifestData.value.items[currentDocIndex.value])
    } else {
      isBatchComplete.value = true
      calcProcessingTime()
    }

    // Show success modal
    showSuccessModal.value = true
  } catch (e) {
    alert('Gagal menyimpan indexing: ' + (e.response?._data?.message || e.message))
  } finally {
    isSaving.value = false
  }
}

const continueToNextDoc = () => {
  showSuccessModal.value = false
  // Next doc is already loaded by saveAndNext
  // If we're at the last doc, go back to staging
  if (!hasMoreDocs.value && currentDocIndex.value >= manifestData.value.items.length - 1) {
    // Already on last doc, just close modal
  }
}

const confirmExit = () => {
  if (confirm('Anda sedang dalam proses indexing. Yakin ingin keluar?')) {
    navigateTo('/intake/staging')
  }
}

const skipDoc = () => {
  if (confirm('Lewati dokumen ini?')) {
    if (currentDocIndex.value < manifestData.value.items.length - 1) {
      currentDocIndex.value++
      loadItemToForm(manifestData.value.items[currentDocIndex.value])
    }
  }
}

const finishBatch = () => {
  navigateTo('/intake/staging')
}

const prevDoc = () => {
  if (currentDocIndex.value > 0) {
    currentDocIndex.value--
    loadItemToForm(manifestData.value.items[currentDocIndex.value])
  }
}

const nextDoc = () => {
  if (currentDocIndex.value < manifestData.value.items.length - 1) {
    currentDocIndex.value++
    loadItemToForm(manifestData.value.items[currentDocIndex.value])
  }
}
</script>

<style scoped>
.glass-viewer {
  background: rgba(255, 255, 255, 0.4);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.6);
}

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #e2e8f0;
  border-radius: 10px;
}

select {
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='%23cbd5e1'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M19 9l-7 7-7-7'%3E%3C/path%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 1rem center;
  background-size: 1rem;
}
</style>
