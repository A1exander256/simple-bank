// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package query

import (
	"context"

	"github.com/google/uuid"
)

const insertUser = `-- name: InsertUser :one
INSERT INTO users(
    guid, username, hashed_password, 
    email, created_at, updated_at
) VALUES(
    $1, $2, $3, $4, now(), now()
) RETURNING guid
`

type InsertUserParams struct {
	Guid           uuid.UUID
	Username       string
	HashedPassword string
	Email          string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.insertUserStmt, insertUser,
		arg.Guid,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
	)
	var guid uuid.UUID
	err := row.Scan(&guid)
	return guid, err
}
