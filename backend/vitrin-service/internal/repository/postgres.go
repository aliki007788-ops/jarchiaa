package repository

import "database/sql"

type VitrinRepository struct {
db *sql.DB
}

func NewVitrinRepository(db *sql.DB) *VitrinRepository {
return &VitrinRepository{db: db}
}


