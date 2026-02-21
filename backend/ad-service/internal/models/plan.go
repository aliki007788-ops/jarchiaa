package models

type Plan struct {
ID string json:"id" db:"id"
Name string json:"name" db:"name"
NameFa string json:"name_fa" db:"name_fa"
Price int64 json:"price" db:"price"
MaxAds int json:"max_ads" db:"max_ads"
}


