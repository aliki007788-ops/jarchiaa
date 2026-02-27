<template>
  <div class="drawer" :class="{ open: isOpen }">
    <!-- هدر پروفایل -->
    <div class="drawer-profile">
      <div class="profile-avatar"></div>
      <div class="profile-info">
        <div class="profile-name">{{ userName || 'کاربر' }}</div>
        <div class="profile-role">{{ userRoleText }}</div>
      </div>
    </div>

    <!-- آیتم‌های منو -->
    <div class="drawer-items">
      <!-- خانه - برای همه -->
      <div class="drawer-item" @click="navigateTo('/')">
        <i class="fas fa-home"></i>
        <span>خانه</span>
      </div>

      <!-- ویترین - برای همه -->
      <div class="drawer-item" @click="navigateTo('/vitrin')">
        <i class="fas fa-store-alt"></i>
        <span>ویترین</span>
      </div>

      <!-- ===== بخش کسب‌وکار (فقط برای نقش business) ===== -->
      <template v-if="userRole === 'business'">
        <div class="drawer-divider"></div>
        <div class="drawer-item" @click="navigateTo('/business')">
          <i class="fas fa-store"></i>
          <span>داشبورد کسب‌وکار</span>
        </div>
        <div class="drawer-item" @click="navigateTo('/business/plans')">
          <i class="fas fa-crown"></i>
          <span>پلن‌های آگهی</span>
        </div>
      </template>

      <!-- ===== بخش مدیریت (فقط برای نقش admin) ===== -->
      <template v-if="userRole === 'admin'">
        <div class="drawer-divider"></div>
        <div class="drawer-item" @click="navigateTo('/admin')">
          <i class="fas fa-crown"></i>
          <span>داشبورد مدیریت</span>
        </div>
        <div class="drawer-item" @click="navigateTo('/admin/modules')">
          <i class="fas fa-sliders-h"></i>
          <span>مدیریت ماژول‌ها</span>
        </div>
      </template>

      <div class="drawer-divider"></div>

      <!-- پروفایل - برای همه -->
      <div class="drawer-item" @click="navigateTo('/profile')">
        <i class="fas fa-user"></i>
        <span>پروفایل من</span>
      </div>

      <!-- کیف پول - برای همه -->
      <div class="drawer-item" @click="navigateTo('/wallet')">
        <i class="fas fa-wallet"></i>
        <span>کیف پول</span>
      </div>

      <!-- خروج - برای همه -->
      <div class="drawer-item" @click="logout">
        <i class="fas fa-sign-out-alt"></i>
        <span>خروج</span>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

export default {
  props: {
    isOpen: {
      type: Boolean,
      default: false
    }
  },
  emits: ['close'],
  setup(props, { emit }) {
    const router = useRouter()
    
    // دریافت نقش کاربر از localStorage با رفرش لحظه‌ای
    const userRole = ref('guest')
    const userName = ref('کاربر')
    
    // آپدیت نقش هر بار که منو باز میشه
    onMounted(() => {
      updateUserInfo()
    })
    
    const updateUserInfo = () => {
      userRole.value = localStorage.getItem('userRole') || 'guest'
      userName.value = localStorage.getItem('userName') || 'کاربر'
      console.log('🔄 نقش کاربر آپدیت شد:', userRole.value)
    }
    
    const userRoleText = computed(() => {
      switch(userRole.value) {
        case 'admin': return 'مدیر سیستم'
        case 'business': return 'کسب‌وکار'
        case 'user': return 'کاربر عادی'
        default: return 'مهمان'
      }
    })

    const navigateTo = (path) => {
      router.push(path)
      emit('close')
    }

    const logout = () => {
      localStorage.removeItem('token')
      localStorage.removeItem('userRole')
      localStorage.removeItem('userName')
      router.push('/login')
      emit('close')
    }

    return {
      userRole,
      userName,
      userRoleText,
      navigateTo,
      logout
    }
  }
}
</script>

<style scoped>
.drawer {
  position: fixed;
  top: 0;
  right: -300px;
  width: 280px;
  height: 100vh;
  background: #1A1A1A;
  z-index: 1000;
  transition: right 0.3s ease;
  border-left: 1px solid #D4AF37;
  overflow-y: auto;
}

.drawer.open {
  right: 0;
}

.drawer-profile {
  padding: 30px 20px 20px 20px;
  background: linear-gradient(135deg, #D4AF37, #B8860B);
  border-bottom: 1px solid #2A2A2A;
}

.profile-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: #0A0A0A;
  border: 3px solid white;
  margin-bottom: 10px;
  background-image: url('/avatar-default.png');
  background-size: cover;
}

.profile-name {
  font-size: 18px;
  font-weight: 700;
  color: #0A0A0A;
}

.profile-role {
  font-size: 12px;
  color: #0A0A0A;
  opacity: 0.8;
  background: rgba(255,255,255,0.2);
  display: inline-block;
  padding: 2px 10px;
  border-radius: 20px;
}

.drawer-items {
  padding: 20px;
}

.drawer-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 15px;
  color: #D4AF37;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
  margin-bottom: 4px;
}

.drawer-item:hover {
  background: #2A2A2A;
  transform: translateX(-5px);
}

.drawer-item i {
  width: 20px;
  font-size: 18px;
}

.drawer-divider {
  height: 1px;
  background: #2A2A2A;
  margin: 15px 0;
}
</style>
