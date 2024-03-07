package postgres

import (
	"github.com/jmoiron/sqlx"
	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) PostgresDB {
	return &postgres{
		db: db,
	}
}
