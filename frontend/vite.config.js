import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  // مشخص کردن فایل ورودی
  build: {
    rollupOptions: {
      input: '/src/main.js', // فایل اصلی رو مشخص کن
    }
  }
});
