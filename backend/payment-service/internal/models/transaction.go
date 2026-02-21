package models

import "time"

type Transaction struct {
ID string json:"id" db:"id"
UserID string json:"user_id" db:"user_id"
Amount int64 json:"amount" db:"amount"
Type string json:"type" db:"type"
Status string json:"status" db:"status"
CreatedAt time.Time json:"created_at" db:"created_at"
}


