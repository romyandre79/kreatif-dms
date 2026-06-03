<template>
  <div class="min-h-screen bg-[#F8FAFC] p-8 space-y-8" v-motion-fade>
    <!-- Breadcrumbs & Actions -->
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-4">
        <button @click="router.back()" class="px-6 py-2.5 bg-slate-100 hover:bg-slate-200 text-slate-600 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all">
          Cancel
        </button>
        <button 
          @click="handleFinalize" 
          :disabled="assigning" 
          class="px-8 py-2.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
        >
          <LucideCheckCircle v-if="!assigning" class="w-4 h-4" />
          <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
          Finalize Assignment
        </button>
      </div>
    </div>

    <!-- Page Title -->
    <div class="space-y-1">
      <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tighter uppercase">Label Preview & Box Assignment</h1>
      <p class="text-xs font-bold text-slate-500 italic">Configure and assign item labels to organizational storage units.</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Main Content -->
      <div class="lg:col-span-8 space-y-8">
        <!-- Recommended Storage Target -->
        <div class="bg-white rounded-3xl p-8 border border-slate-100 shadow-sm space-y-6">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-2xl bg-primary-50 text-primary-500 flex items-center justify-center">
              <LucideArchive class="w-5 h-5" />
            </div>
            <h2 class="text-sm font-black text-[#1E3A5F] uppercase tracking-widest">Recommended Storage Target</h2>
          </div>
          
          <div class="p-8 rounded-2xl bg-orange-50/50 border border-orange-100 flex items-center justify-between transition-all" :class="{'ring-4 ring-orange-500/20': !selectedBox}">
            <div class="space-y-1">
              <p class="text-[9px] font-black text-orange-400 uppercase tracking-widest">Target Rack Location</p>
              <h3 class="text-3xl font-black text-[#1E3A5F] tracking-tighter">
                {{ recommendedLocation?.rack_name || 'AUTO-SELECTING...' }}
              </h3>
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">
                {{ recommendedLocation?.location_detail || 'Calculating optimal zonation...' }}
              </p>
            </div>
            <button 
              @click="focusSearch"
              class="flex items-center gap-2 px-5 py-2.5 bg-white border border-slate-200 rounded-xl text-[10px] font-black text-slate-600 uppercase tracking-widest hover:bg-slate-50 transition-all active:scale-95"
            >
              <LucideEdit3 class="w-3.5 h-3.5" />
              Change
            </button>
          </div>
        </div>

        <!-- Label Generation Preview -->
        <div class="space-y-6">
          <div class="flex items-center justify-between px-2">
            <div class="flex items-center gap-3">
              <LucidePrinter class="w-5 h-5 text-primary-500" />
              <h2 class="text-sm font-black text-[#1E3A5F] uppercase tracking-widest">Label Generation Preview</h2>
            </div>
            <button @click="triggerPrint" class="text-[10px] font-black text-primary-500 uppercase tracking-widest hover:underline flex items-center gap-2">
              <LucideExternalLink class="w-3 h-3" />
              Preview Full Sheet
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <!-- Child Item Label -->
            <div class="space-y-4">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">Child Item Label (50x30mm)</p>
              <div class="aspect-[5/3] bg-white border border-slate-200 shadow-lg rounded-2xl p-6 flex flex-col justify-center relative overflow-hidden group">
                <div class="flex gap-4 items-start">
                  <div class="w-16 h-16 bg-slate-100 rounded-lg flex items-center justify-center shrink-0">
                    <LucideQrCode class="w-10 h-10 text-slate-400" />
                  </div>
                  <div class="space-y-1.5 min-w-0">
                    <p class="text-[7px] font-black text-slate-400 uppercase tracking-widest">PT AKIRADATA</p>
                    <p class="text-sm font-black text-[#1E3A5F] tracking-tight leading-none uppercase truncate">
                      {{ firstDoc?.title || 'AKR-CH-00582' }}
                    </p>
                    <div class="space-y-0.5">
                      <p class="text-[6px] font-black text-slate-300 uppercase tracking-widest">LOCATION CODE</p>
                      <p class="text-[8px] font-black text-slate-600 uppercase">{{ recommendedLocation?.rack_name || 'PENDING' }}</p>
                    </div>
                    <p class="text-[7px] font-bold text-slate-400">System ID: {{ docIDs[0]?.substring(0,8).toUpperCase() || 'DOC-ID' }}</p>
                  </div>
                </div>
                <div class="absolute inset-0 bg-[#1E3A5F]/0 group-hover:bg-[#1E3A5F]/5 transition-colors cursor-zoom-in"></div>
              </div>
            </div>

            <!-- Parent Box Label -->
            <div class="space-y-4">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">Parent Box Label (100x70mm)</p>
              <div class="aspect-[5/3] bg-white border border-slate-200 shadow-lg rounded-2xl p-8 flex flex-col justify-between relative group">
                <div class="flex justify-between items-start">
                  <div class="space-y-1">
                    <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">PT AKIRADATA</p>
                    <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest mt-2">MASTER BOX ID</p>
                    <p class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ selectedBox?.name || 'AKR-BX-2024-092' }}</p>
                  </div>
                  <div class="w-16 h-16 bg-slate-800 rounded-lg p-2 flex items-center justify-center shrink-0">
                    <LucideQrCode class="w-full h-full text-white" />
                  </div>
                </div>
                <div class="border-t border-slate-100 pt-4 space-y-1">
                  <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">HIERARCHY LOCATION</p>
                  <p class="text-sm font-black text-slate-500 tracking-tight uppercase truncate">
                    {{ recommendedLocation?.rack_name || 'SELECT RACK' }} / {{ recommendedLocation?.location_detail || 'LVL-03' }}
                  </p>
                </div>
                <div class="absolute inset-0 bg-[#1E3A5F]/0 group-hover:bg-[#1E3A5F]/5 transition-colors cursor-zoom-in"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Sidebar -->
      <div class="lg:col-span-4">
        <div class="bg-white rounded-3xl shadow-2xl shadow-slate-200/50 border border-slate-100 overflow-hidden sticky top-8">
          <div class="bg-[#1E3A5F] p-6 flex items-center justify-between">
            <h3 class="text-sm font-black text-white uppercase tracking-widest">Kelompokkan ke Boks</h3>
            <button @click="router.back()" class="text-white/60 hover:text-white"><LucideX class="w-5 h-5" /></button>
          </div>
          
          <div class="p-8 space-y-8">
            <!-- Box Selection -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Select Parent Box</label>
              <div class="relative group">
                <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
                <input 
                  ref="searchInput"
                  type="text" 
                  v-model="boxSearchQuery"
                  @input="handleBoxSearch"
                  @keyup.enter="handleSearchEnter"
                  placeholder="Search box ID or location..." 
                  class="w-full pl-12 pr-5 py-4 bg-slate-50 border-none rounded-2xl text-xs font-bold outline-none ring-2 ring-transparent focus:ring-primary-500/10 focus:bg-white transition-all shadow-inner"
                />
                
                <!-- Search Results Dropdown -->
                <div v-if="boxSearchResults.length > 0 && showBoxDropdown" class="absolute top-full left-0 right-0 mt-3 bg-white rounded-2xl shadow-2xl border border-slate-100 z-50 overflow-hidden py-2" v-motion-pop>
                  <button 
                    v-for="box in boxSearchResults" 
                    :key="box.id"
                    @click="selectBox(box)"
                    class="w-full text-left px-6 py-4 hover:bg-slate-50 transition-colors flex items-center justify-between group"
                  >
                    <div>
                      <p class="text-xs font-black text-[#1E3A5F] group-hover:text-primary-500">{{ box.name }}</p>
                      <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ box.rack_name }} • {{ box.department_name }}</p>
                    </div>
                    <LucideArrowRight class="w-4 h-4 text-slate-200 group-hover:translate-x-1 group-hover:text-primary-500 transition-all" />
                  </button>
                </div>
              </div>
            </div>

            <!-- Current Selection Summary -->
            <div class="p-6 bg-slate-50 rounded-2xl border-2 border-dashed border-slate-200 space-y-4 relative overflow-hidden">
              <div class="flex items-center justify-between px-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Current Item Selection</p>
                <span class="text-[9px] font-black text-[#1E3A5F] uppercase">{{ docIDs.length }} Items</span>
              </div>
              <ul class="space-y-3">
                <li v-for="(doc, idx) in selectedDocuments.slice(0, 3)" :key="doc.id" class="flex items-center gap-3">
                  <div class="w-1.5 h-1.5 bg-[#1E3A5F] rounded-full"></div>
                  <div class="flex flex-col">
                    <span class="text-[10px] font-black text-[#1E3A5F] tracking-tight uppercase">{{ doc.title }}</span>
                    <span class="text-[8px] text-slate-400 font-medium">SYS-{{ doc.id.substring(0,8).toUpperCase() }}</span>
                  </div>
                </li>
                <template v-if="selectedDocuments.length === 0">
                  <li v-for="(id, idx) in docIDs.slice(0, 3)" :key="id" class="flex items-center gap-3">
                    <div class="w-1.5 h-1.5 bg-[#1E3A5F] rounded-full"></div>
                    <span class="text-[10px] font-black text-[#1E3A5F] tracking-tight uppercase">AKR-CH-{{ id.substring(0,8).toUpperCase() }}</span>
                  </li>
                </template>
                <li v-if="docIDs.length > 3" class="text-[9px] font-bold text-slate-400 italic pl-4">
                  + {{ docIDs.length - 3 }} other documents
                </li>
              </ul>
              <!-- Background Icon -->
              <LucideLayers class="absolute -right-4 -bottom-4 w-20 h-20 text-slate-200 opacity-20 -rotate-12" />
            </div>

            <!-- Action Buttons -->
            <div class="space-y-3 pt-4">
              <button 
                @click="handleFinalize"
                :disabled="assigning"
                class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 disabled:opacity-50 group"
              >
                <LucidePackage class="w-4 h-4 group-hover:scale-110 transition-transform" />
                Assign to Box
              </button>
              <button @click="saveAsDraft" class="w-full py-5 bg-white border border-slate-200 text-slate-600 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all active:scale-95">
                Save as Draft
              </button>
            </div>

            <!-- Info Box -->
            <div class="p-5 bg-primary-50/50 rounded-2xl flex gap-4">
              <LucideInfo class="w-4 h-4 text-primary-500 shrink-0" />
              <p class="text-[10px] font-bold text-slate-500 leading-relaxed italic">
                Once assigned, parent labels will automatically include these items in the manifest. QR will be synchronized instantly.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Printing Overlay -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100" leave-active-class="transition duration-200 ease-in" leave-from-class="opacity-100" leave-to-class="opacity-0">
      <div v-if="printing" class="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-sm">
        <div class="bg-white dark:bg-slate-900 w-full max-w-sm rounded-3xl shadow-2xl p-10 flex flex-col items-center text-center space-y-6">
          <div class="relative">
            <LucidePrinter class="w-16 h-16 text-[#1E3A5F] animate-bounce" />
            <div class="absolute -top-2 -right-2 w-6 h-6 bg-orange-500 rounded-full flex items-center justify-center border-4 border-white">
              <div class="w-1.5 h-1.5 bg-white rounded-full animate-ping"></div>
            </div>
          </div>
          <div class="space-y-2">
            <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight">Printing Labels...</h3>
            <p class="text-xs text-slate-500 font-medium">Sending {{ docIDs.length + 1 }} labels to the thermal printer</p>
          </div>
          <div class="w-full h-1.5 bg-slate-100 rounded-full overflow-hidden">
            <div class="h-full bg-orange-500 animate-progress"></div>
          </div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest animate-pulse">DO NOT CLOSE THIS PAGE</p>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { 
  LucideChevronRight, LucideX, LucideSearch, LucideArrowRight, 
  LucidePackage, LucideInfo, LucideArchive, LucideEdit3,
  LucidePrinter, LucideQrCode, LucideCheckCircle, LucideLoader2,
  LucideExternalLink, LucideLayers
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'
import { useToast } from '~/composables/useToast'

const route = useRoute()
const router = useRouter()
const { $api } = useApi()
const toast = useToast()

const docIDs = ref([])
const selectedDocuments = ref([])
const assigning = ref(false)
const printing = ref(false)
const boxSearchQuery = ref('')
const boxSearchResults = ref([])
const showBoxDropdown = ref(false)
const selectedBox = ref(null)
const recommendedLocation = ref(null)
const firstDoc = ref(null)
const searchInput = ref(null)

onMounted(async () => {
  const ids = route.query.ids
  if (ids) {
    docIDs.value = ids.split(',')
    
    // Fetch all selected documents for preview & list
    try {
      const promises = docIDs.value.map(id => $api(`/documents/${id}`))
      const responses = await Promise.all(promises)
      selectedDocuments.value = responses.map(res => res.data)
      firstDoc.value = selectedDocuments.value[0]
    } catch (err) {
      console.error('Failed to fetch selected documents:', err)
    }

    // Fetch smart recommendation from backend
    try {
      const res = await $api(`/intake/boxes/recommend?ids=${ids}`)
      if (res.data) {
        selectBox(res.data)
      } else {
        recommendedLocation.value = {
          rack_name: 'AUTO-SELECTING...',
          location_detail: 'No optimal box found in your department. Please search manually.'
        }
      }
    } catch (err) {
      console.error('Failed to get recommendation:', err)
    }
  } else {
    router.back()
  }
})

const handleBoxSearch = async () => {
  if (boxSearchQuery.value.length < 2) {
    boxSearchResults.value = []
    showBoxDropdown.value = false
    return
  }
  
  try {
    const res = await $api(`/intake/boxes/search?q=${boxSearchQuery.value}`)
    boxSearchResults.value = res.data || []
    showBoxDropdown.value = true
  } catch (err) {
    console.error('Failed to search boxes:', err)
  }
}

const selectBox = (box) => {
  selectedBox.value = box
  boxSearchQuery.value = box.name
  showBoxDropdown.value = false
  
  recommendedLocation.value = {
    rack_name: box.rack_name,
    location_detail: box.location_detail
  }
}

const focusSearch = () => {
  searchInput.value?.focus()
}

const handleSearchEnter = () => {
  if (boxSearchResults.value.length > 0) {
    selectBox(boxSearchResults.value[0])
  }
}

const triggerPrint = () => {
  printing.value = true
  setTimeout(() => {
    printing.value = false
    toast.info('Simulasi: Label telah dikirim ke antrean printer.')
  }, 3000)
}

const saveAsDraft = () => {
  toast.success('Draft penugasan berhasil disimpan.')
  router.push('/warehouse/labels')
}

const handleFinalize = async () => {
  if (!selectedBox.value) {
    toast.warning('Silakan cari dan pilih boks penyimpanan terlebih dahulu.')
    return
  }
  if (docIDs.value.length === 0) {
    toast.error('Tidak ada dokumen yang dipilih.')
    return
  }
  
  assigning.value = true
  try {
    // 1. Assign in DB
    await $api('/intake/boxes/assign', {
      method: 'POST',
      body: {
        document_ids: docIDs.value,
        box_id: selectedBox.value.id
      }
    })
    
    // 2. Trigger printing simulation before finishing
    printing.value = true
    await new Promise(resolve => setTimeout(resolve, 3000))
    printing.value = false
    
    toast.success('Penugasan boks dan pencetakan label berhasil!')
    router.push('/warehouse/labels')
  } catch (err) {
    toast.error('Gagal menyelesaikan penugasan: ' + err.message)
  } finally {
    assigning.value = false
  }
}
</script>

<style scoped>
@keyframes progress {
  0% { width: 0; }
  100% { width: 100%; }
}
.animate-progress {
  animation: progress 3s linear infinite;
}
</style>
