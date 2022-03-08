package userpublicrepo

import (
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserPublicRepo struct {
	db *mongo.Client
}

func NewUserPublicRepo(db *mongo.Client) repository.UserPublicRepository {
	return &UserPublicRepo{db: db}
}
