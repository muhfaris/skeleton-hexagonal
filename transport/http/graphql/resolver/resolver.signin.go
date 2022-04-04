package resolver

import (
	"context"
	"net/http"

	rslvmodel "github.com/muhfaris/skeleton-hexagonal/transport/http/graphql/resolver/model"
	"github.com/muhfaris/skeleton-hexagonal/transport/structures"
)

func (r *Resolver) SignIn(ctx context.Context, data structures.DataLoginRead) (*rslvmodel.AccountDataResolver, error) {
	account, err := r.UserPublicService.Login(ctx, data.Data)
	if err != nil {
		return &rslvmodel.AccountDataResolver{}, err
	}

	return &rslvmodel.AccountDataResolver{
		StatusRslv: http.StatusOK,
		DataRslv: rslvmodel.AccountResolver{
			Id:          account.ID,
			Fullname:    account.FullName,
			Username:    account.Username,
			Email:       account.Email,
			AccessToken: account.AccessToken,
		},
	}, nil
}
