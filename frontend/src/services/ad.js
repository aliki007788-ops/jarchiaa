const AdService = {
async getNearbyAds(lat, lng, radius = 1000) {
return API.get(/ads/nearby?lat=${lat}&lng=${lng}&radius=${radius});
},

async getAdDetails(adId) {
return API.get(/ads/${adId});
},

async likeAd(adId) {
return API.post(/ads/${adId}/like);
}
};

window.AdService = AdService;


