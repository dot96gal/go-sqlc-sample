// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: author_books.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createAuthorBook = `-- name: CreateAuthorBook :exec
INSERT INTO author_books (
  author_id, book_id 
) VALUES (
  ?, ?
)
`

type CreateAuthorBookParams struct {
	AuthorID int64
	BookID   int64
}

func (q *Queries) CreateAuthorBook(ctx context.Context, arg CreateAuthorBookParams) error {
	_, err := q.db.ExecContext(ctx, createAuthorBook, arg.AuthorID, arg.BookID)
	return err
}

const deleteAuthorBook = `-- name: DeleteAuthorBook :exec
DELETE FROM author_books
WHERE author_id = ? AND book_id = ?
`

type DeleteAuthorBookParams struct {
	AuthorID int64
	BookID   int64
}

func (q *Queries) DeleteAuthorBook(ctx context.Context, arg DeleteAuthorBookParams) error {
	_, err := q.db.ExecContext(ctx, deleteAuthorBook, arg.AuthorID, arg.BookID)
	return err
}

const getAuthorBook = `-- name: GetAuthorBook :one
SELECT author_id, book_id FROM author_books
WHERE author_id = ? AND book_id = ? LIMIT 1
`

type GetAuthorBookParams struct {
	AuthorID int64
	BookID   int64
}

func (q *Queries) GetAuthorBook(ctx context.Context, arg GetAuthorBookParams) (AuthorBook, error) {
	row := q.db.QueryRowContext(ctx, getAuthorBook, arg.AuthorID, arg.BookID)
	var i AuthorBook
	err := row.Scan(&i.AuthorID, &i.BookID)
	return i, err
}

const listAuthorBooks = `-- name: ListAuthorBooks :many
SELECT
  a.id AS author_id,
  a.name AS author_name,
  a.bio AS author_bio,
  b.id AS book_id,
  b.title AS book_title
FROM authors AS a
INNER JOIN author_books AS ab
ON a.id = ab.author_id
INNER JOIN books AS b
ON ab.book_id = b.id
ORDER BY a.id, b.id
`

type ListAuthorBooksRow struct {
	AuthorID   int64
	AuthorName string
	AuthorBio  sql.NullString
	BookID     int64
	BookTitle  string
}

func (q *Queries) ListAuthorBooks(ctx context.Context) ([]ListAuthorBooksRow, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAuthorBooksRow
	for rows.Next() {
		var i ListAuthorBooksRow
		if err := rows.Scan(
			&i.AuthorID,
			&i.AuthorName,
			&i.AuthorBio,
			&i.BookID,
			&i.BookTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
