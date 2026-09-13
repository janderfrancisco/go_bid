-- name: CreatePerson :one
INSERT INTO people (first_name, last_name)
VALUES ($1, $2)
RETURNING id, first_name, last_name;

-- name: GetPerson :one
SELECT id, first_name, last_name
FROM people
WHERE id = $1
LIMIT 1;

-- name: ListPeople :many
SELECT id, first_name, last_name
FROM people
ORDER BY id;

-- name: UpdatePerson :one
UPDATE people
SET first_name = $1,
    last_name = $2
WHERE id = $3
RETURNING id, first_name, last_name;

-- name: DeletePerson :exec
DELETE FROM people
WHERE id = $1;
