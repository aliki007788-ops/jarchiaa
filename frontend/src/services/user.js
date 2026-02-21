const UserService = {
async getProfile() {
return API.get('/user/profile');
},

async updateProfile(data) {
return API.put('/user/profile', data);
},

async getFavorites() {
return API.get('/user/favorites');
},

async addFavorite(itemId, type) {
return API.post('/user/favorites', { item_id: itemId, type });
},

async removeFavorite(itemId) {
return API.delete(/user/favorites/${itemId});
}
};

window.UserService = UserService;


