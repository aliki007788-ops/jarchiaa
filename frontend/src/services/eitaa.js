const EitaaService = {
async login() {
window.location.href = 'https://eitaa.com/oauth/authorize';
},

shareAd(adId) {
window.open(https://eitaa.com/share?text=${adId});
}
};

window.EitaaService = EitaaService;


