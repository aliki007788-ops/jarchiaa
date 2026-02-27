<template>
  <div class="vitrin-page">
    <div class="vitrin-header">
      <div class="vitrin-title">
        <i class="fas fa-store-alt"></i>
        <span>ویترین</span>
      </div>
      <span class="badge">{{ vitrinItems.length }} آگهی ویژه</span>
    </div>

    <div class="vitrin-filters">
      <span v-for="filter in filters" :key="filter.id"
            class="vitrin-filter" :class="{active: filter.active}"
            @click="selectFilter(filter)">
        {{ filter.name }}
      </span>
    </div>

    <div class="vitrin-grid">
      <div v-for="item in vitrinItems" :key="item.id"
           class="vitrin-item" :class="item.size"
           @click="showBusiness(item.businessId)">
        <div class="vitrin-image" :style="{background: item.color}">
          <span class="vitrin-badge">{{ item.badge }}</span>
          <span class="vitrin-business-id">ID: {{ item.businessId }}</span>
          <div class="vitrin-overlay">
            <div class="title">{{ item.title }}</div>
            <div class="category">{{ item.category }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  setup() {
    const filters = ref([
      { id: 1, name: 'همه', active: true },
      { id: 2, name: 'غذا', active: false },
      { id: 3, name: 'فروشگاه', active: false },
      { id: 4, name: 'خدمات', active: false }
    ])

    const vitrinItems = ref([
      {
        id: 1,
        title: 'پیتزا شیل',
        businessId: 'B-107',
        category: 'رستوران',
        badge: 'ویژه',
        size: 'large',
        color: 'linear-gradient(135deg, #e84342, #ff7675)'
      },
      {
        id: 2,
        title: 'آرایشگاه مدرن',
        businessId: 'B-102',
        category: 'خدمات',
        badge: 'محبوب',
        size: 'vertical',
        color: 'linear-gradient(135deg, #6c5ce7, #a363d9)'
      }
    ])

    const selectFilter = (selected) => {
      filters.value.forEach(f => f.active = (f.id === selected.id))
    }

    const showBusiness = (id) => {
      router.push(`/business/${id}`)
    }

    return { filters, vitrinItems, selectFilter, showBusiness }
  }
}
</script>
