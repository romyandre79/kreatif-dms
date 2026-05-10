<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- State: Reconciliation Report -->
    <template v-if="state === 'report'">
      <!-- Page Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <div class="flex items-center gap-3 text-[9px] font-black text-slate-400 uppercase tracking-widest mb-1">
            <span class="px-2 py-0.5 bg-blue-50 text-blue-500 rounded border border-blue-100">{{ $t('stock.reconciliation.badge') }}</span>
            <span>Report ID: #REC-2024-001</span>
          </div>
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
            {{ $t('stock.reconciliation.title') }}
          </h1>
          <p class="text-xs font-bold text-slate-500">
            {{ $t('stock.reconciliation.subtitle', { date: 'Oct 24, 2023 • 09:42 AM' }) }}
          </p>
        </div>
        <div class="flex items-center gap-4">
          <button class="px-8 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-sm hover:bg-slate-50 transition-all flex items-center gap-2">
            <LucideDownload class="w-4 h-4" />
            {{ $t('stock.reconciliation.btn_export') }}
          </button>
          <button @click="state = 'resolution'" class="px-10 py-3 bg-primary-500 text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-primary-500/20 hover:bg-primary-600 transition-all flex items-center gap-2 active:scale-95">
            <LucideSearch class="w-4 h-4" />
            {{ $t('stock.reconciliation.btn_investigate') }}
          </button>
        </div>
      </div>

      <!-- Summary Cards Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        <div v-for="card in summaryCards" :key="card.type" :class="`glass p-8 rounded-lg border-2 transition-all group ${card.borderColor}`">
          <div class="flex items-center justify-between mb-6">
            <h4 :class="`text-[10px] font-black uppercase tracking-widest ${card.textColor}`">{{ $t(`stock.reconciliation.summary.${card.type}.title`) }}</h4>
            <div :class="`w-8 h-8 rounded-lg flex items-center justify-center ${card.iconBg} ${card.textColor}`">
              <component :is="card.icon" class="w-4 h-4" />
            </div>
          </div>
          <div class="space-y-1">
            <div class="flex items-baseline gap-3">
              <span class="text-4xl font-black text-[#1E3A5F] dark:text-white">{{ card.val }}</span>
              <span :class="`text-[9px] font-black uppercase tracking-widest ${card.subTextColor}`">{{ $t(`stock.reconciliation.summary.${card.type}.${card.subKey}`) }}</span>
            </div>
            <p class="text-[10px] font-bold text-slate-400 leading-tight">{{ $t(`stock.reconciliation.summary.${card.type}.desc`) }}</p>
          </div>
        </div>
      </div>

      <!-- Inventory Table Section -->
      <div class="glass rounded-lg overflow-hidden" v-motion-slide-visible-bottom>
        <div class="p-8 border-b border-slate-100 dark:border-slate-800 flex flex-col md:flex-row md:items-center justify-between gap-6 bg-white/30">
          <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('stock.reconciliation.table.title') }}</h3>
          <div class="flex items-center bg-slate-100 dark:bg-slate-800 p-1 rounded-xl">
            <button v-for="t in ['all', 'location']" :key="t" :class="`px-6 py-2 rounded-lg text-[10px] font-black uppercase tracking-widest transition-all ${tableTab === t ? 'bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white shadow-sm' : 'text-slate-400 hover:text-slate-600'}`" @click="tableTab = t">
              {{ $t(`stock.reconciliation.table.tabs.${t}`) }}
            </button>
          </div>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800 bg-slate-50/30 dark:bg-slate-900/30">
                <th class="p-8 px-10">{{ $t('stock.reconciliation.table.cols.sku') }}</th>
                <th class="p-8 text-center">{{ $t('stock.reconciliation.table.cols.sys_qty') }}</th>
                <th class="p-8 text-center">{{ $t('stock.reconciliation.table.cols.phy_qty') }}</th>
                <th class="p-8 text-center">{{ $t('stock.reconciliation.table.cols.variance') }}</th>
                <th class="p-8 text-center">{{ $t('stock.reconciliation.table.cols.status') }}</th>
                <th class="p-8 px-10">{{ $t('stock.reconciliation.table.cols.location') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
              <tr v-for="item in inventoryItems" :key="item.sku" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all group">
                <td class="p-8 px-10">
                  <div class="space-y-1">
                    <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ item.name }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ item.sku }}</p>
                  </div>
                </td>
                <td class="p-8 text-center text-sm font-black text-slate-700 dark:text-slate-300">{{ item.sys }}</td>
                <td class="p-8 text-center text-sm font-black text-slate-700 dark:text-slate-300">{{ item.phy }}</td>
                <td class="p-8 text-center">
                  <span :class="`text-sm font-black ${item.variance > 0 ? 'text-red-500' : item.variance < 0 ? 'text-orange-500' : 'text-green-500'}`">
                    {{ item.variance > 0 ? '+' + item.variance : item.variance }}
                  </span>
                </td>
                <td class="p-8 text-center">
                  <span :class="`px-3 py-1 rounded-md text-[8px] font-black tracking-widest border ${getStatusStyles(item.status)}`">
                    {{ item.status }}
                  </span>
                </td>
                <td class="p-8 px-10">
                  <div class="space-y-1">
                    <p class="text-[11px] font-bold text-slate-600 dark:text-slate-400">{{ item.location }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase">{{ item.subloc }}</p>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="p-8 text-center border-t border-slate-50 dark:border-slate-800">
          <button class="text-[10px] font-black text-primary-500 uppercase tracking-widest hover:underline transition-all">
            {{ $t('stock.reconciliation.table.load_more', { count: 138 }) }}
          </button>
        </div>
      </div>
    </template>

    <!-- State: Discrepancy Resolution -->
    <template v-else-if="state === 'resolution'">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('stock.resolution.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('stock.resolution.subtitle') }}</p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <div class="lg:col-span-2 space-y-8">
          <div class="glass p-10 rounded-lg space-y-10">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 text-[#1E3A5F] dark:text-white flex items-center justify-center shadow-sm">
                  <LucideClipboardList class="w-5 h-5" />
                </div>
                <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('stock.resolution.items.title') }}</h3>
              </div>
              <span class="px-4 py-1.5 bg-blue-50 text-blue-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-blue-100">
                {{ $t('stock.resolution.items.pending', { count: 2 }) }}
              </span>
            </div>

            <div class="space-y-10">
              <div class="space-y-6">
                <div class="flex items-start gap-6">
                  <div class="w-12 h-12 rounded-2xl bg-orange-50 text-orange-500 flex items-center justify-center shrink-0 shadow-sm border border-orange-100"><LucidePlusSquare class="w-6 h-6" /></div>
                  <div class="flex-grow space-y-3">
                    <div class="flex items-center justify-between">
                      <div class="flex items-center gap-3">
                        <span class="text-[9px] font-black text-orange-500 uppercase tracking-widest bg-orange-50 px-2 py-0.5 rounded">{{ $t('stock.resolution.items.extra.badge') }}</span>
                        <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Ref: #LGL-055</span>
                      </div>
                      <div class="text-right">
                        <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.resolution.items.extra.suggested') }}</p>
                        <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">"{{ $t('stock.resolution.items.extra.action_text', { loc: 'Rak A1' }) }}"</p>
                      </div>
                    </div>
                    <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">EXTRA Item LGL.AMS.B3.055</h4>
                    <p class="text-[11px] font-bold text-slate-500 leading-relaxed max-w-xl">Item found in staging area B3, but system records indicate it belongs to main storage.</p>
                  </div>
                </div>
                <button class="w-full py-4 bg-primary-500 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-primary-500/20 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95">
                  <LucideRefreshCcw class="w-4 h-4" />
                  {{ $t('stock.resolution.items.extra.btn_update', { loc: 'Rak A1' }) }}
                </button>
              </div>
              <div class="h-px bg-slate-50 dark:bg-slate-800"></div>
              <div class="space-y-6">
                <div class="flex items-start gap-6">
                  <div class="w-12 h-12 rounded-2xl bg-red-50 text-red-500 flex items-center justify-center shrink-0 shadow-sm border border-red-100"><LucideClipboardX class="w-6 h-6" /></div>
                  <div class="flex-grow space-y-3">
                    <div class="flex items-center gap-3">
                      <span class="text-[9px] font-black text-red-500 uppercase tracking-widest bg-red-50 px-2 py-0.5 rounded">{{ $t('stock.resolution.items.missing.badge') }}</span>
                      <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Ref: #LGL-012</span>
                    </div>
                    <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">MISSING Item LGL.AMS.A1.012</h4>
                    <p class="text-[11px] font-bold text-slate-500 leading-relaxed max-w-xl">Item not found in its designated location A1. Checked nearby secondary bins with no result.</p>
                  </div>
                </div>
                <div class="p-6 bg-slate-50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800 flex items-center gap-4 group cursor-pointer" @click="item2Checked = !item2Checked">
                  <div :class="`w-6 h-6 rounded border-2 transition-all flex items-center justify-center ${item2Checked ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'border-slate-200 bg-white'}`">
                    <LucideCheck v-if="item2Checked" class="w-4 h-4 text-white" />
                  </div>
                  <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('stock.resolution.items.missing.checkbox') }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="space-y-8">
          <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-right>
            <div class="flex items-center gap-3">
              <LucidePenTool class="w-5 h-5 text-primary-500" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('stock.resolution.details.title') }}</h3>
            </div>
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.resolution.details.note_label') }}</label>
              <textarea :placeholder="$t('stock.resolution.details.note_placeholder')" class="w-full h-48 px-6 py-5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-lg text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all resize-none custom-scrollbar"></textarea>
              <div class="flex gap-3 text-slate-400">
                <LucideInfo class="w-4 h-4 shrink-0" />
                <p class="text-[10px] font-bold leading-relaxed italic">{{ $t('stock.resolution.details.note_hint') }}</p>
              </div>
            </div>
            <div class="space-y-4 pt-6">
              <button @click="state = 'validation'" class="w-full py-5 bg-primary-500 text-white rounded-lg text-[10px] font-black uppercase tracking-widest shadow-xl shadow-primary-500/30 hover:bg-primary-600 transition-all active:scale-95">
                {{ $t('stock.resolution.details.btn_submit') }}
              </button>
              <button class="w-full text-[10px] font-black text-slate-400 uppercase tracking-widest hover:text-[#1E3A5F] transition-all">
                {{ $t('stock.resolution.details.btn_save') }}
              </button>
            </div>
            <div class="pt-10 border-t border-slate-50 dark:border-slate-800">
              <div class="bg-blue-50/30 dark:bg-blue-900/10 p-6 rounded-lg flex items-center gap-5">
                <div class="w-10 h-10 rounded-full bg-primary-500 text-white flex items-center justify-center shadow-lg"><LucideUserCheck class="w-5 h-5" /></div>
                <div>
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.resolution.details.actor') }}</p>
                  <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">Admin Warehouse</p>
                  <p class="text-[8px] font-bold text-primary-500 uppercase">{{ $t('stock.resolution.details.authority') }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!-- State: Final Audit Validation (Head Approval) -->
    <template v-else-if="state === 'validation'">
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('stock.validation.title') }}</h1>
          <p class="text-xs font-bold text-slate-500">{{ $t('stock.validation.subtitle') }}</p>
        </div>
        <span class="px-6 py-2 bg-amber-50 text-amber-500 rounded-full text-[10px] font-black uppercase tracking-widest border border-amber-100 flex items-center gap-3">
          <span class="w-2 h-2 bg-amber-500 rounded-full animate-pulse"></span>
          {{ $t('stock.validation.pending_badge') }}
        </span>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <div class="lg:col-span-2 space-y-8">
          <!-- Unresolved Missing Items -->
          <div class="glass p-10 rounded-lg space-y-8">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-4 bg-amber-50 dark:bg-amber-900/10 px-6 py-3 rounded-2xl border border-amber-100 dark:border-amber-800">
                <LucideAlertTriangle class="w-5 h-5 text-amber-500" />
                <h3 class="text-xs font-black text-amber-800 dark:text-amber-400 uppercase tracking-widest">{{ $t('stock.validation.unresolved.title', { count: 2 }) }}</h3>
              </div>
              <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.validation.unresolved.id', { id: '#ST-2023-0892' }) }}</span>
            </div>

            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                    <th class="py-6 px-4">{{ $t('stock.validation.unresolved.cols.code') }}</th>
                    <th class="py-6 px-4">{{ $t('stock.validation.unresolved.cols.desc') }}</th>
                    <th class="py-6 px-4 text-center">{{ $t('stock.validation.unresolved.cols.sys_qty') }}</th>
                    <th class="py-6 px-4 text-center">{{ $t('stock.validation.unresolved.cols.act_qty') }}</th>
                    <th class="py-6 px-4 text-center">{{ $t('stock.validation.unresolved.cols.variance') }}</th>
                    <th class="py-6 px-4">{{ $t('stock.validation.unresolved.cols.status') }}</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                  <tr v-for="item in missingItems" :key="item.code" class="group">
                    <td class="py-8 px-4">
                      <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ item.code }}</p>
                    </td>
                    <td class="py-8 px-4">
                      <div class="space-y-1">
                        <p class="text-sm font-bold text-slate-600 dark:text-slate-300">{{ item.desc }}</p>
                        <p class="text-[9px] font-bold text-slate-400 uppercase">{{ item.type }}</p>
                      </div>
                    </td>
                    <td class="py-8 px-4 text-center font-black text-[#1E3A5F] dark:text-white">1</td>
                    <td class="py-8 px-4 text-center font-black text-red-500">0</td>
                    <td class="py-8 px-4 text-center font-black text-red-500">-1</td>
                    <td class="py-8 px-4">
                      <div class="relative group">
                        <select class="w-full appearance-none px-6 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-[10px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest outline-none focus:ring-2 focus:ring-primary-500/20">
                          <option>{{ $t('stock.validation.unresolved.lost_status') }}</option>
                        </select>
                        <LucideChevronDown class="absolute right-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- Digital Signature / PIN -->
          <div class="glass p-10 rounded-lg space-y-10">
            <div class="flex items-center gap-4">
              <LucidePenTool class="w-6 h-6 text-primary-500" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('stock.signature.title') }}</h3>
            </div>
            <div class="space-y-8">
              <div class="relative">
                <LucideLock class="absolute left-6 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-300" />
                <input type="password" :placeholder="$t('stock.signature.pin_placeholder')" class="w-full pl-16 pr-8 py-5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-sm font-black outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" />
              </div>
              <div class="aspect-[3/1] bg-slate-50 dark:bg-slate-800/50 rounded-2xl border-2 border-dashed border-slate-200 dark:border-slate-700 flex items-center justify-center relative group overflow-hidden">
                <div class="text-center space-y-2 opacity-30 group-hover:opacity-60 transition-opacity">
                  <LucideSignature class="w-10 h-10 mx-auto text-slate-400" />
                  <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.signature.pad_hint') }}</p>
                </div>
                <div class="absolute inset-x-0 bottom-8 h-px bg-slate-200 dark:bg-slate-700 w-2/3 mx-auto"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar: Finalization Actions -->
        <div class="space-y-8">
          <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-right>
            <div class="p-8 bg-blue-50/50 dark:bg-blue-900/10 rounded-lg border border-blue-100 dark:border-blue-800 space-y-3">
              <div class="flex items-center gap-3 text-primary-500">
                <LucideInfo class="w-5 h-5" />
                <h4 class="text-[10px] font-black uppercase tracking-widest">{{ $t('stock.validation.actions.note_title') }}</h4>
              </div>
              <p class="text-[10px] font-bold text-slate-500 leading-relaxed italic">{{ $t('stock.validation.actions.note_desc') }}</p>
            </div>

            <div class="space-y-4 pt-6">
              <button @click="state = 'report'" class="w-full py-5 bg-primary-500 text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-2xl shadow-primary-500/40 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95">
                <LucideCheckCircle2 class="w-5 h-5" />
                {{ $t('stock.validation.actions.btn_approve') }}
              </button>
              <button class="w-full py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-400 rounded-lg text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
                {{ $t('stock.validation.actions.btn_reject') }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!-- Global Footer -->
    <footer class="pt-20 flex flex-col md:flex-row md:items-center justify-between gap-6 opacity-40 grayscale hover:opacity-100 hover:grayscale-0 transition-all duration-700">
      <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">© 2024 StockTake System. All items logged for audit.</p>
      <div class="flex items-center gap-10">
        <div class="flex items-center gap-3">
          <span class="w-2 h-2 bg-green-500 rounded-full"></span>
          <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">SYSTEM CONNECTED</span>
        </div>
        <button class="text-[10px] font-black text-slate-400 uppercase tracking-widest hover:underline">Help & Support</button>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideDownload, LucideSearch, LucideCheckCircle2, LucideHistory, 
  LucideAlertTriangle, LucideShieldCheck, LucideClipboardList,
  LucidePlusSquare, LucideClipboardX, LucideCheck, LucideRefreshCcw,
  LucidePenTool, LucideInfo, LucideUserCheck, LucideChevronDown,
  LucideLock, LucideSignature
} from 'lucide-vue-next'

const state = ref('report') // 'report', 'resolution', 'validation'
const tableTab = ref('all')
const item2Checked = ref(false)

const summaryCards = [
  { type: 'match', val: 142, subKey: 'valid', icon: LucideCheckCircle2, borderColor: 'border-green-100 hover:border-green-500/30', textColor: 'text-green-500', subTextColor: 'text-green-500', iconBg: 'bg-green-50' },
  { type: 'on_loan', val: 5, subKey: 'checked', icon: LucideHistory, borderColor: 'border-blue-100 hover:border-blue-500/30', textColor: 'text-blue-500', subTextColor: 'text-blue-500', iconBg: 'bg-blue-50' },
  { type: 'missing', val: 3, subKey: 'action', icon: LucideAlertTriangle, borderColor: 'border-orange-100 hover:border-orange-500/30', textColor: 'text-orange-500', subTextColor: 'text-orange-500', iconBg: 'bg-orange-50' },
  { type: 'extra', val: 1, subKey: 'critical', icon: LucideShieldCheck, borderColor: 'border-red-100 hover:border-red-500/30', textColor: 'text-red-500', subTextColor: 'text-red-500', iconBg: 'bg-red-50' }
]

const inventoryItems = [
  { sku: 'SKU-882103-X', name: 'MacBook Pro M2 - Space Gray', sys: 0, phy: 1, variance: 1, status: 'EXTRA', location: 'Main Warehouse', subloc: 'Shelf B4' },
  { sku: 'SKU-11920-P', name: 'Logitech MX Master 3S', sys: 12, phy: 9, variance: -3, status: 'MISSING', location: 'Office Area', subloc: 'Zone C' },
  { sku: 'SKU-7729-M', name: 'Dell UltraSharp 32" 4K', sys: 5, phy: 0, variance: -5, status: 'ON LOAN', location: 'Tech Lab', subloc: 'Checked Out' },
  { sku: 'SKU-4412-A', name: 'iPad Air (5th Gen) 64GB', sys: 25, phy: 25, variance: 0, status: 'MATCH', location: 'Store Front', subloc: 'Cabinet 2' }
]

const missingItems = [
  { code: 'DOC-INV-9901', desc: 'Master Agreement - Site A', type: '(Original)' },
  { code: 'DOC-INV-9905', desc: 'Land Certificate Deed -', type: 'Sector 4' }
]

const getStatusStyles = (status) => {
  switch (status) {
    case 'MATCH': return 'bg-green-50 text-green-500 border-green-100'
    case 'ON LOAN': return 'bg-blue-50 text-blue-500 border-blue-100'
    case 'MISSING': return 'bg-orange-50 text-orange-500 border-orange-100'
    case 'EXTRA': return 'bg-red-50 text-red-500 border-red-100'
    default: return 'bg-slate-50 text-slate-500 border-slate-100'
  }
}
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>

