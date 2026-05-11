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

        <!-- Verification Section (Conditional) -->
        <Transition 
          enter-active-class="transition duration-500 ease-out"
          enter-from-class="transform translate-y-10 opacity-0"
          enter-to-class="transform translate-y-0 opacity-100"
          leave-active-class="transition duration-300 ease-in"
          leave-from-class="transform translate-y-0 opacity-100"
          leave-to-class="transform translate-y-10 opacity-0"
        >
          <div v-if="manifestDetail" class="space-y-10">
            <!-- Manifest Info Card -->
            <div class="glass rounded-lg p-12 border-l-8 border-primary-500 shadow-2xl space-y-10">
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-6">
                  <div class="w-16 h-16 bg-primary-50 dark:bg-primary-900/20 rounded-2xl flex items-center justify-center text-primary-500">
                    <LucideFileSearch class="w-8 h-8" />
                  </div>
                  <div>
                    <p class="text-[10px] font-black text-primary-500 uppercase tracking-[0.2em] mb-1">Manifest Serah Terima</p>
                    <h3 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ manifestDetail.manifest_no }}</h3>
                  </div>
                  <span class="px-4 py-1.5 bg-emerald-100 text-emerald-600 rounded-xl text-[10px] font-black uppercase tracking-widest ml-4">{{ manifestDetail.status }}</span>
                </div>
                <div class="flex items-center gap-6">
                  <div class="text-right">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Pemohon</p>
                    <p class="text-sm font-black text-slate-700 dark:text-slate-200">{{ manifestDetail.owner_name }}</p>
                  </div>
                  <div class="text-right">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Departemen</p>
                    <p class="text-sm font-black text-slate-700 dark:text-slate-200">{{ manifestDetail.department_name }}</p>
                  </div>
                  <button 
                    @click="manifestDetail = null"
                    class="ml-6 px-6 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-all flex items-center gap-2"
                  >
                    <LucidePlus class="w-4 h-4" />
                    Scan Manifest Lain
                  </button>
                </div>
              </div>

              <!-- Document List Table -->
              <div class="space-y-6">
                <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white flex items-center gap-3">
                  Daftar Dokumen yang Diserahkan 
                  <span class="text-primary-500 text-sm font-bold">({{ manifestDetail.items?.length || 0 }} Dokumen)</span>
                </h4>

                <div class="overflow-hidden border border-slate-100 dark:border-slate-800 rounded-3xl">
                  <table class="w-full text-left">
                    <thead>
                      <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                        <th class="px-8 py-5 w-16">No</th>
                        <th class="px-8 py-5">Judul Dokumen</th>
                        <th class="px-8 py-5">Jenis</th>
                        <th class="px-8 py-5">Tgl Dokumen</th>
                        <th class="px-8 py-5">Kondisi</th>
                        <th class="px-8 py-5 text-right">Cek</th>
                      </tr>
                    </thead>
                    <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
                      <tr v-for="(item, index) in manifestDetail.items" :key="item.id" class="group transition-colors" :class="item.physical_status === 'damaged' ? 'bg-red-50/30 dark:bg-red-900/10' : 'hover:bg-slate-50/30 dark:hover:bg-slate-800/20'">
                        <td class="px-8 py-6 text-sm font-bold text-slate-400">{{ index + 1 }}</td>
                        <td class="px-8 py-6">
                          <p class="font-black text-sm text-[#1E3A5F] dark:text-slate-200">{{ item.title }}</p>
                          <p class="text-[10px] text-slate-400 font-bold uppercase mt-1">No: {{ item.document_id.substring(0, 8) }}</p>
                          
                          <!-- Damage Notes Input -->
                          <div v-if="item.physical_status === 'damaged'" class="mt-4 p-4 bg-white dark:bg-slate-950 border border-red-100 dark:border-red-900/30 rounded-xl space-y-2" v-motion-slide-top>
                            <p class="text-[9px] font-black text-red-500 uppercase tracking-widest">Catatan Kerusakan</p>
                            <textarea 
                              v-model="item.notes"
                              rows="2"
                              class="w-full bg-transparent border-none focus:ring-0 text-xs font-bold text-slate-600 dark:text-slate-300 p-0 placeholder:text-slate-300"
                              placeholder="Contoh: Halaman sobek, teks tidak terbaca..."
                            ></textarea>
                          </div>
                        </td>
                        <td class="px-8 py-6">
                          <span class="px-3 py-1 bg-blue-50 dark:bg-blue-900/20 text-blue-500 rounded-md text-[9px] font-black uppercase tracking-tight border border-blue-100 dark:border-blue-800/30">{{ item.type_name }}</span>
                        </td>
                        <td class="px-8 py-6 text-sm font-bold text-slate-500">{{ item.date }}</td>
                        <td class="px-8 py-6">
                          <div class="flex items-center gap-6">
                            <label class="flex items-center gap-3 cursor-pointer group/label">
                              <div class="relative flex items-center justify-center">
                                <input type="radio" :name="`cond-${item.id}`" value="good" v-model="item.physical_status" class="sr-only" />
                                <div class="w-5 h-5 rounded-full border-2 transition-all flex items-center justify-center" :class="item.physical_status === 'good' ? 'border-emerald-500 bg-emerald-500' : 'border-slate-200 group-hover/label:border-emerald-300'">
                                  <LucideCheck v-if="item.physical_status === 'good'" class="w-3 h-3 text-white" />
                                </div>
                              </div>
                              <span class="text-xs font-black uppercase tracking-tight transition-colors" :class="item.physical_status === 'good' ? 'text-emerald-600' : 'text-slate-400 group-hover/label:text-slate-600'">Baik</span>
                            </label>
                            
                            <label class="flex items-center gap-3 cursor-pointer group/label">
                              <div class="relative flex items-center justify-center">
                                <input type="radio" :name="`cond-${item.id}`" value="damaged" v-model="item.physical_status" class="sr-only" />
                                <div class="w-5 h-5 rounded-full border-2 transition-all flex items-center justify-center" :class="item.physical_status === 'damaged' ? 'border-red-500 bg-red-500' : 'border-slate-200 group-hover/label:border-red-300'">
                                  <LucideX v-if="item.physical_status === 'damaged'" class="w-3 h-3 text-white" />
                                </div>
                              </div>
                              <span class="text-xs font-black uppercase tracking-tight transition-colors" :class="item.physical_status === 'damaged' ? 'text-red-600' : 'text-slate-400 group-hover/label:text-slate-600'">Rusak</span>
                            </label>
                          </div>
                        </td>
                        <td class="px-8 py-6 text-right">
                          <div class="inline-flex items-center justify-center w-8 h-8 rounded-lg transition-all" :class="(item.physical_status === 'good' || item.physical_status === 'damaged') ? 'bg-emerald-500 text-white shadow-lg shadow-emerald-500/20' : 'bg-slate-100 text-slate-300'">
                            <LucideCheck class="w-5 h-5" />
                          </div>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
            </div>

            <!-- Physical Verification Summary Bar -->
            <div class="bg-blue-50/50 dark:bg-blue-900/10 rounded-lg p-10 flex flex-col md:flex-row items-center justify-between gap-10">
              <div class="flex-1 space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <h5 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Status Verifikasi Fisik</h5>
                    <p class="text-xs font-bold text-slate-500 mt-1">Progres pengecekan dokumen yang diterima</p>
                  </div>
                  <div class="text-right">
                    <p class="text-3xl font-black text-[#1E3A5F] dark:text-white">{{ verificationProgress }}%</p>
                    <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ verifiedCount }} dari {{ manifestDetail.items.length }} Terverifikasi</p>
                  </div>
                </div>
                <div class="h-4 w-full bg-slate-200 dark:bg-slate-800 rounded-full overflow-hidden">
                  <div class="h-full bg-emerald-500 rounded-full transition-all duration-1000" :style="`width: ${verificationProgress}%`"></div>
                </div>
              </div>

              <div class="flex gap-4">
                 <div class="bg-white dark:bg-slate-900 rounded-3xl p-6 px-10 text-center shadow-sm border border-slate-100 dark:border-slate-800">
                   <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-1">Manifest</p>
                   <p class="text-3xl font-black text-[#1E3A5F] dark:text-white">{{ manifestDetail.items.length }}</p>
                 </div>
                 <div class="bg-white dark:bg-slate-900 rounded-3xl p-6 px-10 text-center shadow-sm border border-slate-100 dark:border-slate-800">
                   <p class="text-[9px] font-black text-emerald-500 uppercase tracking-widest mb-1">Baik</p>
                   <p class="text-3xl font-black text-emerald-500">{{ goodCount }}</p>
                 </div>
                 <div class="bg-white dark:bg-slate-900 rounded-3xl p-6 px-10 text-center shadow-sm border border-slate-100 dark:border-slate-800">
                   <p class="text-[9px] font-black text-red-500 uppercase tracking-widest mb-1">Rusak</p>
                   <p class="text-3xl font-black text-red-500">{{ damagedCount }}</p>
                 </div>
              </div>
            </div>

            <!-- Final Submission Action -->
            <div class="flex items-center justify-end gap-6 pt-6">
              <button 
                @click="handleReceive"
                :disabled="processing || verifiedCount < manifestDetail.items.length"
                class="px-12 py-5 bg-emerald-500 hover:bg-emerald-600 text-white rounded-lg text-sm font-black uppercase tracking-widest shadow-2xl shadow-emerald-500/30 transition-all disabled:opacity-50 disabled:grayscale flex items-center gap-4 group"
              >
                <LucideCheckCircle v-if="!processing" class="w-6 h-6 group-hover:scale-110 transition-transform" />
                <LucideRefreshCw v-else class="w-6 h-6 animate-spin" />
                Selesaikan Penerimaan Berkas Fisik
              </button>
            </div>
          </div>
        </Transition>

        <!-- Procedure Box (Only if no manifest) -->
        <div v-if="!manifestDetail" class="w-full bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100 dark:border-blue-800/30 rounded-lg p-8 space-y-6">
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
        <div v-if="!manifestDetail && pendingManifests.length > 0" class="space-y-6" v-motion-fade>
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
                  <div class="space-y-1">
                    <p class="text-[9px] font-black text-primary-500 uppercase tracking-[0.2em]">{{ m.manifest_no }}</p>
                    <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ m.department_name }}</h4>
                  </div>
                  <div class="flex items-center gap-2">
                    <div class="w-5 h-5 bg-slate-100 dark:bg-slate-800 rounded-full flex items-center justify-center text-[8px] font-black text-slate-400 uppercase">
                      {{ m.owner_name?.charAt(0) }}
                    </div>
                    <p class="text-[10px] font-bold text-slate-400 italic">Oleh: {{ m.owner_name }}</p>
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
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { 
  LucideQrCode, LucideScanLine, LucideInfo, LucideCheckCircle, 
  LucideClock, LucideXCircle, LucideArrowUpRight, LucideFileSearch, LucideX, LucideRefreshCw,
  LucidePlus, LucideCheck
} from 'lucide-vue-next'

const { $api } = useApi()
const auth = useAuthStore()
const scanInput = ref('')
const scannerInput = ref(null)
const processing = ref(false)
const manifestDetail = ref(null)
const pendingManifests = ref([])
const notes = ref('')
const stats = reactive({
  received: 0,
  pending: 0,
  rejected: 0
})

// Computed stats for the manifest
const goodCount = computed(() => {
  if (!manifestDetail.value?.items) return 0
  return manifestDetail.value.items.filter(i => i.physical_status === 'good').length
})

const damagedCount = computed(() => {
  if (!manifestDetail.value?.items) return 0
  return manifestDetail.value.items.filter(i => i.physical_status === 'damaged').length
})

const verifiedCount = computed(() => {
  if (!manifestDetail.value?.items) return 0
  return manifestDetail.value.items.filter(i => i.physical_status === 'good' || i.physical_status === 'damaged').length
})

const verificationProgress = computed(() => {
  if (!manifestDetail.value?.items?.length) return 0
  return Math.round((verifiedCount.value / manifestDetail.value.items.length) * 100)
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
    const payload = {
      manifest_id: manifestDetail.value.id,
      notes: notes.value,
      items: manifestDetail.value.items.map(item => ({
        document_id: item.document_id,
        physical_status: item.physical_status,
        notes: item.notes
      }))
    }

    const res = await $api('/intake/receive', {
      method: 'POST',
      body: payload
    })
    
    if (res && res.success) {
      alert('Berkas fisik berhasil diterima dan dicatat.')
      manifestDetail.value = null
      fetchStats()
      fetchPendingManifests()
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
