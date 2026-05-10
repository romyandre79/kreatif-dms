<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <p class="text-[10px] font-bold text-blue-500 uppercase tracking-widest">{{ $t('rfid.model_f16.breadcrumb') }}</p>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('rfid.model_f16.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('rfid.model_f16.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-8 py-3.5 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all flex items-center gap-3">
          <LucidePrinter class="w-4 h-4" />
          {{ $t('rfid.model_f16.header.print') }}
        </button>
        <button class="px-10 py-3.5 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-4 active:scale-95 group">
          <LucideSave class="w-4 h-4 group-hover:scale-110 transition-transform" />
          {{ $t('rfid.model_f16.header.save') }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Container Preview -->
      <div class="lg:col-span-4 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="flex items-center justify-between">
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.container.title') }}</h3>
            <span class="px-3 py-1 bg-blue-50 text-blue-500 rounded text-[8px] font-black uppercase tracking-widest">{{ $t('rfid.model_f16.container.parent') }}</span>
          </div>

          <div class="aspect-square bg-slate-50/50 dark:bg-slate-900/50 rounded-[3rem] border-4 border-dashed border-slate-100 dark:border-slate-800 flex flex-col items-center justify-center gap-6 group cursor-pointer hover:bg-slate-100/50 transition-all">
            <div class="w-32 h-32 bg-white dark:bg-slate-800 rounded-lg shadow-xl flex items-center justify-center text-slate-100 dark:text-slate-700 group-hover:scale-110 transition-transform">
              <LucideBox class="w-16 h-16" />
            </div>
            <p class="text-xs font-bold text-slate-300 uppercase tracking-widest">{{ $t('rfid.model_f16.container.drop') }}</p>
          </div>

          <div class="space-y-6">
            <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-4 relative overflow-hidden group">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.container.id') }}</label>
              <div class="flex items-center justify-between">
                <p class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight">BX-2023-F16-0922</p>
                <LucideQrCode class="w-6 h-6 text-blue-500/50 group-hover:text-blue-500 transition-colors" />
              </div>
            </div>

            <div class="grid grid-cols-2 gap-6">
              <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-2">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.container.qr') }}</p>
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 bg-green-500 rounded-full"></div>
                  <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('rfid.model_f16.container.linked') }}</p>
                </div>
              </div>
              <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-2">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.container.tag') }}</p>
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 bg-green-500 rounded-full"></div>
                  <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('rfid.model_f16.container.active') }}</p>
                </div>
              </div>
            </div>
          </div>

          <div class="p-8 bg-blue-50/50 dark:bg-blue-900/10 border-l-4 border-blue-500 rounded-2xl">
            <p class="text-[10px] font-bold text-blue-600 dark:text-blue-400 leading-relaxed uppercase tracking-tight">
              {{ $t('rfid.model_f16.container.behavior') }}
            </p>
          </div>
        </div>
      </div>

      <!-- Center Column: Child Items -->
      <div class="lg:col-span-5 space-y-10">
        <div class="glass p-10 rounded-[3.5rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800 flex flex-col h-full">
          <div class="flex items-center justify-between">
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.items.title') }}</h3>
            <span class="px-4 py-1 bg-slate-950 text-white rounded-lg text-[9px] font-black uppercase tracking-widest">{{ $t('rfid.model_f16.items.count', { count: 12 }) }}</span>
          </div>

          <div class="flex-grow space-y-6 overflow-y-auto pr-2 custom-scrollbar">
            <div v-for="item in [
              { title: 'Sertifikat Tanah Hak Milik', ref: 'SHM/JKT-09122', id: '0x923...EF', status: 'ready', icon: LucideFileText },
              { title: 'Berkas Pajak Tahunan 2022', ref: 'TAX/CORP/0023', id: '0x861...AB', status: 'ready', icon: LucideBriefcase },
              { title: 'Akta Pendirian Cabang', ref: 'LAW/EST/2021', status: 'no_tag', icon: LucideFileText }
            ]" :key="item.title" class="p-8 bg-white dark:bg-slate-800/50 border border-slate-100 dark:border-slate-700 rounded-3xl shadow-sm hover:shadow-xl hover:scale-[1.02] transition-all cursor-move group flex items-center justify-between">
              <div class="flex items-center gap-6">
                <div :class="`w-14 h-14 rounded-2xl flex items-center justify-center transition-colors ${item.status === 'ready' ? 'bg-slate-50 text-[#1E3A5F] dark:bg-slate-900 dark:text-white' : 'bg-red-50 text-red-500'}`">
                  <component :is="item.icon" class="w-6 h-6" />
                </div>
                <div class="space-y-1">
                  <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ item.title }}</h4>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">REF: {{ item.ref }}</p>
                </div>
              </div>
              <div class="flex flex-col items-end gap-2">
                <LucideGripVertical class="w-5 h-5 text-slate-200 group-hover:text-slate-400 transition-colors mb-2" />
                <div class="flex items-center gap-3">
                  <span :class="`text-[8px] font-black uppercase tracking-widest flex items-center gap-1.5 ${item.status === 'ready' ? 'text-green-500' : 'text-red-500'}`">
                    <LucideRss v-if="item.status === 'ready'" class="w-3 h-3" />
                    {{ $t(`rfid.model_f16.items.${item.status}`) }}
                  </span>
                  <span v-if="item.id" class="text-[9px] font-bold text-slate-300 uppercase font-mono tracking-tighter">ID: {{ item.id }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Summary & Network -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.summary.title') }}</h3>
          
          <div class="space-y-10">
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('rfid.model_f16.summary.capacity') }}</p>
                <span class="text-[10px] font-black text-blue-500 uppercase">{{ $t('rfid.model_f16.summary.items', { curr: 8, total: 20 }) }}</span>
              </div>
              <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden shadow-inner">
                <div class="h-full bg-blue-500 shadow-lg shadow-blue-500/20" style="width: 40%"></div>
              </div>
            </div>

            <div class="space-y-6">
              <div class="flex gap-4">
                <div class="w-10 h-10 rounded-xl bg-[#1E3A5F] text-white flex items-center justify-center shrink-0 shadow-lg shadow-blue-900/20">
                  <LucideNetwork class="w-5 h-5" />
                </div>
                <div class="space-y-1">
                  <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('rfid.model_f16.summary.hierarchy') }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase leading-relaxed">{{ $t('rfid.model_f16.summary.hierarchy_desc') }}</p>
                </div>
              </div>
              <div class="flex gap-4">
                <div class="w-10 h-10 rounded-xl bg-blue-500 text-white flex items-center justify-center shrink-0 shadow-lg shadow-blue-500/20">
                  <LucideRepeat class="w-5 h-5" />
                </div>
                <div class="space-y-1">
                  <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('rfid.model_f16.summary.sync') }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase leading-relaxed">{{ $t('rfid.model_f16.summary.sync_desc') }}</p>
                </div>
              </div>
            </div>
          </div>

          <div class="pt-10 border-t border-slate-50 dark:border-slate-800 space-y-8">
            <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-lg border border-slate-100 dark:border-slate-800 space-y-4">
              <div class="flex items-center gap-3">
                <LucideInfo class="w-4 h-4 text-blue-500" />
                <h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('rfid.model_f16.summary.network') }}</h4>
              </div>
              <div class="flex items-center justify-between">
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('rfid.model_f16.summary.ldap') }}</p>
                <span class="px-3 py-1 bg-green-50 text-green-500 rounded text-[7px] font-black uppercase tracking-widest">{{ $t('rfid.model_f16.summary.online') }}</span>
              </div>
            </div>

            <button class="w-full py-4 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all">
              {{ $t('rfid.model_f16.summary.btn_log') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucidePrinter, LucideSave, LucideBox, LucideQrCode, 
  LucideFileText, LucideBriefcase, LucideGripVertical, 
  LucideRss, LucideNetwork, LucideRepeat, LucideInfo 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}
</style>

