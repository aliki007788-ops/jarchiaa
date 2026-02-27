package models

import (
    "time"
    "github.com/google/uuid"
)

type Module struct {
    ID          int       `json:"id" db:"id"`
    Key         string    `json:"key" db:"key"`
    Name        string    `json:"name" db:"name"`
    NameFa      string    `json:"name_fa" db:"name_fa"`
    Description string    `json:"description" db:"description"`
    Icon        string    `json:"icon" db:"icon"`
    IsActive    bool      `json:"is_active" db:"is_active"`
    IsCore      bool      `json:"is_core" db:"is_core"`
    Order       int       `json:"order" db:"order"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type SystemConfig struct {
    ID           string    `json:"id" db:"id"`
    Key          string    `json:"key" db:"key"`
    Value        string    `json:"value" db:"value"`
    Type         string    `json:"type" db:"type"` // string, int, bool, json
    Description  string    `json:"description" db:"description"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type PlanConfig struct {
    ID              int     `json:"id" db:"id"`
    Name            string  `json:"name" db:"name"`
    NameFa          string  `json:"name_fa" db:"name_fa"`
    Price           int64   `json:"price" db:"price"`
    MaxAds          int     `json:"max_ads" db:"max_ads"`
    DurationDays    int     `json:"duration_days" db:"duration_days"`
    Modules         []string `json:"modules" db:"modules"` // آرایه‌ای از کلید ماژول‌ها
    IsActive        bool    `json:"is_active" db:"is_active"`
    IsPopular       bool    `json:"is_popular" db:"is_popular"`
    DiscountPercent int     `json:"discount_percent" db:"discount_percent"`
}

type AdminStats struct {
    TotalUsers          int64   `json:"total_users"`
    NewUsersToday       int64   `json:"new_users_today"`
    TotalBusinesses     int64   `json:"total_businesses"`
    PendingBusinesses   int64   `json:"pending_businesses"`
    TotalAds            int64   `json:"total_ads"`
    PendingAds          int64   `json:"pending_ads"`
    TotalTransactions   int64   `json:"total_transactions"`
    TotalRevenue        int64   `json:"total_revenue"`
    RevenueToday        int64   `json:"revenue_today"`
    CommissionEarned    int64   `json:"commission_earned"`
    WithdrawRequests    int64   `json:"withdraw_requests"`
    SystemUptime        string  `json:"system_uptime"`
    ActiveUsersNow      int64   `json:"active_users_now"`
}
