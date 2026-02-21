const NotificationService = {
async getNotifications() {
return API.get('/notifications');
},

async markAsRead(id) {
return API.put(/notifications/${id}/read);
},

showList() {
JarchiaApp.navigateTo('notifications');
}
};

window.NotificationService = NotificationService;


