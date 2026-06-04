<template>
  <div class="min-h-screen bg-[#F8FAFC] p-6 pt-3">
    <div class="max-w-6xl mx-auto space-y-4">
      <!-- Header -->
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div>
            <h1 class="text-2xl lg:text-3xl font-extrabold text-[#1E3A5F] tracking-tight uppercase">Folder Path Preview</h1>
            <p class="text-slate-500 font-medium text-sm">Automatic target generation based on document metadata strings.</p>
          </div>
        </div>
        
        <!-- Storage Health Widget -->
        <div class="bg-white p-4 rounded-2xl border border-slate-100 shadow-sm flex items-center gap-4">
          <div class="text-right">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-tight" v-html="$t('dashboard.stats.warehouse_capacity').replace('\n', '<br>')"></p>
            <p class="text-sm font-black text-[#1E3A5F] mt-1">{{ formatBytes(storageUsed) }} / {{ formatBytes(storageLimitBytes) }}</p>
          </div>
          <div class="w-10 h-10 rounded-full border-4 border-blue-500 border-t-transparent animate-spin-slow flex items-center justify-center">
             <LucideCloud class="w-4 h-4 text-blue-500" />
          </div>
        </div>
      </div>
 
      <div class="grid grid-cols-12 gap-6">
        <!-- Main Explorer Preview -->
        <div class="col-span-8 space-y-4">
          <div class="bg-white rounded-3xl border border-slate-200/60 shadow-lg shadow-blue-900/5 overflow-hidden">
            <!-- Target Hierarchy Bar -->
            <div class="bg-[#F1F5F9]/50 border-b border-slate-200/60 p-5 flex items-center justify-between">
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
            <div class="p-6 min-h-[400px]">
              <div class="space-y-4">
                <div class="flex items-center gap-3">
                  <LucideFolder class="w-5 h-5 text-amber-400 fill-amber-400" />
                  <span class="text-base font-black text-slate-800 tracking-tight uppercase">ARSIP</span>
                  <span class="px-2 py-0.5 bg-slate-100 text-[9px] font-black text-slate-400 rounded-md uppercase">Root</span>
                </div>
 
                <!-- Dynamic Nested Tree -->
                <div class="pt-2 space-y-3">
                  <div 
                    v-for="item in flatPreviewItems" 
                    :key="item.id"
                    :style="{ paddingLeft: `${item.indent * 1.5}rem` }"
                    class="flex flex-col gap-2"
                  >
                    <!-- Folder Item -->
                    <div v-if="item.isFolder" class="flex items-center gap-3">
                      <div class="pl-2 border-l-2 border-slate-200/80 mr-1 py-0.5 flex items-center gap-3">
                        <component 
                          :is="item.isNew ? LucideFolderPlus : LucideFolder" 
                          class="w-5 h-5 text-amber-400 fill-amber-400"
                          :class="{ 'fill-none text-amber-500': item.isNew }"
                        />
                        <span class="text-sm font-semibold text-slate-600 uppercase tracking-tight">{{ item.name }}</span>
                        <span class="text-[9px] text-[#0ea5e9] font-bold">{{ item.tag }}</span>
                        <span v-if="item.isNew" class="px-2 py-0.5 bg-blue-50 text-blue-600 text-[8px] font-black rounded uppercase">New Folder</span>
                      </div>
                    </div>
 
                    <!-- Document Item -->
                    <div v-else-if="item.isDocument" class="pl-2 border-l-2 border-slate-200/80">
                      <div class="bg-slate-50 border border-slate-100 rounded-xl p-3 flex items-center gap-3 max-w-sm ml-2">
                        <div class="w-8 h-8 bg-white rounded-lg shadow-sm flex items-center justify-center text-[#0ea5e9]">
                          <LucideFileText class="w-5 h-5" />
                        </div>
                        <div>
                          <p class="text-[11px] font-black text-slate-800 uppercase tracking-tight">{{ item.name }}</p>
                          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">{{ item.tag }}</p>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
 
          <!-- Bottom Actions -->
          <div class="flex items-center justify-end pt-4">
            <button @click="saveConfig" :disabled="isSaving" class="flex items-center gap-3 px-6 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl shadow-lg shadow-blue-900/10 transition-all font-black uppercase tracking-[0.15em] text-xs group">
              <span v-if="!isSaving">Commit Folder Placement</span>
              <span v-else>Saving...</span>
              <LucideCheckCircle2 class="w-4 h-4 group-hover:scale-110 transition-transform" />
            </button>
          </div>
        </div>
 
        <!-- Settings Sidebar -->
        <div class="col-span-4 space-y-6">
          <!-- Foldering Rules Section -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 text-slate-800">
              <LucideSettings2 class="w-4 h-4 text-[#1E3A5F]" />
              <h2 class="text-xs font-black uppercase tracking-widest text-[#1E3A5F]">Foldering Rules</h2>
            </div>
 
            <!-- Policy Banner -->
            <div class="bg-white p-5 rounded-2xl border border-slate-200/60 shadow-sm space-y-3">
               <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Rule Source Policy</p>
               <div class="flex items-start gap-3 p-3 bg-green-50/50 border border-green-100/80 rounded-xl">
                 <LucideCheckCircle2 class="w-4 h-4 text-green-500 mt-0.5" />
                 <div>
                   <p class="text-xs font-black text-green-700 uppercase">Dept Policy: FINANCE</p>
                   <p class="text-[11px] text-green-600/80 font-medium leading-relaxed mt-0.5">System automatically maps folder targets based on Financial Compliance (ISO-27001).</p>
                 </div>
               </div>
            </div>
 
            <!-- Path Formula Builder -->
            <div class="bg-white p-6 rounded-3xl border border-slate-200/60 shadow-sm space-y-4">
               <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Path Formula Preview</p>
               
               <div class="bg-[#1E3A5F] p-4 rounded-xl flex flex-wrap gap-1.5 min-h-[80px] content-start">
                 <span class="px-2.5 py-1 bg-blue-900/30 text-blue-200 text-[10px] font-black rounded border border-blue-200/10">/ARSIP/</span>
                 
                 <div 
                   v-for="(level, index) in levels" 
                   :key="level.key"
                   class="flex items-center"
                 >
                    <div class="px-2.5 py-1 bg-[#0ea5e9] text-white text-[10px] font-black rounded shadow flex items-center gap-1.5 group cursor-pointer relative">
                      ${{ level.label }}
                      <div class="hidden group-hover:flex absolute -top-8 left-0 gap-1">
                         <button @click="moveLevel(index, -1)" class="p-1 bg-white text-[#1E3A5F] rounded shadow hover:bg-slate-100"><LucideChevronLeft class="w-3 h-3"/></button>
                         <button @click="moveLevel(index, 1)" class="p-1 bg-white text-[#1E3A5F] rounded shadow hover:bg-slate-100"><LucideChevronRight class="w-3 h-3"/></button>
                      </div>
                    </div>
                    <span v-if="index < levels.length - 1" class="text-blue-200/30 px-0.5">/</span>
                 </div>
               </div>
 
               <!-- Add Level Dropdown (Simple version) -->
               <div class="space-y-3">
                 <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Configuration Settings</p>
                 
                 <!-- Create Missing Folders Toggle -->
                 <div class="flex items-center justify-between p-3 bg-slate-50 rounded-xl border border-slate-200/60">
                    <div>
                      <p class="text-xs font-black text-slate-800 uppercase tracking-tight">Create if not exists</p>
                      <p class="text-[10px] text-slate-400 font-medium">Auto-generate missing directories</p>
                    </div>
                    <button 
                      @click="config.create_missing_folders = !config.create_missing_folders"
                      class="w-10 h-5 rounded-full transition-all relative"
                      :class="config.create_missing_folders ? 'bg-[#0ea5e9]' : 'bg-slate-200'"
                    >
                      <div class="absolute top-0.5 w-4 h-4 bg-white rounded-full transition-all" :class="config.create_missing_folders ? 'left-5.5' : 'left-0.5'"></div>
                    </button>
                 </div>
 
                 <!-- Apply Retention Toggle -->
                 <div class="flex items-center justify-between p-3 bg-slate-50 rounded-xl border border-slate-200/60">
                    <div>
                      <p class="text-xs font-black text-slate-800 uppercase tracking-tight">Apply Retention Tag</p>
                      <p class="text-[10px] text-slate-400 font-medium">Add 7-year purge metadata</p>
                    </div>
                    <button 
                      @click="config.apply_retention_tag = !config.apply_retention_tag"
                      class="w-10 h-5 rounded-full transition-all relative"
                      :class="config.apply_retention_tag ? 'bg-[#0ea5e9]' : 'bg-slate-200'"
                    >
                      <div class="absolute top-0.5 w-4 h-4 bg-white rounded-full transition-all" :class="config.apply_retention_tag ? 'left-5.5' : 'left-0.5'"></div>
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
import { useAuthStore } from '~/stores/auth'
import { formatBytes } from '~/utils/format'

const router = useRouter()
const { $api } = useApi()
const toast = useToast()
const auth = useAuthStore()

const isSaving = ref(false)
const isRefreshing = ref(false)
const storageUsed = ref(0)

const storageLimitBytes = computed(() => {
  const role = (auth.user?.role || '').toLowerCase()
  if (role.includes('controller') || role.includes('dc')) {
    return 25 * 1024 * 1024 * 1024 * 1024 // 25 TB
  }
  return 10 * 1024 * 1024 * 1024 * 1024 // 10 TB
})

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

const mockDataMap = {
  company: { name: 'KREATIF HOLDING', tag: '{PT}' },
  branch: { name: 'KREATIF JAKARTA', tag: '{CABANG}' },
  department: { name: 'IT_INFRASTRUCTURE', tag: '{DEPT}' },
  year: { name: '2026', tag: '{TAHUN}' },
  rack: { name: 'RCK-A1-204-B', tag: '{RAK}' },
  box: { name: 'BOX-IT-2026-001', tag: '{BOX}' },
  ordner: { name: 'ORDNER-SERVER-MAINT', tag: '{ORDNER}' },
  type: { name: 'INVOICE_BILLING', tag: '{TIPE}' }
}

const flatPreviewItems = computed(() => {
  const items = []
  levels.value.forEach((level, index) => {
    const key = level.key
    const mockInfo = mockDataMap[key] || { name: key.toUpperCase(), tag: `{${key.toUpperCase()}}` }
    items.push({
      id: key,
      name: mockInfo.name,
      tag: mockInfo.tag,
      indent: index,
      isNew: index === levels.value.length - 1, // mark last level as new folder
      isFolder: true
    })
  })
  
  items.push({
    id: 'document',
    name: 'INV-2026-00192_KREATIF.pdf',
    tag: '2.4 MB • OCR Verified',
    indent: levels.value.length,
    isDocument: true
  })
  
  return items
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
      if (res.data.config) {
        config.value = { ...config.value, ...res.data.config }
      } else {
        config.value = { ...config.value, ...res.data }
      }
      storageUsed.value = res.data.storage_used || 0
      
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
