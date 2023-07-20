package internal

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"github.com/SQLC/generate"
	_ "github.com/lib/pq"
)

func Run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "postgres://yunus:godev@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}

	queries := generate.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, generate.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}
