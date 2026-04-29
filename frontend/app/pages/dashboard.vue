<template>
  <div class="space-y-8">
    <AdminView v-if="user?.role === 'superadmin'" />
    <ManagerView v-else-if="user?.role === 'manajer'" />
    <DocControllerView v-else-if="user?.role === 'admin doc controller' || user?.role === 'kepala doc controller'" />
    <UserView v-else />
  </div>
</template>

<script setup>
import { useAuthStore } from '~/stores/auth'
import AdminView from '~/components/dashboard/AdminView.vue'
import ManagerView from '~/components/dashboard/ManagerView.vue'
import DocControllerView from '~/components/dashboard/DocControllerView.vue'
import UserView from '~/components/dashboard/UserView.vue'

const auth = useAuthStore()
const user = computed(() => auth.user)
</script>
