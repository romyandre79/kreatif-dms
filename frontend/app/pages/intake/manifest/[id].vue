<template>
  <div class="max-w-6xl mx-auto space-y-10 pb-20 relative" v-motion-fade>
    <!-- Toast Notification -->
    <Transition name="slide-fade">
      <div v-if="showToast" class="fixed top-10 left-1/2 -translate-x-1/2 z-[200] w-full max-w-lg px-6">
        <div class="bg-emerald-500 text-white rounded-3xl p-6 shadow-2xl flex items-center gap-6 relative overflow-hidden group border border-white/20">
          <div class="w-14 h-14 bg-white/20 rounded-full flex items-center justify-center flex-shrink-0">
            <div class="w-8 h-8 rounded-full border-2 border-white flex items-center justify-center">
              <LucideCheck class="w-4 h-4 stroke-[4px]" />
            </div>
          </div>
          <div class="flex-1">
            <h4 class="font-black text-sm uppercase tracking-tight leading-tight">
              {{ manifestDetail?.manifest_no }} — {{ manifestDetail?.items?.length }} dokumen berhasil diterima
            </h4>
            <p class="text-[10px] font-bold text-white/80 uppercase mt-1 tracking-widest">Status: Diterima (Antrean Scan)</p>
          </div>
          <button @click="showToast = false" class="p-2 hover:bg-white/20 rounded-xl transition-all">
            <LucideX class="w-5 h-5" />
          </button>
        </div>
      </div>
    </Transition>

    <!-- Header -->
    <div class="flex items-center justify-between">
      <div class="space-y-2">
        <button 
          @click="navigateTo('/intake/inbound')"
          class="flex items-center gap-2 text-slate-400 hover:text-primary-500 font-bold text-xs uppercase tracking-widest transition-colors mb-4"
        >
          <LucideArrowLeft class="w-4 h-4" />
          Kembali ke Inbound
        </button>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
          Verifikasi Manifest Fisik
        </h1>
      </div>

      <div v-if="manifestDetail" class="flex items-center gap-4">
        <div class="text-right">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Status Saat Ini</p>
          <span class="px-4 py-1 bg-orange-50 text-orange-600 rounded-xl text-[10px] font-black uppercase tracking-widest border border-orange-100">
            {{ manifestDetail.status }}
          </span>
        </div>
      </div>
    </div>

    <div v-if="loading" class="glass rounded-lg p-20 flex flex-col items-center justify-center space-y-4">
      <LucideRefreshCw class="w-12 h-12 text-primary-500 animate-spin" />
      <p class="text-slate-500 font-bold uppercase tracking-widest text-xs">Mengambil data manifest...</p>
    </div>

    <div v-else-if="manifestDetail" class="space-y-10">
      <!-- Manifest Info Card -->
      <div class="glass rounded-lg p-12 border-l-8 border-primary-500 shadow-2xl space-y-10">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-6">
            <div class="w-16 h-16 bg-primary-50 dark:bg-primary-900/20 rounded-2xl flex items-center justify-center text-primary-500">
              <LucideFileSearch class="w-8 h-8" />
            </div>
            <div>
              <p class="text-[10px] font-black text-primary-500 uppercase tracking-[0.2em] mb-1">Nomor Manifest</p>
              <h3 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ manifestDetail.manifest_no }}</h3>
            </div>
          </div>
          <div class="flex items-center gap-10">
            <div class="text-right">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Pengirim / Pemohon</p>
              <p class="text-sm font-black text-slate-700 dark:text-slate-200">{{ manifestDetail.owner_name }}</p>
            </div>
            <div class="text-right">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Unit Kerja / Departemen</p>
              <p class="text-sm font-black text-slate-700 dark:text-slate-200">{{ manifestDetail.department_name }}</p>
            </div>
            <div class="text-right">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Tanggal Pengajuan</p>
              <p class="text-sm font-black text-slate-700 dark:text-slate-200">{{ manifestDetail.created_at }}</p>
            </div>
          </div>
        </div>

        <!-- Document List Table -->
        <div class="space-y-6">
          <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white flex items-center gap-3">
            Daftar Lampiran Dokumen
            <span class="text-primary-500 text-sm font-bold">({{ manifestDetail.items?.length || 0 }} Item)</span>
          </h4>

          <div class="overflow-hidden border border-slate-100 dark:border-slate-800 rounded-3xl">
            <table class="w-full text-left">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                  <th class="px-8 py-5 w-16">No</th>
                  <th class="px-8 py-5">Informasi Dokumen</th>
                  <th class="px-8 py-5">Kategori</th>
                  <th class="px-8 py-5">Tgl Dokumen</th>
                  <th class="px-8 py-5">Kondisi Fisik</th>
                  <th class="px-8 py-5 text-right">Verifikasi</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
                <tr v-for="(item, index) in manifestDetail.items" :key="item.id" class="group transition-colors" :class="item.physical_status === 'damaged' ? 'bg-red-50/30 dark:bg-red-900/10' : 'hover:bg-slate-50/30 dark:hover:bg-slate-800/20'">
                  <td class="px-8 py-6 text-sm font-bold text-slate-400">{{ index + 1 }}</td>
                  <td class="px-8 py-6">
                    <p class="font-black text-sm text-[#1E3A5F] dark:text-slate-200">{{ item.title }}</p>
                    <p class="text-[10px] text-slate-400 font-bold uppercase mt-1">ID: {{ item.document_id.substring(0, 8) }}</p>
                    
                    <!-- Damage Notes Input -->
                    <div v-if="item.physical_status === 'damaged'" class="mt-4 p-4 bg-white dark:bg-slate-950 border border-red-100 dark:border-red-900/30 rounded-xl space-y-2" v-motion-slide-top>
                      <p class="text-[9px] font-black text-red-500 uppercase tracking-widest">Catatan Kerusakan</p>
                      <textarea 
                        v-model="item.notes"
                        rows="2"
                        class="w-full bg-transparent border-none focus:ring-0 text-xs font-bold text-slate-600 dark:text-slate-300 p-0 placeholder:text-slate-300"
                        placeholder="Detail kerusakan fisik..."
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

      <!-- General Notes -->
      <div class="glass rounded-lg p-8 space-y-4">
        <h5 class="text-xs font-black text-slate-400 uppercase tracking-widest">Catatan Penerimaan (Opsional)</h5>
        <textarea 
          v-model="mainNotes"
          rows="3"
          class="w-full bg-slate-50/50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 rounded-xl p-4 text-sm font-medium text-slate-600 dark:text-slate-300 focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 outline-none transition-all"
          placeholder="Tambahkan catatan jika diperlukan..."
        ></textarea>
      </div>

      <!-- Final Submission Action -->
      <div class="flex items-center justify-end gap-6 pt-6">
        <button 
          @click="showRejectModal = true"
          class="px-8 py-5 border-2 border-slate-200 dark:border-slate-800 text-slate-400 rounded-lg text-sm font-black uppercase tracking-widest hover:bg-slate-50 transition-all"
        >
          Tolak Manifest
        </button>
        <button 
          @click="showConfirmModal = true"
          :disabled="processing || verifiedCount < manifestDetail.items.length"
          class="px-12 py-5 bg-emerald-500 hover:bg-emerald-600 text-white rounded-lg text-sm font-black uppercase tracking-widest shadow-2xl shadow-emerald-500/30 transition-all disabled:opacity-50 disabled:grayscale flex items-center gap-4 group"
        >
          <LucideCheckCircle v-if="!processing" class="w-6 h-6 group-hover:scale-110 transition-transform" />
          <LucideRefreshCw v-else class="w-6 h-6 animate-spin" />
          Konfirmasi Penerimaan
        </button>
      </div>
    </div>

    <div v-else class="glass rounded-lg p-20 text-center space-y-6">
      <div class="w-20 h-20 bg-red-50 dark:bg-red-900/20 rounded-full flex items-center justify-center text-red-500 mx-auto">
        <LucideXCircle class="w-10 h-10" />
      </div>
      <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Data Tidak Ditemukan</h3>
      <p class="text-slate-500 max-w-md mx-auto">Mohon maaf, data manifest yang Anda cari tidak ditemukan atau sudah diproses sebelumnya.</p>
      <button 
        @click="navigateTo('/intake/inbound')"
        class="px-8 py-3 bg-primary-500 text-white rounded-xl font-black text-xs uppercase tracking-widest shadow-lg shadow-primary-500/20"
      >
        Kembali ke Scanner
      </button>
    </div>

    <!-- Modals Section (Outside main v-if chain) -->
    <div class="modals-container">
      <!-- Confirmation Modal -->
      <Transition name="fade">
        <div v-if="showConfirmModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-slate-900/60 backdrop-blur-sm">
          <div class="bg-white dark:bg-slate-900 rounded-[32px] w-full max-w-lg p-12 shadow-2xl space-y-10" v-motion-pop>
            <div class="flex flex-col items-center text-center space-y-6">
              <div class="w-20 h-20 bg-emerald-50 dark:bg-emerald-900/20 rounded-full flex items-center justify-center text-emerald-500">
                <div class="w-12 h-12 rounded-full border-4 border-emerald-500 flex items-center justify-center">
                  <LucideCheck class="w-6 h-6 stroke-[4px]" />
                </div>
              </div>
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Konfirmasi Penerimaan Dokumen</h3>
            </div>

            <div v-if="manifestDetail" class="bg-slate-50 dark:bg-slate-950/50 rounded-3xl p-8 space-y-4">
              <div class="flex justify-between items-center text-sm">
                <span class="font-bold text-slate-400">Manifest ID</span>
                <span class="font-black text-[#1E3A5F] dark:text-white uppercase">{{ manifestDetail.manifest_no }}</span>
              </div>
              <div class="flex justify-between items-center text-sm">
                <span class="font-bold text-slate-400">Pemohon</span>
                <span class="font-black text-[#1E3A5F] dark:text-white">{{ manifestDetail.owner_name }}</span>
              </div>
              <div class="h-px bg-slate-200 dark:bg-slate-800 my-2"></div>
              <div class="flex justify-between items-center text-sm">
                <span class="font-bold text-slate-400">Jumlah Diterima</span>
                <span class="font-black text-emerald-500">{{ goodCount }} Dokumen</span>
              </div>
              <div class="flex justify-between items-center text-sm">
                <span class="font-bold text-slate-400">Jumlah Ditolak/Rusak</span>
                <span class="font-black text-red-500">{{ damagedCount }} Dokumen</span>
              </div>
            </div>

            <p class="text-xs font-bold text-slate-500 text-center leading-relaxed px-4">
              Setelah diterima, tanggung jawab fisik dokumen berpindah ke Central Document dan masuk Antrean Scan.
            </p>

            <div class="grid grid-cols-2 gap-4 pt-4">
              <button 
                @click="showConfirmModal = false"
                class="py-4 border-2 border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-black text-slate-400 uppercase tracking-widest hover:bg-slate-50 transition-all"
              >
                Batal
              </button>
              <button 
                @click="handleReceive"
                class="py-4 bg-emerald-500 hover:bg-emerald-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-emerald-500/20 transition-all"
              >
                Konfirmasi Terima
              </button>
            </div>
          </div>
        </div>
      </Transition>

      <!-- Reject Manifest Modal -->
      <Transition name="fade">
        <div v-if="showRejectModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-slate-900/60 backdrop-blur-sm">
          <div class="bg-white dark:bg-slate-900 rounded-[32px] w-full max-w-lg p-12 shadow-2xl space-y-10" v-motion-pop>
            <div class="flex flex-col items-center text-center space-y-6">
              <div class="w-20 h-20 bg-red-50 dark:bg-red-900/20 rounded-full flex items-center justify-center text-red-500">
                <LucideXCircle class="w-12 h-12" />
              </div>
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Tolak Manifest Fisik</h3>
              <p class="text-xs font-bold text-slate-400">Berikan alasan penolakan untuk mengembalikan berkas ke pengirim.</p>
            </div>

            <div class="space-y-4">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Alasan Penolakan (Wajib)</p>
              <textarea 
                v-model="rejectReason"
                rows="4"
                class="w-full bg-slate-50 dark:bg-slate-950/50 border border-slate-100 dark:border-slate-800 rounded-3xl p-6 text-sm font-medium text-slate-600 dark:text-slate-300 focus:ring-2 focus:ring-red-500/20 outline-none"
                placeholder="Contoh: Dokumen tidak sesuai manifest, segel rusak..."
              ></textarea>
            </div>

            <div class="grid grid-cols-2 gap-4 pt-4">
              <button 
                @click="showRejectModal = false"
                class="py-4 border-2 border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-black text-slate-400 uppercase tracking-widest hover:bg-slate-50 transition-all"
              >
                Batal
              </button>
              <button 
                @click="handleRejectManifest"
                :disabled="!rejectReason || processing"
                class="py-4 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-red-500/20 transition-all disabled:opacity-50"
              >
                Tolak Sekarang
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { 
  LucideFileSearch, LucideCheck, LucideX, LucideRefreshCw,
  LucideCheckCircle, LucideArrowLeft, LucideXCircle
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'

const route = useRoute()
const { $api } = useApi()
const manifestId = route.params.id

const loading = ref(true)
const processing = ref(false)
const manifestDetail = ref(null)
const mainNotes = ref('')
const showConfirmModal = ref(false)
const showRejectModal = ref(false)
const showToast = ref(false)
const rejectReason = ref('')

// Computed stats
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

const fetchManifestDetail = async () => {
  loading.value = true
  try {
    const res = await $api(`/intake/manifest/${manifestId}`)
    if (res && res.data) {
      manifestDetail.value = res.data
      // Set default physical status to null for forced checking if not already set
      manifestDetail.value.items.forEach(item => {
        if (!item.physical_status) item.physical_status = null
      })
    }
  } catch (err) {
    console.error('Failed to fetch manifest:', err)
  } finally {
    loading.value = false
  }
}

const handleReceive = async () => {
  if (!manifestDetail.value) return
  
  processing.value = true
  showConfirmModal.value = false
  try {
    const payload = {
      manifest_id: manifestDetail.value.id,
      notes: mainNotes.value,
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
      showToast.value = true
      // Wait for toast then redirect
      setTimeout(() => {
        navigateTo('/intake/inbound')
      }, 2000)
    }
  } catch (err) {
    console.error('Receive error:', err)
    alert('Gagal memproses penerimaan: ' + (err.data?.message || err.message))
  } finally {
    processing.value = false
  }
}

const handleRejectManifest = async () => {
  if (!manifestDetail.value || !rejectReason.value) return
  
  processing.value = true
  try {
    const res = await $api(`/intake/reject/${manifestDetail.value.id}`, {
      method: 'POST',
      body: { reason: rejectReason.value }
    })
    
    if (res && res.success) {
      alert('Manifest berhasil ditolak dan dikembalikan.')
      showRejectModal.value = false
      navigateTo('/intake/inbound')
    }
  } catch (err) {
    console.error('Reject error:', err)
    alert('Gagal menolak manifest: ' + (err.data?.message || err.message))
  } finally {
    processing.value = false
  }
}

onMounted(() => {
  fetchManifestDetail()
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

/* Toast Animations */
.slide-fade-enter-active {
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter-from {
  transform: translate(-50%, -40px);
  opacity: 0;
}
.slide-fade-leave-to {
  transform: translate(-50%, -20px);
  opacity: 0;
}
</style>
