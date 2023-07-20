package main

import (
	"log"
	i "github.com/SQLC/internal/app"
)

func main() {
	if err := i.Run(); err != nil {
		log.Fatal(err)
	}
}
