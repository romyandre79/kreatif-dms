<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- State: Reconciliation Report -->
    <template v-if="state === 'report'">
      <!-- Page Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <div class="flex items-center gap-3 text-[9px] font-black text-slate-400 uppercase tracking-widest mb-1">
            <span class="px-2 py-0.5 bg-blue-50 text-blue-500 rounded border border-blue-100">{{ $t('stock.reconciliation.badge') }}</span>
            <span>Session: {{ sessionNo || 'ST-TEMP' }}</span>
          </div>
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
            {{ $t('stock.reconciliation.title') }}
          </h1>
          <p class="text-xs font-bold text-slate-500">
            {{ $t('stock.reconciliation.subtitle', { date: createdDateStr }) }}
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
              <tr v-for="item in filteredItems" :key="item.ID" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all group">
                <td class="p-8 px-10">
                  <div class="space-y-1">
                    <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ item.ItemName?.String || item.ItemName }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ item.SkuCode?.String || item.SkuCode }}</p>
                  </div>
                </td>
                <td class="p-8 text-center text-sm font-black text-slate-700 dark:text-slate-300">{{ item.SystemQty }}</td>
                <td class="p-8 text-center text-sm font-black text-slate-700 dark:text-slate-300">{{ item.PhysicalQty }}</td>
                <td class="p-8 text-center">
                  <span :class="`text-sm font-black ${item.PhysicalQty - item.SystemQty > 0 ? 'text-red-500' : item.PhysicalQty - item.SystemQty < 0 ? 'text-orange-500' : 'text-green-500'}`">
                    {{ item.PhysicalQty - item.SystemQty > 0 ? '+' + (item.PhysicalQty - item.SystemQty) : (item.PhysicalQty - item.SystemQty) }}
                  </span>
                </td>
                <td class="p-8 text-center">
                  <span :class="`px-3 py-1 rounded-md text-[8px] font-black tracking-widest border ${getStatusStyles(item.Status)}`">
                    {{ item.Status }}
                  </span>
                </td>
                <td class="p-8 px-10">
                  <div class="space-y-1">
                    <p class="text-[11px] font-bold text-slate-600 dark:text-slate-400">{{ item.LocationName?.String || 'Unassigned' }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase">{{ item.SubLocation?.String || 'Unassigned' }}</p>
                  </div>
                </td>
              </tr>
              <tr v-if="filteredItems.length === 0">
                <td colspan="6" class="p-16 text-center text-xs font-bold text-slate-400 italic">
                  No items found in this audit session.
                </td>
              </tr>
            </tbody>
          </table>
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
                {{ discrepancies.length }} Discrepancies
              </span>
            </div>

            <div class="space-y-10">
              <div v-for="item in discrepancies" :key="item.ID" class="space-y-6">
                <!-- Extra Item Card -->
                <div v-if="item.Status === 'EXTRA'" class="flex items-start gap-6">
                  <div class="w-12 h-12 rounded-2xl bg-orange-50 text-orange-500 flex items-center justify-center shrink-0 shadow-sm border border-orange-100"><LucidePlusSquare class="w-6 h-6" /></div>
                  <div class="flex-grow space-y-3">
                    <div class="flex items-center justify-between">
                      <div class="flex items-center gap-3">
                        <span class="text-[9px] font-black text-orange-500 uppercase tracking-widest bg-orange-50 px-2 py-0.5 rounded">EXTRA ITEM</span>
                        <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">SKU: {{ item.SkuCode?.String || item.SkuCode }}</span>
                      </div>
                      <div class="text-right">
                        <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">RESOLUTION STATUS</p>
                        <p class="text-[10px] font-black text-emerald-500 uppercase">{{ item.Resolution?.String || 'PENDING ACTION' }}</p>
                      </div>
                    </div>
                    <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ item.ItemName?.String || item.ItemName }}</h4>
                    <p class="text-[11px] font-bold text-slate-500 leading-relaxed max-w-xl">
                      Item scanned here physically but expected elsewhere in system record.
                    </p>
                    
                    <button 
                      v-if="!item.Resolution?.String"
                      @click="resolveItem(item.ID, 'relocated', 'Relocated item to scanned shelf area')"
                      class="w-full py-4 bg-primary-500 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-primary-500/20 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95"
                    >
                      <LucideRefreshCcw class="w-4 h-4" />
                      Move Item & Update Inventory Location
                    </button>
                    <div v-else class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl text-xs font-bold text-slate-500 italic">
                      Resolved: {{ item.Resolution?.String }} (Note: {{ item.ResolutionNote?.String }})
                    </div>
                  </div>
                </div>

                <!-- Missing Item Card -->
                <div v-if="item.Status === 'MISSING'" class="flex items-start gap-6">
                  <div class="w-12 h-12 rounded-2xl bg-red-50 text-red-500 flex items-center justify-center shrink-0 shadow-sm border border-red-100"><LucideClipboardX class="w-6 h-6" /></div>
                  <div class="flex-grow space-y-3">
                    <div class="flex items-center justify-between">
                      <div class="flex items-center gap-3">
                        <span class="text-[9px] font-black text-red-500 uppercase tracking-widest bg-red-50 px-2 py-0.5 rounded">MISSING ITEM</span>
                        <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">SKU: {{ item.SkuCode?.String || item.SkuCode }}</span>
                      </div>
                      <div class="text-right">
                        <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">RESOLUTION STATUS</p>
                        <p class="text-[10px] font-black text-red-500 uppercase">{{ item.Resolution?.String || 'PENDING ACTION' }}</p>
                      </div>
                    </div>
                    <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ item.ItemName?.String || item.ItemName }}</h4>
                    <p class="text-[11px] font-bold text-slate-500 leading-relaxed max-w-xl">
                      Expected physically at location but not detected during scan execution.
                    </p>

                    <button 
                      v-if="!item.Resolution?.String"
                      @click="resolveItem(item.ID, 'confirmed_lost', 'Confirmed lost during physical audit')"
                      class="w-full py-4 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-red-500/20 transition-all flex items-center justify-center gap-3 active:scale-95"
                    >
                      <LucideX class="w-4 h-4" />
                      Mark as Confirmed Lost & Adjust System Qty
                    </button>
                    <div v-else class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl text-xs font-bold text-slate-500 italic">
                      Resolved: {{ item.Resolution?.String }} (Note: {{ item.ResolutionNote?.String }})
                    </div>
                  </div>
                </div>
                
                <div class="h-px bg-slate-50 dark:bg-slate-800"></div>
              </div>

              <div v-if="discrepancies.length === 0" class="p-8 text-center text-xs font-bold text-slate-400 italic">
                No discrepancies found to resolve.
              </div>
            </div>
          </div>
        </div>

        <!-- Right Resolution Sidebar -->
        <div class="space-y-8">
          <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-right>
            <div class="flex items-center gap-3">
              <LucidePenTool class="w-5 h-5 text-primary-500" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('stock.resolution.details.title') }}</h3>
            </div>
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.resolution.details.note_label') }}</label>
              <textarea v-model="generalNote" :placeholder="$t('stock.resolution.details.note_placeholder')" class="w-full h-48 px-6 py-5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-lg text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all resize-none custom-scrollbar"></textarea>
              <div class="flex gap-3 text-slate-400">
                <LucideInfo class="w-4 h-4 shrink-0" />
                <p class="text-[10px] font-bold leading-relaxed italic">{{ $t('stock.resolution.details.note_hint') }}</p>
              </div>
            </div>
            <div class="space-y-4 pt-6">
              <button @click="submitToValidation" class="w-full py-5 bg-primary-500 text-white rounded-lg text-[10px] font-black uppercase tracking-widest shadow-xl shadow-primary-500/30 hover:bg-primary-600 transition-all active:scale-95">
                {{ $t('stock.resolution.details.btn_submit') }}
              </button>
              <button @click="state = 'report'" class="w-full text-[10px] font-black text-slate-400 uppercase tracking-widest hover:text-[#1E3A5F] transition-all">
                Cancel
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
                <h3 class="text-xs font-black text-amber-800 dark:text-amber-400 uppercase tracking-widest">
                  {{ unresolvedMissingItems.length }} UNRESOLVED MISSING ITEMS
                </h3>
              </div>
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
                    <th class="py-6 px-4">RESOLUTION</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                  <tr v-for="item in unresolvedMissingItems" :key="item.ID" class="group">
                    <td class="py-8 px-4">
                      <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ item.SkuCode?.String || item.SkuCode }}</p>
                    </td>
                    <td class="py-8 px-4">
                      <div class="space-y-1">
                        <p class="text-sm font-bold text-slate-600 dark:text-slate-300">{{ item.ItemName?.String || item.ItemName }}</p>
                        <p class="text-[9px] font-bold text-slate-400 uppercase">{{ item.LocationName?.String || 'Unassigned' }}</p>
                      </div>
                    </td>
                    <td class="py-8 px-4 text-center font-black text-[#1E3A5F] dark:text-white">{{ item.SystemQty }}</td>
                    <td class="py-8 px-4 text-center font-black text-red-500">{{ item.PhysicalQty }}</td>
                    <td class="py-8 px-4 text-center font-black text-red-500">{{ item.PhysicalQty - item.SystemQty }}</td>
                    <td class="py-8 px-4">
                      <div class="relative group">
                        <select class="w-full appearance-none px-6 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-[10px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest outline-none focus:ring-2 focus:ring-primary-500/20">
                          <option>DECLARE LOST & WRITE OFF</option>
                        </select>
                        <LucideChevronDown class="absolute right-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
                      </div>
                    </td>
                  </tr>
                  <tr v-if="unresolvedMissingItems.length === 0">
                    <td colspan="6" class="p-8 text-center text-xs font-bold text-slate-400 italic">
                      No unresolved missing items.
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
                <input 
                  v-model="pinVal"
                  type="password" 
                  maxlength="6"
                  :placeholder="$t('stock.signature.pin_placeholder')" 
                  class="w-full pl-16 pr-8 py-5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-sm font-black outline-none focus:ring-4 focus:ring-primary-500/10 transition-all text-center tracking-widest" 
                />
              </div>
              
              <div class="space-y-2">
                <div class="flex items-center justify-between">
                  <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Draw Signature</span>
                  <button type="button" @click="clearSignature" class="text-[10px] font-black text-red-500 uppercase tracking-widest hover:underline">Clear</button>
                </div>
                <div class="aspect-[3/1] bg-slate-50 dark:bg-slate-800/50 rounded-2xl border-2 border-dashed border-slate-200 dark:border-slate-700 flex items-center justify-center relative group overflow-hidden">
                  <canvas 
                    ref="signatureCanvas" 
                    class="absolute inset-0 w-full h-full cursor-crosshair bg-transparent z-10"
                    @mousedown="startDrawing"
                    @mousemove="draw"
                    @mouseup="stopDrawing"
                    @mouseleave="stopDrawing"
                    @touchstart="startDrawingTouch"
                    @touchmove="drawTouch"
                    @touchend="stopDrawing"
                  ></canvas>
                  <div class="text-center space-y-2 opacity-30 group-hover:opacity-60 transition-opacity pointer-events-none">
                    <LucideSignature class="w-10 h-10 mx-auto text-slate-400" />
                    <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.signature.pad_hint') }}</p>
                  </div>
                  <div class="absolute inset-x-0 bottom-8 h-px bg-slate-200 dark:bg-slate-700 w-2/3 mx-auto pointer-events-none"></div>
                </div>
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
              <button @click="handleApprove" :disabled="isApproving" class="w-full py-5 bg-primary-500 text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-2xl shadow-primary-500/40 hover:bg-primary-600 transition-all flex items-center justify-center gap-3 active:scale-95 disabled:opacity-50">
                <LucideCheckCircle2 class="w-5 h-5" />
                {{ isApproving ? 'Approving...' : $t('stock.validation.actions.btn_approve') }}
              </button>
              <button @click="state = 'resolution'" class="w-full py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-400 rounded-lg text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
                Go Back
              </button>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, nextTick, watch } from 'vue'
import { useRoute } from 'vue-router'
import { 
  LucideDownload, LucideSearch, LucideCheckCircle2, LucideHistory, 
  LucideAlertTriangle, LucideShieldCheck, LucideClipboardList,
  LucidePlusSquare, LucideClipboardX, LucideCheck, LucideRefreshCcw,
  LucidePenTool, LucideInfo, LucideUserCheck, LucideChevronDown,
  LucideLock, LucideSignature, LucideX
} from 'lucide-vue-next'

const route = useRoute()
const sessionId = ref(route.query.session_id || '')
const missionId = ref(route.query.mission_id || '')

const { $api } = useApi()
const toast = useToast()

const state = ref('report') // 'report', 'resolution', 'validation'
const tableTab = ref('all')
const items = ref([])
const activeMission = ref(null)

const generalNote = ref('')
const pinVal = ref('')
const isApproving = ref(false)

const signatureCanvas = ref(null)
let ctx = null
let drawing = false

const sessionNo = computed(() => activeMission.value?.session_no || '')
const createdDateStr = computed(() => {
  if (!activeMission.value?.created_at?.Time) return 'Oct 24, 2023 • 09:42 AM'
  return new Date(activeMission.value.created_at.Time).toLocaleString()
})

const matchCount = computed(() => items.value.filter(i => i.Status === 'MATCH').length)
const onLoanCount = computed(() => items.value.filter(i => i.Status === 'ON LOAN').length)
const missingCount = computed(() => items.value.filter(i => i.Status === 'MISSING').length)
const extraCount = computed(() => items.value.filter(i => i.Status === 'EXTRA').length)

const summaryCards = computed(() => [
  { type: 'match', val: matchCount.value, subKey: 'valid', icon: LucideCheckCircle2, borderColor: 'border-green-100 hover:border-green-500/30', textColor: 'text-green-500', subTextColor: 'text-green-500', iconBg: 'bg-green-50' },
  { type: 'on_loan', val: onLoanCount.value, subKey: 'checked', icon: LucideHistory, borderColor: 'border-blue-100 hover:border-blue-500/30', textColor: 'text-blue-500', subTextColor: 'text-blue-500', iconBg: 'bg-blue-50' },
  { type: 'missing', val: missingCount.value, subKey: 'action', icon: LucideAlertTriangle, borderColor: 'border-orange-100 hover:border-orange-500/30', textColor: 'text-orange-500', subTextColor: 'text-orange-500', iconBg: 'bg-orange-50' },
  { type: 'extra', val: extraCount.value, subKey: 'critical', icon: LucideShieldCheck, borderColor: 'border-red-100 hover:border-red-500/30', textColor: 'text-red-500', subTextColor: 'text-red-500', iconBg: 'bg-red-50' }
])

const discrepancies = computed(() => {
  return items.value.filter(i => i.Status === 'MISSING' || i.Status === 'EXTRA')
})

const unresolvedMissingItems = computed(() => {
  return items.value.filter(i => i.Status === 'MISSING' && !i.Resolution?.String)
})

const filteredItems = computed(() => {
  if (tableTab.value === 'location') {
    return items.value.filter(i => i.Status === 'EXTRA' || i.Status === 'MISSING')
  }
  return items.value
})

const fetchSessionItems = async () => {
  if (!sessionId.value) return
  try {
    const res = await $api(`/stock/sessions/${sessionId.value}/items`)
    if (res.success && res.data) {
      items.value = res.data || []
    }
  } catch (err) {
    console.error('Error fetching items:', err)
  }
}

const loadMissionDetails = async () => {
  try {
    const res = await $api('/stock/missions?limit=100')
    if (res.success && res.data && res.data.items) {
      const found = res.data.items.find(m => m.id === missionId.value || m.session_id?.Bytes === sessionId.value)
      if (found) {
        activeMission.value = found
        if (!sessionId.value && found.session_id?.Bytes) {
          sessionId.value = found.session_id.Bytes
        }
      }
    }
  } catch (err) {
    console.error('Error loading mission:', err)
  }
}

const resolveItem = async (itemId, resolution, note) => {
  try {
    const res = await $api(`/stock/items/${itemId}/resolve`, {
      method: 'PUT',
      body: {
        resolution,
        note
      }
    })
    if (res.success) {
      toast.success('Item resolved in draft.')
      await fetchSessionItems()
    } else {
      toast.error(res.message || 'Failed to resolve item.')
    }
  } catch (err) {
    console.error(err)
    toast.error(err.data?.message || 'Error resolving item.')
  }
}

const submitToValidation = () => {
  state.value = 'validation'
  nextTick(() => {
    initCanvas()
  })
}

// Signature Canvas implementation
const initCanvas = () => {
  const canvas = signatureCanvas.value
  if (!canvas) return
  canvas.width = canvas.offsetWidth
  canvas.height = canvas.offsetHeight
  ctx = canvas.getContext('2d')
  ctx.strokeStyle = '#1E3A5F'
  ctx.lineWidth = 3
  ctx.lineCap = 'round'
}

const startDrawing = (e) => {
  drawing = true
  ctx.beginPath()
  ctx.moveTo(e.offsetX, e.offsetY)
}

const draw = (e) => {
  if (!drawing) return
  ctx.lineTo(e.offsetX, e.offsetY)
  ctx.stroke()
}

const stopDrawing = () => {
  drawing = false
}

const startDrawingTouch = (e) => {
  drawing = true
  const rect = signatureCanvas.value.getBoundingClientRect()
  const touch = e.touches[0]
  ctx.beginPath()
  ctx.moveTo(touch.clientX - rect.left, touch.clientY - rect.top)
}

const drawTouch = (e) => {
  if (!drawing) return
  e.preventDefault()
  const rect = signatureCanvas.value.getBoundingClientRect()
  const touch = e.touches[0]
  ctx.lineTo(touch.clientX - rect.left, touch.clientY - rect.top)
  ctx.stroke()
}

const clearSignature = () => {
  const canvas = signatureCanvas.value
  if (!canvas) return
  ctx.clearRect(0, 0, canvas.width, canvas.height)
}

const handleApprove = async () => {
  if (!pinVal.value) {
    toast.warning('Please enter your 6-digit secure authorization PIN.')
    return
  }
  isApproving.value = true
  try {
    const res = await $api(`/stock/sessions/${sessionId.value}/approve`, {
      method: 'POST',
      body: {
        pin: pinVal.value,
        resolution_note: generalNote.value
      }
    })
    if (res.success) {
      toast.success('Session approved & inventory permanently reconciled.')
      navigateTo('/stock/missions')
    } else {
      toast.error(res.message || 'Approval failed.')
    }
  } catch (err) {
    console.error(err)
    toast.error(err.data?.message || 'Failed to authorize. Please check your PIN.')
  } finally {
    isApproving.value = false
  }
}

const getStatusStyles = (status) => {
  switch (status) {
    case 'MATCH': return 'bg-green-50 text-green-500 border-green-100'
    case 'ON LOAN': return 'bg-blue-50 text-blue-500 border-blue-100'
    case 'MISSING': return 'bg-orange-50 text-orange-500 border-orange-100'
    case 'EXTRA': return 'bg-red-50 text-red-500 border-red-100'
    default: return 'bg-slate-50 text-slate-500 border-slate-100'
  }
}

watch(state, (newState) => {
  if (newState === 'validation') {
    nextTick(() => {
      initCanvas()
    })
  }
})

onMounted(async () => {
  if (sessionId.value || missionId.value) {
    await loadMissionDetails()
    await fetchSessionItems()
  }
})
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
