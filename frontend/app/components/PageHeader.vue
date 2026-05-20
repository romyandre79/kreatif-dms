<template>
  <div class="mb-2" v-motion-fade>
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl md:text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tight">
          {{ title }}
        </h1>
        <p v-if="subtitle" class="text-slate-500 dark:text-slate-400 font-medium text-sm md:text-base max-w-2xl leading-relaxed">
          {{ subtitle }}
        </p>
      </div>

      <div class="flex items-center gap-3 self-start md:self-center">
        <slot name="actions" />
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    required: true
  },
  subtitle: {
    type: String,
    default: ''
  }
})
</script>
