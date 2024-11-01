-- name: GetUserByID :one
SELECT * FROM users_secret
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users_secret
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users_secret
WHERE email = $1 LIMIT 1;

-- name: ListUsersByID :many
SELECT * FROM users_secret
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: ListUsersByUsername :many
SELECT * FROM users_secret
ORDER BY username
    LIMIT $1
OFFSET $2;


-- name: UpdateUserByID :one
UPDATE users_secret
SET
    email = $2,
    email_status = $3,
    username = $4
WHERE id = $1
    RETURNING *;

-- name: UpdateUsernameByID :one
UPDATE users_secret
SET username = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateEmailStatusByID :one
UPDATE users_secret
SET email_status = $2
WHERE id = $1
    RETURNING *;

-- name: UpdatePasswordByID :one
UPDATE users_secret
SET pass_hash = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateEmailByID :one
UPDATE users_secret
SET
    email = $2,
    email_status = $3
WHERE id = $1
    RETURNING *;

-- name: InsertUser :one
INSERT INTO users_secret (
    id,
    username,
    email,
    email_status,
    pass_hash
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;


-- name: DeleteUserByID :exec
DELETE FROM users_secret
WHERE id = $1;