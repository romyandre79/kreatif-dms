<template>
  <div class="max-w-7xl mx-auto space-y-10 pb-20" v-motion-fade>
    <!-- Breadcrumbs -->
    <nav class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
      <span>Distribution</span>
      <LucideChevronRight class="w-3 h-3" />
      <span class="text-[#1E3A5F] dark:text-white">Results</span>
    </nav>

    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
        Distribution Result & Audit Trail <span class="text-slate-400">(F-39)</span>
      </h1>
      <button class="px-8 py-3.5 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
        {{ $t('distribution.audit.btn_back') || 'Selesai & Kembali' }}
        <LucideCheckCircle2 class="w-4 h-4" />
      </button>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Left: Success Card -->
      <div class="lg:col-span-8">
        <div class="glass p-10 rounded-lg relative overflow-hidden group border-2 border-green-500/10" v-motion-slide-visible-bottom>
          <div class="flex items-start gap-8">
            <div class="w-20 h-20 rounded-lg bg-green-50 dark:bg-green-900/20 flex items-center justify-center text-green-500 shadow-inner">
              <LucideCheckCircle2 class="w-10 h-10" />
            </div>
            <div class="space-y-2 flex-grow">
              <h2 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Delivery Successful</h2>
              <p class="text-sm font-bold text-slate-500">{{ $t('distribution.audit.success_msg') || 'Dokumen telah berhasil didistribusikan ke penerima eksternal.' }}</p>
            </div>
            <!-- Decorative Check -->
            <LucideCheck class="absolute -right-4 -top-4 w-32 h-32 text-slate-100 dark:text-slate-800/20 -rotate-12 opacity-50" />
          </div>

          <!-- Quick Stats -->
          <div class="grid grid-cols-3 gap-1 mt-12">
            <div class="p-6 bg-slate-50 dark:bg-slate-800/50 rounded-2xl space-y-2">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Channel</p>
              <div class="flex items-center gap-2">
                <LucideLink class="w-3.5 h-3.5 text-blue-500" />
                <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">One-Time Link</span>
              </div>
            </div>
            <div class="p-6 bg-slate-50 dark:bg-slate-800/50 rounded-2xl space-y-2">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Status</p>
              <span class="px-3 py-1 bg-green-50 dark:bg-green-900/30 text-green-500 border border-green-100 dark:border-green-800 rounded-lg text-[9px] font-black uppercase tracking-widest inline-block">Selesai</span>
            </div>
            <div class="p-6 bg-slate-50 dark:bg-slate-800/50 rounded-2xl space-y-2">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Timestamp</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">24 Oct 2023, 14:32:05</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Right: Link Scenario Info -->
      <div class="lg:col-span-4">
        <div class="glass p-10 rounded-lg space-y-8 h-full flex flex-col" v-motion-slide-visible-bottom :delay="100">
          <div class="flex items-center justify-between">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2">
              Link Scenario Info
              <LucideInfo class="w-4 h-4 text-blue-500" />
            </h3>
          </div>

          <div class="space-y-6 flex-grow">
            <div class="flex items-center justify-between">
              <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Token ID</span>
              <span class="px-3 py-1.5 bg-slate-100 dark:bg-slate-800 rounded-lg text-[10px] font-black text-slate-600 dark:text-slate-300 font-mono">TK-992-BX-21</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Expiry Countdown</span>
              <div class="flex items-center gap-1.5 text-red-500">
                <LucideClock class="w-3.5 h-3.5" />
                <span class="text-[11px] font-black uppercase">23h 58m</span>
              </div>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">First Access</span>
              <span class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">Not Accessed Yet</span>
            </div>
          </div>

          <button class="w-full py-4 border-2 border-[#1E3A5F] dark:border-blue-400/30 text-[#1E3A5F] dark:text-blue-400 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-[#1E3A5F] hover:text-white transition-all">
            SALIN TAUTAN (COPY LINK)
          </button>
        </div>
      </div>
    </div>

    <!-- Audit Trail Logs -->
    <div class="space-y-6" v-motion-slide-visible-bottom>
      <div class="flex items-center justify-between px-2">
        <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-3">
          <LucideHistory class="w-5 h-5" />
          Audit Trail Logs
        </h3>
        <div class="flex items-center gap-4">
          <button class="p-2.5 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideFilter class="w-5 h-5" /></button>
          <button class="p-2.5 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideDownload class="w-5 h-5" /></button>
        </div>
      </div>

      <div class="glass rounded-[3rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800">
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                <th class="p-8 pl-10">Performed By</th>
                <th class="p-8">Action</th>
                <th class="p-8">Time</th>
                <th class="p-8">IP / Device</th>
                <th class="p-8 pr-10 text-right">Outcome</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="log in logs" :key="log.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-all cursor-pointer">
                <td class="p-8 pl-10">
                  <div class="flex items-center gap-5">
                    <div :class="`w-10 h-10 rounded-xl flex items-center justify-center text-[10px] font-black shadow-sm ${log.avatarBg} ${log.avatarColor}`">
                      {{ log.initials }}
                    </div>
                    <div>
                      <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ log.user }}</p>
                      <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ log.role }}</p>
                    </div>
                  </div>
                </td>
                <td class="p-8">
                  <div class="flex items-center gap-3">
                    <component :is="log.icon" class="w-4 h-4 text-blue-500" />
                    <p class="text-sm font-bold text-slate-700 dark:text-slate-300">{{ log.action }}</p>
                  </div>
                </td>
                <td class="p-8 text-[11px] font-black text-slate-500 uppercase">{{ log.time }}</td>
                <td class="p-8">
                  <div class="space-y-0.5">
                    <p class="text-[11px] font-black text-slate-700 dark:text-slate-300 font-mono">{{ log.ip }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-tighter">{{ log.device }}</p>
                  </div>
                </td>
                <td class="p-8 pr-10 text-right">
                  <span :class="`px-4 py-1.5 rounded-full text-[9px] font-black uppercase tracking-widest border ${log.outcomeColor}`">
                    {{ log.outcome }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- Pagination -->
        <div class="p-8 bg-slate-50/30 dark:bg-slate-900/30 flex items-center justify-between border-t border-slate-100 dark:border-slate-800">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            Menampilkan 4 dari 4 entri log audit
          </p>
          <div class="flex gap-3">
            <button class="px-6 py-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-400 disabled:opacity-50" disabled>Previous</button>
            <button class="px-6 py-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-400 disabled:opacity-50" disabled>Next</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Info Box Footer -->
    <div class="p-10 bg-slate-50 dark:bg-slate-900/50 rounded-[3rem] border border-slate-100 dark:border-slate-800 flex items-start gap-8 group hover:border-blue-500/30 transition-all" v-motion-slide-visible-bottom>
      <div class="w-16 h-16 rounded-lg bg-[#1E3A5F] text-white flex items-center justify-center shadow-xl shadow-blue-900/20 group-hover:scale-110 transition-transform">
        <LucideShieldCheck class="w-8 h-8" />
      </div>
      <div class="space-y-4 flex-grow">
        <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Informasi Tambahan Distribusi</h4>
        <p class="text-sm font-bold text-slate-500 leading-relaxed max-w-4xl">
          Seluruh jejak audit di atas telah terenkripsi secara permanen (Immutable Ledger). Segala bentuk akses terhadap tautan "One-Time Link" akan tercatat secara real-time pada tabel di atas sebagai "Penerima Akses".
        </p>
        <div class="flex items-center gap-8 pt-2">
          <button class="flex items-center gap-2 text-[10px] font-black text-blue-500 hover:text-blue-600 transition-colors uppercase tracking-widest">
            <LucideDownload class="w-4 h-4" />
            Unduh Laporan Sertifikasi
          </button>
          <button class="flex items-center gap-2 text-[10px] font-black text-blue-500 hover:text-blue-600 transition-colors uppercase tracking-widest">
            <LucidePrinter class="w-4 h-4" />
            Cetak Jejak Audit
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideChevronRight, LucideCheckCircle2, LucideCheck, LucideLink, 
  LucideInfo, LucideClock, LucideHistory, LucideFilter, LucideDownload, 
  LucideShieldCheck, LucidePrinter, LucidePlayCircle, LucideShield, LucideLock
} from 'lucide-vue-next'

const logs = [
  {
    id: 1,
    user: 'Admin Utama',
    role: 'ADMINISTRATOR',
    initials: 'AD',
    avatarBg: 'bg-blue-50',
    avatarColor: 'text-blue-500',
    action: 'Distribution Initiated',
    icon: LucidePlayCircle,
    time: '24 OCT 2023, 14:32:05',
    ip: '192.168.1.184',
    device: 'macOS / Chrome 118',
    outcome: 'SUCCESS',
    outcomeColor: 'bg-green-50 text-green-500 border-green-100'
  },
  {
    id: 2,
    user: 'Admin Utama',
    role: 'ADMINISTRATOR',
    initials: 'AD',
    avatarBg: 'bg-blue-50',
    avatarColor: 'text-blue-500',
    action: 'Token Generated',
    icon: LucideLink,
    time: '24 OCT 2023, 14:31:58',
    ip: '192.168.1.184',
    device: 'macOS / Chrome 118',
    outcome: 'SUCCESS',
    outcomeColor: 'bg-green-50 text-green-500 border-green-100'
  },
  {
    id: 3,
    user: 'System Engine',
    role: 'INTERNAL PROCESS',
    initials: 'SY',
    avatarBg: 'bg-slate-100',
    avatarColor: 'text-slate-500',
    action: 'Integrity Check',
    icon: LucideShield,
    time: '24 OCT 2023, 14:31:50',
    ip: 'Internal Cloud',
    device: 'Server Node #42',
    outcome: 'PASSED',
    outcomeColor: 'bg-green-50 text-green-500 border-green-100'
  },
  {
    id: 4,
    user: 'Auditor 01',
    role: 'AUDITOR',
    initials: 'AU',
    avatarBg: 'bg-indigo-50',
    avatarColor: 'text-indigo-500',
    action: 'Draft Approval',
    icon: LucideLock,
    time: '24 OCT 2023, 14:15:22',
    ip: '10.0.4.88',
    device: 'Windows / Edge 117',
    outcome: 'APPROVED',
    outcomeColor: 'bg-green-50 text-green-500 border-green-100'
  }
]

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";

.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
