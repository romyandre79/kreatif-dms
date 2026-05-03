<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-32">
    <!-- Header -->
    <div class="flex flex-col gap-2" v-motion-fade>
      <h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tight">
        {{ $t('registration.staging.title') }}
      </h1>
      <p class="text-slate-500 font-medium">
        {{ $t('registration.staging.subtitle') }}
      </p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6" v-motion-slide-visible-bottom>
      <div v-for="(stat, i) in stats" :key="i" class="glass p-8 rounded-[2rem] relative overflow-hidden group">
        <div class="flex items-start justify-between">
          <div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">{{ $t(`registration.staging.stats.${stat.key}`) }}</p>
            <div class="flex items-baseline gap-2">
              <span class="text-4xl font-black text-[#1E3A5F] dark:text-white">{{ stat.value }}</span>
              <span class="text-xs font-bold text-green-500">{{ $t('registration.staging.stats.trend', { value: stat.trend }) }}</span>
            </div>
          </div>
          <div :class="`w-10 h-10 rounded-xl ${stat.bg} flex items-center justify-center ${stat.color} shadow-sm group-hover:scale-110 transition-transform`">
            <component :is="stat.icon" class="w-5 h-5" />
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <div class="lg:col-span-2 space-y-8">
        <!-- Upload Zone -->
        <div class="glass rounded-[2.5rem] p-12 border-2 border-dashed border-slate-200 dark:border-slate-800 flex flex-col items-center text-center group hover:border-primary-500/50 hover:bg-primary-50/5 transition-all cursor-pointer" v-motion-slide-visible-bottom>
          <div class="w-20 h-20 rounded-[2rem] bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center mb-6 group-hover:scale-110 transition-transform shadow-inner">
            <LucideFileUp class="w-10 h-10 text-[#1E3A5F] dark:text-primary-400" />
          </div>
          <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white mb-2 uppercase tracking-tight">
            {{ $t('registration.staging.upload.title') }}
          </h2>
          <p class="text-sm font-bold text-slate-400 mb-8 max-w-sm leading-relaxed">
            {{ $t('registration.staging.upload.subtitle') }}
          </p>
          <div class="flex items-center gap-4">
            <button class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3">
              <LucidePlus class="w-4 h-4" />
              {{ $t('registration.staging.upload.btn_choose') }}
            </button>
            <button class="px-10 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-xs font-black uppercase tracking-widest text-[#1E3A5F] dark:text-white hover:bg-slate-50 transition-all flex items-center gap-3">
              <LucideFolderOpen class="w-4 h-4" />
              {{ $t('registration.staging.upload.btn_batch') }}
            </button>
          </div>
        </div>

        <!-- Table -->
        <div class="glass rounded-[2.5rem] overflow-hidden shadow-xl shadow-slate-200/50 dark:shadow-none" v-motion-slide-visible-bottom>
          <div class="px-10 py-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-900/50">
            <h3 class="font-black text-xl text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('registration.staging.table.title') }}</h3>
            <div class="flex items-center gap-4">
              <div class="relative">
                <select class="pl-5 pr-12 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-[10px] font-black uppercase tracking-widest appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 transition-all cursor-pointer">
                  <option>{{ $t('registration.staging.table.sort_by') }}</option>
                </select>
                <LucideChevronDown class="absolute right-4 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-slate-400 pointer-events-none" />
              </div>
              <button class="p-2.5 text-slate-400 hover:text-[#1E3A5F] transition-colors bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl shadow-sm">
                <LucideRefreshCcw class="w-4 h-4" />
              </button>
            </div>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                  <th class="p-8 px-10">{{ $t('registration.staging.table.cols.filename') }}</th>
                  <th class="p-8">{{ $t('registration.staging.table.cols.pages') }}</th>
                  <th class="p-8">{{ $t('registration.staging.table.cols.status') }}</th>
                  <th class="p-8 text-right px-10">{{ $t('registration.staging.table.cols.operator') }}</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                <tr v-for="doc in documents" :key="doc.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-colors cursor-pointer">
                  <td class="p-8 px-10">
                    <div class="flex items-center gap-5">
                      <div :class="`w-12 h-12 rounded-xl flex items-center justify-center ${doc.iconBg} ${doc.iconColor} group-hover:scale-110 transition-transform shadow-sm`">
                        <component :is="doc.icon" class="w-6 h-6" />
                      </div>
                      <div>
                        <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight group-hover:text-primary-500 transition-colors">{{ doc.filename }}</p>
                        <p class="text-[10px] font-bold text-slate-400 mt-1">{{ doc.size }} • {{ doc.time }}</p>
                      </div>
                    </div>
                  </td>
                  <td class="p-8">
                    <span class="text-xs font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest">{{ doc.pages }} {{ $t('registration.staging.table.cols.pages').toLowerCase() }}</span>
                  </td>
                  <td class="p-8">
                    <div v-if="doc.ocrStatus === 'READY'" class="flex items-center gap-2 px-3 py-1 bg-green-50 dark:bg-green-900/20 text-green-500 rounded-full w-fit">
                      <div class="w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse"></div>
                      <span class="text-[10px] font-black uppercase tracking-widest">{{ $t('registration.staging.table.status.ready') }}</span>
                    </div>
                    <div v-else-if="doc.ocrStatus === 'QUEUED'" class="flex items-center gap-2 px-3 py-1 bg-slate-100 dark:bg-slate-800 text-slate-400 rounded-full w-fit">
                      <div class="w-1.5 h-1.5 rounded-full bg-slate-400"></div>
                      <span class="text-[10px] font-black uppercase tracking-widest">{{ $t('registration.staging.table.status.queued') }}</span>
                    </div>
                    <div v-else class="space-y-2 max-w-[120px]">
                      <div class="flex justify-between items-center text-[10px] font-black text-slate-500 uppercase tracking-widest">
                        <span>{{ doc.progress }}%</span>
                      </div>
                      <div class="w-full h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                        <div :style="{ width: doc.progress + '%' }" class="h-full bg-[#1E3A5F] dark:bg-primary-500 rounded-full"></div>
                      </div>
                    </div>
                  </td>
                  <td class="p-8 text-right px-10">
                    <div v-if="doc.operator" class="flex items-center justify-end gap-3">
                      <div class="text-right">
                        <p class="text-[10px] font-black text-slate-700 dark:text-slate-200 uppercase tracking-tighter">{{ doc.operator.name }}</p>
                        <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ doc.operator.role }}</p>
                      </div>
                      <div class="w-9 h-9 rounded-full bg-slate-200 overflow-hidden ring-2 ring-white dark:ring-slate-800 shadow-sm">
                        <img :src="doc.operator.avatar" class="w-full h-full object-cover" />
                      </div>
                    </div>
                    <span v-else class="text-[9px] font-black text-slate-400 uppercase tracking-widest italic">Auto-assigning...</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="px-10 py-6 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
              {{ $t('registration.staging.table.footer', { current: 4, total: 124 }) }}
            </p>
            <div class="flex items-center gap-2">
              <button v-for="p in 3" :key="p" :class="`w-8 h-8 rounded-lg flex items-center justify-center text-[10px] font-black transition-all ${p === 1 ? 'bg-[#1E3A5F] text-white shadow-lg shadow-blue-900/20' : 'text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800'}`">
                {{ p }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Sidebar -->
      <div class="space-y-8">
        <!-- Priority Card -->
        <div class="glass p-10 rounded-[2.5rem] relative overflow-hidden group" v-motion-slide-visible-bottom>
          <div class="absolute top-0 right-0 p-8 opacity-5">
            <LucideZap class="w-24 h-24 text-[#1E3A5F]" />
          </div>
          <div class="flex items-center gap-4 mb-8">
            <div class="w-12 h-12 rounded-2xl bg-orange-50 dark:bg-orange-900/20 flex items-center justify-center text-orange-500 shadow-sm ring-1 ring-orange-100 dark:ring-orange-900/50">
              <LucideBell class="w-6 h-6 animate-bounce" />
            </div>
            <div>
              <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.staging.priority.title') }}</h3>
              <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('registration.staging.priority.subtitle') }}</p>
            </div>
          </div>
          <p class="text-sm font-bold text-slate-500 leading-relaxed mb-10">
            {{ $t('registration.staging.priority.desc', { count: 12, status: $t('registration.staging.priority.status_ready') }) }}
          </p>
          <ul class="space-y-6 mb-10">
            <li v-for="(feat, i) in features" :key="i" class="flex items-center gap-4 group/item">
              <div :class="`w-5 h-5 rounded-full flex items-center justify-center ${feat.bg} ${feat.color} transition-all group-hover/item:scale-110`">
                <component :is="feat.icon" class="w-3 h-3" />
              </div>
              <span class="text-[10px] font-black text-slate-500 uppercase tracking-widest">{{ $t(`registration.staging.priority.features.${feat.key}`) }}</span>
            </li>
          </ul>
          <button class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-[1.5rem] text-xs font-black uppercase tracking-widest transition-all shadow-2xl shadow-blue-900/40 flex items-center justify-center gap-3 active:scale-95 group">
            <LucidePlay class="w-4 h-4 fill-white group-hover:scale-110 transition-transform" />
            {{ $t('registration.staging.priority.btn_start') }}
          </button>
        </div>

        <!-- Engine Status -->
        <div class="glass p-8 rounded-[2rem] space-y-8" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-xl bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center text-[#1E3A5F] dark:text-primary-400 shadow-sm">
                <LucideCpu class="w-6 h-6" />
              </div>
              <div>
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('registration.staging.engine.ocr_title') }}</p>
                <p class="text-xs font-black text-slate-700 dark:text-white uppercase tracking-tight">ABBYY Finereader</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-[10px] font-black text-green-500 uppercase tracking-widest">{{ $t('registration.staging.engine.status_active') }} (99.2%)</p>
            </div>
          </div>
          <div class="w-full h-px bg-slate-100 dark:bg-slate-800"></div>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-xl bg-purple-50 dark:bg-purple-900/20 flex items-center justify-center text-purple-500 shadow-sm">
                <LucideDatabase class="w-6 h-6" />
              </div>
              <div>
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('registration.staging.engine.repo_title') }}</p>
                <p class="text-xs font-black text-slate-700 dark:text-white uppercase tracking-tight">Cloud Instance</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-[10px] font-black text-green-500 uppercase tracking-widest">{{ $t('registration.staging.engine.status_sync') }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideFileUp, LucidePlus, LucideFolderOpen, LucideChevronDown, LucideRefreshCcw,
  LucideFileText, LucideFileArchive, LucideFileImage, LucideZap, LucideBell, 
  LucideCheckCircle2, LucideLayoutGrid, LucideHistory, LucideCpu, LucideDatabase,
  LucidePlay, LucideClock, LucideFileSearch
} from 'lucide-vue-next'

const stats = [
  { key: 'queued', value: '124', trend: '+12%', icon: LucideClock, bg: 'bg-blue-50 dark:bg-blue-900/20', color: 'text-blue-500' },
  { key: 'processing', value: '12', trend: '+5%', icon: LucideRefreshCcw, bg: 'bg-purple-50 dark:bg-purple-900/20', color: 'text-purple-500' },
  { key: 'ready', value: '85', trend: '+8%', icon: LucideCheckCircle2, bg: 'bg-green-50 dark:bg-green-900/20', color: 'text-green-500' }
]

const features = [
  { key: 'auto_class', icon: LucideCheckCircle2, bg: 'bg-green-50', color: 'text-green-500' },
  { key: 'ai_meta', icon: LucideCheckCircle2, bg: 'bg-green-50', color: 'text-green-500' },
  { key: 'incoming', icon: LucideRefreshCcw, bg: 'bg-slate-50', color: 'text-slate-400' }
]

const documents = [
  {
    id: 1,
    filename: 'Legal_Contract_2023_v2.pdf',
    size: '14.2 MB',
    time: '2 mins ago',
    pages: '12',
    ocrStatus: 'PROCESSING',
    progress: 65,
    icon: LucideFileText,
    iconBg: 'bg-red-50',
    iconColor: 'text-red-500',
    operator: {
      name: 'Budi Santoso',
      role: 'Junior Staff',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Budi'
    }
  },
  {
    id: 2,
    filename: 'Invoices_Batch_Oct.tiff',
    size: '45.8 MB',
    time: '5 mins ago',
    pages: '48',
    ocrStatus: 'READY',
    icon: LucideFileArchive,
    iconBg: 'bg-blue-50',
    iconColor: 'text-blue-500',
    operator: {
      name: 'Siska Amelia',
      role: 'Senior Admin',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Siska'
    }
  },
  {
    id: 3,
    filename: 'Identity_Card_Scan.jpg',
    size: '1.2 MB',
    time: '10 mins ago',
    pages: '1',
    ocrStatus: 'QUEUED',
    icon: LucideFileImage,
    iconBg: 'bg-orange-50',
    iconColor: 'text-orange-500',
    operator: null
  }
]
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>

