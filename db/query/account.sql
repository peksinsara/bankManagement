-- name: CreateAccount :one
INSERT INTO
    accounts (
    owner,
    balance,
    currency
)
VALUES
    ('Sara', 100, '$') RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = 1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE owner = 'Sara'
ORDER BY id
    LIMIT 2
OFFSET 3;

/*couldnt run :exec*/
-- name: UpdateAccount :one
UPDATE accounts
SET balance = 999
WHERE owner = 'Sara'
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE owner = 'Sara';

