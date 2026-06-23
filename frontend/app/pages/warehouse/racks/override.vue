<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Page Header -->
    <div class="space-y-1">
      <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
        Manajemen Override Status Rak
      </h1>
      <p class="text-xs font-bold text-slate-500 max-w-2xl">
        Lakukan override manual pada rak yang memerlukan penandaan khusus. Rak yang dioverride akan dikecualikan dari sistem rekomendasi penempatan otomatis.
      </p>
    </div>

    <!-- Main Content Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
      
      <!-- LEFT: Racks List -->
      <div class="lg:col-span-1 space-y-4">
        <div class="flex items-center justify-between px-2">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Daftar Rak Tersedia</h3>
          <span class="px-2 py-0.5 bg-[#1E3A5F]/10 text-[#1E3A5F] dark:bg-blue-900/30 dark:text-blue-300 rounded text-[8px] font-black">{{ racks.length }} RAKS</span>
        </div>
        
        <div class="space-y-3 h-[600px] overflow-y-auto pr-2 custom-scrollbar">
          <div 
            v-for="rack in racks" 
            :key="rack.id"
            @click="selectRack(rack)"
            :class="[
              'p-4 rounded-xl border cursor-pointer transition-all flex items-start gap-4',
              selectedRack?.id === rack.id 
                ? 'bg-white dark:bg-slate-800 border-primary-500 shadow-md ring-1 ring-primary-500' 
                : 'bg-slate-50/50 dark:bg-slate-900/50 border-slate-200 dark:border-slate-800 hover:border-primary-300'
            ]"
          >
            <!-- Icon -->
            <div :class="[
              'w-10 h-10 rounded-lg flex items-center justify-center shrink-0',
              rack.is_full_override ? 'bg-red-100 text-red-500 dark:bg-red-900/30' : 'bg-slate-200 dark:bg-slate-800 text-[#1E3A5F] dark:text-white'
            ]">
              <LucideLock v-if="rack.is_full_override" class="w-5 h-5" />
              <LucideServer v-else class="w-5 h-5" />
            </div>

            <!-- Details -->
            <div class="flex-1 min-w-0 space-y-1">
              <h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase truncate">{{ rack.name }}</h4>
              <p class="text-[9px] font-bold text-slate-500 uppercase truncate">{{ rack.location_detail || rack.department_name }}</p>
            </div>

            <!-- Badges -->
            <div class="text-right shrink-0 flex flex-col items-end gap-1">
              <span v-if="rack.is_full_override" class="px-2 py-0.5 bg-red-500 text-white rounded text-[8px] font-black uppercase tracking-widest shadow-sm">
                FULL OVERRIDE
              </span>
              <span v-else class="text-xs font-black text-[#1E3A5F] dark:text-white">
                {{ Math.round(((rack.current_docs_count || 0) / (rack.max_docs_capacity || 100)) * 100) }}% FILLED
              </span>
              <span v-if="!rack.is_full_override" class="text-[8px] font-bold text-slate-400">Normal</span>
            </div>
          </div>
        </div>
      </div>

      <!-- MIDDLE: Rack Details -->
      <div class="lg:col-span-1 space-y-4">
        <div class="flex items-center justify-between px-2">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Detail Informasi Rak</h3>
          <LucideInfo class="w-4 h-4 text-slate-400" />
        </div>

        <div v-if="selectedRack" class="glass p-6 rounded-2xl space-y-6">
          <div class="aspect-square bg-slate-100 dark:bg-slate-800 rounded-xl overflow-hidden relative border border-slate-200 dark:border-slate-700">
            <img src="https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?q=80&w=400&auto=format&fit=crop" class="w-full h-full object-cover opacity-80" />
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent flex flex-col justify-end p-6">
              <h2 class="text-3xl font-black text-white uppercase tracking-tighter">{{ selectedRack.name }}</h2>
              <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">MASTER IDENTIFIER</p>
            </div>
            
            <!-- Full Override Watermark -->
            <div 
              v-if="form.is_full_override" 
              class="absolute inset-0 flex items-center justify-center bg-red-950/25 backdrop-blur-[2px] transition-all"
              v-motion-pop
            >
              <div class="bg-red-600/90 text-white border-2 border-white/20 px-10 py-5 rounded-2xl shadow-2xl flex items-center gap-4 transform -rotate-12 select-none">
                <LucideLock class="w-7 h-7 text-white" />
                <span class="text-base font-black uppercase tracking-[0.3em] whitespace-nowrap">FULL / CLOSED</span>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div class="bg-slate-50 dark:bg-slate-800/50 p-4 rounded-xl border border-slate-100 dark:border-slate-700/50">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-1">Status Sensor</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white">Active</p>
            </div>
            <div class="bg-slate-50 dark:bg-slate-800/50 p-4 rounded-xl border border-slate-100 dark:border-slate-700/50">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-1">Kapasitas Terpakai</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white">{{ Math.round(((selectedRack.current_docs_count || 0) / (selectedRack.max_docs_capacity || 100)) * 100) }}% / 100%</p>
            </div>
          </div>
          
          <div class="pt-4 border-t border-slate-100 dark:border-slate-800">
             <div class="flex justify-between items-center text-xs">
                <span class="font-bold text-slate-500">Tipe Dokumen</span>
                <span class="font-black text-[#1E3A5F] dark:text-white">Arsip Fisik</span>
             </div>
          </div>
        </div>

        <div v-else class="glass p-10 rounded-2xl flex flex-col items-center justify-center text-center gap-4 h-[400px]">
           <LucideServer class="w-12 h-12 text-slate-300" />
           <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">Pilih rak dari daftar untuk melihat detail</p>
        </div>
      </div>

      <!-- RIGHT: Configuration -->
      <div class="lg:col-span-2 space-y-4">
        <div class="flex items-center justify-between px-2">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Konfigurasi Manual Override</h3>
        </div>

        <div v-if="selectedRack" class="glass p-8 rounded-2xl space-y-8 h-full">
          <p class="text-[11px] font-bold text-slate-500 italic border-l-2 border-primary-500 pl-4 py-1">
            Perubahan ini bersifat permanen hingga dilakukan reset manual oleh administrator.
          </p>

          <!-- Toggle -->
          <div class="flex items-center justify-between p-6 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl shadow-sm">
            <div>
              <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Set Status: FULL</h4>
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">MANUALLY LOCK STATUS</p>
            </div>
            
            <button 
              @click="form.is_full_override = !form.is_full_override"
              :class="[
                'w-14 h-8 rounded-full transition-colors relative',
                form.is_full_override ? 'bg-[#1E3A5F]' : 'bg-slate-200 dark:bg-slate-700'
              ]"
            >
              <div :class="[
                'w-6 h-6 rounded-full bg-white shadow-md absolute top-1 transition-all',
                form.is_full_override ? 'left-7' : 'left-1'
              ]"></div>
            </button>
          </div>

          <!-- Alert -->
          <div v-if="form.is_full_override" class="p-4 bg-red-50 dark:bg-red-900/10 border-l-4 border-red-500 text-red-600 dark:text-red-400 flex items-center gap-3" v-motion-slide-visible-top>
             <LucideAlertTriangle class="w-5 h-5 shrink-0" />
             <span class="text-[10px] font-black uppercase tracking-widest">MANUAL FULL BADGE ACTIVE</span>
          </div>

          <!-- Reason -->
          <div class="space-y-3">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Alasan Override (Audit Trait)</label>
            <textarea 
              v-model="form.override_reason"
              rows="4"
              class="w-full p-4 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-medium focus:ring-2 focus:ring-primary-500 outline-none resize-none placeholder:text-slate-400"
              placeholder="Contoh: Rak sedang dalam perbaikan fisik atau pembersihan berkala..."
              :disabled="!form.is_full_override"
            ></textarea>
          </div>

          <!-- Recommendations Checkbox -->
          <div class="flex items-start gap-3">
             <div 
               @click="form.is_full_override && (form.exclude_recommendation = !form.exclude_recommendation)"
               :class="[
                 'w-5 h-5 rounded border-2 flex items-center justify-center shrink-0 mt-0.5 transition-colors cursor-pointer',
                 form.exclude_recommendation ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-transparent border-slate-300 dark:border-slate-600',
                 !form.is_full_override ? 'opacity-50 cursor-not-allowed' : ''
               ]"
             >
                <LucideCheck v-if="form.exclude_recommendation" class="w-3 h-3 text-white" />
             </div>
             <div>
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white">Exclude from Recommendation Engine</p>
                <p class="text-[10px] font-medium text-slate-500">Sistem tidak akan menyarankan rak ini untuk penempatan dokumen baru.</p>
             </div>
          </div>

          <!-- Actions -->
          <div class="pt-6 space-y-3">
            <button 
              @click="saveOverride"
              :disabled="isSaving || (!form.is_full_override && !selectedRack.is_full_override)"
              class="w-full py-4 bg-[#1E3A5F] disabled:opacity-50 text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center justify-center gap-2"
            >
              <LucideSave v-if="!isSaving" class="w-4 h-4" />
              <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
              Terapkan Override Status
            </button>
            
            <button 
              @click="resetForm"
              class="w-full py-4 bg-transparent text-slate-500 dark:text-slate-400 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-100 dark:hover:bg-slate-800 transition-all"
            >
              Batalkan Perubahan
            </button>
          </div>
        </div>

        <div v-else class="glass p-10 rounded-2xl flex items-center justify-center h-full">
           <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">Pilih rak terlebih dahulu</p>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { 
  LucideServer, LucideLock, LucideInfo, LucideAlertTriangle,
  LucideCheck, LucideSave, LucideLoader2
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

definePageMeta({
  layout: 'default'
})

const { $api } = useApi()
const racks = ref([])
const selectedRack = ref(null)

const form = ref({
  is_full_override: false,
  override_reason: '',
  exclude_recommendation: true
})

const isSaving = ref(false)

const loadRacks = async () => {
  try {
    const res = await $api('/master/racks')
    racks.value = res.data || []
  } catch (error) {
    console.error('Failed to load racks', error)
  }
}

const selectRack = (rack) => {
  selectedRack.value = rack
  form.value = {
    is_full_override: rack.is_full_override || false,
    override_reason: rack.override_reason || '',
    exclude_recommendation: true // Default true when enabled
  }
}

const resetForm = () => {
  if (selectedRack.value) {
    selectRack(selectedRack.value)
  }
}

const saveOverride = async () => {
  if (!selectedRack.value) return
  isSaving.value = true

  try {
    const res = await $api(`/master/racks/${selectedRack.value.id}/override`, {
      method: 'PUT',
      body: {
        is_full_override: form.value.is_full_override,
        override_reason: form.value.is_full_override ? form.value.override_reason : ''
      }
    })

    const updated = res.data

    // Update local state
    const index = racks.value.findIndex(r => r.id === selectedRack.value.id)
    if (index !== -1 && updated) {
      racks.value[index].is_full_override = updated.is_full_override
      racks.value[index].override_reason = updated.override_reason
    }
    
    if (updated) {
      selectedRack.value.is_full_override = updated.is_full_override
      selectedRack.value.override_reason = updated.override_reason
    }
    
    // Show success (use a toast or notification system here)
    alert('Status override berhasil diperbarui!')
  } catch (error) {
    alert('Gagal menyimpan perubahan: ' + error.message)
  } finally {
    isSaving.value = false
  }
}

onMounted(() => {
  loadRacks()
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
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-700 rounded-full;
}
</style>
