// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package core

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, email_status, pas_hash, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, email_status, pas_hash, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, email_status, pas_hash, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users (
    id,
    username,
    email,
    email_status,
    pas_hash
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING id, username, email, email_status, pas_hash, created_at
`

type InsertUserParams struct {
	ID          uuid.UUID
	Username    sql.NullString
	Email       sql.NullString
	EmailStatus sql.NullBool
	PasHash     sql.NullString
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.EmailStatus,
		arg.PasHash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const listUsersByID = `-- name: ListUsersByID :many
SELECT id, username, email, email_status, pas_hash, created_at FROM users
ORDER BY id
    LIMIT $1
OFFSET $2
`

type ListUsersByIDParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsersByID(ctx context.Context, arg ListUsersByIDParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.EmailStatus,
			&i.PasHash,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByUsername = `-- name: ListUsersByUsername :many
SELECT id, username, email, email_status, pas_hash, created_at FROM users
ORDER BY username
    LIMIT $1
OFFSET $2
`

type ListUsersByUsernameParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsersByUsername(ctx context.Context, arg ListUsersByUsernameParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByUsername, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.EmailStatus,
			&i.PasHash,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmailStatusByID = `-- name: UpdateEmailStatusByID :one
UPDATE users
SET email_status = $2
WHERE id = $1
    RETURNING id, username, email, email_status, pas_hash, created_at
`

type UpdateEmailStatusByIDParams struct {
	ID          uuid.UUID
	EmailStatus sql.NullBool
}

func (q *Queries) UpdateEmailStatusByID(ctx context.Context, arg UpdateEmailStatusByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateEmailStatusByID, arg.ID, arg.EmailStatus)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const updatePasswordByID = `-- name: UpdatePasswordByID :one
UPDATE users
SET pas_hash = $2
WHERE id = $1
    RETURNING id, username, email, email_status, pas_hash, created_at
`

type UpdatePasswordByIDParams struct {
	ID      uuid.UUID
	PasHash sql.NullString
}

func (q *Queries) UpdatePasswordByID(ctx context.Context, arg UpdatePasswordByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePasswordByID, arg.ID, arg.PasHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE users
SET
    email = $2,
    email_status = $3,
    username = $4
WHERE id = $1
    RETURNING id, username, email, email_status, pas_hash, created_at
`

type UpdateUserByIDParams struct {
	ID          uuid.UUID
	Email       sql.NullString
	EmailStatus sql.NullBool
	Username    sql.NullString
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserByID,
		arg.ID,
		arg.Email,
		arg.EmailStatus,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}

const updateUsernameByID = `-- name: UpdateUsernameByID :one
UPDATE users
SET username = $2
WHERE id = $1
    RETURNING id, username, email, email_status, pas_hash, created_at
`

type UpdateUsernameByIDParams struct {
	ID       uuid.UUID
	Username sql.NullString
}

func (q *Queries) UpdateUsernameByID(ctx context.Context, arg UpdateUsernameByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUsernameByID, arg.ID, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PasHash,
		&i.CreatedAt,
	)
	return i, err
}
