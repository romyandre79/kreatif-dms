<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('settings.sso.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('settings.sso.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all">
          <LucideEye class="w-4 h-4" />
          {{ $t('settings.sso.header.preview') }}
        </button>
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all">
          <LucideZap class="w-4 h-4" />
          {{ $t('settings.sso.header.test') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3">
          <LucideSave class="w-4 h-4" />
          {{ $t('settings.sso.header.save') }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Connection Settings -->
      <div class="lg:col-span-4 glass p-10 rounded-[3rem] space-y-10">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><LucideNetwork class="w-6 h-6" /></div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('settings.sso.connection.title') }}</h3>
          </div>
          <div class="flex flex-col items-end gap-1">
            <div class="w-12 h-6 bg-[#1E3A5F] rounded-full relative cursor-pointer" @click="ldapActive = !ldapActive">
              <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${ldapActive ? 'left-7' : 'left-1'}`"></div>
            </div>
            <span class="text-[7px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.active') }}</span>
          </div>
        </div>

        <div class="space-y-8">
          <div class="space-y-3">
            <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.provider.title') }}</label>
            <div class="flex bg-slate-100 dark:bg-slate-800 p-1 rounded-xl">
              <button v-for="p in ['ad', 'openldap']" :key="p" :class="`flex-grow py-3 rounded-lg text-[9px] font-black uppercase tracking-widest transition-all ${provider === p ? 'bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white shadow-sm border border-slate-100 dark:border-slate-800' : 'text-slate-400'}`" @click="provider = p">
                <LucideServer class="w-3.5 h-3.5 inline-block mr-2" />
                {{ $t(`settings.sso.connection.provider.${p}`) }}
              </button>
            </div>
          </div>

          <div class="space-y-6">
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.url') }}</label>
              <input type="text" value="ldap://10.20.30.144" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
            </div>
            <div class="grid grid-cols-2 gap-6">
              <div class="space-y-2">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.port') }}</label>
                <input type="text" value="389" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
              </div>
              <div class="space-y-2">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.timeout') }}</label>
                <input type="text" value="30" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
              </div>
            </div>
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.bind_dn') }}</label>
              <input type="text" value="cn=admin,dc=akiradata,dc=co,dc=id" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-[10px] font-bold outline-none" />
            </div>
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.connection.password') }}</label>
              <div class="relative">
                <input type="password" value="secretpassword" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
                <LucideEye class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 cursor-pointer" />
              </div>
            </div>
          </div>

          <div class="p-6 bg-slate-50/50 dark:bg-slate-800/30 rounded-2xl border border-slate-100 dark:border-slate-800 flex gap-4 cursor-pointer group" @click="anonymousBind = !anonymousBind">
            <div :class="`mt-1 w-5 h-5 rounded border flex items-center justify-center transition-all shrink-0 ${anonymousBind ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200'}`"><LucideCheck v-if="anonymousBind" class="w-3.5 h-3.5 text-white" /></div>
            <p class="text-[10px] font-bold text-slate-500 leading-relaxed">{{ $t('settings.sso.connection.anonymous') }}</p>
          </div>
        </div>
      </div>

      <!-- Mapping & Filters -->
      <div class="lg:col-span-8 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-8">
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-xl bg-amber-50 text-amber-500 flex items-center justify-center"><LucideLink2 class="w-5 h-5" /></div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('settings.sso.mapping.title') }}</h3>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-10 gap-y-8">
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.mapping.user_base') }}</label>
              <textarea class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-[10px] font-bold outline-none h-24 resize-none">ou=users,dc=akiradata,dc=co,dc=id</textarea>
            </div>
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.mapping.user_attr') }}</label>
              <input type="text" value="uid" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
            </div>
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.mapping.group_base') }}</label>
              <textarea class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-[10px] font-bold outline-none h-24 resize-none">ou=groups,dc=akiradata,dc=co,dc=id</textarea>
            </div>
            <div class="space-y-6">
              <div class="space-y-2">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.mapping.login_attr') }}</label>
                <input type="text" value="sAMAccountName" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
              </div>
              <div class="space-y-2">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.mapping.lang') }}</label>
                <div class="relative">
                  <select class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none appearance-none cursor-pointer">
                    <option>Bahasa Indonesia (ID)</option>
                    <option>English (US)</option>
                  </select>
                  <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none" />
                </div>
              </div>
            </div>
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.mapping.user_class') }}</label>
              <input type="text" value="inetOrgPerson" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none" />
            </div>
          </div>
        </div>

        <div class="glass p-10 rounded-[3rem] space-y-8">
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><LucideFilter class="w-5 h-5" /></div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('settings.sso.filters.title') }}</h3>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
            <div class="space-y-6">
              <div class="space-y-2">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.filters.inclusion') }}</label>
                <textarea class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-[10px] font-bold outline-none h-28 resize-none font-mono">(&(objectClass=user)(memberOf=cn=DMS_Users,ou=groups,dc=akiradata...))</textarea>
              </div>
              <div class="flex items-center justify-between p-6 bg-slate-50/30 dark:bg-slate-800/20 rounded-2xl border border-slate-100 dark:border-slate-800">
                <div class="space-y-0.5">
                  <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('settings.sso.filters.orphan') }}</p>
                  <p class="text-[8px] font-bold text-slate-400 uppercase">{{ $t('settings.sso.filters.orphan_desc') }}</p>
                </div>
                <div class="w-12 h-6 bg-slate-200 dark:bg-slate-700 rounded-full relative cursor-pointer" @click="autoDisable = !autoDisable">
                  <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${autoDisable ? 'left-7 bg-[#1E3A5F]' : 'left-1'}`"></div>
                </div>
              </div>
            </div>
            <div class="space-y-8">
              <div class="space-y-2">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.filters.role') }}</label>
                <div class="relative">
                  <select class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none appearance-none cursor-pointer">
                    <option>Standard User (Viewer)</option>
                    <option>Admin</option>
                  </select>
                  <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none" />
                </div>
              </div>
              <div class="space-y-4">
                <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.filters.freq.title') }}</label>
                <div class="flex bg-slate-100 dark:bg-slate-800 p-1 rounded-xl">
                  <button v-for="f in ['6h', '12h', 'daily']" :key="f" :class="`flex-grow py-3 rounded-lg text-[9px] font-black uppercase tracking-widest transition-all ${syncFreq === f ? 'bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white shadow-sm border border-slate-100 dark:border-slate-800' : 'text-slate-400'}`" @click="syncFreq = f">
                    {{ $t(`settings.sso.filters.freq.${f}`) }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Last Sync Activity Footer -->
    <div class="bg-slate-100/50 dark:bg-slate-900/50 p-10 rounded-[3rem] border border-slate-100 dark:border-slate-800 flex flex-col md:flex-row items-center justify-between gap-10">
      <div class="flex items-center gap-10 flex-grow">
        <div class="space-y-1">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.footer.last_sync') }}</p>
          <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight">24 Oct 2023 - 14:20:05 WIB</p>
        </div>
        <div class="w-px h-10 bg-slate-200 dark:bg-slate-700 hidden md:block"></div>
        <div class="flex gap-10">
          <div class="text-center space-y-0.5">
            <p class="text-xl font-black text-[#1E3A5F] dark:text-white">1,204</p>
            <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.footer.users') }}</p>
          </div>
          <div class="text-center space-y-0.5">
            <p class="text-xl font-black text-amber-500">42</p>
            <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.footer.groups') }}</p>
          </div>
          <div class="text-center space-y-0.5">
            <p class="text-xl font-black text-red-500">0</p>
            <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('settings.sso.footer.errors') }}</p>
          </div>
        </div>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucideHistory class="w-4 h-4" />
          {{ $t('settings.sso.footer.btn_log') }}
        </button>
        <button class="px-10 py-4 bg-[#1E3A5F] text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center gap-3">
          <LucideRotateCw class="w-4 h-4" />
          {{ $t('settings.sso.footer.btn_run') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideEye, LucideZap, LucideSave, LucideNetwork, LucideServer, 
  LucideCheck, LucideLink2, LucideChevronDown, LucideFilter, 
  LucideHistory, LucideRotateCw 
} from 'lucide-vue-next'

const ldapActive = ref(true)
const provider = ref('ad')
const anonymousBind = ref(false)
const autoDisable = ref(false)
const syncFreq = ref('6h')

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
