const BottomNav = {
render(activePage = 'home') {
return <nav class="bottom-nav"> <div class="nav-item ${activePage === 'home' ? 'active' : ''}" onclick="JarchiaApp.navigateTo('home')"> <i class="fas fa-home"></i> <span>????</span> </div> <div class="nav-item ${activePage === 'vitrin' ? 'active' : ''}" onclick="JarchiaApp.navigateTo('vitrin')"> <i class="fas fa-store-alt"></i> <span>??????</span> </div> <div class="nav-item nav-item-special" onclick="BottomNav.quickAction()"> <div class="special-button"> <i class="fas fa-plus"></i> </div> <span class="nav-label-special">????</span> </div> <div class="nav-item ${activePage === 'business' ? 'active' : ''}" onclick="JarchiaApp.navigateTo('business')"> <i class="fas fa-store"></i> <span>????????</span> </div> <div class="nav-item ${activePage === 'admin' ? 'active' : ''}" onclick="JarchiaApp.navigateTo('admin')"> <i class="fas fa-crown"></i> <span>??????</span> </div> </nav>;
},

quickAction() {
const role = JarchiaApp?.state?.currentRole || 'guest';
if (role === 'business') Toast.info('????? ???? ????');
else if (role === 'admin') JarchiaApp.navigateTo('admin');
else Toast.info('?????? ????');
}
};

window.BottomNav = BottomNav;


