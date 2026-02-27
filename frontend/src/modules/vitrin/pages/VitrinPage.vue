<template>
  <div class="vitrin-page">
    <div class="vitrin-header">
      <h2>✨ ویترین جارچیا</h2>
      <p>آگهی‌های ویژه و پرطرفدار</p>
    </div>

    <!-- گرید ویترین -->
    <div class="vitrin-grid">
      <div 
        v-for="ad in vitrinAds" 
        :key="ad.id"
        class="vitrin-item"
        :class="ad.size"
        @click="goToAdDetail(ad.id)"
      >
        <div class="vitrin-image" :style="{background: ad.color}">
          <span class="vitrin-badge">{{ ad.badge }}</span>
          <div class="vitrin-overlay">
            <h3>{{ ad.title }}</h3>
            <p>{{ ad.business_name }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useVitrin } from '../composables/useVitrin'

export default {
  setup() {
    const router = useRouter()
    const { getVitrinAds } = useVitrin()
    const vitrinAds = ref([])

    onMounted(async () => {
      vitrinAds.value = await getVitrinAds()
    })

    const goToAdDetail = (adId) => {
      router.push(`/ad/${adId}`)
    }

    return { vitrinAds, goToAdDetail }
  }
}
</script>