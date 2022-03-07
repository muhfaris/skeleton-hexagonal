package repository

import (
	"context"

	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
)

// Result wrap response query
type Result struct {
	Data       interface{}
	Error      error
	Pagination interface{}
}

// Interface name in repository must to general case not use prefix to spesified of data source

// UserPublicRepository is wrap for user public case
type UserPublicRepository interface {
	SignIn(ctx context.Context, username, password string) (model.Account, error)
	SignUp(ctx context.Context, account *model.Account) <-chan Result
}

type AccountRepository interface {
	FindByEmail(ctx context.Context, email string) <-chan Result
}
