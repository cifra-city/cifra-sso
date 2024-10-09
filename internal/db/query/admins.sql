-- name: GetAdminByID :one
SELECT * FROM admins
WHERE id = $1 LIMIT 1;

-- name: GetAdminByUID :one
SELECT * FROM admins
WHERE uid = $1 LIMIT 1;

-- name: ListAdmins :many
SELECT * FROM admins
ORDER BY created_at;

-- name: InsertAdmin :one
INSERT INTO admins (
    uid,
    name
) VALUES (
     $1, $2
) RETURNING *;

-- name: UpdateAdminByID :one
UPDATE admins
SET
    name = $2,
    uid = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteAdminByID :exec
DELETE FROM admins
WHERE id = $1;