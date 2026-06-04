<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- State: Checklist -->
    <div v-if="state === 'checklist'" class="space-y-10">
      <!-- Page Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
            {{ $t('warehouse.structure.title') }}
          </h1>
          <p class="text-xs font-bold text-slate-500">
            {{ $t('warehouse.structure.subtitle') }}
          </p>
        </div>
        <button class="px-6 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-sm hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucidePrinter class="w-4 h-4" />
          {{ $t('warehouse.structure.btn_print_batch') }}
        </button>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <div class="lg:col-span-2 space-y-8">
          <!-- Verification Checklist -->
          <div class="glass p-10 rounded-lg space-y-10">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-xl bg-primary-50 dark:bg-primary-900/20 text-primary-500 flex items-center justify-center shadow-sm">
                <LucideCheckSquare class="w-5 h-5" />
              </div>
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.structure.checklist.title') }}</h3>
            </div>

            <div class="space-y-6">
              <div v-for="(item, i) in [1, 2, 3]" :key="i" class="p-8 bg-slate-50/50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 rounded-lg flex items-start gap-6 hover:border-primary-500/30 transition-all group">
                <div class="pt-1">
                  <div class="w-6 h-6 rounded border-2 border-slate-200 dark:border-slate-700 flex items-center justify-center group-hover:border-primary-500 transition-colors cursor-pointer">
                    <LucideCheck v-if="checked[i]" class="w-4 h-4 text-primary-500" />
                  </div>
                </div>
                <div class="space-y-1">
                  <p class="text-sm font-black text-[#1E3A5F] dark:text-white">{{ $t(`warehouse.structure.checklist.item_${item}.title`) }}</p>
                  <p class="text-[11px] font-medium text-slate-500 leading-relaxed">{{ $t(`warehouse.structure.checklist.item_${item}.desc`) }}</p>
                </div>
              </div>
            </div>

            <button @click="state = 'scan'" class="w-full py-5 bg-[#1E3A5F] text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 active:scale-95">
              <LucideScanLine class="w-5 h-5" />
              {{ $t('warehouse.structure.checklist.btn_start') }}
            </button>
          </div>

          <!-- SOP Info -->
          <div class="p-8 bg-primary-50/30 dark:bg-primary-900/10 border border-primary-100/50 dark:border-primary-800/30 rounded-lg flex gap-6">
            <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-800 text-primary-500 flex items-center justify-center shadow-sm shrink-0">
              <LucideInfo class="w-6 h-6" />
            </div>
            <div class="space-y-2">
              <h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.structure.sop.title') }}</h4>
              <p class="text-[11px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed italic">
                {{ $t('warehouse.structure.sop.desc') }}
              </p>
            </div>
          </div>
        </div>

        <!-- Assigned Location Sidebar -->
        <div class="space-y-8">
          <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-right>
            <div class="flex items-center gap-3">
              <LucideMapPin class="w-5 h-5 text-primary-500" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.structure.assigned.title') }}</h3>
            </div>

            <div class="aspect-square bg-slate-100 dark:bg-slate-800 rounded-lg overflow-hidden relative group">
              <img src="https://api.mapbox.com/styles/v1/mapbox/dark-v10/static/106.8456, -6.2088,12,0,0/400x400?access_token=pk.eyJ1IjoiYmFyY2FiaWwiLCJhIjoiY2p3Z3R4Z3Q0MDByZDRicXl4bmZ6eXZwMiJ9.8_nF_E0wK_7w1p_8_v_8_w" class="w-full h-full object-cover grayscale opacity-60 group-hover:grayscale-0 group-hover:opacity-100 transition-all" />
              <div class="absolute inset-0 flex items-center justify-center">
                <div class="w-10 h-10 bg-primary-500/20 rounded-full animate-ping"></div>
                <div class="absolute w-5 h-5 bg-primary-500 rounded-full ring-4 ring-white dark:ring-slate-900"></div>
              </div>
              <div class="absolute bottom-6 left-6 bg-slate-800/80 backdrop-blur-md px-4 py-2 rounded-xl border border-white/10">
                <span class="text-[9px] font-black text-white uppercase tracking-widest">LIVE MAP • {{ $t('warehouse.structure.assigned.map_label') }}</span>
              </div>
            </div>

            <div class="space-y-6">
              <div v-for="(val, key) in { zone: 'Zone A', row: 'Row 12', level: 'Shelf 3', slot: 'R-12-03-A', batch: '#88291' }" :key="key" class="flex items-center justify-between py-2 border-b border-slate-50 dark:border-slate-800/50">
                <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t(`warehouse.structure.assigned.fields.${key}`) }}</span>
                <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ val }}</span>
              </div>
            </div>

            <button class="w-full py-4 bg-slate-50 dark:bg-slate-800 text-slate-500 dark:text-slate-400 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-100 transition-all">
              {{ $t('warehouse.structure.assigned.btn_nav') }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- State: Placement Verification Scan -->
    <div v-else-if="state === 'scan'" class="space-y-10" v-motion-fade>
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('warehouse.scan.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('warehouse.scan.subtitle') }}</p>
      </div>

      <!-- Stepper -->
      <div class="bg-white dark:bg-slate-900 p-8 rounded-lg border border-slate-100 dark:border-slate-800 shadow-sm flex items-center justify-center gap-12">
        <div v-for="s in [1, 2, 3]" :key="s" class="flex items-center gap-4">
          <div :class="`w-10 h-10 rounded-full flex items-center justify-center text-xs font-black transition-all ${step >= s ? 'bg-[#1E3A5F] text-white' : 'bg-slate-100 dark:bg-slate-800 text-slate-400'}`">
            {{ s }}
          </div>
          <span :class="`text-[10px] font-black uppercase tracking-widest ${step >= s ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400'}`">{{ $t(`warehouse.scan.stepper.step_${s}`) }}</span>
          <div v-if="s < 3" class="w-16 h-px bg-slate-100 dark:bg-slate-800 mx-2"></div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <div class="space-y-8">
          <!-- Target Rack Info -->
          <div class="glass p-8 rounded-lg space-y-6">
            <div class="flex items-center justify-between">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.rack_info.title') }}</h4>
              <span class="px-2 py-0.5 bg-green-50 text-green-500 rounded text-[8px] font-black uppercase tracking-widest border border-green-100">{{ $t('warehouse.scan.rack_info.scanned') }}</span>
            </div>
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.rack_info.code_label') }}</label>
              <div class="p-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <LucideQrCode class="w-5 h-5 text-[#1E3A5F] dark:text-white" />
                  <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">RCK-A1-204-B</span>
                </div>
                <LucideCheckCircle2 class="w-4 h-4 text-green-500" />
              </div>
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl space-y-1">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.rack_info.zone') }}</p>
                <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('warehouse.scan.rack_info.zone_val') }}</p>
              </div>
              <div class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl space-y-1">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.rack_info.category') }}</p>
                <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('warehouse.scan.rack_info.cat_val') }}</p>
              </div>
            </div>
          </div>

          <!-- Box Verification -->
          <div class="glass p-8 rounded-lg space-y-6 border-2 border-primary-500/10">
            <div class="flex items-center justify-between">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.box_verify.title') }}</h4>
              <span class="text-[9px] font-black text-primary-500 uppercase tracking-widest animate-pulse">{{ $t('warehouse.scan.box_verify.waiting') }}</span>
            </div>
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.box_verify.code_label') }}</label>
              <div class="relative">
                <LucideQrCode class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
                <input type="text" :placeholder="$t('warehouse.scan.box_verify.placeholder')" class="w-full pl-11 pr-5 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-xs font-bold outline-none shadow-sm" />
              </div>
            </div>
            <div class="p-6 bg-green-50/50 dark:bg-green-900/10 border border-green-100 dark:border-green-800 rounded-2xl space-y-2">
              <h5 class="text-[10px] font-black text-green-600 uppercase tracking-widest">{{ $t('warehouse.scan.box_verify.match_found') }}</h5>
              <p class="text-[9px] font-bold text-green-700/70 leading-relaxed">{{ $t('warehouse.scan.box_verify.match_desc') }}</p>
            </div>
          </div>
        </div>

        <div class="lg:col-span-2 space-y-8">
          <!-- Scanner Visual -->
          <div class="bg-slate-900 rounded-lg aspect-video relative overflow-hidden group shadow-2xl">
            <img src="https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?q=80&w=1000&auto=format&fit=crop" class="w-full h-full object-cover opacity-40 grayscale group-hover:grayscale-0 group-hover:opacity-60 transition-all duration-700" />
            
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="w-64 h-64 border-2 border-white/30 rounded-3xl relative">
                <div class="absolute -top-1 -left-1 w-8 h-8 border-t-4 border-l-4 border-primary-500 rounded-tl-xl"></div>
                <div class="absolute -top-1 -right-1 w-8 h-8 border-t-4 border-r-4 border-primary-500 rounded-tr-xl"></div>
                <div class="absolute -bottom-1 -left-1 w-8 h-8 border-b-4 border-l-4 border-primary-500 rounded-bl-xl"></div>
                <div class="absolute -bottom-1 -right-1 w-8 h-8 border-b-4 border-r-4 border-primary-500 rounded-br-xl"></div>
                
                <div class="absolute top-1/2 left-0 w-full h-0.5 bg-primary-500 shadow-[0_0_15px_rgba(59,130,246,1)] animate-scanner"></div>
              </div>
            </div>

            <div class="absolute inset-x-0 bottom-12 text-center space-y-2">
              <p class="text-xs font-black text-white uppercase tracking-widest drop-shadow-lg">{{ $t('warehouse.scan.scanner.hint') }}</p>
              <p class="text-[8px] font-black text-primary-400 uppercase tracking-widest opacity-60">{{ $t('warehouse.scan.scanner.auto') }}</p>
            </div>
          </div>

          <!-- Summary & Confirm -->
          <div class="glass p-10 rounded-lg flex items-center justify-between" v-motion-slide-visible-bottom>
            <div class="space-y-1">
              <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.scan.summary.title') }}</h4>
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.summary.ready') }}</p>
            </div>
            
            <div class="flex items-center gap-10">
              <div class="text-right">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('warehouse.scan.summary.confidence', { value: '' }).split(':')[0] }}</p>
                <p class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">99.8% Match</p>
              </div>
              <div class="flex gap-4">
                <button @click="state = 'checklist'" class="px-8 py-4 bg-slate-100 dark:bg-slate-800 text-slate-500 dark:text-slate-400 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-200 transition-all">
                  {{ $t('warehouse.scan.summary.btn_cancel') }}
                </button>
                <button @click="state = 'checklist'" class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all active:scale-95">
                  {{ $t('warehouse.scan.summary.btn_confirm') }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer Info Bar -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 pt-4">
        <div class="bg-white dark:bg-slate-900 p-6 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm flex items-center gap-5">
          <div class="w-10 h-10 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400"><LucideUser class="w-5 h-5" /></div>
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.footer.assignee') }}</p>
            <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">John Doe (Op #12)</p>
          </div>
        </div>
        <div class="bg-white dark:bg-slate-900 p-6 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm flex items-center gap-5">
          <div class="w-10 h-10 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400"><LucideHistory class="w-5 h-5" /></div>
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.footer.sync') }}</p>
            <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">Oct 24, 2023 - 14:32</p>
          </div>
        </div>
        <div class="bg-white dark:bg-slate-900 p-6 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm flex items-center gap-5">
          <div class="w-10 h-10 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400"><LucideMapPin class="w-5 h-5" /></div>
          <div>
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.scan.footer.location') }}</p>
            <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">Warehouse Central JKT</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucidePrinter, LucideCheckSquare, LucideCheck, LucideScanLine, 
  LucideInfo, LucideMapPin, LucideQrCode, LucideCheckCircle2,
  LucideUser, LucideHistory
} from 'lucide-vue-next'

const state = ref('checklist') // 'checklist', 'scan'
const step = ref(2)
const checked = ref([true, true, false])
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}

@keyframes scanner {
  0%, 100% { top: 0; }
  50% { top: 100%; }
}

.animate-scanner {
  animation: scanner 3s ease-in-out infinite;
}
</style>

