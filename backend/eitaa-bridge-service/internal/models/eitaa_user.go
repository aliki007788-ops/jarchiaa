package models

type EitaaUser struct {
ID string json:"id" db:"id"
UserID string json:"user_id" db:"user_id"
EitaaID string json:"eitaa_id" db:"eitaa_id"
Username string json:"username" db:"username"
}


