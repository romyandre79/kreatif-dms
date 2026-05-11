import { useAuthStore } from '~/stores/auth'

export const useApi = () => {
  const auth = useAuthStore()
  const config = useRuntimeConfig()

  const fetchWithAuth = async (url: string, options: any = {}) => {
    // Auto-prepend apiBase if relative URL
    const fullUrl = url.startsWith('http') ? url : `${config.public.apiBase}${url.startsWith('/') ? url : '/' + url}`

    // Add auth header if token exists
    if (auth.accessToken) {
      options.headers = {
        ...options.headers,
        'Authorization': `Bearer ${auth.accessToken}`
      }
    }

    try {
      const response = await $fetch.raw(fullUrl, options)
      return response._data
    } catch (err: any) {
      // Handle 401 Unauthorized
      if (err.response?.status === 401) {
        if (auth.refreshToken) {
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
              const retryRes = await $fetch.raw(fullUrl, options)
              return retryRes._data
            }
          } catch (refreshErr) {
            // Refresh failed, logout
            auth.logout()
            throw refreshErr
          }
        } else {
          // No refresh token, logout immediately
          auth.logout()
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
