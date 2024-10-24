-- name: CreateOneUser :exec
INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateOneUser :exec
UPDATE users SET first_name=$2, last_name=$3
WHERE id=$1;
