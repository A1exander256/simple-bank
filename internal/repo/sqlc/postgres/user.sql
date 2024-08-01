-- name: InsertUser :one
INSERT INTO users(
    guid, username, hashed_password, 
    email, created_at, updated_at
) VALUES(
    $1, $2, $3, $4, now(), now()
) RETURNING guid;