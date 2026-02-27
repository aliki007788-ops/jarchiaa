package service

import (
    "errors"
    "time"
    "business-service/models"
    "business-service/repository"
)

type PlanService struct {
    businessRepo *repository.BusinessRepository
    planRepo     *repository.PlanRepository
}

func NewPlanService(businessRepo *repository.BusinessRepository, planRepo *repository.PlanRepository) *PlanService {
    return &PlanService{
        businessRepo: businessRepo,
        planRepo:     planRepo,
    }
}

// دریافت پلن‌های فعال
func (s *PlanService) GetActivePlans() ([]models.Plan, error) {
    return s.planRepo.FindActive()
}

// ثبت‌نام کسب‌وکار جدید
func (s *PlanService) RegisterBusiness(userID string, req *models.RegisterBusinessRequest) (*models.Business, error) {
    // بررسی وجود کسب‌وکار با این کاربر
    existing, _ := s.businessRepo.FindByUserID(userID)
    if existing != nil {
        return nil, errors.New("شما قبلاً یک کسب‌وکار ثبت کرده‌اید")
    }
    
    business := &models.Business{
        UserID:        uuid.MustParse(userID),
        Name:          req.Name,
        Category:      req.Category,
        Description:   req.Description,
        Address:       req.Address,
        Phone:         req.Phone,
        Location:      fmt.Sprintf("POINT(%f %f)", req.Longitude, req.Latitude),
        PlanID:        1, // پلن رایگان
        IsVerified:    false,
        IsActive:      true,
        CommissionRate: 10.0,
    }
    
    if err := s.businessRepo.Create(business); err != nil {
        return nil, err
    }
    
    return business, nil
}

// خرید پلن
func (s *PlanService) PurchasePlan(businessID string, req *models.PurchasePlanRequest) error {
    plan, err := s.planRepo.FindByID(req.PlanID)
    if err != nil {
        return errors.New("پلن یافت نشد")
    }
    
    business, err := s.businessRepo.FindByID(businessID)
    if err != nil {
        return err
    }
    
    // محاسبه قیمت با تخفیف (اگر کاربر قبلاً پلن داشته)
    price := plan.Price
    if business.PlanID > 1 { // اگر پلن فعلی رایگان نیست
        price = int64(float64(plan.Price) * 0.9) // 10% تخفیف
    }
    
    // TODO: بررسی موجودی کیف پول
    
    business.PlanID = plan.ID
    business.PlanExpiresAt = time.Now().AddDate(0, 0, plan.DurationDays)
    
    return s.businessRepo.Update(business)
}

// دریافت آمار کسب‌وکار
func (s *PlanService) GetBusinessStats(businessID string) (*models.BusinessStats, error) {
    business, err := s.businessRepo.FindByID(businessID)
    if err != nil {
        return nil, err
    }
    
    // دریافت آمار از سرویس‌های دیگر
    stats := &models.BusinessStats{
        TotalAds:    0,
        ActiveAds:   0,
        TotalViews:  0,
        TotalClicks: 0,
        TotalLikes:  0,
        ViewsToday:  0,
        ClicksToday: 0,
    }
    
    // TODO: دریافت از ad-service
    
    return stats, nil
}
