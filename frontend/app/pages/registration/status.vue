<template>
  <div class="max-w-7xl mx-auto space-y-6 pb-24">
    <!-- Header -->
    <div class="flex flex-col gap-4" v-motion-fade>
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
            Status Pengajuan Dokumen
          </h1>
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">
            Semua pengajuan di departemen Anda
          </p>
        </div>
        <div class="flex items-center gap-2">
          <span class="w-2 h-2 bg-primary-500 rounded-full animate-pulse"></span>
          <span class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ pagination.total }} Dokumen</span>
        </div>
      </div>

      <!-- Stats Bar -->
      <div class="grid grid-cols-2 md:grid-cols-5 gap-3">
        <button
          v-for="stat in statCards"
          :key="stat.status"
          @click="setStatusFilter(stat.status)"
          :class="[
            'glass rounded-2xl p-3.5 text-left transition-all hover:scale-[1.02] active:scale-[0.98] border',
            activeStatus === stat.status
              ? 'border-primary-300 ring-2 ring-primary-500/20'
              : 'border-transparent'
          ]"
        >
          <p class="text-[10px] font-black uppercase tracking-widest text-slate-400">{{ stat.label }}</p>
          <p class="text-2xl font-black mt-1" :class="stat.color">{{ stat.count }}</p>
        </button>
      </div>

      <!-- Search & Filters -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-3">
        <div class="md:col-span-2 relative group">
          <LucideSearch class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari judul dokumen atau nama pengaju..."
            class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl py-2.5 pl-10 pr-4 text-sm font-medium outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm"
            @input="onSearch"
          />
        </div>
        <div class="relative">
          <select
            v-model="activeStatus"
            @change="loadData"
            class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl py-2.5 px-3 text-sm font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm cursor-pointer"
          >
            <option value="">Semua Status</option>
            <option value="draft">Draft</option>
            <option value="pending">Menunggu Persetujuan</option>
            <option value="approved">Disetujui</option>
            <option value="rejected">Ditolak</option>
            <option value="active">Terarsip</option>
          </select>
          <LucideChevronDown class="absolute right-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
        </div>
      </div>
    </div>

    <!-- Table -->
    <div class="glass rounded-3xl overflow-hidden" v-motion-fade>
      <!-- Loading State -->
      <div v-if="isLoading" class="py-16 flex flex-col items-center justify-center gap-3">
        <div class="w-8 h-8 border-2 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">Memuat data...</p>
      </div>

      <!-- Empty State -->
      <div v-else-if="items.length === 0" class="py-16 flex flex-col items-center justify-center gap-3">
        <LucideFileX class="w-12 h-12 text-slate-200 dark:text-slate-700" />
        <p class="text-sm font-black text-slate-400 uppercase tracking-widest">Tidak ada pengajuan ditemukan</p>
        <p class="text-xs font-medium text-slate-300 dark:text-slate-600">Coba ubah filter atau kata kunci pencarian</p>
      </div>

      <!-- Table Content -->
      <table v-else class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/30 border-b border-slate-100 dark:border-slate-700">
            <th class="px-5 py-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">Dokumen</th>
            <th class="px-4 py-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">Pengaju</th>
            <th class="px-4 py-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">Tipe</th>
            <th class="px-4 py-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">Tanggal Pengajuan</th>
            <th class="px-4 py-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">Status</th>
            <th class="pr-5 py-3 text-right text-[10px] font-black text-slate-400 uppercase tracking-widest">Disetujui Oleh</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <tr
            v-for="item in items"
            :key="item.id"
            class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-colors cursor-pointer"
            @click="navigateTo(`/documents/${item.id}`)"
          >
            <!-- Document info -->
            <td class="px-5 py-3.5">
              <div class="space-y-0.5">
                <p class="text-xs font-black text-[#1E3A5F] dark:text-white leading-tight line-clamp-1">{{ item.title }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">
                  {{ formatFileSize(item.file_size) }}
                  <span v-if="item.sensitivity" class="ml-1.5 px-1.5 py-0.5 rounded bg-slate-100 dark:bg-slate-700">{{ item.sensitivity }}</span>
                </p>
              </div>
            </td>

            <!-- Submitter -->
            <td class="px-4 py-3.5">
              <div class="flex items-center gap-2">
                <div class="w-6 h-6 rounded-full bg-slate-100 dark:bg-slate-700 overflow-hidden ring-1 ring-white dark:ring-slate-600 shadow-sm shrink-0">
                  <img :src="`https://i.pravatar.cc/50?u=${item.owner_email}`" class="w-full h-full object-cover" />
                </div>
                <p class="text-xs font-bold text-[#1E3A5F] dark:text-slate-200 leading-tight">{{ item.owner_name || '—' }}</p>
              </div>
            </td>

            <!-- Type -->
            <td class="px-4 py-3.5">
              <span class="text-[9px] font-black text-slate-500 dark:text-slate-400 uppercase tracking-widest">
                {{ item.type_name || '—' }}
              </span>
            </td>

            <!-- Date -->
            <td class="px-4 py-3.5">
              <div class="space-y-0.5">
                <p class="text-[11px] font-black text-slate-700 dark:text-slate-300 leading-tight">{{ formatDate(item.created_at) }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ formatTime(item.created_at) }}</p>
              </div>
            </td>

            <!-- Status -->
            <td class="px-4 py-3.5">
              <div class="space-y-1">
                <span :class="['px-2.5 py-1 rounded-full text-[9px] font-black uppercase tracking-widest border', statusBadge(item.status).class]">
                  {{ statusBadge(item.status).label }}
                </span>
                <p v-if="item.approval_note" class="text-[8px] font-medium text-slate-400 leading-tight line-clamp-1 max-w-[140px]">
                  {{ item.approval_note }}
                </p>
              </div>
            </td>

            <!-- Approver -->
            <td class="pr-5 py-3.5 text-right">
              <div v-if="item.approver_name" class="space-y-0.5">
                <p class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-200">{{ item.approver_name }}</p>
                <p v-if="item.decided_at" class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ formatDate(item.decided_at) }}</p>
              </div>
              <span v-else class="text-[9px] font-bold text-slate-300 dark:text-slate-600 uppercase tracking-widest">Belum diproses</span>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="!isLoading && items.length > 0" class="px-5 py-4 bg-slate-50/30 dark:bg-slate-800/20 flex items-center justify-between border-t border-slate-100 dark:border-slate-700">
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
          Menampilkan {{ (pagination.page - 1) * pagination.limit + 1 }}–{{ Math.min(pagination.page * pagination.limit, pagination.total) }} dari {{ pagination.total }}
        </p>
        <div class="flex items-center gap-2">
          <button
            :disabled="pagination.page <= 1"
            @click="changePage(pagination.page - 1)"
            class="w-7 h-7 rounded-lg bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 flex items-center justify-center text-slate-400 hover:text-primary-600 hover:border-primary-200 transition-all disabled:opacity-30 disabled:cursor-not-allowed"
          >
            <LucideChevronLeft class="w-3.5 h-3.5" />
          </button>
          <span class="text-[10px] font-black text-slate-500 uppercase tracking-widest px-2">{{ pagination.page }} / {{ pagination.total_pages }}</span>
          <button
            :disabled="pagination.page >= pagination.total_pages"
            @click="changePage(pagination.page + 1)"
            class="w-7 h-7 rounded-lg bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 flex items-center justify-center text-slate-400 hover:text-primary-600 hover:border-primary-200 transition-all disabled:opacity-30 disabled:cursor-not-allowed"
          >
            <LucideChevronRight class="w-3.5 h-3.5" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  LucideSearch, LucideChevronDown, LucideChevronLeft, LucideChevronRight, LucideFileX
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'

const { $api } = useApi()

const isLoading = ref(false)
const searchQuery = ref('')
const activeStatus = ref('')
let searchTimer = null

const items = ref([])
const pagination = ref({ total: 0, page: 1, limit: 20, total_pages: 1 })

// Summary counts per status
const statusCounts = ref({ all: 0, draft: 0, pending: 0, approved: 0, rejected: 0, active: 0 })

const statCards = computed(() => [
  { status: '', label: 'Semua', count: statusCounts.value.all, color: 'text-[#1E3A5F] dark:text-white' },
  { status: 'pending', label: 'Menunggu', count: statusCounts.value.pending, color: 'text-amber-500' },
  { status: 'approved', label: 'Disetujui', count: statusCounts.value.approved, color: 'text-green-600' },
  { status: 'rejected', label: 'Ditolak', count: statusCounts.value.rejected, color: 'text-red-500' },
  { status: 'active', label: 'Terarsip', count: statusCounts.value.active, color: 'text-sky-500' },
])

const loadData = async () => {
  isLoading.value = true
  try {
    const params = new URLSearchParams({
      page: pagination.value.page,
      limit: pagination.value.limit,
    })
    if (activeStatus.value) params.set('status', activeStatus.value)
    if (searchQuery.value.trim()) params.set('q', searchQuery.value.trim())

    const res = await $api(`/documents/submissions?${params}`)
    if (res?.data) {
      items.value = res.data.items || []
      pagination.value = { ...pagination.value, ...res.data.pagination }
    }
  } catch (e) {
    console.error('[SubmissionStatus] Failed to load:', e)
    items.value = []
  } finally {
    isLoading.value = false
  }
}

const loadCounts = async () => {
  const statuses = ['', 'draft', 'pending', 'approved', 'rejected', 'active']
  const results = await Promise.allSettled(
    statuses.map(s => $api(`/documents/submissions?limit=1${s ? `&status=${s}` : ''}`))
  )
  results.forEach((r, i) => {
    if (r.status === 'fulfilled' && r.value?.data?.pagination) {
      const key = statuses[i] || 'all'
      statusCounts.value[key] = r.value.data.pagination.total
    }
  })
}

const setStatusFilter = (status) => {
  activeStatus.value = status
  pagination.value.page = 1
  loadData()
}

const onSearch = () => {
  clearTimeout(searchTimer)
  searchTimer = setTimeout(() => {
    pagination.value.page = 1
    loadData()
  }, 400)
}

const changePage = (page) => {
  pagination.value.page = page
  loadData()
}

// Formatters
const formatDate = (iso) => {
  if (!iso) return '—'
  return new Date(iso).toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' }).toUpperCase()
}
const formatTime = (iso) => {
  if (!iso) return ''
  return new Date(iso).toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}
const formatFileSize = (bytes) => {
  if (!bytes) return '—'
  if (bytes < 1024) return `${bytes} B`
  if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)} KB`
  return `${(bytes / (1024 * 1024)).toFixed(1)} MB`
}

const statusBadge = (status) => {
  const map = {
    draft:    { label: 'Draft',     class: 'bg-slate-50 text-slate-500 border-slate-100' },
    pending:  { label: 'Menunggu',  class: 'bg-amber-50 text-amber-600 border-amber-100' },
    approved: { label: 'Disetujui', class: 'bg-green-50 text-green-600 border-green-100' },
    rejected: { label: 'Ditolak',   class: 'bg-red-50 text-red-500 border-red-100' },
    active:   { label: 'Terarsip',  class: 'bg-sky-50 text-sky-600 border-sky-100' },
  }
  return map[status] || { label: status || '—', class: 'bg-slate-50 text-slate-400 border-slate-100' }
}

// Initial load
onMounted(() => {
  loadData()
  loadCounts()
})
</script>

<style scoped>
.glass {
  background-color: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(40px);
  -webkit-backdrop-filter: blur(40px);
  border: 1px solid rgba(255, 255, 255, 1);
  box-shadow: 0 25px 50px -12px rgba(226, 232, 240, 0.5);
}
:root.dark .glass {
  background-color: rgba(15, 23, 42, 0.7);
  border-color: rgba(255, 255, 255, 0.05);
}
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
