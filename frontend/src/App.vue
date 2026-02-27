<template>
  <div id="app" class="mobile-view">
    <!-- منو و هدر -->
    <header class="header">
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

    <!-- صفحه اصلی (از گراک) -->
    <router-view></router-view>

    <!-- ناوبری پایین -->
    <BottomNav />
  </div>
</template>

<script>
import BottomNav from './components/BottomNav.vue'
import { useStore } from './store'

export default {
  components: { BottomNav },
  setup() {
    const store = useStore()
    
    const toggleDrawer = () => store.commit('toggleDrawer')
    const toggleTheme = () => store.commit('toggleTheme')
    const switchRole = () => store.dispatch('switchRole')
    const showNotifications = () => alert('۳ اعلان جدید')
    const showSearch = () => alert('جستجو')

    return { toggleDrawer, toggleTheme, switchRole, showNotifications, showSearch }
  }
}
</script>
