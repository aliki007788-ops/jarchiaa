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
          pattern="[0-9]*"
          inputmode="numeric"
          @keypress="isNumber"
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
      
      <div class="timer" v-if="timer > 0">{{ formatTimer(timer) }}</div>
      <div class="resend" @click="resendOTP" v-else>
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
import { ref, computed, onUnmounted } from 'vue'
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
    const timer = ref(0)
    let timerInterval = null
    
    const isPhoneValid = computed(() => {
      return phone.value.length === 10
    })
    
    const isOTPValid = computed(() => {
      return otp.value.every(digit => digit && digit.length === 1)
    })
    
    // فقط اعداد مجاز باشن
    const isNumber = (evt) => {
      const charCode = evt.which ? evt.which : evt.keyCode
      if (charCode < 48 || charCode > 57) {
        evt.preventDefault()
      }
    }
    
    const sendOTP = async () => {
      try {
        await login(phone.value)
        step.value = 'otp'
        startTimer(120) // 2 دقیقه تایمر
      } catch (error) {
        alert('خطا در ارسال کد')
      }
    }
    
    const verifyOTP = async () => {
      const code = otp.value.join('')
      try {
        const response = await verify(phone.value, code)
        
        // ✅ ذخیره اطلاعات کاربر در localStorage
        if (response && response.token) {
          localStorage.setItem('token', response.token)
          
          if (response.user) {
            localStorage.setItem('userRole', response.user.role || 'user')
            localStorage.setItem('userName', response.user.full_name || response.user.name || 'کاربر')
            localStorage.setItem('userPhone', response.user.phone || phone.value)
            localStorage.setItem('userId', response.user.id || '')
          }
          
          // اگه کاربر تأیید شده بود بره خانه، وگرنه بره ثبت‌نام
          if (response.user?.isVerified) {
            router.push('/')
          } else {
            step.value = 'register'
          }
        } else {
          alert('خطا در ورود به سیستم')
        }
      } catch (error) {
        console.error('Verify error:', error)
        alert('کد نامعتبر است')
      }
    }
    
    const completeRegistration = async () => {
      try {
        const response = await register({
          phone: phone.value,
          fullName: fullName.value,
          email: email.value
        })
        
        // ✅ بعد از ثبت‌نام کامل، دوباره اطلاعات رو ذخیره کن
        if (response && response.token) {
          localStorage.setItem('token', response.token)
          localStorage.setItem('userRole', response.user?.role || 'user')
          localStorage.setItem('userName', fullName.value)
          localStorage.setItem('userPhone', phone.value)
        }
        
        router.push('/')
      } catch (error) {
        alert('خطا در ثبت‌نام')
      }
    }
    
    const startTimer = (seconds) => {
      timer.value = seconds
      if (timerInterval) clearInterval(timerInterval)
      
      timerInterval = setInterval(() => {
        if (timer.value > 0) {
          timer.value--
        } else {
          clearInterval(timerInterval)
        }
      }, 1000)
    }
    
    const formatTimer = (seconds) => {
      const mins = Math.floor(seconds / 60)
      const secs = seconds % 60
      return `${mins}:${secs < 10 ? '0' : ''}${secs}`
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
    
    // پاکسازی تایمر موقع خروج از کامپوننت
    onUnmounted(() => {
      if (timerInterval) clearInterval(timerInterval)
    })
    
    return {
      step,
      phone,
      otp,
      fullName,
      email,
      timer,
      isPhoneValid,
      isOTPValid,
      isNumber,
      sendOTP,
      verifyOTP,
      completeRegistration,
      moveToNext,
      moveToPrev,
      resendOTP,
      formatTimer
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
  z-index: 2;
}

.input-box .prefix {
  position: absolute;
  left: 20px;
  top: 50%;
  transform: translateY(-50%);
  color: rgba(212, 175, 55, 0.6);
  border-right: 1px solid rgba(212, 175, 55, 0.2);
  padding-right: 10px;
  z-index: 2;
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

.input-box input:focus {
  border-color: #D4AF37;
  box-shadow: 0 0 15px rgba(212, 175, 55, 0.3);
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
  transition: transform 0.3s;
}

.main-btn:hover {
  transform: translateY(-2px);
}

.main-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
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
  outline: none;
}

.otp-input:focus {
  border-width: 2px;
  box-shadow: 0 0 15px rgba(212, 175, 55, 0.3);
}

.timer {
  text-align: center;
  color: #B8860B;
  font-size: 18px;
  margin: 15px 0;
}

.resend {
  text-align: center;
  color: #D4AF37;
  cursor: pointer;
  margin-top: 20px;
  text-decoration: underline;
}

.resend:hover {
  color: #B8860B;
}
</style>
