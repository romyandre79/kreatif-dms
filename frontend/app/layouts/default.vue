<template>
  <div class="min-h-screen bg-[#F8FAFC] dark:bg-[#0A0F1C] font-sans selection:bg-primary-500/30">
    <!-- Mobile Sidebar Overlay -->
    <Transition name="fade">
      <div v-if="isSidebarOpen" 
           @click="isSidebarOpen = false"
           class="fixed inset-0 bg-[#1E3A5F]/40 backdrop-blur-sm z-[100] lg:hidden">
      </div>
    </Transition>

    <!-- Sidebar Navigation -->
    <aside :class="[
      'fixed top-0 left-0 h-full w-[280px] bg-white dark:bg-[#0D121F] border-r border-slate-200/60 dark:border-slate-800/40 z-[110] transition-all duration-500 cubic-bezier(0.4, 0, 0.2, 1)',
      isSidebarOpen ? 'translate-x-0 shadow-2xl' : '-translate-x-full lg:translate-x-0'
    ]">
      <div class="flex flex-col h-full">
        <!-- Logo Section -->
        <div class="p-8 pb-4">
          <NuxtLink to="/dashboard" class="flex items-center gap-3 group transition-transform hover:scale-[1.02] active:scale-95">
            <div class="w-10 h-10 rounded-2xl bg-white dark:bg-[#1A2234] flex items-center justify-center shadow-xl shadow-primary-500/10 ring-1 ring-slate-200/50 dark:ring-slate-800/50 overflow-hidden p-2">
              <img src="/logo.png" alt="Kreatif DMS" class="w-full h-full object-contain" />
            </div>
            <div>
              <h1 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter leading-none">Kreatif</h1>
              <span class="text-[10px] font-bold text-primary-500 uppercase tracking-widest leading-none">DMS Protocol</span>
            </div>
          </NuxtLink>
        </div>

        <!-- Menu Navigation -->
        <nav class="flex-1 overflow-y-auto px-4 py-6 custom-scrollbar space-y-1">
          <div v-for="item in dynamicMenu" :key="item.key" class="space-y-1">
            <!-- Parent Menu with Children -->
            <div v-if="item.children && item.children.length > 0">
              <button 
                @click="toggleSubmenu(item.key)"
                class="w-full flex items-center justify-between px-5 py-3.5 rounded-2xl transition-all duration-300 group"
                :class="openSubmenus.includes(item.key) ? 'bg-slate-50 dark:bg-slate-800/40' : 'hover:bg-slate-50/50 dark:hover:bg-slate-800/20'"
              >
                <div class="flex items-center gap-4">
                  <div class="w-8 h-8 rounded-xl flex items-center justify-center transition-colors"
                       :class="openSubmenus.includes(item.key) ? 'bg-primary-500/10 text-primary-500' : 'bg-slate-100 dark:bg-slate-800 text-slate-400 group-hover:text-primary-500'">
                    <component :is="getIcon(item.icon || item.key)" class="w-4 h-4" />
                  </div>
                  <div class="flex-1 flex items-center justify-between min-w-0">
                    <span class="text-[11px] font-black uppercase tracking-tight text-slate-600 dark:text-slate-300 group-hover:text-primary-500 transition-colors text-left">
                      {{ t('layout.menu.' + item.key) }}
                    </span>
                    <div class="flex items-center gap-2 ml-4">
                      <span v-if="item.badge" class="px-1.5 py-0.5 rounded-full bg-red-500 text-white text-[9px] font-black">{{ item.badge }}</span>
                      <LucideChevronRight class="w-4 h-4 text-slate-400 transition-transform duration-300" 
                                          :class="{ 'rotate-90 text-primary-500': openSubmenus.includes(item.key) }" />
                    </div>
                  </div>
                </div>
              </button>
              
              <Transition name="expand">
                <div v-if="openSubmenus.includes(item.key)" class="mt-1 space-y-1 pl-4">
                  <NuxtLink v-for="child in item.children" :key="child.key" 
                           :to="child.path"
                           @click="isSidebarOpen = false"
                           class="flex items-center gap-4 px-9 py-3 rounded-xl transition-all duration-300 group relative overflow-hidden"
                           :class="route.path === child.path ? 'bg-primary-500/5 text-primary-500' : 'text-slate-400 hover:text-primary-500'">
                    <div v-if="route.path === child.path" class="absolute left-0 top-0 bottom-0 w-1 bg-primary-500 rounded-full"></div>
                    <div class="flex-1 flex items-center justify-between min-w-0">
                      <span class="text-[10px] font-black uppercase tracking-tight truncate text-left">{{ t('layout.menu.' + child.key) }}</span>
                      <span v-if="child.badge" class="px-1.5 py-0.5 rounded-md bg-white/10 text-white/50 text-[8px] font-black">{{ child.badge }}</span>
                    </div>
                  </NuxtLink>
                </div>
              </Transition>
            </div>

            <!-- Single Menu Link -->
            <NuxtLink v-else 
                     :to="item.path"
                     @click="isSidebarOpen = false"
                     class="w-full flex items-center gap-4 px-5 py-3.5 rounded-2xl transition-all duration-300 group relative overflow-hidden"
                     :class="route.path === item.path ? 'bg-primary-500 text-white shadow-xl shadow-primary-500/20' : 'text-slate-500 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800/40'">
              <div class="w-8 h-8 rounded-xl flex items-center justify-center transition-colors"
                   :class="route.path === item.path ? 'bg-white/20' : 'bg-slate-100 dark:bg-slate-800 text-slate-400 group-hover:text-primary-500'">
                <component :is="getIcon(item.icon || item.key)" class="w-4 h-4" />
              </div>
              <div class="flex-1 flex items-center justify-between min-w-0">
                <span class="text-[11px] font-black uppercase tracking-tight truncate text-left">{{ t('layout.menu.' + item.key) }}</span>
                <span v-if="item.badge" :class="`px-1.5 py-0.5 rounded-full text-[9px] font-black ${route.path === item.path ? 'bg-white/20 text-white' : 'bg-red-500 text-white'}`">{{ item.badge }}</span>
              </div>
            </NuxtLink>
          </div>
        </nav>

      </div>
    </aside>

    <!-- Main Content Wrapper -->
    <div class="lg:ml-[280px] flex flex-col min-h-screen">
      <!-- Header -->
      <header class="sticky top-0 z-[90] h-20 bg-white/70 dark:bg-[#0A0F1C]/70 backdrop-blur-xl border-b border-slate-200/60 dark:border-slate-800/40 px-8 flex items-center justify-between">
        <div class="flex items-center gap-6">
          <button @click="isSidebarOpen = true" 
                  class="lg:hidden p-2 rounded-xl hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors">
            <LucideMenu class="w-5 h-5 text-[#1E3A5F] dark:text-white" />
          </button>
          
          <div class="hidden sm:flex items-center gap-4">
            <!-- Breadcrumbs -->
            <div class="flex items-center gap-2 text-xs font-black uppercase tracking-wider text-slate-400 dark:text-slate-500">
              <NuxtLink to="/dashboard" class="hover:text-primary-500 dark:hover:text-primary-400 transition-all duration-300 flex items-center gap-1.5 py-1">
                <LucideHome class="w-3.5 h-3.5 text-slate-400 dark:text-slate-500" />
                <span>DMS</span>
              </NuxtLink>
              
              <template v-for="(crumb, index) in computedBreadcrumbs" :key="index">
                <LucideChevronRight class="w-3.5 h-3.5 text-slate-300 dark:text-slate-700" />
                <NuxtLink 
                  v-if="crumb.path && index < computedBreadcrumbs.length - 1" 
                  :to="crumb.path" 
                  class="hover:text-primary-500 dark:hover:text-primary-400 transition-all duration-300 py-1"
                >
                  {{ crumb.label }}
                </NuxtLink>
                <span v-else class="text-[#1E3A5F] dark:text-white font-extrabold py-1">
                  {{ crumb.label }}
                </span>
              </template>
            </div>
            
            <div class="h-4 w-px bg-slate-200 dark:bg-slate-800/80 mx-2"></div>
            
            <!-- Search -->
            <div class="flex items-center gap-3">
              <LucideSearch class="w-4 h-4 text-slate-400" />
              <input type="text" 
                     v-model="searchQuery"
                     @keyup.enter="handleHeaderSearch"
                     placeholder="Search anything..." 
                     class="bg-transparent border-none focus:ring-0 text-sm font-medium text-slate-600 dark:text-slate-300 placeholder:text-slate-400 w-64" />
            </div>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <!-- Language Switcher -->
          <div class="hidden md:block">
            <LanguageSwitcher />
          </div>

          <!-- Notification Bell -->
          <div class="relative">
            <button @click="showNotifications = !showNotifications"
                    class="p-2.5 rounded-xl bg-slate-50 dark:bg-slate-800/40 hover:bg-primary-500/10 hover:text-primary-500 transition-all group relative">
              <LucideBell class="w-5 h-5 text-slate-500 group-hover:text-primary-500 transition-colors" />
              <div v-if="notifStore.unreadCount > 0" class="absolute top-2 right-2 w-2 h-2 rounded-full bg-primary-500 ring-2 ring-white dark:ring-[#0A0F1C] animate-pulse"></div>
            </button>
            
            <Transition name="fade">
              <div v-if="showNotifications" 
                   class="absolute right-0 mt-3 w-80 bg-white dark:bg-[#0D121F] border border-slate-200 dark:border-slate-800 rounded-lg shadow-2xl overflow-hidden z-[100]">
                <div class="p-4 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-900/50">
                  <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Notifications</h3>
                  <button @click="handleMarkAllRead" class="text-[10px] font-bold text-primary-500 hover:underline">Mark all read</button>
                </div>
                <div class="max-h-96 overflow-y-auto custom-scrollbar">
                  <div v-if="notifStore.notifications.length === 0" class="p-10 text-center text-slate-400 italic text-[11px]">No notifications</div>
                  <div v-for="notif in notifStore.notifications" :key="notif.id"
                       @click="handleNotifClick(notif)"
                       class="p-4 hover:bg-slate-50 dark:hover:bg-slate-800 transition-colors cursor-pointer border-b border-slate-50 dark:border-slate-800 last:border-0"
                       :class="{ 'opacity-60': notif.is_read }">
                    <div class="flex gap-4">
                      <div :class="['w-10 h-10 rounded-xl flex items-center justify-center shrink-0', getNotifUI(notif.type).bg]">
                        <component :is="getIcon(notif.type)" :class="['w-5 h-5', getNotifUI(notif.type).color]" />
                      </div>
                      <div class="flex-1 min-w-0">
                        <p class="text-[11px] font-bold text-slate-700 dark:text-slate-300 line-clamp-1">{{ notif.title }}</p>
                        <p class="text-[10px] text-slate-500 dark:text-slate-400 line-clamp-2 mt-0.5 leading-relaxed">{{ notif.message }}</p>
                        <span class="text-[8px] font-bold text-slate-400 uppercase tracking-widest mt-2 block">{{ formatNotifTime(notif.created_at) }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </Transition>
          </div>

          <!-- User Profile -->
          <div class="relative">
            <button @click="showUserMenu = !showUserMenu" 
                    class="flex items-center gap-3 p-1.5 rounded-2xl bg-slate-50 dark:bg-slate-800/40 hover:bg-white dark:hover:bg-slate-800 transition-all border border-transparent hover:border-slate-200/60 dark:hover:border-slate-700">
              <div class="w-9 h-9 rounded-xl bg-primary-500 text-white flex items-center justify-center font-black text-xs shadow-lg shadow-primary-500/20 overflow-hidden ring-2 ring-primary-500/10">
                <img v-if="auth.user?.avatar_url" :src="auth.user.avatar_url" class="w-full h-full object-cover">
                <span v-else>{{ auth.user?.full_name?.charAt(0) }}</span>
              </div>
              <div class="hidden md:block text-left mr-2">
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight leading-none mb-1">{{ auth.user?.full_name }}</p>
                <p class="text-[9px] font-bold text-primary-500 uppercase tracking-widest leading-none">{{ auth.user?.role }}</p>
              </div>
              <LucideChevronDown class="w-4 h-4 text-slate-400" :class="{ 'rotate-180': showUserMenu }" />
            </button>

            <Transition name="fade">
              <div v-if="showUserMenu" 
                   class="absolute right-0 mt-3 w-56 bg-white dark:bg-[#0D121F] border border-slate-200 dark:border-slate-800 rounded-lg shadow-2xl overflow-hidden z-[100]">
                <div class="p-4 border-b border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50">
                  <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1 text-left">Manage Account</p>
                  <p class="text-[11px] font-bold text-slate-700 dark:text-slate-300 line-clamp-1 text-left">{{ auth.user?.email }}</p>
                </div>
                <div class="p-2">
                  <NuxtLink to="/profile" @click="showUserMenu = false" class="w-full flex items-center gap-3 px-4 py-3 rounded-xl hover:bg-slate-50 dark:hover:bg-slate-800 text-slate-600 dark:text-slate-400 transition-colors">
                    <LucideUser class="w-4 h-4" />
                    <span class="text-[10px] font-black uppercase tracking-tight">My Profile</span>
                  </NuxtLink>
                  <NuxtLink to="/config/params" @click="showUserMenu = false" class="w-full flex items-center gap-3 px-4 py-3 rounded-xl hover:bg-slate-50 dark:hover:bg-slate-800 text-slate-600 dark:text-slate-400 transition-colors">
                    <LucideSettings class="w-4 h-4" />
                    <span class="text-[10px] font-black uppercase tracking-tight">Settings</span>
                  </NuxtLink>
                  <div class="h-px bg-slate-100 dark:bg-slate-800 my-2"></div>
                  <button @click="handleLogout" class="w-full flex items-center gap-3 px-4 py-3 rounded-xl hover:bg-red-50 dark:hover:bg-red-900/20 text-red-500 transition-colors">
                    <LucideLogOut class="w-4 h-4" />
                    <span class="text-[10px] font-black uppercase tracking-tight">Logout System</span>
                  </button>
                </div>
              </div>
            </Transition>
          </div>
        </div>
      </header>

      <!-- Main Content Slot -->
      <main class="flex-1 p-2">
        <slot />
      </main>
    </div>

    <!-- Global Toast Notifications -->
    <ToastContainer />
    
    <!-- Global Cart Bar -->
    <FloatingCart />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notifications'
import { 
  LucideShield, LucideLayoutDashboard, LucidePlusCircle, LucideSearch, 
  LucideRepeat, LucideHome, LucideSettings, LucideHistory,
  LucideChevronDown, LucideChevronRight, LucideMenu, LucideBell,
  LucideLogOut, LucideUser, LucideFileText, LucideBox, LucideTrash2,
  LucideBarChart3, LucideActivity, LucideShieldAlert, LucideCheckCircle2,
  LucideAlertCircle, LucideInfo, LucideUploadCloud, LucideUpload, LucideInbox,
  LucidePieChart, LucideMap, LucideQrCode, LucideShoppingBag,
  LucideLogIn, LucideClock, LucideFileJson, LucideClipboardList,
  LucideScanLine, LucideGitCompare, LucideHardDrive, LucideBuilding2,
  LucideMapPin, LucideUsers, LucideServer, LucideArchive,
  LucideFolder, LucideSettings2, LucideUserCog, LucideBookOpen, LucideLayers,
  LucideLock, LucideMegaphone
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const { t } = useI18n()
const auth = useAuthStore()
const config = useRuntimeConfig()
const route = useRoute()
const notifStore = useNotificationStore()

const isSidebarOpen = ref(false)
const searchQuery = ref('')
const showNotifications = ref(false)
const showUserMenu = ref(false)
const dynamicMenu = ref([])
const openSubmenus = ref([])

const computedBreadcrumbs = computed(() => {
  const list = []
  const currentPath = route.path
  
  if (currentPath === '/dashboard') {
    list.push({
      label: t('layout.menu.dashboard') || 'Dashboard',
      path: '/dashboard'
    })
    return list
  }
  
  // Try to find matching item in dynamicMenu
  for (const item of dynamicMenu.value) {
    if (item.path === currentPath) {
      list.push({
        label: t('layout.menu.' + item.key) || item.label || item.key,
        path: item.path
      })
      return list
    }
    
    if (item.children) {
      for (const child of item.children) {
        if (child.path === currentPath) {
          list.push({
            label: t('layout.menu.' + item.key) || item.label || item.key,
            path: item.path || ''
          })
          list.push({
            label: t('layout.menu.' + child.key) || child.label || child.key,
            path: child.path
          })
          return list
        }
      }
    }
  }
  
  // Fallback for subpaths or pages not explicitly in dynamicMenu
  const segments = currentPath.split('/').filter(Boolean)
  let accumulatedPath = ''
  segments.forEach((segment) => {
    accumulatedPath += `/${segment}`
    let foundMenu = null
    for (const item of dynamicMenu.value) {
      if (item.path === accumulatedPath) {
        foundMenu = item
        break
      }
      if (item.children) {
        for (const child of item.children) {
          if (child.path === accumulatedPath) {
            foundMenu = child
            break
          }
        }
      }
    }
    
    if (foundMenu) {
      list.push({
        label: t('layout.menu.' + foundMenu.key) || foundMenu.label || foundMenu.key,
        path: foundMenu.path
      })
    } else {
      const isId = /^[0-9a-fA-F-]+$/.test(segment) || /^\d+$/.test(segment)
      if (isId) {
        list.push({
          label: `#${segment.substring(0, 8).toUpperCase()}`,
          path: accumulatedPath
        })
      } else {
        const humanized = segment
          .replace(/[-_]/g, ' ')
          .replace(/\b\w/g, c => c.toUpperCase())
        list.push({
          label: humanized,
          path: accumulatedPath
        })
      }
    }
  })
  
  return list
})

// Icon map for dynamic resolution
const iconMap = {
  'LucideLayoutDashboard': LucideLayoutDashboard,
  'LucidePlusCircle': LucidePlusCircle,
  'LucideSearch': LucideSearch,
  'LucideRepeat': LucideRepeat,
  'LucideHome': LucideHome,
  'LucideSettings': LucideSettings,
  'LucideHistory': LucideHistory,
  'LucideFileText': LucideFileText,
  'LucideBox': LucideBox,
  'LucideTrash2': LucideTrash2,
  'LucideBarChart3': LucideBarChart3,
  'LucideActivity': LucideActivity,
  'LucideShieldAlert': LucideShieldAlert,
  'LucideCheckCircle2': LucideCheckCircle2,
  'LucideAlertCircle': LucideAlertCircle,
  'LucideInfo': LucideInfo,
  'LucideUploadCloud': LucideUploadCloud,
  'LucideInbox': LucideInbox,
  'LucidePieChart': LucidePieChart,
  'LucideMap': LucideMap,
  'LucideQrCode': LucideQrCode,
  'LucideShoppingBag': LucideShoppingBag,
  'LucideLogOut': LucideLogOut,
  'LucideLogIn': LucideLogIn,
  'LucideClock': LucideClock,
  'LucideFileJson': LucideFileJson,
  'LucideClipboardList': LucideClipboardList,
  'LucideScanLine': LucideScanLine,
  'LucideGitCompare': LucideGitCompare,
  'LucideHardDrive': LucideHardDrive,
  'LucideBuilding2': LucideBuilding2,
  'LucideMapPin': LucideMapPin,
  'LucideUsers': LucideUsers,
  'LucideServer': LucideServer,
  'LucideArchive': LucideArchive,
  'LucideFolder': LucideFolder,
  'LucideSettings2': LucideSettings2,
  'LucideUserCog': LucideUserCog,
  'LucideBookOpen': LucideBookOpen,
  'LucideLayers': LucideLayers,
  'Dashboard': LucideLayoutDashboard,
  'Registration': LucidePlusCircle,
  'Warehouse': LucideHome,
  'Search': LucideSearch,
  'Circulation': LucideRepeat,
  'Stock Take': LucideBox,
  'Retention': LucideTrash2,
  'Reports': LucideBarChart3,
  'Configuration': LucideSettings,
  'Administration': LucideShieldAlert,
  'cat_admin': LucideShieldAlert,
  'modules': LucideLayers,
  'roles': LucideShieldAlert,
  'users': LucideUserCog,
  'admin_pin': LucideLock,
  'admin_hierarchy': LucideGitCompare,
  'LucideLock': LucideLock,
  'Dashboard': LucideLayoutDashboard,
  'Search': LucideSearch,
  'Submit': LucideUpload,
  'Loans': LucideFileText,
  'Tracking': LucideQrCode,
  'tracking': LucideQrCode,
  'cat_warehouse': LucideHome,
  'Warehouse Management': LucideHome,
  'LucideBell': LucideBell,
  'LucideMegaphone': LucideMegaphone
}

const getIcon = (key) => iconMap[key] || LucideFileText

const getLoansPath = () => {
  const role = auth.user?.role
  if (role === 'admin doc controller' || role === 'kepala doc controller' || role === 'superadmin') {
    return '/circulation/checkout'
  }
  return '/loans/my'
}

const fetchDynamicMenu = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/auth/me/menu`)
    if (res && res.data) {
      // Patch menu items to use correct paths
      dynamicMenu.value = res.data.map(item => {
        let path = item.path
        if (path === '/registration/new') {
          path = '/documents/upload'
        } else if (path === '/circulation/checkout' || path === '/loans/request' || path === '/loans' || path === '/loans/my') {
          path = getLoansPath()
        } else if (path === '/approvals/submissions') {
          path = '/approvals'
        }
        
        return {
          ...item,
          path,
          children: item.children?.map(child => {
            let childPath = child.path
            if (childPath === '/registration/new') {
              childPath = '/documents/upload'
            } else if (childPath === '/circulation/checkout' || childPath === '/loans/request' || childPath === '/loans' || childPath === '/loans/my') {
              childPath = getLoansPath()
            } else if (childPath === '/approvals/submissions') {
              childPath = '/approvals'
            }
            return {
              ...child,
              path: childPath
            }
          })
        }
      })
    }
  } catch (err) {
    console.error('Failed to fetch dynamic menu:', err)
  }
}

const toggleSubmenu = (key) => {
  if (openSubmenus.value.includes(key)) {
    openSubmenus.value = openSubmenus.value.filter(n => n !== key)
  } else {
    openSubmenus.value.push(key)
  }
}

const handleHeaderSearch = () => {
  if (searchQuery.value.trim()) {
    navigateTo(`/documents?q=${encodeURIComponent(searchQuery.value)}`)
  }
}

const handleLogout = () => {
  showUserMenu.value = false
  auth.logout()
}

const getNotifUI = (type) => {
  switch (type) {
    case 'success':
      return { icon: LucideCheckCircle2, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' }
    case 'warning':
      return { icon: LucideAlertCircle, color: 'text-orange-500', bg: 'bg-orange-50 dark:bg-orange-900/20' }
    default:
      return { icon: LucideInfo, color: 'text-blue-500', bg: 'bg-blue-50 dark:bg-blue-900/20' }
  }
}

const formatNotifTime = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

const handleMarkAllRead = async () => {
  await notifStore.markAllAsRead()
}

const handleNotifClick = async (notif) => {
  if (!notif.is_read) {
    await notifStore.markAsRead(notif.id)
  }
  
  showNotifications.value = false
  
  // Navigation logic
  if (notif.entity_type === 'loan_request' || notif.entity_type === 'loan') {
    navigateTo(`/approvals/loans?id=${notif.entity_id}`)
  } else if (notif.entity_type === 'loan_extension' || notif.entity_type === 'extension') {
    navigateTo(`/approvals/extensions?id=${notif.entity_id}`)
  } else if (notif.entity_type === 'document_upload') {
    navigateTo(`/approvals/${notif.entity_id}`)
  } else if (notif.entity_type === 'document' && notif.entity_id) {
    navigateTo(`/documents/${notif.entity_id}`)
  } else if (notif.type === 'doc-pending-approval') {
    navigateTo('/approvals')
  }
}

onMounted(() => {
  isSidebarOpen.value = false
  if (auth.accessToken) {
    fetchDynamicMenu()
    notifStore.fetchNotifications()
    
    // Periodically refresh menu badges and notifications
    setInterval(() => {
      fetchDynamicMenu()
      notifStore.fetchNotifications()
    }, 10000)
  }
})

// Watch for authentication changes to refresh menu
watch(() => auth.isAuthenticated, (newVal) => {
  if (newVal) {
    fetchDynamicMenu()
    notifStore.fetchNotifications()
  } else {
    dynamicMenu.value = []
  }
}, { immediate: true })

// Update browser tab title based on route
watch(() => route.path, () => {
  // Refresh menu badges on navigation
  if (auth.accessToken) {
    fetchDynamicMenu()
  }

  // Find current menu item label and auto-expand its parent
  const findAndExpand = (items) => {
    for (const item of items) {
      if (item.path === route.path) return item.key
      if (item.children) {
        for (const child of item.children) {
          if (child.path === route.path) {
            if (!openSubmenus.value.includes(item.key)) {
              openSubmenus.value.push(item.key)
            }
            return child.key
          }
        }
      }
    }
    return null
  }
  
  const key = findAndExpand(dynamicMenu.value)
  if (key) {
    useHead({
      title: t('layout.menu.' + key)
    })
  }
}, { immediate: true })

if (import.meta.client) {
  window.addEventListener('resize', () => {
    if (window.innerWidth >= 1024) {
      isSidebarOpen.value = false
    }
  })
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(148, 163, 184, 0.2);
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(148, 163, 184, 0.4);
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.expand-enter-active, .expand-leave-active {
  transition: all 0.3s ease-out;
  max-height: 500px;
  overflow: hidden;
}
.expand-enter-from, .expand-leave-to {
  max-height: 0;
  opacity: 0;
}
</style>
