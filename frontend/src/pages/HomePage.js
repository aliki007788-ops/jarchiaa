export default {
  name: 'HomePage',
  template: `
    <div class="home-page">
      <div class="stories-section">
        <h3>داستان‌های امروز</h3>
        <div class="stories-container">
          <div class="story-item">
            <div class="story-ring">
              <i class="fas fa-plus"></i>
            </div>
            <span>شما</span>
          </div>
        </div>
      </div>
      <div class="ads-section">
        <h3>آگهی‌های نزدیک شما</h3>
        <div class="ads-container"></div>
      </div>
    </div>
  `
}
