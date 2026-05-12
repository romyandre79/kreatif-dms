<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.monitoring.integration.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('admin.monitoring.integration.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
          <LucideDownload class="w-4 h-4" />
          {{ $t('admin.monitoring.integration.header.report') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3">
          <LucideRefreshCw class="w-4 h-4" />
          {{ $t('admin.monitoring.integration.header.refresh') }}
        </button>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div v-for="s in [
        { id: 'health', val: '94.2%', color: 'border-green-500', icon: LucideCheckCircle2, iconColor: 'text-green-500' },
        { id: 'latency', val: '12ms', color: 'border-blue-500', icon: LucideClock, iconColor: 'text-blue-500' },
        { id: 'alerts', val: '02', color: 'border-red-500', icon: LucideAlertTriangle, iconColor: 'text-red-500' },
        { id: 'sync', val: '1.2k', color: 'border-[#1E3A5F]', icon: LucideRotateCw, iconColor: 'text-[#1E3A5F]' }
      ]" :key="s.id" :class="`glass p-8 rounded-lg border-l-4 ${s.color} space-y-4 shadow-sm`">
        <div class="flex items-center justify-between">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`admin.monitoring.integration.stats.${s.id}`) }}</p>
          <component :is="s.icon" :class="`w-5 h-5 ${s.iconColor}`" />
        </div>
        <div class="space-y-1">
          <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ s.val }}</p>
          <p :class="`text-[8px] font-black uppercase tracking-widest ${s.iconColor}`">{{ $t(`admin.monitoring.integration.stats.${s.id}_sub`) }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Service Table -->
      <div class="lg:col-span-8 glass rounded-lg overflow-hidden">
        <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.monitoring.integration.table.title') }}</h3>
          <span class="flex items-center gap-2 text-[8px] font-black text-green-500 uppercase tracking-widest">
            <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
            {{ $t('admin.monitoring.integration.table.live') }}
          </span>
        </div>
        <table class="w-full text-left">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800">
              <th class="p-6 pl-10">{{ $t('admin.monitoring.integration.table.cols.name') }}</th>
              <th class="p-6">{{ $t('admin.monitoring.integration.table.cols.host') }}</th>
              <th class="p-6">{{ $t('admin.monitoring.integration.table.cols.status') }}</th>
              <th class="p-6 pr-10 text-right">{{ $t('admin.monitoring.integration.table.cols.latency') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="svc in [
              { id: 'ad', host: '10.0.1.45:389', status: 'ok', latency: '4ms', icon: LucideServer },
              { id: 'sql', host: 'db-prod-cluster.internal', status: 'spike', latency: '142ms', icon: LucideDatabase },
              { id: 's3', host: 's3.ap-southeast-1.aws', status: 'ok', latency: '22ms', icon: LucideFileBox },
              { id: 'scanner', host: '192.168.20.104', status: 'ok', latency: '12ms', icon: LucideScanLine },
              { id: 'printer', host: '192.168.20.115', status: 'fail', latency: 'OFF', icon: LucidePrinter },
              { id: 'rfid', host: 'localhost:8088', status: 'ok', latency: '1ms', icon: LucideRadioTower },
              { id: 'smtp', host: 'smtp.office365.com:587', status: 'ok', latency: '45ms', icon: LucideMail }
            ]" :key="svc.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
              <td class="p-6 pl-10">
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400 group-hover:text-[#1E3A5F] transition-all shadow-sm border border-slate-100 dark:border-slate-800">
                    <component :is="svc.icon" class="w-5 h-5" />
                  </div>
                  <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`admin.monitoring.integration.table.services.${svc.id}`) }}</p>
                </div>
              </td>
              <td class="p-6 text-[10px] font-bold text-slate-400 uppercase tracking-tight font-mono">{{ svc.host }}</td>
              <td class="p-6">
                <span :class="`px-3 py-1 rounded text-[8px] font-black uppercase tracking-widest border ${svc.status === 'ok' ? 'bg-green-50 text-green-500 border-green-100' : svc.status === 'spike' ? 'bg-amber-50 text-amber-500 border-amber-100' : 'bg-red-50 text-red-500 border-red-100'}`">
                  {{ $t(`admin.monitoring.integration.table.status.${svc.status}`) }}
                </span>
              </td>
              <td class="p-6 pr-10 text-right">
                <div class="flex items-center justify-end gap-3">
                  <span :class="`text-[10px] font-black ${svc.status === 'ok' ? 'text-green-500' : svc.status === 'spike' ? 'text-amber-500' : 'text-red-500'}`">{{ svc.latency }}</span>
                  <div :class="`w-4 h-1 rounded-full ${svc.status === 'ok' ? 'bg-green-500' : svc.status === 'spike' ? 'bg-amber-500' : 'bg-red-500'}`"></div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Right Column: Alert Feed & Trend -->
      <div class="lg:col-span-4 space-y-10">
        <div class="glass p-10 rounded-lg space-y-8">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div class="w-8 h-8 rounded-lg bg-red-50 text-red-500 flex items-center justify-center shadow-sm"><LucideActivity class="w-4 h-4" /></div>
              <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.monitoring.integration.feed.title') }}</h3>
            </div>
            <span class="text-[8px] font-black text-slate-300 uppercase">T+ 0.5s</span>
          </div>
          <div class="space-y-6">
            <div v-for="alert in [
              { id: 'offline', color: 'border-red-500', bg: 'bg-red-50/30', icon: LucideAlertCircle, iconColor: 'text-red-500', time: '14:23:45', code: 'ERR_CONN_TIMEOUT' },
              { id: 'threshold', color: 'border-amber-500', bg: 'bg-amber-50/30', icon: LucideHistory, iconColor: 'text-amber-500', time: '14:10:22', code: 'WRN_HIGH_LATENCY' },
              { id: 'heartbeat', color: 'border-blue-500', bg: 'bg-blue-50/30', icon: LucideCheckCircle2, iconColor: 'text-blue-500', time: '13:55:01', code: 'INF_HEARTBEAT_OK' }
            ]" :key="alert.id" :class="`p-6 rounded-2xl border-l-4 ${alert.color} ${alert.bg} space-y-4 transition-all hover:scale-[1.02] cursor-pointer`">
              <div class="flex items-center gap-4">
                <component :is="alert.icon" :class="`w-4 h-4 ${alert.iconColor}`" />
                <h4 :class="`text-[10px] font-black uppercase tracking-widest ${alert.iconColor}`">{{ $t(`admin.monitoring.integration.feed.${alert.id}.title`) }}</h4>
              </div>
              <p class="text-[11px] font-bold text-slate-500 leading-relaxed">{{ $t(`admin.monitoring.integration.feed.${alert.id}.desc`) }}</p>
              <div class="flex items-center justify-between text-[8px] font-black text-slate-400 uppercase tracking-widest">
                <span>{{ alert.time }}</span>
                <span>{{ alert.code }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="glass p-10 rounded-lg space-y-8">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.monitoring.integration.feed.trend') }}</h4>
          <div class="h-48 flex items-end gap-2 px-2">
            <div v-for="i in 15" :key="i" :class="`flex-grow rounded-sm transition-all hover:scale-y-110 cursor-pointer ${i === 8 ? 'bg-blue-400 h-[80%]' : i === 7 ? 'bg-[#1E3A5F] h-[60%]' : 'bg-slate-200 dark:bg-slate-800 h-[' + (Math.random() * 40 + 20) + '%]'}`"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDownload, LucideRefreshCw, LucideCheckCircle2, LucideClock, 
  LucideAlertTriangle, LucideRotateCw, LucideServer, LucideDatabase, 
  LucideFileBox, LucideScanLine, LucidePrinter, LucideRadioTower, 
  LucideMail, LucideActivity, LucideAlertCircle, LucideHistory 
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

