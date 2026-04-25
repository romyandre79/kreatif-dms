<template>
  <div class="min-h-screen bg-slate-50 dark:bg-slate-950 font-sans text-slate-900 dark:text-slate-50">
    <!-- Sidebar Overlay (Mobile only) -->
    <div 
      class="fixed inset-0 z-30 bg-slate-950/60 backdrop-blur-sm lg:hidden transition-all duration-300"
      :class="[isSidebarOpen ? 'opacity-100 visible' : 'opacity-0 invisible pointer-events-none']"
      @click="isSidebarOpen = false"
    ></div>

    <!-- Sidebar -->
    <aside
      class="fixed left-0 top-0 z-40 h-screen w-64 transition-transform duration-300 border-r border-slate-200 dark:border-slate-800 bg-white/80 dark:bg-slate-900/80 backdrop-blur-xl"
      :class="[isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0']"
    >
      <div class="flex flex-col h-full px-4 py-6">
        <div class="flex items-center justify-between mb-10 px-2">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-xl bg-white/10 flex items-center justify-center shadow-lg backdrop-blur-sm border border-white/20 overflow-hidden p-2">
              <img src="/logo.png" alt="Logo" class="w-full h-full object-contain" />
            </div>
            <span class="text-xl font-bold tracking-tight text-slate-900 dark:text-white">{{ config.public.appName }}</span>
          </div>
          <button @click="isSidebarOpen = false" class="lg:hidden p-2 text-slate-400 hover:text-slate-600">
            <LucideX class="w-5 h-5" />
          </button>
        </div>

        <nav class="flex-1 space-y-1">
          <div v-for="item in menuItems" :key="item.name">
            <!-- Simple Link -->
            <NuxtLink
              v-if="!item.children"
              :to="item.path"
              @click="isSidebarOpen = false"
              class="flex items-center justify-between px-4 py-3 rounded-xl transition-all duration-200 group mb-1"
              :class="[
                $route.path === item.path
                  ? 'bg-[#2D9B7B] text-white shadow-lg shadow-[#2D9B7B]/30'
                  : 'text-slate-300 hover:bg-slate-100/10'
              ]"
            >
              <div class="flex items-center gap-3">
                <component :is="item.icon" class="w-5 h-5" />
                <span class="font-medium text-sm">{{ $t(item.key) }}</span>
              </div>
              <span v-if="item.badge" class="bg-red-500 text-white text-[10px] font-black px-1.5 py-0.5 rounded-full">{{ item.badge }}</span>
            </NuxtLink>

            <!-- Collapsible Menu -->
            <div v-else class="mb-1">
              <button
                @click="toggleSubmenu(item.name)"
                class="w-full flex items-center justify-between px-4 py-3 rounded-xl transition-all duration-200 group text-slate-300 hover:bg-slate-100/10"
              >
                <div class="flex items-center gap-3">
                  <component :is="item.icon" class="w-5 h-5" />
                  <span class="font-medium text-sm">{{ $t(item.key) }}</span>
                </div>
                <div class="flex items-center gap-2">
                  <span v-if="item.badge" class="bg-red-500 text-white text-[10px] font-black px-1.5 py-0.5 rounded-full">{{ item.badge }}</span>
                  <LucideChevronDown 
                    class="w-4 h-4 transition-transform duration-200" 
                    :class="{ 'rotate-180': openSubmenus.includes(item.name) }" 
                  />
                </div>
              </button>
              
              <Transition
                enter-active-class="transition duration-200 ease-out"
                enter-from-class="transform scale-95 opacity-0 -translate-y-2"
                enter-to-class="transform scale-100 opacity-100 translate-y-0"
              >
                <div v-if="openSubmenus.includes(item.name)" class="mt-1 ml-4 pl-4 border-l border-slate-700/50 space-y-1">
                  <NuxtLink
                    v-for="sub in item.children"
                    :key="sub.path"
                    :to="sub.path"
                    @click="isSidebarOpen = false"
                    class="flex items-center justify-between px-4 py-2 text-xs font-medium text-slate-400 hover:text-white transition-colors"
                  >
                    <span>{{ $t(sub.key) }}</span>
                    <span v-if="sub.count" class="text-[10px] text-slate-500 font-bold">{{ sub.count }}</span>
                  </NuxtLink>
                </div>
              </Transition>
            </div>
          </div>
        </nav>

        <div class="mt-auto pt-6 border-t border-slate-200 dark:border-slate-800">
          <div class="flex items-center gap-3 px-2">
            <div class="w-10 h-10 rounded-full bg-primary-100 dark:bg-primary-900/30 flex items-center justify-center text-primary-600 dark:text-primary-400 font-black text-xs">
              {{ auth.user?.full_name?.substring(0, 2).toUpperCase() || 'AD' }}
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-sm font-semibold truncate text-slate-700 dark:text-slate-200">{{ auth.user?.full_name || 'Administrator' }}</p>
              <p class="text-xs text-slate-500 dark:text-slate-400 truncate">{{ auth.user?.email || 'admin@kreatif.com' }}</p>
            </div>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="lg:ml-64 min-h-screen flex flex-col">
      <!-- Navbar -->
      <header class="sticky top-0 z-30 h-16 border-b border-slate-200 dark:border-slate-800 bg-white/80 dark:bg-slate-900/80 backdrop-blur-xl">
        <div class="h-full px-4 lg:px-8 flex items-center justify-between">
          <div class="flex items-center gap-4">
            <button @click="isSidebarOpen = true" class="lg:hidden p-2 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800">
              <LucideMenu class="w-6 h-6" />
            </button>
            <div class="hidden sm:flex items-center gap-2 text-xs font-bold text-slate-400">
              <span>DMS</span>
              <LucideChevronRight class="w-3 h-3" />
              <span class="text-slate-900 dark:text-white">{{ $t(currentPageTitleKey) }}</span>
            </div>
          </div>

          <div class="flex items-center gap-5">
            <div class="relative hidden lg:block">
              <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              <input
                type="text"
                v-model="searchQuery"
                @keyup.enter="handleHeaderSearch"
                :placeholder="$t('layout.navbar.search_placeholder')"
                class="w-80 xl:w-96 pl-12 pr-4 py-2.5 text-sm rounded-xl border border-slate-200 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-950/50 focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-medium"
              />
            </div>
            
            <div class="flex items-center gap-2">
              <Dropdown v-model="showNotifications">
                <template #trigger>
                  <button class="p-2.5 rounded-xl hover:bg-slate-100 dark:hover:bg-slate-800 relative text-slate-500 transition-colors">
                    <LucideBell class="w-5 h-5" />
                    <span v-if="notifications.length > 0" class="absolute top-2.5 right-2.5 w-2 h-2 bg-red-500 rounded-full border-2 border-white dark:border-slate-900"></span>
                  </button>
                </template>
                <div class="px-4 py-3 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-900/50">
                  <span class="text-xs font-black uppercase tracking-widest text-slate-400">{{ $t('layout.navbar.notifications') }}</span>
                  <span class="text-[10px] font-bold text-primary-500 hover:underline cursor-pointer">{{ $t('layout.navbar.mark_all_read') }}</span>
                </div>
                <div class="max-h-80 overflow-y-auto">
                  <div v-for="notif in notifications" :key="notif.id" class="px-4 py-4 hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors border-b border-slate-100 dark:border-slate-800 last:border-0 cursor-pointer group">
                    <div class="flex gap-3">
                      <div :class="`w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0 ${notif.bg}`">
                        <component :is="notif.icon" :class="`w-4 h-4 ${notif.color}`" />
                      </div>
                      <div class="flex-1">
                        <p class="text-sm font-bold text-slate-700 dark:text-slate-200 group-hover:text-primary-500 transition-colors">{{ notif.title }}</p>
                        <p class="text-xs text-slate-400 mt-0.5 line-clamp-2">{{ notif.message }}</p>
                        <p class="text-[10px] text-slate-300 dark:text-slate-500 font-bold mt-1.5 uppercase tracking-tighter">{{ notif.time }}</p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="px-4 py-3 bg-slate-50/50 dark:bg-slate-900/50 text-center border-t border-slate-100 dark:border-slate-800">
                  <button class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 dark:hover:text-slate-200">{{ $t('layout.navbar.view_all_notifications') }}</button>
                </div>
              </Dropdown>
              
              <button @click="navigateTo('/settings')" class="p-2.5 rounded-xl hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-500 transition-colors">
                <LucideSettings class="w-5 h-5" />
              </button>
            </div>

            <div class="w-px h-8 bg-slate-200 dark:bg-slate-800 mx-1 hidden sm:block"></div>
            
            <Dropdown v-model="showUserMenu">
              <template #trigger>
                <button class="flex items-center gap-3 p-1 rounded-full hover:bg-slate-100 dark:hover:bg-slate-800 transition-all">
                  <span class="text-sm font-bold hidden xl:block text-slate-700 dark:text-slate-300">{{ auth.user?.full_name?.split(' ')[0] || 'Admin' }}</span>
                  <div class="w-9 h-9 rounded-full bg-primary-100 dark:bg-primary-900/30 flex items-center justify-center text-primary-600 dark:text-primary-400 font-black text-xs">
                    {{ auth.user?.full_name?.substring(0, 2).toUpperCase() || 'AD' }}
                  </div>
                </button>
              </template>
              
              <div class="px-4 py-4 border-b border-slate-100 dark:border-slate-800 flex items-center gap-3">
                <div class="w-10 h-10 rounded-full bg-primary-100 dark:bg-primary-900/30 flex items-center justify-center text-primary-600 dark:text-primary-400 font-black text-sm">
                  {{ auth.user?.full_name?.substring(0, 2).toUpperCase() || 'AD' }}
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-bold text-slate-900 dark:text-white truncate">{{ auth.user?.full_name || 'Administrator' }}</p>
                  <p class="text-xs text-slate-400 truncate">{{ auth.user?.email || 'admin@kreatif.id' }}</p>
                </div>
              </div>

              <div class="py-1">
                <button @click="navigateTo('/profile')" class="w-full flex items-center gap-3 px-4 py-2.5 text-sm font-bold text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors">
                  <LucideUser class="w-4 h-4" /> {{ $t('layout.navbar.my_profile') }}
                </button>
                <button @click="navigateTo('/settings')" class="w-full flex items-center gap-3 px-4 py-2.5 text-sm font-bold text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors">
                  <LucideSettings class="w-4 h-4" /> {{ $t('layout.navbar.account_settings') }}
                </button>
              </div>

              <div class="border-t border-slate-100 dark:border-slate-800 mt-1 py-1">
                <button @click="handleLogout" class="w-full flex items-center gap-3 px-4 py-2.5 text-sm font-bold text-red-500 hover:bg-red-50 dark:hover:bg-red-900/10 transition-colors">
                  <LucideLogOut class="w-4 h-4" /> {{ $t('layout.navbar.sign_out') }}
                </button>
              </div>
            </Dropdown>
          </div>
        </div>
      </header>

      <!-- Page Content -->
      <main class="flex-1 p-4 lg:p-8 overflow-x-hidden">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { 
  LucideFileStack, 
  LucideLayoutDashboard, 
  LucideFiles, 
  LucideUploadCloud, 
  LucideSettings, 
  LucideMenu, 
  LucideX,
  LucideSearch, 
  LucideBell,
  LucideShieldCheck,
  LucideChevronRight,
  LucideLogOut,
  LucideUser,
  LucideCheckCircle2,
  LucideAlertCircle,
  LucideInfo,
  LucideFileText,
  LucideBarChart3,
  LucideBookOpen,
  LucideChevronDown,
  LucideBox,
  LucideHome,
  LucideRepeat,
  LucideTrash2
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import Dropdown from '~/components/Dropdown.vue'

const auth = useAuthStore()
const config = useRuntimeConfig()
const route = useRoute()
const isSidebarOpen = ref(false)
const searchQuery = ref('')

const showNotifications = ref(false)
const showUserMenu = ref(false)

const handleHeaderSearch = () => {
  if (searchQuery.value.trim()) {
    navigateTo(`/documents?q=${encodeURIComponent(searchQuery.value)}`)
  }
}

const notifications = [
  { id: 1, title: 'Document Approved', message: 'Your request for "Q4 Financial Report" has been approved.', time: '2 mins ago', icon: LucideCheckCircle2, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' },
  { id: 2, title: 'Signature Required', message: 'You have a new loan request that requires your signature.', time: '1 hour ago', icon: LucideAlertCircle, color: 'text-orange-500', bg: 'bg-orange-50 dark:bg-orange-900/20' },
  { id: 3, title: 'System Maintenance', message: 'DMS will be unavailable this Sunday for scheduled updates.', time: '3 hours ago', icon: LucideInfo, color: 'text-blue-500', bg: 'bg-blue-50 dark:bg-blue-900/20' },
]

const openSubmenus = ref([])

const toggleSubmenu = (key) => {
  if (openSubmenus.value.includes(key)) {
    openSubmenus.value = openSubmenus.value.filter(n => n !== key)
  } else {
    openSubmenus.value.push(key)
  }
}

const handleLogout = () => {
  showUserMenu.value = false
  auth.logout()
}

onMounted(() => {
  isSidebarOpen.value = false
})

const allMenuItems = {
  admin: [
    { key: 'layout.menu.dashboard', path: '/dashboard', icon: LucideLayoutDashboard },
    { 
      key: 'layout.menu.doc_reg', 
      icon: LucideFileText,
      children: [
        { key: 'layout.menu.inbound_pre', path: '/registration/pre' },
        { key: 'layout.menu.new_doc', path: '/registration/new' },
        { key: 'layout.menu.legacy_mig', path: '/registration/migration' },
        { key: 'layout.menu.staging', path: '/registration/staging' },
      ]
    },
    { 
      key: 'layout.menu.warehouse', 
      icon: LucideHome,
      children: [
        { key: 'layout.menu.structure', path: '/warehouse/structure' },
        { key: 'layout.menu.occupancy', path: '/warehouse/occupancy' },
        { key: 'layout.menu.zonation', path: '/warehouse/zonation' },
        { key: 'layout.menu.print_qr', path: '/warehouse/labels' },
      ]
    },
    { key: 'layout.menu.search', path: '/documents', icon: LucideSearch },
    { 
      key: 'layout.menu.circulation', 
      icon: LucideRepeat,
      children: [
        { key: 'layout.menu.pickup', path: '/circulation/pickup' },
        { key: 'layout.menu.checkout', path: '/circulation/checkout' },
        { key: 'layout.menu.checkin', path: '/circulation/checkin' },
        { key: 'layout.menu.overdue', path: '/circulation/overdue' },
        { key: 'layout.menu.softcopy', path: '/circulation/softcopy' },
      ]
    },
    { 
      key: 'layout.menu.stock_take', 
      icon: LucideBox,
      children: [
        { key: 'layout.menu.audit_missions', path: '/stock/missions' },
        { key: 'layout.menu.scan_exec', path: '/stock/scan' },
        { key: 'layout.menu.reconciliation', path: '/stock/reconciliation' },
      ]
    },
    { 
      key: 'layout.menu.retention', 
      icon: LucideTrash2,
      children: [
        { key: 'layout.menu.approaching', path: '/retention/approaching' },
        { key: 'layout.menu.history', path: '/retention/history' },
      ]
    },
    { 
      key: 'layout.menu.reports', 
      icon: LucideBarChart3,
      children: [
        { key: 'layout.menu.stats', path: '/reports/statistics' },
        { key: 'layout.menu.capacity', path: '/reports/capacity' },
        { key: 'layout.menu.audit_trail', path: '/reports/audit' },
      ]
    },
    { 
      key: 'layout.menu.config', 
      icon: LucideSettings,
      children: [
        { key: 'layout.menu.company', path: '/config/company' },
        { key: 'layout.menu.dept', path: '/config/department' },
        { key: 'layout.menu.params', path: '/config/params' },
      ]
    },
  ],
  user: [
    { key: 'layout.menu.dashboard', path: '/dashboard', icon: LucideLayoutDashboard },
    { key: 'layout.menu.search', path: '/documents', icon: LucideSearch },
    { key: 'layout.menu.submit', path: '/documents/upload', icon: LucideUploadCloud },
    { key: 'layout.menu.loans', path: '/loans', icon: LucideFileText },
    { key: 'layout.menu.tracking', path: '/tracking', icon: LucideBarChart3 },
  ],
  manager: [
    { key: 'layout.menu.dashboard', path: '/dashboard', icon: LucideLayoutDashboard },
    { 
      key: 'layout.menu.approvals', 
      icon: LucideShieldCheck, 
      badge: 5,
      children: [
        { key: 'layout.menu.submissions', path: '/approvals/submissions', count: 2 },
        { key: 'layout.menu.loans', path: '/approvals/loans', count: 2 },
        { key: 'layout.menu.extensions', path: '/approvals/extensions', count: 1 },
      ]
    },
    { key: 'layout.menu.search', path: '/documents', icon: LucideSearch },
    { 
      key: 'layout.menu.loans', 
      icon: LucideBookOpen,
      children: [
        { key: 'layout.menu.fast_track', path: '/loans/fast-track' },
        { key: 'layout.menu.request_loan', path: '/loans/request' },
        { key: 'layout.menu.my_loans', path: '/loans/my' },
        { key: 'layout.menu.softcopy', path: '/loans/softcopy' },
        { key: 'layout.menu.loan_history', path: '/loans/history' },
      ]
    },
    { 
      key: 'layout.menu.submissions', 
      icon: LucideFileStack,
      children: [
        { key: 'layout.menu.submit', path: '/documents/upload' },
        { key: 'layout.menu.submission_status', path: '/documents/status' },
      ]
    },
    { key: 'layout.menu.notifications', path: '/notifications', icon: LucideBell, badge: 8 },
  ],
  doc_controller: [
    { key: 'layout.menu.dashboard', path: '/dashboard', icon: LucideLayoutDashboard },
    { 
      key: 'layout.menu.approvals', 
      icon: LucideShieldCheck, 
      badge: 5,
      children: [
        { key: 'layout.menu.loans', path: '/approvals/loans' },
        { key: 'layout.menu.softcopy', path: '/approvals/softcopy' },
        { key: 'layout.menu.extensions', path: '/approvals/extensions' },
        { key: 'layout.menu.disposal', path: '/approvals/disposal' },
      ]
    },
    { key: 'layout.menu.search', path: '/documents', icon: LucideSearch },
    { 
      key: 'layout.menu.loans', 
      icon: LucideBookOpen,
      children: [
        { key: 'layout.menu.request_loan', path: '/loans/request' },
        { key: 'layout.menu.my_loans', path: '/loans/my' },
        { key: 'layout.menu.softcopy', path: '/loans/softcopy' },
        { key: 'layout.menu.loan_history', path: '/loans/history' },
      ]
    },
    { 
      key: 'layout.menu.reports', 
      icon: LucideBarChart3,
      children: [
        { key: 'layout.menu.stats', path: '/reports/statistics' },
        { key: 'layout.menu.circulation_report', path: '/reports/circulation' },
        { key: 'layout.menu.overdue_monitoring', path: '/reports/overdue' },
      ]
    },
    { 
      key: 'layout.menu.stock_take', 
      icon: LucideBox,
      children: [
        { key: 'layout.menu.reconciliation', path: '/stock/reconciliation' },
      ]
    },
    { key: 'layout.menu.notifications', path: '/notifications', icon: LucideBell, badge: 8 },
  ],
  manager_doc_controller: [
    { key: 'layout.menu.dashboard', path: '/dashboard', icon: LucideLayoutDashboard },
    { 
      key: 'layout.menu.doc_reg', 
      icon: LucideFileText,
      children: [
        { key: 'layout.menu.inbound_pre', path: '/registration/pre' },
        { key: 'layout.menu.new_doc', path: '/registration/new' },
        { key: 'layout.menu.legacy_mig', path: '/registration/migration' },
        { key: 'layout.menu.staging', path: '/registration/staging' },
      ]
    },
    { 
      key: 'layout.menu.warehouse', 
      icon: LucideHome,
      children: [
        { key: 'layout.menu.structure', path: '/warehouse/structure' },
        { key: 'layout.menu.occupancy', path: '/warehouse/occupancy' },
        { key: 'layout.menu.zonation', path: '/warehouse/zonation' },
        { key: 'layout.menu.print_qr', path: '/warehouse/labels' },
      ]
    },
    { key: 'layout.menu.search', path: '/documents', icon: LucideSearch },
    { 
      key: 'layout.menu.circulation', 
      icon: LucideRepeat,
      children: [
        { key: 'layout.menu.pickup', path: '/circulation/pickup' },
        { key: 'layout.menu.checkout', path: '/circulation/checkout' },
        { key: 'layout.menu.checkin', path: '/circulation/checkin' },
        { key: 'layout.menu.overdue', path: '/circulation/overdue' },
        { key: 'layout.menu.softcopy', path: '/circulation/softcopy' },
      ]
    },
    { 
      key: 'layout.menu.stock_take', 
      icon: LucideBox,
      children: [
        { key: 'layout.menu.audit_missions', path: '/stock/missions' },
        { key: 'layout.menu.scan_exec', path: '/stock/scan' },
        { key: 'layout.menu.reconciliation', path: '/stock/reconciliation' },
      ]
    },
    { 
      key: 'layout.menu.retention', 
      icon: LucideTrash2,
      children: [
        { key: 'layout.menu.approaching', path: '/retention/approaching' },
        { key: 'layout.menu.history', path: '/retention/history' },
      ]
    },
    { 
      key: 'layout.menu.reports', 
      icon: LucideBarChart3,
      children: [
        { key: 'layout.menu.stats', path: '/reports/statistics' },
        { key: 'layout.menu.capacity', path: '/reports/capacity' },
        { key: 'layout.menu.audit_trail', path: '/reports/audit' },
      ]
    },
    { 
      key: 'layout.menu.config', 
      icon: LucideSettings,
      children: [
        { key: 'layout.menu.company', path: '/config/company' },
        { key: 'layout.menu.dept', path: '/config/department' },
        { key: 'layout.menu.params', path: '/config/params' },
      ]
    },
  ]
}

const menuItems = computed(() => {
  const role = auth.user?.role || 'user'
  return allMenuItems[role] || allMenuItems.user
})

const currentPageTitleKey = computed(() => {
  const current = menuItems.value.find(m => m.path === route.path)
  if (current) return current.key
  
  // Check children
  for (const item of menuItems.value) {
    if (item.children) {
      const child = item.children.find(c => c.path === route.path)
      if (child) return child.key
    }
  }
  
  return 'layout.menu.dashboard'
})

// Auto-close sidebar on desktop resize
if (import.meta.client) {
  window.addEventListener('resize', () => {
    if (window.innerWidth >= 1024) {
      isSidebarOpen.value = false
    }
  })
}
</script>
