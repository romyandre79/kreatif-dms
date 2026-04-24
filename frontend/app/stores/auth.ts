import { defineStore } from 'pinia'

interface User {
  id: string
  full_name: string
  email: string
  role: string
  avatar_url?: string
  signature_url?: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const accessToken = ref<string | null>(null)
  const refreshToken = ref<string | null>(null)

  const isAuthenticated = computed(() => !!accessToken.value)

  function setTokens(access: string, refresh: string) {
    accessToken.value = access
    refreshToken.value = refresh
    // In a real app, you'd save these to cookies or localStorage
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
