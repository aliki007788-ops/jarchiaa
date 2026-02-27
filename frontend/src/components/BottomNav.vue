<template>
  <nav class="bottom-nav">
    <div class="nav-item" :class="{active: $route.path === '/'}" @click="goTo('/')">
      <i class="fas fa-home"></i><span>خانه</span>
    </div>
    <div class="nav-item" :class="{active: $route.path === '/vitrin'}" @click="goTo('/vitrin')">
      <i class="fas fa-store-alt"></i><span>ویترین</span>
    </div>
    <div class="nav-item nav-item-special" @click="quickAction">
      <div class="special-button"><i class="fas fa-plus"></i></div>
      <span class="nav-label-special">سریع</span>
    </div>
    <div class="nav-item" :class="{active: $route.path === '/business'}" @click="goTo('/business')">
      <i class="fas fa-store"></i><span>کسب‌وکار</span>
    </div>
    <div class="nav-item" :class="{active: $route.path === '/admin'}" @click="goTo('/admin')">
      <i class="fas fa-crown"></i><span>مدیریت</span>
    </div>
  </nav>
</template>

<script>
import { useRouter } from 'vue-router'
import { useStore } from '../store'

export default {
  setup() {
    const router = useRouter()
    const store = useStore()
    
    const goTo = (path) => router.push(path)
    const quickAction = () => {
      const role = store.state.currentRole
      if (role === 'business') alert('ایجاد آگهی جدید')
      else if (role === 'admin') alert('دسترسی سریع به پنل مدیریت')
      else alert('جستجوی سریع')
    }

    return { goTo, quickAction }
  }
}
</script>
