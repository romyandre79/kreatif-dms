<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="space-y-1">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.printing.title') }}</h1>
      <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest font-mono">{{ $t('warehouse.printing.batch', { id: 'RFID-2023-F20-8842' }) }}</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Hardware Status -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.printing.hardware.title') }}</h3>
          
          <div class="flex flex-col items-center text-center space-y-6">
            <div class="w-24 h-24 rounded-3xl bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 flex items-center justify-center text-slate-200 shadow-inner group">
              <LucidePrinter class="w-12 h-12 group-hover:text-blue-500 transition-colors" />
            </div>
            <div class="space-y-1">
              <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.printing.hardware.device') }}</h4>
              <p class="text-[9px] font-bold text-green-500 uppercase tracking-widest">{{ $t('warehouse.printing.hardware.status') }}</p>
            </div>
          </div>

          <div class="space-y-8">
            <div class="space-y-3">
              <div class="flex justify-between text-[10px] font-black uppercase tracking-widest">
                <span class="text-slate-400">{{ $t('warehouse.printing.hardware.stock') }}</span>
                <span class="text-blue-500">84%</span>
              </div>
              <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-blue-500" style="width: 84%"></div>
              </div>
              <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('warehouse.printing.hardware.labels_left', { count: 420 }) }}</p>
            </div>

            <div class="space-y-3">
              <div class="flex justify-between text-[10px] font-black uppercase tracking-widest">
                <span class="text-slate-400">{{ $t('warehouse.printing.hardware.ribbon') }}</span>
                <span class="text-amber-500">32%</span>
              </div>
              <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-amber-500" style="width: 32%"></div>
              </div>
            </div>
          </div>

          <button class="w-full py-4 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all">
            {{ $t('warehouse.printing.hardware.btn_diag') }}
          </button>
        </div>

        <!-- Encoding Config -->
        <div class="glass p-10 rounded-[3rem] space-y-8 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.printing.config.title') }}</h3>
          <div class="grid grid-cols-2 gap-4">
            <div class="p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-2">
              <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">PROTO COL</p>
              <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase leading-tight">{{ $t('warehouse.printing.config.protocol') }}</p>
            </div>
            <div class="p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-2">
              <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">ANTENNA</p>
              <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase leading-tight">{{ $t('warehouse.printing.config.antenna') }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Center Column: Batch Execution Panel -->
      <div class="lg:col-span-6 space-y-10">
        <div class="glass rounded-[3rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800 border-t-8 border-blue-500">
          <div class="p-10 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
            <div class="space-y-1">
              <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.printing.execution.title') }}</h3>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('warehouse.printing.execution.subtitle', { count: '1,200', hub: 'Jakarta Hub' }) }}</p>
            </div>
            <button class="px-8 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-4 group">
              <LucidePlay class="w-4 h-4 fill-current group-hover:scale-110 transition-transform" />
              {{ $t('warehouse.printing.execution.btn_resume') }}
            </button>
          </div>

          <div class="p-10 space-y-12">
            <div class="space-y-8">
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-4">
                    <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center">
                      <LucidePrinter class="w-5 h-5" />
                    </div>
                    <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.printing.execution.printing') }}</p>
                  </div>
                  <div class="flex items-center gap-4">
                    <p class="text-[10px] font-black text-slate-400 uppercase">842 / 1,200</p>
                    <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">70.1%</p>
                  </div>
                </div>
                <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden shadow-inner">
                  <div class="h-full bg-[#1E3A5F]" style="width: 70.1%"></div>
                </div>
              </div>

              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-4 text-slate-400">
                    <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center">
                      <LucideRss class="w-5 h-5" />
                    </div>
                    <p class="text-[10px] font-black uppercase tracking-widest">{{ $t('warehouse.printing.execution.encoding') }}</p>
                  </div>
                  <div class="flex items-center gap-4">
                    <p class="text-[10px] font-black text-slate-400 uppercase">818 / 1,200</p>
                    <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tighter opacity-60">68.2%</p>
                  </div>
                </div>
                <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden shadow-inner">
                  <div class="h-full bg-[#1E3A5F] opacity-40" style="width: 68.2%"></div>
                </div>
              </div>
            </div>

            <div class="flex items-center justify-between pt-6 border-t border-slate-50 dark:border-slate-800">
              <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('warehouse.printing.execution.est', { time: '4m 12s' }) }}</p>
              <p class="text-[8px] font-bold text-slate-300 uppercase tracking-widest">{{ $t('warehouse.printing.execution.delay', { ms: 140 }) }}</p>
            </div>

            <!-- Current Operation -->
            <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 space-y-6">
              <h4 class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.printing.execution.detail.title') }}</h4>
              <div class="flex items-start justify-between">
                <div class="flex items-center gap-6">
                  <div class="w-12 h-12 bg-white dark:bg-slate-800 rounded-xl border border-slate-100 dark:border-slate-700 flex items-center justify-center">
                    <LucideQrCode class="w-6 h-6 text-slate-300" />
                  </div>
                  <div class="space-y-1">
                    <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Asset ID: JKT-HUB-8842-X</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase font-mono tracking-tighter">EPC: 303425789012345678901234</p>
                  </div>
                </div>
                <div class="text-right space-y-1">
                  <p class="text-[10px] font-black text-green-500 uppercase tracking-widest">{{ $t('warehouse.printing.execution.detail.success') }}</p>
                  <div class="flex justify-end">
                    <div class="w-2 h-2 bg-green-500 rounded-full animate-ping"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Execution Stats -->
        <div class="grid grid-cols-3 gap-6">
          <div v-for="s in [
            { id: 'speed', val: '2.4', unit: 'lbl/sec' },
            { id: 'temp', val: '42.5', unit: '°C' },
            { id: 'retries', val: '4', color: 'text-red-500' }
          ]" :key="s.id" class="glass p-8 rounded-[2.5rem] space-y-2 text-center border border-slate-50 dark:border-slate-800">
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`warehouse.printing.execution.stats.${s.id}`) }}</p>
            <p :class="`text-2xl font-black ${s.color || 'text-[#1E3A5F] dark:text-white'} tracking-tighter`">{{ s.val }}<span v-if="s.unit" class="text-[10px] ml-1">{{ s.unit }}</span></p>
          </div>
        </div>
      </div>

      <!-- Right Column: Real-Time Log -->
      <div class="lg:col-span-3 flex flex-col gap-6 h-[850px]">
        <div class="flex items-center justify-between px-4">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.printing.logs.title') }}</h3>
          <LucideChevronDown class="w-4 h-4 text-slate-300" />
        </div>
        
        <div class="flex-grow glass rounded-[3rem] p-6 overflow-y-auto custom-scrollbar space-y-4 border border-slate-50 dark:border-slate-800">
          <div v-for="(log, idx) in [
            { time: '14:22:15', status: 'OK', msg: 'Batch ID 8842 processed. EPC verified.' },
            { time: '14:21:58', status: 'FAIL', msg: 'Asset JKT-8839: Encoding Timeout.', retry: true },
            { time: '14:21:30', status: 'OK', msg: 'Batch ID 8841 processed successfully.' },
            { time: '14:21:05', status: 'OK', msg: 'Asset JKT-8840 EPC update complete.' },
            { time: '14:20:55', status: 'SYSTEM', msg: 'Low ribbon warning threshold reached (35%).' },
            { time: '14:20:44', status: 'OK', msg: 'Label print command executed for JKT-8839.' },
            { time: '14:20:12', status: 'OK', msg: 'Hardware handshake complete. Printer Ready.' }
          ]" :key="idx" :class="`p-6 rounded-2xl border ${log.status === 'FAIL' ? 'bg-red-50 border-red-100' : 'bg-slate-50/50 dark:bg-slate-800/50 border-slate-100 dark:border-slate-700'} space-y-3` ">
            <div class="flex items-center justify-between">
              <span class="text-[9px] font-bold text-slate-400 uppercase font-mono">{{ log.time }}</span>
              <span :class="`text-[8px] font-black uppercase tracking-widest ${log.status === 'OK' ? 'text-green-500' : log.status === 'FAIL' ? 'text-red-500' : 'text-amber-500'}`">{{ log.status }}</span>
            </div>
            <p :class="`text-[10px] font-bold leading-relaxed uppercase tracking-tight ${log.status === 'FAIL' ? 'text-red-600' : 'text-[#1E3A5F] dark:text-white'}`">{{ log.msg }}</p>
            <div v-if="log.retry" class="flex gap-4 pt-2 border-t border-red-200">
              <button class="text-[8px] font-black text-red-600 uppercase tracking-widest hover:underline">{{ $t('warehouse.printing.logs.manual_retry') }}</button>
              <button class="text-[8px] font-black text-slate-400 uppercase tracking-widest hover:underline">{{ $t('warehouse.printing.logs.skip') }}</button>
            </div>
          </div>
        </div>

        <button class="w-full py-5 bg-slate-50 dark:bg-slate-800 border-2 border-slate-100 dark:border-slate-700 text-[#1E3A5F] dark:text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-100 transition-all flex items-center justify-center gap-4">
          <LucideDownload class="w-4 h-4" />
          {{ $t('warehouse.printing.logs.btn_export') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucidePrinter, LucidePlay, LucideRss, LucideQrCode, 
  LucideChevronDown, LucideDownload 
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

