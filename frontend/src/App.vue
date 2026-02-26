<template>
  <div id="app">
    <div class="mobile-view">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
@import './assets/css/main.css';
@import './assets/css/theme.css';
</style>
