<template>
  <div class="p-8">
    <div v-if="pending" class="flex flex-col items-center justify-center h-[60vh] gap-4">
      <LucideLoader2 class="w-12 h-12 animate-spin text-primary-500" />
      <p class="text-slate-400 font-bold animate-pulse uppercase tracking-widest text-xs">Memuat Data...</p>
    </div>
    
    <div v-else class="space-y-8 pb-10" v-motion-fade>
    <PageHeader 
      title="Manajemen Pengumuman"
      subtitle="Kelola pesan pembaruan sistem dan catatan rilis untuk dashboard pengguna."
    >
      <template #actions>
        <button 
          @click="saveAll"
          :disabled="saving"
          class="flex items-center gap-2 px-6 py-3 bg-[#1E3A5F] hover:bg-[#152943] text-white font-bold rounded-xl shadow-lg shadow-blue-900/20 transition-all disabled:opacity-50"
        >
          <LucideSave v-if="!saving" class="w-5 h-5" />
          <LucideLoader2 v-else class="w-5 h-5 animate-spin" />
          {{ saving ? 'Menyimpan...' : 'Simpan Perubahan' }}
        </button>
      </template>
    </PageHeader>

    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Main Settings -->
      <div class="xl:col-span-2 space-y-8">
        <div class="glass rounded-2xl p-8 space-y-6">
          <div class="flex items-center gap-4 mb-4">
            <div class="w-12 h-12 rounded-xl bg-primary-50 dark:bg-primary-900/20 flex items-center justify-center text-primary-500">
              <LucideMegaphone class="w-6 h-6" />
            </div>
            <div>
              <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white">Konten Pengumuman</h3>
              <p class="text-xs text-slate-400 font-bold uppercase tracking-widest">Informasi Utama</p>
            </div>
          </div>

          <div class="space-y-4">
            <div class="space-y-2">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Judul Pengumuman</label>
              <input 
                v-model="settings.announcement_title"
                type="text"
                placeholder="Contoh: Sistem Update v1.2"
                class="w-full px-6 py-4 rounded-2xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-bold text-slate-700 dark:text-slate-200 shadow-sm"
              />
            </div>

            <div class="space-y-2">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Pesan Singkat</label>
              <textarea 
                v-model="settings.announcement_message"
                rows="3"
                placeholder="Pesan singkat yang muncul di banner dashboard..."
                class="w-full px-6 py-4 rounded-2xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-medium text-slate-700 dark:text-slate-200 shadow-sm resize-none"
              ></textarea>
            </div>
          </div>
        </div>

        <div class="glass rounded-2xl p-8 space-y-6">
          <div class="flex items-center gap-4 mb-4">
            <div class="w-12 h-12 rounded-xl bg-orange-50 dark:bg-orange-900/20 flex items-center justify-center text-orange-500">
              <LucideFileText class="w-6 h-6" />
            </div>
            <div>
              <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white">Catatan Rilis Detail</h3>
              <p class="text-xs text-slate-400 font-bold uppercase tracking-widest">Markdown Support</p>
            </div>
          </div>

          <div class="space-y-4">
            <div class="space-y-2">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Konten Markdown</label>
              <textarea 
                v-model="settings.announcement_notes"
                rows="10"
                placeholder="Gunakan markdown untuk membuat catatan rilis yang rapi..."
                class="w-full px-6 py-4 rounded-2xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-mono text-sm text-slate-700 dark:text-slate-200 shadow-sm"
              ></textarea>
            </div>
          </div>
        </div>
      </div>

      <!-- Side Panels -->
      <div class="space-y-8">
        <!-- Status Panel -->
        <div class="glass rounded-2xl p-8 space-y-8">
          <h3 class="font-black text-slate-800 dark:text-white">Status Publikasi</h3>
          
          <div class="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800">
            <div class="flex items-center gap-3">
              <div :class="`w-3 h-3 rounded-full ${settings.announcement_active ? 'bg-green-500 shadow-[0_0_10px_rgba(34,197,94,0.5)]' : 'bg-slate-300'}`"></div>
              <span class="text-sm font-bold text-slate-600 dark:text-slate-400">{{ settings.announcement_active ? 'Aktif' : 'Nonaktif' }}</span>
            </div>
            <Switch v-model="settings.announcement_active" />
          </div>

          <p class="text-xs text-slate-400 font-medium leading-relaxed">
            Saat aktif, pengumuman akan muncul di bagian bawah dashboard untuk semua role pengguna.
          </p>
        </div>

        <!-- Preview Panel -->
        <div class="glass rounded-2xl p-8 space-y-6">
          <h3 class="font-black text-slate-800 dark:text-white">Preview Banner</h3>
          
          <div class="bg-gradient-to-br from-[#1E3A5F] to-[#152943] rounded-xl p-6 text-white relative overflow-hidden shadow-xl">
            <div class="relative z-10">
              <div class="w-10 h-10 bg-white/10 rounded-xl flex items-center justify-center mb-4 ring-1 ring-white/20">
                 <LucideInfo class="w-5 h-5 text-white" />
              </div>
              <h4 class="font-black text-lg mb-2 leading-tight">{{ settings.announcement_title || 'Judul Pengumuman' }}</h4>
              <p class="text-blue-100/70 text-xs font-medium leading-relaxed mb-6 line-clamp-2">
                {{ settings.announcement_message || 'Pesan singkat akan muncul di sini...' }}
              </p>
              <button class="w-full py-3 bg-white text-[#1E3A5F] rounded-xl font-black text-[10px] uppercase tracking-widest">
                Catatan Rilis
              </button>
            </div>
            <div class="absolute -right-4 -bottom-4 w-32 h-32 bg-white/5 rounded-full blur-2xl"></div>
          </div>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideMegaphone, 
  LucideSave, 
  LucideInfo, 
  LucideFileText, 
  LucideLoader2 
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'
import PageHeader from '~/components/PageHeader.vue'
import Switch from '~/components/Switch.vue'
// import { useAlert } from '~/composables/useAlert'
const alert = {
  success: (msg) => console.log('Success:', msg),
  error: (msg) => console.log('Error:', msg)
}

const { $api } = useApi()
const saving = ref(false)

const settings = ref({
  announcement_title: '',
  announcement_message: '',
  announcement_notes: '',
  announcement_active: false
})

// Fetch existing settings
const { data: rawSettings, pending } = useAsyncData('dashboard-settings', async () => {
  try {
    const res = await $api('/master/settings/dashboard')
    console.log('[Announcements] API Raw Response:', res)
    return res.data
  } catch (err) {
    console.error('[Announcements] Fetch error:', err)
    return []
  }
}, { server: false })

watch(rawSettings, (newVal) => {
  if (newVal && Array.isArray(newVal)) {
    console.log('[Announcements] Mapping settings...', newVal)
    newVal.forEach(s => {
      // Handle both string and pgtype.Text object
      let val = typeof s.value === 'object' ? s.value.String : s.value
      
      if (s.key === 'announcement_active') {
        settings.value[s.key] = val === 'true' || val === true
      } else {
        // Convert literal \n to actual newlines for the textarea
        if (typeof val === 'string' && val.includes('\\n')) {
          val = val.replace(/\\n/g, '\n')
        }
        settings.value[s.key] = val || ''
      }
    })
    console.log('[Announcements] Settings mapped:', settings.value)
  }
}, { immediate: true })

const saveAll = async () => {
  saving.value = true
  try {
    const promises = Object.entries(settings.value).map(([key, value]) => {
      return $api('/master/settings/dashboard', {
        method: 'POST',
        body: {
          key,
          value: String(value),
          type: key === 'announcement_active' ? 'boolean' : 'string',
          description: getDesc(key)
        }
      })
    })

    await Promise.all(promises)
    alert.success('Pengumuman berhasil diperbarui')
  } catch (err) {
    console.error(err)
    alert.error('Gagal menyimpan perubahan')
  } finally {
    saving.value = false
  }
}

const getDesc = (key) => {
  switch (key) {
    case 'announcement_title': return 'Judul pengumuman di dashboard'
    case 'announcement_message': return 'Pesan pengumuman di dashboard'
    case 'announcement_notes': return 'Detail catatan rilis di dashboard'
    case 'announcement_active': return 'Status aktif pengumuman di dashboard'
    default: return ''
  }
}

definePageMeta({
  layout: 'default'
})
</script>
