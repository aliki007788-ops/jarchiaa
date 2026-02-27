<template>
  <div class="plans-page">
    <div class="page-header">
      <h2>📊 پلن‌های آگهی</h2>
      <p>پلن مناسب کسب‌وکار خود را انتخاب کنید</p>
    </div>

    <div class="plans-grid">
      <!-- پلن نقره‌ای -->
      <div 
        v-for="plan in plans" 
        :key="plan.id"
        class="plan-card" 
        :class="{ popular: plan.id === 2 }"
        @click="selectPlan(plan)"
      >
        <div v-if="plan.id === 2" class="popular-badge">پرفروش‌ترین</div>
        
        <div class="plan-name">{{ plan.name_fa }}</div>
        <div class="plan-price">
          {{ formatPrice(plan.price) }}
          <span>/ماه</span>
        </div>

        <ul class="plan-features">
          <li>
            <i class="fas fa-check"></i>
            {{ plan.max_ads }} آگهی در ماه
          </li>
          <li>
            <i :class="plan.has_geo_targeting ? 'fas fa-check' : 'fas fa-times'"></i>
            هدف‌گیری جغرافیایی
          </li>
          <li>
            <i :class="plan.has_analytics ? 'fas fa-check' : 'fas fa-times'"></i>
            آمار پیشرفته
          </li>
          <li>
            <i :class="plan.has_priority ? 'fas fa-check' : 'fas fa-times'"></i>
            اولویت نمایش
          </li>
        </ul>

        <button class="plan-btn" :class="{ outline: plan.id !== 2 }">
          انتخاب پلن
        </button>
      </div>
    </div>

    <!-- مودال تأیید خرید -->
    <div v-if="showModal" class="modal" @click.self="showModal = false">
      <div class="modal-content">
        <h3>خرید پلن {{ selectedPlan?.name_fa }}</h3>
        <p>موجودی کیف پول: {{ formatPrice(walletBalance) }}</p>
        <p>قیمت پلن: {{ formatPrice(selectedPlan?.price) }}</p>
        <p v-if="walletBalance < selectedPlan?.price" class="error">
          موجودی کافی نیست
        </p>
        
        <div class="modal-actions">
          <button 
            class="btn-primary" 
            @click="confirmPurchase"
            :disabled="walletBalance < selectedPlan?.price"
          >
            تأیید و پرداخت
          </button>
          <button class="btn-secondary" @click="showModal = false">
            انصراف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useBusiness } from '../composables/useBusiness'

export default {
  setup() {
    const { getPlans, purchasePlan, getWalletBalance } = useBusiness()
    const plans = ref([])
    const selectedPlan = ref(null)
    const showModal = ref(false)
    const walletBalance = ref(0)

    onMounted(async () => {
      plans.value = await getPlans()
      walletBalance.value = await getWalletBalance()
    })

    const selectPlan = (plan) => {
      selectedPlan.value = plan
      showModal.value = true
    }

    const confirmPurchase = async () => {
      try {
        await purchasePlan(selectedPlan.value.id)
        showModal.value = false
        alert('پلن با موفقیت خریداری شد')
      } catch (error) {
        alert('خطا در خرید پلن')
      }
    }

    const formatPrice = (price) => {
      return price?.toLocaleString('fa-IR') + ' تومان'
    }

    return {
      plans,
      selectedPlan,
      showModal,
      walletBalance,
      selectPlan,
      confirmPurchase,
      formatPrice
    }
  }
}
</script>

<style scoped>
.plans-page {
  padding: 20px;
}

.page-header {
  text-align: center;
  margin-bottom: 30px;
}

.page-header h2 {
  color: #D4AF37;
  font-size: 24px;
  margin-bottom: 10px;
}

.plans-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
  margin-top: 30px;
}

.plan-card {
  background: #1A1A1A;
  border-radius: 20px;
  padding: 30px 20px;
  border: 1px solid #2A2A2A;
  position: relative;
  cursor: pointer;
  transition: all 0.3s;
}

.plan-card:hover {
  transform: translateY(-4px);
  border-color: #D4AF37;
  box-shadow: 0 10px 30px rgba(212, 175, 55, 0.2);
}

.plan-card.popular {
  border: 2px solid #D4AF37;
  background: #1A1A1A;
}

.popular-badge {
  position: absolute;
  top: -12px;
  left: 20px;
  background: #D4AF37;
  color: #0A0A0A;
  padding: 4px 16px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.plan-name {
  font-size: 22px;
  font-weight: 700;
  color: #D4AF37;
  margin-bottom: 15px;
}

.plan-price {
  font-size: 32px;
  font-weight: 800;
  margin-bottom: 20px;
}

.plan-price span {
  font-size: 14px;
  font-weight: 400;
  color: #B8860B;
}

.plan-features {
  list-style: none;
  margin: 20px 0;
}

.plan-features li {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 12px;
  color: #F5E6D3;
  font-size: 14px;
}

.plan-features i.fa-check {
  color: #00b894;
}

.plan-features i.fa-times {
  color: #e84342;
}

.plan-btn {
  width: 100%;
  padding: 12px;
  border-radius: 30px;
  border: none;
  background: #D4AF37;
  color: #0A0A0A;
  font-weight: 600;
  cursor: pointer;
  transition: 0.3s;
}

.plan-btn.outline {
  background: transparent;
  border: 1px solid #D4AF37;
  color: #D4AF37;
}

.plan-btn:hover {
  background: #B8860B;
}

.plan-btn.outline:hover {
  background: #D4AF37;
  color: #0A0A0A;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: #1A1A1A;
  padding: 30px;
  border-radius: 20px;
  max-width: 400px;
  width: 90%;
  border: 2px solid #D4AF37;
}

.modal-actions {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.btn-primary {
  flex: 2;
  background: #D4AF37;
  color: #0A0A0A;
  border: none;
  padding: 12px;
  border-radius: 30px;
  font-weight: 600;
  cursor: pointer;
}

.btn-secondary {
  flex: 1;
  background: transparent;
  border: 1px solid #D4AF37;
  color: #D4AF37;
  padding: 12px;
  border-radius: 30px;
  cursor: pointer;
}

.error {
  color: #e84342;
  margin-top: 10px;
}
</style>