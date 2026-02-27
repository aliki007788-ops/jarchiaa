package models

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID           uuid.UUID `json:"id" db:"id"`
    Phone        string    `json:"phone" db:"phone"`
    FullName     string    `json:"full_name" db:"full_name"`
    Email        string    `json:"email" db:"email"`
    Role         string    `json:"role" db:"role"` // user, business, admin
    Wallet       int64     `json:"wallet" db:"wallet"`
    TrustScore   float64   `json:"trust_score" db:"trust_score"`
    IsVerified   bool      `json:"is_verified" db:"is_verified"`
    IsActive     bool      `json:"is_active" db:"is_active"`
    LastSeen     time.Time `json:"last_seen" db:"last_seen"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Session struct {
    ID           uuid.UUID `json:"id" db:"id"`
    UserID       uuid.UUID `json:"user_id" db:"user_id"`
    Token        string    `json:"token" db:"token"`
    RefreshToken string    `json:"refresh_token" db:"refresh_token"`
    ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type LoginRequest struct {
    Phone string `json:"phone"`
}

type VerifyOTPRequest struct {
    Phone string `json:"phone"`
    Code  string `json:"code"`
}

type RegisterRequest struct {
    Phone    string `json:"phone"`
    FullName string `json:"full_name"`
    Email    string `json:"email"`
}

type AuthResponse struct {
    Token        string `json:"token"`
    RefreshToken string `json:"refresh_token"`
    User         User   `json:"user"`
}
