package models

import "time"

type User struct {
ID string json:"id" db:"id"
Phone string json:"phone" db:"phone"
Name string json:"name" db:"name"
Email string json:"email" db:"email"
CreatedAt time.Time json:"created_at" db:"created_at"
}


