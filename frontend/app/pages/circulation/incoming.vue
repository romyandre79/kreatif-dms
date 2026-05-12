<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- State: Validation Form (Step 2) -->
    <template v-if="state === 'validation'">
      <div class="h-[calc(100vh-120px)] flex flex-col lg:flex-row gap-0 overflow-hidden -m-10">
        <!-- Left Side: Document Viewer -->
        <div class="flex-grow bg-slate-100 dark:bg-slate-950 relative flex flex-col">
          <div class="p-6 bg-white/50 dark:bg-slate-900/50 backdrop-blur-md border-b border-slate-200 dark:border-slate-800 flex items-center justify-between z-10">
            <div class="flex items-center gap-2">
              <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideZoomOut class="w-4 h-4" /></button>
              <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideZoomIn class="w-4 h-4" /></button>
              <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideRotateCcw class="w-4 h-4" /></button>
            </div>
            <button class="flex items-center gap-2 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:underline transition-all">
              <LucideFileText class="w-4 h-4" />
              {{ $t('circulation.incoming.viewer.full_ocr') }}
            </button>
          </div>
          <div class="flex-grow overflow-auto p-12 flex justify-center custom-scrollbar">
            <div class="w-full max-w-4xl bg-white dark:bg-slate-900 shadow-2xl rounded-sm p-16 space-y-12 relative min-h-[1200px]" v-motion-slide-visible-top>
              <div class="flex justify-between items-start">
                <div class="w-24 h-24 bg-slate-50 dark:bg-slate-800 rounded border border-slate-100 dark:border-slate-700 flex items-center justify-center text-[8px] font-black text-slate-300">LOGO HERE</div>
                <div class="text-right space-y-2">
                  <h2 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">MEMORANDUM OF UNDERSTANDING</h2>
                  <p class="text-[10px] font-bold text-slate-400">Ref: <span class="bg-amber-100/50 dark:bg-amber-900/20 border-2 border-dashed border-amber-300/50 dark:border-amber-700/50 px-1">LGL.EXT.2026.0045</span></p>
                </div>
              </div>
              <div class="flex justify-between items-end border-b-2 border-[#1E3A5F] pb-4">
                <div class="space-y-1">
                  <p class="text-[10px] font-black text-primary-500 uppercase tracking-widest border-2 border-dashed border-amber-300/50 inline-block px-1">TO:</p>
                  <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase leading-tight">PT AKIRADATA Solutions<br>JL Sudirman Kav 12<br>Jakarta Selatan, 12190</p>
                </div>
                <div class="text-right space-y-1">
                  <p class="text-[10px] font-black text-primary-500 uppercase tracking-widest">DATE:</p>
                  <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase">October 14, 2026</p>
                </div>
              </div>
              <div class="text-center space-y-2 py-4">
                <div class="inline-block border-2 border-dashed border-amber-300/50 p-2">
                  <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">PARTNERSHIP SERVICE AGREEMENT</h3>
                </div>
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-[0.4em]">FOR THE PROVISION OF CLOUD INFRASTRUCTURE SERVICES</p>
              </div>
              <div class="space-y-8 text-[11px] font-medium text-slate-600 dark:text-slate-400 leading-relaxed text-justify">
                <p>This Agreement is made on this 14th day of October 2026, by and between <span class="font-black text-[#1E3A5F] dark:text-white">CloudArch Systems Inc.</span> (hereinafter referred to as the "Service Provider") and <span class="font-black text-[#1E3A5F] dark:text-white">PT AKIRADATA</span> (hereinafter referred to as the "Client").</p>
                <p>WHEREAS, the Service Provider is engaged in the business of providing high-scale enterprise cloud solutions; and WHEREAS, the Client desires to engage the Service Provider for the management and maintenance of its digital archival core.</p>
                <p>NOW, THEREFORE, the parties hereto agree as follows: The initial term of this agreement shall be for a period of twenty-four (24) months commencing from the date of final signature by both parties.</p>
              </div>
              <div class="pt-20 flex justify-between">
                <div class="w-48 border-t border-slate-200 pt-2 text-[8px] font-black text-slate-400 uppercase">Authorized Signature</div>
                <div class="w-48 border-t border-slate-200 pt-2 text-[8px] font-black text-slate-400 uppercase">Date</div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Side: Validation Form -->
        <div class="w-full lg:w-[450px] bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col">
          <div class="flex-grow overflow-y-auto p-10 space-y-12 custom-scrollbar">
            <div class="space-y-1">
              <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('circulation.incoming.form.title') }}</h2>
              <span class="inline-flex items-center gap-2 text-[8px] font-black text-primary-500 uppercase tracking-widest">
                <LucideZap class="w-3 h-3 fill-primary-500" />
                {{ $t('circulation.incoming.form.ocr_badge') }}
              </span>
            </div>
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] border-l-4 border-[#1E3A5F] pl-4">{{ $t('circulation.incoming.form.sections.identity.title') }}</h4>
              <div class="space-y-5">
                <div class="space-y-2">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.incoming.form.sections.identity.ref_no') }}</label>
                  <div class="relative group">
                    <input type="text" value="LGL.EXT.2026.0045" class="w-full px-6 py-4 bg-[#1E3A5F] text-white rounded-xl text-sm font-black tracking-widest outline-none shadow-xl shadow-blue-900/10" readonly />
                    <LucideLock class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-white/40" />
                  </div>
                </div>
                <div class="space-y-2">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.incoming.form.sections.identity.sender') }}</label>
                  <div class="relative">
                    <input type="text" value="CloudArch Systems Inc." class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-sm font-bold outline-none" />
                    <LucideShieldCheck class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-primary-500" />
                  </div>
                </div>
              </div>
            </div>
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] border-l-4 border-[#1E3A5F] pl-4">{{ $t('circulation.incoming.form.sections.content.title') }}</h4>
              <div class="space-y-5">
                <div class="space-y-2">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.incoming.form.sections.content.doc_title') }}</label>
                  <textarea class="w-full px-6 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm font-bold outline-none h-24 resize-none">Cloud Infrastructure Partnership Agreement</textarea>
                  <div class="text-right"><span class="text-[8px] font-black text-slate-300 uppercase">42 / 120 Characters</span></div>
                </div>
                <div class="space-y-2">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.incoming.form.sections.content.subject') }}</label>
                  <textarea class="w-full px-6 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm font-bold outline-none h-32 resize-none">Provision of cloud infrastructure and managed digital archival services for PT AKIRADATA Core systems.</textarea>
                  <div class="text-right"><span class="text-[8px] font-black text-slate-300 uppercase">108 / 250 Characters</span></div>
                </div>
              </div>
            </div>
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] border-l-4 border-[#1E3A5F] pl-4">{{ $t('circulation.incoming.form.sections.flags.title') }}</h4>
              <div class="bg-slate-50 dark:bg-slate-800/50 p-8 rounded-lg space-y-6 border border-slate-100 dark:border-slate-800">
                <div v-for="flag in flags" :key="flag.label" class="flex items-center justify-between">
                  <div class="flex items-center gap-4">
                    <div :class="`w-8 h-8 rounded-lg flex items-center justify-center ${flag.bg} ${flag.color}`"><component :is="flag.icon" class="w-4 h-4" /></div>
                    <span class="text-[10px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest">{{ $t(`circulation.incoming.form.sections.flags.${flag.label}`) }}</span>
                  </div>
                  <span :class="`text-xs font-black uppercase ${flag.color}`">{{ flag.val }}</span>
                </div>
              </div>
            </div>
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] border-l-4 border-[#1E3A5F] pl-4">{{ $t('circulation.incoming.form.sections.manual.title') }}</h4>
              <div class="space-y-4">
                <div v-for="m in [1, 2]" :key="m" class="p-6 bg-slate-50 dark:bg-slate-800/30 rounded-2xl border border-slate-100 dark:border-slate-800 flex items-start gap-4 group cursor-pointer" @click="manualChecked[m-1] = !manualChecked[m-1]">
                  <div :class="`mt-1 w-5 h-5 rounded border-2 transition-all flex items-center justify-center ${manualChecked[m-1] ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200'}`"><LucideCheck v-if="manualChecked[m-1]" class="w-3.5 h-3.5 text-white" /></div>
                  <p class="text-[10px] font-bold text-slate-500 leading-relaxed">{{ $t(`circulation.incoming.form.sections.manual.check_${m}`) }}</p>
                </div>
              </div>
            </div>
          </div>
          <div class="p-10 border-t border-slate-100 dark:border-slate-800 space-y-4 bg-slate-50/30">
            <button class="w-full py-5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">{{ $t('circulation.incoming.form.actions.btn_back') }}</button>
            <button @click="state = 'routing_slip'" class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 active:scale-95">
              <LucideStamp class="w-4 h-4" />
              {{ $t('circulation.incoming.form.actions.btn_register') }}
            </button>
          </div>
        </div>
      </div>
    </template>

    <!-- State: Routing Slip (Step 3) -->
    <template v-else-if="state === 'routing_slip'">
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('circulation.routing_slip.title') }}</h1>
          <div class="flex items-center gap-6">
            <p class="text-xs font-bold text-slate-500">{{ $t('circulation.routing_slip.subtitle') }}</p>
            <div class="flex items-center gap-3">
              <div class="w-32 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden"><div class="h-full bg-[#1E3A5F]" style="width: 60%"></div></div>
              <span class="text-[10px] font-black text-slate-400">3/5</span>
            </div>
          </div>
        </div>
        <div class="flex items-center gap-4">
          <button @click="state = 'validation'" class="px-8 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">{{ $t('circulation.routing_slip.btn_back') }}</button>
          <button @click="state = 'label_generation'" class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-2 active:scale-95"><LucidePrinter class="w-4 h-4" />{{ $t('circulation.routing_slip.btn_print') }}</button>
        </div>
      </div>
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <div class="lg:col-span-2 space-y-10">
          <div class="glass p-10 rounded-lg grid grid-cols-1 md:grid-cols-2 gap-10">
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.routing_slip.doc_summary.title') }}</h4>
              <div class="space-y-4">
                <div class="space-y-1"><p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.routing_slip.doc_summary.ref') }}</p><p class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">LGL.EXT.2026.0045</p></div>
                <div class="space-y-1"><p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.routing_slip.doc_summary.doc_title') }}</p><p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase leading-tight">Partnership Service Agreement</p></div>
              </div>
            </div>
            <div class="space-y-6 md:text-right flex flex-col md:items-end">
              <span class="px-4 py-1 bg-blue-50 text-blue-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-blue-100">{{ $t('circulation.routing_slip.doc_summary.classification') }}</span>
              <div class="space-y-4 mt-2">
                <div class="space-y-1"><p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.routing_slip.doc_summary.sender') }}</p><p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase">CloudArch Systems Inc.</p></div>
                <div class="space-y-1"><p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.routing_slip.doc_summary.received') }}</p><p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase">October 14, 2026</p></div>
              </div>
            </div>
          </div>
          <div class="bg-white dark:bg-slate-900 shadow-2xl rounded-sm p-20 min-h-[1200px] flex flex-col items-center" v-motion-fade>
            <div class="w-full flex justify-between items-start mb-16">
              <div class="space-y-1"><p class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">PT AKIRADATA</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest border-t border-slate-200 pt-1">{{ $t('circulation.routing_slip.slip.header.division') }}</p><h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase mt-6 leading-tight">{{ $t('circulation.routing_slip.slip.header.title') }}</h3></div>
              <div class="text-center space-y-2"><div class="w-24 h-24 bg-slate-900 rounded-lg p-2 flex items-center justify-center"><LucideQrCode class="w-full h-full text-white" /></div><p class="text-[7px] font-black text-slate-400">LGL-2026-0045-AUTO</p></div>
            </div>
            <div class="w-full h-1 bg-slate-900 mb-8"></div>
            <div class="w-full border-2 border-slate-900 divide-y-2 divide-slate-900 mb-12">
              <div v-for="row in [{ label: $t('circulation.routing_slip.slip.form.sender'), val: 'CloudArch Systems Inc.' }, { label: $t('circulation.routing_slip.slip.form.subject'), val: 'Partnership Service Agreement for Cloud Infrastructure' }, { label: $t('circulation.routing_slip.slip.form.date'), val: '14 Oktober 2026' }]" :key="row.label" class="flex"><div class="w-48 bg-slate-50 p-4 text-[9px] font-black uppercase tracking-widest border-r-2 border-slate-900">{{ row.label }}</div><div class="flex-grow p-4 text-[10px] font-bold uppercase">{{ row.val }}</div></div>
              <div class="flex">
                <div class="w-48 bg-slate-50 p-4 text-[9px] font-black uppercase tracking-widest border-r-2 border-slate-900">{{ $t('circulation.routing_slip.slip.form.nature') }}</div>
                <div class="flex-grow p-4 flex gap-6">
                  <div v-for="n in ['urgent', 'normal', 'confidential']" :key="n" class="flex items-center gap-2"><div :class="`w-4 h-4 border-2 border-slate-900 rounded-sm flex items-center justify-center ${n === 'urgent' ? 'bg-slate-900' : ''}`"><LucideCheck v-if="n === 'urgent'" class="w-3 h-3 text-white" /></div><span class="text-[9px] font-bold uppercase">{{ $t(`circulation.routing_slip.slip.form.nature_vals.${n}`) }}</span></div>
                </div>
              </div>
            </div>
            <div class="w-full space-y-4">
              <h4 class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest border-b-2 border-[#1E3A5F] pb-2">{{ $t('circulation.routing_slip.slip.table.title') }}</h4>
              <table class="w-full border-2 border-[#1E3A5F] text-left">
                <thead>
                  <tr class="bg-[#1E3A5F] text-white text-[8px] font-black uppercase tracking-widest">
                    <th class="p-3 border-r border-white/20 w-12">{{ $t('circulation.routing_slip.slip.table.cols.no') }}</th>
                    <th class="p-3 border-r border-white/20">{{ $t('circulation.routing_slip.slip.table.cols.name') }}</th>
                    <th class="p-3 border-r border-white/20 w-24">{{ $t('circulation.routing_slip.slip.table.cols.date') }}</th>
                    <th class="p-3 border-r border-white/20 w-32">{{ $t('circulation.routing_slip.slip.table.cols.sign') }}</th>
                    <th class="p-3">{{ $t('circulation.routing_slip.slip.table.cols.note') }}</th>
                  </tr>
                </thead>
                <tbody class="divide-y-2 divide-[#1E3A5F]"><tr v-for="(r, i) in recipients" :key="i" class="text-[9px] font-bold"><td class="p-6 border-r-2 border-[#1E3A5F] text-center">{{ i + 1 }}</td><td class="p-6 border-r-2 border-[#1E3A5F] space-y-1"><p class="font-black uppercase text-xs">{{ r.name }}</p><p class="text-[7px] text-slate-400 uppercase tracking-widest">{{ r.role }}</p></td><td class="p-6 border-r-2 border-[#1E3A5F]"></td><td class="p-6 border-r-2 border-[#1E3A5F]"></td><td class="p-6 text-slate-300 italic">{{ r.note || '' }}</td></tr></tbody>
              </table>
            </div>
            <div class="mt-auto w-full flex justify-between text-[7px] font-bold text-slate-300 uppercase tracking-widest border-t pt-4"><span>AUDIT TRAIL ID: BF07-7729-103X-E | GENERATED ON OCT 14, 2026 14:22 WIB</span><span>PAGE 1 OF 1</span></div>
          </div>
        </div>
        <div class="space-y-8">
          <div class="glass p-10 rounded-lg space-y-8" v-motion-slide-right>
            <div class="flex items-center gap-3"><LucideShare2 class="w-5 h-5 text-primary-500" /><h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.routing_slip.routing_summary.title') }}</h3></div>
            <div class="space-y-6">
              <div v-for="s in [{ label: $t('circulation.routing_slip.routing_summary.recipients'), val: '3' }, { label: $t('circulation.routing_slip.routing_summary.type'), val: $t('circulation.routing_slip.routing_summary.type_val') }, { label: $t('circulation.routing_slip.routing_summary.deadline'), val: 'October 21, 2026', color: 'text-red-500' }]" :key="s.label" class="flex items-center justify-between py-2 border-b border-slate-50 dark:border-slate-800/50"><span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ s.label }}</span><span :class="`text-xs font-black uppercase tracking-tight ${s.color || 'text-[#1E3A5F] dark:text-white'}`">{{ s.val }}</span></div>
            </div>
            <div class="pt-10 flex gap-4"><div v-for="r in recipients" :key="r.name" class="w-10 h-10 rounded-full border-2 border-white dark:border-slate-800 bg-slate-100 dark:bg-slate-700 flex items-center justify-center text-[10px] font-black text-[#1E3A5F] dark:text-white shadow-sm -ml-4 first:ml-0">{{ r.name.split(' ').map(n => n[0]).join('') }}</div></div>
          </div>
        </div>
      </div>
    </template>

    <!-- State: QR Label Generation (Step 4) -->
    <template v-else-if="state === 'label_generation'">
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('circulation.labels.title') }}</h1>
          <div class="flex items-center gap-6">
            <p class="text-xs font-bold text-slate-500">{{ $t('circulation.labels.subtitle') }}</p>
            <div class="flex items-center gap-3">
              <div class="w-32 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden"><div class="h-full bg-[#1E3A5F]" style="width: 80%"></div></div>
              <span class="text-[10px] font-black text-slate-400">4/5</span>
            </div>
          </div>
        </div>
        <div class="flex items-center gap-4">
          <button @click="state = 'routing_slip'" class="px-8 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">{{ $t('circulation.routing_slip.btn_back') }}</button>
          <button @click="state = 'signature'" class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-2 active:scale-95"><LucideCheckCircle2 class="w-4 h-4" />Selesaikan & Tanda Tangan</button>
        </div>
      </div>
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
        <div class="lg:col-span-7 space-y-10">
          <div class="glass p-12 rounded-lg flex flex-col items-center text-center space-y-10">
            <div class="relative bg-white p-12 rounded-lg shadow-2xl border border-slate-100"><div class="absolute top-4 right-4 px-3 py-1 bg-green-50 text-green-500 rounded-lg text-[8px] font-black uppercase tracking-widest border border-green-100">{{ $t('circulation.labels.preview.quality') }}</div><div class="w-64 h-64 border-4 border-slate-50 p-4"><LucideQrCode class="w-full h-full text-slate-900" /></div></div>
            <div class="space-y-3"><h3 class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">LGL.EXT.2026.0045</h3><div class="flex items-center justify-center gap-6 text-[10px] font-black text-slate-400 uppercase tracking-widest"><span class="flex items-center gap-2"><LucideCalendar class="w-3.5 h-3.5" /> {{ $t('circulation.labels.preview.generated', { date: 'Oct 24, 2023' }) }}</span><span class="flex items-center gap-2"><LucideFileText class="w-3.5 h-3.5" /> {{ $t('circulation.labels.preview.type') }}</span></div></div>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="glass p-8 rounded-lg space-y-6"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.labels.output.title') }}</h4><div class="space-y-4"><div v-for="opt in ['print', 'pdf', 'rfid_tag']" :key="opt" class="flex items-center gap-4 group cursor-pointer" @click="outputOption = opt"><div :class="`w-5 h-5 rounded-full border-2 flex items-center justify-center transition-all ${outputOption === opt ? 'border-[#1E3A5F] bg-[#1E3A5F]' : 'border-slate-200'}`"><div class="w-1.5 h-1.5 bg-white rounded-full" v-if="outputOption === opt"></div></div><span class="text-xs font-bold text-slate-600 dark:text-slate-300">{{ $t(`circulation.labels.output.${opt}`) }}</span></div></div></div>
            <div class="glass p-8 rounded-lg space-y-6"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.labels.printer.title') }}</h4><div class="space-y-6"><div class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl border border-slate-100 dark:border-slate-800"><p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('circulation.labels.printer.name') }}</p></div><button class="w-full py-4 bg-slate-100 dark:bg-slate-800 text-slate-500 dark:text-slate-400 rounded-xl text-[9px] font-black uppercase tracking-widest hover:bg-slate-200 transition-all flex items-center justify-center gap-2"><LucideRotateCcw class="w-3.5 h-3.5" />{{ $t('circulation.labels.printer.btn_test') }}</button></div></div>
          </div>
        </div>
        <div class="lg:col-span-5 space-y-8" v-if="rfidEnabled" v-motion-slide-right>
          <div class="bg-green-50 dark:bg-green-900/10 p-6 rounded-2xl border border-green-100 dark:border-green-800 flex items-center justify-between"><div class="flex items-center gap-4"><LucideRadioTower class="w-5 h-5 text-green-500" /><span class="text-[10px] font-black text-green-600 uppercase tracking-widest">{{ $t('circulation.labels.rfid.ready') }}</span></div><div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div></div>
          <div class="glass p-10 rounded-lg space-y-10"><p class="text-xs font-bold text-slate-500 leading-relaxed">{{ $t('circulation.labels.rfid.desc') }}</p><div class="space-y-6"><h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.labels.rfid.encoding.title') }}</h4><div class="space-y-4"><div v-for="lvl in ['ref_only', 'full', 'url']" :key="lvl" class="flex items-center gap-4 group cursor-pointer" @click="encodingLevel = lvl"><div :class="`w-5 h-5 rounded-full border-2 flex items-center justify-center transition-all ${encodingLevel === lvl ? 'border-[#1E3A5F] bg-[#1E3A5F]' : 'border-slate-200'}`"><div class="w-1.5 h-1.5 bg-white rounded-full" v-if="encodingLevel === lvl"></div></div><span class="text-[11px] font-bold text-slate-600 dark:text-slate-300">{{ $t(`circulation.labels.rfid.encoding.${lvl}`) }}</span></div></div></div><div class="space-y-6 pt-10 border-t border-slate-50 dark:border-slate-800"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.labels.rfid.design_preview') }}</h4><div class="bg-slate-50 dark:bg-slate-800 rounded-3xl p-10 flex items-center justify-center relative group overflow-hidden border border-slate-100 dark:border-slate-800"><div class="bg-white p-4 shadow-xl border border-slate-100 flex items-start gap-4"><div class="space-y-1 shrink-0"><p class="text-[4px] font-black text-slate-400 uppercase">PT AKIRADATA</p><p class="text-[6px] font-black text-[#1E3A5F]">LGL.EXT.2<br>026.0045</p></div><div class="w-10 h-10 bg-slate-900 rounded-sm"></div><div class="absolute bottom-6 left-1/2 -translate-x-1/2 flex items-center gap-1 opacity-20"><LucideRadioTower class="w-2 h-2" /><span class="text-[4px] font-black uppercase">RFID ENABLED SMART LABEL</span><LucideZap class="w-2 h-2" /></div></div></div><p class="text-center text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('circulation.labels.rfid.standard') }}</p></div></div>
        </div>
      </div>
    </template>

    <!-- State: Digital Signature (Step 5) -->
    <template v-else-if="state === 'signature'">
      <div class="flex flex-col md:flex-row md:items-start justify-between gap-10">
        <div class="flex-grow space-y-12">
          <div class="space-y-2"><p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">DOCUMENT ARCHIVE ID</p><h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter uppercase">LGL.EXT.2026.0045</h1><p class="text-sm font-bold text-slate-500">Surat Permohonan Kerjasama Bisnis</p></div>
          <div class="flex items-center gap-6"><span class="px-4 py-1.5 bg-green-50 text-green-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2"><LucideCheckCircle2 class="w-3.5 h-3.5" />{{ $t('circulation.signature.signed_by', { name: 'Head Legal' }) }}</span><span class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.signature.next', { name: 'GM LEGAL' }) }}</span></div>
          <div class="glass p-12 rounded-lg space-y-10">
            <div class="flex flex-col md:flex-row md:items-center justify-between gap-6"><div class="space-y-1"><h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('circulation.signature.title') }}</h3><p class="text-xs font-bold text-slate-500">{{ $t('circulation.signature.subtitle') }}</p></div><div class="flex bg-slate-100 dark:bg-slate-800 p-1 rounded-xl"><button v-for="t in ['digital', 'upload', 'typed']" :key="t" :class="`px-4 py-2 rounded-lg text-[9px] font-black uppercase tracking-widest transition-all ${sigTab === t ? 'bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white shadow-sm' : 'text-slate-400'}`" @click="sigTab = t">{{ $t(`circulation.signature.tabs.${t}`) }}</button></div></div>
            <div class="aspect-[2/1] bg-slate-50 dark:bg-slate-950 rounded-lg border-2 border-dashed border-slate-200 dark:border-slate-800 flex items-center justify-center relative group overflow-hidden"><button class="absolute top-6 right-6 text-[9px] font-black text-red-500 uppercase tracking-widest flex items-center gap-2 hover:opacity-70 transition-all"><LucideTrash2 class="w-3.5 h-3.5" />{{ $t('circulation.signature.clear') }}</button><div class="absolute inset-0 flex items-center justify-center pointer-events-none opacity-[0.05] tracking-[1em] select-none"><span class="text-4xl font-black uppercase">PLACE SIGNATURE HERE</span></div><LucideSignature class="w-32 h-32 text-[#1E3A5F] dark:text-white opacity-40 group-hover:scale-110 transition-transform duration-700" /></div>
            <div class="flex gap-4 p-8 bg-slate-50/50 dark:bg-slate-800/30 rounded-2xl border border-slate-100 dark:border-slate-800 group cursor-pointer" @click="sigAgreed = !sigAgreed"><div :class="`mt-1 w-6 h-6 rounded border-2 flex items-center justify-center transition-all shrink-0 ${sigAgreed ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200'}`"><LucideCheck v-if="sigAgreed" class="w-4 h-4 text-white" /></div><p class="text-xs font-bold text-slate-500 leading-relaxed">{{ $t('circulation.signature.agreement', { name: 'Head Legal Name', date: 'April 5, 2026, 14:35 WIB' }) }}</p></div>
            <div class="space-y-4"><label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.signature.notes_label') }}</label><textarea :placeholder="$t('circulation.signature.notes_placeholder')" class="w-full h-32 px-6 py-5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all resize-none custom-scrollbar"></textarea><div class="text-right"><span class="text-[8px] font-black text-slate-300">0 / 200 Characters</span></div></div>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4"><button v-for="s in [{ id: 'approved', color: 'text-green-500', bg: 'bg-green-50/50', border: 'border-green-200', dot: 'bg-green-500' }, { id: 'conditional', color: 'text-amber-500', bg: 'bg-amber-50/50', border: 'border-amber-200', dot: 'bg-amber-500' }, { id: 'clarify', color: 'text-orange-500', bg: 'bg-orange-50/50', border: 'border-orange-200', dot: 'bg-orange-500' }, { id: 'rejected', color: 'text-red-500', bg: 'bg-red-50/50', border: 'border-red-200', dot: 'bg-red-500' }]" :key="s.id" :class="`p-4 rounded-xl border ${s.border} ${s.bg} flex flex-col items-center gap-2 transition-all hover:scale-105 active:scale-95 group`" @click="approvalStatus = s.id"><div :class="`w-4 h-4 rounded-full border-2 ${s.id === approvalStatus ? s.dot : 'border-slate-200'} flex items-center justify-center`"><div class="w-1 h-1 bg-white rounded-full"></div></div><span :class="`text-[9px] font-black uppercase tracking-widest ${s.color}`">{{ s.id }}</span></button></div>
          </div>
          <div class="flex flex-col md:flex-row items-center justify-between gap-6 px-4"><button class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-3 hover:text-[#1E3A5F] transition-all"><LucideAnchor class="w-4 h-4" />{{ $t('circulation.signature.actions.hold') }}</button><button @click="state = 'rescan'" class="px-16 py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center gap-4 active:scale-95">{{ $t('circulation.signature.actions.finish') }}<LucideArrowRight class="w-5 h-5" /></button></div>
        </div>
        <div class="w-full lg:w-96 space-y-8">
          <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-right>
            <div class="flex items-center justify-between"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('signature.sidebar.identity') }}</h4><div class="relative w-16 h-16"><svg class="w-full h-full -rotate-90" viewBox="0 0 36 36"><path class="text-slate-100 dark:text-slate-800 stroke-current" stroke-width="3" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" /><path class="text-[#1E3A5F] stroke-current" stroke-width="3" stroke-dasharray="33, 100" stroke-linecap="round" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" /></svg><div class="absolute inset-0 flex flex-col items-center justify-center"><span class="text-xs font-black text-[#1E3A5F] dark:text-white">1/3</span></div></div></div>
            <div class="flex items-center gap-5"><div class="w-14 h-14 rounded-2xl bg-slate-50 dark:bg-slate-800 text-[#1E3A5F] dark:text-white flex items-center justify-center text-lg font-black border border-slate-100 dark:border-slate-700">HL</div><div class="space-y-0.5"><p class="text-base font-black text-[#1E3A5F] dark:text-white">Head Legal Name</p><p class="text-[10px] font-bold text-slate-400 uppercase">Head Legal, Legal Department</p></div></div>
            <div class="space-y-6 pt-6 border-t border-slate-50 dark:border-slate-800"><div v-for="d in [{ label: $t('circulation.signature.sidebar.location'), val: 'Central Document', type: 'select' }, { label: $t('circulation.signature.sidebar.timestamp'), val: 'April 5, 2026, 14:35' }, { label: $t('circulation.signature.sidebar.ip'), val: '192.168.1.104' }]" :key="d.label" class="flex items-center justify-between"><span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ d.label }}</span><div v-if="d.type === 'select'" class="flex items-center gap-2 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ d.val }} <LucideChevronDown class="w-3 h-3" /></div><span v-else class="text-[10px] font-black text-slate-500 uppercase tracking-tight">{{ d.val }}</span></div></div>
          </div>
          <div class="glass p-10 rounded-lg space-y-6"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.signature.sidebar.viewport') }}</h4><div class="aspect-[1/1.4] bg-slate-50 dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700 p-4 shadow-inner"><div class="w-full h-full bg-white dark:bg-slate-900 rounded shadow-lg overflow-hidden flex flex-col p-4 space-y-2 opacity-50"><div class="h-2 bg-slate-100 dark:bg-slate-800 rounded w-1/2 mx-auto"></div><div v-for="l in 20" :key="l" class="h-1 bg-slate-50 dark:bg-slate-800 rounded" :style="{ width: Math.random() * 40 + 60 + '%' }"></div></div></div></div>
          <div class="p-8 bg-orange-50/50 dark:bg-orange-900/10 rounded-lg border border-orange-100 dark:border-orange-800 flex gap-4"><LucideInfo class="w-5 h-5 text-orange-500 shrink-0" /><p class="text-[10px] font-bold text-slate-500 leading-relaxed italic">{{ $t('circulation.signature.sidebar.info') }}</p></div>
        </div>
      </div>
    </template>

    <!-- State: Rescan Signed Routing Slip (Step 5 Final) -->
    <template v-else-if="state === 'rescan'">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('circulation.rescan.title') }}</h1>
        <div class="flex items-center gap-6">
          <p class="text-xs font-bold text-slate-500">{{ $t('circulation.rescan.subtitle') }}</p>
          <div class="flex items-center gap-4">
            <div class="w-32 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden"><div class="h-full bg-[#1E3A5F]" style="width: 100%"></div></div>
            <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">VERIFICATION STAGE</span>
          </div>
        </div>
      </div>
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
        <div class="lg:col-span-4 space-y-8">
          <div class="glass p-8 rounded-lg space-y-6">
            <div class="flex items-center justify-between"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.rescan.device.title') }}</h4><span class="px-3 py-1 bg-green-50 text-green-500 rounded-full text-[8px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-1.5"><span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>{{ $t('circulation.rescan.device.ready') }}</span></div>
            <div class="flex items-center gap-5 p-6 bg-slate-50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800"><div class="w-12 h-12 rounded-xl bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white flex items-center justify-center shadow-sm"><LucidePrinter class="w-6 h-6" /></div><div class="space-y-0.5"><p class="text-sm font-black text-[#1E3A5F] dark:text-white">{{ $t('circulation.rescan.device.name') }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-tight">{{ $t('circulation.rescan.device.details') }}</p></div></div>
          </div>
          <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 rounded-lg space-y-6"><h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.rescan.checklist.title') }}</h4><ul class="space-y-6"><li v-for="i in [1, 2, 3]" :key="i" class="flex gap-4"><div class="w-6 h-6 rounded-full bg-blue-50 text-[#1E3A5F] flex items-center justify-center text-[10px] font-black shrink-0">{{ i }}</div><p class="text-[11px] font-bold text-slate-500 leading-relaxed">{{ $t(`circulation.rescan.checklist.item_${i}`) }}</p></li></ul></div>
          <div class="space-y-4"><button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center justify-center gap-3"><LucideScanLine class="w-5 h-5" />{{ $t('circulation.rescan.buttons.scan') }}</button><button class="w-full py-5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-500 dark:text-slate-400 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-3"><LucideUpload class="w-5 h-5" />{{ $t('circulation.rescan.buttons.upload') }}</button></div>
          <div class="space-y-3 px-2"><div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest"><span class="text-slate-400">{{ $t('circulation.rescan.ocr.title') }}</span><span class="text-[#1E3A5F] dark:text-white">{{ $t('circulation.rescan.ocr.val') }}</span></div><div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden"><div class="h-full bg-[#1E3A5F]" style="width: 85%"></div></div></div>
        </div>
        <div class="lg:col-span-8 space-y-10">
          <div class="bg-slate-100 dark:bg-slate-950 rounded-lg aspect-[1/1.3] relative shadow-inner overflow-hidden border border-slate-200 dark:border-slate-800 flex items-center justify-center p-12">
            <div class="absolute top-10 right-10 z-20" v-motion-pop><div class="bg-white/90 dark:bg-slate-900/90 backdrop-blur-md p-6 rounded-lg border border-green-100 dark:border-green-800 flex items-center gap-5 shadow-2xl"><div class="w-10 h-10 rounded-full bg-green-500 text-white flex items-center justify-center shadow-lg shadow-green-500/20"><LucideCheck class="w-5 h-5" /></div><div class="space-y-0.5"><p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('circulation.rescan.detection.title') }}</p><p class="text-[9px] font-bold text-green-500 uppercase tracking-tight">{{ $t('circulation.rescan.detection.desc', { count: 4, total: 4 }) }}</p></div></div></div>
            <div class="w-full h-full bg-white shadow-2xl rounded-sm p-12 space-y-8 relative overflow-hidden"><div class="flex justify-between items-start"><div class="flex items-center gap-4"><div class="w-12 h-12 bg-slate-900 rounded-sm"></div><div class="space-y-0.5"><p class="text-[10px] font-black">PT AKIRADATA</p><p class="text-[6px] font-bold text-slate-400">Logistics & Distribution</p></div></div><div class="text-right space-y-1"><p class="text-[9px] font-black">LEMBAR DISPOSISI</p><p class="text-[6px] font-bold text-slate-400 uppercase">Ref: LGL/2023/REC/055</p></div></div><div class="w-full h-px bg-slate-900"></div><div class="space-y-3"><div v-for="l in ['SURAT DARI:', 'PERIHAL:', 'INSTRUKSI / DISPOSISI:']" :key="l" class="flex gap-4"><span class="text-[8px] font-black w-24">{{ l }}</span><div class="flex-grow h-4 bg-slate-50 rounded"></div></div></div><div class="h-48 border-2 border-slate-100 rounded-xl relative p-8 flex items-center justify-center"><p class="text-[10px] font-bold text-slate-300 italic text-center max-w-sm">"Segera lakukan audit dokumen pendukung dan verifikasi lapangan. Laporkan hasil dalam 2 hari kerja."</p><div class="absolute bottom-4 right-8 text-center space-y-1"><LucideSignature class="w-12 h-12 text-[#1E3A5F] opacity-70" /><p class="text-[6px] font-black uppercase">Direktur Operasional</p></div></div><div class="absolute top-[28%] right-[10%] w-24 h-12 border-2 border-primary-500/30 bg-primary-500/5 rounded flex items-center justify-center"><span class="text-[6px] font-black text-primary-500 bg-white px-1">SIG_01</span></div><div class="absolute top-[60%] left-[10%] w-24 h-12 border-2 border-primary-500/30 bg-primary-500/5 rounded flex items-center justify-center"><span class="text-[6px] font-black text-primary-500 bg-white px-1">SIG_02</span></div><div class="absolute bottom-[10%] right-[30%] w-24 h-12 border-2 border-primary-500/30 bg-primary-500/5 rounded flex items-center justify-center"><span class="text-[6px] font-black text-primary-500 bg-white px-1">SIG_03</span></div></div>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
            <div class="glass p-10 rounded-lg space-y-8"><h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.rescan.quality.title') }}</h4><div class="space-y-6"><div v-for="i in [1, 2, 3]" :key="i" class="flex items-center gap-4 group cursor-pointer"><div class="w-6 h-6 bg-[#1E3A5F] rounded flex items-center justify-center shadow-lg shadow-blue-900/10"><LucideCheck class="w-4 h-4 text-white" /></div><span class="text-xs font-bold text-slate-600 dark:text-slate-300">{{ $t(`circulation.rescan.quality.item_${i}`) }}</span></div></div></div>
            <div class="bg-blue-50/30 dark:bg-blue-900/10 p-10 rounded-lg border border-blue-100 dark:border-blue-800 space-y-8"><h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.rescan.verification.title') }}</h4><div class="flex gap-4 group cursor-pointer" @click="finalAgreed = !finalAgreed"><div :class="`mt-1 w-6 h-6 rounded border-2 transition-all shrink-0 ${finalAgreed ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200'}`"><LucideCheck v-if="finalAgreed" class="w-4 h-4 text-white" /></div><p class="text-[10px] font-bold text-slate-500 leading-relaxed">{{ $t('circulation.rescan.verification.agreement') }}</p></div><div class="flex items-center gap-4 pt-4"><button @click="state = 'signature'" class="px-6 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2"><LucideRotateCcw class="w-3.5 h-3.5" />{{ $t('circulation.rescan.verification.btn_back') }}</button><button @click="state = 'appending'" class="flex-grow py-5 bg-[#1E3A5F] text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 active:scale-95">{{ $t('circulation.rescan.verification.btn_next') }}<LucideArrowRight class="w-4 h-4" /></button></div></div>
          </div>
        </div>
      </div>
    </template>

    <!-- State: Document Appending (Step 5.5) -->
    <template v-else-if="state === 'appending'">
      <div class="flex items-center justify-between">
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('circulation.appending.title') }}</h1>
          <p class="text-xs font-bold text-slate-500">{{ $t('circulation.appending.subtitle') }}</p>
        </div>
        <span class="px-4 py-1.5 bg-green-50 text-green-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2">
          <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
          {{ $t('circulation.appending.ready') }}
        </span>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
        <!-- Original Document Column -->
        <div class="lg:col-span-3 space-y-6">
          <div class="flex items-center justify-between px-2">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2"><LucideFileText class="w-3.5 h-3.5" />{{ $t('circulation.appending.original.title') }}</h4>
            <span class="text-[8px] font-black text-slate-300 uppercase">{{ $t('circulation.appending.original.pages', { count: 5 }) }}</span>
          </div>
          <div class="aspect-[1/1.4] bg-slate-900 rounded-lg shadow-2xl overflow-hidden relative group cursor-zoom-in">
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
            <div class="absolute inset-0 flex flex-col items-center justify-center text-white p-8 text-center space-y-4">
              <p class="text-lg font-black tracking-widest opacity-20">PDE: SAFE WORK</p>
              <div class="w-12 h-12 rounded-full bg-white/10 backdrop-blur-md flex items-center justify-center group-hover:scale-110 transition-transform"><LucideMaximize2 class="w-5 h-5" /></div>
            </div>
          </div>
          <div class="p-5 bg-white dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 flex items-center gap-4 shadow-sm">
            <div class="w-10 h-10 rounded-lg bg-red-50 text-red-500 flex items-center justify-center"><LucideFileText class="w-6 h-6" /></div>
            <div class="space-y-0.5 overflow-hidden">
              <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white truncate">{{ $t('circulation.appending.original.filename') }}</p>
              <p class="text-[8px] font-bold text-slate-400 uppercase">{{ $t('circulation.appending.original.info') }}</p>
            </div>
          </div>
        </div>

        <!-- Operations Column -->
        <div class="lg:col-span-4 glass p-10 rounded-lg space-y-10">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.appending.ops.title') }}</h4>
          <div class="space-y-8">
            <div class="space-y-4">
              <p class="text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.appending.ops.position.title') }}</p>
              <div class="space-y-4">
                <div v-for="pos in ['front', 'end', 'insert']" :key="pos" class="flex items-center gap-4 group cursor-pointer" @click="appendPos = pos">
                  <div :class="`w-5 h-5 rounded-full border-2 flex items-center justify-center transition-all ${appendPos === pos ? 'border-[#1E3A5F] bg-[#1E3A5F]' : 'border-slate-200'}`"><div class="w-1.5 h-1.5 bg-white rounded-full" v-if="appendPos === pos"></div></div>
                  <span class="text-xs font-bold text-slate-600 dark:text-slate-300">{{ $t(`circulation.appending.ops.position.${pos}`) }}</span>
                </div>
              </div>
            </div>
            <div class="space-y-4">
              <p class="text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.appending.ops.options.title') }}</p>
              <div class="space-y-4">
                <div v-for="opt in ['bookmark', 'metadata', 'version']" :key="opt" class="flex items-center gap-4 group cursor-pointer" @click="appendOptions[opt] = !appendOptions[opt]">
                  <div :class="`w-5 h-5 rounded border flex items-center justify-center transition-all ${appendOptions[opt] ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'border-slate-200'}`"><LucideCheck v-if="appendOptions[opt]" class="w-3.5 h-3.5 text-white" /></div>
                  <span class="text-xs font-bold text-slate-600 dark:text-slate-300">{{ $t(`circulation.appending.ops.options.${opt}`) }}</span>
                </div>
              </div>
            </div>
          </div>
          <div class="pt-10 space-y-4">
            <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex flex-col items-center gap-1 group">
              <div class="flex items-center gap-3"><LucideAnchor class="w-4 h-4 group-hover:rotate-12 transition-transform" />{{ $t('circulation.appending.ops.btn_execute') }}</div>
            </button>
            <p class="text-center text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('circulation.appending.ops.finalize') }}</p>
          </div>
        </div>

        <!-- Preview Column -->
        <div class="lg:col-span-5 space-y-8">
          <div class="glass p-10 rounded-lg space-y-8">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.appending.preview.title') }}</h4>
            <div class="grid grid-cols-2 gap-4">
              <div v-for="p in [1, 2, 3, 4]" :key="p" class="aspect-[1/1.4] bg-slate-50 dark:bg-slate-800 rounded-xl border border-slate-100 dark:border-slate-700 p-4 relative group cursor-pointer">
                <div class="w-full h-full bg-white dark:bg-slate-900 rounded shadow-sm opacity-50 flex flex-col p-2 space-y-1">
                  <div v-for="l in 10" :key="l" class="h-0.5 bg-slate-50 dark:bg-slate-800 rounded" :style="{ width: Math.random() * 50 + 50 + '%' }"></div>
                </div>
                <div class="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity bg-black/5 rounded-xl"><LucideMaximize2 class="w-6 h-6 text-slate-400" /></div>
                <span class="absolute bottom-4 left-1/2 -translate-x-1/2 text-[7px] font-black text-slate-300">PAGE {{ p }}</span>
              </div>
            </div>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="p-8 bg-green-50/50 dark:bg-green-900/10 rounded-lg border border-green-100 dark:border-green-800 space-y-4">
              <h4 class="text-[9px] font-black text-green-600 uppercase tracking-widest">{{ $t('circulation.appending.preview.verification.title') }}</h4>
              <div class="space-y-3">
                <div v-for="v in ['original', 'appended', 'signatures']" :key="v" class="flex items-center gap-3">
                  <div class="w-1.5 h-1.5 bg-green-500 rounded-full"></div>
                  <span class="text-[10px] font-bold text-green-600/80">{{ $t(`circulation.appending.preview.verification.${v}`) }}</span>
                </div>
              </div>
            </div>
            <div class="p-8 bg-slate-50 dark:bg-slate-800/50 rounded-lg border border-slate-100 dark:border-slate-800 flex flex-col justify-center gap-1">
              <p class="text-xs font-black text-[#1E3A5F] dark:text-white">LGL.EXT.2026.0045_...</p>
              <p class="text-[9px] font-bold text-slate-400 uppercase">{{ $t('circulation.appending.preview.summary') }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer Info -->
      <div class="flex flex-col md:flex-row items-center justify-between gap-10 p-10 bg-slate-50/50 dark:bg-slate-900/50 rounded-lg border border-slate-100 dark:border-slate-800">
        <div class="flex items-center gap-8 flex-grow">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-xl bg-white dark:bg-slate-800 flex items-center justify-center text-[#1E3A5F] dark:text-white shadow-sm"><LucideAnchor class="w-6 h-6" /></div>
            <div class="space-y-0.5">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.appending.footer.path') }}</p>
              <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white tracking-tight">/ARCHIVE/2026/LGL/EXTERNAL/CONTRACTS/</p>
            </div>
          </div>
          <div class="w-px h-12 bg-slate-200 dark:bg-slate-700 hidden md:block"></div>
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-xl bg-amber-50 text-amber-600 flex items-center justify-center shadow-sm"><LucideCalendar class="w-6 h-6" /></div>
            <div class="space-y-0.5">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.appending.footer.retention') }}</p>
              <p class="text-[11px] font-black text-amber-600 tracking-tight">{{ $t('circulation.appending.footer.retention_val') }}</p>
            </div>
          </div>
        </div>
        <div class="flex items-center gap-6">
          <button @click="state = 'rescan'" class="text-[10px] font-black text-slate-400 uppercase tracking-widest hover:text-red-500 transition-all">{{ $t('circulation.appending.footer.btn_back') }}</button>
          <button @click="state = 'distribution'" class="px-16 py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center gap-4 active:scale-95">
            {{ $t('circulation.appending.footer.btn_save') }}
            <LucideArrowRight class="w-5 h-5" />
          </button>
        </div>
      </div>
    </template>

    <!-- State: Distribution (Step 6) -->
    <template v-else-if="state === 'distribution'">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('circulation.distribution.title') }}</h1>
        <div class="flex items-center justify-between">
          <p class="text-xs font-bold text-slate-500">{{ $t('circulation.distribution.subtitle') }}</p>
          <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Ref ID: 0045/LGL/2026</span>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
        <div class="lg:col-span-8 space-y-10">
          <!-- Recipients Table -->
          <div class="glass rounded-lg overflow-hidden">
            <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.distribution.recipients.title') }}</h3>
              <div class="flex items-center gap-4">
                <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideZap class="w-4 h-4" /></button>
                <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideUpload class="w-4 h-4" /></button>
              </div>
            </div>
            <table class="w-full text-left">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-800/30 text-[9px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800">
                  <th class="p-8 w-20">{{ $t('circulation.distribution.recipients.cols.no') }}</th>
                  <th class="p-8">{{ $t('circulation.distribution.recipients.cols.name') }}</th>
                  <th class="p-8">{{ $t('circulation.distribution.recipients.cols.status') }}</th>
                  <th class="p-8">{{ $t('circulation.distribution.recipients.cols.result') }}</th>
                  <th class="p-8">{{ $t('circulation.distribution.recipients.cols.notif') }}</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="(r, i) in [
                  { name: 'Head Legal', email: 'head.legal@akiradata.co.id' },
                  { name: 'GM Legal', email: 'gm.legal@akiradata.co.id' },
                  { name: 'Section Head', email: 'sec.head@akiradata.co.id' }
                ]" :key="i" class="text-xs font-bold group hover:bg-slate-50/50 transition-all">
                  <td class="p-8 text-slate-300 font-black">{{ i + 1 }}</td>
                  <td class="p-8 space-y-1">
                    <p class="font-black text-[#1E3A5F] dark:text-white uppercase">{{ r.name }}</p>
                    <p class="text-[9px] text-slate-400">{{ r.email }}</p>
                  </td>
                  <td class="p-8">
                    <span class="flex items-center gap-2 text-green-500 text-[10px] font-black uppercase tracking-widest">
                      <LucideCheckCircle2 class="w-3.5 h-3.5" />
                      Completed
                    </span>
                  </td>
                  <td class="p-8">
                    <span class="px-3 py-1 bg-green-50 text-green-500 rounded-lg text-[8px] font-black uppercase tracking-widest border border-green-100">APPROVED</span>
                  </td>
                  <td class="p-8">
                    <span class="flex items-center gap-2 text-blue-500 text-[10px] font-black uppercase tracking-widest">
                      <span class="w-1.5 h-1.5 bg-blue-500 rounded-full animate-pulse"></span>
                      Ready to Send
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Notification Preview -->
          <div class="glass rounded-lg overflow-hidden">
            <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800 cursor-pointer hover:bg-slate-50/30 transition-all">
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><LucideFileText class="w-5 h-5" /></div>
                <div class="space-y-0.5">
                  <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.distribution.preview.title') }}</h3>
                  <p class="text-[10px] font-bold text-slate-400">{{ $t('circulation.distribution.preview.subtitle') }}</p>
                </div>
              </div>
              <LucideChevronDown class="w-5 h-5 text-slate-400" />
            </div>
            <div class="p-12 space-y-8 bg-white dark:bg-slate-900/50">
              <div class="grid grid-cols-1 gap-6">
                <div class="flex items-start gap-10 pb-6 border-b border-slate-50 dark:border-slate-800">
                  <span class="w-24 text-[10px] font-black text-slate-300 uppercase tracking-widest">Subject:</span>
                  <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">[DMS] Dokumen Masuk Tersedia: Surat Permohonan Kerjasama Bisnis - LGL.EXT.2026.0045</span>
                </div>
                <div class="flex items-start gap-10">
                  <span class="w-24 text-[10px] font-black text-slate-300 uppercase tracking-widest">To:</span>
                  <span class="text-xs font-bold text-slate-500 italic">Recipient Name (Recipient Title)</span>
                </div>
              </div>
              <div class="p-10 bg-slate-50 dark:bg-slate-800/30 rounded-lg space-y-8 text-xs font-medium text-slate-600 dark:text-slate-400 leading-relaxed">
                <p>Yth. Bapak/Ibu,</p>
                <p>Pemberitahuan bahwa dokumen digital telah tersedia untuk referensi Anda dalam sistem Manajemen Dokumen BP-07.</p>
                <div class="p-8 bg-white dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-3">
                  <p><span class="font-black text-[#1E3A5F] dark:text-white">No. Dokumen:</span> LGL.EXT.2026.0045</p>
                  <p><span class="font-black text-[#1E3A5F] dark:text-white">Judul:</span> Surat Permohonan Kerjasama Bisnis</p>
                  <p><span class="font-black text-[#1E3A5F] dark:text-white">Prioritas:</span> Normal</p>
                </div>
                <p>Silahkan masuk ke panel Distribusi Digital untuk melihat detail lengkap dan dokumen terkait.</p>
                <div class="pt-8 space-y-1">
                  <p>Terima kasih,</p>
                  <p class="font-black text-[#1E3A5F] dark:text-white">DMS Administrator - PT AKIRADATA</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar Actions -->
        <div class="lg:col-span-4 space-y-8">
          <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-right>
            <div class="flex items-center justify-between">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.distribution.sidebar.title') }}</h4>
              <div class="w-12 h-12 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center shadow-lg shadow-blue-500/10"><LucideCheckCircle2 class="w-6 h-6" /></div>
            </div>
            <div class="space-y-6">
              <div v-for="s in [
                { label: $t('circulation.distribution.sidebar.total'), val: '03', icon: LucideShare2 },
                { label: $t('circulation.distribution.sidebar.size'), val: '2.4 MB (PDF)', icon: LucideFileText },
                { label: $t('circulation.distribution.sidebar.channel'), val: 'Email', icon: LucideZap }
              ]" :key="s.label" class="flex items-center justify-between group cursor-default">
                <div class="flex items-center gap-4">
                  <div class="w-8 h-8 rounded-lg bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400 group-hover:text-[#1E3A5F] transition-all"><component :is="s.icon" class="w-4 h-4" /></div>
                  <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ s.label }}</span>
                </div>
                <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase">{{ s.val }}</span>
              </div>
            </div>
            <div class="pt-10 border-t border-slate-50 dark:border-slate-800 flex gap-4 cursor-pointer group" @click="distAgreed = !distAgreed">
              <div :class="`mt-1 w-6 h-6 rounded border-2 flex items-center justify-center transition-all shrink-0 ${distAgreed ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200'}`"><LucideCheck v-if="distAgreed" class="w-4 h-4 text-white" /></div>
              <p class="text-[10px] font-bold text-slate-400 leading-relaxed">{{ $t('circulation.distribution.sidebar.auth') }}</p>
            </div>
          </div>

          <!-- History -->
          <div class="glass p-10 rounded-lg space-y-8">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.distribution.sidebar.history') }}</h4>
            <div class="space-y-8 relative before:absolute before:left-3 before:top-2 before:bottom-2 before:w-px before:bg-slate-100 dark:before:bg-slate-800">
              <div class="relative pl-10 space-y-1">
                <div class="absolute left-1.5 top-1.5 w-3 h-3 rounded-full bg-green-500 border-4 border-white dark:border-slate-900 z-10"></div>
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">Final Approval Received</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase">System • 10 minutes ago</p>
              </div>
              <div class="relative pl-10 space-y-1">
                <div class="absolute left-1.5 top-1.5 w-3 h-3 rounded-full bg-slate-300 border-4 border-white dark:border-slate-900 z-10"></div>
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">Document Digitized</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase">Admin • 1 hour ago</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="flex items-center justify-between px-4 mt-12">
        <button @click="state = 'appending'" class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-3 hover:text-red-500 transition-all">
          <LucideArrowRight class="w-4 h-4 rotate-180" />
          {{ $t('circulation.distribution.footer.btn_back') }}
        </button>
        <button class="px-20 py-6 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center gap-4 active:scale-95 group">
          {{ $t('circulation.distribution.footer.btn_send') }}
          <LucideShare2 class="w-5 h-5 group-hover:translate-x-1 group-hover:-translate-y-1 transition-transform" />
        </button>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideZoomOut, LucideZoomIn, LucideRotateCcw, LucideFileText, 
  LucideZap, LucideLock, LucideShieldCheck, LucideCheck, LucideStamp,
  LucideScanLine, LucideFileCheck, LucideCopyCheck, LucideChevronRight,
  LucidePrinter, LucideQrCode, LucideShare2, LucideRadioTower, LucideCalendar,
  LucideMaximize2, LucideCheckCircle2, LucideTrash2, LucideSignature,
  LucideChevronDown, LucideInfo, LucideAnchor, LucideArrowRight, LucideUpload
} from 'lucide-vue-next'

const state = ref('validation') // 'validation', 'routing_slip', 'label_generation', 'signature', 'rescan', 'appending', 'distribution'
const manualChecked = ref([false, false])
const rfidEnabled = ref(true)
const outputOption = ref('print')
const encodingLevel = ref('full')
const sigTab = ref('digital')
const sigAgreed = ref(false)
const approvalStatus = ref('approved')
const finalAgreed = ref(false)
const appendPos = ref('end')
const appendOptions = ref({ bookmark: true, metadata: true, version: false })
const distAgreed = ref(false)

const flags = [
  { label: 'ocr_quality', val: '90%', icon: LucideScanLine, bg: 'bg-green-50', color: 'text-green-500' },
  { label: 'completeness', val: '80%', icon: LucideFileCheck, bg: 'bg-orange-50', color: 'text-orange-500' },
  { label: 'duplicate', val: 'UNIQUE', icon: LucideCopyCheck, bg: 'bg-blue-50', color: 'text-blue-500' }
]

const recipients = [
  { name: 'Budi Santoso', role: 'Head of Legal Department', note: 'Harap diperiksa kesesuaian SLA...' },
  { name: 'Siti Aminah', role: 'GM Legal & Compliance' },
  { name: 'Anwar Ibrahim', role: 'Section Head Administration' }
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
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>

