-- name: CreateUser :one
INSERT INTO "users" (
    email,
    password
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetOneUserByEmail :one
SELECT  * FROM "users" WHERE email = $1 LIMIT 1;

-- name: GetOneUserById :one
SELECT  * FROM "users" WHERE  id = $1 LIMIT 1;

-- name: UpdateUserPassword :exec
UPDATE "users" SET password = $1 WHERE email = $2;

