package accountrepo

import (
	"context"

	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func (q *AccountRepo) FindByEmail(ctx context.Context, email string) <-chan repository.Result {
	result := make(chan repository.Result)
	go func() {
		var account model.Account
		err := q.db.Database("skeleton_db").Collection("accounts").FindOne(ctx, bson.M{"email": email}).Decode(&account)
		if err != nil {
			result <- repository.Result{Error: err}
			return
		}

		result <- repository.Result{Data: account}
	}()

	return result
}
