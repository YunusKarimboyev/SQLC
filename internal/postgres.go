package internal

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Author struct {
	ID   int
	Name string
	Bio  sql.NullString
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

type AuthorParams struct {
	ID   int
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg AuthorParams) (Author, error) {
	var i Author

	const createAuthor = `-- name: CreateAuthor :one
	INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING id, name, bio`
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}


func (q *Queries) ListAuthorsByIDs(ctx context.Context, ids []int) ([]Author, error) {
	var items []Author

	const listAuthors = `-- name: ListAuthorsByIDs :many
	SELECT id, bio, birth_year FROM authors WHERE id = ANY($1::int[])`

	rows, err := q.db.QueryContext(ctx, listAuthors, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Bio, &i.Name); err != nil {
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


func (q *Queries) CountAuthors(ctx context.Context) (int, error) {
	var i int

	const countAuthors = `-- name: CountAuthors :one
	SELECT count(*) FROM authors`

	row := q.db.QueryRowContext(ctx, countAuthors)

	err := row.Scan(&i)
	return i, err
}


func (q *Queries) UpdateAuthor(ctx context.Context, arg AuthorParams) error {

	const updateAuthor = `-- name: UpdateAuthor :exec
	UPDATE authors SET bio = $2 WHERE id = $1`

	_, err := q.db.ExecContext(ctx, updateAuthor, arg.ID, arg.Bio)
	return err
}


func (q *Queries) DeleteAuthor(ctx context.Context, id int) error {

	const deleteAuthor = `-- name: DeleteAuthor :exec
	DELETE FROM authors WHERE id = $1`
	
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}
