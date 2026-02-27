package service

import (
    "admin-service/models"
    "admin-service/repository"
    "encoding/json"
    "fmt"
)

type ModuleService struct {
    moduleRepo     *repository.ModuleRepository
    configRepo     *repository.ConfigRepository
    planRepo       *repository.PlanRepository
    eventPublisher *EventPublisher
}

func NewModuleService(
    moduleRepo *repository.ModuleRepository,
    configRepo *repository.ConfigRepository,
    planRepo *repository.PlanRepository,
    eventPublisher *EventPublisher,
) *ModuleService {
    return &ModuleService{
        moduleRepo:     moduleRepo,
        configRepo:     configRepo,
        planRepo:       planRepo,
        eventPublisher: eventPublisher,
    }
}

// دریافت همه ماژول‌ها
func (s *ModuleService) GetAllModules() ([]models.Module, error) {
    return s.moduleRepo.FindAll()
}

// تغییر وضعیت ماژول
func (s *ModuleService) ToggleModule(moduleKey string, isActive bool) error {
    module, err := s.moduleRepo.FindByKey(moduleKey)
    if err != nil {
        return fmt.Errorf("ماژول یافت نشد: %s", moduleKey)
    }
    
    if module.IsCore {
        return fmt.Errorf("ماژول هسته را نمی‌توان غیرفعال کرد")
    }
    
    module.IsActive = isActive
    if err := s.moduleRepo.Update(module); err != nil {
        return err
    }
    
    // انتشار رویداد برای همه سرویس‌ها
    event := map[string]interface{}{
        "type":    "module_toggle",
        "module":  moduleKey,
        "active":  isActive,
        "user_id": "admin",
    }
    s.eventPublisher.Publish("module.events", event)
    
    return nil
}

// تعریف پلن جدید
func (s *ModuleService) CreatePlan(plan *models.PlanConfig) error {
    // اعتبارسنجی ماژول‌ها
    for _, moduleKey := range plan.Modules {
        module, err := s.moduleRepo.FindByKey(moduleKey)
        if err != nil || module == nil {
            return fmt.Errorf("ماژول %s معتبر نیست", moduleKey)
        }
    }
    
    return s.planRepo.Create(plan)
}

// ویرایش پلن
func (s *ModuleService) UpdatePlan(planID int, plan *models.PlanConfig) error {
    existing, err := s.planRepo.FindByID(planID)
    if err != nil {
        return err
    }
    
    // اعتبارسنجی ماژول‌ها
    for _, moduleKey := range plan.Modules {
        module, err := s.moduleRepo.FindByKey(moduleKey)
        if err != nil || module == nil {
            return fmt.Errorf("ماژول %s معتبر نیست", moduleKey)
        }
    }
    
    existing.Name = plan.Name
    existing.NameFa = plan.NameFa
    existing.Price = plan.Price
    existing.MaxAds = plan.MaxAds
    existing.DurationDays = plan.DurationDays
    existing.Modules = plan.Modules
    existing.IsActive = plan.IsActive
    existing.IsPopular = plan.IsPopular
    existing.DiscountPercent = plan.DiscountPercent
    
    return s.planRepo.Update(existing)
}

// تأیید کسب‌وکار
func (s *ModuleService) VerifyBusiness(businessID string, verify bool) error {
    // TODO: فراخوانی business-service
    event := map[string]interface{}{
        "type":        "business_verification",
        "business_id": businessID,
        "verified":    verify,
        "timestamp":   time.Now(),
    }
    s.eventPublisher.Publish("business.events", event)
    return nil
}

// تأیید آگهی
func (s *ModuleService) VerifyAd(adID string, verify bool) error {
    event := map[string]interface{}{
        "type":     "ad_verification",
        "ad_id":    adID,
        "verified": verify,
        "timestamp": time.Now(),
    }
    s.eventPublisher.Publish("ad.events", event)
    return nil
}

// دریافت آمار سیستم
func (s *ModuleService) GetSystemStats() (*models.AdminStats, error) {
    // TODO: جمع‌آوری آمار از همه سرویس‌ها
    stats := &models.AdminStats{
        TotalUsers:        12540,
        NewUsersToday:     234,
        TotalBusinesses:   856,
        PendingBusinesses: 23,
        TotalAds:          3421,
        PendingAds:        45,
        TotalRevenue:      12500000,
        RevenueToday:      450000,
        CommissionEarned:  1875000,
        WithdrawRequests:  12,
        ActiveUsersNow:    345,
    }
    return stats, nil
}

// دریافت تنظیمات سیستم
func (s *ModuleService) GetConfig(key string) (string, error) {
    config, err := s.configRepo.FindByKey(key)
    if err != nil {
        return "", err
    }
    return config.Value, nil
}

// تنظیم مقدار جدید
func (s *ModuleService) SetConfig(key, value, valueType string) error {
    config, err := s.configRepo.FindByKey(key)
    if err != nil {
        // ایجاد جدید
        config = &models.SystemConfig{
            Key:   key,
            Value: value,
            Type:  valueType,
        }
        return s.configRepo.Create(config)
    }
    
    config.Value = value
    config.Type = valueType
    return s.configRepo.Update(config)
}
