-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;


-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;


-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;


-- name: UpdateAuthor :one
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1
RETURNING *;


-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-----------------------------------------
-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;


-- name: ListAuthorsByIDs :many
SELECT * FROM authors
WHERE id = ANY($1::int[]);


-- name: CountAuthors :one
SELECT count(*) FROM authors;


-- name: UpdateAuthor :exec
UPDATE authors SET bio = $2
WHERE id = $1;

