const StoriesService = {
async getStories() {
return API.get('/stories');
},

async createStory(formData) {
return API.upload('/stories', formData);
}
};

window.StoriesService = StoriesService;


