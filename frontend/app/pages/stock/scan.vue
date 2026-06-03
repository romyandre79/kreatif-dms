<template>
  <div class="space-y-10 pb-20 relative min-h-[600px]" v-motion-fade>
    <!-- State: Mission Selection / Pre-start -->
    <template v-if="state === 'selection'">
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('stock.scan_execution.title') }}
        </h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('stock.scan_execution.subtitle') }}</p>
      </div>

      <div v-if="activeMission" class="glass rounded-lg p-10 flex flex-col lg:flex-row items-center gap-16 shadow-2xl shadow-blue-900/5 overflow-hidden relative">
        <LucideQrCode class="absolute -top-10 -right-10 w-64 h-64 text-slate-50 dark:text-slate-800/20 rotate-12 -z-10" />

        <div class="w-full lg:w-1/3 space-y-8">
          <div class="aspect-square bg-blue-50 dark:bg-blue-900/10 rounded-lg flex items-center justify-center relative overflow-hidden group">
            <div class="absolute inset-0 bg-primary-500/5 animate-pulse"></div>
            <LucideScanBarcode class="w-24 h-24 text-primary-500 group-hover:scale-110 transition-transform duration-700" />
            <div class="absolute inset-x-8 top-1/2 h-0.5 bg-primary-400 shadow-[0_0_15px_rgba(59,130,246,0.8)] animate-scanner"></div>
          </div>
          <div class="space-y-3 px-2">
            <div class="flex items-center justify-between text-[10px] font-black text-slate-400 uppercase tracking-widest">
              <span>{{ $t('stock.scan_execution.active_mission.progress') }}</span>
              <span>{{ progressPercent.toFixed(0) }}%</span>
            </div>
            <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
              <div class="h-full bg-primary-500 transition-all duration-1000" :style="`width: ${progressPercent}%`"></div>
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
                {{ activeMission.title }}
              </h2>
              <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">
                {{ $t('stock.scan_execution.active_mission.id', { id: activeMission.session_no || 'ST-TEMP' }) }}
              </p>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-12">
            <div class="space-y-2">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <LucideMapPin class="w-3.5 h-3.5" />
                {{ $t('stock.scan_execution.active_mission.target') }}
              </div>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ activeMission.target_area?.String || 'All Areas' }}</p>
            </div>
            <div class="space-y-2">
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <LucideBox class="w-3.5 h-3.5" />
                {{ $t('stock.scan_execution.active_mission.expected') }}
              </div>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ expectedCount }} SKUs</p>
            </div>
          </div>

          <div class="flex flex-col sm:flex-row gap-4 pt-4">
            <button @click="startAuditMode" class="flex-grow flex items-center justify-center gap-3 py-5 bg-primary-500 text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-2xl shadow-primary-500/40 hover:bg-primary-600 transition-all active:scale-95 group">
              <LucidePlayCircle class="w-5 h-5 group-hover:scale-110 transition-transform" />
              {{ $t('stock.scan_execution.active_mission.btn_start') }}
            </button>
            <button @click="navigateTo('/stock/missions')" class="px-10 py-5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-lg text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
              Back to List
            </button>
          </div>
        </div>
      </div>

      <!-- No Active Mission State -->
      <div v-else class="glass rounded-lg p-16 text-center space-y-6">
        <LucideAlertTriangle class="w-16 h-16 text-amber-500 mx-auto animate-pulse" />
        <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase">No Active Mission Selected</h3>
        <p class="text-xs font-bold text-slate-400 max-w-md mx-auto">Please go to the Audit Mission List page and select a mission to begin or resume scanning.</p>
        <button @click="navigateTo('/stock/missions')" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-[#152943] transition-all">
          Go to Missions
        </button>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div v-for="(stat, i) in ['last_scan', 'verification', 'team']" :key="i" class="bg-white dark:bg-slate-900 p-8 rounded-lg border border-slate-100 dark:border-slate-800 shadow-sm space-y-4 group hover:shadow-xl transition-all">
          <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400 group-hover:text-primary-500 transition-colors">
            <LucideHistory v-if="stat === 'last_scan'" class="w-5 h-5" />
            <LucideShieldCheck v-else-if="stat === 'verification'" class="w-5 h-5" />
            <LucideUsers v-else class="w-5 h-5" />
          </div>
          <div class="space-y-1">
            <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t(`stock.scan_execution.stats.${stat}`) }}</h4>
            <p v-if="stat === 'last_scan'" class="text-xs font-bold text-slate-400">{{ lastScanTimeStr }}</p>
            <p v-else-if="stat === 'verification'" class="text-xs font-bold text-slate-400">{{ extraCount }} Items to Verify</p>
            <p v-else class="text-xs font-bold text-slate-400">{{ $t('stock.scan_execution.stats.active_today', { count: 2 }) }}</p>
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

      <div class="bg-white dark:bg-slate-900 rounded-lg p-16 shadow-2xl border border-slate-100 dark:border-slate-800 max-w-4xl mx-auto flex flex-col items-center text-center space-y-10 relative overflow-hidden" v-motion-fade>
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] bg-primary-500/5 rounded-full blur-3xl -z-10"></div>
        <div class="relative">
          <div class="w-40 h-40 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-primary-500 shadow-inner">
            <LucideQrCode class="w-16 h-16 animate-pulse" />
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
            <input 
              v-model="scanSkuCode"
              type="text" 
              ref="scanInputRef"
              @keyup.enter="handleScanSubmit" 
              :placeholder="$t('stock.scan_execution.blind_audit.card.waiting')" 
              class="w-full pl-16 pr-10 py-5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-2xl text-xs font-black outline-none shadow-inner focus:ring-4 focus:ring-primary-500/10 transition-all text-center uppercase" 
            />
            <div class="absolute right-6 top-1/2 -translate-y-1/2 w-2 h-2 bg-green-500 rounded-full animate-pulse shadow-[0_0_8px_rgba(34,197,94,0.5)]"></div>
          </div>
          <button @click="handleScanSubmit" :disabled="isScanning" class="w-full py-5 bg-primary-500 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-primary-500/20 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95 disabled:opacity-50">
            <LucideCheckCircle2 class="w-5 h-5" />
            {{ isScanning ? 'Verifying...' : $t('stock.scan_execution.blind_audit.card.btn_verify') }}
          </button>
          <button @click="state = 'active_session'" class="w-full py-4 bg-white dark:bg-slate-800 border border-slate-100 dark:border-slate-700 text-slate-400 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
            <LucideHistory class="w-4 h-4" />
            View Active Session Stats
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
            <p class="text-[9px] font-black text-slate-500 uppercase">
              <template v-if="tag === 'rak'">{{ activeMission.target_area?.String || 'Not Detected' }}</template>
              <template v-else-if="tag === 'zona'">Main Zone</template>
              <template v-else>Active</template>
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
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">CURRENT SESSION LOCATION</p>
            <p class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ activeMission?.target_area?.String || 'ALL AREAS' }}</p>
          </div>
        </div>
        <div class="flex items-center gap-4">
          <button @click="navigateTo(`/stock/reconciliation?session_id=${sessionId}&mission_id=${missionId}`)" class="px-6 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white rounded-xl text-xs font-black uppercase tracking-widest transition-all">
            Proceed to Reconciliation
          </button>
          <span class="px-4 py-1.5 bg-green-50 text-green-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2">
            <span class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></span>
            {{ $t('stock.scan_execution.active_session.header.live') }}
          </span>
        </div>
      </div>

      <div class="bg-white dark:bg-slate-900 rounded-lg p-16 shadow-2xl border border-slate-100 dark:border-slate-800 text-center space-y-8" v-motion-fade>
        <p class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.scan_execution.active_session.counter.title') }}</p>
        <h2 class="text-[120px] font-black text-[#1E3A5F] dark:text-white leading-none tracking-tighter" v-motion-pop>{{ verifiedCount }}</h2>
        <p class="text-[10px] font-bold text-slate-400 italic">{{ $t('stock.scan_execution.active_session.counter.blind_hint') }}</p>
      </div>

      <div v-if="lastScannedItem" class="bg-green-500 p-6 rounded-2xl shadow-xl shadow-green-500/20 flex items-center justify-between text-white" v-motion-slide-visible-bottom>
        <div class="flex items-center gap-6">
          <div class="w-10 h-10 rounded-xl bg-white/20 flex items-center justify-center border border-white/20"><LucideCheckCircle2 class="w-6 h-6" /></div>
          <div class="space-y-0.5">
            <p class="text-[9px] font-black text-white/70 uppercase tracking-widest">{{ $t('stock.scan_execution.active_session.last_scan.title') }}</p>
            <p class="text-sm font-black uppercase tracking-tight">{{ lastScannedItem.SkuCode?.String || lastScannedItem.SkuCode }} - {{ lastScannedItem.ItemName?.String || lastScannedItem.ItemName }}</p>
          </div>
        </div>
        <span class="px-4 py-1.5 bg-white/20 rounded-lg text-[9px] font-black uppercase tracking-widest border border-white/30">{{ lastScannedItem.Status }}</span>
      </div>

      <button @click="state = 'blind_audit'" class="w-full py-6 bg-primary-500 text-white rounded-lg text-sm font-black uppercase tracking-widest shadow-2xl shadow-primary-500/40 hover:bg-primary-600 transition-all flex items-center justify-center gap-4 group active:scale-95">
        <LucideScanBarcode class="w-6 h-6 group-hover:scale-110 transition-transform" />
        {{ $t('stock.scan_execution.active_session.btn_continue') }}
      </button>

      <div class="space-y-6" v-motion-slide-visible-bottom>
        <div class="flex items-center justify-between px-2">
          <div class="flex items-center gap-3">
            <LucideHistory class="w-5 h-5 text-slate-400" />
            <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('stock.scan_execution.active_session.history.title') }}</h3>
          </div>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-lg border border-slate-100 dark:border-slate-800 shadow-sm divide-y divide-slate-50 dark:divide-slate-800/50 overflow-hidden">
          <div v-for="(item, i) in historyItems" :key="i" class="p-8 flex items-center justify-between hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all group">
            <div class="flex items-center gap-6">
              <div :class="`w-10 h-10 rounded-full flex items-center justify-center group-hover:scale-110 transition-transform ${item.status === 'MATCH' ? 'bg-green-50 text-green-500' : 'bg-red-50 text-red-500'}`">
                <LucideCheckCircle2 v-if="item.status === 'MATCH'" class="w-5 h-5" />
                <LucideAlertTriangle v-else class="w-5 h-5" />
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ item.id }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ item.name }}</p>
              </div>
            </div>
            <div class="text-right">
              <p :class="`text-[10px] font-black uppercase tracking-widest ${item.status === 'MATCH' ? 'text-green-600' : 'text-red-600'}`">{{ item.status }}</p>
              <p class="text-[9px] font-bold text-slate-400">{{ item.time }}</p>
            </div>
          </div>
          <div v-if="historyItems.length === 0" class="p-8 text-center text-xs font-bold text-slate-400 italic">
            No items scanned in this session yet.
          </div>
        </div>
      </div>

      <!-- Session Progress Bar (Bottom) -->
      <div class="fixed bottom-0 left-0 right-0 lg:left-72 bg-white dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 px-10 py-6 z-20 flex items-center justify-between shadow-[0_-10px_30px_rgba(0,0,0,0.05)]" v-motion-slide-bottom>
        <div class="space-y-1.5 w-1/3">
          <div class="flex items-center justify-between text-[8px] font-black text-slate-400 uppercase tracking-widest">
            <span>PROGRESS</span>
            <span>{{ verifiedCount }} / {{ expectedCount }} Item</span>
          </div>
          <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
            <div class="h-full bg-primary-500" :style="`width: ${progressPercent}%`"></div>
          </div>
        </div>
        <div class="flex items-center gap-12 text-center">
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mb-1">FOUND MATCH</p>
            <p class="text-sm font-black text-green-500">{{ matchedCount }}</p>
          </div>
          <div class="w-px h-8 bg-slate-100 dark:bg-slate-800"></div>
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mb-1">ON LOAN</p>
            <p class="text-sm font-black text-blue-500">{{ loanCount }}</p>
          </div>
          <div class="w-px h-8 bg-slate-100 dark:bg-slate-800"></div>
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mb-1">MISPLACED / EXTRA</p>
            <p class="text-sm font-black text-red-500">{{ extraCount }}</p>
          </div>
        </div>
      </div>
    </template>

    <!-- Misplaced / Extra Alert Overlay -->
    <div v-if="showMisplacedAlert" class="fixed inset-0 bg-[#1E3A5F]/40 backdrop-blur-sm z-50 flex items-center justify-center p-10" @click.self="showMisplacedAlert = false">
      <div class="bg-white dark:bg-slate-900 rounded-lg w-full max-w-4xl overflow-hidden shadow-2xl border-4 border-red-500" v-motion-pop>
        <div class="bg-red-500 px-10 py-6 flex items-center justify-center gap-4 text-white">
          <LucideAlertTriangle class="w-6 h-6 animate-bounce" />
          <h3 class="text-xl font-black uppercase tracking-widest">DISCREPANCY ALERT: UNEXPECTED FIND</h3>
        </div>
        
        <div class="p-16 flex flex-col md:flex-row gap-16">
          <!-- Scanned Info -->
          <div class="w-full md:w-1/3 space-y-6">
            <div class="aspect-square bg-slate-100 dark:bg-slate-800 rounded-3xl overflow-hidden relative border-4 border-slate-50 dark:border-slate-700 shadow-inner flex items-center justify-center">
              <LucideFileX2 class="w-16 h-16 text-red-500" />
              <div class="absolute top-4 right-4 px-3 py-1 bg-red-500 text-white rounded-lg text-[8px] font-black uppercase tracking-widest shadow-lg">
                DISCREPANCY
              </div>
            </div>
            <div class="space-y-1 text-center">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">SCANNED BARCODE ID</p>
              <p class="text-lg font-black text-red-500 tracking-tight">{{ alertSku }}</p>
            </div>
          </div>

          <!-- Info & Actions -->
          <div class="flex-grow space-y-10">
            <div class="space-y-6">
              <span class="inline-flex items-center gap-2 px-4 py-1.5 bg-red-50 dark:bg-red-900/20 text-red-600 rounded-full text-[9px] font-black uppercase tracking-widest border border-red-100 dark:border-red-800">
                <span class="w-2 h-2 bg-red-600 rounded-full animate-ping"></span>
                Extra Item Found
              </span>
              <div class="space-y-3">
                <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase leading-tight tracking-tight">
                  {{ alertMessage }}
                </h3>
                <p class="text-sm font-medium text-slate-500 leading-relaxed">
                  This item is physically present in the target area, but system records indicate otherwise. Please return the item to its expected location or reconcile this in the system.
                </p>
              </div>
            </div>

            <div class="space-y-4">
              <button @click="showMisplacedAlert = false" class="w-full py-5 bg-primary-500 text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-xl shadow-primary-500/30 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95">
                <LucideCheckCircle2 class="w-5 h-5" />
                {{ $t('stock.scan_execution.misplaced_alert.btn_acknowledge') }}
              </button>
              <button @click="navigateTo(`/stock/reconciliation?session_id=${sessionId}&mission_id=${missionId}`)" class="w-full py-5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-lg text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
                <LucideTarget class="w-5 h-5" />
                Open Reconciliation Report
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { 
  LucideQrCode, LucideScanBarcode, LucideMapPin, LucideBox, 
  LucidePlayCircle, LucideHistory, LucideShieldCheck, LucideUsers,
  LucideAlertTriangle, LucideArrowLeftRight, LucideCheckCircle2,
  LucideKeyboard, LucideInfo, LucideLayoutGrid, LucideFileX2,
  LucideTarget, LucideX
} from 'lucide-vue-next'

const route = useRoute()
const sessionId = ref(route.query.session_id || '')
const missionId = ref(route.query.mission_id || '')

const { $api } = useApi()
const toast = useToast()

const state = ref('selection') // 'selection', 'blind_audit', 'active_session'
const showMisplacedAlert = ref(false)
const scanSkuCode = ref('')
const isScanning = ref(false)
const items = ref([])
const activeMission = ref(null)
const lastScannedItem = ref(null)

const alertSku = ref('')
const alertMessage = ref('')

const scanInputRef = ref(null)

const expectedCount = computed(() => items.value.length)
const verifiedCount = computed(() => items.value.filter(i => i.PhysicalQty > 0).length)
const matchedCount = computed(() => items.value.filter(i => i.Status === 'MATCH').length)
const loanCount = computed(() => items.value.filter(i => i.Status === 'ON LOAN').length)
const extraCount = computed(() => items.value.filter(i => i.Status === 'EXTRA').length)
const progressPercent = computed(() => expectedCount.value > 0 ? (verifiedCount.value / expectedCount.value) * 100 : 0)

const lastScanTimeStr = computed(() => {
  if (!lastScannedItem.value) return 'No scan activity yet'
  return 'Just now'
})

const fetchSessionItems = async () => {
  if (!sessionId.value) return
  try {
    const res = await $api(`/stock/sessions/${sessionId.value}/items`)
    if (res.success && res.data) {
      items.value = res.data || []
    }
  } catch (err) {
    console.error('Error fetching session items:', err)
  }
}

const loadMissionDetails = async () => {
  try {
    const res = await $api('/stock/missions?limit=100')
    if (res.success && res.data && res.data.items) {
      const found = res.data.items.find(m => m.id === missionId.value || m.session_id?.Bytes === sessionId.value)
      if (found) {
        activeMission.value = found
        if (!sessionId.value && found.session_id?.Bytes) {
          sessionId.value = found.session_id.Bytes
        }
      }
    }
  } catch (err) {
    console.error('Error loading mission details:', err)
  }
}

const startAuditMode = () => {
  state.value = 'blind_audit'
  nextTick(() => {
    if (scanInputRef.value) {
      scanInputRef.value.focus()
    }
  })
}

const handleScanSubmit = async () => {
  if (!scanSkuCode.value) {
    toast.warning('Please enter or scan a barcode/SKU code.')
    return
  }
  isScanning.value = true
  try {
    const res = await $api('/stock/scan', {
      method: 'POST',
      body: {
        session_id: sessionId.value,
        sku_code: scanSkuCode.value.trim(),
        current_rack: activeMission.value?.target_area?.String || ''
      }
    })
    if (res.success && res.data) {
      const scanResult = res.data
      lastScannedItem.value = scanResult.item
      
      await fetchSessionItems()

      if (!scanResult.is_match) {
        // Discrepancy or Extra item!
        alertSku.value = scanSkuCode.value.trim()
        alertMessage.value = scanResult.message || 'Misplaced or extra item detected.'
        showMisplacedAlert.value = true
        toast.warning('Discrepancy alert!')
      } else {
        toast.success('Scan verified successfully!')
        state.value = 'active_session'
      }
      scanSkuCode.value = ''
    } else {
      toast.error(res.message || 'Verification failed.')
    }
  } catch (err) {
    console.error(err)
    toast.error(err.data?.message || 'Error processing scan.')
  } finally {
    isScanning.value = false
    nextTick(() => {
      if (scanInputRef.value) {
        scanInputRef.value.focus()
      }
    })
  }
}

const historyItems = computed(() => {
  return items.value
    .filter(i => i.PhysicalQty > 0)
    .map(i => ({
      id: i.SkuCode?.String || i.SkuCode,
      name: i.ItemName?.String || i.ItemName,
      status: i.Status,
      time: 'Just now'
    }))
    .reverse()
})

onMounted(async () => {
  if (sessionId.value || missionId.value) {
    await loadMissionDetails()
    await fetchSessionItems()
  }
})
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
