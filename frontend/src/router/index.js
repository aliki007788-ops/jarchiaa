export function initRouter() {
    window.addEventListener('hashchange', handleRoute);
    handleRoute();
}

function handleRoute() {
    const hash = window.location.hash || '#home';
    const page = hash.replace('#', '');
    showPage(page);
}
