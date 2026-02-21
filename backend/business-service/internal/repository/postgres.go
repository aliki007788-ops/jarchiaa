package repository

import "database/sql"

type BusinessRepository struct {
db *sql.DB
}

func NewBusinessRepository(db *sql.DB) *BusinessRepository {
return &BusinessRepository{db: db}
}


