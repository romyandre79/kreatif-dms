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

    <!-- Main Scanner Card -->
    <div class="glass rounded-[3rem] p-16 flex flex-col items-center justify-center space-y-10 shadow-2xl shadow-slate-200/50 relative overflow-hidden">
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

      <!-- Procedure Box -->
      <div class="w-full max-w-2xl bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100 dark:border-blue-800/30 rounded-[2rem] p-8 space-y-6">
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
    </div>

    <!-- Stats Row -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <!-- Diterima -->
      <div class="glass-card rounded-[2rem] p-8 flex items-center justify-between group hover:-translate-y-1 transition-all duration-300">
        <div class="flex items-center gap-6">
          <div class="w-14 h-14 bg-emerald-50 dark:bg-emerald-900/20 rounded-2xl flex items-center justify-center text-emerald-500 shadow-sm group-hover:bg-emerald-500 group-hover:text-white transition-all duration-500">
            <LucideCheckCircle class="w-7 h-7" />
          </div>
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Diterima Hari Ini</p>
            <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stats.received }}</p>
          </div>
        </div>
        <LucideArrowUpRight class="w-5 h-5 text-slate-200 group-hover:text-emerald-500 transition-colors" />
      </div>

      <!-- Antrean -->
      <div class="glass-card rounded-[2rem] p-8 flex items-center justify-between group hover:-translate-y-1 transition-all duration-300">
        <div class="flex items-center gap-6">
          <div class="w-14 h-14 bg-orange-50 dark:bg-orange-900/20 rounded-2xl flex items-center justify-center text-orange-500 shadow-sm group-hover:bg-orange-500 group-hover:text-white transition-all duration-500">
            <LucideClock class="w-7 h-7" />
          </div>
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Antrean Scan</p>
            <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stats.pending }}</p>
          </div>
        </div>
        <LucideArrowUpRight class="w-5 h-5 text-slate-200 group-hover:text-orange-500 transition-colors" />
      </div>

      <!-- Ditolak -->
      <div class="glass-card rounded-[2rem] p-8 flex items-center justify-between group hover:-translate-y-1 transition-all duration-300">
        <div class="flex items-center gap-6">
          <div class="w-14 h-14 bg-red-50 dark:bg-red-900/20 rounded-2xl flex items-center justify-center text-red-500 shadow-sm group-hover:bg-red-500 group-hover:text-white transition-all duration-500">
            <LucideXCircle class="w-7 h-7" />
          </div>
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Ditolak</p>
            <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stats.rejected }}</p>
          </div>
        </div>
        <LucideArrowUpRight class="w-5 h-5 text-slate-200 group-hover:text-red-500 transition-colors" />
      </div>
    </div>

    <!-- Verification Section (Conditional) -->
    <Transition 
      enter-active-class="transition duration-500 ease-out"
      enter-from-class="transform translate-y-10 opacity-0"
      enter-to-class="transform translate-y-0 opacity-100"
      leave-active-class="transition duration-300 ease-in"
      leave-from-class="transform translate-y-0 opacity-100"
      leave-to-class="transform translate-y-10 opacity-0"
    >
      <div v-if="manifestDetail" class="glass rounded-[3rem] p-12 border-l-8 border-primary-500 shadow-2xl space-y-10">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-6">
            <div class="w-16 h-16 bg-primary-50 dark:bg-primary-900/20 rounded-2xl flex items-center justify-center text-primary-500">
              <LucideFileSearch class="w-8 h-8" />
            </div>
            <div>
              <p class="text-[10px] font-black text-primary-500 uppercase tracking-[0.2em] mb-1">Hasil Scan Ditemukan</p>
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ manifestDetail.title }}</h3>
            </div>
          </div>
          <button @click="manifestDetail = null" class="p-3 hover:bg-slate-50 dark:hover:bg-slate-800 rounded-full transition-colors">
            <LucideX class="w-6 h-6 text-slate-400" />
          </button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2">Jenis Dokumen</p>
            <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ manifestDetail.type_name }}</p>
          </div>
          <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2">Pengirim</p>
            <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ manifestDetail.owner_name }}</p>
          </div>
          <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2">Departemen</p>
            <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ manifestDetail.department_name }}</p>
          </div>
          <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2">Status Saat Ini</p>
            <span class="px-3 py-1 bg-orange-100 text-orange-600 rounded-md text-[10px] font-black uppercase tracking-tighter">{{ manifestDetail.status }}</span>
          </div>
        </div>

        <!-- Action Bar -->
        <div class="flex items-center justify-end gap-6 pt-6 border-t border-slate-100 dark:border-slate-800">
          <div class="flex-1 max-w-md">
             <input 
              v-model="notes"
              type="text" 
              placeholder="Tambahkan catatan penerimaan (opsional)..."
              class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-2xl text-sm font-medium focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all outline-none"
            />
          </div>
          <button 
            @click="handleReject"
            :disabled="processing"
            class="px-10 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-xs font-black uppercase tracking-widest text-red-500 hover:bg-red-50 transition-all disabled:opacity-50"
          >
            Tolak Berkas
          </button>
          <button 
            @click="handleReceive"
            :disabled="processing"
            class="px-12 py-4 bg-emerald-500 hover:bg-emerald-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-emerald-500/20 transition-all disabled:opacity-50 flex items-center gap-3"
          >
            <LucideCheckCircle v-if="!processing" class="w-5 h-5" />
            <LucideRefreshCw v-else class="w-5 h-5 animate-spin" />
            Terima Berkas Fisik
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { 
  LucideQrCode, LucideScanLine, LucideInfo, LucideCheckCircle, 
  LucideClock, LucideXCircle, LucideArrowUpRight, LucideFileSearch, LucideX, LucideRefreshCw
} from 'lucide-vue-next'

const { $api } = useApi()
const auth = useAuthStore()
const scanInput = ref('')
const scannerInput = ref(null)
const processing = ref(false)
const manifestDetail = ref(null)
const notes = ref('')
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
  try {
    const res = await $api(`/intake/manifest/${scanInput.value}`)
    if (res && res.data) {
      manifestDetail.value = res.data
      notes.value = ''
    } else {
      alert('Manifest tidak ditemukan atau sudah diproses.')
    }
  } catch (err) {
    console.error('Scan error:', err)
    alert('Gagal mengambil data manifest: ' + (err.data?.message || err.message))
  } finally {
    processing.value = false
    scanInput.value = ''
  }
}

const handleReceive = async () => {
  if (!manifestDetail.value) return
  
  processing.value = true
  try {
    const res = await $api('/intake/receive', {
      method: 'POST',
      body: {
        document_id: manifestDetail.value.document_id,
        notes: notes.value
      }
    })
    
    if (res && res.success) {
      alert('Berkas fisik berhasil diterima dan dicatat.')
      manifestDetail.value = null
      fetchStats()
    }
  } catch (err) {
    console.error('Receive error:', err)
    alert('Gagal memproses penerimaan: ' + (err.data?.message || err.message))
  } finally {
    processing.value = false
  }
}

const handleReject = async () => {
  if (!manifestDetail.value) return
  if (!notes.value) {
    alert('Mohon isi catatan alasan penolakan.')
    return
  }
  
  alert('Fungsi penolakan fisik sedang dikembangkan.')
}

onMounted(() => {
  // Access control
  const role = auth.user?.role?.toLowerCase()
  const allowedRoles = ['superadmin', 'admin doc controller', 'kepala doc controller']
  
  if (!role || !allowedRoles.includes(role)) {
    console.warn('[Access Control] Unauthorized access to inbound registration for role:', role)
    navigateTo('/dashboard')
    return
  }

  // Fetch initial stats
  fetchStats()

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
