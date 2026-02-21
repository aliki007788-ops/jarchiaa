const AuthService = {
token: localStorage.getItem('jarchia_token'),

async sendOTP(phone) {
try {
const response = await API.post('/auth/otp', { phone });
Toast.success('?? ????? ????? ??');
return response;
} catch (error) {
Toast.error('??? ?? ????? ??');
}
},

async verifyOTP(phone, code) {
try {
const response = await API.post('/auth/verify', { phone, code });
if (response.token) {
this.token = response.token;
localStorage.setItem('jarchia_token', response.token);
JarchiaApp.state.user = response.user;
JarchiaApp.state.isLoggedIn = true;
JarchiaApp.navigateTo('home');
}
return response;
} catch (error) {
Toast.error('?? ??????? ???');
}
},

logout() {
this.token = null;
localStorage.removeItem('jarchia_token');
JarchiaApp.logout();
}
};

window.AuthService = AuthService;


