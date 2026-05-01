import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'

interface User {
  id: string
  full_name: string
  email: string
  role: string
  avatar_url?: string
  signature_url?: string
  is_mfa_enabled?: boolean
  pin_status?: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const accessToken = ref<string | null>(null)
  const refreshToken = ref<string | null>(null)

  // Initialize from localStorage (Client-side only)
  if (typeof window !== 'undefined') {
    try {
      const savedUser = localStorage.getItem('kreatif_user')
      const savedAccess = localStorage.getItem('kreatif_access_token')
      const savedRefresh = localStorage.getItem('kreatif_refresh_token')

      if (savedUser) user.value = JSON.parse(savedUser)
      if (savedAccess) accessToken.value = savedAccess
      if (savedRefresh) refreshToken.value = savedRefresh
    } catch (e) {
      console.error('Auth store initialization failed', e)
    }
  }

  // Watch for changes and persist (Client-side only)
  if (typeof window !== 'undefined') {
    watch(user, (val) => {
      if (val) localStorage.setItem('kreatif_user', JSON.stringify(val))
      else localStorage.removeItem('kreatif_user')
    }, { deep: true })

    watch(accessToken, (val) => {
      if (val) localStorage.setItem('kreatif_access_token', val)
      else localStorage.removeItem('kreatif_access_token')
    })

    watch(refreshToken, (val) => {
      if (val) localStorage.setItem('kreatif_refresh_token', val)
      else localStorage.removeItem('kreatif_refresh_token')
    })
  }

  const isAuthenticated = computed(() => !!accessToken.value)

  function setTokens(access: string, refresh: string) {
    accessToken.value = access
    refreshToken.value = refresh
    if (typeof window !== 'undefined') {
      localStorage.setItem('token', access) // legacy support
    }
  }

  function setUser(userData: any) {
    user.value = {
      id: userData.id,
      email: userData.email,
      full_name: userData.full_name,
      role: userData.role,
      avatar_url: userData.avatar_url,
      signature_url: userData.signature_url,
      is_mfa_enabled: userData.is_mfa_enabled,
      pin_status: userData.pin_status
    }
  }

  function logout() {
    user.value = null
    accessToken.value = null
    refreshToken.value = null
    if (typeof window !== 'undefined') {
      localStorage.removeItem('token')
      localStorage.removeItem('kreatif_user')
      localStorage.removeItem('kreatif_access_token')
      localStorage.removeItem('kreatif_refresh_token')
    }
    navigateTo('/login')
  }

  const token = computed(() => accessToken.value)

  return {
    user,
    accessToken,
    refreshToken,
    token,
    isAuthenticated,
    setTokens,
    setUser,
    logout
  }
})
