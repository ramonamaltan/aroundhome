-- name: GetPartner :one
SELECT * FROM partners
WHERE id = $1 LIMIT 1;

-- name: ListPartners :many
SELECT * FROM partners
WHERE servicename = $1 AND material like $2
ORDER BY radius ASC;

-- name: CreatePartner :one
INSERT INTO partners (
    partnername, servicename, latitude, longitude, material, radius, rating
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
