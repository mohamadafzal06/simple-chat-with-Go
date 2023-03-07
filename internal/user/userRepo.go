package user

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	//ExecContext(ctx context.Context, query string) (*sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	//PrepareContext(context.Context, string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	//	QueryContext(context.Context, string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	//	QueryRowContext(context.Context, string, args ...any) (*sql.Row, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertedID int
	query := "INSERT INOT users(username, password, email) VALUES($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(lastInsertedID)
	if err != nil {
		return &User{}, fmt.Errorf("cannot insert user to db: %w\n", err)
	}

	user.ID = int64(lastInsertedID)
	return user, nil
}
