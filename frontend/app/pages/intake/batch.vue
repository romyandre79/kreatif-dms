<template>
  <div class="max-w-7xl mx-auto space-y-10 pb-20 relative min-h-screen" v-motion-fade>
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div class="space-y-2">
        <div class="flex items-center gap-4">
          <button @click="navigateTo('/intake/staging')" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all">
            <LucideArrowLeft class="w-6 h-6 text-slate-400" />
          </button>
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Batch Digitalisasi</h1>
        </div>
        <p class="text-slate-500 font-medium italic ml-12">Memproses {{ manifestIds.length }} manifest secara kolektif</p>
      </div>
    </div>

    <!-- Batch List -->
    <div class="glass rounded-[32px] overflow-hidden border border-slate-100 dark:border-slate-800 shadow-2xl shadow-slate-200/20">
      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-900/50 border-b border-slate-100 dark:border-slate-800 text-[10px] font-black text-slate-400 uppercase tracking-widest">
              <th class="px-8 py-6 w-16 text-center">No</th>
              <th class="px-8 py-6">Manifest ID</th>
              <th class="px-8 py-6">Pemohon</th>
              <th class="px-8 py-6">Departemen</th>
              <th class="px-8 py-6 text-center">Dokumen</th>
              <th class="px-8 py-6 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr v-if="pending" class="animate-pulse">
              <td colspan="6" class="px-8 py-20 text-center text-slate-400 font-bold uppercase tracking-widest">Loading batch data...</td>
            </tr>
            <tr 
              v-for="(m, index) in batchManifests" 
              :key="m.id" 
              @click="openDetail(m)"
              class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-all cursor-pointer"
            >
              <td class="px-8 py-6 text-center text-xs font-black text-slate-300">
                {{ index + 1 }}
              </td>
              <td class="px-8 py-6">
                <p class="font-black text-sm uppercase tracking-tight text-[#1E3A5F] dark:text-slate-200 group-hover:text-primary-500 transition-colors">
                  {{ m.manifest_no }}
                </p>
              </td>
              <td class="px-8 py-6">
                <p class="font-black text-sm text-[#1E3A5F] dark:text-slate-300">{{ m.owner_name }}</p>
              </td>
              <td class="px-8 py-6">
                <p class="font-bold text-xs text-slate-500">{{ m.department_name }}</p>
              </td>
              <td class="px-8 py-6 text-center">
                <p class="text-lg font-black text-[#1E3A5F] dark:text-white">{{ m.total_items }}</p>
              </td>
              <td class="px-8 py-6 text-right">
                <button class="p-2 text-slate-300 group-hover:text-primary-500 transition-all">
                  <LucideChevronRight class="w-5 h-5" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Detail Drawer (Slider) -->
    <Transition name="drawer">
      <div v-if="selectedManifest" class="fixed inset-0 z-[100] flex justify-end">
        <!-- Backdrop -->
        <div class="absolute inset-0 bg-slate-900/40 backdrop-blur-sm" @click="selectedManifest = null"></div>
        
        <!-- Drawer Content -->
        <div class="relative w-full max-w-md bg-white dark:bg-slate-950 h-full shadow-2xl flex flex-col overflow-hidden border-l border-white/10">
          <!-- Drawer Header -->
          <div class="p-8 border-b border-slate-50 dark:border-slate-900 flex items-center justify-between">
            <div class="space-y-1">
              <h2 class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight">Detail Manifest</h2>
              <div class="flex items-center gap-3">
                <span class="text-lg font-black text-orange-500 uppercase">{{ selectedManifest.manifest_no }}</span>
                <span class="px-2 py-0.5 bg-blue-500 text-white text-[9px] font-black rounded uppercase tracking-widest">{{ selectedManifest.source }}</span>
              </div>
            </div>
            <button @click="selectedManifest = null" class="w-10 h-10 flex items-center justify-center rounded-xl bg-slate-50 dark:bg-slate-900 text-slate-400 hover:text-slate-600 transition-all">
              <LucideX class="w-5 h-5" />
            </button>
          </div>

          <!-- Drawer Body -->
          <div class="flex-1 overflow-y-auto p-8 space-y-10 custom-scrollbar">
            <!-- Info Manifest -->
            <section class="space-y-6">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-300 uppercase tracking-widest">
                <LucideInfo class="w-3 h-3" />
                <span>Info Manifest</span>
              </div>
              <div class="grid grid-cols-1 gap-4">
                <div v-for="(val, label) in manifestInfo" :key="label" class="flex justify-between items-center py-2 border-b border-slate-50 dark:border-slate-900">
                  <span class="text-xs font-bold text-slate-400">{{ label }}</span>
                  <div class="flex items-center gap-2">
                    <div v-if="label === 'Prioritas'" class="w-2 h-2 rounded-full bg-red-500"></div>
                    <span class="text-sm font-black text-[#1E3A5F] dark:text-slate-200">{{ val }}</span>
                  </div>
                </div>
              </div>
            </section>

            <!-- Dokumen Dalam Manifest -->
            <section class="space-y-6">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-300 uppercase tracking-widest">
                <LucideFileText class="w-3 h-3" />
                <span>Dokumen Dalam Manifest</span>
              </div>
              <div class="space-y-4">
                <div v-for="doc in manifestDetail?.items" :key="doc.id" class="p-4 rounded-2xl border border-slate-100 dark:border-slate-900 flex items-center gap-4 hover:border-primary-500/30 transition-all group">
                  <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-900 flex items-center justify-center text-slate-400">
                    <LucideFileText class="w-5 h-5" />
                  </div>
                  <div class="flex-1 space-y-0.5">
                    <p class="text-xs font-black text-[#1E3A5F] dark:text-slate-200 uppercase tracking-tight line-clamp-1">{{ doc.title }}</p>
                    <p :class="['text-[9px] font-black uppercase tracking-widest', doc.status === 'received' ? 'text-emerald-500' : 'text-amber-500']">
                      {{ doc.status === 'received' ? 'Kondisi Baik' : 'Perlu Perbaikan Fisik' }}
                    </p>
                  </div>
                  <LucideExternalLink class="w-4 h-4 text-slate-200 group-hover:text-primary-500 transition-all" />
                </div>
              </div>
            </section>

            <!-- Timeline -->
            <section class="space-y-6">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-300 uppercase tracking-widest">
                <LucideActivity class="w-3 h-3" />
                <span>Timeline</span>
              </div>
              <div class="space-y-8 pl-4 relative">
                <div class="absolute left-[19px] top-2 bottom-2 w-0.5 bg-slate-100 dark:bg-slate-900"></div>
                <div v-for="(step, i) in manifestDetail?.timeline" :key="i" class="relative flex gap-6 items-start">
                  <div :class="['w-3 h-3 rounded-full mt-1.5 z-10 outline outline-4', i === 0 ? 'bg-orange-500 outline-orange-500/20' : 'bg-slate-200 outline-white dark:outline-slate-950']"></div>
                  <div class="space-y-1">
                    <p :class="['text-xs font-black uppercase tracking-tight', i === 0 ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400']">{{ step.status }}</p>
                    <p class="text-[10px] font-bold text-slate-400">{{ step.timestamp }}</p>
                  </div>
                </div>
              </div>
            </section>
          </div>

          <!-- Drawer Footer -->
          <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-900">
            <button 
              @click="navigateTo({ path: '/intake/indexing', query: { manifest_no: selectedManifest.manifest_no } })"
              class="w-full py-4 bg-emerald-500 hover:bg-emerald-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-lg shadow-emerald-500/20 transition-all flex items-center justify-center gap-3"
            >
              <LucideScanLine class="w-4 h-4" />
              Digitalisasi Sekarang
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { 
  LucideArrowLeft, LucideChevronRight, LucideX, LucideInfo, 
  LucideFileText, LucideActivity, LucideExternalLink, LucideScanLine
} from 'lucide-vue-next'

const route = useRoute()
const { $api } = useApi()

const manifestIds = computed(() => {
  const ids = route.query.ids
  return ids ? ids.split(',') : []
})

const batchManifests = ref([])
const pending = ref(true)

// Fetch all selected manifests
onMounted(async () => {
  try {
    const promises = manifestIds.value.map(id => $api(`/intake/manifest/${id}`))
    const results = await Promise.all(promises)
    batchManifests.value = results.map(r => r.data)
  } catch (e) {
    console.error('Failed to fetch batch data', e)
  } finally {
    pending.value = false
  }
})

const selectedManifest = ref(null)
const manifestDetail = ref(null)

const openDetail = async (m) => {
  selectedManifest.value = m
  // We already have some data, but let's re-fetch to get full items/timeline if needed
  // In our case, GetManifest already returns full detail
  manifestDetail.value = m
}

const manifestInfo = computed(() => {
  if (!selectedManifest.value) return {}
  return {
    'Pemohon': selectedManifest.value.owner_name,
    'Departemen': selectedManifest.value.department_name,
    'Prioritas': selectedManifest.value.priority || 'Normal',
    'Estimasi': selectedManifest.value.estimate || `${selectedManifest.value.total_items} Dokumen`
  }
})
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

.drawer-enter-active,
.drawer-leave-active {
  transition: all 0.5s cubic-bezier(0.16, 1, 0.3, 1);
}

.drawer-enter-from,
.drawer-leave-to {
  transform: translateX(100%);
  opacity: 0;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #e2e8f0;
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: #1e293b;
}
</style>
