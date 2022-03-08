package accountrepo

import (
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepo struct {
	db *mongo.Client
}

func NewAccountRepo(client *mongo.Client) repository.AccountRepository {
	return &AccountRepo{db: client}
}
