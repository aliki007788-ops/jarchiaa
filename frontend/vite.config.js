import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { VitePWA } from 'vite-plugin-pwa';

export default defineConfig({
plugins: [
vue(),
VitePWA({
registerType: 'autoUpdate',
manifest: {
name: '??????',
short_name: '??????',
theme_color: '#D4AF37',
background_color: '#0A0A0A',
display: 'standalone',
dir: 'rtl',
lang: 'fa'
}
})
],
server: {
port: 3000,
host: true
}
});


