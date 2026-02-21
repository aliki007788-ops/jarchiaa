const Header = {
render() {
return <header class="header"> <div class="menu-btn" onclick="DrawerManager.toggle()"> <span></span> <span></span> <span></span> </div> <div class="logo" onclick="JarchiaApp.navigateTo('home')"> <i class="fas fa-bullhorn"></i> ?????? </div> <div class="header-icons"> <i class="fas fa-adjust" onclick="ThemeManager.toggle()"></i> <i class="fas fa-bell" onclick="NotificationService.showList()"></i> <i class="fas fa-search" onclick="SearchService.show()"></i> </div> </header>;
}
};

window.Header = Header;


