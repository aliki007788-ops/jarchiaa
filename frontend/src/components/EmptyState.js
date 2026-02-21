const EmptyState = {
render(title = '????? ???? ???', message = '') {
return <div class="empty-state"> <i class="fas fa-box-open empty-icon"></i> <h3>${title}</h3> <p>${message}</p> </div>;
}
};

window.EmptyState = EmptyState;


