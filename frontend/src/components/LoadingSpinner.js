const LoadingSpinner = {
show() {
const spinner = document.createElement('div');
spinner.className = 'loading-spinner';
spinner.innerHTML = '<div class="spinner"></div><div class="loading-text">?? ??? ????????...</div>';
document.body.appendChild(spinner);
},

hide() {
document.querySelector('.loading-spinner')?.remove();
}
};

window.LoadingSpinner = LoadingSpinner;


