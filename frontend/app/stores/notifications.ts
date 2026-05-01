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
  const notifications = ref<Notification[]>([])
  const unreadCount = ref(0)
  const loading = ref(false)

  async function fetchNotifications() {
    const auth = useAuthStore()
    if (!auth.token) return

    loading.value = true
    try {
      const response = await $fetch<any>('/api/v1/notifications', {
        headers: {
          Authorization: `Bearer ${auth.token}`
        }
      })
      notifications.value = response.notifications || []
      unreadCount.value = response.unread_count || 0
    } catch (error) {
      console.error('Failed to fetch notifications:', error)
    } finally {
      loading.value = false
    }
  }

  async function markAsRead(id: string) {
    const auth = useAuthStore()
    try {
      await $fetch(`/api/v1/notifications/${id}/read`, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${auth.token}`
        }
      })
      await fetchNotifications()
    } catch (error) {
      console.error('Failed to mark notification as read:', error)
    }
  }

  async function markAllAsRead() {
    const auth = useAuthStore()
    try {
      await $fetch('/api/v1/notifications/read-all', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${auth.token}`
        }
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
