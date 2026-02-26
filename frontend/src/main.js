import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/main.css'

const app = createApp(App)
app.use(router)
app.mount('#app')

console.log('✅ جارچیا با موفقیت راه‌اندازی شد!')
