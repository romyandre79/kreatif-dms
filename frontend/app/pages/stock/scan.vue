<template>
  <div class="space-y-10 pb-20 relative min-h-[600px]" v-motion-fade>
    <!-- State: Mission Selection -->
    <template v-if="state === 'selection'">
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('stock.scan_execution.title') }}
        </h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('stock.scan_execution.subtitle') }}</p>
      </div>

      <div class="glass rounded-[3rem] p-10 flex flex-col lg:flex-row items-center gap-16 shadow-2xl shadow-blue-900/5 overflow-hidden relative">
        <LucideQrCode class="absolute -top-10 -right-10 w-64 h-64 text-slate-50 dark:text-slate-800/20 rotate-12 -z-10" />

        <div class="w-full lg:w-1/3 space-y-8">
          <div class="aspect-square bg-blue-50 dark:bg-blue-900/10 rounded-[2.5rem] flex items-center justify-center relative overflow-hidden group">
            <div class="absolute inset-0 bg-primary-500/5 animate-pulse"></div>
            <LucideScanBarcode class="w-24 h-24 text-primary-500 group-hover:scale-110 transition-transform duration-700" />
            <div class="absolute inset-x-8 top-1/2 h-0.5 bg-primary-400 shadow-[0_0_15px_rgba(59,130,246,0.8)] animate-scanner"></div>
          </div>
          <div class="space-y-3 px-2">
            <div class="flex items-center justify-between text-[10px] font-black text-slate-400 uppercase tracking-widest">
              <span>{{ $t('stock.scan_execution.active_mission.progress') }}</span>
              <span>0%</span>
            </div>
            <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
              <div class="h-full bg-primary-500 w-0 transition-all duration-1000"></div>
            </div>
          </div>
        </div>

        <div class="flex-grow space-y-10 w-full lg:w-auto">
          <div class="space-y-4">
            <span class="px-4 py-1.5 bg-primary-50 dark:bg-primary-900/20 text-primary-600 dark:text-primary-400 rounded-full text-[10px] font-black uppercase tracking-widest border border-primary-100 dark:border-primary-800">
              {{ $t('stock.scan_execution.active_mission.badge') }}
            </span>
            <div class="space-y-1">
              <h2 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
                {{ $t('stock.scan_execution.active_mission.title') }}
              </h2>
              <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">
                {{ $t('stock.scan_execution.active_mission.id', { id: 'ST-2023-001' }) }}
              </p>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-12">
            <div class="space-y-2">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <LucideMapPin class="w-3.5 h-3.5" />
                {{ $t('stock.scan_execution.active_mission.target') }}
              </div>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Dept. Legal</p>
            </div>
            <div class="space-y-2">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <LucideBox class="w-3.5 h-3.5" />
                {{ $t('stock.scan_execution.active_mission.expected') }}
              </div>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">142 SKUs</p>
            </div>
          </div>

          <div class="flex flex-col sm:flex-row gap-4 pt-4">
            <button @click="state = 'blind_audit'" class="flex-grow flex items-center justify-center gap-3 py-5 bg-primary-500 text-white rounded-[1.5rem] text-xs font-black uppercase tracking-widest shadow-2xl shadow-primary-500/40 hover:bg-primary-600 transition-all active:scale-95 group">
              <LucidePlayCircle class="w-5 h-5 group-hover:scale-110 transition-transform" />
              {{ $t('stock.scan_execution.active_mission.btn_start') }}
            </button>
            <button class="px-10 py-5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-[1.5rem] text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
              {{ $t('stock.scan_execution.active_mission.btn_details') }}
            </button>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div v-for="(stat, i) in ['last_scan', 'verification', 'team']" :key="i" class="bg-white dark:bg-slate-900 p-8 rounded-[2rem] border border-slate-100 dark:border-slate-800 shadow-sm space-y-4 group hover:shadow-xl transition-all">
          <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400 group-hover:text-primary-500 transition-colors">
            <LucideHistory v-if="stat === 'last_scan'" class="w-5 h-5" />
            <LucideShieldCheck v-else-if="stat === 'verification'" class="w-5 h-5" />
            <LucideUsers v-else class="w-5 h-5" />
          </div>
          <div class="space-y-1">
            <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t(`stock.scan_execution.stats.${stat}`) }}</h4>
            <p v-if="stat === 'last_scan'" class="text-xs font-bold text-slate-400">{{ $t('stock.scan_execution.stats.hours_ago', { count: 2 }) }}</p>
            <p v-else-if="stat === 'verification'" class="text-xs font-bold text-slate-400">{{ $t('stock.scan_execution.stats.pending') }}</p>
            <p v-else class="text-xs font-bold text-slate-400">{{ $t('stock.scan_execution.stats.active_today', { count: 4 }) }}</p>
          </div>
        </div>
      </div>
    </template>

    <!-- State: Blind Audit Active -->
    <template v-else-if="state === 'blind_audit'">
      <div class="bg-amber-50 dark:bg-amber-900/10 border border-amber-100 dark:border-amber-800 rounded-2xl p-6 flex items-center gap-6 shadow-sm shadow-amber-500/5" v-motion-slide-top>
        <div class="w-10 h-10 rounded-xl bg-white dark:bg-amber-950 flex items-center justify-center text-amber-500 shadow-sm">
          <LucideAlertTriangle class="w-5 h-5" />
        </div>
        <div class="space-y-0.5">
          <h4 class="text-xs font-black text-amber-800 dark:text-amber-400 uppercase tracking-widest">{{ $t('stock.scan_execution.blind_audit.alert.title') }}</h4>
          <p class="text-[11px] font-bold text-amber-700/70 dark:text-amber-500/70">{{ $t('stock.scan_execution.blind_audit.alert.desc') }}</p>
        </div>
      </div>

      <div class="bg-white dark:bg-slate-900 rounded-[3rem] p-16 shadow-2xl border border-slate-100 dark:border-slate-800 max-w-4xl mx-auto flex flex-col items-center text-center space-y-10 relative overflow-hidden" v-motion-fade>
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] bg-primary-500/5 rounded-full blur-3xl -z-10"></div>
        <div class="relative">
          <div class="w-40 h-40 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-primary-500 shadow-inner">
            <LucideQrCode class="w-16 h-16" />
          </div>
          <div class="absolute -bottom-3 left-1/2 -translate-x-1/2 px-4 py-1 bg-primary-500 text-white rounded-full text-[8px] font-black uppercase tracking-widest shadow-lg shadow-primary-500/30 flex items-center gap-1.5">
            <LucideArrowLeftRight class="w-3 h-3" /> READY
          </div>
        </div>
        <div class="space-y-3">
          <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('stock.scan_execution.blind_audit.card.title') }}</h2>
          <p class="text-sm font-medium text-slate-500">{{ $t('stock.scan_execution.blind_audit.card.subtitle') }}</p>
        </div>
        <div class="w-full max-w-md space-y-4">
          <div class="relative group">
            <LucideScanBarcode class="absolute left-6 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300 group-focus-within:text-primary-500 transition-colors" />
            <input type="text" readonly :placeholder="$t('stock.scan_execution.blind_audit.card.waiting')" class="w-full pl-16 pr-10 py-5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-2xl text-xs font-black outline-none shadow-inner" />
            <div class="absolute right-6 top-1/2 -translate-y-1/2 w-2 h-2 bg-green-500 rounded-full animate-pulse shadow-[0_0_8px_rgba(34,197,94,0.5)]"></div>
          </div>
          <button @click="state = 'active_session'" class="w-full py-5 bg-primary-500 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-primary-500/20 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95">
            <LucideCheckCircle2 class="w-5 h-5" />
            {{ $t('stock.scan_execution.blind_audit.card.btn_verify') }}
          </button>
          <button class="w-full py-4 bg-white dark:bg-slate-800 border border-slate-100 dark:border-slate-700 text-slate-400 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
            <LucideKeyboard class="w-4 h-4" />
            {{ $t('stock.scan_execution.blind_audit.card.btn_manual') }}
          </button>
        </div>
        <div class="w-full pt-12 border-t border-slate-50 dark:border-slate-800 grid grid-cols-3 gap-8 text-center">
          <div v-for="tag in ['rak', 'zona', 'level']" :key="tag" class="space-y-1.5">
            <div class="flex items-center justify-center gap-2 text-slate-400">
              <LucideInfo v-if="tag === 'rak'" class="w-4 h-4" />
              <LucideMapPin v-else-if="tag === 'zona'" class="w-4 h-4" />
              <LucideLayoutGrid v-else class="w-4 h-4" />
              <span class="text-[9px] font-black uppercase tracking-widest">{{ $t(`stock.scan_execution.blind_audit.footer.${tag}`) }}</span>
            </div>
            <p class="text-[9px] font-black text-slate-400 uppercase">
              <template v-if="tag === 'rak'">{{ $t('stock.scan_execution.blind_audit.footer.not_detected') }}</template>
              <template v-else-if="tag === 'zona'">{{ $t('stock.scan_execution.blind_audit.footer.not_determined') }}</template>
              <template v-else>{{ $t('stock.scan_execution.blind_audit.footer.required') }}</template>
            </p>
          </div>
        </div>
      </div>
    </template>

    <!-- State: Active Session -->
    <template v-else-if="state === 'active_session'">
      <div class="bg-white dark:bg-slate-900 p-6 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm flex items-center justify-between" v-motion-slide-top>
        <div class="flex items-center gap-6">
          <div class="w-12 h-12 rounded-2xl bg-primary-50 dark:bg-primary-900/20 text-primary-500 flex items-center justify-center shadow-sm"><LucideMapPin class="w-6 h-6" /></div>
          <div>
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.scan_execution.active_session.header.label') }}</p>
            <p class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('stock.scan_execution.active_session.header.val') }}</p>
          </div>
        </div>
        <span class="px-4 py-1.5 bg-green-50 text-green-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2">
          <span class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></span>
          {{ $t('stock.scan_execution.active_session.header.live') }}
        </span>
      </div>

      <div class="bg-white dark:bg-slate-900 rounded-[3rem] p-16 shadow-2xl border border-slate-100 dark:border-slate-800 text-center space-y-8" v-motion-fade>
        <p class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.scan_execution.active_session.counter.title') }}</p>
        <h2 class="text-[120px] font-black text-[#1E3A5F] dark:text-white leading-none tracking-tighter" v-motion-pop>14</h2>
        <p class="text-[10px] font-bold text-slate-400 italic">{{ $t('stock.scan_execution.active_session.counter.blind_hint') }}</p>
      </div>

      <div class="bg-green-500 p-6 rounded-2xl shadow-xl shadow-green-500/20 flex items-center justify-between text-white" v-motion-slide-visible-bottom>
        <div class="flex items-center gap-6">
          <div class="w-10 h-10 rounded-xl bg-white/20 flex items-center justify-center border border-white/20"><LucideCheckCircle2 class="w-6 h-6" /></div>
          <div class="space-y-0.5">
            <p class="text-[9px] font-black text-white/70 uppercase tracking-widest">{{ $t('stock.scan_execution.active_session.last_scan.title') }}</p>
            <p class="text-sm font-black uppercase tracking-tight">{{ $t('stock.scan_execution.active_session.last_scan.result', { id: 'LGL.AMS.2024.001' }) }}</p>
          </div>
        </div>
        <span class="px-4 py-1.5 bg-white/20 rounded-lg text-[9px] font-black uppercase tracking-widest border border-white/30">{{ $t('stock.scan_execution.active_session.last_scan.verified') }}</span>
      </div>

      <button @click="showMisplacedAlert = true" class="w-full py-6 bg-primary-500 text-white rounded-[2rem] text-sm font-black uppercase tracking-widest shadow-2xl shadow-primary-500/40 hover:bg-primary-600 transition-all flex items-center justify-center gap-4 group active:scale-95">
        <LucideScanBarcode class="w-6 h-6 group-hover:scale-110 transition-transform" />
        {{ $t('stock.scan_execution.active_session.btn_continue') }}
      </button>

      <div class="space-y-6" v-motion-slide-visible-bottom>
        <div class="flex items-center justify-between px-2">
          <div class="flex items-center gap-3">
            <LucideHistory class="w-5 h-5 text-slate-400" />
            <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('stock.scan_execution.active_session.history.title') }}</h3>
          </div>
          <button class="text-[10px] font-black text-primary-500 uppercase tracking-widest hover:underline">{{ $t('stock.scan_execution.active_session.history.view_all') }}</button>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 shadow-sm divide-y divide-slate-50 dark:divide-slate-800/50 overflow-hidden">
          <div v-for="(item, i) in historyItems" :key="i" class="p-8 flex items-center justify-between hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all group">
            <div class="flex items-center gap-6">
              <div class="w-10 h-10 rounded-full bg-green-50 dark:bg-green-900/10 text-green-500 flex items-center justify-center group-hover:scale-110 transition-transform"><LucideCheckCircle2 class="w-5 h-5" /></div>
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ item.id }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ item.name }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-[10px] font-black text-green-600 uppercase tracking-widest">{{ $t('stock.scan_execution.active_session.history.match') }}</p>
              <p class="text-[9px] font-bold text-slate-400">{{ $t('stock.scan_execution.active_session.history.time_ago', { count: item.time }) }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Session Progress Bar (Bottom) -->
      <div class="fixed bottom-0 left-0 right-0 lg:left-72 bg-white dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 px-10 py-6 z-20 flex items-center justify-between shadow-[0_-10px_30px_rgba(0,0,0,0.05)]" v-motion-slide-bottom>
        <div class="space-y-1.5 w-1/3">
          <div class="flex items-center justify-between text-[8px] font-black text-slate-400 uppercase tracking-widest">
            <span>{{ $t('stock.scan_execution.misplaced_alert.footer.progress') }}</span>
            <span>142 / 500 Item</span>
          </div>
          <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
            <div class="h-full bg-primary-500" style="width: 28%"></div>
          </div>
        </div>
        <div class="flex items-center gap-12 text-center">
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('stock.scan_execution.misplaced_alert.footer.found') }}</p>
            <p class="text-sm font-black text-green-500">128</p>
          </div>
          <div class="w-px h-8 bg-slate-100 dark:bg-slate-800"></div>
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('stock.scan_execution.misplaced_alert.footer.misplaced') }}</p>
            <p class="text-sm font-black text-red-500">14</p>
          </div>
        </div>
      </div>
    </template>

    <!-- Misplaced Alert Overlay -->
    <div v-if="showMisplacedAlert" class="fixed inset-0 bg-[#1E3A5F]/40 backdrop-blur-sm z-50 flex items-center justify-center p-10" @click.self="showMisplacedAlert = false">
      <div class="bg-white dark:bg-slate-900 rounded-[3rem] w-full max-w-4xl overflow-hidden shadow-2xl border-4 border-red-500" v-motion-pop>
        <div class="bg-red-500 px-10 py-6 flex items-center justify-center gap-4 text-white">
          <LucideAlertTriangle class="w-6 h-6 animate-bounce" />
          <h3 class="text-xl font-black uppercase tracking-widest">{{ $t('stock.scan_execution.misplaced_alert.header') }}</h3>
        </div>
        
        <div class="p-16 flex flex-col md:flex-row gap-16">
          <!-- Left: Scanned Image -->
          <div class="w-full md:w-1/3 space-y-6">
            <div class="aspect-square bg-slate-100 dark:bg-slate-800 rounded-3xl overflow-hidden relative border-4 border-slate-50 dark:border-slate-700 shadow-inner group">
              <img src="https://images.unsplash.com/photo-1568667256549-094345857637?q=80&w=400&auto=format&fit=crop" class="w-full h-full object-cover opacity-50 grayscale group-hover:grayscale-0 group-hover:opacity-100 transition-all" />
              <div class="absolute inset-0 flex items-center justify-center">
                <LucideFileX2 class="w-16 h-16 text-red-500" />
              </div>
              <div class="absolute top-4 right-4 px-3 py-1 bg-red-500 text-white rounded-lg text-[8px] font-black uppercase tracking-widest shadow-lg">
                {{ $t('stock.scan_execution.misplaced_alert.badge') }}
              </div>
            </div>
            <div class="space-y-1 text-center">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.scan_execution.misplaced_alert.scanned_id') }}</p>
              <p class="text-lg font-black text-red-500 tracking-tight">LGL.AMS.B3.055</p>
            </div>
          </div>

          <!-- Right: Info & Actions -->
          <div class="flex-grow space-y-10">
            <div class="space-y-6">
              <span class="inline-flex items-center gap-2 px-4 py-1.5 bg-red-50 dark:bg-red-900/20 text-red-600 rounded-full text-[9px] font-black uppercase tracking-widest border border-red-100 dark:border-red-800">
                <span class="w-2 h-2 bg-red-600 rounded-full animate-ping"></span>
                {{ $t('stock.scan_execution.misplaced_alert.discrepancy') }}
              </span>
              <div class="space-y-3">
                <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase leading-tight tracking-tight">
                  {{ $t('stock.scan_execution.misplaced_alert.target_loc', { loc: 'Rak B3' }).split('Rak B3')[0] }}
                  <span class="text-primary-500 border-b-4 border-primary-500/20">Rak B3</span>.
                </h3>
                <p class="text-sm font-medium text-slate-500 leading-relaxed">
                  {{ $t('stock.scan_execution.misplaced_alert.instruction', { current: 'Rak A1' }) }}
                </p>
              </div>
            </div>

            <div class="space-y-4">
              <button @click="showMisplacedAlert = false" class="w-full py-5 bg-primary-500 text-white rounded-[1.5rem] text-xs font-black uppercase tracking-widest shadow-xl shadow-primary-500/30 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95">
                <LucideCheckCircle2 class="w-5 h-5" />
                {{ $t('stock.scan_execution.misplaced_alert.btn_acknowledge') }}
              </button>
              <button @click="navigateTo('/stock/reconciliation')" class="w-full py-5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-[1.5rem] text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
                <LucideTarget class="w-5 h-5" />
                {{ $t('stock.scan_execution.misplaced_alert.btn_reconcile') }}
              </button>
            </div>

            <div class="p-6 bg-amber-50/50 dark:bg-amber-900/10 border border-amber-100/50 dark:border-amber-800/30 rounded-2xl flex gap-4">
              <LucideInfo class="w-4 h-4 text-amber-500 shrink-0" />
              <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed italic">
                {{ $t('stock.scan_execution.misplaced_alert.log_info', { name: 'John Doe' }) }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideQrCode, LucideScanBarcode, LucideMapPin, LucideBox, 
  LucidePlayCircle, LucideHistory, LucideShieldCheck, LucideUsers,
  LucideAlertTriangle, LucideArrowLeftRight, LucideCheckCircle2,
  LucideKeyboard, LucideInfo, LucideLayoutGrid, LucideFileX2,
  LucideTarget
} from 'lucide-vue-next'

const state = ref('selection') // 'selection', 'blind_audit', 'active_session'
const showMisplacedAlert = ref(false)

const historyItems = [
  { id: 'LGL.AMS.2024.001', name: 'Legal Documents Cabinet A', time: 2 },
  { id: 'OPS.JAK.2023.442', name: 'Operational Records', time: 5 },
  { id: 'FIN.TAX.2024.089', name: 'Finance & Tax Archive', time: 12 }
]
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}

@keyframes scanner {
  0%, 100% { top: 20%; }
  50% { top: 80%; }
}

.animate-scanner {
  animation: scanner 2.5s ease-in-out infinite;
}
</style>

