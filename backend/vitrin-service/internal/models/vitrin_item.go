package models

type VitrinItem struct {
ID string json:"id" db:"id"
BusinessID string json:"business_id" db:"business_id"
Title string json:"title" db:"title"
MediaURL string json:"media_url" db:"media_url"
SizeType string json:"size_type" db:"size_type"
Views int64 json:"views" db:"views"
}


