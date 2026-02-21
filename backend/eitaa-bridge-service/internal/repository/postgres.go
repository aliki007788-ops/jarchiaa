package repository

import "database/sql"

type EitaaRepository struct {
db *sql.DB
}

func NewEitaaRepository(db *sql.DB) *EitaaRepository {
return &EitaaRepository{db: db}
}


