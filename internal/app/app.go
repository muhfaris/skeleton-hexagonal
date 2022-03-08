package app

import (
	accountrepo "github.com/muhfaris/skeleton-hexagonal/core/repository/mongodb/account"
	userpublicrepo "github.com/muhfaris/skeleton-hexagonal/core/repository/mongodb/userpublic"
	"github.com/muhfaris/skeleton-hexagonal/core/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceApp struct {
	UserPublicService services.UserPublicService
}

func NewServiceApp(db *mongo.Client) *ServiceApp {
	userPublicRepo := userpublicrepo.NewUserPublicRepo(db)
	accountRepo := accountrepo.NewAccountRepo(db)

	return &ServiceApp{
		UserPublicService: services.NewUserPublicnService(userPublicRepo, accountRepo),
	}
}
