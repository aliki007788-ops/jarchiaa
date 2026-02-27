<template>
  <div class="login-box">
    <h1>جارچیا</h1>
    <input v-model="phone" placeholder="شماره موبایل">
    <button @click="login">ورود</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      phone: ''
    }
  },
  methods: {
    login() {
      // ذخیره اطلاعات
      localStorage.setItem('token', '123456')
      localStorage.setItem('userRole', 'user')
      localStorage.setItem('userName', 'کاربر')
      
      // انتقال به صفحه اصلی
      window.location.href = '/'
    }
  }
}
</script>

<style>
.login-box {
  max-width: 400px;
  margin: 100px auto;
  padding: 20px;
  text-align: center;
}
.login-box h1 {
  color: #D4AF37;
  margin-bottom: 30px;
}
.login-box input {
  width: 100%;
  padding: 12px;
  margin-bottom: 20px;
  background: #1A1A1A;
  border: 1px solid #D4AF37;
  color: white;
  border-radius: 8px;
}
.login-box button {
  width: 100%;
  padding: 12px;
  background: #D4AF37;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
}
</style>
