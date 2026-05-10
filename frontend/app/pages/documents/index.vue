<template>
  <div class="flex h-[calc(100vh-theme(spacing.16))] -m-8 overflow-hidden bg-[#F8FAFC]">
    <!-- Left Sidebar: Document Explorer -->
    <aside class="w-80 border-r border-slate-200 bg-white flex flex-col overflow-hidden">
      <div class="p-6 space-y-6">
        <h2 class="text-lg font-bold text-slate-800">{{ $t('documents.explorer.title') }}</h2>
        
        <div class="relative">
          <LucideSearch class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input
            type="text"
            :placeholder="$t('documents.explorer.search_placeholder')"
            class="w-full pl-10 pr-4 py-2 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all"
          />
        </div>
      </div>

      <!-- Folder Tree -->
      <div class="flex-1 overflow-y-auto px-4 pb-6 custom-scrollbar">
        <div class="space-y-1">
          <div v-for="node in folderTree" :key="node.id" class="space-y-1">
            <button 
              @click="toggleFolder(node)"
              class="w-full flex items-center gap-2 px-3 py-2 rounded-lg text-sm font-bold transition-colors"
              :class="node.expanded ? 'text-primary-600 bg-primary-50/50' : 'text-slate-600 hover:bg-slate-50'"
            >
              <LucideChevronDown v-if="node.children && node.children.length" class="w-4 h-4 transition-transform" :class="{ '-rotate-90': !node.expanded }" />
              <div v-else class="w-4"></div>
              <LucideFolder v-if="!node.expanded" class="w-4 h-4 text-amber-400 fill-amber-400" />
              <LucideFolderOpen v-else class="w-4 h-4 text-amber-400 fill-amber-400" />
              <span class="uppercase tracking-tight">{{ node.name }}</span>
            </button>

            <!-- Recursive Children -->
            <div v-if="node.expanded && node.children" class="ml-4 pl-4 border-l border-slate-100 space-y-1">
              <div v-for="child in node.children" :key="child.id">
                <button 
                  @click="toggleFolder(child)"
                  class="w-full flex items-center justify-between px-3 py-2 rounded-lg text-sm font-bold transition-colors"
                  :class="[
                    child.active && !query ? 'text-primary-600 bg-primary-50' : 'text-slate-500 hover:bg-slate-50',
                    child.expanded ? 'text-primary-600' : ''
                  ]"
                >
                  <div class="flex items-center gap-2">
                    <LucideChevronDown v-if="child.children && child.children.length" class="w-3.5 h-3.5 transition-transform" :class="{ '-rotate-90': !child.expanded }" />
                    <div v-else class="w-3.5"></div>
                    <LucideFolder v-if="!child.expanded" class="w-4 h-4 text-amber-400 fill-amber-400" />
                    <LucideFolderOpen v-else class="w-4 h-4 text-amber-400 fill-amber-400" />
                    <span class="uppercase tracking-tight">{{ child.name }}</span>
                  </div>
                  <span v-if="child.count" class="text-[10px] text-slate-400 bg-slate-100 px-1.5 py-0.5 rounded-md">({{ child.count }})</span>
                </button>

                <!-- Second Level -->
                <div v-if="child.expanded && child.children" class="ml-4 pl-4 border-l border-slate-100 space-y-1">
                  <button 
                    v-for="subChild in child.children" 
                    :key="subChild.id"
                    class="w-full flex items-center gap-2 px-3 py-2 rounded-lg text-xs font-bold transition-colors"
                    :class="subChild.active && !query ? 'text-primary-600 bg-primary-50' : 'text-slate-400 hover:bg-slate-50 hover:text-slate-600'"
                  >
                    <span class="uppercase tracking-tight">{{ subChild.name }}</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-1 flex flex-col overflow-hidden relative">
      <!-- HEADER: SEARCH RESULTS MODE -->
      <header v-if="query" class="bg-white border-b border-slate-200 px-8 py-8 flex items-center justify-between" v-motion-fade>
        <div>
          <h1 class="text-2xl font-black text-[#1E3A5F]" v-html="$t('documents.search.results_for', { query, count: searchResults.length })"></h1>
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-2">{{ $t('documents.search.stats_desc', { total: '15.234', time: '0.4' }) }}</p>
        </div>
        <div class="relative w-80">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input
            type="text"
            :placeholder="$t('documents.search.within_placeholder')"
            class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-2xl text-xs font-bold focus:outline-none focus:ring-4 focus:ring-primary-500/5 transition-all"
          />
        </div>
      </header>

      <!-- HEADER: EXPLORER MODE -->
      <header v-else class="bg-white border-b border-slate-200 px-8 py-4 flex items-center justify-between" v-motion-fade>
        <div class="flex items-center gap-2 text-xs font-bold text-slate-400">
          <span>{{ $t('documents.mock.accounting') }}</span>
          <LucideChevronRight class="w-3 h-3" />
          <span>2026</span>
          <LucideChevronRight class="w-3 h-3" />
          <span class="text-primary-600 uppercase">{{ $t('documents.mock.february') }}</span>
        </div>
        <div class="text-xs font-bold text-slate-400">{{ $t('documents.explorer.stats_total', { count: 128 }) }}</div>
      </header>

      <!-- FILTERS BAR (Only in Explorer Mode) -->
      <div v-if="!query" class="bg-white border-b border-slate-100 px-8 py-5 flex flex-wrap items-center gap-4">
        <div class="relative flex-1 min-w-[300px]">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input
            type="text"
            :placeholder="$t('documents.filters.search_within')"
            class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-2xl text-sm focus:outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all font-medium"
          />
        </div>

        <div class="flex items-center gap-3">
          <div class="flex items-center gap-2">
            <span class="text-xs font-bold text-slate-400">{{ $t('documents.filters.doc_type') }}</span>
            <button class="flex items-center gap-2 px-4 py-2.5 bg-white border border-slate-200 rounded-xl text-xs font-black text-slate-700 hover:bg-slate-50">{{ $t('documents.filters.all') }} <LucideChevronDown class="w-3.5 h-3.5" /></button>
          </div>
          <div class="flex items-center gap-2">
            <span class="text-xs font-bold text-slate-400">{{ $t('documents.filters.year') }}</span>
            <button class="flex items-center gap-2 px-4 py-2.5 bg-white border border-slate-200 rounded-xl text-xs font-black text-slate-700 hover:bg-slate-50">2026 <LucideChevronDown class="w-3.5 h-3.5" /></button>
          </div>
          <div class="flex items-center gap-2">
            <span class="text-xs font-bold text-slate-400">{{ $t('documents.filters.status') }}</span>
            <button class="flex items-center gap-2 px-4 py-2.5 bg-white border border-slate-200 rounded-xl text-xs font-black text-slate-700 hover:bg-slate-50">{{ $t('documents.filters.available') }} <LucideChevronDown class="w-3.5 h-3.5" /></button>
          </div>
          <button class="flex items-center gap-2 px-4 py-2.5 bg-white border border-slate-200 rounded-xl text-xs font-black text-slate-700 hover:bg-slate-50"><LucideFilter class="w-3.5 h-3.5" /> {{ $t('documents.filters.sort') }}</button>
        </div>
      </div>

      <!-- VIEW AREA -->
      <div class="flex-1 overflow-y-auto p-8 custom-scrollbar">
        <!-- SEARCH RESULTS VIEW -->
        <div v-if="query" class="space-y-4 max-w-5xl mx-auto">
          <div v-if="searchResults.length > 0" class="space-y-6">
            <div v-for="res in searchResults" :key="res.id" class="bg-white p-8 rounded-lg border border-slate-100 hover:border-primary-300 hover:shadow-xl hover:shadow-primary-900/5 transition-all group relative overflow-hidden">
              <div class="flex items-start justify-between">
                <div class="space-y-4 flex-1 pr-10">
                  <div class="space-y-1">
                    <h3 class="text-lg font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors uppercase tracking-tight cursor-pointer" @click="viewDocument(res)">{{ res.name }}</h3>
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ res.path }}</p>
                  </div>
                  <p class="text-sm text-slate-500 leading-relaxed line-clamp-2 italic" v-html="res.snippet"></p>
                  <div class="flex flex-wrap gap-2 pt-2">
                    <span v-for="tag in res.tags" :key="tag" class="px-3 py-1 bg-slate-50 border border-slate-100 text-slate-400 text-[9px] font-black rounded-lg uppercase tracking-widest">{{ tag }}</span>
                    <div class="flex items-center gap-1.5 ml-2 px-3 py-1 bg-green-50 text-green-600 rounded-lg text-[9px] font-black uppercase tracking-widest border border-green-100">
                      <div class="w-1 h-1 rounded-full bg-green-500"></div>
                      {{ $t('documents.filters.available') }}
                    </div>
                  </div>
                </div>
                <div class="flex items-center gap-2">
                  <button @click="viewDocument(res)" class="p-3 bg-slate-50 rounded-2xl text-slate-400 hover:text-primary-600 hover:bg-primary-50 transition-all"><LucideEye class="w-5 h-5" /></button>
                  <button class="p-3 bg-slate-50 rounded-2xl text-slate-400 hover:text-primary-600 hover:bg-primary-50 transition-all"><LucideShoppingCart class="w-5 h-5" /></button>
                </div>
              </div>
            </div>
          </div>

          <!-- EMPTY SEARCH STATE -->
          <div v-else class="h-full flex flex-col items-center justify-center py-20 text-center space-y-6" v-motion-fade>
            <div class="relative">
              <div class="w-32 h-32 bg-slate-100 rounded-full flex items-center justify-center">
                <LucideSearch class="w-16 h-16 text-slate-300" />
              </div>
              <div class="absolute -bottom-2 -right-2 w-10 h-10 bg-white rounded-full shadow-lg flex items-center justify-center border border-slate-50">
                <LucideX class="w-5 h-5 text-red-400" />
              </div>
            </div>
            <div class="space-y-2">
              <h2 class="text-2xl font-black text-[#1E3A5F]">{{ $t('documents.search.no_found') }}</h2>
              <p class="text-sm font-bold text-slate-400 max-w-md">{{ $t('documents.search.no_found_desc') }}</p>
            </div>
            <div class="text-xs font-bold text-slate-400 bg-white px-6 py-3 rounded-2xl border border-slate-100 shadow-sm flex items-center gap-3">
              {{ $t('documents.search.did_you_mean') }} 
              <span class="text-primary-600 cursor-pointer hover:underline">"kontrak"</span>, 
              <span class="text-primary-600 cursor-pointer hover:underline">"kontak"</span> ?
            </div>
          </div>
        </div>

        <!-- EXPLORER TABLE VIEW -->
        <div v-else class="bg-white rounded-lg shadow-sm border border-slate-100 overflow-hidden">
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-[0.15em] border-b border-slate-100">
                <th class="pl-8 pr-4 py-5 w-16">
                  <div class="flex items-center justify-center">
                    <input type="checkbox" class="w-5 h-5 rounded-md border-2 border-slate-200 text-primary-600 focus:ring-primary-500/20" />
                  </div>
                </th>
                <th class="px-6 py-5">{{ $t('documents.table.name') }}</th>
                <th class="px-6 py-5">{{ $t('documents.table.type') }}</th>
                <th class="px-6 py-5">{{ $t('documents.table.date') }}</th>
                <th class="px-6 py-5">{{ $t('documents.table.dept') }}</th>
                <th class="px-6 py-5">{{ $t('documents.table.status') }}</th>
                <th class="pr-8 pl-4 py-5 text-right">{{ $t('documents.table.actions') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="doc in documents" :key="doc.id" class="group hover:bg-slate-50/50 transition-colors">
                <td class="pl-8 pr-4 py-6">
                  <div class="flex items-center justify-center">
                    <input type="checkbox" v-model="doc.selected" class="w-5 h-5 rounded-md border-2 border-slate-200 text-primary-600 focus:ring-primary-500/20" />
                  </div>
                </td>
                <td class="px-6 py-6">
                  <div class="flex items-center gap-4">
                    <div :class="`w-12 h-12 rounded-xl flex items-center justify-center shadow-sm ${doc.iconBg}`">
                      <LucideFileText :class="`w-6 h-6 ${doc.iconColor}`" />
                    </div>
                    <div>
                      <p class="text-sm font-black text-slate-800 tracking-tight uppercase group-hover:text-primary-600 transition-colors cursor-pointer" @click="viewDocument(doc)">{{ doc.name }}</p>
                      <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-1">{{ doc.id }}</p>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-6">
                  <span :class="`inline-block px-3 py-1 rounded-lg text-[10px] font-black uppercase tracking-tight ${doc.typeColor}`">{{ doc.type }}</span>
                </td>
                <td class="px-6 py-6"><p class="text-xs font-bold text-slate-600">{{ doc.date }}</p></td>
                <td class="px-6 py-6"><p class="text-xs font-bold text-slate-500">{{ doc.dept }}</p></td>
                <td class="px-6 py-6">
                  <div class="flex items-center gap-2">
                    <div :class="`w-2 h-2 rounded-full ${doc.statusColor}`"></div>
                    <span class="text-xs font-bold text-slate-700">{{ doc.status }}</span>
                  </div>
                </td>
                <td class="pr-8 pl-4 py-6 text-right">
                  <div class="flex items-center justify-end gap-2">
                    <button @click="viewDocument(doc)" class="p-2.5 rounded-xl text-slate-400 hover:text-primary-500 hover:bg-primary-50 transition-all"><LucideEye class="w-5 h-5" /></button>
                    <button class="p-2.5 rounded-xl text-slate-400 hover:text-primary-500 hover:bg-primary-50 transition-all"><LucideShoppingCart class="w-5 h-5" /></button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Floating Cart Bar -->
      <Transition
        enter-active-class="transition duration-500 ease-out"
        enter-from-class="translate-y-20 opacity-0"
        enter-to-class="translate-y-0 opacity-100"
      >
        <div v-if="selectedCount > 0" class="absolute bottom-10 left-1/2 -translate-x-1/2 w-full max-w-lg">
          <div class="bg-[#1E3A5F] text-white rounded-2xl px-6 py-4 flex items-center justify-between shadow-2xl shadow-blue-900/40 border border-white/10 backdrop-blur-xl">
            <div class="flex items-center gap-4">
              <div class="relative">
                <LucideShoppingCart class="w-6 h-6 text-blue-200" />
                <span class="absolute -top-2 -right-2 w-5 h-5 bg-red-500 text-[10px] font-black flex items-center justify-center rounded-full border-2 border-[#1E3A5F]">{{ selectedCount }}</span>
              </div>
              <div>
                <p class="text-sm font-black tracking-tight">{{ $t('documents.cart.items_in_cart', { count: selectedCount }) }}</p>
                <p class="text-[10px] font-bold text-blue-300/80 uppercase tracking-widest">{{ $t('documents.cart.ready_desc') }}</p>
              </div>
            </div>
            <button @click="navigateTo('/loans/cart')" class="flex items-center gap-2 px-6 py-2.5 bg-blue-500 hover:bg-blue-400 text-white text-xs font-black rounded-xl transition-all group">
              {{ $t('documents.cart.btn_view') }} <LucideArrowRight class="w-4 h-4 group-hover:translate-x-1 transition-transform" />
            </button>
          </div>
        </div>
      </Transition>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { 
  LucideSearch, LucideChevronDown, LucideChevronRight, LucideFolder, LucideFolderOpen, LucideFileText, 
  LucideEye, LucideShoppingCart, LucideFilter, LucideArrowRight, LucideX
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'
import { formatDate } from '~/utils/format'

const { t } = useI18n()
const { $api } = useApi()
const route = useRoute()
const query = computed(() => route.query.q)

const searchResults = ref([])
const documents = ref([])
const isLoading = ref(false)

const getStatusStyles = (status) => {
  const s = status?.toLowerCase() || ''
  if (s === 'active' || s === 'available') return { label: 'Tersedia', color: 'bg-green-500' }
  if (s === 'pending' || s === 'processing') return { label: 'Diproses', color: 'bg-orange-500' }
  if (s === 'rejected') return { label: 'Ditolak', color: 'bg-red-500' }
  if (s === 'on_loan') return { label: 'Dipinjam', color: 'bg-blue-500' }
  if (s === 'draft') return { label: 'Draf', color: 'bg-slate-400' }
  return { label: status, color: 'bg-slate-300' }
}

const getTypeStyles = (mimeType) => {
  const m = mimeType?.toLowerCase() || ''
  if (m.includes('pdf')) return { bg: 'bg-red-50', color: 'text-red-500', typeBg: 'bg-red-50 text-red-600 border border-red-100' }
  if (m.includes('image')) return { bg: 'bg-blue-50', color: 'text-blue-500', typeBg: 'bg-blue-50 text-blue-600 border border-blue-100' }
  if (m.includes('excel') || m.includes('sheet')) return { bg: 'bg-green-50', color: 'text-green-500', typeBg: 'bg-green-50 text-green-600 border border-green-100' }
  return { bg: 'bg-slate-50', color: 'text-slate-500', typeBg: 'bg-slate-50 text-slate-600 border border-slate-100' }
}

const fetchDocuments = async () => {
  isLoading.value = true
  try {
    const res = await $api('/documents?limit=50')
    if (res && res.data) {
      documents.value = res.data.map(doc => {
        const statusStyle = getStatusStyles(doc.status)
        const typeStyle = getTypeStyles(doc.mime_type)
        return {
          id: doc.id.substring(0, 8).toUpperCase(),
          original_id: doc.id,
          name: doc.title,
          type: doc.type_name || 'Dokumen',
          typeColor: typeStyle.typeBg,
          date: formatDate(doc.created_at),
          dept: doc.department_name || '-',
          status: statusStyle.label,
          statusColor: statusStyle.color,
          iconBg: typeStyle.bg,
          iconColor: typeStyle.color,
          selected: false
        }
      })
    }
  } catch (err) {
    console.error('Failed to fetch documents:', err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  fetchDocuments()
})

const updateSearchResults = (q) => {
  if (!q) {
    searchResults.value = []
    return
  }
  
  if (q.toLowerCase() === 'kontrak cleaning') {
    searchResults.value = [
      {
        id: 'LGL-AKR-001-2026-005',
        name: t('documents.mock.search.result_1_name'),
        path: `${t('documents.mock.root_folder')} > ${t('documents.mock.legal')} > 2026 > ${t('documents.mock.docs.type_contract')}`,
        snippet: t('documents.mock.search.result_1_snippet'),
        tags: [t('documents.mock.legal'), t('documents.mock.docs.type_contract'), '2026']
      },
      {
        id: 'LGL-AKR-2025-015',
        name: t('documents.mock.search.result_2_name'),
        path: `${t('documents.mock.root_folder')} > ${t('documents.mock.legal')} > 2025 > ADDENDUM`,
        snippet: t('documents.mock.search.result_2_snippet'),
        tags: [t('documents.mock.legal'), 'ADDENDUM', '2025']
      },
      {
        id: 'HR-SOP-2026-001',
        name: t('documents.mock.search.result_3_name'),
        path: `${t('documents.mock.root_folder')} > GA / HR > 2026 > SOP`,
        snippet: t('documents.mock.search.result_3_snippet'),
        tags: ['HR', 'SOP', '2026']
      }
    ]
  } else {
    searchResults.value = []
  }
}

watch(query, (newVal) => {
  updateSearchResults(newVal)
}, { immediate: true })

const folderTree = ref([
  {
    id: 1,
    name: 'SEMUA ARSIP',
    expanded: true,
    children: [
      {
        id: 2,
        name: 'DEPARTEMEN',
        expanded: true,
        count: documents.value.length,
        children: []
      }
    ]
  }
])

const selectedCount = computed(() => documents.value.filter(d => d.selected).length)

const toggleFolder = (node) => {
  if (node.children) {
    node.expanded = !node.expanded
  }
}

const viewDocument = (doc) => {
  navigateTo(`/documents/${doc.original_id || doc.id}`)
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: #CBD5E1; }
</style>


<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #E2E8F0;
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #CBD5E1;
}
</style>
