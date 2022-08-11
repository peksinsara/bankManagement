// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO
    accounts (
    owner,
    balance,
    currency
)
VALUES
    ('Sara', 100, '$') RETURNING id, owner, balance, currency, created_at
`

func (q *Queries) CreateAccount(ctx context.Context) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE owner = 'Sara'
`

func (q *Queries) DeleteAccount(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAccount)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = 1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE owner = 'Sara'
ORDER BY id
    LIMIT 2
OFFSET 3
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET balance = 999
WHERE owner = 'Sara'
RETURNING id, owner, balance, currency, created_at
`

//couldnt run :exec
func (q *Queries) UpdateAccount(ctx context.Context) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}