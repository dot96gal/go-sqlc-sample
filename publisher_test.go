package main

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dot96gal/go-sqlc-sample/internal/sqlc"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestCreatePublisher(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		expected sqlc.Publisher
	}{
		{
			scenario: "create publisher",
			input:    "hoge",
			expected: sqlc.Publisher{
				Name: "hoge",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.scenario, func(t *testing.T) {
			queries := sqlc.New(db)

			// test with transaction
			tx, err := db.Begin()
			if err != nil {
				t.Error(err)
			}
			t.Cleanup(func() {
				err = tx.Rollback()
				if err != nil {
					t.Error(err)
				}
			})

			queries = queries.WithTx(tx)

			// crete publisher
			ctx := context.Background()
			result, err := queries.CreatePublisher(ctx, tt.input)
			if err != nil {
				t.Error(err)
			}

			insertedPublisherID, err := result.LastInsertId()
			if err != nil {
				t.Error(err)
			}

			// get publisher
			publisher, err := queries.GetPublisher(ctx, insertedPublisherID)
			if err != nil {
				t.Error(err)
			}

			if publisher.Name != tt.expected.Name {
				t.Errorf("got=%v, want=%v", publisher.Name, tt.expected.Name)
			}
		})
	}
}

func TestUpdatePublisher(t *testing.T) {
	tests := []struct {
		scenario string
		input    sqlc.UpdatePublisherParams
		expected sqlc.Publisher
	}{
		{
			scenario: "update publisher",
			input: sqlc.UpdatePublisherParams{
				Name: "Updated: hoge",
			},
			expected: sqlc.Publisher{
				Name: "Updated: hoge",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.scenario, func(t *testing.T) {
			queries := sqlc.New(db)

			// test with transaction
			tx, err := db.Begin()
			if err != nil {
				t.Error(err)
			}
			t.Cleanup(func() {
				err = tx.Rollback()
				if err != nil {
					t.Error(err)
				}
			})

			queries = queries.WithTx(tx)

			// crete publisher
			ctx := context.Background()
			input := "hoge"
			result, err := queries.CreatePublisher(ctx, input)
			if err != nil {
				t.Error(err)
			}

			insertedPublisherID, err := result.LastInsertId()
			if err != nil {
				t.Error(err)
			}

			// update publisher
			tt.input.ID = insertedPublisherID
			err = queries.UpdatePublisher(ctx, tt.input)
			if err != nil {
				t.Error(err)
			}

			// get publisher
			publisher, err := queries.GetPublisher(ctx, insertedPublisherID)
			if err != nil {
				t.Error(err)
			}

			if publisher.Name != tt.expected.Name {
				t.Errorf("got=%v, want=%v", publisher.Name, tt.expected.Name)
			}
		})
	}
}

func TestDeletePublisher(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		expected error
	}{
		{
			scenario: "delete publisher",
			input:    "hoge",
			expected: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.scenario, func(t *testing.T) {
			queries := sqlc.New(db)

			// test with transaction
			tx, err := db.Begin()
			if err != nil {
				t.Error(err)
			}
			t.Cleanup(func() {
				err = tx.Rollback()
				if err != nil {
					t.Error(err)
				}
			})

			queries = queries.WithTx(tx)

			// crete publisher
			ctx := context.Background()
			result, err := queries.CreatePublisher(ctx, tt.input)
			if err != nil {
				t.Error(err)
			}

			insertedPublisherID, err := result.LastInsertId()
			if err != nil {
				t.Error(err)
			}

			// delete publisher
			err = queries.DeletePublisher(ctx, insertedPublisherID)
			if err != nil {
				t.Error(err)
			}

			// get publisher
			_, err = queries.GetPublisher(ctx, insertedPublisherID)
			if err != tt.expected {
				t.Errorf("got=%v, want=%v", err, tt.expected)
			}
		})
	}
}

func TestListPublisher(t *testing.T) {
	tests := []struct {
		scenario string
		input    []string
		expected []sqlc.Publisher
	}{
		{
			scenario: "list publisher",
			input: []string{
				"hoge",
				"fuga",
			},
			expected: []sqlc.Publisher{
				{
					Name: "hoge",
				},
				{
					Name: "fuga",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.scenario, func(t *testing.T) {
			queries := sqlc.New(db)

			// test with transaction
			tx, err := db.Begin()
			if err != nil {
				t.Error(err)
			}
			t.Cleanup(func() {
				err = tx.Rollback()
				if err != nil {
					t.Error(err)
				}
			})

			queries = queries.WithTx(tx)

			// crete publisher
			ctx := context.Background()
			for _, input := range tt.input {
				_, err := queries.CreatePublisher(ctx, input)
				if err != nil {
					t.Error(err)
				}
			}

			// list publisher
			results, err := queries.ListPublishers(ctx)
			if err != nil {
				t.Error(err)
			}

			for i := 0; i < len(tt.expected); i++ {
				if results[i].Name != tt.expected[i].Name {
					t.Errorf("got=%v, want=%v", results[i].Name, tt.expected[i].Name)
				}
			}
		})
	}
}
