import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'

interface User {
  id: string
  full_name: string
  email: string
  role: string
  avatar_url?: string
  signature_url?: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = useLocalStorage<User | null>('auth_user', null)
  const accessToken = useLocalStorage<string | null>('auth_token', null)
  const refreshToken = useLocalStorage<string | null>('refresh_token', null)

  const isAuthenticated = computed(() => !!accessToken.value)

  function setTokens(access: string, refresh: string) {
    accessToken.value = access
    refreshToken.value = refresh
    // Also set legacy 'token' for pages still using it directly
    if (typeof window !== 'undefined') {
      localStorage.setItem('token', access)
    }
  }

  function setUser(userData: User) {
    user.value = {
      id: userData.id,
      email: userData.email,
      full_name: userData.full_name,
      role: userData.role,
      avatar_url: userData.avatar_url,
      signature_url: userData.signature_url
    }
  }

  function logout() {
    user.value = null
    accessToken.value = null
    refreshToken.value = null
    if (typeof window !== 'undefined') {
      localStorage.removeItem('token')
    }
    navigateTo('/login')
  }

  return {
    user,
    accessToken,
    refreshToken,
    isAuthenticated,
    setTokens,
    setUser,
    logout
  }
})
