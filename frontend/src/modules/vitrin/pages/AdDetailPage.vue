<template>
  <div class="ad-detail-page" v-if="ad">
    <!-- تصویر آگهی -->
    <div class="ad-detail-image" :style="{background: ad.color}">
      <button class="back-btn" @click="goBack">
        <i class="fas fa-arrow-right"></i>
      </button>
      <span class="ad-badge">{{ ad.badge }}</span>
    </div>

    <!-- اطلاعات آگهی -->
    <div class="ad-detail-info">
      <h1>{{ ad.title }}</h1>
      <p class="ad-description">{{ ad.description }}</p>
      
      <div class="ad-stats-large">
        <div class="stat" @click="likeAd">
          <i :class="isLiked ? 'fas fa-heart' : 'far fa-heart'"></i>
          <span>{{ ad.likes }}</span>
        </div>
        <div class="stat">
          <i class="far fa-eye"></i>
          <span>{{ ad.views }}</span>
        </div>
        <div class="stat">
          <i class="far fa-comment"></i>
          <span>{{ ad.comments.length }}</span>
        </div>
      </div>

      <!-- اطلاعات کسب‌وکار (با کلیک می‌ره به صفحه کسب‌وکار) -->
      <div class="business-card" @click="goToBusinessPage(ad.business_id)">
        <div class="business-avatar"></div>
        <div>
          <h3>{{ ad.business_name }}</h3>
          <p>مشاهده همه آگهی‌های این کسب‌وکار</p>
        </div>
        <i class="fas fa-chevron-left"></i>
      </div>

      <!-- بخش نظرات -->
      <CommentsSection 
        :adId="ad.id" 
        :comments="ad.comments"
        @newComment="addComment"
      />
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useVitrin } from '../composables/useVitrin'
import CommentsSection from '../components/CommentsSection.vue'

export default {
  components: { CommentsSection },
  setup() {
    const route = useRoute()
    const router = useRouter()
    const { getAdById, likeAd: apiLikeAd, addComment: apiAddComment } = useVitrin()
    
    const ad = ref(null)
    const isLiked = ref(false)

    onMounted(async () => {
      const adId = route.params.id
      ad.value = await getAdById(adId)
    })

    const likeAd = async () => {
      if (!isLiked.value) {
        await apiLikeAd(ad.value.id)
        ad.value.likes++
        isLiked.value = true
      }
    }

    const addComment = async (text) => {
      const newComment = await apiAddComment(ad.value.id, text)
      ad.value.comments.push(newComment)
    }

    const goToBusinessPage = (businessId) => {
      router.push(`/business/${businessId}`)
    }

    const goBack = () => {
      router.back()
    }

    return { ad, isLiked, likeAd, addComment, goToBusinessPage, goBack }
  }
}
</script>