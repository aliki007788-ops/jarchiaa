const StatsBadge = {
render(views = 0, likes = 0, comments = 0) {
return <div class="stats-badge"> <span><i class="far fa-eye"></i> ${Utils.toPersianNumber(views)}</span> <span><i class="far fa-heart"></i> ${Utils.toPersianNumber(likes)}</span> <span><i class="far fa-comment"></i> ${Utils.toPersianNumber(comments)}</span> </div>;
}
};

window.StatsBadge = StatsBadge;


