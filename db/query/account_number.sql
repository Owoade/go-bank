-- name: CreateAccountNumber :one 
INSERT INTO "account_numbers" (
    account_id,
    account_name,
    account_number,
    bank_name
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserAccountNumbers :many 
SELECT * FROM "account_numbers" WHERE account_id = $1 LIMIT 1;