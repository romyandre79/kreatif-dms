<template>
  <div class="h-full flex flex-col space-y-6 p-8">
    <!-- Header -->
    <header class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-black text-[#1E3A5F] uppercase tracking-tight">Pusat Notifikasi</h1>
        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">Daftar aktivitas dan permintaan yang memerlukan perhatian Anda</p>
      </div>
      <div class="flex items-center gap-3">
        <button 
          @click="markAllAsRead" 
          class="px-5 py-2.5 bg-white border border-slate-200 rounded-xl text-xs font-black text-slate-600 uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2"
          :disabled="notifStore.unreadCount === 0"
        >
          <LucideCheckCircle2 class="w-4 h-4" />
          Tandai Semua Dibaca
        </button>
      </div>
    </header>

    <!-- Content -->
    <div class="flex-1 bg-white rounded-3xl border border-slate-200 shadow-sm overflow-hidden flex flex-col">
      <!-- Tabs/Filter -->
      <div class="px-8 py-5 border-b border-slate-100 flex items-center gap-8">
        <button class="text-xs font-black uppercase tracking-widest text-primary-600 relative py-2">
          Semua Notifikasi
          <div class="absolute bottom-0 left-0 right-0 h-1 bg-primary-600 rounded-full"></div>
        </button>
        <button class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors py-2">
          Belum Dibaca ({{ notifStore.unreadCount }})
        </button>
      </div>

      <!-- List -->
      <div class="flex-1 overflow-y-auto custom-scrollbar p-6">
        <div v-if="notifStore.notifications.length === 0" class="h-full flex flex-col items-center justify-center py-20 text-center space-y-4">
          <div class="w-20 h-20 bg-slate-50 rounded-full flex items-center justify-center">
            <LucideBellOff class="w-10 h-10 text-slate-300" />
          </div>
          <p class="text-sm font-bold text-slate-400">Tidak ada notifikasi untuk saat ini</p>
        </div>

        <div v-else class="space-y-3">
          <div 
            v-for="notif in notifStore.notifications" 
            :key="notif.id"
            @click="handleNotifClick(notif)"
            class="group p-6 rounded-2xl border transition-all cursor-pointer relative overflow-hidden"
            :class="[
              notif.is_read ? 'bg-white border-slate-100 opacity-60' : 'bg-slate-50/50 border-primary-100 shadow-sm'
            ]"
          >
            <div v-if="!notif.is_read" class="absolute left-0 top-0 bottom-0 w-1.5 bg-primary-500"></div>
            
            <div class="flex gap-6">
              <div :class="['w-14 h-14 rounded-2xl flex items-center justify-center shrink-0 shadow-sm', getNotifUI(notif.type).bg]">
                <component :is="getIcon(notif.type)" :class="['w-7 h-7', getNotifUI(notif.type).color]" />
              </div>
              
              <div class="flex-1 min-w-0 space-y-1">
                <div class="flex items-center justify-between">
                  <h3 class="text-base font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ notif.title }}</h3>
                  <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ formatTime(notif.created_at) }}</span>
                </div>
                <p class="text-sm text-slate-500 font-medium leading-relaxed">{{ notif.message || notif.body }}</p>
                
                <div v-if="notif.entity_id" class="pt-4 flex items-center gap-3">
                  <button class="px-4 py-2 bg-primary-500 text-white text-[10px] font-black rounded-lg uppercase tracking-widest hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/20">
                    Lihat Detail
                  </button>
                  <span class="text-[9px] font-black text-slate-300 uppercase tracking-widest">ID: {{ notif.entity_id }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideBell, LucideCheckCircle2, LucideBellOff, LucideInfo, 
  LucideAlertCircle, LucideFileText, LucideUser, LucideSettings
} from 'lucide-vue-next'
import { useNotificationStore } from '@/stores/notifications'
import { timeAgo } from '~/utils/format'

const notifStore = useNotificationStore()

const getIcon = (type) => {
  switch (type) {
    case 'doc-pending-approval': return LucideFileText
    case 'user-pending-approval': return LucideUser
    case 'system': return LucideSettings
    case 'success': return LucideCheckCircle2
    case 'warning': return LucideAlertCircle
    default: return LucideInfo
  }
}

const getNotifUI = (type) => {
  switch (type) {
    case 'success':
      return { bg: 'bg-emerald-50 border border-emerald-100', color: 'text-emerald-500' }
    case 'warning':
      return { bg: 'bg-amber-50 border border-amber-100', color: 'text-amber-500' }
    case 'doc-pending-approval':
      return { bg: 'bg-indigo-50 border border-indigo-100', color: 'text-indigo-500' }
    default:
      return { bg: 'bg-blue-50 border border-blue-100', color: 'text-blue-500' }
  }
}

const formatTime = (dateStr) => {
  if (!dateStr) return ''
  return timeAgo(dateStr)
}

const markAllAsRead = async () => {
  await notifStore.markAllAsRead()
}

const handleNotifClick = async (notif) => {
  if (!notif.is_read) {
    await notifStore.markAsRead(notif.id)
  }
  
  // Navigation logic based on entity
  if (notif.entity_type === 'document' && notif.entity_id) {
    navigateTo(`/documents/${notif.entity_id}`)
  }
}

onMounted(() => {
  notifStore.fetchNotifications()
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #E2E8F0;
  border-radius: 10px;
}
</style>
