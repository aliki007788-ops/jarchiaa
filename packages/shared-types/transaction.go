package types

import "time"

type Transaction struct {
ID string json:"id"
UserID string json:"user_id"
Amount int64 json:"amount"
Type string json:"type"
Status string json:"status"
CreatedAt time.Time json:"created_at"
}


