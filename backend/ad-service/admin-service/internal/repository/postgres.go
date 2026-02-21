package repository

import "database/sql"

type AdminRepository struct {
db *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
return &AdminRepository{db: db}
}


