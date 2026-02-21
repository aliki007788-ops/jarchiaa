const RoleSwitcher = {
render() {
return <div class="role-switcher"> <div class="role-option user active" onclick="RoleSwitcher.switch('user')">?????</div> <div class="role-option business" onclick="RoleSwitcher.switch('business')">????????</div> <div class="role-option admin" onclick="RoleSwitcher.switch('admin')">????</div> </div>;
},

switch(role) {
JarchiaApp.state.currentRole = role;
JarchiaApp.updateUI();
}
};

window.RoleSwitcher = RoleSwitcher;


