package internal

import (
	"context"
	"database/sql"
	"log"

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

	// create an author
	createResponse, err := queries.CreateAuthor(ctx, generate.CreateAuthorParams{
		Name: "Brian treysi",
		Bio:  "C Programming Language",
	})
	if err != nil {
		return err
	}
	log.Println(createResponse) // create author

	// list all authors
	// authorsList, err := queries.ListAuthors(ctx)
	// if err != nil {
	// 	return err
	// }
	// log.Println(authorsList)    // get all authors

	// get the author by id
	// getID, err := queries.GetAuthor(ctx, 4)
	// if err != nil {
	// 	return err
	// }
	// log.Println(getID)          // get author by ID

	// update  author
	// updateResponse, err := queries.UpdateAuthor(ctx, generate.UpdateAuthorParams{
	// 	ID:     2,
	// 	Name:   "Demobek",
	// 	Bio:    "Test uchun",
	// })
	// if err != nil {
	// 	return err
	// }
	// log.Println(updateResponse) // update author

	// delete author
	// err = queries.DeleteAuthor(ctx, 7)
	// if err != nil {
	// 	return err
	// }

	// prints true
	// log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))

	return nil
}
