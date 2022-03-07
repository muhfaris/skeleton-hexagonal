package userpublicrepo

import (
	"context"

	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
)

func (q *UserPublicRepo) SignUp(ctx context.Context, account *model.Account) <-chan repository.Result {
	result := make(chan repository.Result)
	go func() {
		sql := `
			INSERT INTO "public"."accounts" ( "id", "full_name", "username", "email", "password", "role", "created_at") 
			VALUES ($1 , $2, $3, $4, $5, $6, now() ); `
		_, err := q.db.Exec(ctx, sql, account.ID, account.FullName, account.Username, account.Email, account.HashPassword, account.Role)
		if err != nil {
			result <- repository.Result{Error: err}
			return
		}

		result <- repository.Result{Data: account}
	}()

	return result
}
