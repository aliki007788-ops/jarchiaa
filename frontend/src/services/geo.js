const GeoService = {
async getCurrentLocation() {
return new Promise((resolve, reject) => {
navigator.geolocation.getCurrentPosition(
pos => resolve({ lat: pos.coords.latitude, lng: pos.coords.longitude }),
err => reject(err)
);
});
}
};

window.GeoService = GeoService;


