package accountrepo

import (
	"github.com/jackc/pgx/v4"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
)

type AccountRepo struct {
	db *pgx.Conn
}

func NewAccountRepo(db *pgx.Conn) repository.AccountRepository {
	return &AccountRepo{db: db}
}
