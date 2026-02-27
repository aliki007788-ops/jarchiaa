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
      // ذخیره در localStorage
      localStorage.setItem('token', '123456')
      localStorage.setItem('userRole', 'user')
      localStorage.setItem('userName', 'کاربر')

      // برو به خانه
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
}
.login-box h1 {
  color: gold;
  text-align: center;
}
.login-box input {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  background: #333;
  border: 1px solid gold;
  color: white;
}
.login-box button {
  width: 100%;
  padding: 10px;
  background: gold;
  border: none;
  cursor: pointer;
}
</style>
