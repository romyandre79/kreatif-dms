<template>
  <div class="flex h-[calc(100vh-theme(spacing.20)-1rem)] -m-2 overflow-hidden bg-[#F8FAFC]">
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
        <div v-if="isLoadingTree" class="p-4 space-y-4">
          <div v-for="i in 5" :key="i" class="h-8 bg-slate-50 animate-pulse rounded-lg"></div>
        </div>
        <div v-else class="space-y-1">
          <ExplorerNode 
            v-for="node in folderTree" 
            :key="node.id" 
            :node="node" 
            @select="handleNodeSelect"
          />
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
          <template v-if="breadcrumbs.length">
            <template v-for="(crumb, i) in breadcrumbs" :key="crumb.id">
              <button 
                @click="handleBreadcrumbClick(crumb, i)"
                class="hover:text-primary-600 transition-colors uppercase"
                :class="{ 'text-primary-600': i === breadcrumbs.length - 1 }"
              >
                {{ crumb.name }}
              </button>
              <LucideChevronRight v-if="i < breadcrumbs.length - 1" class="w-3 h-3" />
            </template>
          </template>
          <span v-else>{{ $t('documents.explorer.root') }}</span>
        </div>
        <div class="text-xs font-bold text-slate-400">
          {{ $t('documents.explorer.stats_total', { count: selectedNode?.count || documents.length }) }}
        </div>
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
      <div class="flex-1 overflow-y-auto p-2 custom-scrollbar">
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
                    <input 
                      type="checkbox" 
                      v-model="isAllSelected"
                      class="w-5 h-5 rounded-md border-2 border-slate-200 text-primary-600 focus:ring-primary-500/20" 
                    />
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
                <td class="pl-8 py-6">
                  <input 
                    type="checkbox" 
                    :checked="cartStore.isInCart(doc.original_id || doc.id)"
                    @change="toggleDocSelection(doc)"
                    :disabled="!isDocSelectable(doc)"
                    class="w-5 h-5 rounded-md border-2 border-slate-200 text-primary-600 focus:ring-primary-500/20 disabled:opacity-30 disabled:cursor-not-allowed" 
                  />
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
                    <div :class="`w-2 h-2 rounded-full ${isDocRequested(doc) ? 'bg-amber-500' : doc.statusColor}`"></div>
                    <span class="text-xs font-bold" :class="isDocRequested(doc) ? 'text-amber-600 font-bold' : 'text-slate-700'">
                      {{ isDocRequested(doc) ? 'Requested' : doc.status }}
                    </span>
                  </div>
                </td>
                <td class="pr-8 pl-4 py-6 text-right">
                  <div class="flex items-center justify-end gap-2">
                    <button @click="viewDocument(doc)" class="p-2.5 rounded-xl text-slate-400 hover:text-primary-500 hover:bg-primary-50 transition-all">
                      <LucideEye class="w-5 h-5" />
                    </button>
                    <button 
                      @click.stop="handleQuickLoan(doc)"
                      :disabled="!isDocSelectable(doc)"
                      class="p-2.5 rounded-xl transition-all"
                      :class="[
                        doc.selected ? 'text-primary-600 bg-primary-50 ring-1 ring-primary-100' : 
                        (!isDocSelectable(doc))
                          ? 'text-slate-200 cursor-not-allowed bg-slate-50/50'
                          : 'text-slate-400 hover:text-primary-500 hover:bg-primary-50'
                      ]"
                    >
                      <LucideShoppingCart class="w-5 h-5" :class="{ 'opacity-50': !isDocSelectable(doc) }" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

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
import ExplorerNode from '~/components/ExplorerNode.vue'
import { useCartStore } from '~/stores/cart'

const { t } = useI18n()
const { $api } = useApi()
const route = useRoute()
const cartStore = useCartStore()
const query = computed(() => route.query.q)
const selectedCount = computed(() => cartStore.count)

const searchResults = ref([])
const documents = ref([])
const isLoading = ref(false)
const isLoadingTree = ref(false)
const folderTree = ref([])
const selectedNode = ref(null)
const breadcrumbs = ref([])

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

const fetchDocuments = async (filters = {}) => {
  isLoading.value = true
  try {
    let url = '/documents?limit=50'
    Object.keys(filters).forEach(key => {
      url += `&${key}=${filters[key]}`
    })
    
    const res = await $api(url)
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
          rawStatus: doc.status,
          physicalStatus: doc.physical_status,
          statusColor: statusStyle.color,
          iconBg: typeStyle.bg,
          iconColor: typeStyle.color,
          company_name: doc.company_name?.String || doc.company_name || doc.company || '',
          branch_name: doc.branch_name?.String || doc.branch_name || doc.branch || '',
          selected: cartStore.isInCart(doc.id)
        }
      })
    }
  } catch (err) {
    console.error('Failed to fetch documents:', err)
  } finally {
    isLoading.value = false
  }
}

const fetchExplorerTree = async () => {
  isLoadingTree.value = true
  try {
    const res = await $api('/documents/explorer/tree')
    if (res && res.data) {
      folderTree.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch explorer tree:', err)
  } finally {
    isLoadingTree.value = false
  }
}

const handleNodeSelect = ({ node, path }) => {
  // Clear active status from all nodes
  const clearActive = (nodes) => {
    nodes.forEach(n => {
      n.active = false
      if (n.children) clearActive(n.children)
    })
  }
  clearActive(folderTree.value)
  
  node.active = true
  selectedNode.value = node
  breadcrumbs.value = path
  
  // Build filters based on ALL nodes in the path
  const filters = {}
  path.forEach(item => {
    if (item.type === 'department') filters.department_id = item.id
    else if (item.type === 'company') filters.company_id = item.id
    else if (item.type === 'branch') filters.branch_id = item.id
    else if (item.type === 'rack') filters.rack_id = item.id
    else if (item.type === 'box') filters.box_id = item.id
    else if (item.type === 'ordner') filters.ordner_id = item.id
    else if (item.type === 'year') filters.year = item.id.replace('year-', '')
    else if (item.type === 'type') filters.type_id = item.id
  })
  
  fetchDocuments(filters)
}

const handleBreadcrumbClick = (crumb, index) => {
  const newPath = breadcrumbs.value.slice(0, index + 1)
  handleNodeSelect({ node: crumb.node, path: newPath })
}

onMounted(() => {
  fetchDocuments()
  fetchExplorerTree()
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

const isAllSelected = computed({
  get: () => documents.value.length > 0 && documents.value.every(doc => cartStore.isInCart(doc.original_id || doc.id)),
  set: (val) => {
    documents.value.forEach(doc => {
      if (val) {
        cartStore.addItem(doc)
      } else {
        cartStore.removeItem(doc.original_id || doc.id)
      }
    })
  }
})

const toggleFolder = (node) => {
  if (node.children) {
    node.expanded = !node.expanded
  }
}

const viewDocument = (doc) => {
  navigateTo(`/documents/${doc.original_id || doc.id}`)
}

const handleQuickLoan = (doc) => {
  if (cartStore.count > 0) {
    cartStore.toggleItem(doc)
    doc.selected = cartStore.isInCart(doc.id)
  } else {
    cartStore.addItem(doc)
    navigateTo('/loans/cart')
  }
}

const toggleDocSelection = (doc) => {
  cartStore.toggleItem(doc)
  doc.selected = cartStore.isInCart(doc.id)
}

const isDocRequested = (doc) => {
  return cartStore.requestedLoanDocIds?.includes(doc.original_id || doc.id)
}

const isDocSelectable = (doc) => {
  if (isDocRequested(doc)) return false
  return !(doc.rawStatus === 'on_loan' || doc.physicalStatus === 'damaged' || doc.physicalStatus === 'missing')
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
