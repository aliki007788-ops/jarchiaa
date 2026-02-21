package types

type Ad struct {
ID string json:"id"
BusinessID string json:"business_id"
Title string json:"title"
MediaURLs []string json:"media_urls"
Price int64 json:"price"
}


