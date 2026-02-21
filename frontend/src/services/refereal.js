const ReferralService = {
async getCode() {
return API.get('/referral/code');
},

async getStats() {
return API.get('/referral/stats');
}
};

window.ReferralService = ReferralService;


