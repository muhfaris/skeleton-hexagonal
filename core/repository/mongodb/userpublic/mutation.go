package userpublicrepo

import (
	"context"

	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func (q *UserPublicRepo) SignUp(ctx context.Context, account *model.Account) <-chan repository.Result {
	result := make(chan repository.Result)

	go func() {
		collection := q.db.Database("skeleton_db").Collection("accounts")

		data := bson.M{
			"id":        account.ID,
			"full_name": account.FullName,
			"username":  account.Username,
			"email":     account.Email,
			"password":  account.HashPassword,
			"role":      account.Role,
		}

		_, err := collection.InsertOne(ctx, data)
		if err != nil {
			result <- repository.Result{Error: err}
			return
		}

		result <- repository.Result{Data: account}
	}()

	return result
}
