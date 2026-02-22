// main.js
import './assets/js/main.js';
import './assets/css/main.css';

// اگه از Vue استفاده می‌کنی
import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);
app.mount('#app');

console.log('جارچیا با موفقیت راه‌اندازی شد!');
