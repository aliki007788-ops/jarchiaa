const AppointmentService = {
async getBusinesses() {
return API.get('/appointments/businesses');
},

async bookAppointment(data) {
return API.post('/appointments/book', data);
}
};

window.AppointmentService = AppointmentService;


