const AdminService = {
async getModules() {
return API.get('/admin/modules');
},

async toggleModule(moduleId, isActive) {
return API.put(/admin/modules/${moduleId}/toggle, { is_active: isActive });
},

async verifyBusiness(businessId, verified) {
return API.put(/admin/businesses/${businessId}/verify, { verified });
}
};

window.AdminService = AdminService;


