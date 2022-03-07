package app

import (
	"github.com/jackc/pgx/v4"
	accountrepo "github.com/muhfaris/skeleton-hexagonal/core/repository/mysql/account"
	userpublicrepo "github.com/muhfaris/skeleton-hexagonal/core/repository/mysql/userpublic"
	"github.com/muhfaris/skeleton-hexagonal/core/services"
)

type ServiceApp struct {
	UserPublicService services.UserPublicService
}

func NewServiceApp(db *pgx.Conn) *ServiceApp {
	userPublicRepo := userpublicrepo.NewUserPublicRepo(db)
	accountRepo := accountrepo.NewAccountRepo(db)

	return &ServiceApp{
		UserPublicService: services.NewUserPublicnService(userPublicRepo, accountRepo),
	}
}
