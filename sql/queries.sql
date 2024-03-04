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

-- name: CreateInventory :exec
INSERT INTO inventories ( ID, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: GetInventories :many
SELECT * FROM inventories;

-- name: UpdateLocation :exec
UPDATE users SET latitude = $2, longitude = $3 WHERE id = $1;

-- name: CreateItem :exec
INSERT INTO items ( ID, description, score, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: GetItems :many
SELECT * FROM items;

-- name: CreateInfected :exec
INSERT INTO Infecteds ( user_id_reported, user_id_notified, created_at, updated_at)
VALUES ($1, $2, $3, $4);
