<template>
  <div class="flex flex-col h-screen bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <!-- Top Filter Bar -->
    <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-6 flex items-end gap-8 shadow-sm relative z-20">
      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Event Type</label>
        <div class="relative w-64">
          <select class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-[11px] font-black text-[#1E3A5F] dark:text-white appearance-none outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight">
            <option>All Events</option>
            <option>Deleted</option>
            <option>Viewed</option>
            <option>Downloaded</option>
            <option>Modified</option>
          </select>
          <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
        </div>
      </div>

      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">User</label>
        <div class="relative w-64 group">
          <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
          <input type="text" placeholder="Search Administrator..." class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl pl-12 pr-6 py-3.5 text-[11px] font-black outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight" />
        </div>
      </div>

      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Date Range</label>
        <div class="flex items-center gap-3">
          <input type="text" placeholder="mm/dd/yyyy" class="w-36 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl px-5 py-3.5 text-[11px] font-black outline-none focus:ring-2 focus:ring-blue-500/20 transition-all text-center" />
          <input type="text" placeholder="mm/dd" class="w-24 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl px-5 py-3.5 text-[11px] font-black outline-none focus:ring-2 focus:ring-blue-500/20 transition-all text-center" />
        </div>
      </div>

      <button class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
        <LucideFilter class="w-4 h-4" />
        Apply
      </button>
    </header>

    <div class="flex flex-grow overflow-hidden">
      <!-- Main Content: History Log Table -->
      <main class="flex-grow flex flex-col p-10 overflow-hidden gap-8">
        <div class="bg-white dark:bg-slate-900 rounded-[3rem] shadow-sm border border-slate-100 dark:border-slate-800 flex flex-col overflow-hidden flex-grow">
          <!-- Table Header -->
          <div class="px-10 py-8 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Riwayat Aktivitas Sistem</h3>
            <div class="flex items-center gap-6">
              <button class="p-3 text-slate-400 hover:text-blue-500 transition-colors"><LucideRefreshCw class="w-5 h-5" /></button>
              <button class="p-3 text-slate-400 hover:text-blue-500 transition-colors"><LucideDownload class="w-5 h-5" /></button>
            </div>
          </div>

          <!-- Table -->
          <div class="overflow-auto flex-grow custom-scrollbar">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                  <th class="p-8 pl-10">User</th>
                  <th class="p-8 text-center">Event</th>
                  <th class="p-8">Document</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="log in logs" :key="log.id" 
                  class="group transition-all cursor-pointer border-l-4"
                  :class="log.selected ? 'bg-blue-50/50 dark:bg-blue-900/10 border-blue-500' : 'hover:bg-slate-50/30 dark:hover:bg-slate-800/30 border-transparent'"
                >
                  <td class="p-8 pl-10">
                    <div class="flex items-center gap-4">
                      <div :class="`w-10 h-10 rounded-xl flex items-center justify-center text-white text-[10px] font-black ${log.userBg}`">
                        {{ log.initials }}
                      </div>
                      <div class="space-y-0.5">
                        <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ log.userName }}</p>
                        <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ log.role }}</p>
                      </div>
                    </div>
                  </td>
                  <td class="p-8">
                    <div class="flex justify-center">
                      <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${log.badge}`">
                        {{ log.event }}
                      </span>
                    </div>
                  </td>
                  <td class="p-8">
                    <div class="flex items-center gap-3">
                      <LucideFileText class="w-4 h-4 text-slate-400" />
                      <p class="text-[11px] font-bold text-slate-500 dark:text-slate-300 uppercase tracking-tight truncate max-w-[200px]">{{ log.docName }}</p>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Pagination -->
          <div class="px-10 py-8 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Showing 1 to 5 of 1,248 entries</p>
            <div class="flex items-center gap-2">
              <button class="w-8 h-8 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 text-[10px] font-black">1</button>
              <button class="w-8 h-8 rounded-lg bg-[#1E3A5F] text-white text-[10px] font-black">2</button>
              <button class="w-8 h-8 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 text-[10px] font-black">3</button>
              <span class="text-slate-300 mx-2">...</span>
              <button class="w-8 h-8 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 text-[10px] font-black">125</button>
            </div>
          </div>
        </div>
      </main>

      <!-- Right Sidebar: Event Details -->
      <aside class="w-[450px] bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-10 overflow-y-auto custom-scrollbar gap-10">
        <div class="flex items-center justify-between pb-8 border-b border-slate-50 dark:border-slate-800">
          <h2 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Event Details</h2>
          <button class="p-2 text-slate-300 hover:text-red-500 transition-colors"><LucideX class="w-6 h-6" /></button>
        </div>

        <div class="space-y-12">
          <!-- Action Summary -->
          <div class="p-8 bg-red-50 dark:bg-red-900/10 border border-red-100 dark:border-red-800 rounded-[2.5rem] flex items-center gap-6">
            <div class="w-12 h-12 bg-white dark:bg-slate-800 rounded-2xl flex items-center justify-center text-red-500 shadow-sm">
              <LucideTrash2 class="w-6 h-6" />
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-black text-red-400 uppercase tracking-widest">Action Performed</p>
              <p class="text-lg font-black text-red-600 dark:text-red-400 uppercase tracking-tight">Document Deletion</p>
            </div>
          </div>

          <!-- Document Info -->
          <section class="space-y-6">
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
              <LucideFileText class="w-4 h-4 text-blue-500" />
              Document Info
            </h4>
            <div class="p-8 bg-slate-50 dark:bg-slate-800/50 rounded-[2rem] space-y-6 border border-slate-100 dark:border-slate-800">
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">File Name</p>
                <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">SK_Direksi_2023_V2.pdf</p>
              </div>
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Full Repository Path</p>
                <p class="text-[11px] font-bold text-slate-500 dark:text-slate-400 font-mono break-all">/ROOT/LEGAL_INTERNAL/SK/2023/SK_Direksi_2023_V2.pdf</p>
              </div>
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Object GUID</p>
                <p class="text-[11px] font-bold text-blue-500 font-mono uppercase tracking-tighter">8f12-08ab-4431-fe92-13bc98218821</p>
              </div>
            </div>
          </section>

          <!-- User Metadata -->
          <section class="space-y-6">
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
              <LucideUser class="w-4 h-4 text-blue-500" />
              User Metadata
            </h4>
            <div class="grid grid-cols-2 gap-8 px-2">
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Login ID</p>
                <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase">a.saputra.01</p>
              </div>
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Access Level</p>
                <p class="text-[11px] font-black text-red-500 uppercase tracking-tight">Level 4 (Admin)</p>
              </div>
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Department</p>
                <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase">Security Office</p>
              </div>
              <div class="space-y-1">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Terminal ID</p>
                <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase font-mono">WS-JKT-A01</p>
              </div>
            </div>
          </section>

          <!-- Network Details -->
          <section class="space-y-6">
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
              <LucideNetwork class="w-4 h-4 text-blue-500" />
              Network Details
            </h4>
            <div class="space-y-6 px-2">
              <div class="flex justify-between items-center text-[11px]">
                <span class="font-black text-slate-400 uppercase tracking-widest">IP Address</span>
                <span class="font-bold text-[#1E3A5F] dark:text-white font-mono uppercase tracking-tighter">192.168.1.142</span>
              </div>
              <div class="flex justify-between items-center text-[11px]">
                <span class="font-black text-slate-400 uppercase tracking-widest">MAC Address</span>
                <span class="font-bold text-[#1E3A5F] dark:text-white font-mono uppercase tracking-tighter">00:1A:28:3C:40:5E</span>
              </div>
              <div class="space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Browser Agent</p>
                <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed italic">Chrome/119.0.5993.89 Win10 x64</p>
              </div>
            </div>
          </section>

          <!-- Blockchain Evidence -->
          <div class="p-10 bg-blue-50 dark:bg-blue-900/10 border-l-8 border-blue-500 rounded-[2.5rem] space-y-6" v-motion-slide-visible-bottom>
            <div class="flex items-center gap-4">
              <LucideShieldCheck class="w-6 h-6 text-blue-500" />
              <h5 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Blockchain Hash Evidence</h5>
            </div>
            <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed uppercase">
              This log entry has been cryptographically signed and anchored to the internal private chain for immutability.
            </p>
            <div class="space-y-1 pt-2 border-t border-blue-100 dark:border-blue-800">
              <p class="text-[8px] font-black text-blue-400 uppercase tracking-widest">SHA-256</p>
              <p class="text-[10px] font-bold text-blue-500 break-all font-mono tracking-tighter">0x9f3d9b1884c7d659a2...</p>
            </div>
          </div>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideChevronDown, LucideSearch, LucideFilter, LucideRefreshCw, 
  LucideDownload, LucideFileText, LucideX, LucideTrash2, 
  LucideUser, LucideNetwork, LucideShieldCheck 
} from 'lucide-vue-next'

const logs = [
  { id: 1, userName: 'Aditya Saputra', role: 'Sec-Admin-01', initials: 'AS', userBg: 'bg-[#1E3A5F]', event: 'Deleted', badge: 'bg-red-50 text-red-500', docName: 'SK_Direksi_2023_V2.pdf', selected: true },
  { id: 2, userName: 'Rina Maharani', role: 'Fin-Officer', initials: 'RM', userBg: 'bg-blue-400', event: 'Viewed', badge: 'bg-blue-50 text-blue-500', docName: 'Laporan_Audit_Q3.xls' },
  { id: 3, userName: 'Budi Hartono', role: 'IT-Ops', initials: 'BH', userBg: 'bg-amber-400', event: 'Downloaded', badge: 'bg-green-50 text-green-500', docName: 'Log_RFID_Archive.zip' },
  { id: 4, userName: 'Siti Aminah', role: 'Legal-Staff', initials: 'SA', userBg: 'bg-indigo-400', event: 'Modified', badge: 'bg-amber-50 text-amber-600', docName: 'Draft_MoU_Vendor.docx' },
  { id: 5, userName: 'System Process', role: 'Auto-Archive', initials: 'SYS', userBg: 'bg-slate-400', event: 'Archive', badge: 'bg-slate-50 text-slate-500', docName: 'Backup_2022_Q4' }
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
</style>
