import './assets/css/main.css';
import { initRouter } from './router.js';
import { loadUserData } from './services/user.js';

window.onload = () => {
    initRouter();
    loadUserData();
};
