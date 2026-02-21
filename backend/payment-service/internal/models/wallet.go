package models

type Wallet struct {
ID string json:"id" db:"id"
UserID string json:"user_id" db:"user_id"
Balance int64 json:"balance" db:"balance"
}


