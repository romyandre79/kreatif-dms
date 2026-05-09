<template>
  <div class="flex flex-col h-full bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <div class="flex flex-1 overflow-hidden">
      <!-- Main Content Area -->
      <main class="flex-grow flex flex-col overflow-hidden">
        <!-- Header -->
        <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
          <div class="space-y-1">
            <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
              Parameter Sistem
            </h1>
            <p class="text-sm font-bold text-slate-500 uppercase tracking-tighter">Konfigurasi variabel global dan pengaturan operasional platform.</p>
          </div>
          <div class="flex items-center gap-4">
            <button @click="fetchSettings" class="p-3 text-slate-400 hover:text-blue-500 transition-colors cursor-pointer" title="Refresh">
              <LucideRefreshCw :class="{ 'animate-spin': loading }" class="w-5 h-5" />
            </button>
          </div>
        </header>

        <!-- Category Tabs -->
        <div class="px-10 py-4 bg-white dark:bg-slate-900 border-b border-slate-100 dark:border-slate-800 flex items-center gap-2 overflow-x-auto no-scrollbar">
          <button 
            v-for="cat in categories" :key="cat.id"
            @click="activeCategory = cat.id"
            :class="[
              'px-6 py-2.5 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all whitespace-nowrap',
              activeCategory === cat.id 
                ? 'bg-[#1E3A5F] text-white shadow-lg shadow-blue-900/20' 
                : 'text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800 hover:text-slate-600'
            ]"
          >
            {{ cat.label }}
          </button>
        </div>

        <!-- Settings Table -->
        <div class="flex-grow p-8 overflow-auto custom-scrollbar">
          <div class="bg-white dark:bg-slate-900 rounded-[1.5rem] shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden">
            <div v-if="loading" class="p-20 flex flex-col items-center justify-center space-y-4">
               <LucideRefreshCw class="w-10 h-10 text-blue-500 animate-spin" />
               <p class="text-xs font-black text-slate-400 uppercase tracking-widest">Memuat Pengaturan...</p>
            </div>

            <table v-else class="w-full text-left border-collapse">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                  <th class="p-5 pl-10 w-1/4">Parameter Key</th>
                  <th class="p-5">Value / Configuration</th>
                  <th class="p-5 w-1/3">Description</th>
                  <th class="p-5 pr-10 text-right">Action</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="s in filteredSettings" :key="s.key" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/20 transition-all">
                  <td class="p-5 pl-10">
                    <div class="flex flex-col">
                      <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ s.key }}</span>
                      <span class="text-[9px] font-bold text-slate-400 uppercase mt-0.5">{{ s.value_type }}</span>
                    </div>
                  </td>
                  <td class="p-5">
                    <div v-if="editingKey === s.key" class="max-w-md">
                      <!-- Input based on type -->
                      <input v-if="s.value_type === 'string' || s.value_type === 'integer'" 
                             v-model="editValue" 
                             class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-2 text-xs font-bold outline-none shadow-lg shadow-blue-500/10" />
                      
                      <select v-else-if="s.value_type === 'boolean'" 
                              v-model="editValue"
                              class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-2 text-xs font-bold outline-none shadow-lg shadow-blue-500/10">
                        <option value="true">True / Enabled</option>
                        <option value="false">False / Disabled</option>
                      </select>

                      <textarea v-else-if="s.value_type === 'json'" 
                                v-model="editValue" 
                                rows="4"
                                class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-2 text-xs font-mono outline-none shadow-lg shadow-blue-500/10"></textarea>
                    </div>
                    <div v-else>
                      <span v-if="s.value_type === 'boolean'" 
                            :class="s.value === 'true' ? 'text-green-500 bg-green-50' : 'text-slate-400 bg-slate-50'"
                            class="px-3 py-1 rounded-lg text-[10px] font-black uppercase tracking-widest border border-current opacity-80">
                        {{ s.value === 'true' ? 'Enabled' : 'Disabled' }}
                      </span>
                      <span v-else class="text-xs font-bold text-slate-600 dark:text-slate-300 break-all">{{ s.value }}</span>
                    </div>
                  </td>
                  <td class="p-5">
                    <p class="text-[11px] font-medium text-slate-400 leading-relaxed">{{ s.description || 'No description available' }}</p>
                  </td>
                  <td class="p-5 pr-10 text-right">
                    <div v-if="editingKey === s.key" class="flex justify-end gap-2">
                       <button @click="saveSetting(s)" class="p-2.5 bg-blue-500 text-white rounded-xl hover:bg-blue-600 shadow-lg shadow-blue-500/20 active:scale-95 transition-all"><LucideSave class="w-4 h-4" /></button>
                       <button @click="editingKey = ''" class="p-2.5 bg-slate-200 text-slate-500 rounded-xl hover:bg-slate-300 active:scale-95 transition-all"><LucideX class="w-4 h-4" /></button>
                    </div>
                    <button v-else @click="startEdit(s)" class="p-2.5 text-slate-400 hover:text-blue-500 hover:bg-white dark:hover:bg-slate-800 rounded-xl transition-all shadow-sm border border-slate-100 dark:border-slate-800">
                      <LucidePencil class="w-4 h-4" />
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>

      <!-- Help Sidebar -->
      <aside class="w-80 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-8 overflow-y-auto custom-scrollbar">
        <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3 mb-8">
          <LucideInfo class="w-5 h-5 text-blue-500" />
          SISTEM PARAMETER
        </h3>
        <div class="space-y-6">
          <div class="p-6 rounded-2xl bg-slate-50 dark:bg-slate-800/40 border border-slate-100 dark:border-slate-800 space-y-3">
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Kategori General</h4>
            <p class="text-[10px] font-bold text-slate-500 leading-relaxed">Berisi informasi dasar sistem seperti Nama Aplikasi dan Versi yang ditampilkan pada dokumen manifest.</p>
          </div>
          <div class="p-6 rounded-2xl bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100 dark:border-blue-900/30 space-y-3">
            <h4 class="text-[11px] font-black text-blue-600 dark:text-blue-400 uppercase tracking-tight">Peringatan Keamanan</h4>
            <p class="text-[10px] font-bold text-blue-500/80 leading-relaxed">Perubahan pada parameter sistem dapat berdampak langsung pada operasional seluruh pengguna. Pastikan nilai yang dimasukkan valid.</p>
          </div>
        </div>

        <div class="mt-auto pt-8 border-t border-slate-100 dark:border-slate-800">
          <button @click="applyChanges" class="w-full py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/30 flex items-center justify-center gap-3 transition-all active:scale-95 group">
            <LucideRefreshCw class="w-4 h-4 group-hover:rotate-180 transition-transform duration-500" />
            Reload App Cache
          </button>
        </div>
      </aside>
    </div>

    <!-- Error Toast -->
    <div v-if="error" class="fixed bottom-10 right-10 z-[100] bg-red-500 text-white px-6 py-4 rounded-2xl shadow-2xl flex items-center gap-4 animate-bounce">
       <LucideBan class="w-6 h-6" />
       <div>
         <p class="text-xs font-black uppercase tracking-widest">Notification</p>
         <p class="text-sm font-bold">{{ error }}</p>
       </div>
       <button @click="error = ''" class="ml-4 hover:opacity-70"><LucideX class="w-4 h-4" /></button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { 
  LucideSettings2, LucideRefreshCw, LucideInfo, LucidePencil, 
  LucideSave, LucideX, LucideBan, LucideShieldCheck
} from 'lucide-vue-next'

const loading = ref(false)
const error = ref('')
const settings = ref([])
const activeCategory = ref('general')
const editingKey = ref('')
const editValue = ref('')

const categories = [
  { id: 'general', label: 'General Info' },
  { id: 'ocr', label: 'OCR Engine' },
  { id: 'storage', label: 'Storage Config' },
  { id: 'security', label: 'System Security' },
  { id: 'api', label: 'API & Integration' }
]

const fetchSettings = async () => {
  loading.value = true
  try {
    const config = useRuntimeConfig()
    const auth = useAuthStore()
    const res = await $fetch(`${config.public.apiBase}/master/settings/${activeCategory.value}`, {
      headers: { Authorization: `Bearer ${auth.accessToken}` }
    })
    settings.value = res.data || []
  } catch (err) {
    error.value = 'Gagal memuat pengaturan sistem'
  } finally {
    loading.value = false
  }
}

const filteredSettings = computed(() => settings.value)

const startEdit = (s) => {
  editingKey.value = s.key
  editValue.value = s.value
}

const saveSetting = async (s) => {
  try {
    const config = useRuntimeConfig()
    const auth = useAuthStore()
    
    await $fetch(`${config.public.apiBase}/master/settings/${activeCategory.value}`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${auth.accessToken}` },
      body: {
        category: activeCategory.value,
        key: s.key,
        value: editValue.value,
        value_type: s.value_type,
        description: s.description
      }
    })
    
    editingKey.value = ''
    fetchSettings()
    error.value = 'Pengaturan berhasil diperbarui'
    setTimeout(() => error.value = '', 3000)
  } catch (err) {
    error.value = 'Gagal menyimpan pengaturan'
  }
}

const applyChanges = () => {
  window.location.reload()
}

watch(activeCategory, fetchSettings)

onMounted(fetchSettings)
</script>

<style scoped>
@reference "../../assets/css/main.css";
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { @apply bg-transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { @apply bg-slate-200 dark:bg-slate-800 rounded-full; }
.no-scrollbar::-webkit-scrollbar { display: none; }
</style>
