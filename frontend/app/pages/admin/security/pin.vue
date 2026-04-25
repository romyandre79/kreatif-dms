<template>
  <div class="h-[calc(100vh-120px)] flex flex-col lg:flex-row gap-10 -m-10 overflow-hidden" v-motion-fade>
    <!-- Left Column: Summary & Policy -->
    <div class="w-full lg:w-96 bg-slate-50 dark:bg-slate-950 border-r border-slate-200 dark:border-slate-800 flex flex-col">
      <div class="flex-grow overflow-y-auto p-10 space-y-10 custom-scrollbar">
        <!-- Compliance Summary -->
        <div class="glass p-8 rounded-[2.5rem] space-y-8 shadow-sm">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.security.pin.summary.title') }}</h4>
          <div class="flex items-center justify-between gap-6">
            <div class="space-y-1">
              <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">94%</p>
              <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('admin.security.pin.summary.compliance') }}</p>
            </div>
            <div class="space-y-1 text-right">
              <p class="text-3xl font-black text-red-500 tracking-tighter">12</p>
              <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('admin.security.pin.summary.lockouts') }}</p>
            </div>
          </div>
          <button class="w-full p-4 bg-blue-50 dark:bg-blue-900/10 rounded-2xl flex items-center justify-between group transition-all hover:bg-blue-100/50">
            <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('admin.security.pin.summary.non_compliant', { count: 2 }) }}</span>
            <LucideArrowRight class="w-4 h-4 text-[#1E3A5F] group-hover:translate-x-1 transition-transform" />
          </button>
        </div>

        <!-- Enforcement Rules -->
        <div class="glass p-8 rounded-[3rem] space-y-10">
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><LucideShieldCheck class="w-5 h-5" /></div>
            <h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.security.pin.rules.title') }}</h4>
          </div>
          <div class="space-y-8">
            <div v-for="rule in [
              { id: 'mandatory', active: true },
              { id: 'security', active: true },
              { id: 'archive', active: false }
            ]" :key="rule.id" class="flex items-start justify-between gap-4">
              <div class="space-y-1">
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`admin.security.pin.rules.${rule.id}`) }}</p>
                <p class="text-[9px] font-bold text-slate-400 leading-relaxed">{{ $t(`admin.security.pin.rules.${rule.id}_desc`) }}</p>
              </div>
              <div :class="`w-10 h-5 rounded-full relative cursor-pointer transition-all ${rule.active ? 'bg-[#1E3A5F]' : 'bg-slate-200 dark:bg-slate-700'}`">
                <div :class="`absolute top-0.5 w-4 h-4 bg-white rounded-full transition-all ${rule.active ? 'left-5.5' : 'left-0.5'}`"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Config Grid -->
        <div class="grid grid-cols-2 gap-4">
          <div v-for="c in [
            { id: 'length', val: 6 },
            { id: 'attempts', val: 3 },
            { id: 'duration', val: 15 },
            { id: 'expiry', val: 90 }
          ]" :key="c.id" class="p-6 bg-white dark:bg-slate-900 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-2 shadow-sm">
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`admin.security.pin.config.${c.id}`) }}</p>
            <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t(`admin.security.pin.config.${c.id}_val`, { count: c.val }) }}</p>
          </div>
        </div>

        <div class="p-6 bg-blue-50/30 dark:bg-blue-900/10 rounded-[2.5rem] border border-blue-100 dark:border-blue-800 flex gap-4">
          <LucideInfo class="w-5 h-5 text-blue-500 shrink-0" />
          <p class="text-[9px] font-bold text-slate-500 leading-relaxed italic">{{ $t('admin.security.pin.config.reuse_info') }}</p>
        </div>

        <!-- Numpad Preview -->
        <div class="glass p-10 rounded-[3rem] space-y-8 flex flex-col items-center">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest w-full">{{ $t('admin.security.pin.preview.title') }}</h4>
          <div class="flex gap-2">
            <div v-for="i in 6" :key="i" :class="`w-10 h-14 rounded-lg border-2 flex items-center justify-center transition-all ${pin.length >= i ? 'border-[#1E3A5F] bg-[#1E3A5F]/5' : 'border-slate-100 dark:border-slate-800 bg-white dark:bg-slate-900'}`">
              <div v-if="pin.length >= i" class="w-2 h-2 bg-[#1E3A5F] rounded-full"></div>
            </div>
          </div>
          <div class="grid grid-cols-3 gap-3 w-full max-w-[240px]">
            <button v-for="n in [1, 2, 3, 4, 5, 6, 7, 8, 9]" :key="n" @click="press(n)" class="aspect-square bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-lg font-black text-[#1E3A5F] dark:text-white shadow-sm hover:scale-105 active:scale-95 transition-all">{{ n }}</button>
            <button @click="pin = ''" class="aspect-square flex items-center justify-center text-red-400 hover:text-red-600 transition-all"><LucideDelete class="w-6 h-6" /></button>
            <button @click="press(0)" class="aspect-square bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-lg font-black text-[#1E3A5F] dark:text-white shadow-sm hover:scale-105 active:scale-95 transition-all">0</button>
            <button class="aspect-square flex items-center justify-center text-green-500"><LucideCheckCircle2 class="w-8 h-8" /></button>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Column: User List -->
    <div class="flex-grow flex flex-col bg-white dark:bg-slate-900">
      <div class="p-10 space-y-8">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
          <div class="space-y-1">
            <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.security.pin.title') }}</h1>
            <p class="text-xs font-bold text-slate-400">{{ $t('admin.security.pin.subtitle') }}</p>
          </div>
          <div class="flex items-center gap-4">
            <div class="relative">
              <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              <input type="text" placeholder="Cari nama atau peran..." class="pl-12 pr-6 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none w-64 focus:ring-4 focus:ring-primary-500/10 transition-all" />
            </div>
            <button class="p-3 bg-slate-50 dark:bg-slate-800 rounded-xl text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideFilter class="w-5 h-5" /></button>
          </div>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[9px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                <th class="p-8">{{ $t('admin.security.pin.list.cols.name') }}</th>
                <th class="p-8">{{ $t('admin.security.pin.list.cols.role') }}</th>
                <th class="p-8">{{ $t('admin.security.pin.list.cols.status') }}</th>
                <th class="p-8">{{ $t('admin.security.pin.list.cols.update') }}</th>
                <th class="p-8">{{ $t('admin.security.pin.list.cols.failed') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="user in users" :key="user.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
                <td class="p-8">
                  <div class="flex items-center gap-5">
                    <div v-if="user.avatar" class="w-12 h-12 rounded-2xl overflow-hidden border-2 border-white dark:border-slate-800 shadow-sm"><img :src="user.avatar" class="w-full h-full object-cover" /></div>
                    <div v-else class="w-12 h-12 rounded-2xl bg-slate-100 dark:bg-slate-800 text-slate-400 flex items-center justify-center text-xs font-black">{{ user.initials }}</div>
                    <div class="space-y-0.5">
                      <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ user.name }}</p>
                      <p class="text-[9px] font-bold text-slate-400">{{ user.email }}</p>
                    </div>
                  </div>
                </td>
                <td class="p-8 text-xs font-bold text-slate-500 uppercase">{{ user.role }}</td>
                <td class="p-8">
                  <span :class="`px-4 py-1 rounded-full text-[9px] font-black uppercase tracking-widest border ${user.statusColor}`">{{ $t(`admin.security.pin.list.status.${user.status}`) }}</span>
                </td>
                <td class="p-8 text-[10px] font-bold text-slate-500 uppercase">{{ user.lastUpdate }}</td>
                <td class="p-8 text-center font-black" :class="user.failed > 0 ? 'text-red-500' : 'text-slate-300'">{{ user.failed }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideShieldCheck, LucideArrowRight, LucideInfo, LucideSearch, 
  LucideFilter, LucideDelete, LucideCheckCircle2 
} from 'lucide-vue-next'

const pin = ref('')
const press = (n) => {
  if (pin.value.length < 6) pin.value += n
}

const users = [
  { id: 1, name: 'Bambang Susilo', email: 'bambang.s@akiradata.co.id', role: 'Finance Approval', initials: 'BS', status: 'active', statusColor: 'bg-green-50 text-green-500 border-green-100', lastUpdate: '12 Des 2023', failed: 0 },
  { id: 2, name: 'Dewi Lestari', email: 'dewi.l@akiradata.co.id', role: 'System Admin', initials: 'DL', status: 'locked', statusColor: 'bg-red-50 text-red-500 border-red-100', lastUpdate: '05 Jan 2024', failed: 3 },
  { id: 3, name: 'Agus Pratama', email: 'agus.p@akiradata.co.id', role: 'Document Officer', initials: 'AP', status: 'expired', statusColor: 'bg-amber-50 text-amber-500 border-amber-100', lastUpdate: '10 Agu 2023', failed: 0 },
  { id: 4, name: 'Siti Nurhaliza', email: 'siti.n@akiradata.co.id', role: 'Staff Operasional', initials: 'SN', status: 'not_set', statusColor: 'bg-slate-50 text-slate-400 border-slate-100', lastUpdate: 'Never', failed: 0 }
]

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>
