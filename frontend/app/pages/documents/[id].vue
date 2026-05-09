<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20">
    <!-- Top Navigation & Header -->
    <div class="flex items-center justify-between" v-motion-fade>
      <div class="space-y-4">
        <button 
          @click="navigateTo('/documents')" 
          class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-primary-600 transition-colors uppercase tracking-widest group"
        >
          <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
          {{ $t('documents.detail.btn_back') }}
        </button>
        <div class="space-y-1">
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ doc?.title || $t('documents.detail.loading') }}</h1>
          <p v-if="doc?.file_name" class="text-sm font-bold text-slate-400 uppercase tracking-widest">{{ doc?.file_name }}</p>
        </div>
      </div>
    </div>

    <!-- Error State -->
    <div v-if="docRes?.error" class="mb-10 p-6 bg-red-50 border border-red-100 rounded-3xl flex items-center gap-4 text-red-600" v-motion-fade>
      <div class="p-3 bg-red-100 rounded-2xl">
        <LucideAlertCircle class="w-6 h-6" />
      </div>
      <div>
        <h3 class="font-black uppercase tracking-tight">Gagal Memuat Dokumen</h3>
        <p class="text-xs font-bold opacity-70">{{ docRes.error.message || 'Cek koneksi API' }}</p>
      </div>
    </div>

    <!-- Standard PDF Viewer Wrapper -->
    <div 
      class="glass rounded-lg overflow-hidden shadow-2xl shadow-slate-200/50 dark:shadow-none border border-white transition-all duration-500"
      :class="isEnlarged ? 'fixed inset-0 z-[1000] h-screen w-screen bg-[#525659]' : ''"
      v-motion-fade
    >
      <!-- Viewer Header/Toolbar -->
      <div class="bg-[#525659] p-2 flex items-center justify-between text-white border-b border-white/10">
        <div class="flex items-center gap-4">
          <div class="flex items-center gap-2 px-2">
            <span class="text-xs font-bold opacity-80 uppercase tracking-widest">{{ $t('documents.detail.viewer.page') }}</span>
            <input 
              type="text" 
              v-model="currentPage" 
              class="w-10 bg-[#323639] border-none rounded px-2 py-1 text-xs font-bold text-center focus:ring-1 focus:ring-blue-400" 
            />
            <span class="text-xs font-bold opacity-80">/ {{ doc?.metadata?.page_count || 1 }}</span>
          </div>
        </div>
        <div class="flex items-center gap-6">
          <div class="flex items-center gap-4 bg-[#323639] px-4 py-1.5 rounded-full border border-white/5 shadow-inner">
            <button @click="zoomOut" class="hover:text-blue-400 transition-colors"><LucideMinus class="w-4 h-4" /></button>
            <span class="text-xs font-bold w-12 text-center tracking-tighter">{{ zoomLevel }}%</span>
            <button @click="zoomIn" class="hover:text-blue-400 transition-colors"><LucidePlus class="w-4 h-4" /></button>
          </div>
          <div class="flex items-center gap-4 border-l border-white/10 pl-6">
            <button @click="printDocument" class="hover:text-blue-400 transition-colors"><LucidePrinter class="w-5 h-5" /></button>
            <button @click="downloadDocument" class="hover:text-blue-400 transition-colors"><LucideDownload class="w-5 h-5" /></button>
            <button @click="toggleEnlarge" class="hover:text-blue-400 transition-colors">
              <component :is="isEnlarged ? LucideMinimize2 : LucideMaximize2" class="w-5 h-5" />
            </button>
          </div>
        </div>
      </div>
      
      <!-- Viewer Content Area -->
      <div 
        class="bg-[#8E9194] p-0 flex flex-col justify-start overflow-auto relative transition-all"
        :class="isEnlarged ? 'h-[calc(100vh-60px)]' : 'min-h-[600px]'"
      >
        <!-- Debug Info (Only in Dev) -->
        <div class="absolute top-4 right-4 bg-black/50 text-white text-[8px] px-2 py-1 rounded z-10 font-mono" v-if="!isEnlarged">
          MIME: {{ doc?.mime_type }}
        </div>

        <div class="flex justify-center p-4 min-h-full">
          <iframe 
            v-if="doc?.mime_type === 'application/pdf' || (doc?.mime_type === 'application/octet-stream' && doc?.file_name?.toLowerCase().endsWith('.pdf'))"
            :src="documentUrl" 
            class="w-full border-none transition-transform duration-300 shadow-2xl"
            :class="isEnlarged ? 'h-full' : 'h-[800px]'"
            :style="{ transform: `scale(${zoomLevel / 100})`, transformOrigin: 'top center' }"
          ></iframe>
          <div v-else-if="doc?.mime_type?.startsWith('image/')" class="flex justify-center w-full">
            <img 
              :src="documentUrl" 
              class="max-w-full shadow-2xl transition-transform duration-300 object-contain" 
              :style="{ transform: `scale(${zoomLevel / 100})`, transformOrigin: 'top center' }"
            />
          </div>
          <div v-else class="flex flex-col items-center justify-center text-white/50 py-20 w-full">
            <LucideFileText class="w-20 h-20 mb-4 opacity-20" />
            <p class="text-xs font-black uppercase tracking-[0.2em]">{{ $t('documents.detail.viewer.no_preview') }}</p>
            <p class="text-[10px] mt-2 opacity-50 uppercase tracking-widest">{{ doc?.mime_type }}</p>
          </div>
        </div>
        
      </div>
    </div>

    <!-- Actions Bar -->
    <div class="flex flex-wrap items-center justify-between gap-6 py-4" v-motion-fade>
      <div class="flex items-center gap-6">
        <div class="space-y-1">
          <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ doc?.title }}</h3>
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ doc?.id }}</p>
        </div>
        <span class="px-4 py-1.5 bg-green-50 text-green-600 rounded-full text-[10px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2">
          <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
          {{ $t('documents.detail.actions.available') }}
        </span>
      </div>

      <div class="flex items-center gap-3">
        <!-- Approval Actions (If Pending) -->
        <template v-if="doc?.status === 'pending' || doc?.status === 'processing'">
          <button 
            @click="approveDocument"
            class="flex items-center gap-3 px-8 py-3.5 bg-green-600 hover:bg-green-700 text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-green-900/20 transition-all"
          >
            <LucideCheck class="w-4 h-4" />
            Setujui Dokumen
          </button>
          <button 
            @click="rejectDocument"
            class="flex items-center gap-3 px-8 py-3.5 bg-red-50 border border-red-200 text-red-600 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-red-100 transition-all"
          >
            <LucideX class="w-4 h-4" />
            Tolak
          </button>
        </template>

        <!-- Standard Actions (If Active) -->
        <template v-else>
          <button class="flex items-center gap-3 px-8 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all">
            <LucideShoppingCart class="w-4 h-4" />
            {{ $t('documents.detail.actions.add_to_cart') }}
          </button>
          <button class="flex items-center gap-3 px-8 py-3.5 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all">
            <LucideFileDown class="w-4 h-4" />
            {{ $t('documents.detail.actions.request_digital') }}
          </button>
        </template>
        
        <button class="p-3.5 bg-white border border-slate-200 rounded-xl text-slate-400 hover:text-slate-600 transition-all">
          <LucideMoreHorizontal class="w-5 h-5" />
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
      <div class="lg:col-span-7 space-y-10">
        <div class="flex flex-wrap gap-2" v-motion-fade>
          <span v-for="tag in doc?.tags || []" :key="tag" class="px-4 py-1.5 bg-white border border-slate-100 text-slate-400 text-[10px] font-black rounded-lg uppercase tracking-widest hover:border-primary-500 hover:text-primary-600 transition-colors cursor-pointer">#{{ tag }}</span>
          <span v-if="!(doc?.tags?.length)" class="text-[10px] font-bold text-slate-300 italic uppercase tracking-widest">{{ $t('documents.detail.metadata.no_tags') }}</span>
        </div>
        <div class="space-y-6" v-motion-fade>
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
            <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.metadata.title') }}</h2>
          </div>
          <div class="glass rounded-[2rem] overflow-hidden border border-slate-100">
            <table class="w-full text-sm">
              <tbody class="divide-y divide-slate-50">
                <tr v-for="(val, label) in metadata" :key="label" class="group">
                  <td class="px-8 py-4 w-1/3 bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ label }}</td>
                  <td class="px-8 py-4 font-bold text-[#1E3A5F]">{{ val }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div class="space-y-6" v-motion-fade>
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
            <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.location.title') }}</h2>
          </div>
          <div class="bg-slate-50/50 border border-slate-100 rounded-[2rem] p-8 flex items-center justify-between relative overflow-hidden group">
            <div class="space-y-6 relative z-10">
              <div class="grid grid-cols-2 gap-x-12 gap-y-6">
                <div class="flex items-start gap-4">
                  <LucideHome class="w-5 h-5 text-slate-400 mt-0.5" />
                  <div class="space-y-1">
                    <p class="text-xs font-bold text-slate-700">{{ doc?.branch_name || '-' }} / {{ doc?.department_name || '-' }}</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('documents.detail.location.building_room') }}</p>
                  </div>
                </div>
                <div class="flex items-start gap-4">
                  <LucideLayers class="w-5 h-5 text-slate-400 mt-0.5" />
                  <div class="space-y-1">
                    <p class="text-xs font-bold text-slate-700">{{ doc?.rack_name || '-' }} / {{ doc?.box_name || '-' }}</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('documents.detail.location.cabinet_shelf') }}</p>
                  </div>
                </div>
              </div>
              <div class="flex items-center gap-4 bg-white/60 p-4 rounded-2xl border border-white/80 shadow-sm w-fit">
                <LucideBox class="w-5 h-5 text-primary-600" />
                <p class="text-sm font-black text-primary-600 uppercase tracking-tight">{{ doc?.ordner_name || 'No Folder' }}</p>
              </div>
            </div>
            <div class="relative w-24 h-24 bg-white rounded-2xl shadow-xl flex items-center justify-center p-2 group-hover:scale-105 transition-transform">
              <LucideQrCode class="w-full h-full text-slate-300" />
              <div class="absolute bottom-1 right-1 bg-primary-500 text-white p-1 rounded-md"><LucideLocateFixed class="w-3 h-3" /></div>
              <span class="absolute -top-3 -right-3 bg-white px-2 py-1 rounded-md text-[8px] font-black text-slate-400 shadow-sm border border-slate-50 uppercase tracking-widest">{{ $t('documents.detail.location.scan') }}</span>
            </div>
            <LucideMapPin class="absolute top-4 right-4 w-5 h-5 text-slate-200 group-hover:text-primary-200 transition-colors" />
          </div>
        </div>
        <div class="space-y-6" v-motion-fade>
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
            <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.related.title') }}</h2>
          </div>
          <div class="space-y-4">
            <div v-for="rel in relatedDocs" :key="rel.name" class="group flex items-center justify-between p-5 bg-white border border-slate-100 rounded-3xl hover:border-primary-200 hover:shadow-lg hover:shadow-primary-500/5 transition-all cursor-pointer">
              <div class="flex items-center gap-5">
                <div :class="`w-12 h-12 rounded-2xl flex items-center justify-center ${rel.bg}`"><LucideFileText :class="`w-6 h-6 ${rel.color}`" /></div>
                <div class="space-y-1">
                  <p class="text-sm font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors uppercase tracking-tight">{{ rel.name }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ rel.desc }}</p>
                </div>
              </div>
              <LucideChevronRight class="w-5 h-5 text-slate-300 group-hover:translate-x-1 transition-transform" />
            </div>
          </div>
        </div>
      </div>
      <div class="lg:col-span-5 space-y-8" v-motion-fade>
        <div class="flex items-center gap-3">
          <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
          <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.loan_history.title') }}</h2>
        </div>
        <div class="relative pl-12 space-y-12">
          <div class="absolute left-6 top-4 bottom-4 w-0.5 bg-slate-100"></div>
          <div v-for="(loan, index) in timeline" :key="index" class="relative group">
            <div class="absolute -left-6 top-1.5 w-6 h-6 rounded-full bg-white border-4 border-slate-100 group-hover:border-primary-200 transition-colors z-10"></div>
            <div class="space-y-4">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 rounded-2xl bg-slate-100 border-2 border-white shadow-sm overflow-hidden flex-shrink-0"><img :src="`https://i.pravatar.cc/150?u=${loan.name}`" class="w-full h-full object-cover" /></div>
                <div>
                  <p class="text-sm font-black text-[#1E3A5F]">{{ loan.name }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Dept. {{ loan.dept }}</p>
                </div>
              </div>
              <div class="bg-slate-50/50 p-4 rounded-2xl border border-slate-100 group-hover:bg-primary-50/30 group-hover:border-primary-100 transition-all">
                <p class="text-[10px] font-bold text-slate-500 leading-relaxed uppercase tracking-widest" v-html="$t('documents.detail.loan_history.returned', { date: loan.date }) + ' • ' + $t('documents.detail.loan_history.duration', { days: loan.duration })"></p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
const { t } = useI18n()
const route = useRoute()
const config = useRuntimeConfig()
const { $api } = useApi()

import { ref, computed } from 'vue'
import { 
  LucideArrowLeft, LucideMenu, LucideMinus, LucidePlus, LucidePrinter, LucideDownload, 
  LucideShoppingCart, LucideFileDown, LucideMoreHorizontal, LucideHome, LucideLayers, LucideBox,
  LucideQrCode, LucideLocateFixed, LucideMapPin, LucideChevronRight, LucideFileText, LucideAlertCircle, LucideMaximize2, LucideMinimize2,
  LucideCheck, LucideX
} from 'lucide-vue-next'

const isEnlarged = ref(false)
const currentPage = ref(1)
const zoomLevel = ref(100)

const toggleEnlarge = () => {
  isEnlarged.value = !isEnlarged.value
}

const zoomIn = () => {
  if (zoomLevel.value < 200) zoomLevel.value += 25
}

const zoomOut = () => {
  if (zoomLevel.value > 50) zoomLevel.value -= 25
}

const printDocument = () => {
  const printWindow = window.open(documentUrl.value, '_blank')
  printWindow?.print()
}

const downloadDocument = () => {
  const downloadUrl = `${documentUrl.value}&download=true`
  window.location.href = downloadUrl
}

// Data Fetching
const { data: docRes } = await useAsyncData(`doc-${route.params.id}`, () => 
  $api(`/documents/${route.params.id}`)
)

const { data: loansRes } = await useAsyncData(`loans-${route.params.id}`, () => 
  $api(`/documents/${route.params.id}/loans`)
)

const doc = computed(() => docRes.value?.data)
const loans = computed(() => loansRes.value?.data || [])

const approveDocument = async () => {
  // TODO: Call API to approve
  alert('Dokumen disetujui')
}

const rejectDocument = async () => {
  // TODO: Call API to reject
  alert('Dokumen ditolak')
}

const documentUrl = computed(() => {
  if (!doc.value) return ''
  const auth = useAuthStore()
  return `${config.public.apiBase}/documents/${doc.value.id}/preview?token=${auth.accessToken}`
})

const metadata = computed(() => {
  if (!doc.value) return {}
  
  // Base fields
  const base = {
    [t('documents.detail.metadata.labels.doc_type')]: doc.value.type_name?.String || doc.value.type_name || (doc.value.status === 'processing' ? 'Processing...' : 'General'),
    [t('documents.detail.metadata.labels.company')]: doc.value.company_name?.String || doc.value.company_name || '-',
    [t('documents.detail.metadata.labels.department')]: doc.value.department_name?.String || doc.value.department_name || '-',
    [t('documents.detail.metadata.labels.year')]: doc.value.created_at ? new Date(doc.value.created_at).getFullYear() : '-',
    [t('documents.detail.metadata.labels.date_issued')]: doc.value.created_at ? new Date(doc.value.created_at).toLocaleDateString() : '-',
    [t('documents.detail.metadata.labels.retention')]: `${doc.value.retention_years || 7} Years`,
    [t('documents.detail.metadata.labels.version')]: `v${doc.value.current_version || 1}.0`
  }

  // Extract from JSONB metadata
  const custom = {}
  if (doc.value.metadata) {
    const meta = typeof doc.value.metadata === 'string' ? JSON.parse(doc.value.metadata) : doc.value.metadata
    if (meta.urgency) custom['Urgency'] = meta.urgency
    if (meta.document_date) custom['Document Date'] = meta.document_date
    if (meta.page_count) custom['Total Pages'] = meta.page_count
    if (meta.sensitivity) custom['Sensitivity'] = meta.sensitivity
  }

  return { ...base, ...custom }
})

const relatedDocs = computed(() => {
  // Simple logic: docs from same dept (this would ideally be a separate API call)
  return [] 
})

const timeline = computed(() => {
  const history = []
  
  // Always include upload event
  if (doc.value) {
    history.push({ 
      name: doc.value.owner_name || 'Owner', 
      dept: doc.value.department_name || 'Upload', 
      date: new Date(doc.value.created_at).toLocaleDateString(), 
      duration: 0,
      status: 'uploaded'
    })
  }

  // Include real loans
  loans.value.forEach(l => {
    history.push({
      name: l.user_name,
      dept: l.department_name,
      date: l.borrow_date ? new Date(l.borrow_date).toLocaleDateString() : 'Pending',
      duration: l.return_date ? 'Returned' : 'In Use',
      status: l.status
    })
  })

  return history
})
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
