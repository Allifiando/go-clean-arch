package repo

import (
	"context"
	"database/sql"
	"fmt"

	"go-clean-arch/src/domain"
)

type Repo struct {
	Conn *sql.DB
	*Queries
}

// InitRepo ...
func InitRepo(db *sql.DB) domain.Repo {
	return &Repo{
		Conn:    db,
		Queries: New(db),
	}
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func (store *Repo) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
