<template>
  <div class="max-w-7xl mx-auto space-y-10 pb-20" v-motion-fade>
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Staging Area</h1>
        <p class="text-slate-500 font-medium italic">Antrean dokumen yang siap untuk proses digitalisasi & scanning</p>
      </div>
      <div class="flex items-center gap-4">
        <div class="flex -space-x-3">
          <div v-for="(dc, i) in (stats?.active_dc_list || [])" :key="i" class="w-10 h-10 rounded-full border-4 border-white dark:border-slate-900 bg-slate-200 flex items-center justify-center overflow-hidden">
            <img v-if="dc.avatar" :src="dc.avatar" class="w-full h-full object-cover" />
            <div v-else class="text-[10px] font-black text-slate-400 uppercase">{{ dc.name?.substring(0, 2) }}</div>
          </div>
        </div>
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-none">
          {{ stats?.active_dc_count || 0 }} Document<br/>Controllers Aktif
        </p>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div v-for="(stat, key) in statCards" :key="key" class="glass rounded-[32px] p-8 flex flex-col gap-6 group hover:-translate-y-1 transition-all duration-300">
        <div class="flex items-center justify-between">
          <div :class="['w-14 h-14 rounded-2xl flex items-center justify-center shadow-sm group-hover:text-white transition-all duration-500', stat.colorClass]">
            <component :is="stat.icon" class="w-7 h-7" />
          </div>
          <div class="text-right">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ stat.label }}</p>
            <div class="flex items-baseline justify-end gap-2">
              <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stat.value }}</p>
              <span v-if="stat.subValue" :class="['text-[10px] font-black uppercase tracking-widest', stat.subColor]">{{ stat.subValue }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Table Card / Empty State -->
    <div v-if="!pending && manifests.length === 0" class="glass rounded-[32px] p-20 flex flex-col items-center justify-center space-y-10 border border-slate-100 dark:border-slate-800 shadow-2xl shadow-slate-200/20" v-motion-fade>
      <div class="w-32 h-32 bg-slate-50 dark:bg-slate-900 rounded-full flex items-center justify-center text-slate-200 shadow-inner">
        <LucidePackage class="w-16 h-16" />
      </div>
      <div class="text-center space-y-4 max-w-lg">
        <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Tidak Ada Dokumen dalam Antrean</h3>
        <p class="text-sm font-medium text-slate-400 leading-relaxed italic">
          Dokumen akan masuk ke sini setelah Admin melakukan penerimaan fisik di menu <span class="text-orange-500 font-bold">Inbound Pre-Registration</span>
        </p>
      </div>
      <div class="flex items-center gap-4">
        <button 
          @click="navigateTo('/intake/inbound')"
          class="px-8 py-4 border-2 border-orange-500/20 hover:border-orange-500/40 text-orange-600 rounded-2xl text-xs font-black uppercase tracking-widest transition-all flex items-center gap-3 group"
        >
          <LucideArrowRight class="w-4 h-4 group-hover:translate-x-1 transition-transform" />
          Ke Inbound Pre-Registration
        </button>
        <button 
          @click="refresh"
          class="px-8 py-4 bg-slate-100 dark:bg-slate-800 text-slate-400 hover:text-slate-600 rounded-2xl text-xs font-black uppercase tracking-widest transition-all"
        >
          Refresh Halaman
        </button>
      </div>
    </div>

    <div v-else class="space-y-6">
      <!-- Tabs & Search -->
      <div class="flex flex-col md:flex-row items-center justify-between gap-6">
        <div class="flex items-center p-1.5 bg-slate-100/50 dark:bg-slate-900/50 rounded-2xl w-full md:w-auto">
          <button 
            v-for="tab in tabs" 
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'px-8 py-3 rounded-xl text-xs font-black uppercase tracking-widest transition-all',
              activeTab === tab.id ? 'bg-white dark:bg-slate-800 text-primary-500 shadow-sm' : 'text-slate-400 hover:text-slate-600'
            ]"
          >
            {{ tab.label }} <span class="ml-2 opacity-50">({{ tab.count }})</span>
          </button>
        </div>

        <div class="flex items-center gap-4 w-full md:w-auto">
           <div class="relative flex-1 md:w-64">
             <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300" />
             <input 
              v-model="search"
              type="text"
              placeholder="Cari Manifest ID..."
              class="w-full pl-12 pr-4 py-3 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold text-slate-600 dark:text-slate-300 focus:ring-2 focus:ring-primary-500/10 outline-none transition-all"
             />
           </div>
           <button class="p-3 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-slate-400 hover:text-primary-500 transition-all">
             <LucideFilter class="w-5 h-5" />
           </button>
        </div>
      </div>

      <!-- Main Table Card -->
      <div class="glass rounded-[32px] overflow-hidden border border-slate-100 dark:border-slate-800 shadow-2xl shadow-slate-200/20">
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-900/50 border-b border-slate-100 dark:border-slate-800 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <th class="px-8 py-6 w-16">
                  <input 
                    type="checkbox" 
                    :checked="isAllSelected"
                    @change="toggleSelectAll"
                    class="w-4 h-4 rounded border-slate-300 text-primary-500 focus:ring-primary-500/20" 
                  />
                </th>
                <th class="px-8 py-6 w-16 text-center">No</th>
                <th class="px-8 py-6">Manifest ID</th>
                <th class="px-8 py-6 text-center">Sumber</th>
                <th class="px-8 py-6">Pemohon / Pengaju</th>
                <th class="px-8 py-6">Departemen</th>
                <th class="px-8 py-6 text-center">Tgl Terima</th>
                <th class="px-8 py-6">Waktu Tunggu</th>
                <th class="px-8 py-6 text-center">Dokumen</th>
                <th class="px-8 py-6 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
              <tr v-if="pending" class="animate-pulse">
                <td colspan="10" class="px-8 py-20 text-center text-slate-400 font-bold uppercase tracking-widest">Loading data...</td>
              </tr>
              <tr v-else-if="filteredManifests.length === 0" class="text-center">
                <td colspan="10" class="px-8 py-20 text-slate-400 font-bold uppercase tracking-widest italic opacity-50">Tidak ada hasil pencarian</td>
              </tr>
              <tr 
                v-for="(m, index) in paginatedManifests" 
                :key="m.id" 
                class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-all cursor-default"
                :class="{'bg-red-50/30 dark:bg-red-900/10': isOverdue(m.received_at)}"
              >
                <td class="px-8 py-6">
                  <input 
                    type="checkbox" 
                    :value="m.id"
                    v-model="selectedIds"
                    class="w-4 h-4 rounded border-slate-300 text-primary-500 focus:ring-primary-500/20" 
                  />
                </td>
                <td class="px-8 py-6 text-center text-xs font-black text-slate-300">
                  {{ ((currentPage - 1) * itemsPerPage) + index + 1 }}
                </td>
                <td class="px-8 py-6">
                  <p class="font-black text-sm uppercase tracking-tight text-[#1E3A5F] dark:text-slate-200 group-hover:text-primary-500 transition-colors">
                    {{ m.manifest_no }}
                  </p>
                </td>
                <td class="px-8 py-6 text-center">
                  <span 
                    :class="[
                      'px-3 py-1 rounded-md text-[9px] font-black uppercase tracking-widest border shadow-sm',
                      m.source === 'MIG' ? 'bg-amber-50 text-amber-600 border-amber-100' : 'bg-blue-50 text-blue-600 border-blue-100'
                    ]"
                  >
                    {{ m.source }}
                  </span>
                </td>
                <td class="px-8 py-6">
                  <p class="font-black text-sm text-[#1E3A5F] dark:text-slate-300">{{ m.owner_name }}</p>
                </td>
                <td class="px-8 py-6">
                  <p class="font-bold text-xs text-slate-500">{{ m.department_name }}</p>
                </td>
                <td class="px-8 py-6 text-center text-xs font-bold text-slate-500">
                  {{ formatDate(m.received_at) }}
                </td>
                <td class="px-8 py-6">
                  <div class="flex items-center gap-2">
                    <div :class="['w-2 h-2 rounded-full', getWaitTimeColor(m.received_at)]"></div>
                    <p :class="['text-xs font-black uppercase tracking-tight', getWaitTimeTextColor(m.received_at)]">
                      {{ getWaitTime(m.received_at) }}
                      <span v-if="isOverdue(m.received_at)" class="ml-1 text-[8px] font-black text-red-500">TERLAMBAT</span>
                    </p>
                  </div>
                </td>
                <td class="px-8 py-6 text-center">
                  <p class="text-lg font-black text-[#1E3A5F] dark:text-white">{{ m.total_items || 0 }}</p>
                </td>
                <td class="px-8 py-6 text-right">
                  <button 
                    @click="startDigitalization(m)"
                    class="px-6 py-2.5 bg-orange-500 hover:bg-orange-600 text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-lg shadow-orange-500/20 transition-all hover:-translate-y-0.5 active:translate-y-0"
                  >
                    Digitalisasi
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Footer Pagination -->
        <div class="bg-slate-50/50 dark:bg-slate-900/50 p-6 flex items-center justify-between border-t border-slate-100 dark:border-slate-800">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            Menampilkan {{ ((currentPage - 1) * itemsPerPage) + 1 }}-{{ Math.min(currentPage * itemsPerPage, filteredManifests.length) }} dari {{ filteredManifests.length }} antrean
          </p>
          <div class="flex items-center gap-2">
            <button 
              @click="currentPage > 1 && (currentPage--)"
              :disabled="currentPage === 1"
              class="w-8 h-8 rounded-lg flex items-center justify-center text-xs font-black bg-white dark:bg-slate-800 text-slate-400 hover:text-primary-500 disabled:opacity-30 disabled:hover:text-slate-400 transition-all"
            >
              <LucideChevronLeft class="w-4 h-4" />
            </button>
            <button 
              v-for="p in totalPages" 
              :key="p" 
              @click="currentPage = p"
              :class="[
                'w-8 h-8 rounded-lg flex items-center justify-center text-xs font-black transition-all', 
                currentPage === p ? 'bg-primary-500 text-white shadow-lg shadow-primary-500/20' : 'bg-white dark:bg-slate-800 text-slate-400 hover:text-primary-500'
              ]"
            >
              {{ p }}
            </button>
            <button 
              @click="currentPage < totalPages && (currentPage++)"
              :disabled="currentPage === totalPages || totalPages === 0"
              class="w-8 h-8 rounded-lg flex items-center justify-center text-xs font-black bg-white dark:bg-slate-800 text-slate-400 hover:text-primary-500 disabled:opacity-30 disabled:hover:text-slate-400 transition-all"
            >
              <LucideChevronRight class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Bulk Action Bar (Floating) -->
    <Transition name="slide-up">
      <div v-if="selectedIds.length > 0" class="fixed bottom-10 left-1/2 -translate-x-1/2 z-50 w-full max-w-4xl px-6">
        <div class="bg-[#1E3A5F] dark:bg-slate-900 shadow-2xl shadow-blue-900/40 rounded-2xl p-4 flex items-center justify-between border border-white/10 backdrop-blur-xl">
          <div class="flex items-center gap-6 pl-4">
            <div class="w-10 h-10 bg-emerald-500/20 rounded-xl flex items-center justify-center text-emerald-500 shadow-inner">
              <LucideLayers class="w-5 h-5" />
            </div>
            <div class="space-y-0.5">
              <p class="text-xs font-black text-white uppercase tracking-widest">{{ selectedIds.length }} Manifest dipilih</p>
              <p class="text-[10px] font-bold text-slate-400 italic">Siap untuk diproses secara kolektif</p>
            </div>
          </div>

          <div class="flex items-center gap-4">
            <button 
              @click="selectedIds = []"
              class="px-6 py-3 text-xs font-black text-slate-400 hover:text-white uppercase tracking-widest transition-all"
            >
              Batal
            </button>
            <button 
              @click="startBatchDigitalization"
              class="px-8 py-3 bg-emerald-500 hover:bg-emerald-600 text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-lg shadow-emerald-500/20 transition-all flex items-center gap-3 active:scale-95"
            >
              <LucideScanText class="w-4 h-4" />
              Mulai Batch Digitalisasi
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { 
  LucideSearch, LucideFilter, LucideHourglass, LucideClock, 
  LucideCheckCircle, LucideAlertTriangle, LucideArrowUpRight,
  LucidePackage, LucideArrowRight, LucideChevronLeft, LucideChevronRight,
  LucideLayers, LucideScanText
} from 'lucide-vue-next'

const { $api } = useApi()
const activeTab = ref('all')
const search = ref('')
const currentPage = ref(1)
const itemsPerPage = ref(10)
const selectedIds = ref([])

const tabs = computed(() => [
  { id: 'all', label: 'Semua', count: manifests.value?.length || 0 },
  { id: 'inbound', label: 'Inbound', count: manifests.value?.filter(m => m.source === 'INB').length || 0 },
  { id: 'legacy', label: 'Migrasi Legacy', count: manifests.value?.filter(m => m.source === 'MIG').length || 0 },
])

const { data: manifests, pending, refresh } = useAsyncData('intake-staging', async () => {
  const res = await $api('/intake/staging')
  return res.data || []
}, {
  default: () => []
})

const { data: stats } = useAsyncData('intake-staging-stats', async () => {
  const res = await $api('/intake/staging/stats')
  return res.data || {}
})

const statCards = computed(() => ({
  active: {
    label: 'Antrean Saat Ini',
    value: stats.value?.active_count || 0,
    subValue: 'Aktif',
    subColor: 'text-emerald-500',
    icon: LucideHourglass,
    colorClass: 'bg-emerald-50 text-emerald-500 group-hover:bg-emerald-500'
  },
  wait: {
    label: 'Rata-rata Waktu Tunggu',
    value: stats.value?.avg_wait_time || '1h 24m',
    icon: LucideClock,
    colorClass: 'bg-amber-50 text-amber-500 group-hover:bg-amber-500'
  },
  done: {
    label: 'Selesai Hari Ini',
    value: stats.value?.completed_today || 0,
    subValue: '+10%',
    subColor: 'text-emerald-500',
    icon: LucideCheckCircle,
    colorClass: 'bg-emerald-50 text-emerald-500 group-hover:bg-emerald-500'
  },
  late: {
    label: 'Terlambat',
    value: stats.value?.overdue_count || 0,
    subValue: 'Kritikal',
    subColor: 'text-red-500',
    icon: LucideAlertTriangle,
    colorClass: 'bg-red-50 text-red-500 group-hover:bg-red-500'
  }
}))

const filteredManifests = computed(() => {
  let list = manifests.value || []
  
  if (activeTab.value === 'inbound') list = list.filter(m => m.source === 'INB')
  if (activeTab.value === 'legacy') list = list.filter(m => m.source === 'MIG')
  
  if (search.value) {
    const q = search.value.toLowerCase()
    list = list.filter(m => 
      m.manifest_no.toLowerCase().includes(q) || 
      m.owner_name.toLowerCase().includes(q) ||
      m.department_name.toLowerCase().includes(q)
    )
  }
  
  return list
})

const totalPages = computed(() => Math.ceil(filteredManifests.value.length / itemsPerPage.value))

const paginatedManifests = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredManifests.value.slice(start, end)
})

// Reset page and selection on search or tab change
watch([search, activeTab], () => {
  currentPage.value = 1
  selectedIds.value = []
})

// Selection logic
const isAllSelected = computed(() => {
  return paginatedManifests.value.length > 0 && paginatedManifests.value.every(m => selectedIds.value.includes(m.id))
})

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    const idsToRemove = paginatedManifests.value.map(m => m.id)
    selectedIds.value = selectedIds.value.filter(id => !idsToRemove.includes(id))
  } else {
    const newIds = paginatedManifests.value.map(m => m.id).filter(id => !selectedIds.value.includes(id))
    selectedIds.value = [...selectedIds.value, ...newIds]
  }
}

const toggleSelect = (id) => {
  const index = selectedIds.value.indexOf(id)
  if (index > -1) {
    selectedIds.value.splice(index, 1)
  } else {
    selectedIds.value.push(id)
  }
}

const startBatchDigitalization = () => {
  if (selectedIds.value.length === 0) return
  navigateTo({
    path: '/intake/batch',
    query: { ids: selectedIds.value.join(',') }
  })
}

// Formatting Helpers
const formatDate = (dateStr) => {
  if (!dateStr) return '-'
  return dateStr.split(' ')[0] // Just the date part
}

const getWaitTime = (receivedAt) => {
  if (!receivedAt) return '-'
  
  // Format: DD/MM/YYYY HH:mm
  const parts = receivedAt.split(' ')
  if (parts.length < 2) return receivedAt
  
  const dateParts = parts[0].split('/')
  const timeParts = parts[1].split(':')
  
  const date = new Date(
    parseInt(dateParts[2]), 
    parseInt(dateParts[1]) - 1, 
    parseInt(dateParts[0]), 
    parseInt(timeParts[0]), 
    parseInt(timeParts[1])
  )
  
  const diffMs = new Date() - date
  if (diffMs < 0) return '0m'
  
  const diffDays = Math.floor(diffMs / 86400000)
  const diffHrs = Math.floor((diffMs % 86400000) / 3600000)
  const diffMins = Math.floor((diffMs % 3600000) / 60000)
  
  let result = ''
  if (diffDays > 0) result += `${diffDays}d `
  if (diffHrs > 0 || diffDays > 0) result += `${diffHrs}h `
  result += `${diffMins}m`
  
  return result
}

const isOverdue = (receivedAt) => {
  if (!receivedAt) return false
  const parts = receivedAt.split(' ')
  if (parts.length < 2) return false
  
  const dateParts = parts[0].split('/')
  const timeParts = parts[1].split(':')
  const date = new Date(
    parseInt(dateParts[2]), 
    parseInt(dateParts[1]) - 1, 
    parseInt(dateParts[0]), 
    parseInt(timeParts[0]), 
    parseInt(timeParts[1])
  )
  const diffMs = new Date() - date
  return diffMs > (4 * 3600000) // 4 hours threshold
}

const getWaitTimeColor = (receivedAt) => {
  if (isOverdue(receivedAt)) return 'bg-red-500 animate-pulse'
  return 'bg-emerald-500'
}

const getWaitTimeTextColor = (receivedAt) => {
  if (isOverdue(receivedAt)) return 'text-red-500'
  return 'text-emerald-500'
}

const startDigitalization = (manifest) => {
  navigateTo({
    path: '/intake/indexing',
    query: { manifest_no: manifest.manifest_no }
  })
}
</script>

<style scoped>
.glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(40px);
  -webkit-backdrop-filter: blur(40px);
  border: 1px solid rgba(255, 255, 255, 0.8);
}

.dark .glass {
  background: rgba(13, 18, 31, 0.7);
  border-color: rgba(30, 41, 59, 0.5);
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.5s cubic-bezier(0.16, 1, 0.3, 1);
}

.slide-up-enter-from,
.slide-up-leave-to {
  transform: translate(-50%, 100px);
  opacity: 0;
}
</style>
