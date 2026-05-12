<template>
  <div class="space-y-8 pb-20" v-motion-fade>
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">Print QR Label</h1>
        <p class="text-xs font-bold text-slate-500 italic">Cetak label identitas fisik untuk dokumen yang telah didigitalisasi</p>
      </div>
      
      <div class="flex items-center gap-4">
        <div class="relative w-full md:w-80">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input 
            type="text" 
            v-model="searchQuery"
            placeholder="Cari System ID atau Title..." 
            class="w-full pl-11 pr-5 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all shadow-sm" 
          />
        </div>
      </div>
    </div>

    <!-- Tabs & Status -->
    <div class="flex items-center justify-between border-b border-slate-200 dark:border-slate-800 pb-px">
      <div class="flex gap-10">
        <button 
          v-for="tab in tabs" 
          :key="tab.id"
          @click="activeTab = tab.id"
          class="pb-4 text-xs font-black uppercase tracking-widest transition-all relative"
          :class="activeTab === tab.id ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400 hover:text-slate-600'"
        >
          {{ tab.name }}
          <span v-if="tab.count > 0" class="ml-2 px-2 py-0.5 bg-orange-100 text-orange-600 rounded-full text-[9px]">{{ tab.count }}</span>
          <div v-if="activeTab === tab.id" class="absolute bottom-0 left-0 right-0 h-1 bg-orange-500 rounded-t-full" v-motion-pop></div>
        </button>
      </div>

      <div v-if="selectedItems.length > 0" class="pb-4 flex items-center gap-4" v-motion-slide-right>
        <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ selectedItems.length }} item terpilih</span>
        <button 
          @click="handleBulkPrint"
          class="px-6 py-2 bg-[#1E3A5F] text-white rounded-lg text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#2A4B7C] transition-all flex items-center gap-2"
        >
          <LucidePrinter class="w-3.5 h-3.5" />
          {{ activeTab === 'digitized' ? 'Cetak Label' : 'Cetak Ulang' }}
        </button>
      </div>
    </div>

    <!-- Table Content -->
    <div class="glass-container rounded-3xl overflow-hidden shadow-2xl shadow-slate-200/50 dark:shadow-none border border-white dark:border-slate-800">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
              <th class="p-6 pl-8 w-12">
                <input 
                  type="checkbox" 
                  :checked="isAllSelected" 
                  @change="toggleSelectAll"
                  class="w-4 h-4 rounded border-slate-300 text-primary-600 focus:ring-primary-500"
                />
              </th>
              <th class="p-6">System ID</th>
              <th class="p-6">Title</th>
              <th class="p-6">Department</th>
              <th class="p-6 text-right pr-8">Manifest Ref</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50 bg-white dark:bg-[#0D121F]">
            <tr v-if="loading" v-for="i in 5" :key="i" class="animate-pulse">
              <td colspan="5" class="p-8 px-10"><div class="h-4 bg-slate-100 dark:bg-slate-800 rounded-full w-full"></div></td>
            </tr>
            <tr v-else-if="filteredDocs.length === 0" class="text-center py-20">
              <td colspan="5" class="p-20">
                <div class="flex flex-col items-center gap-4 text-slate-400">
                  <LucideInbox class="w-12 h-12 opacity-20" />
                  <p class="text-xs font-bold italic tracking-wide uppercase">Tidak ada dokumen di antrian ini</p>
                </div>
              </td>
            </tr>
            <tr 
              v-for="doc in filteredDocs" 
              :key="doc.document_id"
              class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all cursor-pointer"
              @click="toggleSelectItem(doc.document_id)"
            >
              <td class="p-6 pl-8" @click.stop>
                <input 
                  type="checkbox" 
                  :checked="selectedItems.includes(doc.document_id)" 
                  @change="toggleSelectItem(doc.document_id)"
                  class="w-4 h-4 rounded border-slate-300 text-primary-600 focus:ring-primary-500"
                />
              </td>
              <td class="p-6">
                <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight group-hover:text-primary-500 transition-colors">
                  SYS-{{ doc.document_id.substring(0, 8).toUpperCase() }}
                </span>
              </td>
              <td class="p-6">
                <p class="text-xs font-bold text-slate-600 dark:text-slate-300">{{ doc.title }}</p>
              </td>
              <td class="p-6">
                <span class="px-3 py-1 bg-blue-50 dark:bg-blue-900/20 text-blue-500 dark:text-blue-400 rounded-md text-[9px] font-black uppercase tracking-widest border border-blue-100 dark:border-blue-800/50">
                  {{ doc.department_name }}
                </span>
              </td>
              <td class="p-6 text-right pr-8">
                <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ doc.manifest_no }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Print Progress Modal -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100" leave-active-class="transition duration-200 ease-in" leave-from-class="opacity-100" leave-to-class="opacity-0">
      <div v-if="printing" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/60 backdrop-blur-md">
        <div class="bg-white dark:bg-slate-900 w-full max-w-sm rounded-3xl shadow-2xl p-10 flex flex-col items-center text-center">
          <LucidePrinter class="w-12 h-12 text-[#1E3A5F] dark:text-white animate-bounce mb-6" />
          <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight mb-2">Sedang Mencetak...</h3>
          <p class="text-xs text-slate-500 font-medium mb-8">Mohon tunggu, label sedang dikirim ke printer thermal</p>
          <div class="w-full h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
            <div class="h-full bg-orange-500 animate-progress"></div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { 
  LucideSearch, LucidePrinter, LucideInbox, LucideChevronDown 
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const loading = ref(true)
const printing = ref(false)
const activeTab = ref('digitized')
const searchQuery = ref('')
const selectedItems = ref([])
const documents = ref([])
const stats = ref({
  waiting_count: 0,
  printed_count: 0,
  stored_count: 0
})

const tabs = computed(() => [
  { id: 'digitized', name: 'Menunggu Label', count: stats.value.waiting_count },
  { id: 'labeled', name: 'Sudah Cetak', count: stats.value.printed_count },
  { id: 'archived', name: 'Sudah Tersimpan', count: stats.value.stored_count }
])

const fetchData = async () => {
  loading.value = true
  try {
    const [docsRes, statsRes] = await Promise.all([
      $api(`/intake/labels?status=${activeTab.value}`),
      $api('/intake/labels/stats')
    ])
    documents.value = docsRes.data || []
    stats.value = statsRes.data || { waiting_count: 0, printed_count: 0, stored_count: 0 }
  } catch (err) {
    console.error('Failed to fetch data:', err)
  } finally {
    loading.value = false
  }
}

const filteredDocs = computed(() => {
  if (!searchQuery.value) return documents.value
  const q = searchQuery.value.toLowerCase()
  return documents.value.filter(d => 
    d.title.toLowerCase().includes(q) || 
    d.document_id.toLowerCase().includes(q)
  )
})

const isAllSelected = computed(() => {
  return filteredDocs.value.length > 0 && selectedItems.value.length === filteredDocs.value.length
})

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedItems.value = []
  } else {
    selectedItems.value = filteredDocs.value.map(d => d.document_id)
  }
}

const toggleSelectItem = (id) => {
  if (selectedItems.value.includes(id)) {
    selectedItems.value = selectedItems.value.filter(i => i !== id)
  } else {
    selectedItems.value.push(id)
  }
}

const handleBulkPrint = () => {
  if (selectedItems.value.length === 0) return
  
  // Navigate to assignment page with selected IDs
  navigateTo({
    path: '/warehouse/labels-print-assignment',
    query: {
      ids: selectedItems.value.join(',')
    }
  })
}

watch(activeTab, () => {
  selectedItems.value = []
  fetchData()
})

onMounted(fetchData)
</script>

<style scoped>
.glass-container {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}
.animate-progress {
  width: 0;
  animation: progress 2s ease-in-out infinite;
}
@keyframes progress {
  0% { width: 0; }
  100% { width: 100%; }
}
</style>
