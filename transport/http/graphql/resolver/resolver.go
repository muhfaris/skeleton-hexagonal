package resolver

import (
	"github.com/jackc/pgx/v4"
	accountrepo "github.com/muhfaris/skeleton-hexagonal/core/repository/mysql/account"
	userpublicrepo "github.com/muhfaris/skeleton-hexagonal/core/repository/mysql/userpublic"
	"github.com/muhfaris/skeleton-hexagonal/core/services"
)

type Resolver struct {
	UserPublicService services.UserPublicService
}

func NewResolver(db *pgx.Conn) *Resolver {
	userPublicRepo := userpublicrepo.NewUserPublicRepo(db)
	accountRepo := accountrepo.NewAccountRepo(db)

	return &Resolver{
		UserPublicService: services.NewUserPublicnService(userPublicRepo, accountRepo),
	}
}
