const DrawerManager = {
toggle() {
document.getElementById('drawerMenu').classList.toggle('open');
document.getElementById('drawerOverlay').classList.toggle('open');
},

close() {
document.getElementById('drawerMenu').classList.remove('open');
document.getElementById('drawerOverlay').classList.remove('open');
},

navigate(page) {
this.close();
JarchiaApp.navigateTo(page);
},

logout() {
this.close();
JarchiaApp.logout();
}
};

window.DrawerManager = DrawerManager;


