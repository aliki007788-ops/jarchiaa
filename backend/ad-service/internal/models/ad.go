package models

type Ad struct {
ID string json:"id" db:"id"
BusinessID string json:"business_id" db:"business_id"
Title string json:"title" db:"title"
MediaURLs []string json:"media_urls" db:"media_urls"
Price int64 json:"price" db:"price"
}


