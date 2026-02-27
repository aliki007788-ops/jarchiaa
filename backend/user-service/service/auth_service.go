package service

import (
    "crypto/rand"
    "fmt"
    "time"
    "user-service/models"
    "user-service/repository"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo  *repository.UserRepository
    sessionRepo *repository.SessionRepository
    jwtSecret []byte
}

func NewAuthService(userRepo *repository.UserRepository, sessionRepo *repository.SessionRepository, jwtSecret string) *AuthService {
    return &AuthService{
        userRepo:    userRepo,
        sessionRepo: sessionRepo,
        jwtSecret:   []byte(jwtSecret),
    }
}

// تولید کد OTP 6 رقمی
func (s *AuthService) GenerateOTP() string {
    b := make([]byte, 3)
    rand.Read(b)
    return fmt.Sprintf("%06d", (int(b[0])<<16|int(b[1])<<8|int(b[2]))%1000000)
}

// ارسال OTP (اینجا شبیه‌سازی شده)
func (s *AuthService) SendOTP(phone, code string) error {
    // TODO: اتصال به سرویس پیامک (Kavenegar/Farazsms)
    fmt.Printf("📱 OTP for %s: %s\n", phone, code)
    return nil
}

// ثبت‌نام یا ورود با شماره تلفن
func (s *AuthService) LoginOrRegister(req *models.LoginRequest) error {
    user, err := s.userRepo.FindByPhone(req.Phone)
    
    if err != nil {
        // کاربر جدید
        user = &models.User{
            Phone:      req.Phone,
            Role:       "user",
            IsActive:   true,
            IsVerified: false,
            Wallet:     0,
        }
        if err := s.userRepo.Create(user); err != nil {
            return err
        }
    }
    
    // تولید و ارسال OTP
    code := s.GenerateOTP()
    // TODO: ذخیره OTP در Redis با زمان 2 دقیقه
    return s.SendOTP(req.Phone, code)
}

// تأیید OTP و تولید توکن
func (s *AuthService) VerifyOTP(req *models.VerifyOTPRequest) (*models.AuthResponse, error) {
    // TODO: بررسی OTP از Redis
    // if storedCode != req.Code { return nil, errors.New("کد نامعتبر") }
    
    user, err := s.userRepo.FindByPhone(req.Phone)
    if err != nil {
        return nil, err
    }
    
    // تولید JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID.String(),
        "phone":   user.Phone,
        "role":    user.Role,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })
    
    tokenString, err := token.SignedString(s.jwtSecret)
    if err != nil {
        return nil, err
    }
    
    // تولید Refresh Token
    refreshToken := s.GenerateOTP() + fmt.Sprintf("%d", time.Now().UnixNano())
    
    // ذخیره سشن
    session := &models.Session{
        UserID:       user.ID,
        Token:        tokenString,
        RefreshToken: refreshToken,
        ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
    }
    s.sessionRepo.Create(session)
    
    return &models.AuthResponse{
        Token:        tokenString,
        RefreshToken: refreshToken,
        User:         *user,
    }, nil
}

// ثبت‌نام کامل (بعد از OTP)
func (s *AuthService) CompleteRegistration(req *models.RegisterRequest) (*models.User, error) {
    user, err := s.userRepo.FindByPhone(req.Phone)
    if err != nil {
        return nil, err
    }
    
    user.FullName = req.FullName
    user.Email = req.Email
    user.IsVerified = true
    
    if err := s.userRepo.Update(user); err != nil {
        return nil, err
    }
    
    return user, nil
}

// اعتبارسنجی توکن
func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return s.jwtSecret, nil
    })
}

// دریافت کاربر از توکن
func (s *AuthService) GetUserFromToken(tokenString string) (*models.User, error) {
    token, err := s.ValidateToken(tokenString)
    if err != nil {
        return nil, err
    }
    
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    
    userID, _ := uuid.Parse(claims["user_id"].(string))
    return s.userRepo.FindByID(userID)
}