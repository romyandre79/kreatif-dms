<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('deployment.readiness.title') }}</h1>
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('deployment.readiness.subtitle', { id: 'F-55-ARCHIVE-NODE', time: '14:02' }) }}</p>
      </div>
      <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 group">
        <LucideZap class="w-4 h-4 fill-current group-hover:scale-110 transition-transform" />
        {{ $t('deployment.readiness.header.run') }}
      </button>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm border-t-8 border-blue-500">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('deployment.readiness.stats.overall') }}</p>
        <div class="flex items-center gap-4">
          <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">92%</p>
          <span class="text-[8px] font-black text-blue-500 uppercase tracking-widest">{{ $t('deployment.readiness.stats.optimal') }}</span>
        </div>
      </div>
      <div v-for="s in [
        { id: 'app', val: 'Stable', color: 'text-green-500', icon: LucideCheckCircle2 },
        { id: 'db', val: 'Healthy', color: 'text-green-500', icon: LucideDatabase }
      ]" :key="s.id" class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`deployment.readiness.stats.${s.id}`) }}</p>
        <div class="flex items-center gap-3">
          <component :is="s.icon" :class="`w-5 h-5 ${s.color}`" />
          <p :class="`text-xl font-black ${s.color} uppercase tracking-tight`">{{ s.val }}</p>
        </div>
      </div>
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('deployment.readiness.stats.storage') }}</p>
        <div class="space-y-2">
          <p class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight">14.2 TB <span class="text-[10px] text-slate-400">Free</span></p>
          <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
            <div class="h-full bg-[#1E3A5F]" style="width: 65%"></div>
          </div>
        </div>
      </div>
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('deployment.readiness.stats.network') }}</p>
        <div class="flex items-center gap-3">
          <LucideAlertTriangle class="w-5 h-5 text-amber-500" />
          <p class="text-xl font-black text-amber-500 uppercase tracking-tight">{{ $t('deployment.readiness.stats.network_val') }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Center Area: Topology Map -->
      <div class="lg:col-span-8 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 border border-slate-50 dark:border-slate-800">
          <div class="flex items-center gap-4 border-b border-slate-50 dark:border-slate-800 pb-8">
            <LucideNetwork class="w-6 h-6 text-[#1E3A5F]" />
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('deployment.readiness.topology.title') }}</h3>
          </div>
          
          <div class="grid grid-cols-2 md:grid-cols-3 gap-8">
            <div v-for="node in ['app', 'db', 'file', 'scanner', 'printer', 'ad']" :key="node" class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 flex flex-col items-center justify-center text-center space-y-4 hover:scale-105 transition-all cursor-pointer group">
              <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center text-[#1E3A5F] dark:text-white shadow-sm border border-slate-100 dark:border-slate-800 group-hover:bg-[#1E3A5F] group-hover:text-white transition-all">
                <component :is="node === 'app' ? LucideServer : node === 'db' ? LucideDatabase : node === 'file' ? LucideHardDrive : node === 'scanner' ? LucideScanLine : node === 'printer' ? LucidePrinter : LucideUsers" class="w-6 h-6" />
              </div>
              <div class="space-y-0.5">
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`deployment.readiness.topology.nodes.${node}`) }}</p>
                <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t(`deployment.readiness.topology.nodes.${node}_sub`) }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Server Inventory -->
        <div class="glass rounded-[3rem] overflow-hidden">
          <div class="p-10 flex items-center gap-4 border-b border-slate-50 dark:border-slate-800">
            <LucideBox class="w-6 h-6 text-[#1E3A5F]" />
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('deployment.readiness.inventory.title') }}</h3>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                <th class="p-8 pl-10">{{ $t('deployment.readiness.inventory.cols.host') }}</th>
                <th class="p-8">{{ $t('deployment.readiness.inventory.cols.os') }}</th>
                <th class="p-8">{{ $t('deployment.readiness.inventory.cols.cpu') }}</th>
                <th class="p-8">{{ $t('deployment.readiness.inventory.cols.disk') }}</th>
                <th class="p-8 pr-10 text-right">{{ $t('deployment.readiness.inventory.cols.role') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="s in [
                { host: 'AKR-APP-01', os: 'Ubuntu 22.04 LTS', cpu: '16 vCPU / 32 GB', disk: '45%', role: 'app', color: 'bg-blue-500' },
                { host: 'AKR-DB-01', os: 'RHEL 9.1', cpu: '32 vCPU / 64 GB', disk: '78%', role: 'db', color: 'bg-amber-500' },
                { host: 'AKR-STOR-01', os: 'TrueNAS Scale', cpu: '8 vCPU / 128 GB', disk: '82%', role: 'file', color: 'bg-[#1E3A5F]' }
              ]" :key="s.host" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
                <td class="p-8 pl-10 font-black text-xs text-blue-500 uppercase font-mono">{{ s.host }}</td>
                <td class="p-8 text-[11px] font-bold text-slate-500 uppercase">{{ s.os }}</td>
                <td class="p-8 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ s.cpu }}</td>
                <td class="p-8">
                  <div class="flex items-center gap-4">
                    <div class="w-24 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden shadow-inner">
                      <div :class="`h-full ${s.color}`" :style="{ width: s.disk }"></div>
                    </div>
                    <span class="text-[10px] font-black text-slate-400">{{ s.disk }}</span>
                  </div>
                </td>
                <td class="p-8 pr-10 text-right">
                  <span class="px-3 py-1 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 text-[8px] font-black text-slate-400 uppercase tracking-widest rounded">{{ $t(`deployment.readiness.inventory.roles.${s.role}`) }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Column: Checklist & Advisory -->
      <div class="lg:col-span-4 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('deployment.readiness.checklist.title') }}</h3>
          <div class="space-y-8">
            <div v-for="item in [
              { id: 'tls', status: 'done', icon: LucideCheckCircle2, color: 'text-green-500' },
              { id: 'firewall', status: 'done', icon: LucideCheckCircle2, color: 'text-green-500' },
              { id: 'backup', status: 'progress', icon: LucideRotateCw, color: 'text-amber-500', spin: true },
              { id: 'monitor', status: 'pending', icon: LucideCircle, color: 'text-slate-200' }
            ]" :key="item.id" class="flex items-start gap-6 group">
              <div :class="`w-8 h-8 rounded-lg flex items-center justify-center border-2 border-slate-100 dark:border-slate-800 bg-white dark:bg-slate-900 group-hover:border-[#1E3A5F] transition-all ${item.color}`">
                <component :is="item.icon" :class="`w-4 h-4 ${item.spin ? 'animate-spin' : ''}`" />
              </div>
              <div class="space-y-1">
                <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`deployment.readiness.checklist.${item.id}`) }}</p>
                <p class="text-[9px] font-bold text-slate-400 leading-relaxed">{{ $t(`deployment.readiness.checklist.${item.id}_desc`) }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Historical Readiness -->
        <div class="bg-[#1E3A5F] p-10 rounded-[3rem] space-y-8 text-white relative overflow-hidden group shadow-2xl shadow-blue-900/40">
          <LucideHistory class="absolute top-0 right-0 w-32 h-32 text-white/5 -rotate-12 translate-x-10 -translate-y-10 group-hover:scale-110 transition-transform" />
          <h4 class="text-[10px] font-black uppercase tracking-widest opacity-60">{{ $t('deployment.readiness.historical') }}</h4>
          <div class="h-32 flex items-end gap-3">
            <div v-for="i in 5" :key="i" :class="`flex-grow rounded-md transition-all hover:opacity-100 cursor-pointer ${i === 5 ? 'bg-white h-full shadow-lg' : 'bg-white/20 h-[' + (40 + i*10) + '%] opacity-40'}`"></div>
          </div>
          <div class="flex justify-between text-[8px] font-black uppercase tracking-widest opacity-40">
            <span>WK-01</span>
            <span>CURRENT</span>
          </div>
        </div>

        <!-- Technical Advisory -->
        <div class="p-10 bg-slate-100/50 dark:bg-slate-900/50 rounded-[3rem] border border-slate-200 dark:border-slate-800 space-y-6">
          <div class="flex items-center gap-4">
            <LucideInfo class="w-6 h-6 text-[#1E3A5F]" />
            <h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('deployment.readiness.advisory.title') }}</h4>
          </div>
          <p class="text-[10px] font-bold text-slate-500 leading-relaxed uppercase">{{ $t('deployment.readiness.advisory.desc') }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideZap, LucideCheckCircle2, LucideDatabase, LucideAlertTriangle, 
  LucideNetwork, LucideServer, LucideHardDrive, LucideScanLine, 
  LucidePrinter, LucideUsers, LucideBox, LucideRotateCw, LucideCircle, 
  LucideHistory, LucideInfo 
} from 'lucide-vue-next'

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

