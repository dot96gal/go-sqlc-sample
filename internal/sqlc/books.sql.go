// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: books.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createBook = `-- name: CreateBook :exec
INSERT INTO
  books (uuid, title, publisher_uuid)
VALUES
  (?, ?, ?)
`

type CreateBookParams struct {
	Uuid          uuid.UUID
	Title         string
	PublisherUuid uuid.UUID
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) error {
	_, err := q.db.ExecContext(ctx, createBook, arg.Uuid, arg.Title, arg.PublisherUuid)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE
  uuid = ?
`

func (q *Queries) DeleteBook(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBook, argUuid)
	return err
}

const getBook = `-- name: GetBook :one
SELECT
  uuid, title, publisher_uuid
FROM
  books
WHERE
  uuid = ?
LIMIT
  1
`

func (q *Queries) GetBook(ctx context.Context, argUuid uuid.UUID) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, argUuid)
	var i Book
	err := row.Scan(&i.Uuid, &i.Title, &i.PublisherUuid)
	return i, err
}

const getBookPublisher = `-- name: GetBookPublisher :one
SELECT
  b.uuid AS book_uuid,
  b.title AS book_title,
  p.uuid AS publisher_uuid,
  p.name AS publisher_name
FROM
  books AS b
  INNER JOIN publishers AS p ON b.publisher_uuid = p.uuid
WHERE
  b.uuid = ?
LIMIT
  1
`

type GetBookPublisherRow struct {
	BookUuid      uuid.UUID
	BookTitle     string
	PublisherUuid uuid.UUID
	PublisherName string
}

func (q *Queries) GetBookPublisher(ctx context.Context, argUuid uuid.UUID) (GetBookPublisherRow, error) {
	row := q.db.QueryRowContext(ctx, getBookPublisher, argUuid)
	var i GetBookPublisherRow
	err := row.Scan(
		&i.BookUuid,
		&i.BookTitle,
		&i.PublisherUuid,
		&i.PublisherName,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT
  uuid, title, publisher_uuid
FROM
  books
ORDER BY
  uuid
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(&i.Uuid, &i.Title, &i.PublisherUuid); err != nil {
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

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET
  title = ?
WHERE
  uuid = ?
`

type UpdateBookParams struct {
	Title string
	Uuid  uuid.UUID
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook, arg.Title, arg.Uuid)
	return err
}
