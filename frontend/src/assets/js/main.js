const JarchiaApp = {
state: {
user: null,
isLoggedIn: false,
currentRole: 'guest',
currentPage: 'login',
theme: 'dark'
},

init() {
this.loadState();
this.initEventListeners();
this.checkAuth();
this.initTheme();
},

navigateTo(page) {
this.state.currentPage = page;
this.saveState();
this.renderPage(page);
},

renderPage(page) {
document.querySelectorAll('.page').forEach(p => p.classList.remove('active'));
const targetPage = document.getElementById(${page}Page);
if (targetPage) targetPage.classList.add('active');
},

initTheme() {
const savedTheme = localStorage.getItem('jarchia_theme') || 'dark';
this.setTheme(savedTheme);
},

setTheme(theme) {
this.state.theme = theme;
document.body.className = theme === 'dark' ? 'dark-mode' : 'light-mode';
localStorage.setItem('jarchia_theme', theme);
},

logout() {
localStorage.removeItem('jarchia_token');
localStorage.removeItem('jarchia_state');
this.state.user = null;
this.state.isLoggedIn = false;
this.state.currentRole = 'guest';
this.navigateTo('login');
}
};

document.addEventListener('DOMContentLoaded', () => {
window.JarchiaApp = JarchiaApp;
JarchiaApp.init();
});

function toggleTheme() {
JarchiaApp.setTheme(JarchiaApp.state.theme === 'dark' ? 'light' : 'dark');
}

function toggleDrawer() {
document.getElementById('drawerMenu').classList.toggle('open');
document.getElementById('drawerOverlay').classList.toggle('open');
}


