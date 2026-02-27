 <template>
  <div class="home-page">
    <!-- استوری‌ها -->
    <div class="stories-section">
      <div class="section-title">داستان‌های امروز</div>
      <div class="stories-container">
        <div v-for="story in stories" :key="story.id" class="story-item">
          <div class="story-ring">
            <div class="story-avatar" :style="{background: story.color}">
              <i :class="story.icon"></i>
            </div>
          </div>
          <span class="story-name">{{ story.name }}</span>
        </div>
      </div>
    </div>

    <!-- آگهی‌ها -->
    <div class="section-title">آگهی‌های نزدیک شما</div>
    <AdCard v-for="ad in ads" :key="ad.id" :ad="ad" @click="showBusiness(ad.businessId)" />
  </div>
</template>

<script>
import AdCard from '../components/AdCard.vue'
import { ref } from 'vue'

export default {
  components: { AdCard },
  setup() {
    const stories = ref([
      { id: 1, name: 'شما', icon: 'fas fa-plus', color: '#D4AF37' },
      { id: 2, name: 'نانوایی', icon: 'fas fa-store', color: '#00b894' },
      { id: 3, name: 'رستوران', icon: 'fas fa-utensils', color: '#e84342' }
    ])

    const ads = ref([
      {
        id: 1,
        title: 'نان بربری تازه',
        businessId: 'B-101',
        businessName: 'نانوایی بربری‌زاده',
        price: '۴۵,۰۰۰',
        distance: 150,
        views: 234,
        likes: 45,
        comments: 12,
        badge: 'featured'
      },
      // ... بقیه آگهی‌ها
    ])

    const showBusiness = (id) => {
      router.push(`/business/${id}`)
    }

    return { stories, ads, showBusiness }
  }
}
</script>
