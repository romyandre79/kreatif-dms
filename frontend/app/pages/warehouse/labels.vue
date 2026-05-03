<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('registration.migration.print_qr.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('registration.migration.print_qr.subtitle') }}</p>
      </div>
      
      <div class="flex items-center gap-4">
        <div class="relative w-full md:w-80">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input type="text" :placeholder="$t('registration.migration.print_qr.search_placeholder')" class="w-full pl-11 pr-5 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all shadow-sm" />
        </div>
        <button class="w-11 h-11 bg-white dark:bg-slate-900 text-red-500 border border-slate-200 dark:border-slate-800 rounded-xl flex items-center justify-center relative shadow-sm hover:bg-red-50 transition-all">
          <LucideBell class="w-5 h-5" />
          <span class="absolute top-3 right-3 w-2 h-2 bg-red-600 rounded-full border-2 border-white dark:border-slate-900"></span>
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
      <div class="lg:col-span-2 space-y-8">
        <!-- Label Configuration -->
        <div class="glass p-10 rounded-[2.5rem] space-y-10">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-xl bg-primary-50 dark:bg-primary-900/20 text-primary-500 flex items-center justify-center shadow-sm">
              <LucideSettings2 class="w-5 h-5" />
            </div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.title') }}</h3>
          </div>

          <div class="space-y-10">
            <div class="space-y-4">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.type.label') }}</label>
              <div class="flex gap-4">
                <button @click="labelType = 'doc'" :class="`flex-grow flex items-center justify-center gap-3 py-4 px-6 rounded-2xl text-xs font-black transition-all ${labelType === 'doc' ? 'bg-white dark:bg-slate-800 border-2 border-[#1E3A5F] text-[#1E3A5F] dark:text-white shadow-lg' : 'bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 text-slate-400'}`">
                  <LucideFileText class="w-4 h-4" />
                  {{ $t('registration.migration.print_qr.config.type.doc') }}
                </button>
                <button @click="labelType = 'rack'" :class="`flex-grow flex items-center justify-center gap-3 py-4 px-6 rounded-2xl text-xs font-black transition-all ${labelType === 'rack' ? 'bg-white dark:bg-slate-800 border-2 border-[#1E3A5F] text-[#1E3A5F] dark:text-white shadow-lg' : 'bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 text-slate-400'}`">
                  <LucideLayoutGrid class="w-4 h-4" />
                  {{ $t('registration.migration.print_qr.config.type.rack') }}
                </button>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <div class="space-y-4">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.mode.label') }}</label>
                <div class="relative">
                  <select class="w-full pl-5 pr-12 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 transition-all">
                    <option>{{ $t('registration.migration.print_qr.config.mode.single') }}</option>
                  </select>
                  <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
                </div>
              </div>
              <div class="space-y-4">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.copies') }}</label>
                <div class="flex items-center bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl overflow-hidden">
                  <button @click="copies = Math.max(1, copies - 1)" class="p-4 hover:bg-slate-100 dark:hover:bg-slate-700 transition-colors text-slate-400"><LucideMinus class="w-4 h-4" /></button>
                  <input type="text" v-model="copies" class="w-full text-center bg-transparent text-sm font-black outline-none" />
                  <button @click="copies++" class="p-4 hover:bg-slate-100 dark:hover:bg-slate-700 transition-colors text-slate-400"><LucidePlus class="w-4 h-4" /></button>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.printer.label') }}</label>
              <div class="relative">
                <LucidePrinter class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
                <select class="w-full pl-12 pr-12 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold appearance-none outline-none">
                  <option>Zebra ZT411 - Warehouse Floor 1</option>
                </select>
                <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              </div>
              <p class="text-[9px] font-black text-green-500 uppercase tracking-widest flex items-center gap-1.5 px-1">
                <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
                {{ $t('registration.migration.print_qr.config.printer.status') }}
              </p>
            </div>
          </div>
        </div>

        <!-- Data Source -->
        <div class="glass p-10 rounded-[2.5rem] space-y-6">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-xl bg-primary-50 dark:bg-primary-900/20 text-primary-500 flex items-center justify-center shadow-sm">
              <LucideDatabase class="w-5 h-5" />
            </div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.print_qr.source.title') }}</h3>
          </div>
          <div class="space-y-3">
            <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.source.label') }}</label>
            <div class="flex gap-4">
              <input type="text" :placeholder="$t('registration.migration.print_qr.source.placeholder')" class="flex-grow px-5 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none shadow-inner" />
              <button class="px-8 py-4 bg-slate-800 text-white rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-900 transition-all shadow-lg shadow-slate-900/20 active:scale-95">
                {{ $t('registration.migration.print_qr.source.btn_fetch') }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Label Preview Column -->
      <div class="space-y-8">
        <div class="glass p-10 rounded-[2.5rem] space-y-10" v-motion-slide-right>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <LucideEye class="w-5 h-5 text-primary-500" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.title') }}</h3>
            </div>
            <button class="text-[9px] font-black text-primary-500 uppercase tracking-widest hover:underline">{{ $t('registration.migration.print_qr.preview.settings') }}</button>
          </div>

          <!-- Label Render -->
          <div class="aspect-[4/3] bg-white border border-slate-100 shadow-2xl rounded-2xl p-8 flex flex-col relative overflow-hidden group">
            <div class="flex-grow flex items-start justify-between">
              <div class="space-y-1">
                <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.type_label') }}</p>
                <p class="text-xl font-black text-slate-800 tracking-tight">DOC-2023-0842</p>
              </div>
              <div class="w-16 h-16 bg-slate-800 rounded-lg p-2 flex items-center justify-center shadow-lg">
                <LucideQrCode class="w-full h-full text-white" />
              </div>
            </div>
            <div class="mt-auto flex items-end justify-between border-t border-slate-50 pt-6">
              <div class="space-y-1">
                <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.loc_label') }}</p>
                <p class="text-2xl font-black text-slate-800 tracking-tighter">RACK-A1-04</p>
              </div>
              <div class="text-right">
                <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.area_label') }}</p>
                <p class="text-[10px] font-black text-slate-500 uppercase">Zone B14</p>
              </div>
            </div>
            <!-- Label Size Marker -->
            <div class="absolute bottom-2 left-1/2 -translate-x-1/2 text-[7px] font-black text-slate-300 uppercase tracking-widest opacity-0 group-hover:opacity-100 transition-opacity">
              {{ $t('registration.migration.print_qr.preview.thermal_hint') }}
            </div>
          </div>

          <div class="p-6 bg-primary-50/30 dark:bg-primary-900/10 border border-primary-100/50 dark:border-primary-800/30 rounded-2xl flex gap-4">
            <LucideInfo class="w-4 h-4 text-primary-500 flex-shrink-0" />
            <p class="text-[10px] font-bold text-slate-500 leading-relaxed italic">
              {{ $t('registration.migration.print_qr.preview.info') }}
            </p>
          </div>

          <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-[1.5rem] text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 active:scale-95">
            <LucidePrinter class="w-5 h-5" />
            {{ $t('registration.migration.print_qr.preview.btn_print') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Recent Print Jobs -->
    <div class="space-y-6" v-motion-slide-visible-bottom>
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('registration.migration.print_qr.history.title') }}</h3>
        <button class="text-[10px] font-black text-primary-500 uppercase tracking-widest hover:underline">{{ $t('registration.migration.print_qr.history.view_all') }}</button>
      </div>
      <div class="glass rounded-[2.5rem] overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="text-[9px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                <th class="p-8 px-10">{{ $t('registration.migration.print_qr.history.cols.time') }}</th>
                <th class="p-8">{{ $t('registration.migration.print_qr.history.cols.id') }}</th>
                <th class="p-8">{{ $t('registration.migration.print_qr.history.cols.printer') }}</th>
                <th class="p-8 text-center">{{ $t('registration.migration.print_qr.history.cols.status') }}</th>
                <th class="p-8 text-right px-10">{{ $t('registration.migration.print_qr.history.cols.action') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
              <tr v-for="job in printHistory" :key="job.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all">
                <td class="p-8 px-10 text-xs font-bold text-slate-500 uppercase">{{ job.time }}</td>
                <td class="p-8 text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ job.labelId }}</td>
                <td class="p-8 text-xs font-bold text-slate-500">{{ job.printer }}</td>
                <td class="p-8 text-center">
                  <span class="px-3 py-1 bg-green-50 text-green-500 rounded-md text-[8px] font-black uppercase tracking-widest border border-green-100">SUCCESS</span>
                </td>
                <td class="p-8 text-right px-10">
                  <button class="text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:underline">{{ $t('registration.migration.print_qr.history.actions.reprint') }}</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="p-6 text-center border-t border-slate-100 dark:border-slate-800">
          <button class="text-[9px] font-black text-slate-400 uppercase tracking-widest hover:text-slate-600 transition-colors">
            {{ $t('registration.migration.print_qr.history.view_logs') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideSearch, LucideBell, LucideSettings2, LucideFileText, 
  LucideLayoutGrid, LucideChevronDown, LucideMinus, LucidePlus, 
  LucidePrinter, LucideDatabase, LucideEye, LucideQrCode, LucideInfo
} from 'lucide-vue-next'

const labelType = ref('doc')
const copies = ref(1)

const printHistory = [
  { id: 1, time: 'Oct 24, 14:20', labelId: 'DOC-2023-0841', printer: 'Zebra ZT411', status: 'SUCCESS' },
  { id: 2, time: 'Oct 24, 13:55', labelId: 'RACK-B2-01', printer: 'Zebra ZT411', status: 'SUCCESS' }
]
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>

