import { createRouter, createWebHistory } from 'vue-router'

// صفحات
import LoginPage from '../modules/user/pages/LoginPage.vue'
import HomePage from '../modules/user/pages/HomePage.vue'
import BusinessDashboard from '../modules/business/pages/BusinessDashboard.vue'
import AdminDashboard from '../modules/admin/pages/AdminDashboard.vue'

const routes = [
  { path: '/login', component: LoginPage },
  { path: '/', component: HomePage, meta: { requiresAuth: true } },
  { 
    path: '/business', 
    component: BusinessDashboard, 
    meta: { requiresAuth: true, role: 'business' }  // 👈 فقط کسب‌وکار
  },
  { 
    path: '/admin', 
    component: AdminDashboard, 
    meta: { requiresAuth: true, role: 'admin' }     // 👈 فقط مدیر
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// ===== گارد امنیتی =====
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  const userRole = localStorage.getItem('userRole') // 'user', 'business', 'admin'
  
  // اگر صفحه نیاز به لاگین دارد و توکن نیست
  if (to.meta.requiresAuth && !token) {
    next('/login')
    return
  }
  
  // اگر صفحه نیاز به نقش خاصی دارد
  if (to.meta.role && to.meta.role !== userRole) {
    // کاربر عادی به صفحه کسب‌وکار نره
    if (userRole === 'user') {
      next('/') // برگرد به خانه
    } 
    // کسب‌وکار به صفحه مدیریت نره
    else if (userRole === 'business' && to.meta.role === 'admin') {
      next('/business') // برگرد به داشبوردش
    }
    else {
      next('/')
    }
    return
  }
  
  next()
})

export default router
