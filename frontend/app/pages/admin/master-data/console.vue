<template>
  <div class="flex h-screen bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <!-- Left Sidebar: Entity Navigation -->
    <aside class="w-72 bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 flex flex-col">
      <div class="p-8">
        <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] mb-8">Entitas Master</h3>
        <nav class="space-y-2">
          <button v-for="item in menuItems" :key="item.name" 
            class="w-full flex items-center justify-between px-6 py-4 rounded-2xl transition-all group"
            :class="item.active ? 'bg-[#1E3A5F] text-white shadow-lg shadow-blue-900/20' : 'text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800'"
          >
            <div class="flex items-center gap-4">
              <component :is="item.icon" class="w-5 h-5 transition-transform group-hover:scale-110" />
              <span class="text-sm font-black tracking-tight uppercase">{{ item.name }}</span>
            </div>
            <span :class="`text-[10px] font-black px-2 py-0.5 rounded-lg ${item.active ? 'bg-white/20' : 'bg-slate-100 dark:bg-slate-800 text-slate-400'}`">
              {{ item.count }}
            </span>
          </button>
        </nav>
      </div>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-grow flex flex-col overflow-hidden">
      <!-- Header -->
      <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">Manajemen PT</h1>
          <p class="text-sm font-bold text-slate-500 uppercase tracking-tighter">Pengaturan entitas legal dan struktur organisasi tingkat atas.</p>
        </div>
        <div class="flex items-center gap-4">
          <button class="px-6 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2">
            <LucideUpload class="w-4 h-4" />
            Import CSV
          </button>
          <button class="px-8 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
            <LucidePlus class="w-4 h-4" />
            Add Row
          </button>
        </div>
      </header>

      <!-- Table Section -->
      <div class="flex-grow p-10 overflow-auto custom-scrollbar">
        <div class="bg-white dark:bg-slate-900 rounded-[3rem] shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden">
          <!-- Filter Bar -->
          <div class="px-8 py-6 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <div class="relative w-96 group">
              <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
              <input type="text" placeholder="Filter entitas..." class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl pl-12 pr-6 py-3 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20 transition-all" />
            </div>
            <div class="flex items-center gap-4">
              <button class="p-3 text-slate-400 hover:text-blue-500 transition-colors"><LucideDownload class="w-5 h-5" /></button>
              <button class="p-3 text-slate-400 hover:text-blue-500 transition-colors"><LucideSettings2 class="w-5 h-5" /></button>
            </div>
          </div>

          <!-- Table -->
          <div class="overflow-x-auto">
            <table class="w-full text-left">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                  <th class="p-8 pl-10">Entity ID</th>
                  <th class="p-8">Nama Perusahaan</th>
                  <th class="p-8">NPWP Status</th>
                  <th class="p-8">Lokasi Pusat</th>
                  <th class="p-8">Status</th>
                  <th class="p-8 pr-10 text-right">Action</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="row in companies" :key="row.id" 
                  class="group transition-all cursor-pointer"
                  :class="row.id === 'PT-NEW-003' ? 'bg-blue-50/50 dark:bg-blue-900/10' : 'hover:bg-slate-50/50 dark:hover:bg-slate-800/30'"
                >
                  <td class="p-8 pl-10">
                    <span class="text-sm font-black text-blue-500 font-mono tracking-tighter">{{ row.id }}</span>
                  </td>
                  <td class="p-8">
                    <div v-if="row.editing" class="relative group">
                      <input type="text" v-model="row.name" class="w-full bg-white dark:bg-slate-800 border-2 border-blue-500 rounded-xl px-4 py-3 text-sm font-black text-[#1E3A5F] dark:text-white uppercase outline-none shadow-lg shadow-blue-900/10" />
                    </div>
                    <p v-else class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ row.name }}</p>
                  </td>
                  <td class="p-8">
                    <span :class="`px-3 py-1.5 rounded-lg text-[9px] font-black uppercase tracking-widest border ${row.npwpStatus === 'VERIFIED' ? 'bg-green-50 text-green-500 border-green-100' : 'bg-red-50 text-red-500 border-red-100'}`">
                      {{ row.npwpStatus }}
                    </span>
                  </td>
                  <td class="p-8">
                    <p class="text-sm font-bold text-slate-500 uppercase tracking-tight">{{ row.location }}</p>
                  </td>
                  <td class="p-8">
                    <div class="flex items-center gap-2">
                      <div :class="`w-2 h-2 rounded-full ${row.status === 'Aktif' ? 'bg-green-500' : 'bg-amber-500'}`"></div>
                      <span class="text-[11px] font-bold text-slate-700 dark:text-slate-300">{{ row.status }}</span>
                    </div>
                  </td>
                  <td class="p-8 pr-10 text-right">
                    <div v-if="row.editing" class="flex justify-end gap-3">
                      <button class="p-2.5 bg-blue-500 text-white rounded-xl hover:bg-blue-600 shadow-sm"><LucideSave class="w-4 h-4" /></button>
                      <button class="p-2.5 bg-red-500 text-white rounded-xl hover:bg-red-600 shadow-sm"><LucideTrash2 class="w-4 h-4" /></button>
                    </div>
                    <div v-else class="flex justify-end gap-3 opacity-0 group-hover:opacity-100 transition-opacity">
                      <button class="p-2.5 text-slate-400 hover:text-blue-500 hover:bg-white rounded-xl transition-all"><LucidePencil class="w-4 h-4" /></button>
                      <button class="p-2.5 text-slate-400 hover:text-red-500 hover:bg-white rounded-xl transition-all"><LucideBan class="w-4 h-4" /></button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Pagination -->
          <div class="px-10 py-8 bg-slate-50/50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Menampilkan 1-10 dari 42 entitas</p>
            <div class="flex items-center gap-2">
              <button class="w-8 h-8 flex items-center justify-center text-slate-400 hover:text-blue-500 transition-colors"><LucideChevronLeft class="w-4 h-4" /></button>
              <button class="w-8 h-8 rounded-lg bg-[#1E3A5F] text-white text-[10px] font-black">1</button>
              <button class="w-8 h-8 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 text-[10px] font-black">2</button>
              <button class="w-8 h-8 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 text-[10px] font-black">3</button>
              <button class="w-8 h-8 flex items-center justify-center text-slate-400 hover:text-blue-500 transition-colors"><LucideChevronRight class="w-4 h-4" /></button>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Right Sidebar: Rules & Changes -->
    <aside class="w-80 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-8 overflow-y-auto custom-scrollbar">
      <!-- Validation Rules -->
      <section class="space-y-8 mb-12">
        <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
          <LucideClipboardCheck class="w-5 h-5 text-blue-500" />
          Validation Rules
        </h3>
        <div class="space-y-6">
          <div v-for="rule in rules" :key="rule.title" class="p-6 rounded-2xl border-l-4 border-red-500 bg-red-50/50 dark:bg-red-900/10 space-y-3 relative group">
            <LucideInfo class="absolute top-4 right-4 w-3 h-3 text-slate-300" />
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ rule.title }}</h4>
            <p class="text-[10px] font-bold text-slate-500 leading-relaxed">{{ rule.desc }}</p>
            <div class="h-1 bg-slate-200 dark:bg-slate-800 rounded-full overflow-hidden">
              <div class="h-full bg-red-500 w-[70%]"></div>
            </div>
            <div class="flex justify-between items-center text-[8px] font-black uppercase tracking-widest text-red-500">
              <span>Mandatory</span>
              <span>Active</span>
            </div>
          </div>
          <!-- Second Rule Example -->
          <div class="p-6 rounded-2xl border-l-4 border-blue-500 bg-blue-50/50 dark:bg-blue-900/10 space-y-4">
             <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Tax Compliance</h4>
             <p class="text-[10px] font-bold text-slate-500 leading-relaxed">NPWP wajib dilampirkan dalam format PDF (max 2MB) untuk setiap entitas PT baru.</p>
             <div class="flex gap-2">
               <span class="px-2 py-0.5 bg-blue-100 dark:bg-blue-800 text-[8px] font-black text-blue-600 dark:text-blue-300 rounded uppercase">Regex Check</span>
               <span class="px-2 py-0.5 bg-blue-100 dark:bg-blue-800 text-[8px] font-black text-blue-600 dark:text-blue-300 rounded uppercase">LDAP</span>
             </div>
          </div>
        </div>
      </section>

      <!-- Pending Changes -->
      <section class="space-y-8 mt-auto pt-8 border-t border-slate-100 dark:border-slate-800">
        <h3 class="text-xs font-black text-slate-400 uppercase tracking-[0.2em]">Pending Changes</h3>
        <div class="space-y-4">
          <div v-for="change in pendingChanges" :key="change.title" class="flex items-start gap-4">
            <div :class="`w-8 h-8 rounded-lg flex items-center justify-center ${change.type === 'update' ? 'bg-blue-50 text-blue-500' : 'bg-red-50 text-red-500'}`">
              <component :is="change.type === 'update' ? LucidePencil : LucideTrash" class="w-3.5 h-3.5" />
            </div>
            <div>
              <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ change.title }}</p>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ change.meta }}</p>
            </div>
          </div>
        </div>
        <div class="space-y-4">
          <button class="w-full py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/30 flex items-center justify-center gap-3 transition-all active:scale-95 group">
            <LucideRefreshCw class="w-4 h-4 group-hover:rotate-180 transition-transform duration-500" />
            Publish Changes
          </button>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest text-center">Sync ke production tersedia 24/7</p>
        </div>
      </section>
    </aside>
  </div>
</template>

<script setup>
import { 
  LucideBuilding2, LucideUsers, LucideFileText, LucideArchive, 
  LucideClock, LucideUpload, LucidePlus, LucideSearch, LucideDownload, 
  LucideSettings2, LucideCheckCircle2, LucidePencil, LucideBan, 
  LucideSave, LucideTrash2, LucideChevronLeft, LucideChevronRight,
  LucideClipboardCheck, LucideInfo, LucideRefreshCw, LucideTrash
} from 'lucide-vue-next'

const menuItems = [
  { name: 'PT (Perusahaan)', icon: LucideBuilding2, count: 42, active: true },
  { name: 'Departemen', icon: LucideUsers, count: 128 },
  { name: 'Tipe Dokumen', icon: LucideFileText, count: 15 },
  { name: 'Lokasi Arsip', icon: LucideArchive, count: 24 },
  { name: 'Kode Retensi', icon: LucideClock, count: 8 }
]

const companies = [
  { id: 'PT-AKR-001', name: 'PT Akiradata Utama', npwpStatus: 'VERIFIED', location: 'Jakarta Selatan', status: 'Aktif' },
  { id: 'PT-AKR-002', name: 'Akiradata Logistics', npwpStatus: 'VERIFIED', location: 'Surabaya', status: 'Aktif' },
  { id: 'PT-NEW-003', name: 'Akiradata Digital Solusi', npwpStatus: 'PENDING', location: 'Draft Mode', status: 'Editing', editing: true }
]

const rules = [
  { 
    title: 'Naming Convention', 
    desc: "Prefix harus diawali dengan 'PT' dan minimal 3 kata unik untuk menghindari duplikasi database."
  }
]

const pendingChanges = [
  { title: 'Update: Akiradata Digital Solusi', meta: 'Modified 2 mins ago by Admin', type: 'update' },
  { title: 'Remove: Legacy Test Corp', meta: 'Scheduled for deletion', type: 'remove' }
]

definePageMeta({
  layout: 'none'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}

.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-sm;
}
</style>
