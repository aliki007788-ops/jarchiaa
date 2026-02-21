const LoginPage = {
render() {
return <div class="login-container"> <div class="login-logo">??????</div> <div class="input-box"> <span class="flag-icon">????</span> <span class="prefix">+98</span> <input type="tel" id="phone" placeholder="9120000000"> </div> <button class="main-btn" onclick="LoginPage.submit()">????</button> </div>;
},

submit() {
const phone = document.getElementById('phone').value;
if (!Utils.validatePhone(phone)) {
Toast.error('????? ?????? ????? ????');
return;
}
AuthService.sendOTP(phone);
}
};

window.LoginPage = LoginPage;


