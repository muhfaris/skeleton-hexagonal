package userpublicrepo

import (
	"github.com/jackc/pgx/v4"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
)

type UserPublicRepo struct {
	db *pgx.Conn
}

func NewUserPublicRepo(db *pgx.Conn) repository.UserPublicRepository {
	return &UserPublicRepo{db: db}
}
