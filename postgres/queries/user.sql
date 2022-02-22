-- name: GetAllUsers :many
SELECT * FROM "users";

-- name: GetAUserById :one
SELECT * FROM "users" where id = $1 LIMIT 1;

-- name: GetAUserByEmail :one
SELECT * FROM "users" where email = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO "users" (username, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE "users" SET username = $1, email = $2 WHERE id = $3 RETURNING *;

-- name: DeleteUser :exec
DELETE from "users" where id = $1;
