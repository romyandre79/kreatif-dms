<template>
  <div class="flex flex-col h-full bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <div class="flex flex-1 overflow-hidden">
      <!-- Main Content Area -->
      <main class="flex-grow flex flex-col overflow-hidden">
        <!-- Header -->
        <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
          <div class="space-y-1">
            <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
              {{ $t('admin.config.rack.title') }}
            </h1>
            <p class="text-sm font-bold text-slate-500 uppercase tracking-tighter">{{ $t('admin.config.rack.subtitle') }}</p>
          </div>
          <div class="flex items-center gap-4">
            <button @click="addRow" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95 cursor-pointer">
              <LucidePlus class="w-4 h-4" />
              {{ $t('admin.config.rack.add_btn') }}
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
                  :placeholder="$t('admin.config.rack.filter_placeholder')" 
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
            <div v-else class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                    <th class="p-4 pl-8">{{ $t('admin.config.rack.table.name') }}</th>
                    <th class="p-4">{{ $t('admin.config.rack.table.dept') }}</th>
                    <th class="p-4">{{ $t('admin.config.rack.table.branch') }}</th>
                    <th class="p-4">{{ $t('admin.config.rack.table.location') }}</th>
                    <th class="p-4 pr-10 text-right">{{ $t('admin.metadata.table.actions') }}</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                  <tr v-if="filteredRacks.length === 0" class="hover:bg-transparent">
                     <td colspan="5" class="p-20 text-center">
                        <div class="flex flex-col items-center justify-center space-y-3 opacity-30">
                           <LucideArchive class="w-12 h-12 text-slate-300" />
                           <p class="text-xs font-black uppercase tracking-[0.2em] text-slate-400">{{ $t('common.no_data') }}</p>
                        </div>
                     </td>
                  </tr>
                  <tr v-else v-for="row in paginatedRacks" :key="row.id" 
                    class="group transition-all"
                    :class="row.isNew ? 'bg-blue-50/50 dark:bg-blue-900/10' : 'hover:bg-slate-50/50 dark:hover:bg-slate-800/30'"
                  >
                    <td class="p-4 pl-8">
                      <div v-if="row.editing" class="max-w-md">
                        <input 
                          type="text" 
                          v-model="row.name" 
                          class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-3 text-sm font-black text-[#1E3A5F] dark:text-white uppercase outline-none shadow-lg shadow-blue-900/10" 
                          :placeholder="$t('admin.config.rack.table.name')"
                        />
                      </div>
                      <p v-else class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ row.name }}</p>
                    </td>
                    <td class="p-4">
                      <div v-if="row.editing && row.isNew" class="max-w-xs">
                        <select v-model="row.department_id" class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl px-4 py-2 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20">
                          <option value="" disabled>{{ $t('admin.config.rack.select_dept') }}</option>
                          <option v-for="dept in departments" :key="dept.id" :value="dept.id">
                            {{ dept.name }} ({{ dept.branch_name }})
                          </option>
                        </select>
                      </div>
                      <span v-else class="text-sm font-bold text-blue-500 bg-blue-50 dark:bg-blue-900/20 px-3 py-1 rounded-lg">{{ row.department_name }}</span>
                    </td>
                    <td class="p-4">
                      <span class="text-sm font-bold text-slate-500 uppercase tracking-tight">{{ row.branch_name }}</span>
                    </td>
                    <td class="p-4">
                      <div v-if="row.editing" class="max-w-md">
                        <input 
                          type="text" 
                          v-model="row.location_detail" 
                          class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl px-4 py-2 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20" 
                          :placeholder="$t('admin.config.rack.table.location')"
                        />
                      </div>
                      <p v-else class="text-xs font-bold text-slate-400 uppercase tracking-tight">{{ row.location_detail || '-' }}</p>
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
                  {{ $t('admin.config.rack.displaying', { start: filteredRacks.length > 0 ? startIndex + 1 : 0, end: endIndex, total: filteredRacks.length }) }}
                </p>
                <div class="flex items-center gap-2 ml-4">
                  <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.config.rack.rows') }}:</span>
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

      <!-- Right Sidebar: Rules -->
      <aside class="w-80 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-8 overflow-y-auto custom-scrollbar">
        <section class="space-y-8">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
            <LucideClipboardCheck class="w-5 h-5 text-blue-500" />
            RACK RULES
          </h3>
          <div class="space-y-6">
            <div class="p-6 rounded-2xl border-l-4 border-blue-500 bg-blue-50/50 dark:bg-blue-900/10 space-y-3 shadow-sm">
              <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Ownership</h4>
              <p class="text-[10px] font-bold text-slate-500 leading-relaxed">Rak dimiliki oleh satu departemen tertentu untuk menjaga kerahasiaan dokumen.</p>
            </div>
            <div class="p-6 rounded-2xl border-l-4 border-amber-500 bg-amber-50/50 dark:bg-amber-900/10 space-y-3 shadow-sm">
              <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Naming Convention</h4>
              <p class="text-[10px] font-bold text-slate-500 leading-relaxed">Gunakan format penamaan yang memudahkan identifikasi (Contoh: RAK-FIN-01).</p>
            </div>
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
import { useI18n } from 'vue-i18n'
import { 
  LucidePlus, LucideSearch, LucideRefreshCw, LucideArchive,
  LucidePencil, LucideTrash2, LucideSave, LucideX,
  LucideChevronLeft, LucideChevronRight, LucideClipboardCheck,
  LucideBan
} from 'lucide-vue-next'

const { t } = useI18n()
const searchQuery = ref('')
const loading = ref(false)
const error = ref('')
const racks = ref([])
const departments = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(10)

const config = useRuntimeConfig()
const auth = useAuthStore()

const fetchData = async () => {
  loading.value = true
  try {
    const [racksRes, deptsRes] = await Promise.all([
      $fetch(`${config.public.apiBase}/master/racks`, {
        headers: { Authorization: `Bearer ${auth.accessToken}` }
      }),
      $fetch(`${config.public.apiBase}/master/departments`, {
        headers: { Authorization: `Bearer ${auth.accessToken}` }
      })
    ])

    racks.value = (racksRes.data || []).map(r => ({
      ...r,
      editing: false,
      isNew: false
    }))
    departments.value = deptsRes.data || []
  } catch (err) {
    error.value = err.data?.message || t('admin.config.rack.error_fetch')
  } finally {
    loading.value = false
  }
}

const addRow = () => {
  racks.value.unshift({
    id: '', 
    department_id: departments.value.length > 0 ? departments.value[0].id : '',
    name: '',
    location_detail: '',
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
    racks.value = racks.value.filter(r => r !== row)
  } else {
    row.editing = false
    fetchData()
  }
}

const saveRow = async (row) => {
  if (!row.name.trim()) {
    error.value = t('admin.config.rack.error_name')
    return
  }
  if (!row.department_id) {
    error.value = t('admin.config.rack.error_dept')
    return
  }

  try {
    const method = row.isNew ? 'POST' : 'PUT'
    const url = row.isNew 
      ? `${config.public.apiBase}/master/racks` 
      : `${config.public.apiBase}/master/racks/${row.id}`

    await $fetch(url, {
      method,
      headers: { Authorization: `Bearer ${auth.accessToken}` },
      body: {
        name: row.name,
        department_id: row.department_id,
        location_detail: row.location_detail
      }
    })

    row.editing = false
    row.isNew = false
    fetchData()
  } catch (err) {
    error.value = err.data?.message || t('admin.config.rack.error_save')
  }
}

const deleteRow = async (id) => {
  if (!confirm(t('admin.config.rack.confirm_delete'))) return

  try {
    await $fetch(`${config.public.apiBase}/master/racks/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${auth.accessToken}` }
    })
    fetchData()
  } catch (err) {
    error.value = t('admin.config.rack.error_delete')
  }
}

const filteredRacks = computed(() => {
  if (!searchQuery.value) return racks.value
  const q = searchQuery.value.toLowerCase()
  return racks.value.filter(r => 
    (r.name || '').toLowerCase().includes(q) || 
    (r.department_name || '').toLowerCase().includes(q) ||
    (r.branch_name || '').toLowerCase().includes(q)
  )
})

const totalPages = computed(() => Math.ceil(filteredRacks.value.length / itemsPerPage.value) || 1)
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, filteredRacks.value.length))
const paginatedRacks = computed(() => filteredRacks.value.slice(startIndex.value, endIndex.value))

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
