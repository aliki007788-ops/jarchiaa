package service

import (
"time"
"github.com/redis/go-redis/v9"
)

type OTPService struct {
redis *redis.Client
}

func (s *OTPService) StoreOTP(phone, otp string) error {
return s.redis.Set(ctx, "otp:"+phone, otp, 2*time.Minute).Err()
}


