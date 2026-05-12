<template>
  <div class="flex flex-col h-full bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar: Template List -->
      <aside class="w-80 bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 flex flex-col overflow-hidden">
        <header class="p-8 border-b border-slate-50 dark:border-slate-800">
          <h1 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Email Templates</h1>
          <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-1">Manage system notifications</p>
        </header>

        <div class="flex-grow overflow-y-auto custom-scrollbar p-4 space-y-2">
          <div v-if="loading" class="p-10 flex flex-col items-center justify-center opacity-30">
            <LucideRefreshCw class="w-8 h-8 animate-spin" />
          </div>
          <button 
            v-for="tpl in templates" 
            :key="tpl.id"
            @click="selectTemplate(tpl)"
            class="w-full text-left p-4 rounded-2xl transition-all group relative overflow-hidden"
            :class="selectedTemplate?.id === tpl.id ? 'bg-[#1E3A5F] text-white shadow-xl shadow-blue-900/20' : 'hover:bg-slate-50 dark:hover:bg-slate-800 text-slate-600 dark:text-slate-400'"
          >
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 rounded-xl flex items-center justify-center" :class="selectedTemplate?.id === tpl.id ? 'bg-white/10' : 'bg-slate-100 dark:bg-slate-800 text-slate-400'">
                <LucideMail class="w-4 h-4" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-[11px] font-black uppercase tracking-tight truncate">{{ tpl.name }}</p>
                <p class="text-[9px] font-bold opacity-60 uppercase tracking-widest truncate">{{ tpl.slug }}</p>
              </div>
            </div>
            <div v-if="selectedTemplate?.id === tpl.id" class="absolute right-4 top-1/2 -translate-y-1/2">
              <LucideChevronRight class="w-4 h-4 text-white/50" />
            </div>
          </button>
        </div>
      </aside>

      <!-- Main: Editor -->
      <main class="flex-grow flex flex-col bg-slate-50/50 dark:bg-slate-900/50 overflow-hidden">
        <div v-if="!selectedTemplate" class="flex-grow flex flex-col items-center justify-center space-y-4 opacity-20">
          <LucideLayout class="w-20 h-20" />
          <p class="text-xs font-black uppercase tracking-[0.3em]">Pilih template untuk diedit</p>
        </div>

        <template v-else>
          <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
            <div class="space-y-1">
              <div class="flex items-center gap-2">
                <span class="px-2 py-0.5 rounded bg-blue-500/10 text-blue-500 text-[8px] font-black uppercase tracking-widest">{{ selectedTemplate.slug }}</span>
                <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ selectedTemplate.name }}</h2>
              </div>
              <p class="text-sm font-bold text-slate-500 uppercase tracking-tighter">Sesuaikan subjek dan isi konten email sistem.</p>
            </div>
            <div class="flex items-center gap-4">
              <button @click="saveTemplate" :disabled="saving" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95 disabled:opacity-50 cursor-pointer">
                <LucideRefreshCw v-if="saving" class="w-4 h-4 animate-spin" />
                <LucideSave v-else class="w-4 h-4" />
                Simpan Perubahan
              </button>
            </div>
          </header>

          <div class="flex-grow overflow-hidden flex">
            <!-- Left: Form -->
            <div class="flex-grow overflow-y-auto custom-scrollbar p-10 space-y-8">
              <div class="space-y-3">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Email Subject</label>
                <input 
                  type="text" 
                  v-model="editForm.subject"
                  class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl px-6 py-4 text-sm font-bold text-[#1E3A5F] dark:text-white outline-none focus:ring-4 focus:ring-blue-500/10 transition-all shadow-sm"
                  placeholder="Enter subject..."
                />
              </div>

              <div class="space-y-3">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">HTML Content</label>
                <div class="relative rounded-2xl border border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 shadow-sm overflow-hidden min-h-[400px] flex flex-col">
                  <div class="flex items-center gap-2 p-3 bg-slate-50 dark:bg-slate-900 border-b border-slate-100 dark:border-slate-800">
                    <div class="flex gap-1.5 ml-2">
                      <div class="w-2 h-2 rounded-full bg-red-400"></div>
                      <div class="w-2 h-2 rounded-full bg-amber-400"></div>
                      <div class="w-2 h-2 rounded-full bg-green-400"></div>
                    </div>
                    <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest ml-4">Source Editor</span>
                  </div>
                  <textarea 
                    v-model="editForm.body_html"
                    class="flex-grow w-full p-8 text-xs font-mono text-slate-600 dark:text-slate-300 bg-transparent outline-none resize-none leading-relaxed"
                    spellcheck="false"
                  ></textarea>
                </div>
              </div>
            </div>

            <!-- Right: Preview & Help -->
            <aside class="w-96 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col overflow-hidden">
              <div class="p-8 border-b border-slate-50 dark:border-slate-800">
                <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
                  <LucideEye class="w-4 h-4 text-blue-500" />
                  Live Preview
                </h3>
              </div>
              <div class="flex-grow p-6 overflow-y-auto custom-scrollbar bg-slate-100/50 dark:bg-slate-800/20">
                <div class="bg-white rounded-xl shadow-lg border border-slate-200 overflow-hidden scale-90 origin-top">
                  <div class="p-4 bg-slate-50 border-b border-slate-100">
                    <p class="text-[10px] text-slate-400 font-mono line-clamp-1">Subject: {{ previewSubject }}</p>
                  </div>
                  <div class="p-6 overflow-hidden" v-html="previewBody"></div>
                </div>

                <div class="mt-12 space-y-6">
                  <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1 flex items-center gap-2">
                    <LucideInfo class="w-3.5 h-3.5" />
                    Available Placeholders
                  </h4>
                  <div class="grid grid-cols-1 gap-2">
                    <div v-for="p in selectedTemplate.placeholders" :key="p" class="group flex items-center justify-between p-3 rounded-xl bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100/50 dark:border-blue-800/50">
                      <code class="text-[10px] font-black text-blue-600 dark:text-blue-300">{{ '{{' + p + '}}' }}</code>
                      <button @click="copyPlaceholder(p)" class="p-1.5 opacity-0 group-hover:opacity-100 transition-opacity hover:text-blue-500 cursor-pointer">
                        <LucideCopy class="w-3 h-3" />
                      </button>
                    </div>
                  </div>
                  <p class="text-[10px] font-bold text-slate-400 leading-relaxed italic p-4 bg-slate-50 dark:bg-slate-800/40 rounded-xl">Gunakan placeholder di atas dalam subjek atau isi email untuk menampilkan data dinamis dari sistem.</p>
                </div>
              </div>
            </aside>
          </div>
        </template>
      </main>
    </div>

    <!-- Toast Notification -->
    <Transition name="slide-up">
      <div v-if="toast" class="fixed bottom-10 left-1/2 -translate-x-1/2 z-[100] bg-[#1E3A5F] text-white px-8 py-4 rounded-2xl shadow-2xl flex items-center gap-4">
        <LucideCheckCircle2 class="w-5 h-5 text-green-400" />
        <span class="text-xs font-black uppercase tracking-widest">{{ toast }}</span>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  LucideMail, LucideChevronRight, LucideRefreshCw, LucideSave, 
  LucideLayout, LucideEye, LucideInfo, LucideCopy, LucideCheckCircle2 
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const config = useRuntimeConfig()

const loading = ref(false)
const saving = ref(false)
const templates = ref([])
const selectedTemplate = ref(null)
const toast = ref('')

const editForm = ref({
  subject: '',
  body_html: ''
})

const fetchTemplates = async () => {
  loading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/email-templates`)
    templates.value = res || []
    if (templates.value.length > 0 && !selectedTemplate.value) {
      selectTemplate(templates.value[0])
    }
  } catch (err) {
    console.error('Failed to fetch templates:', err)
  } finally {
    loading.value = false
  }
}

const selectTemplate = (tpl) => {
  selectedTemplate.value = tpl
  editForm.value = {
    subject: tpl.subject,
    body_html: tpl.body_html
  }
}

const saveTemplate = async () => {
  if (!selectedTemplate.value) return
  saving.value = true
  try {
    await $api(`${config.public.apiBase}/master/email-templates/${selectedTemplate.value.id}`, {
      method: 'PUT',
      body: editForm.value
    })
    
    // Update local state
    const idx = templates.value.findIndex(t => t.id === selectedTemplate.value.id)
    if (idx !== -1) {
      templates.value[idx] = { ...templates.value[idx], ...editForm.value }
    }
    
    showToast('Template berhasil disimpan')
  } catch (err) {
    showToast('Gagal menyimpan template', true)
  } finally {
    saving.value = false
  }
}

const showToast = (msg) => {
  toast.value = msg
  setTimeout(() => toast.value = '', 3000)
}

const copyPlaceholder = (p) => {
  navigator.clipboard.writeText(`{{${p}}}`)
  showToast(`Placeholder {{${p}}} disalin`)
}

const previewSubject = computed(() => {
  let sub = editForm.value.subject
  const mockData = {
    fullName: 'Romy Andre',
    docTitle: 'Laporan Keuangan Q1.pdf',
    notes: 'Mohon revisi bagian halaman 5.',
    manifestNo: 'INB-20240512-001',
    itemCount: '12',
    reason: 'Dokumen fisik tidak lengkap.',
    resetLink: '#'
  }
  for (const [k, v] of Object.entries(mockData)) {
    sub = sub.replace(new RegExp(`{{${k}}}`, 'g'), v)
  }
  return sub
})

const previewBody = computed(() => {
  let body = editForm.value.body_html
  const mockData = {
    fullName: 'Romy Andre',
    docTitle: 'Laporan Keuangan Q1.pdf',
    notes: 'Mohon revisi bagian halaman 5.',
    manifestNo: 'INB-20240512-001',
    itemCount: '12',
    reason: 'Dokumen fisik tidak lengkap.',
    resetLink: '#'
  }
  for (const [k, v] of Object.entries(mockData)) {
    body = body.replace(new RegExp(`{{${k}}}`, 'g'), v)
  }
  return body
})

onMounted(() => {
  fetchTemplates()
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(148, 163, 184, 0.2);
  border-radius: 10px;
}

.slide-up-enter-active, .slide-up-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.slide-up-enter-from, .slide-up-leave-to {
  transform: translate(-50%, 100%);
  opacity: 0;
}
</style>
