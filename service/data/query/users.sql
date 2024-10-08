-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;


-- name: ListUsersByID :many
SELECT * FROM users
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: ListUsersByUsername :many
SELECT * FROM users
ORDER BY username
    LIMIT $1
OFFSET $2;


-- name: UpdateUserByID :one
UPDATE users
SET
    email = $2,
    email_status = $3,
    username = $4
WHERE id = $1
    RETURNING *;

-- name: UpdateUsernameByID :one
UPDATE users
SET username = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateEmailStatusByID :one
UPDATE users
SET email_status = $2
WHERE id = $1
    RETURNING *;

-- name: UpdatePasswordByID :one
UPDATE users
SET pas_hash = $2
WHERE id = $1
    RETURNING *;


-- name: InsertUser :one
INSERT INTO users (
    id,
    username,
    email,
    email_status,
    pas_hash
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING *;


-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;