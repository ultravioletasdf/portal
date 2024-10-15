-- name: GetApp :one
SELECT * FROM apps WHERE id = ?;
