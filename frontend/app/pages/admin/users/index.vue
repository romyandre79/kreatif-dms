<template>
  <div class="flex h-[calc(100vh-theme(spacing.32))] gap-10" v-motion-fade>
    <!-- Left Section: Main Content -->
    <div class="flex-grow space-y-10 overflow-y-auto pr-4 pb-20 custom-scrollbar">
      <!-- Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('users.list.title') }}</h1>
          <p class="text-xs font-bold text-slate-500">{{ $t('users.list.subtitle') }}</p>
        </div>
        <div class="flex items-center gap-4">
          <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
            <LucideFilter class="w-4 h-4" />
            {{ $t('users.list.header.filter') }}
          </button>
          <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3">
            <LucideUserPlus class="w-4 h-4" />
            {{ $t('users.list.header.add') }}
          </button>
        </div>
      </div>

      <!-- Stats Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div v-for="s in [
          { id: 'total', val: '1,248', color: 'text-blue-500', icon: LucideUsers },
          { id: 'active', val: '1,120', color: 'text-green-500', icon: LucideCheckCircle2 },
          { id: 'privileged', val: '42', color: 'text-amber-500', icon: LucideShieldCheck },
          { id: 'inactive', val: '128', color: 'text-red-500', icon: LucideUserX }
        ]" :key="s.id" class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm group hover:scale-105 transition-all cursor-pointer">
          <div class="flex items-center justify-between">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`users.list.stats.${s.id}`) }}</p>
            <div :class="`w-8 h-8 rounded-lg flex items-center justify-center bg-slate-50 dark:bg-slate-800 ${s.color} shadow-sm border border-slate-100 dark:border-slate-800 group-hover:bg-[#1E3A5F] group-hover:text-white transition-all`">
              <component :is="s.icon" class="w-4 h-4" />
            </div>
          </div>
          <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ s.val }}</p>
        </div>
      </div>

      <!-- User Table -->
      <div class="glass rounded-[3rem] overflow-hidden shadow-sm">
        <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('users.list.table.title') }}</h3>
          <div class="flex items-center gap-6">
            <span class="flex items-center gap-2 text-[8px] font-black text-green-500 uppercase tracking-widest">
              <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
              {{ $t('users.list.table.online') }}
            </span>
            <span class="flex items-center gap-2 text-[8px] font-black text-slate-300 uppercase tracking-widest">
              <span class="w-1.5 h-1.5 bg-slate-300 rounded-full"></span>
              {{ $t('users.list.table.offline') }}
            </span>
          </div>
        </div>
        <table class="w-full text-left">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800">
              <th class="p-6 pl-10">{{ $t('users.list.table.cols.role') }}</th>
              <th class="p-6">{{ $t('users.list.table.cols.dept') }}</th>
              <th class="p-6">{{ $t('users.list.table.cols.status') }}</th>
              <th class="p-6">{{ $t('users.list.table.cols.login') }}</th>
              <th class="p-6 pr-10 text-right">{{ $t('users.list.table.cols.reporting') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="u in [
              { id: 1, role: 'IT MANAGER', dept: 'Information Technology', status: 'ok', time: 'Hari ini, 09:42', ip: '192.168.1.110', boss: 'Direktur Utama', online: true },
              { id: 2, role: 'STAFF ADMIN', dept: 'Finance & Accounting', status: 'ok', time: 'Kemarin, 17:15', ip: '10.0.12.45', boss: 'Finance Manager', online: false },
              { id: 3, role: 'STAFF ADMIN', dept: 'General Affairs', status: 'fail', time: '05 Feb 2024', ip: 'Akun Dinonaktifkan', boss: 'GA Supervisor', online: false },
              { id: 4, role: 'ARCHIVE OFFICER', dept: 'Operations', status: 'ok', time: '08 Feb 2024', ip: '', boss: 'Ops Manager', online: false }
            ]" :key="u.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer" @click="selectedUser = u">
              <td class="p-6 pl-10">
                <div class="px-4 py-2 bg-slate-50 dark:bg-slate-800 rounded-lg text-[9px] font-black text-[#1E3A5F] dark:text-white inline-block border border-slate-100 dark:border-slate-800 group-hover:bg-[#1E3A5F] group-hover:text-white transition-all uppercase tracking-tight">
                  {{ u.role }}
                </div>
              </td>
              <td class="p-6 text-[11px] font-bold text-slate-500 uppercase">{{ u.dept }}</td>
              <td class="p-6">
                <span :class="`px-3 py-1 rounded text-[8px] font-black uppercase tracking-widest border ${u.status === 'ok' ? 'bg-green-50 text-green-500 border-green-100' : 'bg-red-50 text-red-500 border-red-100'}`">
                  {{ u.status === 'ok' ? 'SELESAI' : 'GAGAL' }}
                </span>
              </td>
              <td class="p-6">
                <div class="space-y-0.5">
                  <p class="text-[11px] font-black text-slate-800 dark:text-white uppercase tracking-tight">{{ u.time }}</p>
                  <p class="text-[8px] font-bold text-slate-400 uppercase font-mono">{{ u.ip }}</p>
                </div>
              </td>
              <td class="p-6 pr-10 text-right">
                <div class="flex items-center justify-end gap-3">
                  <span class="text-[10px] font-black text-slate-400 uppercase">{{ u.boss }}</span>
                  <div class="w-8 h-8 rounded-full bg-slate-100 dark:bg-slate-800 border border-white dark:border-slate-700 shadow-sm flex items-center justify-center">
                    <LucideUser class="w-4 h-4 text-slate-300" />
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="p-8 flex items-center justify-between border-t border-slate-50 dark:border-slate-800 text-[10px] font-black text-slate-400 uppercase tracking-widest">
          <span>MENAMPILKAN 1-10 DARI 1,248 PENGGUNA</span>
          <div class="flex items-center gap-4">
            <LucideChevronLeft class="w-4 h-4 cursor-pointer hover:text-[#1E3A5F]" />
            <div class="flex items-center gap-2">
              <span class="w-8 h-8 rounded-lg bg-[#1E3A5F] text-white flex items-center justify-center">1</span>
              <span class="w-8 h-8 rounded-lg hover:bg-slate-50 flex items-center justify-center cursor-pointer">2</span>
              <span class="w-8 h-8 rounded-lg hover:bg-slate-50 flex items-center justify-center cursor-pointer">3</span>
              <span>...</span>
              <span class="w-8 h-8 rounded-lg hover:bg-slate-50 flex items-center justify-center cursor-pointer">125</span>
            </div>
            <LucideChevronRight class="w-4 h-4 cursor-pointer hover:text-[#1E3A5F]" />
          </div>
        </div>
      </div>
    </div>

    <!-- Right Section: Detail Sidebar -->
    <div class="w-[450px] bg-white dark:bg-slate-900 border-l border-slate-100 dark:border-slate-800 overflow-y-auto custom-scrollbar p-10 space-y-12 shrink-0">
      <div class="flex items-center justify-between">
        <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('users.list.detail.title') }}</h3>
        <LucideX class="w-5 h-5 text-slate-300 cursor-pointer hover:text-red-500 transition-all" />
      </div>

      <!-- User Profile Header -->
      <div class="flex flex-col items-center text-center space-y-6 pt-6">
        <div class="relative">
          <div class="w-32 h-32 rounded-[2.5rem] bg-blue-50 dark:bg-blue-900/10 flex items-center justify-center text-4xl font-black text-[#1E3A5F] dark:text-white shadow-xl shadow-blue-900/5">
            BP
          </div>
          <div class="absolute bottom-2 right-2 w-6 h-6 bg-green-500 rounded-full border-4 border-white dark:border-slate-900 animate-pulse shadow-lg"></div>
        </div>
        <div class="space-y-1">
          <h4 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Bambang Pamungkas</h4>
          <p class="text-[10px] font-black text-blue-500 uppercase tracking-widest">IT MANAGER</p>
          <p class="text-[11px] font-bold text-slate-400">b.pamungkas@akiradata.co.id</p>
        </div>
      </div>

      <!-- Action Grid -->
      <div class="grid grid-cols-2 gap-4">
        <button v-for="action in [
          { id: 'reset', icon: LucideRefreshCw },
          { id: 'role', icon: LucideUsers },
          { id: 'reporting', icon: LucideNetwork },
          { id: 'deactivate', icon: LucideUserX, color: 'text-red-500' }
        ]" :key="action.id" class="p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800 flex flex-col items-center justify-center gap-4 group hover:bg-[#1E3A5F] transition-all">
          <component :is="action.icon" :class="`w-5 h-5 ${action.color || 'text-slate-400'} group-hover:text-white transition-all`" />
          <span :class="`text-[8px] font-black uppercase tracking-widest ${action.color || 'text-slate-500'} group-hover:text-white`">{{ $t(`users.list.detail.actions.${action.id}`) }}</span>
        </button>
      </div>

      <!-- System Info -->
      <div class="space-y-8">
        <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">{{ $t('users.list.detail.system.title') }}</h4>
        <div class="space-y-6">
          <div class="flex items-center justify-between px-2">
            <span class="text-[11px] font-bold text-slate-500 uppercase">{{ $t('users.list.detail.system.ldap') }}</span>
            <span class="text-[11px] font-black text-[#1E3A5F] dark:text-white">12 Feb 2024, 08:00</span>
          </div>
          <div class="flex items-center justify-between px-2">
            <span class="text-[11px] font-bold text-slate-500 uppercase">{{ $t('users.list.detail.system.access') }}</span>
            <div class="flex gap-1.5">
              <div class="w-2.5 h-2.5 rounded-full bg-blue-500"></div>
              <div class="w-2.5 h-2.5 rounded-full bg-[#1E3A5F]"></div>
              <div class="w-2.5 h-2.5 rounded-full bg-green-500"></div>
            </div>
          </div>
          <div class="flex items-center justify-between px-2">
            <span class="text-[11px] font-bold text-slate-500 uppercase">{{ $t('users.list.detail.system.status') }}</span>
            <span class="px-3 py-1 bg-green-50 text-green-500 border border-green-100 rounded text-[8px] font-black uppercase tracking-widest">
              {{ $t('users.list.detail.system.verified') }}
            </span>
          </div>
        </div>
      </div>

      <!-- Reporting Structure -->
      <div class="p-8 bg-blue-50/50 dark:bg-blue-900/10 rounded-[2.5rem] space-y-6">
        <h4 class="text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('users.list.detail.reporting.title') }}</h4>
        <div class="flex items-center gap-5">
          <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center text-xl font-black text-[#1E3A5F] dark:text-white shadow-sm border border-slate-100 dark:border-slate-800">
            DP
          </div>
          <div class="space-y-1">
            <h5 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Darmawan Pratama</h5>
            <p class="text-[9px] font-bold text-slate-400 uppercase">{{ $t('users.list.detail.reporting.supervisor') }}</p>
          </div>
        </div>
        <button class="w-full py-4 bg-white dark:bg-slate-800 text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl shadow-sm border border-slate-100 dark:border-slate-800 hover:bg-slate-50 transition-all">
          {{ $t('users.list.detail.reporting.update') }}
        </button>
      </div>

      <!-- Recent Logs -->
      <div class="space-y-8">
        <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">{{ $t('users.list.detail.logs.title') }}</h4>
        <div class="space-y-10 pl-6 border-l-2 border-slate-50 dark:border-slate-800">
          <div v-for="(log, idx) in [
            { id: 'login', time: 'Hari ini, 09:42', loc: 'Jakarta Pusat' },
            { id: 'download', idVal: 'AR-99', time: 'Kemarin, 14:20', loc: 'Head Office' },
            { id: 'password', time: '02 Feb 2024, 10:11', loc: '' }
          ]" :key="idx" class="relative">
            <div class="absolute -left-[31px] top-0 w-4 h-4 rounded-full bg-white dark:bg-slate-900 border-4 border-[#1E3A5F] shadow-sm"></div>
            <div class="space-y-1">
              <h5 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`users.list.detail.logs.${log.id}`, { id: log.idVal }) }}</h5>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ log.time }} <span v-if="log.loc" class="ml-2">• {{ log.loc }}</span></p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideFilter, LucideUserPlus, LucideUsers, LucideCheckCircle2, 
  LucideShieldCheck, LucideUserX, LucideUser, LucideChevronLeft, 
  LucideChevronRight, LucideX, LucideRefreshCw, LucideNetwork, 
  LucideExternalLink 
} from 'lucide-vue-next'

const selectedUser = ref(null)

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
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
