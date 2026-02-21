const Utils = {
formatPrice(price) {
return new Intl.NumberFormat('fa-IR').format(price) + ' ?????';
},

toPersianNumber(str) {
const persianDigits = ['?', '?', '?', '?', '?', '?', '?', '?', '?', '?'];
return str.toString().replace(/\d/g, d => persianDigits[d]);
},

validatePhone(phone) {
return /^09[0-9]{9}$/.test(phone);
},

copyToClipboard(text) {
navigator.clipboard.writeText(text);
}
};

window.Utils = Utils;


