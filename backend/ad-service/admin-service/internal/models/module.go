package models

type Module struct {
ID string json:"id" db:"id"
Key string json:"key" db:"key"
NameFa string json:"name_fa" db:"name_fa"
IsActive bool json:"is_active" db:"is_active"
}


