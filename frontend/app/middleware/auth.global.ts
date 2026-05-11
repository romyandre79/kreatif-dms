import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware((to, from) => {
  const auth = useAuthStore()

  // Define public routes
  const publicRoutes = ['/', '/login', '/registration']
  
  // If the route is not public and the user is not authenticated, redirect to login
  if (!publicRoutes.includes(to.path) && !auth.isAuthenticated) {
    return navigateTo('/login')
  }

  // If the user is authenticated and tries to access login, redirect to dashboard
  if (to.path === '/login' && auth.isAuthenticated) {
    return navigateTo('/dashboard')
  }
})
