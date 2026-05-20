<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
      <!-- Left Column: Details -->
      <div class="lg:col-span-8 space-y-10">
        <!-- Document Review Table -->
        <div class="bg-white rounded-3xl overflow-hidden border border-slate-100 shadow-sm">
          <div class="px-8 py-6 border-b border-slate-50">
            <h2 class="text-sm font-black text-primary-900 uppercase tracking-tight">Review Dokumen Pilihan ({{ selectedDocs.length }})</h2>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="text-[10px] font-bold text-slate-400 uppercase tracking-[0.15em] border-b border-slate-50 bg-slate-50/20">
                <th class="px-8 py-5 w-1/3">NO. DOKUMEN</th>
                <th class="px-8 py-5 w-1/3">JUDUL DOKUMEN</th>
                <th class="px-8 py-5 w-1/3">KATEGORI</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="doc in selectedDocs" :key="doc.id" class="text-xs group hover:bg-slate-50/50 transition-colors">
                <td class="px-8 py-6 font-bold text-slate-400 uppercase tracking-tighter">{{ doc.id }}</td>
                <td class="px-8 py-6 font-black text-slate-700 uppercase tracking-tight">{{ doc.name || doc.title }}</td>
                <td class="px-8 py-6">
                  <span class="text-slate-500 font-bold">{{ doc.type || doc.category_name || 'General' }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Loan Details Form -->
        <div class="bg-white rounded-3xl p-10 space-y-10 border border-slate-100 shadow-sm">
          <h2 class="text-base font-black text-primary-900 uppercase tracking-tight">Detail Peminjaman</h2>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-12 gap-y-10">
            <!-- Pickup Method -->
            <div class="space-y-4">
              <label class="text-[11px] font-black text-slate-400 uppercase tracking-widest">Metode Pengambilan</label>
              <div class="flex flex-col gap-4">
                <label class="flex items-center gap-4 cursor-pointer group">
                  <div class="relative w-5 h-5 rounded-full border-2 border-slate-200 group-hover:border-primary-500 flex items-center justify-center transition-all">
                    <div class="w-2.5 h-2.5 rounded-full bg-primary-500 scale-100 transition-transform"></div>
                  </div>
                  <span class="text-sm font-bold text-slate-600">Courier Service</span>
                </label>
              </div>
            </div>

            <!-- Warehouse Location -->
            <div class="space-y-4">
              <label class="text-[11px] font-black text-slate-400 uppercase tracking-widest">Lokasi Gudang Arsip</label>
              <div class="p-5 bg-primary-50/50 border border-primary-100/50 rounded-2xl flex items-start gap-4">
                <LucideMapPin class="w-5 h-5 text-primary-600 mt-1" />
                <div class="space-y-1">
                  <span class="text-sm font-black text-primary-900 tracking-tight leading-tight block">{{ selectedWarehouse }}</span>
                </div>
              </div>
            </div>

            <!-- Pickup Date -->
            <div class="space-y-4">
              <label class="text-[11px] font-black text-slate-400 uppercase tracking-widest">Tanggal Pengambilan</label>
              <div class="relative">
                <input 
                  type="date" 
                  v-model="form.pickup_date"
                  class="w-full px-6 py-4 bg-slate-50/50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 outline-none transition-all placeholder:text-slate-300"
                />
              </div>
            </div>

            <!-- Time Slot -->
            <div class="space-y-4">
              <label class="text-[11px] font-black text-slate-400 uppercase tracking-widest">Slot Waktu</label>
              <div class="relative">
                <select 
                  v-model="form.time_slot"
                  class="w-full px-6 py-4 bg-slate-50/50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 appearance-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 outline-none transition-all"
                >
                  <option value="">Pilih Slot Waktu</option>
                  <option value="09:00">09:00 - 11:00</option>
                  <option value="13:00">13:00 - 15:00</option>
                </select>
                <LucideChevronDown class="absolute right-6 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 pointer-events-none" />
              </div>
            </div>
          </div>

          <!-- Notes -->
          <div class="space-y-4 pt-4">
            <label class="text-[11px] font-black text-slate-400 uppercase tracking-widest">Catatan Tambahan</label>
            <textarea 
              v-model="form.notes"
              placeholder="Catatan tambahan untuk petugas arsip..." 
              rows="4" 
              class="w-full p-6 bg-slate-50/50 border border-slate-200 rounded-[2rem] text-sm font-bold text-slate-700 outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- Right Column: Summary -->
      <div class="lg:col-span-4 space-y-8" v-motion-slide-visible-bottom>
        <div class="bg-white rounded-[2rem] overflow-hidden border border-slate-100 shadow-2xl shadow-blue-900/5 flex flex-col">
          <div class="bg-[#1E3A5F] px-8 py-6">
            <h2 class="text-base font-black text-white uppercase tracking-tight">Ringkasan Peminjaman</h2>
          </div>
          
          <div class="p-8 flex-1 space-y-10">
            <div class="space-y-6">
              <div class="flex justify-between items-center">
                <span class="text-sm font-bold text-slate-400">Total Dokumen</span>
                <span class="text-sm font-black text-slate-700">{{ selectedDocs.length }} File</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-bold text-slate-400">Durasi Pinjam</span>
                <span class="text-sm font-black text-slate-700">{{ cartStore.loanDuration }} Hari Kalender</span>
              </div>
              <div class="flex justify-between items-start">
                <span class="text-sm font-bold text-slate-400">Estimasi Kembali</span>
                <div class="text-right">
                  <p class="text-sm font-black text-primary-600 leading-tight">{{ estimatedReturnDate }}</p>
                </div>
              </div>
            </div>

            <div class="flex items-center justify-between pt-4">
              <span class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">SECURITY LEVEL</span>
              <span class="px-3 py-1 bg-amber-50 text-amber-600 rounded-full text-[10px] font-black uppercase tracking-widest border border-amber-100/50">CONFIDENTIAL</span>
            </div>

            <div class="pt-6 space-y-4 text-center">
              <button @click="submitLoanRequest" :disabled="isSubmitting" class="w-full flex items-center justify-center gap-4 py-6 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-sm font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all group disabled:opacity-50 disabled:cursor-not-allowed">
                <div class="p-1.5 bg-white/10 rounded-lg">
                   <LucideClipboardCheck class="w-5 h-5 animate-pulse" v-if="isSubmitting" />
                   <LucideClipboardCheck class="w-5 h-5" v-else />
                </div>
                <span>{{ isSubmitting ? 'Submitting Request...' : 'Confirm & Request Loan' }}</span>
              </button>
              
              <button @click="navigateTo('/loans/cart')" class="inline-flex items-center gap-2 text-xs font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors pt-2">
                <LucideArrowLeft class="w-4 h-4" />
                Back to Cart
              </button>
            </div>
          </div>
        </div>

        <!-- Info Box -->
        <div class="p-6 bg-[#EBF5FF] border-l-4 border-[#3B82F6] rounded-r-2xl flex gap-4 shadow-sm">
          <LucideInfo class="w-5 h-5 text-[#3B82F6] shrink-0 mt-0.5" />
          <p class="text-[11px] font-bold text-blue-800 leading-relaxed">
             Permohonan ini akan diteruskan ke Atasan langsung dan Admin Arsip untuk persetujuan.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideArrowLeft, 
  LucideMapPin, 
  LucideChevronDown, 
  LucideClipboardCheck, 
  LucideInfo 
} from 'lucide-vue-next'
import { useCartStore } from '~/stores/cart'
import { useApi } from '~/composables/useApi'
import { useToast } from '~/composables/useToast'

const cartStore = useCartStore()
const selectedDocs = computed(() => cartStore.selectedItems)
const { $api } = useApi()
const config = useRuntimeConfig()
const toast = useToast()
const isSubmitting = ref(false)

const getFutureDate = (daysAhead) => {
  const date = new Date()
  date.setDate(date.getDate() + daysAhead)
  return date.toISOString().split('T')[0]
}

const form = ref({
  pickup_date: getFutureDate(2),
  time_slot: '',
  notes: ''
})

const estimatedReturnDate = computed(() => {
  if (!form.value.pickup_date) return 'Tentukan Tanggal Ambil'
  const pickup = new Date(form.value.pickup_date)
  pickup.setDate(pickup.getDate() + cartStore.loanDuration)
  const options = { day: '2-digit', month: 'short', year: 'numeric' }
  return pickup.toLocaleDateString('id-ID', options)
})

const selectedWarehouse = computed(() => {
  if (!selectedDocs.value || selectedDocs.value.length === 0) {
    return 'Gudang Central Archive'
  }
  
  // Extract location details (branch and company) from each document
  const locations = selectedDocs.value.map(doc => {
    const company = doc.company_name || 'Kreatif Holding'
    const branch = doc.branch_name || ''
    return branch ? `Gudang Arsip ${company} (${branch})` : `Gudang Arsip ${company}`
  })
  
  // Get unique locations
  const uniqueLocations = [...new Set(locations)]
  return uniqueLocations.join(', ')
})

const submitLoanRequest = async () => {
  if (!selectedDocs.value || selectedDocs.value.length === 0) {
    toast.error('Keranjang peminjaman kosong.')
    return
  }

  isSubmitting.value = true
  try {
    const ids = selectedDocs.value.map(doc => doc.id)
    const originalIds = selectedDocs.value.map(doc => doc.original_id).filter(Boolean)
    const documentIds = [...new Set([...ids, ...originalIds])]

    const payload = {
      document_ids: documentIds,
      purpose: 'Peminjaman Dokumen Fisik',
      duration_days: cartStore.loanDuration,
      notes: form.value.notes,
      pickup_date: form.value.pickup_date,
      time_slot: form.value.time_slot
    }

    const res = await $api(`${config.public.apiBase}/loans`, {
      method: 'POST',
      body: payload
    })

    if (res && (res.success || res.status === 'success' || res.code === 201)) {
      toast.success('Permintaan peminjaman berhasil diajukan!')
      cartStore.clearCart()
      navigateTo('/loans/my')
    } else {
      toast.error(res.message || 'Gagal mengajukan peminjaman.')
    }
  } catch (err) {
    console.error(err)
    toast.error(err.data?.message || 'Gagal mengajukan peminjaman. Silakan coba lagi.')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
