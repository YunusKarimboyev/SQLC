// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package generate

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        int32
	Name      string
	Bio       sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
