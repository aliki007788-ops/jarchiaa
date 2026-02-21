const ModalBase = {
show(options) {
const modal = document.createElement('div');
modal.className = 'modal show';
modal.innerHTML = <div class="modal-content"> <div class="modal-header"> <h3>${options.title || ''}</h3> <button class="close-btn" onclick="ModalBase.close()">?</button> </div> <div class="modal-body">${options.content || ''}</div> <div class="modal-footer"> ${options.buttons?.map(btn =>
<button class="modal-btn ${btn.type}" onclick="${btn.onClick}">${btn.text}</button>
).join('')} </div> </div>;
document.body.appendChild(modal);
},

close() {
document.querySelector('.modal')?.remove();
}
};

window.ModalBase = ModalBase;


