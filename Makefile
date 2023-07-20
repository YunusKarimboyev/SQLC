gen:
	sqlc generate

build:
	go build ./...

run:
	go run cmd/main.go