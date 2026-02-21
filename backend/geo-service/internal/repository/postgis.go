package repository

import "database/sql"

type GeoRepository struct {
db *sql.DB
}

func NewGeoRepository(db *sql.DB) *GeoRepository {
return &GeoRepository{db: db}
}


