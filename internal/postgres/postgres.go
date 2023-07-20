package internal

import (
	"context"
	"database/sql"

	"github.com/SQLC/internal/repo"
)

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

func (q *Queries) CreateAuthor(ctx context.Context, arg *repo.AuthorParams) (*repo.Author, error) {
	var i repo.Author

	const createAuthor = `-- name: CreateAuthor :one
	INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING id, name, bio, created_at`
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)

	err := row.Scan(&i.ID, &i.Name, &i.Bio, &i.CreatedAt)
	return &i, err
}

func (q *Queries) CountAuthors(ctx context.Context) (int, error) {
	var i int

	const countAuthors = `-- name: CountAuthors :one
	SELECT count(*) FROM authors`

	row := q.db.QueryRowContext(ctx, countAuthors)

	err := row.Scan(&i)
	return i, err
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg *repo.AuthorParams) error {

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
