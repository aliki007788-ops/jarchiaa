const AnalyticsService = {
async trackEvent(eventName, data = {}) {
return API.post('/analytics/track', { event: eventName, data });
}
};

window.AnalyticsService = AnalyticsService;


