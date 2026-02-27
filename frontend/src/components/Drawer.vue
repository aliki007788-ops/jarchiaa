<template>
  <div class="drawer-items">
    <!-- آیتم‌های عمومی -->
    <div class="drawer-item" @click="navigate('/')">
      <i class="fas fa-home"></i> خانه
    </div>
    
    <!-- آیتم‌های مخصوص کسب‌وکار -->
    <div v-if="userRole === 'business'" class="drawer-item" @click="navigate('/business')">
      <i class="fas fa-store"></i> پنل کسب‌وکار
    </div>
    
    <!-- آیتم‌های مخصوص مدیر -->
    <div v-if="userRole === 'admin'" class="drawer-item" @click="navigate('/admin')">
      <i class="fas fa-crown"></i> پنل مدیریت
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useRouter } from 'vue-router'

export default {
  setup() {
    const router = useRouter()
    const userRole = computed(() => localStorage.getItem('userRole'))
    
    const navigate = (path) => {
      router.push(path)
      closeDrawer()
    }
    
    return { userRole, navigate }
  }
}
</script>
