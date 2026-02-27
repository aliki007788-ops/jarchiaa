 <template>
  <div id="app" class="mobile-view">
    <!-- ذرات درخشان (فقط در لاگین) -->
    <div class="particles-container" :class="{active: $route.path === '/login'}">
      <div v-for="n in 200" :key="n" class="particle" :class="getParticleClass()"
           :style="getParticleStyle()"></div>
    </div>
    
    <!-- هدر -->
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

    <!-- صفحه اصلی -->
    <router-view></router-view>

    <!-- ناوبری پایین (فقط در صفحات اصلی) -->
    <BottomNav v-if="!isLoginPage" />
  </div>
</template>

<script>
import BottomNav from './components/BottomNav.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from './store'

export default {
  name: 'App',
  components: { BottomNav },
  setup() {
    const route = useRoute()
    const store = useStore()
    
    const isLoginPage = computed(() => route.path === '/login')
    
    const toggleDrawer = () => store.commit('toggleDrawer')
    const toggleTheme = () => store.commit('toggleTheme')
    const switchRole = () => store.dispatch('switchRole')
    const showNotifications = () => alert('۳ اعلان جدید')
    const showSearch = () => alert('جستجو')

    // ذرات درخشان
    const getParticleClass = () => Math.random() > 0.5 ? 'gold' : 'white'
    const getParticleStyle = () => ({
      width: Math.random() * 3 + 1 + 'px',
      height: Math.random() * 3 + 1 + 'px',
      left: Math.random() * 100 + '%',
      animationDuration: Math.random() * 6 + 4 + 's',
      animationDelay: -(Math.random() * 10) + 's'
    })

    return { 
      toggleDrawer, toggleTheme, switchRole, 
      showNotifications, showSearch, isLoginPage,
      getParticleClass, getParticleStyle
    }
  }
}
</script>

<style src="./assets/css/main.css"></style>
