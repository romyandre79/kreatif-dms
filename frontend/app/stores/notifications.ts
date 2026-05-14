import { defineStore } from 'pinia'
import { useAuthStore } from './auth'

interface Notification {
  id: string
  title: string
  message: string
  type: string
  is_read: boolean
  created_at: string
}

export const useNotificationStore = defineStore('notification', () => {
  const { $api } = useApi()
  const notifications = ref<Notification[]>([])
  const unreadCount = ref(0)
  const loading = ref(false)

  async function fetchNotifications() {
    loading.value = true
    try {
      const response = await $api('/notifications')
      notifications.value = response.notifications || []
      unreadCount.value = response.unread_count || 0
    } catch (error) {
      console.error('Failed to fetch notifications:', error)
    } finally {
      loading.value = false
    }
  }

  async function markAsRead(id: string) {
    try {
      await $api(`/notifications/${id}/read`, {
        method: 'POST'
      })
      await fetchNotifications()
    } catch (error) {
      console.error('Failed to mark notification as read:', error)
    }
  }

  async function markAllAsRead() {
    try {
      await $api('/notifications/read-all', {
        method: 'POST'
      })
      await fetchNotifications()
    } catch (error) {
      console.error('Failed to mark all as read:', error)
    }
  }

  return {
    notifications,
    unreadCount,
    loading,
    fetchNotifications,
    markAsRead,
    markAllAsRead
  }
})
