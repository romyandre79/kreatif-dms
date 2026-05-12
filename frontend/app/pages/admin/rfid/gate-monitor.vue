<template>
  <div class="flex h-screen bg-[#F1F5F9] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <!-- Left Sidebar: Live Event Stream -->
    <aside class="w-80 bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 flex flex-col p-6 overflow-hidden">
      <div class="flex items-center justify-between mb-8">
        <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Live Event Stream</h3>
        <span class="px-2 py-0.5 bg-blue-50 dark:bg-blue-900/30 text-[8px] font-black text-blue-500 uppercase tracking-widest rounded-md animate-pulse">Real-time</span>
      </div>

      <div class="space-y-4 overflow-y-auto custom-scrollbar pr-2 flex-grow">
        <div v-for="event in events" :key="event.id" 
          class="p-5 rounded-2xl border border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 space-y-3 relative group hover:bg-white dark:hover:bg-slate-800 transition-all border-l-4"
          :class="event.border"
        >
          <div class="flex items-center justify-between">
            <p class="text-[9px] font-black text-[#1E3A5F] dark:text-blue-400 font-mono tracking-tighter">{{ event.tag }}</p>
            <p class="text-[8px] font-bold text-slate-400 uppercase">{{ event.time }}</p>
          </div>
          <div class="space-y-1">
            <p v-if="event.user" class="text-[10px] font-bold text-slate-500 dark:text-slate-300 uppercase tracking-tight">User: <span class="font-black text-[#1E3A5F] dark:text-white">{{ event.user }}</span></p>
            <p v-if="event.doc" class="text-[10px] font-bold text-slate-500 dark:text-slate-300 uppercase tracking-tight">Doc: <span class="font-black text-[#1E3A5F] dark:text-white">{{ event.doc }}</span></p>
            <p v-if="event.event" class="text-[10px] font-bold text-slate-500 dark:text-slate-300 uppercase tracking-tight">Event: <span class="font-black text-[#1E3A5F] dark:text-white">{{ event.event }}</span></p>
          </div>
          <div class="flex items-center gap-2">
            <div :class="`w-1.5 h-1.5 rounded-full ${event.dot}`"></div>
            <p :class="`text-[8px] font-black uppercase tracking-widest ${event.statusColor}`">{{ event.status }}</p>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content Area: Gate Security Monitor -->
    <main class="flex-grow flex flex-col overflow-hidden p-10 gap-10">
      <div class="flex flex-col gap-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">Gate Security Monitor</h1>
        <div class="flex items-center gap-2 text-slate-500">
          <LucideMapPin class="w-4 h-4" />
          <p class="text-xs font-bold uppercase tracking-widest">Zone 4: Southeast Exit Point <span class="text-slate-400">(F-35B)</span></p>
        </div>
      </div>

      <!-- Incident Alert Module -->
      <div class="bg-white dark:bg-slate-900 rounded-lg shadow-2xl shadow-red-900/10 border border-red-500/20 overflow-hidden relative group max-w-4xl" v-motion-pop>
        <div class="bg-red-500 px-8 py-5 flex items-center justify-between text-white">
          <div class="flex items-center gap-3">
            <LucideShieldAlert class="w-6 h-6 animate-pulse" />
            <span class="text-[11px] font-black uppercase tracking-[0.2em]">Incident Detected</span>
          </div>
          <span class="px-3 py-1 bg-white/20 rounded-lg text-[9px] font-black uppercase tracking-widest">Alarm Active</span>
        </div>

        <div class="p-12 flex gap-12 items-start">
          <div class="w-48 aspect-[3/4] bg-slate-100 dark:bg-slate-800 rounded-3xl overflow-hidden shadow-2xl relative">
            <img src="https://images.unsplash.com/photo-1568667256549-094345857637?auto=format&fit=crop&q=80&w=400" class="w-full h-full object-cover grayscale opacity-50" />
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="bg-red-500/80 px-4 py-2 text-white text-[10px] font-black uppercase tracking-widest -rotate-12 border-2 border-white">Top Secret</div>
            </div>
            <p class="absolute bottom-4 left-0 w-full text-center text-[9px] font-black text-white uppercase tracking-widest bg-black/50 py-2">ID: DOC-RFID-44910-X</p>
          </div>

          <div class="flex-grow space-y-10">
            <div class="space-y-2">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Status</p>
              <h2 class="text-6xl font-black text-red-500 tracking-tighter uppercase">Invalid Exit</h2>
            </div>

            <div class="grid grid-cols-2 gap-x-12 gap-y-8">
              <div class="space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Subject Detected</p>
                <p class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase leading-tight">Anton Wijaya<br/><span class="text-sm text-slate-500">(EMP_412)</span></p>
              </div>
              <div class="space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Document Class</p>
                <p class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase">Level 4:<br/>Restricted Internal</p>
              </div>
              <div class="space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Detection Time</p>
                <p class="text-lg font-black text-[#1E3A5F] dark:text-white font-mono tracking-tight">14:21:30.442</p>
              </div>
              <div class="space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Gate ID</p>
                <p class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase">F-35b -<br/>Perimeter South</p>
              </div>
            </div>

            <div class="pt-8 border-t border-slate-100 dark:border-slate-800 flex items-center gap-4">
              <div class="flex -space-x-3">
                <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=100" class="w-10 h-10 rounded-full border-2 border-white dark:border-slate-900 object-cover" />
                <img src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?auto=format&fit=crop&q=80&w=100" class="w-10 h-10 rounded-full border-2 border-white dark:border-slate-900 object-cover" />
              </div>
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest italic">2 Security Personnel responding...</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Bottom Grid: Heatmap & Health -->
      <div class="grid grid-cols-2 gap-10 max-w-4xl">
        <div class="glass p-10 rounded-lg space-y-8">
          <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Frequency Heatmap</h3>
          <div class="h-40 flex items-end justify-between gap-2 px-4">
            <div v-for="(val, i) in [30, 45, 60, 55, 90, 40, 35]" :key="i" 
              class="flex-grow rounded-lg transition-all"
              :class="i === 4 ? 'bg-red-500 h-[90%]' : 'bg-blue-100 dark:bg-blue-900/40'"
              :style="{ height: val + '%' }"
            ></div>
          </div>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest italic text-center">Spike detected at current timestamp.</p>
        </div>

        <div class="glass p-10 rounded-lg space-y-8">
          <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Gate Hardware Health</h3>
          <div class="space-y-8">
            <div class="space-y-3">
              <div class="flex justify-between items-center text-[10px] font-black uppercase tracking-widest">
                <span class="text-slate-400">Antenna Power</span>
                <span class="text-green-500">100%</span>
              </div>
              <div class="h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-green-500 w-full"></div>
              </div>
            </div>
            <div class="space-y-3">
              <div class="flex justify-between items-center text-[10px] font-black uppercase tracking-widest">
                <span class="text-slate-400">Database Latency</span>
                <span class="text-green-500">12ms</span>
              </div>
              <div class="h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-green-500 w-1/4"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Right Sidebar: Command & Control -->
    <aside class="w-80 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-8 gap-10">
      <div class="space-y-8">
        <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Command & Control</h3>
        
        <div class="space-y-4">
          <button class="w-full py-6 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-lg shadow-xl shadow-blue-900/20 transition-all flex flex-col items-center gap-3 active:scale-95">
            <LucideCheckCircle class="w-6 h-6" />
            <span class="text-[10px] font-black uppercase tracking-[0.2em]">Acknowledge Alarm</span>
          </button>
          <button class="w-full py-6 bg-red-500 hover:bg-red-600 text-white rounded-lg shadow-xl shadow-red-900/20 transition-all flex flex-col items-center gap-3 active:scale-95 animate-pulse">
            <LucideLock class="w-6 h-6" />
            <span class="text-[10px] font-black uppercase tracking-[0.2em]">Lock User Access</span>
          </button>
          <button class="w-full py-6 bg-white dark:bg-slate-800 border-2 border-slate-200 dark:border-slate-700 text-[#1E3A5F] dark:text-white rounded-lg transition-all flex flex-col items-center gap-3 hover:bg-slate-50">
            <LucideBell class="w-6 h-6" />
            <span class="text-[10px] font-black uppercase tracking-[0.2em]">Notify Security</span>
          </button>
        </div>
      </div>

      <div class="space-y-6">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">System Overrides</p>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-5 bg-slate-50 dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700">
            <span class="text-[10px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest">Emergency Open</span>
            <LucideWaves class="w-5 h-5 text-slate-400" />
          </div>
          <div class="flex items-center justify-between p-5 bg-slate-50 dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700">
            <span class="text-[10px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest">Hard Lockdown</span>
            <LucideLock class="w-5 h-5 text-slate-400" />
          </div>
        </div>
      </div>

      <div class="mt-auto pt-8 border-t border-slate-50 dark:border-slate-800 flex items-center gap-4">
        <div class="w-2 h-2 rounded-full bg-blue-500 animate-pulse"></div>
        <div class="space-y-0.5">
          <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Console Syncing...</p>
          <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">Terminal ID: DMS-SEC-NODE-01</p>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { 
  LucideMapPin, LucideShieldAlert, LucideCheckCircle, LucideLock, 
  LucideBell, LucideWaves 
} from 'lucide-vue-next'

const events = [
  { id: 1, tag: 'RFID:TAG_99821', user: 'Budi Santoso', time: '14:22:01', status: 'Valid Entry - Main Lobby', statusColor: 'text-green-500', dot: 'bg-green-500', border: 'border-green-500' },
  { id: 2, tag: 'RFID:TAG_77123', doc: 'Confidential_Report_Q3.pdf', time: '14:21:45', status: 'Authorized Removal', statusColor: 'text-green-500', dot: 'bg-green-500', border: 'border-green-500' },
  { id: 3, tag: 'RFID:TAG_UNKNOWN', event: 'Unauthorized Proximity', time: '14:21:30', status: 'Invalid Exit Detected', statusColor: 'text-red-500', dot: 'bg-red-500', border: 'border-red-500' },
  { id: 4, tag: 'RFID:TAG_11204', user: 'Siti Aminah', time: '14:20:12', status: 'Valid Entry - Server Room', statusColor: 'text-green-500', dot: 'bg-green-500', border: 'border-green-500' },
  { id: 5, tag: 'RFID:TAG_88321', user: 'Ahmad Dhani', time: '14:18:55', status: 'Valid Exit - Main Gate', statusColor: 'text-green-500', dot: 'bg-green-500', border: 'border-green-500' }
]

definePageMeta({
  layout: 'none'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-sm;
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

