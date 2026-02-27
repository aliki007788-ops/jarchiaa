package models

import (
    "time"
    "github.com/google/uuid"
)

type Business struct {
    ID              uuid.UUID `json:"id" db:"id"`
    UserID          uuid.UUID `json:"user_id" db:"user_id"`
    Name            string    `json:"name" db:"name"`
    Category        string    `json:"category" db:"category"`
    Description     string    `json:"description" db:"description"`
    Address         string    `json:"address" db:"address"`
    Phone           string    `json:"phone" db:"phone"`
    Location        string    `json:"location" db:"location"` // POINT(lat lng)
    PlanID          int       `json:"plan_id" db:"plan_id"`
    PlanExpiresAt   time.Time `json:"plan_expires_at" db:"plan_expires_at"`
    IsVerified      bool      `json:"is_verified" db:"is_verified"`
    IsActive        bool      `json:"is_active" db:"is_active"`
    TotalEarned     int64     `json:"total_earned" db:"total_earned"`
    CommissionRate  float64   `json:"commission_rate" db:"commission_rate"`
    CreatedAt       time.Time `json:"created_at" db:"created_at"`
    UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Plan struct {
    ID              int     `json:"id" db:"id"`
    Name            string  `json:"name" db:"name"`
    NameFa          string  `json:"name_fa" db:"name_fa"`
    Price           int64   `json:"price" db:"price"`
    MaxAds          int     `json:"max_ads" db:"max_ads"`
    DurationDays    int     `json:"duration_days" db:"duration_days"`
    HasGeoTargeting bool    `json:"has_geo_targeting" db:"has_geo_targeting"`
    HasAnalytics    bool    `json:"has_analytics" db:"has_analytics"`
    HasPriority     bool    `json:"has_priority" db:"has_priority"`
    IsActive        bool    `json:"is_active" db:"is_active"`
}

type RegisterBusinessRequest struct {
    Name        string `json:"name"`
    Category    string `json:"category"`
    Description string `json:"description"`
    Address     string `json:"address"`
    Phone       string `json:"phone"`
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
}

type PurchasePlanRequest struct {
    PlanID int `json:"plan_id"`
}

type BusinessStats struct {
    TotalAds        int     `json:"total_ads"`
    ActiveAds       int     `json:"active_ads"`
    TotalViews      int64   `json:"total_views"`
    TotalClicks     int64   `json:"total_clicks"`
    TotalLikes      int64   `json:"total_likes"`
    TotalComments   int64   `json:"total_comments"`
    ViewsToday      int64   `json:"views_today"`
    ClicksToday     int64   `json:"clicks_today"`
    EngagementRate  float64 `json:"engagement_rate"`
    WalletBalance   int64   `json:"wallet_balance"`
}
