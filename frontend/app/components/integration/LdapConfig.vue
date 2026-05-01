<template>
  <div class="grid grid-cols-1 gap-8">
    <div class="space-y-6">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">LDAP Service Status</p>
          <button 
            @click="modelValue.is_active = !modelValue.is_active" 
            :class="`w-12 h-6 rounded-full transition-all relative ${modelValue.is_active ? 'bg-primary-500' : 'bg-slate-300 dark:bg-slate-700'}`"
          >
            <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${modelValue.is_active ? 'left-7' : 'left-1'}`"></div>
          </button>
        </div>
        <span v-if="modelValue.is_active" class="text-[9px] font-black text-green-500 uppercase tracking-widest animate-pulse">Service Active</span>
      </div>

      <div class="space-y-3 text-left">
        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Provider Type</label>
        <div class="grid grid-cols-2 gap-4">
          <button 
            @click="modelValue.driver = 'active_directory'" 
            :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${modelValue.driver === 'active_directory' ? 'bg-blue-50 border-blue-200 text-blue-600 dark:bg-blue-900/20 dark:border-blue-800' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-400'}`"
          >
            <LucideServer class="w-4 h-4" /> Active Directory
          </button>
          <button 
            @click="modelValue.driver = 'openldap'" 
            :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${modelValue.driver === 'openldap' ? 'bg-blue-50 border-blue-200 text-blue-600 dark:bg-blue-900/20 dark:border-blue-800' : 'bg-white dark:bg-slate-900 border-slate-100 dark:border-slate-800 text-slate-400'}`"
          >
            <LucideHash class="w-4 h-4" /> OpenLDAP
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Timeout (s)</label>
          <input v-model="config.timeout" type="number" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500" placeholder="30">
        </div>
        <div class="md:col-span-2 space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Bind DN</label>
          <input v-model="config.bind_dn" type="text" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500" placeholder="cn=admin,dc=example,dc=com">
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-left">
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Password</label>
          <div class="relative">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="config.bind_password" 
              class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 pr-14"
              placeholder="••••••••"
            >
            <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
              <LucideEye v-if="!showPassword" class="w-4 h-4" />
              <LucideEyeOff v-else class="w-4 h-4" />
            </button>
          </div>
        </div>
        <div class="space-y-2">
          <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Base DN (Users/Groups)</label>
          <input v-model="config.base_dn" type="text" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500" placeholder="dc=example,dc=com">
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { LucideServer, LucideHash, LucideEye, LucideEyeOff } from 'lucide-vue-next'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  }
})

const showPassword = ref(false)

const config = computed({
  get: () => props.modelValue.config || {},
  set: (val) => {
    props.modelValue.config = val
  }
})
</script>
