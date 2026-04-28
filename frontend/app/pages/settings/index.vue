<template>
  <div class="space-y-8">
    <PageHeader 
      :title="$t('layout.navbar.account_settings')"
      :subtitle="$t('placeholders.coming_soon_desc', { page: $t('layout.navbar.account_settings') })"
    />
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center space-y-4" v-motion-slide-visible-bottom>
      <div class="w-20 h-20 rounded-2xl bg-slate-50 dark:bg-slate-900/20 flex items-center justify-center text-slate-500">
        <LucideSettings class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-bold">{{ $t('placeholders.coming_soon_title', { page: $t('layout.navbar.account_settings') }) }}</h2>
      <p class="text-slate-500 max-w-md">{{ $t('placeholders.coming_soon_desc', { page: $t('layout.navbar.account_settings') }) }}</p>
    </div>
  </div>
</template>

<script setup>
import { LucideSettings } from 'lucide-vue-next'
import PageHeader from '~/components/PageHeader.vue'
</script>
