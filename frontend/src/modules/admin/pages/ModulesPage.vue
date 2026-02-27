<template>
  <div class="admin-modules">
    <div class="admin-header">
      <h2><i class="fas fa-crown"></i> پنل مدیریت ماژول‌ها</h2>
      <div class="admin-stats">
        <div class="stat-card">
          <span class="stat-number">{{ stats.totalUsers }}</span>
          <span>کاربران</span>
        </div>
        <div class="stat-card">
          <span class="stat-number">{{ stats.totalBusinesses }}</span>
          <span>کسب‌وکار</span>
        </div>
        <div class="stat-card">
          <span class="stat-number">{{ formatPrice(stats.totalRevenue) }}</span>
          <span>درآمد</span>
        </div>
      </div>
    </div>

    <!-- مدیریت ماژول‌ها -->
    <div class="modules-section">
      <h3>🔧 مدیریت ماژول‌ها</h3>
      <p class="info">با غیرفعال کردن هر ماژول، دسترسی به آن برای همه کاربران قطع می‌شود</p>

      <div class="modules-grid">
        <div v-for="module in modules" :key="module.id" class="module-card">
          <div class="module-info">
            <div class="module-icon">
              <i :class="'fas fa-' + module.icon"></i>
            </div>
            <div>
              <h4>{{ module.name_fa }}</h4>
              <p>{{ module.description }}</p>
              <span v-if="module.is_core" class="core-badge">هسته</span>
            </div>
          </div>
          
          <div class="module-toggle">
            <label class="switch">
              <input 
                type="checkbox" 
                :checked="module.is_active"
                :disabled="module.is_core"
                @change="toggleModule(module)"
              >
              <span class="slider"></span>
            </label>
            <span :class="{ 'active-text': module.is_active, 'inactive-text': !module.is_active }">
              {{ module.is_active ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- مدیریت پلن‌ها -->
    <div class="plans-section">
      <div class="section-header">
        <h3>📊 مدیریت پلن‌های آگهی</h3>
        <button class="btn-primary" @click="showNewPlanModal = true">
          <i class="fas fa-plus"></i> پلن جدید
        </button>
      </div>

      <div class="plans-table">
        <table>
          <thead>
            <tr>
              <th>نام پلن</th>
              <th>قیمت</th>
              <th>تعداد آگهی</th>
              <th>مدت (روز)</th>
              <th>وضعیت</th>
              <th>عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="plan in plans" :key="plan.id">
              <td>{{ plan.name_fa }}</td>
              <td>{{ formatPrice(plan.price) }}</td>
              <td>{{ plan.max_ads }}</td>
              <td>{{ plan.duration_days }}</td>
              <td>
                <span :class="plan.is_active ? 'badge-active' : 'badge-inactive'">
                  {{ plan.is_active ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td>
                <button class="btn-icon" @click="editPlan(plan)">
                  <i class="fas fa-edit"></i>
                </button>
                <button class="btn-icon" @click="togglePlanStatus(plan)">
                  <i :class="plan.is_active ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- درخواست‌های تأیید -->
    <div class="pending-section">
      <h3>⏳ درخواست‌های تأیید</h3>

      <div class="pending-tabs">
        <button 
          :class="['tab-btn', { active: pendingTab === 'businesses' }]"
          @click="pendingTab = 'businesses'"
        >
          کسب‌وکارها ({{ pendingBusinesses.length }})
        </button>
        <button 
          :class="['tab-btn', { active: pendingTab === 'ads' }]"
          @click="pendingTab = 'ads'"
        >
          آگهی‌ها ({{ pendingAds.length }})
        </button>
      </div>

      <div v-if="pendingTab === 'businesses'" class="pending-list">
        <div v-for="business in pendingBusinesses" :key="business.id" class="pending-item">
          <div class="pending-info">
            <h4>{{ business.name }}</h4>
            <p>{{ business.category }} • {{ business.phone }}</p>
          </div>
          <div class="pending-actions">
            <button class="btn-approve" @click="verifyBusiness(business.id, true)">
              <i class="fas fa-check"></i> تأیید
            </button>
            <button class="btn-reject" @click="verifyBusiness(business.id, false)">
              <i class="fas fa-times"></i> رد
            </button>
          </div>
        </div>
      </div>

      <div v-else class="pending-list">
        <div v-for="ad in pendingAds" :key="ad.id" class="pending-item">
          <div class="pending-info">
            <h4>{{ ad.title }}</h4>
            <p>{{ ad.business_name }} • {{ formatPrice(ad.price) }}</p>
          </div>
          <div class="pending-actions">
            <button class="btn-approve" @click="verifyAd(ad.id, true)">تأیید</button>
            <button class="btn-reject" @click="verifyAd(ad.id, false)">رد</button>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال ایجاد پلن جدید -->
    <div v-if="showNewPlanModal" class="modal" @click.self="showNewPlanModal = false">
      <div class="modal-content large">
        <h3>ایجاد پلن جدید</h3>
        
        <form @submit.prevent="createPlan">
          <div class="form-group">
            <label>نام پلن (فارسی)</label>
            <input v-model="newPlan.name_fa" required>
          </div>
          
          <div class="form-group">
            <label>نام پلن (انگلیسی)</label>
            <input v-model="newPlan.name" required>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label>قیمت (تومان)</label>
              <input type="number" v-model="newPlan.price" required>
            </div>
            <div class="form-group">
              <label>تعداد آگهی در ماه</label>
              <input type="number" v-model="newPlan.max_ads" required>
            </div>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label>مدت (روز)</label>
              <input type="number" v-model="newPlan.duration_days" required>
            </div>
            <div class="form-group">
              <label>درصد تخفیف</label>
              <input type="number" v-model="newPlan.discount_percent">
            </div>
          </div>
          
          <div class="form-group">
            <label>ماژول‌های فعال</label>
            <div class="modules-checkbox">
              <label v-for="module in modules" :key="module.key">
                <input 
                  type="checkbox" 
                  :value="module.key"
                  v-model="newPlan.modules"
                >
                {{ module.name_fa }}
              </label>
            </div>
          </div>
          
          <div class="form-checkbox">
            <label>
              <input type="checkbox" v-model="newPlan.is_popular">
              پلن محبوب
            </label>
          </div>
          
          <div class="modal-actions">
            <button type="submit" class="btn-primary">ایجاد پلن</button>
            <button type="button" class="btn-secondary" @click="showNewPlanModal = false">
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useAdmin } from '../composables/useAdmin'

export default {
  setup() {
    const { 
      getStats, 
      getModules, 
      getPlans, 
      toggleModule,
      createPlan,
      updatePlan,
      verifyBusiness,
      verifyAd,
      getPendingBusinesses,
      getPendingAds
    } = useAdmin()

    const stats = ref({})
    const modules = ref([])
    const plans = ref([])
    const pendingBusinesses = ref([])
    const pendingAds = ref([])
    const pendingTab = ref('businesses')
    const showNewPlanModal = ref(false)

    const newPlan = ref({
      name: '',
      name_fa: '',
      price: 0,
      max_ads: 0,
      duration_days: 30,
      modules: [],
      is_popular: false,
      discount_percent: 0
    })

    onMounted(async () => {
      stats.value = await getStats()
      modules.value = await getModules()
      plans.value = await getPlans()
      pendingBusinesses.value = await getPendingBusinesses()
      pendingAds.value = await getPendingAds()
    })

    const toggleModuleHandler = async (module) => {
      const newState = !module.is_active
      await toggleModule(module.key, newState)
      module.is_active = newState
    }

    const createPlanHandler = async () => {
      await createPlan(newPlan.value)
      showNewPlanModal.value = false
      plans.value = await getPlans()
    }

    const editPlan = (plan) => {
      // TODO: ویرایش پلن
    }

    const togglePlanStatus = async (plan) => {
      // TODO: فعال/غیرفعال کردن پلن
    }

    const verifyBusinessHandler = async (id, verified) => {
      await verifyBusiness(id, verified)
      pendingBusinesses.value = await getPendingBusinesses()
    }

    const verifyAdHandler = async (id, verified) => {
      await verifyAd(id, verified)
      pendingAds.value = await getPendingAds()
    }

    const formatPrice = (price) => {
      return price?.toLocaleString('fa-IR') + ' تومان'
    }

    return {
      stats,
      modules,
      plans,
      pendingBusinesses,
      pendingAds,
      pendingTab,
      showNewPlanModal,
      newPlan,
      toggleModule: toggleModuleHandler,
      createPlan: createPlanHandler,
      editPlan,
      togglePlanStatus,
      verifyBusiness: verifyBusinessHandler,
      verifyAd: verifyAdHandler,
      formatPrice
    }
  }
}
</script>

<style scoped>
.admin-modules {
  padding: 20px;
}

.admin-header {
  background: linear-gradient(135deg, #e84342, #c0392b);
  border-radius: 20px;
  padding: 25px;
  margin-bottom: 25px;
  color: white;
}

.admin-header h2 {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.admin-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 15px;
}

.stat-card {
  background: rgba(255,255,255,0.1);
  padding: 15px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
}

.stat-number {
  font-size: 24px;
  font-weight: 700;
}

.modules-section, .plans-section, .pending-section {
  background: #1A1A1A;
  border-radius: 20px;
  padding: 20px;
  margin-bottom: 25px;
  border: 1px solid #2A2A2A;
}

.modules-section h3, .plans-section h3, .pending-section h3 {
  color: #D4AF37;
  margin-bottom: 10px;
}

.info {
  color: #B8860B;
  font-size: 13px;
  margin-bottom: 20px;
}

.modules-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 15px;
}

.module-card {
  background: #2A2A2A;
  border-radius: 15px;
  padding: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border: 1px solid #3A3A3A;
}

.module-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.module-icon {
  width: 40px;
  height: 40px;
  background: #D4AF37;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #0A0A0A;
}

.core-badge {
  background: #e84342;
  color: white;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 10px;
  margin-top: 5px;
  display: inline-block;
}

.module-toggle {
  display: flex;
  align-items: center;
  gap: 10px;
}

.switch {
  position: relative;
  display: inline-block;
  width: 50px;
  height: 24px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
  border-radius: 24px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 20px;
  width: 20px;
  left: 2px;
  bottom: 2px;
  background-color: white;
  transition: .4s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: #D4AF37;
}

input:checked + .slider:before {
  transform: translateX(26px);
}

.active-text {
  color: #00b894;
}

.inactive-text {
  color: #e84342;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.plans-table {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th {
  text-align: right;
  padding: 12px;
  color: #D4AF37;
  border-bottom: 2px solid #D4AF37;
}

td {
  padding: 12px;
  border-bottom: 1px solid #2A2A2A;
}

.badge-active {
  background: #00b894;
  color: white;
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.badge-inactive {
  background: #e84342;
  color: white;
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.btn-icon {
  background: transparent;
  border: none;
  color: #D4AF37;
  cursor: pointer;
  padding: 5px;
  margin: 0 3px;
}

.pending-tabs {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.tab-btn {
  background: #2A2A2A;
  border: 1px solid #D4AF37;
  color: #B8860B;
  padding: 8px 16px;
  border-radius: 20px;
  cursor: pointer;
}

.tab-btn.active {
  background: #D4AF37;
  color: #0A0A0A;
  font-weight: 600;
}

.pending-item {
  background: #2A2A2A;
  border-radius: 12px;
  padding: 15px;
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.pending-info h4 {
  color: #D4AF37;
  margin-bottom: 5px;
}

.pending-actions {
  display: flex;
  gap: 10px;
}

.btn-approve {
  background: #00b894;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 20px;
  cursor: pointer;
}

.btn-reject {
  background: #e84342;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 20px;
  cursor: pointer;
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

.modal-content.large {
  width: 600px;
  max-width: 90%;
  max-height: 80vh;
  overflow-y: auto;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  color: #D4AF37;
  margin-bottom: 5px;
}

.form-group input, .form-group select {
  width: 100%;
  padding: 10px;
  background: #2A2A2A;
  border: 1px solid #D4AF37;
  border-radius: 8px;
  color: #F5E6D3;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.modules-checkbox {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
  padding: 10px;
  background: #2A2A2A;
  border-radius: 8px;
}

.modules-checkbox label {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #F5E6D3;
}

.form-checkbox {
  margin: 15px 0;
}

.modal-actions {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}
</style>