<template>
  <div class="login-container">
    <div class="login-logo">جارچیا</div>
    
    <!-- مرحله ۱: شماره تلفن -->
    <div v-if="step === 'phone'">
      <div class="input-box">
        <span class="flag-icon">🇮🇷</span>
        <span class="prefix">+98</span>
        <input 
          type="tel" 
          v-model="phone" 
          placeholder="9120000000" 
          maxlength="10"
        >
      </div>
      <button class="main-btn" @click="sendOTP">ارسال کد</button>
    </div>
    
    <!-- مرحله ۲: کد OTP -->
    <div v-else>
      <div class="otp-container">
        <input v-for="i in 6" :key="i" type="text" maxlength="1" class="otp-input" v-model="otp[i-1]">
      </div>
      <button class="main-btn" @click="verifyOTP">ورود</button>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

export default {
  setup() {
    const router = useRouter()
    const step = ref('phone')
    const phone = ref('')
    const otp = ref(['', '', '', '', '', ''])

    const sendOTP = () => {
      if (phone.value.length === 10) {
        step.value = 'otp'
      }
    }

    const verifyOTP = () => {
      // ذخیره در localStorage
      localStorage.setItem('token', '123456')
      localStorage.setItem('userRole', 'user')
      localStorage.setItem('userName', 'کاربر تست')
      
      // رفرش کن! (مهم)
      window.location.href = '/'
    }

    return { step, phone, otp, sendOTP, verifyOTP }
  }
}
</script>

<style>
.login-container {
  padding: 20px;
}
.login-logo {
  font-size: 3rem;
  text-align: center;
  margin: 50px 0;
  color: #D4AF37;
}
.input-box input {
  width: 100%;
  padding: 15px;
  border-radius: 10px;
  background: #1A1A1A;
  border: 1px solid #D4AF37;
  color: white;
}
.main-btn {
  width: 100%;
  padding: 15px;
  background: #D4AF37;
  border: none;
  border-radius: 10px;
  font-size: 1.2rem;
  cursor: pointer;
  margin-top: 20px;
}
.otp-container {
  display: flex;
  gap: 10px;
  justify-content: center;
  margin: 30px 0;
}
.otp-input {
  width: 50px;
  height: 60px;
  text-align: center;
  background: #1A1A1A;
  border: 1px solid #D4AF37;
  color: #D4AF37;
  font-size: 24px;
  border-radius: 8px;
}
</style>
