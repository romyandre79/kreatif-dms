<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-32">
    <!-- Header -->
    <div v-if="!isBulkReviewMode && !showProgressModal && !showBulkSuccessModal && !showManifestPreview" class="flex items-center gap-6" v-motion-fade>
      <div class="w-14 h-14 rounded-2xl bg-[#1E3A5F]/5 dark:bg-white/5 border border-slate-200 dark:border-white/10 flex items-center justify-center text-[#1E3A5F] dark:text-white shadow-sm">
        <LucideFileUp class="w-7 h-7" />
      </div>
      <div>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ $t('upload.title') }}</h1>
        <p class="text-slate-500 font-medium mt-1">{{ $t('upload.subtitle') }}</p>
      </div>
    </div>

    <!-- Tabs -->
    <div v-if="!isBulkReviewMode && !showProgressModal && !showBulkSuccessModal && !showManifestPreview" class="flex border-b border-slate-200 dark:border-slate-800" v-motion-fade>
      <button v-for="tab in [{key: 'manual', label: $t('upload.tabs.manual')}, {key: 'bulk', label: $t('upload.tabs.bulk')}]" :key="tab.key" @click="activeTab = tab.key" :class="`px-8 py-4 text-sm font-black uppercase tracking-widest transition-all relative ${activeTab === tab.key ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400 hover:text-slate-600'}`"><div class="flex items-center gap-2"><component :is="tab.key === 'manual' ? LucidePenTool : LucideFileSpreadsheet" class="w-4 h-4" />{{ tab.label }}</div><div v-if="activeTab === tab.key" class="absolute bottom-0 left-0 w-full h-1 bg-[#1E3A5F] dark:bg-primary-500 rounded-t-full"></div></button>
    </div>

    <!-- Bulk Review Mode -->
    <div v-if="isBulkReviewMode && !showProgressModal && !showBulkSuccessModal && !showManifestPreview" class="space-y-8" v-motion-fade>
      <div class="flex flex-col gap-2"><div class="flex items-center gap-3"><LucideFileSpreadsheet class="w-8 h-8 text-[#1E3A5F] dark:text-white" /><h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight">DMS_Bulk_Upload_Batch_01.xlsx</h1></div><p class="text-slate-500 font-medium">{{ $t('upload.bulk.review.subtitle') }}</p></div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="glass p-8 rounded-[2rem] flex flex-col items-center justify-center text-center"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">{{ $t('upload.bulk.review.stats.total') }}</p><div class="flex items-baseline gap-2"><span class="text-4xl font-black text-[#1E3A5F] dark:text-white">45</span><span class="text-sm font-bold text-slate-400">{{ $t('upload.bulk.review.stats.items') }}</span></div></div>
        <div class="glass p-8 rounded-[2rem] flex flex-col items-center justify-center text-center border-l-4 border-green-500"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">{{ $t('upload.bulk.review.stats.valid') }}</p><div class="flex items-baseline gap-2"><span class="text-4xl font-black text-green-500">42</span><span class="text-xs font-black text-green-600/60 uppercase tracking-widest">{{ $t('upload.bulk.review.stats.success_rate', { rate: 93 }) }}</span></div></div>
        <div class="glass p-8 rounded-[2rem] flex flex-col items-center justify-center text-center border-l-4 border-red-500"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">{{ $t('upload.bulk.review.stats.errors') }}</p><div class="flex items-baseline gap-2"><span class="text-4xl font-black text-red-500">3</span><span class="text-xs font-black text-red-600/60 uppercase tracking-widest">{{ $t('upload.bulk.review.stats.action_required') }}</span></div></div>
      </div>
      <div class="glass rounded-[2rem] overflow-hidden shadow-xl shadow-slate-200/50 dark:shadow-none">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-900/50 border-b border-slate-100 dark:border-slate-800">
              <th class="p-6 w-12"><div class="w-5 h-5 rounded border-2 border-slate-200 dark:border-slate-700"></div></th>
              <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('upload.bulk.review.table.status') }}</th>
              <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('upload.bulk.review.table.type') }}</th>
              <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('upload.bulk.review.table.subject') }}</th>
              <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('upload.bulk.review.table.category') }}</th>
              <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('upload.bulk.review.table.date') }}</th>
              <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-widest text-right">{{ $t('upload.bulk.review.table.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
            <template v-for="(row, i) in bulkData" :key="i">
              <tr :class="`group hover:bg-slate-50/30 dark:hover:bg-slate-800/20 transition-colors ${row.hasError ? 'bg-red-50/20' : ''}`"><td class="p-6"><div class="w-5 h-5 rounded border-2 border-slate-200 dark:border-slate-700"></div></td><td class="p-6"><LucideCheckCircle2 v-if="!row.hasError" class="w-5 h-5 text-green-500" /><LucideXCircle v-else class="w-5 h-5 text-red-500" /></td><td class="p-6 text-sm font-bold text-slate-600 dark:text-slate-300">{{ row.type }}</td><td class="p-6"><span v-if="row.subject" class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ row.subject }}</span><span v-else class="text-sm font-bold text-red-500 italic">{{ $t('upload.bulk.review.table.empty') }}</span></td><td class="p-6"><span :class="`text-xs font-black uppercase tracking-widest ${row.category === 'Unknown' ? 'text-red-500' : 'text-slate-500'}`">{{ row.category }}</span></td><td class="p-6 text-xs font-bold text-slate-500">{{ row.date }}</td><td class="p-6"><div class="flex items-center justify-end gap-3 opacity-0 group-hover:opacity-100 transition-opacity"><button class="p-2 text-slate-400 hover:text-primary-500 transition-colors"><LucidePencil class="w-4 h-4" /></button><button class="p-2 text-slate-400 hover:text-red-500 transition-colors"><LucideTrash2 class="w-4 h-4" /></button></div></td></tr>
              <tr v-if="row.hasError" class="bg-red-50/40 dark:bg-red-950/10"><td colspan="7" class="p-4 px-10 border-t border-red-100/50 dark:border-red-900/30"><div class="flex items-center gap-3"><LucideAlertCircle class="w-4 h-4 text-red-500" /><p class="text-[10px] font-black text-red-600 uppercase tracking-widest">{{ row.errorMsg }}</p></div></td></tr>
            </template>
          </tbody>
        </table>
        <div class="p-6 bg-slate-50/30 dark:bg-slate-900/30 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between"><div class="flex items-center gap-4"><span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Bulk Actions:</span><div class="relative"><select class="pl-4 pr-10 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-[10px] font-black uppercase tracking-widest appearance-none outline-none focus:border-primary-500"><option>Select action...</option><option>Delete Selected</option><option>Change Category</option></select><LucideChevronDown class="absolute right-3 top-1/2 -translate-y-1/2 w-3 h-3 text-slate-400 pointer-events-none" /></div><button class="px-6 py-2.5 bg-[#1E3A5F] text-white rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-[#152943] transition-all">Apply</button></div><button class="flex items-center gap-2 text-[10px] font-black text-slate-500 hover:text-[#1E3A5F] transition-colors uppercase tracking-widest"><LucideRefreshCcw class="w-3.5 h-3.5" />{{ $t('upload.bulk.review.footer.revalidate') }}</button></div>
      </div>
    </div>

    <!-- Manual Entry / Bulk Upload Main Content -->
    <div v-if="!isBulkReviewMode && !showProgressModal && !showBulkSuccessModal && !showManifestPreview" class="grid grid-cols-1 lg:grid-cols-3 gap-10">
      <div class="lg:col-span-2 space-y-8">
        <div v-if="activeTab === 'manual'" class="glass rounded-3xl overflow-hidden shadow-xl shadow-slate-200/50 dark:shadow-none" v-motion-slide-visible-bottom>
          <div class="px-10 py-8 border-b border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50"><h2 class="font-black text-xl text-[#1E3A5F] dark:text-white">{{ $t('upload.form.title') }}</h2></div>
          <div class="p-10 space-y-8">
            <div class="grid grid-cols-1 md:grid-cols-3 items-start gap-6"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300 mt-3">{{ $t('upload.form.type') }}<span class="text-red-500 ml-1">*</span></label><div class="md:col-span-2 space-y-2"><div class="relative"><select v-model="form.type" :class="`w-full pl-5 pr-12 py-3.5 bg-white dark:bg-slate-900 border ${errors.type ? 'border-red-500 bg-red-50/10' : 'border-slate-200 dark:border-slate-800'} rounded-2xl text-sm font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all cursor-pointer`"><option value="" disabled selected>{{ $t('upload.form.type_placeholder') }}</option><option>Invoice / Receipt</option><option>Legal Contract</option><option>Personnel File</option><option>Financial Report</option></select><LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" /></div><div v-if="errors.type" class="flex items-center gap-1.5 px-1"><LucideAlertCircle class="w-3 h-3 text-red-500" /><p class="text-[10px] font-black text-red-500 uppercase tracking-widest">{{ errors.type }}</p></div></div></div>
            <div class="grid grid-cols-1 md:grid-cols-3 items-start gap-6"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300 mt-3">{{ $t('upload.form.subject') }}<span class="text-red-500 ml-1">*</span></label><div class="md:col-span-2 space-y-2"><textarea v-model="form.title" :placeholder="$t('upload.form.subject_placeholder')" rows="2" :class="`w-full px-5 py-3.5 bg-white dark:bg-slate-900 border ${errors.title ? 'border-red-500 bg-red-50/10' : 'border-slate-200 dark:border-slate-800'} rounded-2xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none`"></textarea><div class="flex items-center justify-between px-1"><div v-if="errors.title" class="flex items-center gap-1.5"><p class="text-[10px] font-black text-red-500 uppercase tracking-widest">{{ errors.title }}</p></div><p :class="`text-[10px] font-black uppercase tracking-widest ml-auto ${form.title.length > 200 ? 'text-red-500' : 'text-slate-400'}`">{{ $t('upload.form.char_limit', { count: form.title.length }) }}</p></div></div></div>
            <div class="grid grid-cols-1 md:grid-cols-3 items-start gap-6"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300 mt-3">{{ $t('upload.form.category') }}</label><div class="md:col-span-2 space-y-2"><div class="relative"><select v-model="form.category" :class="`w-full pl-5 pr-12 py-3.5 bg-white dark:bg-slate-900 border ${errors.category ? 'border-red-500 bg-red-50/10' : 'border-slate-200 dark:border-slate-800'} rounded-2xl text-sm font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all cursor-pointer`"><option value="" disabled selected>{{ $t('upload.form.category_placeholder') }}</option><option>Confidential</option><option>Internal Use</option><option>Public</option></select><LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" /></div></div></div>
            <div class="grid grid-cols-1 md:grid-cols-3 items-start gap-6"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300 mt-3">{{ $t('upload.form.department') }}</label><div class="md:col-span-2 space-y-2"><div class="flex items-center justify-between px-5 py-3.5 bg-slate-50 dark:bg-slate-950/50 border border-slate-100 dark:border-slate-800 rounded-2xl"><span class="text-sm font-bold text-slate-500">Finance</span><LucideCheckCircle2 class="w-4 h-4 text-green-500" /></div><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest italic px-1">{{ $t('upload.form.department_hint') }}</p></div></div>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-8"><div class="space-y-3"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300">{{ $t('upload.form.count') }}</label><div class="flex items-center border border-slate-200 dark:border-slate-800 rounded-2xl overflow-hidden bg-white dark:bg-slate-900 transition-all"><button @click="form.count > 1 && form.count--" class="p-3.5 text-slate-400 hover:text-[#1E3A5F] transition-colors border-r border-slate-100 dark:border-slate-800"><LucideMinus class="w-4 h-4" /></button><input type="text" v-model.number="form.count" class="w-full text-center text-sm font-black bg-transparent outline-none" /><button @click="form.count++" class="p-3.5 text-slate-400 hover:text-[#1E3A5F] transition-colors border-l border-slate-100 dark:border-slate-800"><LucidePlus class="w-4 h-4" /></button></div></div><div class="md:col-span-2 space-y-3"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300">{{ $t('upload.form.date') }}</label><div class="relative"><input type="text" v-model="form.date" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"/><LucideCalendar class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" /></div></div></div>
            <div class="grid grid-cols-1 md:grid-cols-3 items-start gap-6"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300 mt-3">{{ $t('upload.form.notes') }}<span v-if="form.urgency === 'Urgent'" class="text-orange-500 text-[10px] block mt-0.5">(Justification required)</span></label><div class="md:col-span-2 space-y-2"><textarea v-model="form.notes" rows="4" :placeholder="form.urgency === 'Urgent' ? $t('upload.form.notes_urgent_placeholder') : $t('upload.form.notes_placeholder')" :class="`w-full px-5 py-4 bg-white dark:bg-slate-900 border ${errors.notes ? 'border-orange-500 bg-orange-50/10' : 'border-slate-200 dark:border-slate-800'} rounded-3xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none`"></textarea><p v-if="errors.notes" class="text-[10px] font-black text-orange-600 uppercase tracking-widest px-1">{{ errors.notes }}</p></div></div>
            <div class="space-y-6"><label class="text-sm font-black text-[#1E3A5F] dark:text-slate-300">{{ $t('upload.form.urgency') }}</label><div class="flex flex-wrap gap-8 px-1"><label v-for="level in [{key: 'Normal', label: $t('upload.form.urgency_levels.normal')}, {key: 'High', label: $t('upload.form.urgency_levels.high')}, {key: 'Urgent', label: $t('upload.form.urgency_levels.urgent')}]" :key="level.key" class="flex items-center gap-3 cursor-pointer group"><div class="relative flex items-center justify-center"><input type="radio" v-model="form.urgency" name="urgency" :value="level.key" class="peer appearance-none w-5 h-5 border-2 border-slate-200 dark:border-slate-700 rounded-full checked:border-primary-500 transition-all" /><div class="absolute w-2.5 h-2.5 rounded-full bg-primary-500 transform scale-0 peer-checked:scale-100 transition-transform"></div></div><span :class="`text-sm font-bold transition-colors ${form.urgency === level.key ? 'text-slate-900 dark:text-white' : 'text-slate-400 group-hover:text-slate-600'}`">{{ level.label }}</span></label></div><Transition enter-active-class="transition duration-300 ease-out" enter-from-class="transform -translate-y-4 opacity-0" enter-to-class="transform translate-y-0 opacity-100"><div v-if="form.urgency === 'Urgent'" class="flex items-center gap-4 p-4 bg-orange-50 dark:bg-orange-950/20 border border-orange-100 dark:border-orange-900/50 rounded-2xl"><LucideAlertTriangle class="w-5 h-5 text-orange-500 flex-shrink-0" /><p class="text-xs font-bold text-orange-700 dark:text-orange-400">{{ $t('upload.form.urgent_warning') }}</p></div></Transition></div>
          </div>
        </div>

        <!-- Tab 2: Bulk Upload (Excel) -->
        <div v-else class="glass rounded-3xl overflow-hidden shadow-xl shadow-slate-200/50 dark:shadow-none p-10 space-y-10" v-motion-slide-visible-bottom>
          <div class="space-y-2"><h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white">{{ $t('upload.bulk.title') }}</h2><p class="text-slate-500 font-medium">{{ $t('upload.bulk.subtitle') }}</p></div>
          <div class="space-y-6"><div class="flex items-center gap-3"><div class="w-7 h-7 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center text-xs font-black">1</div><h3 class="font-black text-[#1E3A5F] dark:text-white">{{ $t('upload.bulk.step1_title') }}</h3></div><div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 rounded-3xl space-y-6"><div class="flex items-center gap-5"><div class="w-16 h-16 rounded-2xl bg-green-50 dark:bg-green-900/20 flex items-center justify-center"><LucideFileSpreadsheet class="w-8 h-8 text-green-500" /></div><div><p class="text-lg font-black text-[#1E3A5F] dark:text-white">{{ $t('upload.bulk.step1_filename') }}</p><p class="text-xs font-bold text-slate-400 mt-1">{{ $t('upload.bulk.step1_version') }}</p></div></div><button class="flex items-center justify-center gap-3 w-full py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-sm font-black transition-all shadow-lg shadow-blue-900/20"><LucideDownload class="w-5 h-5" />{{ $t('upload.bulk.step1_download') }}</button><div class="flex items-start gap-3 p-4 bg-blue-50/30 dark:bg-blue-900/10 rounded-2xl border border-blue-100/50 dark:border-blue-800/30"><LucideInfo class="w-4 h-4 text-blue-500 mt-0.5 flex-shrink-0" /><p class="text-[11px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed">{{ $t('upload.bulk.step1_info') }}</p></div></div></div>
          <div class="space-y-6"><div class="flex items-center gap-3"><div class="w-7 h-7 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center text-xs font-black">2</div><h3 class="font-black text-[#1E3A5F] dark:text-white">{{ $t('upload.bulk.step2_title') }}</h3></div><div class="group relative border-2 border-dashed border-slate-200 dark:border-slate-800 rounded-[2.5rem] p-12 flex flex-col items-center text-center hover:border-primary-500/50 hover:bg-primary-50/5 dark:hover:bg-primary-950/5 transition-all cursor-pointer"><input type="file" class="absolute inset-0 opacity-0 cursor-pointer" accept=".xlsx,.xls" /><div class="w-16 h-16 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center mb-6 group-hover:scale-110 transition-transform"><LucideUploadCloud class="w-8 h-8 text-slate-400 group-hover:text-primary-500" /></div><h4 class="text-lg font-black text-[#1E3A5F] dark:text-white mb-2">{{ $t('upload.bulk.step2_dropzone') }}</h4><p class="text-sm font-bold text-slate-400 mb-8">{{ $t('upload.bulk.step2_formats') }}</p><button class="px-10 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-xs font-black uppercase tracking-widest text-[#1E3A5F] dark:text-white hover:bg-slate-50 transition-all shadow-sm">{{ $t('upload.bulk.step2_browse') }}</button></div></div>
          <div class="p-6 bg-orange-50/50 dark:bg-orange-950/10 border border-orange-100 dark:border-orange-900/30 rounded-3xl space-y-4"><div class="flex items-center gap-2 text-orange-600 dark:text-orange-400"><LucideAlertTriangle class="w-5 h-5" /><span class="text-xs font-black uppercase tracking-widest">{{ $t('upload.bulk.validation_rules_title') }}</span></div><ul class="space-y-2.5"><li v-for="rule in $tm('upload.bulk.rules')" :key="rule" class="flex items-center gap-3"><div class="w-1 h-1 rounded-full bg-orange-300"></div><p class="text-xs font-bold text-orange-800/70 dark:text-orange-400/70">{{ $rt(rule) }}</p></li></ul></div>
        </div>
      </div>

      <!-- Right Panel -->
      <div class="space-y-8">
        <div class="glass p-8 rounded-3xl" v-motion-slide-visible-bottom><div class="flex items-center gap-3 mb-8"><div class="w-8 h-8 rounded-lg bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center text-blue-500"><LucideHelpCircle class="w-5 h-5" /></div><h3 class="font-black text-slate-800 dark:text-white">{{ $t('upload.side.next_title') }}</h3></div><div class="space-y-8"><div v-for="(step, i) in $tm('upload.side.steps')" :key="i" class="flex gap-6 relative"><div v-if="i !== 3" class="absolute left-3 top-8 w-0.5 h-10 bg-slate-100 dark:bg-slate-800"></div><div :class="`w-6 h-6 rounded-full flex-shrink-0 flex items-center justify-center border-2 z-10 ${i === 0 ? 'bg-[#1E3A5F] border-[#1E3A5F] text-white' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-300'}`"><LucideCheck v-if="i === 0" class="w-3 h-3" /><span v-else class="text-[10px] font-black">{{ i + 1 }}</span></div><div><p :class="`text-sm font-black ${i === 0 ? 'text-slate-800 dark:text-white' : 'text-slate-400'}`">{{ $rt(step.title) }}</p><p class="text-[11px] font-bold text-slate-400 mt-1 leading-relaxed">{{ $rt(step.desc) }}</p></div></div></div></div>
        <div class="glass p-8 rounded-3xl border-l-4 border-primary-500/30" v-motion-slide-visible-bottom><div class="flex items-center gap-3 mb-8"><div class="w-8 h-8 rounded-lg bg-primary-50 dark:bg-primary-900/20 flex items-center justify-center text-primary-500"><LucideLightbulb class="w-5 h-5" /></div><h3 class="font-black text-slate-800 dark:text-white">{{ $t('upload.side.tips_title') }}</h3></div><ul class="space-y-4"><li v-for="tip in $tm('upload.side.tips')" :key="tip" class="flex items-start gap-3"><LucideCheck class="w-4 h-4 text-primary-500 mt-0.5 flex-shrink-0" /><p class="text-xs font-bold text-slate-500 dark:text-slate-400 leading-relaxed">{{ $rt(tip) }}</p></li></ul></div>
        <div class="glass p-8 rounded-3xl" v-motion-slide-visible-bottom><div class="flex items-center gap-3 mb-8"><div class="w-8 h-8 rounded-lg bg-slate-50 dark:bg-slate-900/20 flex items-center justify-center text-slate-500"><LucideHistory class="w-5 h-5" /></div><h3 class="font-black text-slate-800 dark:text-white">{{ $t('upload.side.recent_title') }}</h3></div><div class="space-y-6"><div v-for="sub in recents" :key="sub.id" class="flex items-center justify-between group"><div><p class="text-sm font-black text-slate-700 dark:text-slate-200 group-hover:text-primary-500 transition-colors">{{ sub.id }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">{{ sub.time }}</p></div><span :class="`px-2.5 py-1 rounded-md text-[9px] font-black uppercase tracking-tighter ${sub.statusBg}`">{{ sub.status }}</span></div></div><button class="w-full mt-10 py-3 border border-slate-100 dark:border-slate-800 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-400 hover:text-primary-500 hover:border-primary-500/30 transition-all">{{ $t('upload.side.view_history') }}</button></div>
      </div>
    </div>

    <!-- Manifest Preview View -->
    <div v-if="showManifestPreview" class="min-h-screen bg-slate-100/50 dark:bg-slate-950 p-8 space-y-8" v-motion-fade>
      <!-- Toolbar -->
      <div class="flex items-center justify-between bg-white dark:bg-slate-900 p-6 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-sm">
        <div class="flex items-center gap-4">
          <LucideFileText class="w-6 h-6 text-[#1E3A5F]" />
          <h1 class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('manifest.title') }}</h1>
          <span class="px-3 py-1 bg-slate-100 dark:bg-slate-800 rounded-full text-[10px] font-black text-slate-500">REG-2026-0234 to REG-2026-0275</span>
        </div>
        <div class="flex items-center gap-3 px-4 py-2 bg-green-50 dark:bg-green-900/20 rounded-full border border-green-100 dark:border-green-900/50">
          <LucideCheckCircle2 class="w-4 h-4 text-green-500" />
          <span class="text-xs font-black text-green-700 dark:text-green-400 uppercase tracking-tight">{{ $t('manifest.status_approved', { name: 'John Doe', date: '7 Mar 2026' }) }}</span>
        </div>
      </div>

      <!-- Paper Container -->
      <div class="max-w-4xl mx-auto bg-white dark:bg-slate-900 shadow-2xl p-16 relative min-h-[1120px] flex flex-col">
        <!-- Header -->
        <div class="flex justify-between items-start mb-12">
          <div class="space-y-1">
            <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-widest">{{ $t('manifest.header_title') }}</h2>
            <p class="text-lg font-bold text-slate-800 dark:text-slate-200">PT. AKIRADATA</p>
            <p class="text-xs font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('manifest.header_subtitle') }}</p>
          </div>
          <div class="w-16 h-16 bg-slate-50 dark:bg-slate-800 rounded-lg flex items-center justify-center text-slate-300 font-black italic">Logo</div>
        </div>

        <div class="w-full h-px bg-slate-200 dark:bg-slate-800 mb-12"></div>

        <!-- Metadata Section -->
        <div class="grid grid-cols-2 gap-12 mb-16">
          <div class="grid grid-cols-2 gap-y-6">
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('manifest.manifest_id') }}</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">MNF-2026-0050</p>
            </div>
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('manifest.generated_date') }}</p>
              <p class="text-sm font-bold text-slate-700 dark:text-slate-300">07 March 2026, 14:22</p>
            </div>
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('manifest.submitter') }}</p>
              <p class="text-sm font-bold text-slate-700 dark:text-slate-300">Budi Qartono</p>
            </div>
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('manifest.approver') }}</p>
              <p class="text-sm font-bold text-slate-700 dark:text-slate-300">John Doe</p>
            </div>
            <div class="space-y-1 pt-4 col-span-2">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('manifest.total_docs') }}</p>
              <p class="text-2xl font-black text-[#1E3A5F] dark:text-white">{{ $t('manifest.files', { count: 42 }) }}</p>
            </div>
          </div>
          <div class="flex flex-col items-center justify-center">
            <div class="w-32 h-32 border-2 border-slate-100 dark:border-slate-800 rounded-2xl flex items-center justify-center mb-2">
              <LucideQrCode class="w-24 h-24 text-slate-800 dark:text-white" />
            </div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">MNF-2026-0050</p>
          </div>
        </div>

        <!-- Listing Table -->
        <div class="space-y-4 mb-16 flex-grow">
          <div class="flex items-center gap-3">
            <div class="w-1 h-4 bg-[#1E3A5F]"></div>
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('manifest.listing_title') }}</h3>
          </div>
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="bg-slate-50 dark:bg-slate-800 border-y border-slate-200 dark:border-slate-700">
                <th class="p-3 text-[9px] font-black text-slate-400 uppercase tracking-tight">{{ $t('manifest.table.no') }}</th>
                <th class="p-3 text-[9px] font-black text-slate-400 uppercase tracking-tight">{{ $t('manifest.table.reg_id') }}</th>
                <th class="p-3 text-[9px] font-black text-slate-400 uppercase tracking-tight">{{ $t('manifest.table.title') }}</th>
                <th class="p-3 text-[9px] font-black text-slate-400 uppercase tracking-tight">{{ $t('manifest.table.category') }}</th>
                <th class="p-3 text-[9px] font-black text-slate-400 uppercase tracking-tight">{{ $t('manifest.table.qty') }}</th>
                <th class="p-3 text-[9px] font-black text-slate-400 uppercase tracking-tight">{{ $t('manifest.table.urgency') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
              <tr><td class="p-3 text-[10px] font-bold">1</td><td class="p-3 text-[10px] font-black text-[#1E3A5F]">REG-2026-0234</td><td class="p-3 text-[10px] font-bold">Q1 Financial Report Final</td><td class="p-3 text-[10px] font-bold">Finance</td><td class="p-3 text-[10px] font-bold">2</td><td class="p-3 text-[10px] font-bold">Normal</td></tr>
              <tr class="bg-red-50/30"><td class="p-3 text-[10px] font-bold text-red-500">2</td><td class="p-3 text-[10px] font-black text-red-500">REG-2026-0235</td><td class="p-3 text-[10px] font-black text-red-500">Confidential Merger Agreement</td><td class="p-3 text-[10px] font-black text-red-500">Legal</td><td class="p-3 text-[10px] font-black text-red-500">1</td><td class="p-3 text-[10px] font-black text-red-500">High Urgency</td></tr>
              <tr><td class="p-3 text-[10px] font-bold">3</td><td class="p-3 text-[10px] font-black text-[#1E3A5F]">REG-2026-0236</td><td class="p-3 text-[10px] font-bold">Employee Payroll Summary Feb</td><td class="p-3 text-[10px] font-bold">HR</td><td class="p-3 text-[10px] font-bold">4</td><td class="p-3 text-[10px] font-bold">Normal</td></tr>
              <tr><td class="p-3 text-[10px] font-bold">...</td><td class="p-3 text-[10px] font-black text-[#1E3A5F]">...</td><td class="p-3 text-[10px] font-bold">...</td><td class="p-3 text-[10px] font-bold">...</td><td class="p-3 text-[10px] font-bold">...</td><td class="p-3 text-[10px] font-bold">...</td></tr>
              <tr><td class="p-3 text-[10px] font-bold">42</td><td class="p-3 text-[10px] font-black text-[#1E3A5F]">REG-2026-0275</td><td class="p-3 text-[10px] font-bold">IT Asset Disposal Log</td><td class="p-3 text-[10px] font-bold">Operations</td><td class="p-3 text-[10px] font-bold">1</td><td class="p-3 text-[10px] font-bold">Normal</td></tr>
            </tbody>
          </table>
        </div>

        <!-- Delivery Instructions -->
        <div class="space-y-6">
          <h3 class="text-xs font-black text-slate-800 dark:text-white uppercase tracking-widest">{{ $t('manifest.delivery_title') }}</h3>
          <p class="text-[11px] font-medium text-slate-500 leading-relaxed">{{ $t('manifest.delivery_text') }}</p>
          <div class="p-4 bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100 dark:border-blue-800 rounded-xl">
            <p class="text-[11px] font-black text-[#1E3A5F] dark:text-blue-400">{{ $t('manifest.delivery_location') }}</p>
          </div>
          <p class="text-[11px] font-medium text-slate-500 italic">{{ $t('manifest.delivery_qr_hint') }}</p>
          <div class="p-4 bg-orange-50/50 dark:bg-orange-900/10 border border-orange-100 dark:border-orange-800 rounded-xl">
            <p class="text-[11px] font-black text-orange-600 dark:text-orange-400 leading-relaxed">{{ $t('manifest.warning') }}</p>
          </div>
        </div>

        <!-- Paper Footer -->
        <div class="mt-20 flex justify-between items-end border-t border-slate-100 pt-6">
          <div class="space-y-1">
            <p class="text-[9px] font-bold text-slate-400">{{ $t('manifest.footer_info') }}</p>
            <p class="text-[9px] font-bold text-slate-400">{{ $t('manifest.verified_on', { date: '07/03/2026 14:22:10' }) }}</p>
          </div>
          <p class="text-[9px] font-bold text-slate-400 uppercase">{{ $t('manifest.page', { current: 1, total: 3 }) }}</p>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('manifest.version') }}</p>
        </div>
      </div>

      <!-- Bottom Action Bar -->
      <div class="fixed bottom-0 left-0 right-0 p-8 glass border-t border-slate-200/50 flex justify-center gap-4">
        <button @click="showManifestPreview = false" class="px-10 py-4 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all">{{ $t('manifest.btn_back') }}</button>
        <button class="px-12 py-4 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl hover:bg-[#152943] transition-all">{{ $t('manifest.btn_print') }}</button>
        <button class="px-10 py-4 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-[#1E3A5F] hover:bg-slate-50 transition-all">{{ $t('manifest.btn_download') }}</button>
      </div>
    </div>

    <!-- Rest of the modals (Footer Bar, Progress, Success Modals) -->
    <div v-if="!showProgressModal && !showBulkSuccessModal && !showManifestPreview" class="fixed bottom-0 left-0 lg:left-64 right-0 p-6 glass border-t border-slate-200/50 dark:border-white/5 z-20" v-motion-slide-bottom>
      <div class="max-w-7xl mx-auto flex items-center justify-between">
        <div v-if="isBulkReviewMode" class="flex items-center gap-3"><LucideAlertTriangle class="w-5 h-5 text-red-500" /><p class="text-sm font-black text-slate-700 dark:text-slate-300"><span class="text-red-500">{{ $t('upload.bulk.review.footer.errors_found', { count: 3 }) }}</span></p></div>
        <button v-else @click="resetForm()" class="px-8 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 dark:text-slate-400 hover:bg-slate-50 transition-all">{{ $t('upload.form.save_draft') }}</button>
        <div class="flex items-center gap-4">
          <template v-if="isBulkReviewMode">
            <button @click="isBulkReviewMode = false" class="px-6 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 dark:text-slate-300 hover:bg-slate-100 transition-all">{{ $t('upload.bulk.review.footer.upload_different') }}</button>
            <button class="px-6 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest text-[#1E3A5F] dark:text-white hover:bg-slate-50 transition-all">{{ $t('upload.bulk.review.footer.revalidate') }}</button>
            <button disabled class="flex items-center gap-3 px-8 py-3.5 bg-slate-200 dark:bg-slate-800 text-slate-400 rounded-xl text-xs font-black uppercase tracking-widest cursor-not-allowed opacity-70"><LucideLock class="w-4 h-4" />{{ $t('upload.bulk.review.footer.submit_all', { count: 42 }) }}</button>
          </template>
          <template v-else>
            <button @click="navigateTo('/dashboard')" class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 mr-4">{{ $t('upload.form.cancel') }}</button>
            <button @click="handleSubmit" :disabled="submitting" class="flex items-center gap-3 px-10 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all group disabled:opacity-50"><span v-if="submitting">{{ $t('upload.form.processing') }}</span><span v-else>{{ $t('upload.form.submit') }}</span><LucideArrowRight v-if="!submitting" class="w-4 h-4 group-hover:translate-x-1 transition-transform" /></button>
          </template>
        </div>
      </div>
    </div>

    <!-- Success Modals with "Print Summary" link -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100">
      <div v-if="showBulkSuccessModal" class="fixed inset-0 z-[70] flex items-center justify-center p-4 bg-slate-950/60 backdrop-blur-md">
        <div class="bg-white dark:bg-slate-900 w-full max-w-2xl rounded-[2.5rem] shadow-2xl relative overflow-hidden flex flex-col max-h-[90vh]" v-motion-pop>
          <div class="p-8 sm:p-12 overflow-y-auto custom-scrollbar flex flex-col items-center text-center">
            <div class="w-16 h-16 rounded-full bg-green-500 flex items-center justify-center text-white mb-6 shadow-xl"><LucideCheck class="w-8 h-8" /></div>
            <p class="text-xs font-black text-green-500 uppercase tracking-[0.2em] mb-2">{{ $t('upload.bulk.success.status') }}</p>
            <h3 class="text-3xl font-black text-[#1E3A5F] dark:text-white mb-8">{{ $t('upload.bulk.success.title') }}</h3>
            <div class="w-full p-6 bg-blue-50/30 dark:bg-blue-900/10 rounded-3xl border border-blue-100 mb-6 flex gap-5">
              <div class="w-12 h-12 rounded-2xl bg-white flex items-center justify-center flex-shrink-0"><LucideFileText class="w-6 h-6 text-[#1E3A5F]" /></div>
              <div class="text-left"><p class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest mb-1">{{ $t('upload.bulk.success.summary_title') }}</p><p class="text-xs font-medium text-slate-500 leading-relaxed">{{ $t('upload.bulk.success.summary_text', { count: 42 }) }}</p></div>
            </div>
            <div class="w-full p-8 border border-slate-100 rounded-3xl mb-6">
              <div class="flex items-center justify-between mb-6">
                <div class="flex items-center gap-3"><LucideLayoutGrid class="w-5 h-5 text-slate-400" /><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('upload.bulk.success.ids_title') }}</p></div>
                <div class="flex gap-2"><button class="px-4 py-2 bg-slate-50 rounded-lg text-[9px] font-black uppercase tracking-widest text-slate-600 flex items-center gap-2"><LucideCopy class="w-3 h-3" /> {{ $t('upload.bulk.success.copy_all') }}</button><button class="px-4 py-2 bg-slate-50 rounded-lg text-[9px] font-black uppercase tracking-widest text-slate-600 flex items-center gap-2"><LucideDownload class="w-3 h-3" /> {{ $t('upload.bulk.success.download_csv') }}</button></div>
              </div>
              <div class="grid grid-cols-3 gap-3"><div v-for="i in 12" :key="i" class="py-3 px-4 bg-white border border-slate-100 rounded-xl text-[10px] font-bold text-slate-500">REG-2026-02{{ 33 + i }}</div></div>
            </div>
            <div class="w-full p-6 bg-orange-50/30 rounded-3xl border border-orange-100 flex gap-5">
              <div class="w-12 h-12 rounded-2xl bg-white flex items-center justify-center flex-shrink-0"><LucideClipboardCheck class="w-6 h-6 text-orange-500" /></div>
              <div class="text-left"><p class="text-[10px] font-black text-orange-600 uppercase tracking-widest mb-1">{{ $t('upload.bulk.success.next_steps_title') }}</p><p class="text-xs font-medium text-slate-500 leading-relaxed">{{ $t('upload.bulk.success.next_steps_text') }}</p></div>
            </div>
          </div>
          <div class="p-8 border-t border-slate-100 bg-slate-50/50 flex items-center justify-between">
            <button @click="resetForm(); showBulkSuccessModal = false" class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-[#1E3A5F]">{{ $t('upload.bulk.success.submit_another') }}</button>
            <div class="flex gap-4">
              <button @click="showBulkSuccessModal = false; showManifestPreview = true" class="px-8 py-3.5 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-[#1E3A5F] hover:bg-slate-50 flex items-center gap-2"><LucidePrinter class="w-4 h-4" /> {{ $t('upload.success_modal.print') }}</button>
              <button @click="resetForm(); showBulkSuccessModal = false" class="px-10 py-3.5 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl">{{ $t('upload.bulk.success.close') }}</button>
            </div>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Progress Modal (Re-using existing code) -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100"><div v-if="showProgressModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-slate-950/60 backdrop-blur-md"><div class="bg-white dark:bg-slate-900 w-full max-w-md rounded-[2.5rem] shadow-2xl p-10 flex flex-col items-center text-center"><div class="w-16 h-16 rounded-full bg-blue-50 flex items-center justify-center mb-6"><LucideRefreshCcw class="w-8 h-8 text-[#1E3A5F] animate-spin" /></div><h3 class="text-lg font-black text-[#1E3A5F] uppercase tracking-widest mb-2">{{ $t('upload.bulk.progress.title') }}</h3><p class="text-slate-500 text-sm font-medium mb-10">{{ $t('upload.bulk.progress.subtitle') }}</p><div class="w-full space-y-4 mb-10"><div class="flex items-end justify-between mb-2"><span class="text-4xl font-black text-[#1E3A5F]">{{ uploadProgress }}%</span><span class="text-xs font-black text-blue-500 uppercase tracking-widest">{{ $t('upload.bulk.progress.action') }}</span></div><div class="w-full h-3 bg-slate-100 rounded-full overflow-hidden"><div :style="{ width: `${uploadProgress}%` }" class="h-full bg-[#1E3A5F] transition-all duration-300 ease-out"></div></div></div><button @click="showProgressModal = false; uploadProgress = 0" class="px-10 py-3 bg-slate-50 rounded-xl text-xs font-black uppercase tracking-widest text-slate-500">{{ $t('upload.bulk.progress.cancel') }}</button></div></div></Transition>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import { 
  LucideFileUp, LucidePenTool, LucideFileSpreadsheet, LucideChevronDown, LucideCheckCircle2, LucideMinus, LucidePlus, LucideCalendar, LucideHelpCircle, LucideCheck, LucideLightbulb, LucideHistory, LucideArrowRight, LucideAlertCircle, LucideAlertTriangle, LucideDownload, LucideInfo, LucideUploadCloud, LucideX, LucideCopy, LucideLayoutGrid, LucidePrinter, LucideXCircle, LucidePencil, LucideTrash2, LucideRefreshCcw, LucideFileText, LucideClock, LucideLock, LucideClipboardCheck, LucideQrCode
} from 'lucide-vue-next'

const { t, tm, rt } = useI18n()

const activeTab = ref('manual')
const submitting = ref(false)
const showManualSuccessModal = ref(false)
const showBulkSuccessModal = ref(false)
const showManifestPreview = ref(false)
const isBulkReviewMode = ref(false)
const showProgressModal = ref(false)
const uploadProgress = ref(0)

const formatDate = (date) => {
  const d = new Date(date);
  const day = String(d.getDate()).padStart(2, '0');
  const month = String(d.getMonth() + 1).padStart(2, '0');
  const year = d.getFullYear();
  return `${day}.${month}.${year}`;
}

const form = reactive({ type: '', title: '', category: '', count: 1, date: formatDate(new Date()), notes: '', urgency: 'Normal' })
const errors = reactive({ type: '', title: '', category: '', count: '', date: '', notes: '' })

const bulkData = [
  { type: 'Invoice', subject: 'INV-2023-001-A', category: 'Finance', date: '12.10.2023', hasError: false },
  { type: 'Contract', subject: '', category: 'Legal', date: '15.10.2023', hasError: true, errorMsg: 'Validation Error: Missing Subject.' },
  { type: 'Memorandum', subject: 'MEMO-Q4-PLAN', category: 'Internal', date: '14.10.2023', hasError: false },
  { type: 'Tax Form', subject: 'TAX-2022-W9', category: 'Unknown', date: '10.10.2023', hasError: true, errorMsg: 'Validation Error: Invalid Category.' },
]

const validate = () => {
  if (activeTab.value === 'bulk') return true
  let isValid = true
  if (!form.type) { errors.type = 'Required'; isValid = false }
  if (!form.title) { errors.title = 'Required'; isValid = false }
  return isValid
}

const handleSubmit = () => {
  if (validate()) {
    if (activeTab.value === 'bulk' && !isBulkReviewMode.value) { isBulkReviewMode.value = true; return }
    if (isBulkReviewMode.value) { startBulkUpload(); return }
    submitting.value = true
    setTimeout(() => { submitting.value = false; showBulkSuccessModal.value = true }, 1200)
  }
}

const startBulkUpload = () => {
  showProgressModal.value = true
  uploadProgress.value = 0
  const interval = setInterval(() => {
    uploadProgress.value += 10
    if (uploadProgress.value >= 100) {
      clearInterval(interval)
      setTimeout(() => { showProgressModal.value = false; showBulkSuccessModal.value = true }, 500)
    }
  }, 300)
}

const resetForm = () => { 
  Object.assign(form, { type: '', title: '', category: '', count: 1, date: formatDate(new Date()), notes: '', urgency: 'Normal' })
  isBulkReviewMode.value = false
  showManifestPreview.value = false
}

const recents = [{ id: 'INV-2023-0045', time: '2h ago', status: 'Approved', statusBg: 'bg-green-50 text-green-500' }, { id: 'PO-AKR-009', time: 'Yesterday', status: 'Pending', statusBg: 'bg-orange-50 text-orange-500' }, { id: 'Travel_Budi', time: 'Oct 25', status: 'Draft', statusBg: 'bg-slate-100 text-slate-500' }]
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 6px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>
