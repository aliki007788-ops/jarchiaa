const ThemeManager = {
init() {
this.loadTheme();
},

loadTheme() {
const savedTheme = localStorage.getItem('jarchia_theme') || 'dark';
this.applyTheme(savedTheme);
},

applyTheme(theme) {
document.body.className = theme === 'dark' ? 'dark-mode' : 'light-mode';
localStorage.setItem('jarchia_theme', theme);
},

toggle() {
const current = document.body.className.includes('dark') ? 'dark' : 'light';
this.applyTheme(current === 'dark' ? 'light' : 'dark');
}
};

document.addEventListener('DOMContentLoaded', () => {
window.ThemeManager = ThemeManager;
ThemeManager.init();
});


