package repository

import (
"database/sql"
_ "github.com/lib/pq"
)

func NewPostgresDB(connStr string) (*sql.DB, error) {
return sql.Open("postgres", connStr)
}


