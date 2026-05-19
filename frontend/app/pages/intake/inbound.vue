<template>
  <div class="max-w-6xl mx-auto space-y-10 pb-20" v-motion-fade>
    <!-- Header -->
    <div class="space-y-2">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
        {{ $t('intake.inbound.title') || 'Penerimaan Dokumen Fisik' }}
      </h1>
      <p class="text-slate-500 font-medium">
        {{ $t('intake.inbound.subtitle') || 'Scan QR Code pada Manifest Serah Terima untuk memulai proses penerimaan' }}
      </p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-4 gap-10 items-start">
      <!-- Left: Scanner & Results -->
      <div class="lg:col-span-3 space-y-10">
        <!-- Main Scanner Card -->
        <div class="glass rounded-lg p-16 flex flex-col items-center justify-center space-y-10 shadow-2xl shadow-slate-200/50 relative overflow-hidden">
          <!-- Decorative Background Elements -->
          <div class="absolute top-0 right-0 w-64 h-64 bg-primary-500/5 rounded-full -mr-32 -mt-32 blur-3xl"></div>
          <div class="absolute bottom-0 left-0 w-64 h-64 bg-emerald-500/5 rounded-full -ml-32 -mb-32 blur-3xl"></div>

          <!-- Icon Section -->
          <div class="w-32 h-32 bg-slate-50 dark:bg-slate-900 rounded-3xl flex items-center justify-center text-slate-300 shadow-inner relative group">
            <LucideQrCode class="w-16 h-16 transition-transform group-hover:scale-110 duration-500" />
            <div class="absolute -top-2 -right-2 w-6 h-6 bg-primary-500 rounded-full animate-ping opacity-20"></div>
            <div class="absolute -top-2 -right-2 w-6 h-6 bg-primary-500 rounded-full flex items-center justify-center shadow-lg">
              <div class="w-2 h-2 bg-white rounded-full"></div>
            </div>
          </div>

          <!-- Text Section -->
          <div class="text-center space-y-4 max-w-md">
            <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
              Scan Manifest Serah Terima
            </h2>
            <p class="text-sm font-medium text-slate-400 leading-relaxed italic">
              Arahkan scanner ke QR Code pada dokumen Manifest fisik yang Anda terima dari kurir atau unit kerja.
            </p>
          </div>

          <!-- Scanner Input Box -->
          <div class="w-full max-w-xl relative group">
            <div class="absolute inset-0 bg-primary-500/10 rounded-2xl blur-xl group-hover:bg-primary-500/20 transition-all duration-500"></div>
            <div class="relative bg-white dark:bg-slate-900 border-2 border-dashed border-slate-200 dark:border-slate-800 rounded-2xl p-6 flex items-center gap-6 transition-all group-hover:border-primary-300">
              <div class="w-12 h-12 bg-slate-50 dark:bg-slate-800 rounded-xl flex items-center justify-center text-slate-400 group-hover:text-primary-500 transition-colors">
                <LucideScanLine class="w-6 h-6" />
              </div>
              <input 
                v-model="scanInput"
                @keyup.enter="handleScan"
                type="text"
                ref="scannerInput"
                class="flex-1 bg-transparent border-none focus:ring-0 text-sm font-bold text-slate-600 dark:text-slate-300 placeholder:text-slate-300 placeholder:font-normal uppercase tracking-widest"
                placeholder="Menunggu input scanner..."
                autofocus
              />
            </div>
          </div>
        </div>


        <!-- Procedure Box -->
        <div class="w-full bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100 dark:border-blue-800/30 rounded-lg p-8 space-y-6">
          <div class="flex items-center gap-3">
            <div class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center text-white">
              <LucideInfo class="w-3.5 h-3.5" />
            </div>
            <h3 class="text-xs font-black text-blue-800 dark:text-blue-300 uppercase tracking-widest">Prosedur Penerimaan:</h3>
          </div>
          <ol class="space-y-4">
            <li v-for="(step, i) in procedures" :key="i" class="flex gap-4">
              <span class="text-xs font-black text-blue-400/60">{{ i + 1 }}.</span>
              <p class="text-xs font-bold text-slate-600 dark:text-slate-400 leading-relaxed">
                {{ step }}
              </p>
            </li>
          </ol>
        </div>

        <!-- Pending Manifests List -->
        <div v-if="pendingManifests.length > 0" class="space-y-6" v-motion-fade>
          <div class="flex items-center justify-between">
            <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Antrean Manifest</h3>
            <span class="px-4 py-1.5 bg-primary-50 text-primary-600 rounded-xl text-[10px] font-black uppercase tracking-widest">{{ pendingManifests.length }} Manifest</span>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div 
              v-for="m in pendingManifests" 
              :key="m.id" 
              @click="selectManifest(m.manifest_no)"
              class="glass p-8 rounded-lg border-l-4 border-primary-500 hover:shadow-xl transition-all cursor-pointer group relative overflow-hidden"
            >
              <div class="absolute top-0 right-0 w-24 h-24 bg-primary-500/5 rounded-full -mr-12 -mt-12 blur-2xl"></div>
              <div class="flex items-center justify-between relative">
                <div class="space-y-3">
                  <div class="flex items-center gap-3">
                    <div class="space-y-1">
                      <p class="text-[9px] font-black text-primary-500 uppercase tracking-[0.2em]">{{ m.manifest_no }}</p>
                      <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ m.department_name }}</h4>
                    </div>
                    <span 
                      :class="[
                        'px-2 py-0.5 rounded-md text-[8px] font-black uppercase tracking-widest border',
                        m.status === 'pending' ? 'bg-orange-50 text-orange-500 border-orange-100' : 'bg-emerald-50 text-emerald-500 border-emerald-100'
                      ]"
                    >
                      {{ m.status }}
                    </span>
                  </div>
                  <div class="flex items-center justify-between">
                    <div class="flex items-center gap-2">
                      <div class="w-5 h-5 bg-slate-100 dark:bg-slate-800 rounded-full flex items-center justify-center text-[8px] font-black text-slate-400 uppercase">
                        {{ m.owner_name?.charAt(0) }}
                      </div>
                      <p class="text-[10px] font-bold text-slate-400 italic">Oleh: {{ m.owner_name }}</p>
                    </div>
                    <div class="flex items-center gap-1.5 text-slate-300">
                      <LucideClock class="w-3 h-3" />
                      <span class="text-[9px] font-bold">{{ m.created_at }}</span>
                    </div>
                  </div>
                </div>
                <div class="w-12 h-12 rounded-2xl bg-slate-50 dark:bg-slate-900 flex items-center justify-center text-slate-300 group-hover:bg-primary-500 group-hover:text-white transition-all duration-500 shadow-inner group-hover:shadow-primary-500/20 group-hover:-translate-y-1">
                  <LucideArrowUpRight class="w-6 h-6" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right: Sidebar Stats -->
      <div class="space-y-8">
        <!-- Diterima -->
        <div class="glass rounded-lg p-8 flex flex-col gap-6 group hover:-translate-y-1 transition-all duration-300">
          <div class="w-14 h-14 bg-emerald-50 dark:bg-emerald-900/20 rounded-2xl flex items-center justify-center text-emerald-500 shadow-sm group-hover:bg-emerald-500 group-hover:text-white transition-all duration-500">
            <LucideCheckCircle class="w-7 h-7" />
          </div>
          <div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Diterima Hari Ini</p>
            <div class="flex items-end justify-between">
              <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stats.received }}</p>
              <LucideArrowUpRight class="w-5 h-5 text-emerald-500 opacity-0 group-hover:opacity-100 transition-opacity" />
            </div>
          </div>
        </div>

        <!-- Antrean -->
        <div class="glass rounded-lg p-8 flex flex-col gap-6 group hover:-translate-y-1 transition-all duration-300">
          <div class="w-14 h-14 bg-orange-50 dark:bg-orange-900/20 rounded-2xl flex items-center justify-center text-orange-500 shadow-sm group-hover:bg-orange-500 group-hover:text-white transition-all duration-500">
            <LucideClock class="w-7 h-7" />
          </div>
          <div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Antrean Verifikasi</p>
            <div class="flex items-end justify-between">
              <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stats.pending }}</p>
              <LucideArrowUpRight class="w-5 h-5 text-orange-500 opacity-0 group-hover:opacity-100 transition-opacity" />
            </div>
          </div>
        </div>

        <!-- Ditolak -->
        <div class="glass rounded-lg p-8 flex flex-col gap-6 group hover:-translate-y-1 transition-all duration-300">
          <div class="w-14 h-14 bg-red-50 dark:bg-red-900/20 rounded-2xl flex items-center justify-center text-red-500 shadow-sm group-hover:bg-red-500 group-hover:text-white transition-all duration-500">
            <LucideXCircle class="w-7 h-7" />
          </div>
          <div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Ditolak</p>
            <div class="flex items-end justify-between">
              <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stats.rejected }}</p>
              <LucideArrowUpRight class="w-5 h-5 text-red-500 opacity-0 group-hover:opacity-100 transition-opacity" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Scan Error Modal (High Fidelity) -->
    <Transition name="fade">
      <div v-if="showErrorModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-slate-900/60 backdrop-blur-sm">
        <div class="bg-white dark:bg-slate-900 rounded-[32px] w-full max-w-lg shadow-2xl relative overflow-hidden flex flex-col items-center p-12 space-y-10" v-motion-pop>
          <!-- Red Accent Bar -->
          <div class="absolute top-0 left-0 bottom-0 w-2 bg-red-500"></div>

          <!-- Icon Section -->
          <div class="w-24 h-24 bg-red-50 dark:bg-red-900/20 rounded-full flex items-center justify-center text-red-500">
            <LucideAlertTriangle class="w-12 h-12" />
          </div>

          <!-- Text Section -->
          <div class="text-center space-y-4">
            <h3 class="text-2xl font-black text-red-600 uppercase tracking-tight">Manifest Tidak Ditemukan</h3>
            <p class="text-sm font-bold text-slate-500 leading-relaxed max-w-sm mx-auto">
              Kode QR tidak cocok dengan data di sistem. Mohon periksa kembali label pada fisik dokumen.
            </p>
          </div>

          <!-- Scanned Value Box -->
          <div class="w-full bg-red-50/50 dark:bg-red-950/20 rounded-2xl p-6 text-center space-y-2 border border-red-100 dark:border-red-900/30">
            <p class="text-[9px] font-black text-red-400 uppercase tracking-[0.2em]">Scanned Value</p>
            <p class="text-lg font-black text-red-600 tracking-widest uppercase">{{ lastScannedValue }}</p>
          </div>

          <!-- Possible Causes -->
          <div class="w-full space-y-4 text-left">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Kemungkinan Penyebab:</p>
            <ul class="space-y-3">
              <li class="flex items-start gap-4">
                <span class="text-xs font-black text-emerald-500">1.</span>
                <p class="text-xs font-bold text-slate-600 dark:text-slate-400 leading-tight">Manifest belum di-approve oleh Atasan</p>
              </li>
              <li class="flex items-start gap-4">
                <span class="text-xs font-black text-emerald-500">2.</span>
                <p class="text-xs font-bold text-slate-600 dark:text-slate-400 leading-tight">QR Code rusak, kotor, atau tidak terbaca sempurna</p>
              </li>
              <li class="flex items-start gap-4">
                <span class="text-xs font-black text-emerald-500">3.</span>
                <p class="text-xs font-bold text-slate-600 dark:text-slate-400 leading-tight">Manifest sudah pernah diproses/diterima sebelumnya</p>
              </li>
            </ul>
          </div>

          <!-- Footer Info -->
          <div class="w-full pt-6 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">ID Perangkat: SCAN-JKT-04</p>
            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Sesi Aktif: 02h 15m</p>
          </div>

          <!-- Actions -->
          <div class="grid grid-cols-2 gap-4 w-full">
            <button 
              @click="showErrorModal = false"
              class="py-4 bg-[#1E3A5F] hover:bg-[#2a4d7d] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all flex items-center justify-center gap-2"
            >
              <LucideRefreshCw class="w-4 h-4" />
              Scan Ulang
            </button>
            <button 
              @click="showErrorModal = false"
              class="py-4 border-2 border-[#1E3A5F]/20 dark:border-slate-800 rounded-2xl text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-2"
            >
              <LucideSearch class="w-4 h-4" />
              Cari Manual
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { 
  LucideQrCode, LucideScanLine, LucideInfo, LucideCheckCircle, 
  LucideClock, LucideXCircle, LucideArrowUpRight, LucideRefreshCw,
  LucideAlertTriangle, LucideSearch
} from 'lucide-vue-next'

const { $api } = useApi()
const auth = useAuthStore()
const scanInput = ref('')
const scannerInput = ref(null)
const processing = ref(false)
const showErrorModal = ref(false)
const lastScannedValue = ref('')
const pendingManifests = ref([])
const stats = reactive({
  received: 0,
  pending: 0,
  rejected: 0
})

const procedures = [
  'Verifikasi fisik Manifest dengan data yang muncul setelah scan.',
  'Hitung jumlah box/amplop sesuai dengan daftar lampiran.',
  'Lakukan konfirmasi "Terima" atau "Tolak" pada sistem.'
]

const fetchPendingManifests = async () => {
  try {
    const res = await $api('/intake/pending')
    if (res && res.data) {
      pendingManifests.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch pending manifests:', err)
  }
}

const selectManifest = async (manifestNo) => {
  scanInput.value = manifestNo
  handleScan()
}

const fetchStats = async () => {
  try {
    const res = await $api('/intake/stats')
    if (res && res.data) {
      stats.received = res.data.received || 0
      stats.pending = res.data.pending || 0
      stats.rejected = res.data.rejected || 0
    }
  } catch (err) {
    console.error('Failed to fetch intake stats:', err)
  }
}

const handleScan = async () => {
  if (!scanInput.value) return
  
  processing.value = true
  lastScannedValue.value = scanInput.value
  try {
    // We just try to fetch to check existence, then redirect
    const res = await $api(`/intake/manifest/${scanInput.value}`)
    if (res && res.data) {
      navigateTo(`/intake/manifest/${scanInput.value}`)
    } else {
      showErrorModal.value = true
    }
  } catch (err) {
    console.error('Scan error:', err)
    showErrorModal.value = true
  } finally {
    processing.value = false
    scanInput.value = ''
  }
}


onMounted(() => {
  // Access control
  const role = auth.user?.role?.toLowerCase()
  const allowedRoles = ['superadmin', 'admin doc controller', 'kepala doc controller', 'kepala dc', 'admin dc']
  
  if (!role || !allowedRoles.includes(role)) {
    console.warn('[Access Control] Unauthorized access to inbound registration for role:', role)
    navigateTo('/dashboard')
    return
  }

  // Fetch initial stats
  fetchStats()
  fetchPendingManifests()

  // Focus scanner input on load
  if (scannerInput.value) {
    scannerInput.value.focus()
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

.glass-card {
  background: white;
  border: 1px solid rgba(226, 232, 240, 0.6);
  box-shadow: 0 20px 40px -15px rgba(0, 0, 0, 0.03);
}

.dark .glass {
  background: rgba(13, 18, 31, 0.7);
  border-color: rgba(30, 41, 59, 0.5);
}

.dark .glass-card {
  background: #0D121F;
  border-color: rgba(30, 41, 59, 0.5);
}
</style>
