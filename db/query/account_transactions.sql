-- name: CreateTransaction :one 
INSERT INTO "transactions" (
    account_id,
    amount,
    type
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccountTransactions :many 
SELECT * FROM "transactions" WHERE account_id = $1 LIMIT $2 OFFSET $3;;
