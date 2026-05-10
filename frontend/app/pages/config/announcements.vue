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
          @click="openCreateModal"
          class="flex items-center gap-2 px-6 py-3 bg-[#1E3A5F] hover:bg-[#152943] text-white font-bold rounded-xl shadow-lg shadow-blue-900/20 transition-all"
        >
          <LucidePlus class="w-5 h-5" />
          Tambah Pengumuman
        </button>
      </template>
    </PageHeader>

    <!-- Announcement List -->
    <div class="glass rounded-3xl overflow-hidden border border-slate-200/60 dark:border-slate-800/40 shadow-xl">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-900/50 border-b border-slate-200/60 dark:border-slate-800/40">
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">Status</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">Judul</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">Pesan Singkat</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-200/60 dark:divide-slate-800/40">
          <tr v-for="ann in announcements" :key="ann.id" class="hover:bg-slate-50/50 dark:hover:bg-white/[0.02] transition-colors group">
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <Switch 
                  :model-value="ann.is_active" 
                  @update:model-value="(val) => toggleActive(ann, val)"
                />
                <span :class="`text-[10px] font-black uppercase tracking-widest ${ann.is_active ? 'text-green-500' : 'text-slate-400'}`">
                  {{ ann.is_active ? 'Aktif' : 'Nonaktif' }}
                </span>
              </div>
            </td>
            <td class="px-6 py-4">
              <p class="font-bold text-slate-700 dark:text-slate-200">{{ ann.title }}</p>
              <p class="text-[10px] text-slate-400 font-medium">{{ formatDate(ann.created_at) }}</p>
            </td>
            <td class="px-6 py-4">
              <p class="text-sm text-slate-500 line-clamp-1 max-w-md">{{ ann.message }}</p>
            </td>
            <td class="px-6 py-4 text-right">
              <div class="flex items-center justify-end gap-2">
                <button 
                  @click="editAnnouncement(ann)"
                  class="p-2.5 rounded-xl bg-blue-50 dark:bg-blue-900/20 text-blue-600 hover:bg-blue-600 hover:text-white transition-all shadow-sm"
                >
                  <LucideEdit2 class="w-4 h-4" />
                </button>
                <button 
                  @click="confirmDelete(ann)"
                  class="p-2.5 rounded-xl bg-red-50 dark:bg-red-900/20 text-red-600 hover:bg-red-600 hover:text-white transition-all shadow-sm"
                >
                  <LucideTrash2 class="w-4 h-4" />
                </button>
              </div>
            </td>
          </tr>
          <tr v-if="!announcements.length">
            <td colspan="4" class="px-6 py-20 text-center">
              <div class="flex flex-col items-center gap-4">
                <div class="w-16 h-16 rounded-2xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-slate-400">
                  <LucideMegaphone class="w-8 h-8" />
                </div>
                <p class="text-slate-400 font-bold uppercase tracking-widest text-xs">Belum ada pengumuman</p>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Edit/Create Modal (Teleported to Body to escape clipping) -->
    <ClientOnly>
      <Teleport to="body">
      <Transition name="fade">
        <div v-if="showModal" class="fixed inset-0 z-[9999] flex items-center justify-center p-4 sm:p-10 bg-slate-950/60 backdrop-blur-md">
          <div class="bg-white dark:bg-[#0D121F] w-full max-w-6xl max-h-[92vh] rounded-lg shadow-2xl border border-white/10 flex flex-col overflow-hidden">
            <!-- Modal Header -->
            <div class="p-6 sm:p-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-900/30 rounded-t-[2.5rem]">
              <div class="flex items-center gap-4">
                <div :class="`w-12 h-12 rounded-2xl flex items-center justify-center ${isEditing ? 'bg-blue-500/10 text-blue-500' : 'bg-green-500/10 text-green-500'}`">
                  <component :is="isEditing ? LucideEdit2 : LucidePlus" class="w-6 h-6" />
                </div>
                <div>
                  <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white">{{ isEditing ? 'Edit Pengumuman' : 'Tambah Pengumuman' }}</h3>
                  <p class="text-xs text-slate-400 font-bold uppercase tracking-widest">Konfigurasi Konten & Rilis</p>
                </div>
              </div>
              <button @click="closeModal" class="p-3 rounded-2xl hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 transition-all">
                <LucideX class="w-6 h-6" />
              </button>
            </div>

            <!-- Modal Body (Scrollable) -->
            <div class="flex-1 overflow-y-auto p-6 sm:p-8 space-y-8 custom-scrollbar">
            <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
              <!-- Left: Content Info -->
              <div class="lg:col-span-5 space-y-6">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Judul Pengumuman</label>
                  <input 
                    v-model="form.title"
                    type="text"
                    placeholder="Contoh: Sistem Update v1.2"
                    class="w-full px-6 py-4 rounded-2xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-bold text-slate-700 dark:text-slate-200 shadow-sm"
                  />
                </div>

                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Pesan Singkat Dashboard</label>
                  <textarea 
                    v-model="form.message"
                    rows="3"
                    placeholder="Pesan singkat yang muncul di banner..."
                    class="w-full px-6 py-4 rounded-2xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-medium text-slate-700 dark:text-slate-200 resize-none shadow-sm"
                  ></textarea>
                </div>

                <div class="flex items-center justify-between p-6 bg-slate-50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm">
                  <div>
                    <p class="font-bold text-slate-700 dark:text-slate-200">Aktifkan Sekarang</p>
                    <p class="text-[10px] text-slate-400 font-bold uppercase tracking-widest">Satu pengumuman aktif saja</p>
                  </div>
                  <Switch v-model="form.is_active" />
                </div>

                <!-- Live Preview small -->
                <div class="space-y-4">
                  <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Live Preview</label>
                  <div class="bg-gradient-to-br from-[#1E3A5F] to-[#152943] rounded-3xl p-6 text-white relative overflow-hidden shadow-xl border border-white/5">
                    <div class="relative z-10">
                      <h4 class="font-black text-lg mb-1 leading-tight">{{ form.title || 'Judul' }}</h4>
                      <p class="text-blue-100/70 text-[10px] font-medium leading-relaxed mb-4 line-clamp-2">
                        {{ form.message || 'Pesan...' }}
                      </p>
                      <div class="inline-flex px-4 py-2 bg-white/10 rounded-lg font-black text-[9px] uppercase tracking-widest ring-1 ring-white/20">
                        Catatan Rilis
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Right: WYSIWYG Editor -->
              <div class="lg:col-span-7 space-y-4">
                <div class="flex items-center justify-between ml-1">
                  <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Catatan Rilis (WYSIWYG Editor)</label>
                </div>
                
                <div class="quill-wrapper bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-sm">
                  <ClientOnly>
                    <QuillEditor 
                      v-model:content="form.notes"
                      content-type="html"
                      theme="snow"
                      placeholder="Tulis detail pembaruan di sini..."
                      class="min-h-[400px]"
                        :toolbar="[
                          [{ 'font': [] }, { 'size': ['small', false, 'large', 'huge'] }],
                          ['bold', 'italic', 'underline', 'strike'],
                          [{ 'color': [] }, { 'background': [] }],
                          [{ 'script': 'sub'}, { 'script': 'super' }],
                          [{ 'header': 1 }, { 'header': 2 }, 'blockquote', 'code-block'],
                          [{ 'list': 'ordered'}, { 'list': 'bullet' }, { 'indent': '-1'}, { 'indent': '+1' }],
                          [{ 'direction': 'rtl' }, { 'align': [] }],
                          ['link', 'image'],
                          ['clean']
                        ]"
                    />
                  </ClientOnly>
                </div>
              </div>
            </div>
          </div>

            <!-- Modal Footer -->
            <div class="p-6 sm:p-8 border-t border-slate-100 dark:border-slate-800 flex items-center justify-end gap-4 bg-slate-50/50 dark:bg-slate-900/30 rounded-b-[2.5rem]">
              <button 
                @click="closeModal"
                class="px-8 py-3.5 rounded-xl border border-slate-200 dark:border-slate-700 font-bold text-slate-500 hover:bg-slate-100 dark:hover:bg-slate-800 transition-all"
              >
                Batal
              </button>
              <button 
                @click="saveAnnouncement"
                :disabled="saving"
                class="px-10 py-3.5 bg-primary-500 hover:bg-primary-600 text-white font-black rounded-xl shadow-xl shadow-primary-500/20 transition-all disabled:opacity-50 flex items-center gap-2"
              >
                <LucideLoader2 v-if="saving" class="w-5 h-5 animate-spin" />
                {{ saving ? 'Menyimpan...' : (isEditing ? 'Update Pengumuman' : 'Simpan Pengumuman') }}
              </button>
            </div>
        </div>
      </div>
        </Transition>
      </Teleport>
    </ClientOnly>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideMegaphone, LucidePlus, LucideEdit2, LucideTrash2, 
  LucideLoader2, LucideX, LucideInfo, LucideCheck,
  LucideHeading3, LucideBold, LucideList
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'
import PageHeader from '~/components/PageHeader.vue'
import Switch from '~/components/Switch.vue'
import { parseMarkdown } from '~/utils/format'

const { $api } = useApi()
const saving = ref(false)
const showModal = ref(false)
const isEditing = ref(false)

const form = ref({
  id: null,
  title: '',
  message: '',
  notes: '',
  is_active: false
})

// Fetch announcements
const { data: announcements, pending, refresh } = useAsyncData('announcements-list', async () => {
  try {
    const res = await $api('/master/announcements')
    return res.data || []
  } catch (err) {
    console.error('Fetch error:', err)
    return []
  }
}, { server: false })

const openCreateModal = () => {
  isEditing.value = false
  form.value = {
    id: null,
    title: '',
    message: '',
    notes: '',
    is_active: false
  }
  showModal.value = true
}

const editAnnouncement = (ann) => {
  isEditing.value = true
  form.value = { ...ann }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

const saveAnnouncement = async () => {
  if (!form.value.title || !form.value.message) return
  
  saving.value = true
  try {
    const url = isEditing.value ? `/master/announcements/${form.value.id}` : '/master/announcements'
    const method = isEditing.value ? 'PUT' : 'POST'
    
    await $api(url, {
      method,
      body: form.value
    })
    
    await refresh()
    closeModal()
  } catch (err) {
    console.error(err)
  } finally {
    saving.value = false
  }
}

const toggleActive = async (ann, val) => {
  try {
    await $api(`/master/announcements/${ann.id}`, {
      method: 'PUT',
      body: {
        ...ann,
        is_active: val
      }
    })
    await refresh()
  } catch (err) {
    console.error(err)
  }
}

const confirmDelete = async (ann) => {
  if (confirm('Apakah Anda yakin ingin menghapus pengumuman ini?')) {
    try {
      await $api(`/master/announcements/${ann.id}`, {
        method: 'DELETE'
      })
      await refresh()
    } catch (err) {
      console.error(err)
    }
  }
}

const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  background: white;
}
.dark .glass {
  background: #0D121F;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(148, 163, 184, 0.2);
  border-radius: 10px;
}

/* Quill Dark Mode & Styling */
:deep(.ql-toolbar) {
  border: none !important;
  border-bottom: 1px solid #f1f5f9 !important;
  background: #f8fafc;
  padding: 12px !important;
}
:deep(.ql-container) {
  border: none !important;
  font-family: inherit;
  font-size: 14px;
}
:deep(.ql-editor) {
  padding: 20px !important;
  min-height: 400px;
}
:deep(.ql-editor.ql-blank::before) {
  color: #94a3b8;
  font-style: normal;
  left: 20px;
}

.dark :deep(.ql-toolbar) {
  background: #0f172a;
  border-bottom: 1px solid #1e293b !important;
}
.dark :deep(.ql-toolbar .ql-stroke) {
  stroke: #94a3b8;
}
.dark :deep(.ql-toolbar .ql-fill) {
  fill: #94a3b8;
}
.dark :deep(.ql-toolbar .ql-picker) {
  color: #94a3b8;
}
.dark :deep(.ql-editor) {
  color: #e2e8f0;
}
.dark :deep(.ql-editor.ql-blank::before) {
  color: #475569;
}
</style>
