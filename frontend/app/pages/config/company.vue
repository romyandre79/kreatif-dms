<template>
  <div class="flex flex-col h-full bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <!-- Hidden File Input for Import -->
    <input type="file" ref="fileInput" class="hidden" accept=".csv" @change="handleFileImport" />

    <div class="flex flex-1 overflow-hidden">
      <!-- Main Content Area -->
      <main class="flex-grow flex flex-col overflow-hidden">
        <!-- Header -->
        <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
          <div class="space-y-1">
            <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
              {{ activeEntityLabel }}
            </h1>
            <p class="text-sm font-bold text-slate-500 uppercase tracking-tighter">Pengaturan entitas master dan struktur organisasi sistem.</p>
          </div>
          <div class="flex items-center gap-4">
            <button @click="triggerImport" class="px-6 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2 shadow-sm cursor-pointer">
              <LucideUpload class="w-4 h-4" />
              Import CSV
            </button>
            <button @click="addRow" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95 cursor-pointer">
              <LucidePlus class="w-4 h-4" />
              Add Row
            </button>
          </div>
        </header>

        <!-- Table Section -->
        <div class="flex-grow p-2.5 overflow-auto custom-scrollbar">
          <div class="bg-white dark:bg-slate-900 rounded-[1.5rem] shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden">
            <!-- Filter Bar -->
            <div class="px-8 py-6 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
              <div class="relative w-96 group">
                <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
                <input 
                  type="text" 
                  v-model="searchQuery"
                  placeholder="Filter data..." 
                  class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl pl-12 pr-6 py-3 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20 transition-all" 
                />
              </div>
              <div class="flex items-center gap-4">
                <button @click="fetchCompanies" class="p-3 text-slate-400 hover:text-blue-500 transition-colors cursor-pointer" title="Refresh"><LucideRefreshCw class="w-5 h-5" /></button>
                <button @click="exportCSV" class="p-3 text-slate-400 hover:text-blue-500 transition-colors cursor-pointer" title="Export CSV"><LucideDownload class="w-5 h-5" /></button>
              </div>
            </div>

            <!-- Loading State -->
            <div v-if="loading" class="p-20 flex flex-col items-center justify-center space-y-4">
               <LucideRefreshCw class="w-10 h-10 text-blue-500 animate-spin" />
               <p class="text-xs font-black text-slate-400 uppercase tracking-widest">Loading data...</p>
            </div>

            <!-- Table -->
            <div v-else class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                    <th class="p-4 pl-8">ID / Code</th>
                    <th class="p-4">Nama / Deskripsi</th>
                    <th class="p-4" v-if="activeEntity === 'pt'">NPWP Status</th>
                    <th class="p-4">Lokasi / Detail</th>
                    <th class="p-4">Status</th>
                    <th class="p-4 pr-10 text-right">Action</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                  <tr v-if="filteredCompanies.length === 0" class="hover:bg-transparent">
                     <td colspan="6" class="p-20 text-center">
                        <div class="flex flex-col items-center justify-center space-y-3 opacity-30">
                           <LucideArchive class="w-12 h-12 text-slate-300" />
                           <p class="text-xs font-black uppercase tracking-[0.2em] text-slate-400">No Data Found</p>
                        </div>
                     </td>
                  </tr>
                  <tr v-else v-for="row in paginatedCompanies" :key="row.id" 
                    class="group transition-all"
                    :class="row.isNew ? 'bg-blue-50/50 dark:bg-blue-900/10' : 'hover:bg-slate-50/50 dark:hover:bg-slate-800/30'"
                  >
                    <td class="p-4 pl-8">
                      <div v-if="row.editing" class="max-w-[120px]">
                        <input 
                          type="text" 
                          v-model="row.entity_id" 
                          class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-2 text-xs font-black text-blue-500 font-mono outline-none shadow-md" 
                        />
                      </div>
                      <span v-else class="text-sm font-black text-blue-500 font-mono tracking-tighter">{{ row.entity_id }}</span>
                    </td>
                    <td class="p-4">
                      <div v-if="row.editing" class="relative group max-w-md">
                        <input 
                          type="text" 
                          v-model="row.name" 
                          class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-3 text-sm font-black text-[#1E3A5F] dark:text-white uppercase outline-none shadow-lg shadow-blue-900/10" 
                          @keyup.enter="saveRow(row)"
                        />
                      </div>
                      <p v-else class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ row.name === 'null' ? '-' : (row.name || '-') }}</p>
                    </td>
                    <td class="p-4" v-if="activeEntity === 'pt'">
                      <div v-if="row.editing">
                         <select v-model="row.npwp_status" class="bg-white dark:bg-slate-800 border border-slate-200 rounded-lg px-2 py-1 text-[10px] font-black uppercase">
                           <option value="PENDING">PENDING</option>
                           <option value="VERIFIED">VERIFIED</option>
                         </select>
                      </div>
                      <span v-else :class="`px-3 py-1.5 rounded-lg text-[9px] font-black uppercase tracking-widest border shadow-sm ${row.npwp_status === 'VERIFIED' ? 'bg-green-50 text-green-500 border-green-100' : 'bg-red-50 text-red-500 border-red-100'}`">
                        {{ row.npwp_status === 'null' ? 'PENDING' : (row.npwp_status || 'PENDING') }}
                      </span>
                    </td>
                    <td class="p-4">
                      <div v-if="row.editing">
                        <input type="text" v-model="row.location" class="w-full bg-white dark:bg-slate-800 border border-slate-200 rounded-xl px-4 py-2 text-xs font-bold" />
                      </div>
                      <p v-else class="text-sm font-bold text-slate-500 uppercase tracking-tight">{{ (row.location === 'null' || !row.location) ? '-' : row.location }}</p>
                    </td>
                    <td class="p-4">
                      <div v-if="row.editing">
                         <select v-model="row.status" class="bg-white dark:bg-slate-800 border border-slate-200 rounded-lg px-2 py-1 text-[10px] font-black uppercase outline-none focus:ring-1 focus:ring-blue-500">
                           <option value="Aktif">Aktif</option>
                           <option value="Nonaktif">Nonaktif</option>
                           <option value="Pending">Pending</option>
                           <option value="Editing">Editing</option>
                         </select>
                      </div>
                      <div v-else class="flex items-center gap-2">
                        <div :class="`w-2.5 h-2.5 rounded-full ${row.status === 'Aktif' ? 'bg-green-500' : (row.status === 'Pending' ? 'bg-amber-500' : 'bg-slate-300')}`"></div>
                        <span class="text-[11px] font-bold text-slate-700 dark:text-slate-300">{{ row.status }}</span>
                      </div>
                    </td>
                    <td class="p-4 pr-10 text-right">
                      <div v-if="row.editing" class="flex justify-end gap-3">
                        <button @click="saveRow(row)" class="p-2.5 bg-blue-500 text-white rounded-xl hover:bg-blue-600 shadow-lg shadow-blue-500/20 active:scale-95 transition-all cursor-pointer"><LucideSave class="w-4 h-4" /></button>
                        <button @click="cancelEdit(row)" class="p-2.5 bg-slate-200 text-slate-500 rounded-xl hover:bg-slate-300 active:scale-95 transition-all cursor-pointer"><LucideX class="w-4 h-4" /></button>
                      </div>
                      <div v-else class="flex justify-end gap-3 transition-all">
                        <button @click="editRow(row)" class="p-2.5 text-slate-400 hover:text-blue-500 hover:bg-white dark:hover:bg-slate-800 rounded-xl transition-all shadow-sm border border-slate-100 dark:border-slate-800 cursor-pointer"><LucidePencil class="w-4 h-4" /></button>
                        <button @click="deleteRow(row.id)" class="p-2.5 text-slate-400 hover:text-red-500 hover:bg-white dark:hover:bg-slate-800 rounded-xl transition-all shadow-sm border border-slate-100 dark:border-slate-800 cursor-pointer"><LucideTrash2 class="w-4 h-4" /></button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>

            <!-- Pagination -->
            <div class="px-10 py-2.5 bg-slate-50/50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between">
              <div class="flex items-center gap-4">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">
                  Menampilkan {{ filteredCompanies.length > 0 ? startIndex + 1 : 0 }}-{{ endIndex }} dari {{ filteredCompanies.length }} entitas
                </p>
                <div class="flex items-center gap-2 ml-4">
                  <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Baris:</span>
                  <select v-model="itemsPerPage" class="bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg px-2 py-1 text-[10px] font-black outline-none focus:ring-1 focus:ring-blue-500/30 transition-all cursor-pointer">
                    <option :value="5">5</option>
                    <option :value="10">10</option>
                    <option :value="20">20</option>
                    <option :value="50">50</option>
                  </select>
                </div>
              </div>
              
              <div class="flex items-center gap-2">
                <button 
                  @click="currentPage--" 
                  :disabled="currentPage === 1"
                  class="w-8 h-8 flex items-center justify-center text-slate-400 hover:text-blue-500 transition-colors disabled:opacity-30 disabled:hover:text-slate-400 cursor-pointer disabled:cursor-not-allowed"
                >
                  <LucideChevronLeft class="w-4 h-4" />
                </button>
                
                <div class="flex items-center gap-1">
                  <button 
                    v-for="page in totalPages" 
                    :key="page"
                    @click="currentPage = page"
                    class="w-8 h-8 rounded-lg text-[10px] font-black transition-all cursor-pointer"
                    :class="currentPage === page ? 'bg-[#1E3A5F] text-white shadow-md' : 'hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400'"
                  >
                    {{ page }}
                  </button>
                </div>

                <button 
                  @click="currentPage++" 
                  :disabled="currentPage === totalPages"
                  class="w-8 h-8 flex items-center justify-center text-slate-400 hover:text-blue-500 transition-colors disabled:opacity-30 disabled:hover:text-slate-400 cursor-pointer disabled:cursor-not-allowed"
                >
                  <LucideChevronRight class="w-4 h-4" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </main>

      <!-- Right Sidebar: Rules & Changes -->
      <aside class="w-80 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-8 overflow-y-auto custom-scrollbar">
        <!-- Validation Rules -->
        <section class="space-y-8 mb-12">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
            <LucideClipboardCheck class="w-5 h-5 text-blue-500" />
            VALIDATION RULES
          </h3>
          <div class="space-y-6">
            <div v-for="rule in rules" :key="rule.title" class="p-6 rounded-2xl border-l-4 border-red-500 bg-red-50/50 dark:bg-red-900/10 space-y-3 relative group shadow-sm">
              <LucideInfo class="absolute top-4 right-4 w-3.5 h-3.5 text-slate-300" />
              <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ rule.title }}</h4>
              <p class="text-[10px] font-bold text-slate-500 leading-relaxed">{{ rule.desc }}</p>
              <div class="h-1 bg-slate-200 dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-red-500 w-[70%]"></div>
              </div>
              <div class="flex justify-between items-center text-[8px] font-black uppercase tracking-widest text-red-500">
                <span>Mandatory</span>
                <span>Active</span>
              </div>
            </div>

            <div class="p-6 rounded-2xl border-l-4 border-blue-500 bg-blue-50/50 dark:bg-blue-900/10 space-y-4 shadow-sm">
               <div class="flex justify-between items-start">
                 <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Tax Compliance</h4>
                 <LucideInfo class="w-3.5 h-3.5 text-slate-300" />
               </div>
               <p class="text-[10px] font-bold text-slate-500 leading-relaxed">NPWP wajib dilampirkan dalam format PDF (max 2MB) untuk setiap entitas PT baru yang didaftarkan.</p>
               <div class="flex gap-2">
                 <span class="px-2 py-1 bg-blue-100 dark:bg-blue-800 text-[8px] font-black text-blue-600 dark:text-blue-300 rounded uppercase tracking-widest shadow-sm">Regex Check</span>
                 <span class="px-2 py-1 bg-blue-100 dark:bg-blue-800 text-[8px] font-black text-blue-600 dark:text-blue-300 rounded uppercase tracking-widest shadow-sm">LDAP</span>
               </div>
            </div>
          </div>
        </section>

        <!-- Pending Changes -->
        <section class="space-y-8 mt-auto pt-8 border-t border-slate-100 dark:border-slate-800">
          <h3 class="text-xs font-black text-slate-400 uppercase tracking-[0.2em]">Pending Changes</h3>
          <div class="space-y-4">
            <div class="flex items-start gap-4 p-2 hover:bg-slate-50 dark:hover:bg-slate-800 rounded-xl transition-colors cursor-default">
              <div class="w-9 h-9 rounded-xl flex items-center justify-center shadow-sm bg-blue-50 text-blue-500 border border-blue-100">
                <LucidePencil class="w-4 h-4" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight truncate">Sync Service</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest truncate">Live connection ready</p>
              </div>
            </div>
          </div>
          <div class="space-y-4 mt-6">
            <button @click="publishChanges" class="w-full py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/30 flex items-center justify-center gap-3 transition-all active:scale-95 group cursor-pointer">
              <LucideRefreshCw class="w-4 h-4 group-hover:rotate-180 transition-transform duration-500" />
              Publish Changes
            </button>
          </div>
        </section>
      </aside>
    </div>

    <!-- Error Toast -->
    <div v-if="error" class="fixed bottom-10 right-10 z-[100] bg-red-500 text-white px-6 py-4 rounded-2xl shadow-2xl flex items-center gap-4 animate-bounce">
       <LucideBan class="w-6 h-6" />
       <div>
         <p class="text-xs font-black uppercase tracking-widest">Notification</p>
         <p class="text-sm font-bold">{{ error }}</p>
       </div>
       <button @click="error = ''" class="ml-4 hover:opacity-70 cursor-pointer"><LucideX class="w-4 h-4" /></button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { 
  LucideBuilding2, LucideUsers, LucideFileText, LucideArchive, 
  LucideClock, LucideUpload, LucidePlus, LucideSearch, LucideDownload, 
  LucideSettings2, LucideCheckCircle2, LucidePencil, LucideBan, 
  LucideSave, LucideTrash2, LucideChevronLeft, LucideChevronRight,
  LucideClipboardCheck, LucideInfo, LucideRefreshCw, LucideTrash,
  LucideX
} from 'lucide-vue-next'

const route = useRoute()
const activeEntity = ref(route.query.entity || 'pt')
const searchQuery = ref('')
const loading = ref(false)
const error = ref('')
const companies = ref([])
const fileInput = ref(null)

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(10)

const activeEntityLabel = computed(() => {
  switch (activeEntity.value) {
    case 'pt': return 'Master Perusahaan'
    case 'dept': return 'Master Departemen'
    case 'type': return 'Master Tipe Dokumen'
    case 'location': return 'Master Lokasi Arsip'
    case 'retention': return 'Master Kode Retensi'
    default: return 'Master Data'
  }
})

const fetchCompanies = async () => {
  loading.value = true
  try {
    const config = useRuntimeConfig()
    const auth = useAuthStore()
    const res = await $fetch(`${config.public.apiBase}/master/companies`, {
      headers: {
        Authorization: `Bearer ${auth.accessToken}`
      }
    })
    const data = res.data || []
    companies.value = data.map(c => ({
      ...c,
      editing: false,
      isNew: false
    }))
  } catch (err) {
    if (err.status === 401) {
       error.value = 'Sesi habis, silakan login kembali'
       setTimeout(() => navigateTo('/login'), 2000)
    } else {
       error.value = 'Failed to fetch data'
    }
  } finally {
    loading.value = false
  }
}

// --- Import Logic ---
const triggerImport = () => {
  fileInput.value.click()
}

const handleFileImport = (event) => {
  const file = event.target.files[0]
  if (!file) return

  const reader = new FileReader()
  reader.onload = async (e) => {
    const content = e.target.result
    const rows = content.split('\n').map(row => row.split(','))
    
    if (rows.length < 2) return

    loading.value = true
    let successCount = 0
    
    for (let i = 1; i < rows.length; i++) {
      const rowData = rows[i]
      if (rowData.length < 2) continue

      const entity_id = (rowData[0] || '').trim().replace(/"/g, '')
      const name = (rowData[1] || '').trim().replace(/"/g, '')
      const npwp_status = (rowData[2] || 'PENDING').trim().replace(/"/g, '')
      const location = (rowData[3] || '').trim().replace(/"/g, '')
      const status = (rowData[4] || 'Aktif').trim().replace(/"/g, '')

      if (!name || !entity_id) continue

      const existing = companies.value.find(c => c.entity_id === entity_id)

      try {
        const config = useRuntimeConfig()
        const auth = useAuthStore()
        
        const method = existing ? 'PUT' : 'POST'
        const url = existing 
          ? `${config.public.apiBase}/master/companies/${existing.id}`
          : `${config.public.apiBase}/master/companies`

        await $fetch(url, {
          method,
          headers: { Authorization: `Bearer ${auth.accessToken}` },
          body: {
            entity_id, name, npwp_status, location, status,
            address: ''
          }
        })
        successCount++
      } catch (err) {
        console.error('Failed to import row:', rows[i], err)
      }
    }

    error.value = `Berhasil mengimpor ${successCount} data.`
    fetchCompanies()
    event.target.value = ''
  }
  reader.readAsText(file)
}

// --- Export Logic ---
const exportCSV = () => {
  if (companies.value.length === 0) {
    error.value = 'Tidak ada data untuk diunduh'
    return
  }

  const headers = ['Entity ID', 'Nama / Deskripsi', 'NPWP Status', 'Lokasi / Detail', 'Status']
  const csvContent = [
    headers.join(','),
    ...companies.value.map(c => [
      c.entity_id,
      `"${c.name}"`,
      c.npwp_status,
      `"${c.location}"`,
      c.status
    ].join(','))
  ].join('\n')

  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.setAttribute('href', url)
  link.setAttribute('download', `${activeEntity.value}_export_${new Date().toISOString().split('T')[0]}.csv`)
  link.click()
}

const addRow = () => {
  companies.value.unshift({
    id: '', 
    entity_id: 'NEW-000',
    name: '',
    npwp_status: 'PENDING',
    location: '',
    status: 'Editing',
    editing: true,
    isNew: true
  })
  currentPage.value = 1
}

const editRow = (row) => {
  row.editing = true
}

const cancelEdit = (row) => {
  if (row.isNew) {
    companies.value = companies.value.filter(c => c !== row)
  } else {
    row.editing = false
    fetchCompanies()
  }
}

const saveRow = async (row) => {
  if (!row.name.trim()) {
    error.value = 'Name is required'
    return
  }

  try {
    const config = useRuntimeConfig()
    const auth = useAuthStore()
    const method = row.isNew ? 'POST' : 'PUT'
    const url = row.isNew 
      ? `${config.public.apiBase}/master/companies` 
      : `${config.public.apiBase}/master/companies/${row.id}`

    await $fetch(url, {
      method,
      headers: {
        Authorization: `Bearer ${auth.accessToken}`
      },
      body: {
        name: row.name,
        entity_id: row.entity_id,
        npwp_status: row.npwp_status,
        location: row.location,
        status: row.status,
        address: ''
      }
    })

    row.editing = false
    row.isNew = false
    fetchCompanies()
  } catch (err) {
    if (err.status === 401) {
       error.value = 'Sesi habis, silakan login kembali'
       setTimeout(() => navigateTo('/login'), 2000)
    } else {
       error.value = err.data?.message || 'Failed to save data'
    }
  }
}

const deleteRow = async (id) => {
  if (!confirm('Are you sure you want to delete this item?')) return

  try {
    const config = useRuntimeConfig()
    const auth = useAuthStore()
    await $fetch(`${config.public.apiBase}/master/companies/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${auth.accessToken}`
      }
    })
    fetchCompanies()
  } catch (err) {
    if (err.status === 401) {
       error.value = 'Sesi habis, silakan login kembali'
       setTimeout(() => navigateTo('/login'), 2000)
    } else {
       error.value = 'Failed to delete item'
    }
  }
}

const filteredCompanies = computed(() => {
  if (!searchQuery.value) return companies.value
  return companies.value.filter(c => 
    (c.name || '').toLowerCase().includes(searchQuery.value.toLowerCase()) || 
    (c.entity_id || '').toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const totalPages = computed(() => Math.ceil(filteredCompanies.value.length / itemsPerPage.value) || 1)
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, filteredCompanies.value.length))
const paginatedCompanies = computed(() => filteredCompanies.value.slice(startIndex.value, endIndex.value))

watch([searchQuery, itemsPerPage], () => {
  currentPage.value = 1
})

watch(() => route.query.entity, (newEntity) => {
  if (newEntity) {
    activeEntity.value = newEntity
  } else {
    activeEntity.value = 'pt'
  }
})

const rules = [
  { 
    title: 'Standard Format', 
    desc: "Pastikan ID / Code bersifat unik dan nama entitas sesuai dengan dokumen legal perusahaan."
  }
]

onMounted(() => {
  fetchCompanies()
})
</script>

<style scoped>
@reference "../../assets/css/main.css";

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  @apply bg-slate-300 dark:bg-slate-700;
}
</style>
