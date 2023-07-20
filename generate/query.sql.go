// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package generate

import (
	"context"
	"database/sql"
)

const countAuthors = `-- name: CountAuthors :one
SELECT count(*) FROM authors
`

func (q *Queries) CountAuthors(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAuthors)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING id, name, bio, created_at, updated_at, deleted_at
`

type CreateAuthorParams struct {
	Name string
	Bio  string
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio, created_at, updated_at, deleted_at FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio, created_at, updated_at, deleted_at FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1
RETURNING id, name, bio, created_at, updated_at, deleted_at
`

type UpdateAuthorParams struct {
	ID   int32
	Name string
	Bio  sql.NullString
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor, arg.ID, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
