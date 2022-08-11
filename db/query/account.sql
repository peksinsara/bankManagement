INSERT INTO
    accounts (
    owner,
    balance,
    currency
)
VALUES
    ('Sara', 100, '$') RETURNING *;