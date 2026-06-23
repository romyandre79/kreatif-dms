<template>
  <div class="flex flex-col h-full bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>

    <div class="flex flex-1 overflow-hidden">
      <!-- Main Content Area -->
      <main class="flex-grow flex flex-col overflow-hidden">
        <!-- Header -->
        <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
          <div class="space-y-1">
            <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
              {{ $t('layout.menu.floor') }}
            </h1>
            <p class="text-sm font-bold text-slate-500 uppercase tracking-tighter">Manajemen Master Lantai</p>
          </div>
          <div class="flex items-center gap-4">
            <button @click="addRow" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95 cursor-pointer">
              <LucidePlus class="w-4 h-4" />
              Tambah Lantai
            </button>
          </div>
        </header>

        <!-- Table Section -->
        <div class="flex-grow p-2.5 overflow-auto custom-scrollbar">
          <div class="bg-white dark:bg-slate-900 rounded-lg shadow-sm border border-slate-100 dark:border-slate-800">
            <!-- Filter Bar -->
            <div class="px-8 py-6 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
              <div class="relative w-96 group">
                <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
                <input 
                  type="text" 
                  v-model="searchQuery"
                  placeholder="Cari lantai berdasarkan nama atau kode..." 
                  class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl pl-12 pr-6 py-3 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20 transition-all" 
                />
              </div>
              <div class="flex items-center gap-4">
                <button @click="fetchData" class="p-3 text-slate-400 hover:text-blue-500 transition-colors cursor-pointer" title="Refresh"><LucideRefreshCw class="w-5 h-5" /></button>
              </div>
            </div>

            <!-- Loading State -->
            <div v-if="loading" class="p-20 flex flex-col items-center justify-center space-y-4">
               <LucideRefreshCw class="w-10 h-10 text-blue-500 animate-spin" />
               <p class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ $t('common.loading') }}</p>
            </div>

            <!-- Table -->
            <div v-else class="overflow-x-auto" :class="{ 'overflow-visible relative z-50': hasEditingRow }">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                    <th class="p-4 pl-8">Nama Lantai</th>
                    <th class="p-4">Kode Lantai</th>
                    <th class="p-4 pr-10 text-right">{{ $t('admin.metadata.table.actions') }}</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                  <tr v-if="filteredFloors.length === 0" class="hover:bg-transparent">
                     <td colspan="3" class="p-20 text-center">
                        <div class="flex flex-col items-center justify-center space-y-3 opacity-30">
                           <LucideArchive class="w-12 h-12 text-slate-300" />
                           <p class="text-xs font-black uppercase tracking-[0.2em] text-slate-400">{{ $t('common.no_data') }}</p>
                        </div>
                     </td>
                  </tr>
                  <tr v-else v-for="row in paginatedFloors" :key="row.id" 
                    class="group transition-all"
                    :class="[
                      row.isNew ? 'bg-blue-50/50 dark:bg-blue-900/10' : 'hover:bg-slate-50/50 dark:hover:bg-slate-800/30',
                      row.editing ? 'relative z-20' : ''
                    ]"
                  >
                    <td class="p-4 pl-8">
                      <div v-if="row.editing" class="max-w-md">
                        <input 
                          type="text" 
                          v-model="row.name" 
                          class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-3 text-sm font-black text-[#1E3A5F] dark:text-white uppercase outline-none shadow-lg shadow-blue-900/10" 
                          placeholder="Contoh: Lantai 02-B (Warehouse Main)"
                        />
                      </div>
                      <p v-else class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ row.name }}</p>
                    </td>
                    <td class="p-4">
                      <div v-if="row.editing" class="max-w-md">
                        <input 
                          type="text" 
                          v-model="row.code" 
                          class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl px-4 py-2 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20" 
                          placeholder="Contoh: Floor 02-B"
                        />
                      </div>
                      <p v-else class="text-xs font-bold text-slate-400 uppercase tracking-tight">{{ row.code || '-' }}</p>
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
            <div class="px-10 py-2.5 bg-slate-50/50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between relative z-0">
              <div class="flex items-center gap-4">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">
                  Menampilkan {{ filteredFloors.length > 0 ? startIndex + 1 : 0 }} - {{ endIndex }} dari {{ filteredFloors.length }} data
                </p>
                <div class="flex items-center gap-2 ml-4">
                  <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Baris:</span>
                  <select v-model="itemsPerPage" class="bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg px-2 py-1 text-[10px] font-black outline-none focus:ring-1 focus:ring-blue-500/30 transition-all cursor-pointer">
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
import { useI18n } from 'vue-i18n'
import { 
  LucidePlus, LucideSearch, LucideRefreshCw, LucideArchive,
  LucidePencil, LucideTrash2, LucideSave, LucideX,
  LucideChevronLeft, LucideChevronRight,
  LucideBan
} from 'lucide-vue-next'

const { t } = useI18n()
const searchQuery = ref('')
const loading = ref(false)
const error = ref('')
const floors = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(10)

const config = useRuntimeConfig()
const { $api } = useApi()

const fetchData = async () => {
  loading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/floors`)
    floors.value = (res.data || []).map(r => ({
      ...r,
      editing: false,
      isNew: false
    }))
  } catch (err) {
    error.value = err.data?.message || "Gagal mengambil data lantai"
  } finally {
    loading.value = false
  }
}


const addRow = () => {
  floors.value.unshift({
    id: '', 
    name: '',
    code: '',
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
    floors.value = floors.value.filter(r => r !== row)
  } else {
    row.editing = false
    fetchData()
  }
}

const saveRow = async (row) => {
  if (!row.name.trim()) {
    error.value = "Nama lantai harus diisi"
    return
  }
  if (!row.code.trim()) {
    error.value = "Kode lantai harus diisi"
    return
  }

  try {
    const method = row.isNew ? 'POST' : 'PUT'
    const url = row.isNew 
      ? `${config.public.apiBase}/master/floors` 
      : `${config.public.apiBase}/master/floors/${row.id}`

    await $api(url, {
      method,
      body: {
        name: row.name,
        code: row.code
      }
    })

    row.editing = false
    row.isNew = false
    fetchData()
  } catch (err) {
    error.value = err.data?.message || "Gagal menyimpan data lantai"
  }
}

const deleteRow = async (id) => {
  if (!confirm("Apakah Anda yakin ingin menghapus data lantai ini?")) return

  try {
    await $api(`${config.public.apiBase}/master/floors/${id}`, {
      method: 'DELETE'
    })
    fetchData()
  } catch (err) {
    error.value = "Gagal menghapus data lantai. Mungkin sedang digunakan."
  }
}

const hasEditingRow = computed(() => floors.value.some(r => r.editing))

const filteredFloors = computed(() => {
  if (!searchQuery.value) return floors.value
  const q = searchQuery.value.toLowerCase()
  return floors.value.filter(r => 
    (r.name || '').toLowerCase().includes(q) || 
    (r.code || '').toLowerCase().includes(q)
  )
})

const totalPages = computed(() => Math.ceil(filteredFloors.value.length / itemsPerPage.value) || 1)
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, filteredFloors.value.length))
const paginatedFloors = computed(() => filteredFloors.value.slice(startIndex.value, endIndex.value))

watch([searchQuery, itemsPerPage], () => {
  currentPage.value = 1
})

onMounted(() => {
  fetchData()
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
