import { createRouter, createWebHistory } from 'vue-router'

// ========== پنل کاربری (User Panel) ==========
import LoginPage from '../modules/user/pages/LoginPage.vue'
import HomePage from '../modules/user/pages/HomePage.vue'
import ProfilePage from '../modules/user/pages/ProfilePage.vue'
import WalletPage from '../modules/user/pages/WalletPage.vue'
import FavoritesPage from '../modules/user/pages/FavoritesPage.vue'
import AppointmentsPage from '../modules/user/pages/AppointmentsPage.vue'

// ========== ویترین (Vitrin) ==========
import VitrinPage from '../modules/vitrin/pages/VitrinPage.vue'
import AdDetailPage from '../modules/vitrin/pages/AdDetailPage.vue'

// ========== پنل کسب‌وکار (Business Panel) ==========
import BusinessDashboard from '../modules/business/pages/BusinessDashboard.vue'
import BusinessPage from '../modules/business/pages/BusinessPage.vue'
import PlansPage from '../modules/business/pages/PlansPage.vue'
import CreateAdPage from '../modules/business/pages/CreateAdPage.vue'
import BusinessStatsPage from '../modules/business/pages/BusinessStatsPage.vue'
import BusinessAppointmentsPage from '../modules/business/pages/BusinessAppointmentsPage.vue'

// ========== پنل مدیریت (Admin Panel) ==========
import AdminDashboard from '../modules/admin/pages/AdminDashboard.vue'
import ModulesPage from '../modules/admin/pages/ModulesPage.vue'
import UsersManagementPage from '../modules/admin/pages/UsersManagementPage.vue'
import BusinessesManagementPage from '../modules/admin/pages/BusinessesManagementPage.vue'
import PlansManagementPage from '../modules/admin/pages/PlansManagementPage.vue'
import ReportsPage from '../modules/admin/pages/ReportsPage.vue'

// ========== تعریف مسیرها ==========
const routes = [
  // ------------------------------------------------------
  // مسیرهای عمومی (نیاز به لاگین ندارند)
  // ------------------------------------------------------
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
    meta: { requiresAuth: false }
  },

  // ------------------------------------------------------
  // مسیرهای پنل کاربری (نیاز به لاگین دارند)
  // ------------------------------------------------------
  {
    path: '/',
    name: 'Home',
    component: HomePage,
    meta: { requiresAuth: true, allowedRoles: ['user', 'business', 'admin'] }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: ProfilePage,
    meta: { requiresAuth: true, allowedRoles: ['user', 'business', 'admin'] }
  },
  {
    path: '/wallet',
    name: 'Wallet',
    component: WalletPage,
    meta: { requiresAuth: true, allowedRoles: ['user', 'business'] }
  },
  {
    path: '/favorites',
    name: 'Favorites',
    component: FavoritesPage,
    meta: { requiresAuth: true, allowedRoles: ['user'] }
  },
  {
    path: '/appointments',
    name: 'Appointments',
    component: AppointmentsPage,
    meta: { requiresAuth: true, allowedRoles: ['user'] }
  },

  // ------------------------------------------------------
  // مسیرهای ویترین (نیاز به لاگین دارند)
  // ------------------------------------------------------
  {
    path: '/vitrin',
    name: 'Vitrin',
    component: VitrinPage,
    meta: { requiresAuth: true, allowedRoles: ['user', 'business', 'admin'] }
  },
  {
    path: '/ad/:id',
    name: 'AdDetail',
    component: AdDetailPage,
    meta: { requiresAuth: true, allowedRoles: ['user', 'business', 'admin'] },
    props: true // ارسال id به کامپوننت به عنوان prop
  },

  // ------------------------------------------------------
  // مسیرهای پنل کسب‌وکار (فقط نقش business)
  // ------------------------------------------------------
  {
    path: '/business',
    name: 'BusinessDashboard',
    component: BusinessDashboard,
    meta: { requiresAuth: true, allowedRoles: ['business', 'admin'] }
  },
  {
    path: '/business/:id',
    name: 'BusinessPage',
    component: BusinessPage,
    meta: { requiresAuth: true, allowedRoles: ['user', 'business', 'admin'] },
    props: true
  },
  {
    path: '/business/plans',
    name: 'Plans',
    component: PlansPage,
    meta: { requiresAuth: true, allowedRoles: ['business'] }
  },
  {
    path: '/business/create-ad',
    name: 'CreateAd',
    component: CreateAdPage,
    meta: { requiresAuth: true, allowedRoles: ['business'] }
  },
  {
    path: '/business/stats',
    name: 'BusinessStats',
    component: BusinessStatsPage,
    meta: { requiresAuth: true, allowedRoles: ['business'] }
  },
  {
    path: '/business/appointments',
    name: 'BusinessAppointments',
    component: BusinessAppointmentsPage,
    meta: { requiresAuth: true, allowedRoles: ['business'] }
  },

  // ------------------------------------------------------
  // مسیرهای پنل مدیریت (فقط نقش admin)
  // ------------------------------------------------------
  {
    path: '/admin',
    name: 'AdminDashboard',
    component: AdminDashboard,
    meta: { requiresAuth: true, allowedRoles: ['admin'] }
  },
  {
    path: '/admin/modules',
    name: 'Modules',
    component: ModulesPage,
    meta: { requiresAuth: true, allowedRoles: ['admin'] }
  },
  {
    path: '/admin/users',
    name: 'UsersManagement',
    component: UsersManagementPage,
    meta: { requiresAuth: true, allowedRoles: ['admin'] }
  },
  {
    path: '/admin/businesses',
    name: 'BusinessesManagement',
    component: BusinessesManagementPage,
    meta: { requiresAuth: true, allowedRoles: ['admin'] }
  },
  {
    path: '/admin/plans',
    name: 'PlansManagement',
    component: PlansManagementPage,
    meta: { requiresAuth: true, allowedRoles: ['admin'] }
  },
  {
    path: '/admin/reports',
    name: 'Reports',
    component: ReportsPage,
    meta: { requiresAuth: true, allowedRoles: ['admin'] }
  },

  // ------------------------------------------------------
  // مسیر پیش‌فرض (404)
  // ------------------------------------------------------
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

// ========== ایجاد نمونه روتر ==========
const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// ========== گارد امنیتی (Navigation Guard) ==========
router.beforeEach((to, from, next) => {
  // گرفتن توکن و نقش کاربر از localStorage
  const token = localStorage.getItem('token')
  const userRole = localStorage.getItem('userRole') || 'guest'
  
  // اگر صفحه نیاز به لاگین دارد ولی توکن نیست
  if (to.meta.requiresAuth && !token) {
    next({
      path: '/login',
      query: { redirect: to.fullPath } // ذخیره مسیر برای برگشت بعد از لاگین
    })
    return
  }
  
  // اگر صفحه نیاز به نقش خاصی دارد
  if (to.meta.allowedRoles && !to.meta.allowedRoles.includes(userRole)) {
    // بر اساس نقش فعلی، صفحه مناسب رو نشون بده
    if (userRole === 'admin') {
      next('/admin')
    } else if (userRole === 'business') {
      next('/business')
    } else {
      next('/')
    }
    return
  }
  
  // اگر همه چیز اوکی بود، برو به صفحه مورد نظر
  next()
})

// ========== export روتر ==========
export default router
