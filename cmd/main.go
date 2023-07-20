package main

import (
	"log"

	i "github.com/SQLC/internal"
)

func main() {
	if err := i.Run(); err != nil {
		log.Fatal(err)
	}
}
