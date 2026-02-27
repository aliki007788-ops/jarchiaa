<template>
  <div class="login-container">
    <div class="login-logo">جارچیا</div>
    
    <!-- مرحله ۱: وارد کردن شماره تلفن -->
    <div v-if="step === 'phone'" class="step">
      <div class="input-box">
        <span class="flag-icon">🇮🇷</span>
        <span class="prefix">+98</span>
        <input 
          type="tel" 
          v-model="phone" 
          placeholder="9120000000" 
          maxlength="10"
          @keyup.enter="sendOTP"
        >
      </div>
      
      <button 
        class="main-btn" 
        @click="sendOTP"
        :disabled="!isPhoneValid"
      >
        ارسال کد تأیید
      </button>
    </div>
    
    <!-- مرحله ۲: وارد کردن کد OTP -->
    <div v-else class="step">
      <div class="otp-container">
        <input 
          v-for="(digit, index) in 6" 
          :key="index"
          type="text"
          maxlength="1"
          class="otp-input"
          v-model="otp[index]"
          @input="moveToNext($event, index)"
          @keydown.backspace="moveToPrev($event, index)"
        >
      </div>
      
      <button 
        class="main-btn" 
        @click="login"
        :disabled="!isOTPValid"
      >
        ورود به جارچیا
      </button>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'

export default {
  setup() {
    const step = ref('phone')
    const phone = ref('')
    const otp = ref(['', '', '', '', '', ''])
    
    const isPhoneValid = computed(() => {
      return phone.value.length === 10
    })
    
    const isOTPValid = computed(() => {
      return otp.value.every(digit => digit && digit.length === 1)
    })
    
    const sendOTP = () => {
      if (isPhoneValid.value) {
        step.value = 'otp'
      }
    }
    
    const login = () => {
      // ذخیره اطلاعات کاربر در localStorage
      localStorage.setItem('token', 'jarchia-token-123456')
      localStorage.setItem('userRole', 'user')
      localStorage.setItem('userName', 'کاربر جارچیا')
      localStorage.setItem('userPhone', phone.value)
      
      // رفرش صفحه و رفتن به خانه
      window.location.href = '/'
    }
    
    const moveToNext = (event, index) => {
      if (event.target.value && index < 5) {
        const next = event.target.nextElementSibling
        if (next) next.focus()
      }
    }
    
    const moveToPrev = (event, index) => {
      if (!event.target.value && index > 0) {
        const prev = event.target.previousElementSibling
        if (prev) prev.focus()
      }
    }
    
    return {
      step,
      phone,
      otp,
      isPhoneValid,
      isOTPValid,
      sendOTP,
      login,
      moveToNext,
      moveToPrev
    }
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 20px;
  background: #0A0A0A;
}

.login-logo {
  font-size: 4rem;
  font-weight: 800;
  text-align: center;
  margin-bottom: 50px;
  background: linear-gradient(to bottom, #D4AF37, #FFF, #D4AF37);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}

.input-box {
  position: relative;
  margin-bottom: 25px;
}

.input-box .flag-icon {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 1.5rem;
  z-index: 2;
}

.input-box .prefix {
  position: absolute;
  left: 20px;
  top: 50%;
  transform: translateY(-50%);
  color: #D4AF37;
  border-right: 1px solid #D4AF37;
  padding-right: 10px;
  z-index: 2;
}

.input-box input {
  width: 100%;
  padding: 18px 60px 18px 85px;
  border-radius: 40px;
  background: #1A1A1A;
  border: 2px solid #D4AF37;
  color: #D4AF37;
  font-size: 1.2rem;
  outline: none;
  text-align: left;
}

.main-btn {
  width: 100%;
  padding: 18px;
  border-radius: 45px;
  background: #D4AF37;
  border: none;
  color: #000;
  font-weight: 800;
  font-size: 1.3rem;
  cursor: pointer;
  box-shadow: 0 0 25px rgba(212, 175, 55, 0.5);
  transition: all 0.3s;
}

.main-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 0 35px rgba(212, 175, 55, 0.8);
}

.main-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
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
  background: #1A1A1A;
  border: 2px solid #D4AF37;
  border-radius: 12px;
  color: #D4AF37;
  font-size: 24px;
  text-align: center;
  outline: none;
}

.otp-input:focus {
  border-width: 3px;
  box-shadow: 0 0 15px rgba(212, 175, 55, 0.5);
}
</style>
