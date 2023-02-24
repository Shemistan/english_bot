package storage

import (
	"github.com/Shemistan/english_bot/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) domain.Storage {
	return &storage{db: db}
}
