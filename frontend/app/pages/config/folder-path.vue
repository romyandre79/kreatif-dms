<template>
  <div class="min-h-screen bg-[#F8FAFC] p-8">
    <div class="max-w-6xl mx-auto space-y-8">
      <!-- Header -->
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <button @click="router.back()" class="p-3 bg-white border border-slate-200 rounded-2xl text-slate-400 hover:text-slate-600 transition-all shadow-sm">
            <LucideArrowLeft class="w-5 h-5" />
          </button>
          <div>
            <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">Folder Path Preview</h1>
            <p class="text-slate-500 font-medium">Automatic target generation based on document metadata strings.</p>
          </div>
        </div>
        
        <!-- Storage Health Widget (Mock) -->
        <div class="bg-white p-4 rounded-2xl border border-slate-100 shadow-sm flex items-center gap-4">
          <div class="text-right">
            <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Storage Health</p>
            <p class="text-sm font-black text-[#1E3A5F]">84.2 GB / 100 GB Used</p>
          </div>
          <div class="w-10 h-10 rounded-full border-4 border-blue-500 border-t-transparent animate-spin-slow flex items-center justify-center">
             <LucideCloud class="w-4 h-4 text-blue-500" />
          </div>
        </div>
      </div>

      <div class="grid grid-cols-12 gap-8">
        <!-- Main Explorer Preview -->
        <div class="col-span-8 space-y-6">
          <div class="bg-white rounded-[2.5rem] border border-slate-200 shadow-xl shadow-blue-900/5 overflow-hidden">
            <!-- Target Hierarchy Bar -->
            <div class="bg-[#F1F5F9]/50 border-b border-slate-200 p-8 flex items-center justify-between">
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 bg-white rounded-xl shadow-sm border border-slate-100 flex items-center justify-center text-primary-600">
                  <LucideLayers class="w-5 h-5" />
                </div>
                <div>
                  <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Target Hierarchy</p>
                  <p class="text-sm font-bold text-slate-600">/ARSIP/{{ hierarchyFormula }}</p>
                </div>
              </div>
              <button @click="fetchPreview" class="p-2 hover:bg-slate-200 rounded-lg transition-colors text-slate-400">
                <LucideRotateCcw class="w-5 h-5" :class="{ 'animate-spin': isRefreshing }" />
              </button>
            </div>

            <!-- Preview Tree View -->
            <div class="p-10 min-h-[500px]">
              <div class="space-y-4">
                <div class="flex items-center gap-3">
                  <LucideFolder class="w-6 h-6 text-amber-400 fill-amber-400" />
                  <span class="text-lg font-black text-slate-800 tracking-tight uppercase">ARSIP</span>
                  <span class="px-2 py-0.5 bg-slate-100 text-[10px] font-black text-slate-400 rounded-md uppercase">Root</span>
                </div>

                <!-- Recursive Tree (Mocked based on levels) -->
                <div class="ml-4 pl-8 border-l-2 border-slate-100 space-y-6 pt-2">
                  <div class="flex items-center gap-3">
                    <LucideFolder class="w-5 h-5 text-amber-400 fill-amber-400" />
                    <span class="text-sm font-bold text-slate-600 uppercase tracking-tight">AKIRADATA_PRATAMA</span>
                    <span class="text-[10px] text-blue-400 font-bold">{PT}</span>
                  </div>

                  <div class="ml-4 pl-8 border-l-2 border-slate-100 space-y-6">
                    <div class="flex items-center gap-3">
                      <LucideFolder class="w-5 h-5 text-amber-400 fill-amber-400" />
                      <span class="text-sm font-bold text-slate-600 uppercase tracking-tight">FINANCE_TAX</span>
                      <span class="text-[10px] text-blue-400 font-bold">{Dept}</span>
                    </div>

                    <div class="ml-4 pl-8 border-l-2 border-slate-100 space-y-6">
                      <div class="flex items-center gap-3">
                        <LucideFolder class="w-5 h-5 text-amber-400 fill-amber-400" />
                        <span class="text-sm font-bold text-slate-600 uppercase tracking-tight">2024</span>
                        <span class="text-[10px] text-blue-400 font-bold">{Tahun}</span>
                      </div>

                      <div class="ml-4 pl-8 border-l-2 border-slate-100 space-y-4">
                        <div class="flex items-center gap-3">
                          <LucideFolderPlus class="w-5 h-5 text-amber-400" />
                          <span class="text-sm font-bold text-slate-600 uppercase tracking-tight">INV3ICE_3иж</span>
                          <span class="text-[10px] text-blue-400 font-bold">{Tipe}</span>
                          <span class="px-2 py-0.5 bg-blue-50 text-blue-600 text-[8px] font-black rounded uppercase">New Folder</span>
                        </div>

                        <!-- Mock Document -->
                        <div class="bg-slate-50 border border-slate-100 rounded-2xl p-4 flex items-center gap-4 max-w-sm">
                          <div class="w-10 h-10 bg-white rounded-xl shadow-sm flex items-center justify-center text-blue-500">
                            <LucideFileText class="w-6 h-6" />
                          </div>
                          <div>
                            <p class="text-xs font-black text-slate-800 uppercase tracking-tight">INV-2024-00192_PT_AD.pdf</p>
                            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-1">2.4 MB • OCR Verified</p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Bottom Actions -->
          <div class="flex items-center justify-between pt-4">
            <button @click="router.back()" class="flex items-center gap-2 text-sm font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest">
              <LucideArrowLeft class="w-4 h-4" /> Back to Naming
            </button>
            <button @click="saveConfig" :disabled="isSaving" class="flex items-center gap-3 px-12 py-5 bg-[#3B82F6] hover:bg-blue-500 text-white rounded-2xl shadow-xl shadow-blue-500/30 transition-all font-black uppercase tracking-[0.15em] text-sm group">
              <span v-if="!isSaving">Commit Folder Placement</span>
              <span v-else>Saving...</span>
              <LucideCheckCircle2 class="w-5 h-5 group-hover:scale-110 transition-transform" />
            </button>
          </div>
        </div>

        <!-- Settings Sidebar -->
        <div class="col-span-4 space-y-8">
          <!-- Foldering Rules Section -->
          <div class="space-y-6">
            <div class="flex items-center gap-3 text-slate-800">
              <LucideSettings2 class="w-5 h-5" />
              <h2 class="text-sm font-black uppercase tracking-widest">Foldering Rules</h2>
            </div>

            <!-- Policy Banner -->
            <div class="bg-white p-6 rounded-3xl border border-slate-100 shadow-sm space-y-4">
               <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Rule Source Policy</p>
               <div class="flex items-start gap-4 p-4 bg-green-50/50 border border-green-100 rounded-2xl">
                 <LucideCheckCircle2 class="w-5 h-5 text-green-500 mt-1" />
                 <div>
                   <p class="text-sm font-black text-green-700 uppercase">Dept Policy: FINANCE</p>
                   <p class="text-xs text-green-600/80 font-medium leading-relaxed mt-1">System automatically maps folder targets based on Financial Compliance (ISO-27001).</p>
                 </div>
               </div>
            </div>

            <!-- Path Formula Builder -->
            <div class="bg-white p-8 rounded-[2.5rem] border border-slate-100 shadow-sm space-y-6">
               <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Path Formula Preview</p>
               
               <div class="bg-[#1E3A5F] p-6 rounded-2xl flex flex-wrap gap-2 min-h-[100px] content-start">
                 <span class="px-3 py-1.5 bg-blue-900/30 text-blue-200 text-[10px] font-black rounded-lg border border-blue-200/20">/ARSIP/</span>
                 
                 <div 
                   v-for="(level, index) in levels" 
                   :key="level.key"
                   class="flex items-center"
                 >
                    <div class="px-3 py-1.5 bg-blue-400 text-white text-[10px] font-black rounded-lg shadow-lg flex items-center gap-2 group cursor-pointer relative">
                      ${{ level.label }}
                      <div class="hidden group-hover:flex absolute -top-8 left-0 gap-1">
                         <button @click="moveLevel(index, -1)" class="p-1 bg-white text-[#1E3A5F] rounded shadow hover:bg-slate-100"><LucideChevronLeft class="w-3 h-3"/></button>
                         <button @click="moveLevel(index, 1)" class="p-1 bg-white text-[#1E3A5F] rounded shadow hover:bg-slate-100"><LucideChevronRight class="w-3 h-3"/></button>
                      </div>
                    </div>
                    <span v-if="index < levels.length - 1" class="text-blue-200/30 px-1">/</span>
                 </div>
               </div>

               <!-- Add Level Dropdown (Simple version) -->
               <div class="space-y-4">
                 <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Configuration Settings</p>
                 
                 <!-- Create Missing Folders Toggle -->
                 <div class="flex items-center justify-between p-4 bg-slate-50 rounded-2xl border border-slate-100">
                    <div>
                      <p class="text-xs font-black text-slate-800 uppercase tracking-tight">Create if not exists</p>
                      <p class="text-[10px] text-slate-400 font-medium">Auto-generate missing directories</p>
                    </div>
                    <button 
                      @click="config.create_missing_folders = !config.create_missing_folders"
                      class="w-12 h-6 rounded-full transition-all relative"
                      :class="config.create_missing_folders ? 'bg-[#3B82F6]' : 'bg-slate-200'"
                    >
                      <div class="absolute top-1 w-4 h-4 bg-white rounded-full transition-all" :class="config.create_missing_folders ? 'left-7' : 'left-1'"></div>
                    </button>
                 </div>

                 <!-- Apply Retention Toggle -->
                 <div class="flex items-center justify-between p-4 bg-slate-50 rounded-2xl border border-slate-100">
                    <div>
                      <p class="text-xs font-black text-slate-800 uppercase tracking-tight">Apply Retention Tag</p>
                      <p class="text-[10px] text-slate-400 font-medium">Add 7-year purge metadata</p>
                    </div>
                    <button 
                      @click="config.apply_retention_tag = !config.apply_retention_tag"
                      class="w-12 h-6 rounded-full transition-all relative"
                      :class="config.apply_retention_tag ? 'bg-[#3B82F6]' : 'bg-slate-200'"
                    >
                      <div class="absolute top-1 w-4 h-4 bg-white rounded-full transition-all" :class="config.apply_retention_tag ? 'left-7' : 'left-1'"></div>
                    </button>
                 </div>
               </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { 
  LucideArrowLeft, LucideChevronLeft, LucideChevronRight, LucideFolder, LucideFileText, 
  LucideLayers, LucideRotateCcw, LucideCheckCircle2, LucideSettings2, LucideCloud, LucideFolderPlus
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'
import { useToast } from '~/composables/useToast'

const router = useRouter()
const { $api } = useApi()
const toast = useToast()

const isSaving = ref(false)
const isRefreshing = ref(false)

const allPossibleLevels = [
  { key: 'company', label: 'PT' },
  { key: 'branch', label: 'CABANG' },
  { key: 'department', label: 'DEPT' },
  { key: 'year', label: 'TAHUN' },
  { key: 'rack', label: 'RAK' },
  { key: 'box', label: 'BOX' },
  { key: 'ordner', label: 'ORDNER' },
  { key: 'type', label: 'TIPE' }
]

const levels = ref([
  { key: 'company', label: 'SCOMP' },
  { key: 'department', label: 'SDEPT' },
  { key: 'year', label: 'SYEAR' },
  { key: 'type', label: 'SDOC_TYPE' }
])

const config = ref({
  folder_path: 'company,department,year,type',
  create_missing_folders: true,
  apply_retention_tag: false
})

const hierarchyFormula = computed(() => {
  return levels.value.map(l => `{${l.label.replace('S', '')}}`).join('/')
})

const moveLevel = (index, delta) => {
  const newIndex = index + delta
  if (newIndex < 0 || newIndex >= levels.value.length) return
  
  const temp = levels.value[index]
  levels.value[index] = levels.value[newIndex]
  levels.value[newIndex] = temp
  
  updateConfigPath()
}

const updateConfigPath = () => {
  config.value.folder_path = levels.value.map(l => l.key).join(',')
}

const fetchConfig = async () => {
  try {
    const res = await $api('/documents/explorer/config')
    if (res && res.data) {
      config.value = { ...config.value, ...res.data }
      
      // Sync levels array
      const pathKeys = config.value.folder_path.split(',')
      levels.value = pathKeys.map(k => {
        const found = allPossibleLevels.find(p => p.key === k)
        return found ? { key: k, label: 'S' + found.label } : { key: k, label: 'S' + k.toUpperCase() }
      })
    }
  } catch (err) {
    console.error('Failed to fetch config:', err)
  }
}

const saveConfig = async () => {
  isSaving.value = true
  try {
    await $api('/documents/explorer/config', {
      method: 'POST',
      body: config.value
    })
    toast.success('Configuration saved successfully!')
  } catch (err) {
    toast.error('Failed to save configuration')
  } finally {
    isSaving.value = false
  }
}

const fetchPreview = async () => {
  isRefreshing.value = true
  // Mock refreshing for 1s
  setTimeout(() => {
    isRefreshing.value = false
  }, 1000)
}

onMounted(() => {
  fetchConfig()
})
</script>

<style scoped>
@keyframes spin-slow {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
.animate-spin-slow {
  animation: spin-slow 8s linear infinite;
}
</style>
