<template>
  <div class="login-container">
    <div class="login-logo">جارچیا</div>
    <div class="input-box">
      <span class="flag-icon">🇮🇷</span>
      <span class="prefix">+98</span>
      <input type="tel" id="phone" placeholder="9120000000">
    </div>
    <button class="main-btn" @click="submit">ورود</button>
  </div>
</template>

<script>
export default {
  name: 'LoginPage',
  methods: {
    submit() {
      const phone = document.getElementById('phone').value
      console.log('Login with:', phone)
    }
  }
}
</script>

<style scoped>
.login-container {
  padding: 40px 20px;
  text-align: center;
}
.login-logo {
  font-size: 32px;
  color: #D4AF37;
  margin-bottom: 40px;
}
.input-box {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-bottom: 20px;
}
.input-box input {
  padding: 12px;
  border: 1px solid #D4AF37;
  border-radius: 8px;
  background: #1A1A1A;
  color: #F5E6D3;
  text-align: center;
}
.main-btn {
  width: 100%;
  padding: 14px;
  background: #D4AF37;
  color: #0A0A0A;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
}
</style>
