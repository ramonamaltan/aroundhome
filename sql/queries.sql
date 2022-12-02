-- name: GetPartner :one
SELECT * FROM partners
WHERE id = $1 LIMIT 1;

-- name: ListPartners :many
SELECT * FROM partners;
-- ORDER BY name;
