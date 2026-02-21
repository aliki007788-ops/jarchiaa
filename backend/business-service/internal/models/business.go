package models

type Business struct {
ID string json:"id" db:"id"
UserID string json:"user_id" db:"user_id"
Name string json:"name" db:"name"
Verified bool json:"verified" db:"verified"
}


