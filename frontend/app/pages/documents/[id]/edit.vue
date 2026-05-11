<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-32">
    <!-- Rejection Alert -->
    <div v-if="doc?.rejection_reason" class="glass p-8 rounded-[2rem] border-l-8 border-red-500 bg-red-50/30 flex items-start gap-6 shadow-xl shadow-red-500/5" v-motion-slide-visible-bottom>
      <div class="w-16 h-16 rounded-2xl bg-red-500 flex items-center justify-center text-white shadow-lg shadow-red-500/20">
        <LucideAlertTriangle class="w-8 h-8" />
      </div>
      <div class="space-y-2">
        <p class="text-[10px] font-black text-red-500 uppercase tracking-[0.3em]">{{ $t('approvals.actions.rejection_reason') }}</p>
        <h3 class="text-xl font-black text-slate-800 tracking-tight">{{ doc.rejection_reason }}</h3>
        <p v-if="doc.rejection_notes" class="text-sm font-medium text-slate-500 leading-relaxed italic">"{{ doc.rejection_notes }}"</p>
      </div>
    </div>

    <!-- Header -->
    <div class="flex items-center gap-6" v-motion-fade>
      <button @click="navigateTo('/dashboard')" class="w-12 h-12 rounded-xl bg-white border border-slate-200 flex items-center justify-center text-slate-400 hover:text-[#1E3A5F] hover:border-[#1E3A5F] transition-all">
        <LucideArrowLeft class="w-5 h-5" />
      </button>
      <div>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight">Revisi Dokumen</h1>
        <p class="text-slate-500 font-medium mt-1">Perbaiki data dokumen Anda sesuai arahan manajer dan kirim ulang untuk persetujuan.</p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Main Form -->
      <div class="lg:col-span-2 space-y-8">
        <div class="glass p-10 rounded-[3rem] space-y-10" v-motion-slide-visible-bottom>
          <div class="space-y-8">
            <!-- Title -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-[0.2em] ml-2">Judul Dokumen</label>
              <input 
                v-model="form.title"
                type="text"
                class="w-full bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl p-5 text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all outline-none"
                placeholder="Contoh: Invoice PT. Maju Jaya - Jan 2024"
              />
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <!-- Type -->
              <div class="space-y-3">
                <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-[0.2em] ml-2">Jenis Dokumen</label>
                <div class="relative">
                  <select 
                    v-model="form.type_id"
                    class="w-full bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl p-5 text-sm font-bold appearance-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all outline-none"
                  >
                    <option value="" disabled>Pilih Jenis...</option>
                    <option v-for="t in docTypes" :key="t.id" :value="t.id">{{ t.name }}</option>
                  </select>
                  <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
                </div>
              </div>

              <!-- Date -->
              <div class="space-y-3">
                <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-[0.2em] ml-2">Tanggal Dokumen</label>
                <input 
                  v-model="form.date"
                  type="date"
                  class="w-full bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl p-5 text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all outline-none"
                />
              </div>
            </div>

            <!-- Notes -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-[0.2em] ml-2">Catatan Tambahan</label>
              <textarea 
                v-model="form.notes"
                rows="4"
                class="w-full bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl p-6 text-sm font-medium focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all outline-none resize-none"
                placeholder="Berikan konteks tambahan untuk mempermudah persetujuan..."
              ></textarea>
            </div>
          </div>
        </div>
      </div>

      <!-- Sidebar / Config -->
      <div class="space-y-8">
        <div class="glass p-8 rounded-[2.5rem] space-y-8" v-motion-slide-visible-bottom>
          <div class="space-y-6">
            <!-- Sensitivity -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-[0.2em] ml-2">Sensitivitas</label>
              <div class="grid grid-cols-2 gap-3">
                <button 
                  v-for="cat in ['Internal', 'Confidential', 'Secret']" 
                  :key="cat"
                  @click="form.category = cat"
                  :class="`py-3 px-4 rounded-xl text-[10px] font-black uppercase tracking-widest border-2 transition-all ${form.category === cat ? 'bg-[#1E3A5F] text-white border-[#1E3A5F] shadow-lg shadow-blue-900/20' : 'bg-white text-slate-400 border-slate-100 hover:border-slate-200'}`"
                >
                  {{ cat }}
                </button>
              </div>
            </div>

            <!-- Urgency -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-[0.2em] ml-2">Urgensi</label>
              <div class="grid grid-cols-2 gap-3">
                <button 
                  v-for="urg in ['Normal', 'High', 'Critical']" 
                  :key="urg"
                  @click="form.urgency = urg"
                  :class="`py-3 px-4 rounded-xl text-[10px] font-black uppercase tracking-widest border-2 transition-all ${form.urgency === urg ? 'bg-orange-500 text-white border-orange-500 shadow-lg shadow-orange-500/20' : 'bg-white text-slate-400 border-slate-100 hover:border-slate-200'}`"
                >
                  {{ urg }}
                </button>
              </div>
            </div>
          </div>

          <div class="pt-6 border-t border-slate-100 dark:border-slate-800">
             <button 
                @click="handleUpdate"
                :disabled="submitting"
                class="w-full py-5 bg-[#1E3A5F] text-white rounded-[1.5rem] text-sm font-black uppercase tracking-[0.2em] shadow-2xl shadow-blue-900/30 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 disabled:opacity-50"
              >
                <LucideRefreshCcw v-if="!submitting" class="w-5 h-5" />
                <LucideLoader2 v-else class="w-5 h-5 animate-spin" />
                {{ submitting ? 'Memproses...' : 'Perbarui & Kirim Ulang' }}
              </button>
          </div>
        </div>

        <!-- File Preview (Static) -->
        <div class="glass p-6 rounded-3xl bg-slate-50/50 dark:bg-slate-950/30 border border-slate-200 dark:border-slate-800 flex items-center gap-4">
          <div class="w-12 h-12 rounded-xl bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 flex items-center justify-center text-slate-400">
            <LucideFileText class="w-6 h-6" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-xs font-black text-[#1E3A5F] dark:text-white truncate uppercase">{{ doc?.file_name }}</p>
            <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ (doc?.file_size / 1024 / 1024).toFixed(2) }} MB • PDF</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideArrowLeft, 
  LucideAlertTriangle, 
  LucideChevronDown, 
  LucideRefreshCcw, 
  LucideLoader2,
  LucideFileText
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'

const route = useRoute()
const { $api } = useApi()
const submitting = ref(false)
const docTypes = ref([])

const form = reactive({ 
  title: '', 
  type_id: '', 
  category: 'Internal', 
  urgency: 'Normal',
  date: '', 
  notes: ''
})

// Fetch Document Data
const { data: docRes } = await useAsyncData(`edit-doc-${route.params.id}`, () => 
  $api(`/documents/${route.params.id}`)
)

const doc = computed(() => docRes.value?.data)

// Pre-fill form
watchEffect(() => {
  if (doc.value) {
    form.title = doc.value.title
    form.type_id = doc.value.type_id
    form.category = doc.value.sensitivity || 'Internal'
    form.notes = doc.value.description || ''
    
    // Extract from metadata
    if (doc.value.metadata) {
      form.urgency = doc.value.metadata.urgency || 'Normal'
      form.date = doc.value.metadata.document_date || ''
    }
  }
})

// Fetch Doc Types
const fetchDocTypes = async () => {
  const res = await $api('/master/document-types')
  if (res && res.data) docTypes.value = res.data
}

onMounted(() => {
  fetchDocTypes()
})

const handleUpdate = async () => {
  if (!form.title || !form.type_id) {
    alert('Judul dan Jenis Dokumen wajib diisi')
    return
  }

  submitting.value = true
  try {
    const res = await $api(`/documents/${route.params.id}`, {
      method: 'PUT',
      body: {
        title: form.title,
        description: form.notes,
        type_id: form.type_id,
        sensitivity: form.category,
        urgency: form.urgency,
        document_date: form.date,
        metadata: {
          // Keep other metadata if any
          ...doc.value?.metadata
        }
      }
    })

    if (res && res.success) {
      alert('Dokumen berhasil diperbarui dan dikirim ulang untuk persetujuan.')
      navigateTo('/dashboard')
    } else {
      alert('Gagal memperbarui dokumen: ' + (res?.message || 'Unknown error'))
    }
  } catch (err) {
    console.error(err)
    alert('Update error: ' + (err.data?.message || err.message))
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.glass {
  background-color: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(40px);
  -webkit-backdrop-filter: blur(40px);
  border: 1px solid rgba(255, 255, 255, 1);
}
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #E2E8F0;
  border-radius: 10px;
}
</style>
