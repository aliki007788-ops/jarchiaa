<template>
  <div class="drawer" :class="{ open: isOpen }">
    <!-- هدر پروفایل -->
    <div class="drawer-profile">
      <div class="profile-avatar"></div>
      <div class="profile-info">
        <div class="profile-name">{{ userName }}</div>
        <div class="profile-role">{{ userRoleText }}</div>
      </div>
    </div>

    <!-- آیتم‌های منو -->
    <div class="drawer-items">
      <!-- خانه - برای همه -->
      <div class="drawer-item" @click="navigateTo('/')">
        <i class="fas fa-home"></i> خانه
      </div>

      <!-- ویترین - برای همه -->
      <div class="drawer-item" @click="navigateTo('/vitrin')">
        <i class="fas fa-store-alt"></i> ویترین
      </div>

      <!-- ===== بخش کسب‌وکار (فقط برای نقش business) ===== -->
      <template v-if="userRole === 'business'">
        <div class="drawer-divider"></div>
        <div class="drawer-item" @click="navigateTo('/business')">
          <i class="fas fa-store"></i> داشبورد کسب‌وکار
        </div>
        <div class="drawer-item" @click="navigateTo('/business/plans')">
          <i class="fas fa-crown"></i> پلن‌های آگهی
        </div>
      </template>

      <!-- ===== بخش مدیریت (فقط برای نقش admin) ===== -->
      <template v-if="userRole === 'admin'">
        <div class="drawer-divider"></div>
        <div class="drawer-item" @click="navigateTo('/admin')">
          <i class="fas fa-crown"></i> داشبورد مدیریت
        </div>
        <div class="drawer-item" @click="navigateTo('/admin/modules')">
          <i class="fas fa-sliders-h"></i> مدیریت ماژول‌ها
        </div>
      </template>

      <div class="drawer-divider"></div>

      <!-- پروفایل - برای همه -->
      <div class="drawer-item" @click="navigateTo('/profile')">
        <i class="fas fa-user"></i> پروفایل من
      </div>

      <!-- کیف پول - برای همه -->
      <div class="drawer-item" @click="navigateTo('/wallet')">
        <i class="fas fa-wallet"></i> کیف پول
      </div>

      <!-- خروج - برای همه -->
      <div class="drawer-item" @click="logout">
        <i class="fas fa-sign-out-alt"></i> خروج
      </div>
    </div>
  </div>
</template>
