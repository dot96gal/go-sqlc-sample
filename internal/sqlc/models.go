// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type Publisher struct {
	ID   int64
	Name string
}
