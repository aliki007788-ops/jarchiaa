const WalletService = {
async getBalance() {
return PaymentService.getBalance();
},

async deposit(amount) {
return PaymentService.deposit(amount);
},

async withdraw(amount, cardNumber) {
return PaymentService.withdraw(amount, cardNumber);
}
};

window.WalletService = WalletService;


