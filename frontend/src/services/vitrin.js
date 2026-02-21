const VitrinService = {
async getItems(page = 1, category = 'all') {
return API.get(/vitrin?page=${page}&category=${category});
},

async getItemDetails(itemId) {
return API.get(/vitrin/${itemId});
}
};

window.VitrinService = VitrinService;


