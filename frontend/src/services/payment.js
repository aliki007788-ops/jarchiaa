const PaymentService = {
async getBalance() {
return API.get('/wallet/balance');
},

async deposit(amount, gateway = 'zarinpal') {
return API.post('/wallet/deposit', { amount, gateway });
},

async withdraw(amount, cardNumber) {
return API.post('/wallet/withdraw', { amount, card_number: cardNumber });
}
};

window.PaymentService = PaymentService;


