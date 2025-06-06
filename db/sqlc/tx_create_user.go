package db

import (
	"context"
	"fmt"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error // callback to run after user is successfully created
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			return err
		}

		return arg.AfterCreate(result.User)
	})

	return result, err
}
