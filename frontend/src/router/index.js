import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../pages/HomePage.vue'
import LoginPage from '../pages/LoginPage.vue'
import VitrinPage from '../pages/VitrinPage.vue'
import BusinessPage from '../pages/BusinessPage.vue'
import AdminPage from '../pages/AdminPage.vue'

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/vitrin', name: 'Vitrin', component: VitrinPage },
  { path: '/business', name: 'Business', component: BusinessPage },
  { path: '/business/:id', name: 'BusinessDetail', component: BusinessPage },
  { path: '/admin', name: 'Admin', component: AdminPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
