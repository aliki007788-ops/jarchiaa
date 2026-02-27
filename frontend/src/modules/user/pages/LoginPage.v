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
    <div v-else-if="step === 'otp'" class="step">
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
        @click="verifyOTP"
        :disabled="!isOTPValid"
      >
        تأیید و ورود
      </button>
      
      <div class="resend" @click="resendOTP">
        ارسال مجدد کد
      </div>
    </div>
    
    <!-- مرحله ۳: ثبت‌نام (اگر کاربر جدید باشد) -->
    <div v-else-if="step === 'register'" class="step">
      <div class="input-box">
        <input 
          type="text" 
          v-model="fullName" 
          placeholder="نام و نام خانوادگی"
        >
      </div>
      
      <div class="input-box">
        <input 
          type="email" 
          v-model="email" 
          placeholder="ایمیل (اختیاری)"
        >
      </div>
      
      <button 
        class="main-btn" 
        @click="completeRegistration"
        :disabled="!fullName"
      >
        تکمیل ثبت‌نام
      </button>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useAuth } from '../composables/useAuth'
import { useRouter } from 'vue-router'

export default {
  setup() {
    const { login, verify, register } = useAuth()
    const router = useRouter()
    
    const step = ref('phone')
    const phone = ref('')
    const otp = ref(['', '', '', '', '', ''])
    const fullName = ref('')
    const email = ref('')
    
    const isPhoneValid = computed(() => {
      return phone.value.length === 10
    })
    
    const isOTPValid = computed(() => {
      return otp.value.every(digit => digit && digit.length === 1)
    })
    
    const sendOTP = async () => {
      try {
        await login(phone.value)
        step.value = 'otp'
      } catch (error) {
        alert('خطا در ارسال کد')
      }
    }
    
    const verifyOTP = async () => {
      const code = otp.value.join('')
      try {
        const response = await verify(phone.value, code)
        if (response.user.isVerified) {
          router.push('/')
        } else {
          step.value = 'register'
        }
      } catch (error) {
        alert('کد نامعتبر است')
      }
    }
    
    const completeRegistration = async () => {
      try {
        await register({
          phone: phone.value,
          fullName: fullName.value,
          email: email.value
        })
        router.push('/')
      } catch (error) {
        alert('خطا در ثبت‌نام')
      }
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
    
    const resendOTP = async () => {
      await sendOTP()
    }
    
    return {
      step,
      phone,
      otp,
      fullName,
      email,
      isPhoneValid,
      isOTPValid,
      sendOTP,
      verifyOTP,
      completeRegistration,
      moveToNext,
      moveToPrev,
      resendOTP
    }
  }
}
</script>

<style scoped>
.login-container {
  min-height: 80vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 20px;
}

.login-logo {
  font-size: 4rem;
  font-weight: 800;
  text-align: center;
  margin-bottom: 30px;
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
}

.input-box .prefix {
  position: absolute;
  left: 20px;
  top: 50%;
  transform: translateY(-50%);
  color: rgba(212, 175, 55, 0.6);
  border-right: 1px solid rgba(212, 175, 55, 0.2);
  padding-right: 10px;
}

.input-box input {
  width: 100%;
  padding: 18px 60px 18px 85px;
  border-radius: 40px;
  background: rgba(20, 20, 20, 0.9);
  border: 1px solid rgba(212, 175, 55, 0.3);
  color: #D4AF37;
  font-size: 1.2rem;
  outline: none;
  text-align: left;
}

.main-btn {
  width: 100%;
  padding: 18px;
  border-radius: 45px;
  background: linear-gradient(45deg, #D4AF37, #B8860B);
  border: none;
  color: #000;
  font-weight: 800;
  font-size: 1.3rem;
  cursor: pointer;
  box-shadow: 0 0 25px rgba(212, 175, 55, 0.5);
  margin-bottom: 15px;
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
  border: 1px solid #D4AF37;
  border-radius: 12px;
  color: #D4AF37;
  font-size: 24px;
  text-align: center;
}

.resend {
  text-align: center;
  color: #D4AF37;
  cursor: pointer;
  margin-top: 20px;
}
</style>