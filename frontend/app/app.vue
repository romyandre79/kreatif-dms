<script setup lang="ts">
const { t } = useI18n()
const config = useRuntimeConfig()

useHead({
  titleTemplate: (title) => title ? `${title} - ${config.public.appName}` : config.public.appName,
})
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<style>
/* Transitions */
.page-enter-active,
.page-leave-active {
  transition: all 0.2s;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
  filter: blur(1rem);
}
</style>
