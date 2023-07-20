package repo

import (
	"time"
)

type Author struct {
	ID        int
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthorParams struct {
	ID        int
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
