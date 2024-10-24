-- name: CreateVerificationUser :exec
INSERT INTO verifications (id, email, code, expires_at, type)
VALUES ($1, $2, $3, $4, $5);

-- name: GetVerificationUser :one
SELECT * FROM verifications
WHERE email=$1;

-- name: DeleteVerificationUser :exec
DELETE FROM verifications
WHERE email=$1;
