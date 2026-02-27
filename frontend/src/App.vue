<template>
  <div id="app" class="mobile-view">
    <!-- ذرات درخشان (فقط در لاگین) -->
    <div class="particles-container" :class="{active: $route.path === '/login'}">
      <div v-for="n in 200" :key="n" class="particle" :class="getParticleClass()"
           :style="getParticleStyle()"></div>
    </div>
    
    <!-- هدر - فقط وقتی لاگین هستی نمایش بده -->
    <header class="header" v-if="isLoggedIn">
      <div class="menu-btn" @click="toggleDrawer">
        <span></span><span></span><span></span>
      </div>
      <div class="logo" @click="switchRole">
        <i class="fas fa-bullhorn"></i> جارچیا
      </div>
      <div class="header-icons">
        <i class="fas fa-adjust" @click="toggleTheme"></i>
        <i class="fas fa-bell" @click="showNotifications"></i>
        <i class="fas fa-search" @click="showSearch"></i>
      </div>
    </header>

    <!-- صفحه اصلی -->
    <router-view></router-view>

    <!-- ناوبری پایین (فقط در صفحات اصلی و وقتی لاگین هستی) -->
    <BottomNav v-if="!isLoginPage && isLoggedIn" />
  </div>
</template>

<script>
import BottomNav from './components/BottomNav.vue'
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from './store'

export default {
  name: 'App',
  components: { BottomNav },
  setup() {
    const route = useRoute()
    const router = useRouter()
    const store = useStore()
    
    const isLoginPage = computed(() => route.path === '/login')
    
    // بررسی لاگین بودن کاربر
    const isLoggedIn = computed(() => {
      return !!localStorage.getItem('token')
    })
    
    const toggleDrawer = () => {
      if (isLoggedIn.value) {
        store.commit('toggleDrawer')
      }
    }
    
    const toggleTheme = () => store.commit('toggleTheme')
    
    const switchRole = () => {
      if (isLoggedIn.value) {
        store.dispatch('switchRole')
      }
    }
    
    const showNotifications = () => {
      if (isLoggedIn.value) {
        alert('۳ اعلان جدید')
      }
    }
    
    const showSearch = () => {
      if (isLoggedIn.value) {
        alert('جستجو')
      }
    }

    // ذرات درخشان
    const getParticleClass = () => Math.random() > 0.5 ? 'gold' : 'white'
    const getParticleStyle = () => ({
      width: Math.random() * 3 + 1 + 'px',
      height: Math.random() * 3 + 1 + 'px',
      left: Math.random() * 100 + '%',
      animationDuration: Math.random() * 6 + 4 + 's',
      animationDelay: -(Math.random() * 10) + 's'
    })

    // اگه توکن نبود و صفحه لاگین نیست، بفرست به لاگین
    if (!isLoggedIn.value && !isLoginPage.value) {
      router.push('/login')
    }

    return { 
      toggleDrawer, 
      toggleTheme, 
      switchRole, 
      showNotifications, 
      showSearch, 
      isLoginPage,
      isLoggedIn,
      getParticleClass, 
      getParticleStyle
    }
  }
}
</script>

<style src="./assets/css/main.css"></style>
