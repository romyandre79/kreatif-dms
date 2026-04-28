import { defineStore } from 'pinia'
import { useAuthStore } from './auth'

export const useNotificationStore = defineStore('notification', {
  state: () => ({
    notifications: [],
    unreadCount: 0,
    loading: false
  }),

  actions: {
    async fetchNotifications() {
      const auth = useAuthStore()
      if (!auth.token) return

      this.loading = true
      try {
        const response = await $fetch('/api/v1/notifications', {
          headers: {
            Authorization: `Bearer ${auth.token}`
          }
        })
        this.notifications = response.notifications || []
        this.unreadCount = response.unread_count || 0
      } catch (error) {
        console.error('Failed to fetch notifications:', error)
      } finally {
        this.loading = false
      }
    },

    async markAsRead(id) {
      const auth = useAuthStore()
      try {
        await $fetch(`/api/v1/notifications/${id}/read`, {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${auth.token}`
          }
        })
        await this.fetchNotifications()
      } catch (error) {
        console.error('Failed to mark notification as read:', error)
      }
    },

    async markAllAsRead() {
      const auth = useAuthStore()
      try {
        await $fetch('/api/v1/notifications/read-all', {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${auth.token}`
          }
        })
        await this.fetchNotifications()
      } catch (error) {
        console.error('Failed to mark all as read:', error)
      }
    }
  }
})
