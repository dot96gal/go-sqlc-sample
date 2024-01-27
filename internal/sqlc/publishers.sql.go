// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: publishers.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createPublisher = `-- name: CreatePublisher :exec
INSERT INTO
  publishers (uuid, name)
VALUES
  (?, ?)
`

type CreatePublisherParams struct {
	Uuid uuid.UUID
	Name string
}

func (q *Queries) CreatePublisher(ctx context.Context, arg CreatePublisherParams) error {
	_, err := q.db.ExecContext(ctx, createPublisher, arg.Uuid, arg.Name)
	return err
}

const deletePublisher = `-- name: DeletePublisher :exec
DELETE FROM publishers
WHERE
  uuid = ?
`

func (q *Queries) DeletePublisher(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePublisher, argUuid)
	return err
}

const getPublisher = `-- name: GetPublisher :one
SELECT
  uuid, name
FROM
  publishers
WHERE
  uuid = ?
LIMIT
  1
`

func (q *Queries) GetPublisher(ctx context.Context, argUuid uuid.UUID) (Publisher, error) {
	row := q.db.QueryRowContext(ctx, getPublisher, argUuid)
	var i Publisher
	err := row.Scan(&i.Uuid, &i.Name)
	return i, err
}

const getPublisherBooks = `-- name: GetPublisherBooks :many
SELECT
  p.uuid AS publisher_uuid,
  p.name AS publisher_name,
  b.uuid AS book_uuid,
  b.title AS book_title
FROM
  publishers AS p
  INNER JOIN books AS b ON p.uuid = b.publisher_uuid
WHERE
  p.uuid = ?
ORDER BY
  p.uuid,
  b.uuid
`

type GetPublisherBooksRow struct {
	PublisherUuid uuid.UUID
	PublisherName string
	BookUuid      uuid.UUID
	BookTitle     string
}

func (q *Queries) GetPublisherBooks(ctx context.Context, argUuid uuid.UUID) ([]GetPublisherBooksRow, error) {
	rows, err := q.db.QueryContext(ctx, getPublisherBooks, argUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPublisherBooksRow
	for rows.Next() {
		var i GetPublisherBooksRow
		if err := rows.Scan(
			&i.PublisherUuid,
			&i.PublisherName,
			&i.BookUuid,
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

const listPublishers = `-- name: ListPublishers :many
SELECT
  uuid, name
FROM
  publishers
ORDER BY
  uuid
`

func (q *Queries) ListPublishers(ctx context.Context) ([]Publisher, error) {
	rows, err := q.db.QueryContext(ctx, listPublishers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Publisher
	for rows.Next() {
		var i Publisher
		if err := rows.Scan(&i.Uuid, &i.Name); err != nil {
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

const updatePublisher = `-- name: UpdatePublisher :exec
UPDATE publishers
SET
  name = ?
WHERE
  uuid = ?
`

type UpdatePublisherParams struct {
	Name string
	Uuid uuid.UUID
}

func (q *Queries) UpdatePublisher(ctx context.Context, arg UpdatePublisherParams) error {
	_, err := q.db.ExecContext(ctx, updatePublisher, arg.Name, arg.Uuid)
	return err
}