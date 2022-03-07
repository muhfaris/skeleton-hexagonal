package accountrepo

import (
	"context"

	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
)

func (q *AccountRepo) FindByEmail(ctx context.Context, email string) <-chan repository.Result {
	result := make(chan repository.Result)
	go func() {
		var account model.Account
		sql := `
			SELECT
				"id",
				"full_name",
				"username",
				"email",
				"password",
				"role"
			FROM "public"."accounts"
			WHERE 
				"email" = $1; `

		if err := q.db.QueryRow(ctx, sql, email).Scan(
			&account.ID,
			&account.FullName,
			&account.Email,
			&account.Username,
			&account.Password,
			&account.Role,
		); err != nil {
			result <- repository.Result{Error: err}
			return
		}

		result <- repository.Result{Data: account}
	}()

	return result
}
