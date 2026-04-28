import { useAuthStore } from '~/stores/auth'

export const useApi = () => {
  const auth = useAuthStore()
  const config = useRuntimeConfig()

  const fetchWithAuth = async (url: string, options: any = {}) => {
    // Add auth header if token exists
    if (auth.accessToken) {
      options.headers = {
        ...options.headers,
        'Authorization': `Bearer ${auth.accessToken}`
      }
    }

    try {
      const response = await $fetch.raw(url, options)
      return response._data
    } catch (err: any) {
      // Handle 401 Unauthorized
      if (err.response?.status === 401 && auth.refreshToken) {
        try {
          // Attempt to refresh token
          const refreshRes: any = await $fetch(`${config.public.apiBase}/auth/refresh`, {
            method: 'POST',
            body: { refresh_token: auth.refreshToken }
          })

          if (refreshRes && refreshRes.data) {
            // Update tokens
            auth.setTokens(refreshRes.data.access_token, refreshRes.data.refresh_token)
            
            // Retry original request with new token
            options.headers['Authorization'] = `Bearer ${auth.accessToken}`
            const retryRes = await $fetch.raw(url, options)
            return retryRes._data
          }
        } catch (refreshErr) {
          // Refresh failed, logout
          auth.logout()
          throw refreshErr
        }
      }
      
      // Handle 403 Forbidden
      if (err.response?.status === 403) {
        console.error('Permission denied:', url)
      }

      throw err
    }
  }

  return {
    fetchWithAuth,
    $api: fetchWithAuth
  }
}
