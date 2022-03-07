package userpublicrepo

import (
	"context"

	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
)

func (q *UserPublicRepo) SignIn(ctx context.Context, username, password string) (model.Account, error) {
	return model.Account{}, nil
}
