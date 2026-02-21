const BusinessService = {
async register(data) {
return API.post('/business/register', data);
},

async getProfile() {
return API.get('/business/profile');
},

async getPlans() {
return API.get('/business/plans');
},

async purchasePlan(planId) {
return API.post('/business/purchase-plan', { plan_id: planId });
},

async getAds() {
return API.get('/business/ads');
},

async createAd(formData) {
return API.upload('/business/ads', formData);
}
};

window.BusinessService = BusinessService;


