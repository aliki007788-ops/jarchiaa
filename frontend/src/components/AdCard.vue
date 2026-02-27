<template>
  <div class="ad-card" @click="$emit('click')">
    <div class="ad-image" :style="{background: ad.color || '#D4AF37'}">
      <div v-if="ad.badge" class="ad-badges">
        <span class="badge" :class="ad.badge">{{ ad.badge }}</span>
      </div>
      <span class="ad-distance">
        <i class="fas fa-location-dot"></i> {{ ad.distance }} متر
      </span>
    </div>
    <div class="ad-content">
      <div class="ad-header">
        <span class="ad-title">{{ ad.title }}</span>
        <span class="ad-price">{{ ad.price }} تومان</span>
      </div>
      <div class="ad-business">
        <i class="fas fa-store"></i> {{ ad.businessName }}
      </div>
      <div class="ad-stats">
        <div class="stat-item"><i class="far fa-eye"></i> {{ ad.views }}</div>
        <div class="stat-item"><i class="far fa-heart"></i> {{ ad.likes }}</div>
        <div class="stat-item"><i class="far fa-comment"></i> {{ ad.comments }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    ad: { type: Object, required: true }
  }
}
</script>
