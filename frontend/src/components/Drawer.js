const Drawer = {
render() {
const user = JarchiaApp?.state?.user || {};
const role = JarchiaApp?.state?.currentRole || 'guest';

text
return `
  <div class="drawer-overlay" onclick="DrawerManager.close()"></div>
  <div class="drawer">
    <div class="drawer-profile">
      <div class="profile-header">
        <div class="profile-avatar"></div>
        <div>
          <div class="profile-name">${user.name || '?????'}</div>
          <div class="profile-role">${role}</div>
        </div>
      </div>
    </div>
    <div class="drawer-wallet" onclick="JarchiaApp.navigateTo('wallet')">
      <span>??? ???</span>
      <span class="wallet-balance">${user.wallet || 0}</span>
    </div>
    <div class="drawer-items">
      <div class="drawer-item" onclick="DrawerManager.navigate('home')">
        <i class="fas fa-home"></i> ????
      </div>
      <div class="drawer-item" onclick="DrawerManager.navigate('vitrin')">
        <i class="fas fa-store-alt"></i> ??????
      </div>
      ${role === 'business' ? `
        <div class="drawer-item" onclick="DrawerManager.navigate('business')">
          <i class="fas fa-store"></i> ??? ????????
        </div>
      ` : ''}
      ${role === 'admin' ? `
        <div class="drawer-item" onclick="DrawerManager.navigate('admin')">
          <i class="fas fa-crown"></i> ??? ??????
        </div>
      ` : ''}
      <div class="drawer-item" onclick="DrawerManager.navigate('profile')">
        <i class="fas fa-user"></i> ???????
      </div>
      <div class="drawer-item" onclick="DrawerManager.logout()">
        <i class="fas fa-sign-out-alt"></i> ????
      </div>
    </div>
  </div>
`;
}
};

window.Drawer = Drawer;


