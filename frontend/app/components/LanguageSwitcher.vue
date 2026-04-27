<template>
  <Dropdown v-model="isOpen">
    <template #trigger>
      <button 
        class="flex items-center gap-2 p-2.5 rounded-xl hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-500 transition-colors"
        title="Switch Language"
      >
        <LucideLanguages class="w-5 h-5" />
        <span class="text-xs font-black uppercase tracking-tighter hidden md:block">{{ locale }}</span>
      </button>
    </template>

    <div class="px-4 py-3 border-b border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50">
      <span class="text-xs font-black uppercase tracking-widest text-slate-400">Pilih Bahasa / Select Language</span>
    </div>
    
    <div class="py-1">
      <button
        v-for="loc in locales"
        :key="loc.code"
        @click="changeLocale(loc.code)"
        class="w-full flex items-center justify-between px-4 py-3 text-sm font-bold transition-colors"
        :class="[
          locale === loc.code 
            ? 'text-primary-600 bg-primary-50 dark:bg-primary-900/10' 
            : 'text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800/50'
        ]"
      >
        <span>{{ loc.name }}</span>
        <LucideCheck v-if="locale === loc.code" class="w-4 h-4" />
      </button>
    </div>
  </Dropdown>
</template>

<script setup>
import { ref } from 'vue'
import { LucideLanguages, LucideCheck } from 'lucide-vue-next'
import Dropdown from './Dropdown.vue'

const { locale, locales, setLocale } = useI18n()
const isOpen = ref(false)

const changeLocale = (code) => {
  setLocale(code)
  isOpen.value = false
}
</script>
