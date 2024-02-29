-- name: CreateUser :exec
INSERT INTO users ( ID, name, age, gender, latitude, longitude, infected, contamination_notification, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET name = $2, age = $3, gender = $4, latitude = $5, longitude = $6 WHERE id = $1;
