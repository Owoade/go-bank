-- name: CreateAccount :one 
INSERT INTO "accounts" (
    user_id,
    balance
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUserAccount :one 
SELECT * FROM "accounts" WHERE user_id = $1 LIMIT 1;

-- name: GetAccountById :one
SELECT * FROM "accounts" WHERE id = $1 LIMIT 1;

-- name: CreditAccount :one
UPDATE "accounts" SET balance = balance + $1 WHERE id = $2 RETURNING *; 

-- name: DebitAccount :one 
UPDATE "accounts" SET balance = balance - $1 WHERE id = $2 RETURNING *;