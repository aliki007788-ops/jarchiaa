const EitaaSSO = {
init() {
this.checkEitaaEnvironment();
this.autoLogin();
},

text
checkEitaaEnvironment() {
    this.isEitaa = navigator.userAgent.includes('Eitaa');
},

async autoLogin() {
    if (this.isEitaa && window.EitaaWebView) {
        try {
            const userData = await window.EitaaWebView.getUser();
            this.handleUserData(userData);
        } catch (error) {
            console.error('Auto login failed:', error);
        }
    }
},

handleUserData(userData) {
    window.parent.postMessage({
        type: 'EITAA_USER',
        data: userData
    }, '*');
},

async login() {
    const redirectUri = encodeURIComponent(window.location.href);
    const authUrl = `https://eitaa.com/oauth/authorize?client_id=${EITAA_CLIENT_ID}&redirect_uri=${redirectUri}&response_type=code`;
    window.location.href = authUrl;
}
};

window.EitaaSSO = EitaaSSO;


