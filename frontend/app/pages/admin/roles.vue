<template>
  <div class="h-[calc(100vh-120px)] flex flex-col -m-10 overflow-hidden" v-motion-fade>
    <!-- Warning Alert Bar -->
    <div class="bg-amber-100 dark:bg-amber-900/30 border-b border-amber-200 dark:border-amber-800 p-4 px-10 flex flex-col md:flex-row md:items-center justify-between gap-4 z-20">
      <div class="flex items-center gap-3">
        <LucideAlertTriangle class="w-5 h-5 text-amber-600" />
        <p class="text-[11px] font-bold text-amber-800 dark:text-amber-200 uppercase tracking-tight">{{ $t('admin.roles.alert.msg', { name: selectedRole?.name || 'Admin' }) }}</p>
      </div>
      <div class="flex items-center gap-6">
        <button class="text-[10px] font-black text-amber-600 dark:text-amber-400 uppercase tracking-widest hover:underline">{{ $t('admin.roles.alert.btn_cancel') }}</button>
        <button class="px-6 py-2 bg-amber-600 text-white rounded-lg text-[10px] font-black uppercase tracking-widest hover:bg-amber-700 transition-all shadow-lg shadow-amber-600/20">{{ $t('admin.roles.alert.btn_review') }}</button>
      </div>
    </div>

    <div class="flex flex-grow overflow-hidden">
      <!-- Left Side: Role Hierarchy -->
      <div class="w-80 bg-slate-50 dark:bg-slate-950 border-r border-slate-200 dark:border-slate-800 flex flex-col">
        <div class="p-8 space-y-6">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.hierarchy.title') }}</h4>
          <div class="relative">
            <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
            <input type="text" :placeholder="$t('admin.roles.hierarchy.filter')" class="w-full pl-12 pr-6 py-3 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" />
          </div>
        </div>
        <div class="flex-grow overflow-y-auto px-4 space-y-2 custom-scrollbar">
          <div v-for="role in roles" :key="role.id" :class="`p-4 px-6 rounded-2xl flex items-center justify-between cursor-pointer transition-all group ${selectedRole?.id === role.id ? 'bg-[#1E3A5F] text-white shadow-xl shadow-blue-900/20' : 'text-slate-500 hover:bg-white dark:hover:bg-slate-900'}`" @click="selectedRole = role">
            <div class="flex items-center gap-4">
              <component :is="role.icon" class="w-5 h-5" />
              <span :class="`text-xs font-black uppercase tracking-tight ${role.disabled ? 'opacity-50' : ''}`">{{ role.name }} <span v-if="role.disabled" class="text-[8px] font-bold">(Disabled)</span></span>
            </div>
            <div v-if="role.disabled" class="w-4 h-4 text-slate-300"><LucideLock class="w-full h-full" /></div>
            <div v-else :class="`w-4 h-4 rounded-full border-2 transition-colors ${selectedRole?.id === role.id ? 'bg-white border-white' : 'border-slate-200 group-hover:border-slate-300'}`"><div v-if="selectedRole?.id === role.id" class="w-1.5 h-1.5 bg-[#1E3A5F] rounded-full m-0.5"></div></div>
          </div>
        </div>
      </div>

      <!-- Center: Permission Matrix -->
      <div class="flex-grow bg-white dark:bg-slate-900 flex flex-col">
        <div class="p-8 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between">
          <h2 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.roles.matrix.title', { name: selectedRole?.name || 'Admin' }) }}</h2>
          <div class="flex items-center gap-4 cursor-pointer" @click="fullAccess = !fullAccess">
            <div :class="`w-10 h-5 rounded-full relative transition-all ${fullAccess ? 'bg-[#1E3A5F]' : 'bg-slate-200 dark:bg-slate-700'}`">
              <div :class="`absolute top-0.5 w-4 h-4 bg-white rounded-full transition-all ${fullAccess ? 'left-5.5' : 'left-0.5'}`"></div>
            </div>
            <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.matrix.full_access') }}</span>
          </div>
        </div>

        <div class="flex-grow overflow-auto custom-scrollbar">
          <table class="w-full text-left border-collapse">
            <thead class="sticky top-0 bg-slate-50/80 dark:bg-slate-800/80 backdrop-blur-md z-10">
              <tr class="text-[9px] font-black text-slate-400 uppercase tracking-widest">
                <th class="p-6 pl-10 border-b border-slate-100 dark:border-slate-800">{{ $t('admin.roles.matrix.headers.modul') }}</th>
                <th v-for="h in ['view', 'pre', 'dl', 'prn', 'up', 'ed', 'del', 'app', 'arc']" :key="h" class="p-6 text-center border-b border-slate-100 dark:border-slate-800">{{ $t(`admin.roles.matrix.headers.${h}`) }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <template v-for="section in matrixSections" :key="section.id">
                <tr class="bg-slate-50/30 dark:bg-slate-800/10">
                  <td colspan="10" class="p-4 pl-10 text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">{{ $t(`admin.roles.matrix.sections.${section.id}`) }}</td>
                </tr>
                <tr v-for="mod in section.modules" :key="mod" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-all group">
                  <td class="p-6 pl-10 text-xs font-black text-slate-600 dark:text-slate-300 uppercase tracking-tight">{{ $t(`admin.roles.matrix.modules.${mod}`) }}</td>
                  <td v-for="h in ['view', 'pre', 'dl', 'prn', 'up', 'ed', 'del', 'app', 'arc']" :key="h" class="p-6 text-center">
                    <div class="flex items-center justify-center">
                      <div v-if="permissions[mod]?.[h] !== undefined" :class="`w-5 h-5 rounded border-2 flex items-center justify-center transition-all cursor-pointer ${permissions[mod][h] ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200 group-hover:border-slate-300'}`" @click="permissions[mod][h] = !permissions[mod][h]">
                        <LucideCheck v-if="permissions[mod][h]" class="w-3.5 h-3.5 text-white" />
                      </div>
                      <div v-else class="w-5 h-0.5 bg-slate-100 dark:bg-slate-800 rounded-full opacity-30"></div>
                    </div>
                  </td>
                </tr>
              </template>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Side: Role Details -->
      <div class="w-96 bg-slate-50 dark:bg-slate-950 border-l border-slate-200 dark:border-slate-800 flex flex-col">
        <div class="p-10 space-y-12 overflow-y-auto custom-scrollbar">
          <div class="flex items-center gap-6">
            <div class="w-16 h-16 rounded-2xl bg-blue-50 text-[#1E3A5F] flex items-center justify-center shadow-lg shadow-blue-900/10"><LucideShieldCheck class="w-8 h-8" /></div>
            <div class="space-y-0.5">
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase leading-none tracking-tighter">{{ selectedRole?.name || 'Admin' }}</h3>
              <p class="text-[10px] font-black text-primary-500 uppercase tracking-widest">{{ $t('admin.roles.details.authority') }}</p>
            </div>
          </div>

          <div class="space-y-6">
            <p class="text-xs font-bold text-slate-500 leading-relaxed">{{ $t('admin.roles.details.desc') }}</p>
            
            <div class="p-8 bg-white dark:bg-slate-900 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 space-y-6 shadow-sm">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.details.risk.title') }}</h4>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div class="h-2 flex-grow bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                    <div class="h-full bg-red-500" style="width: 100%"></div>
                  </div>
                  <span class="ml-4 text-[10px] font-black text-red-500 uppercase">{{ $t('admin.roles.details.risk.val') }}</span>
                </div>
                <p class="text-[9px] font-bold text-slate-400 italic leading-relaxed">{{ $t('admin.roles.details.risk.desc') }}</p>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-6">
            <div class="p-6 bg-white dark:bg-slate-900 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-2 shadow-sm">
              <p class="text-[18px] font-black text-[#1E3A5F] dark:text-white tracking-tighter">124</p>
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.details.stats.active_perms') }}</p>
            </div>
            <div class="p-6 bg-white dark:bg-slate-900 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-2 shadow-sm">
              <p class="text-[18px] font-black text-[#1E3A5F] dark:text-white tracking-tighter">05</p>
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.details.stats.active_users') }}</p>
            </div>
          </div>

          <div class="space-y-6">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.details.restricted.title') }}</h4>
            <div class="space-y-4">
              <div v-for="i in [1, 2]" :key="i" class="flex items-center gap-4 opacity-50 grayscale group cursor-not-allowed">
                <div class="w-8 h-8 rounded-full border-2 border-red-200 flex items-center justify-center text-red-400 group-hover:bg-red-50 transition-all"><LucideBan class="w-4 h-4" /></div>
                <span class="text-[11px] font-bold text-slate-500 uppercase">{{ $t(`admin.roles.details.restricted.item_${i}`) }}</span>
              </div>
            </div>
          </div>

          <div class="pt-10 border-t border-slate-100 dark:border-slate-800 space-y-6">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.roles.details.history.title') }}</h4>
            <div class="flex items-center gap-5 p-5 bg-white dark:bg-slate-900 rounded-2xl border border-slate-50 dark:border-slate-800">
              <div class="w-10 h-10 rounded-xl bg-slate-900 text-white flex items-center justify-center text-sm font-black">RH</div>
              <div class="space-y-0.5">
                <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">R. Hartono</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase">12 Okt 2023 • 14:20 WIB</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { 
  LucideAlertTriangle, LucideSearch, LucideShield, LucideUser, 
  LucideUsers, LucideBriefcase, LucideLock, LucideBuilding, 
  LucideCheck, LucideShieldCheck, LucideBan, LucideChevronDown
} from 'lucide-vue-next'

const selectedRole = ref(null)
const fullAccess = ref(false)

const roles = [
  { id: 1, name: 'Admin', icon: LucideShield },
  { id: 2, name: 'User Umum', icon: LucideUser },
  { id: 3, name: 'Section Head', icon: LucideBriefcase },
  { id: 4, name: 'Head', icon: LucideBriefcase },
  { id: 5, name: 'Group Head', icon: LucideUsers, disabled: true },
  { id: 6, name: 'Management', icon: LucideBuilding }
]

const matrixSections = [
  { id: 'search', modules: ['pencarian_global', 'navigasi_folder'] },
  { id: 'workflow', modules: ['antrian_tugas'] },
  { id: 'admin', modules: ['manajemen_user'] }
]

const permissions = reactive({
  pencarian_global: { view: true, pre: true, dl: true, prn: true },
  navigasi_folder: { view: true, pre: true, dl: true, prn: true, up: true, ed: true },
  antrian_tugas: { view: true, pre: true, dl: false, prn: true, app: true },
  manajemen_user: { view: true, up: true, ed: true, del: true }
})

onMounted(() => {
  selectedRole.value = roles[0]
})

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 4px; height: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>
