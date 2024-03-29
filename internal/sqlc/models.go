// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	Uuid uuid.UUID
	Name string
	Bio  sql.NullString
}

type AuthorBook struct {
	AuthorUuid uuid.UUID
	BookUuid   uuid.UUID
}

type Book struct {
	Uuid          uuid.UUID
	Title         string
	PublisherUuid uuid.UUID
}

type Publisher struct {
	Uuid uuid.UUID
	Name string
}
