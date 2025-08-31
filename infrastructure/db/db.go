package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yuita-yoshihiko/daredemo-design-backend/config"
)

var ErrNotFound = errors.New("record not found")

type txKeyType struct{}

var txKey = txKeyType{}

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", config.Conf.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(300 * time.Second)
	return db, nil
}

type DBManager interface {
	DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error)
}

type DBManagerImpl struct {
	db *sqlx.DB
}

func NewDBManager(db *sqlx.DB) DBManager {
	return &DBManagerImpl{db: db}
}

func (m *DBManagerImpl) DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	ctx = context.WithValue(ctx, txKey, tx)
	v, err := f(ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return v, nil
}

type DBUtils interface {
	ConnectionFromCtx(ctx context.Context) Executor
	Error(error) error
}

type dbUtil struct {
	db *sqlx.DB
}

func NewDBUtil(db *sqlx.DB) DBUtils {
	return &dbUtil{db: db}
}

func (u *dbUtil) ConnectionFromCtx(ctx context.Context) Executor {
	tx, ok := ctx.Value(txKey).(*sqlx.Tx)
	if ok {
		return &TxExecutor{tx}
	}
	return u.db
}

func (u *dbUtil) Error(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: %v", ErrNotFound, err)
	}
	return err
}

type Executor interface {
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error)
}

type TxExecutor struct {
	*sqlx.Tx
}

func (e *TxExecutor) NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error) {
	return e.Tx.NamedExecContext(ctx, query, arg)
}

func (e *TxExecutor) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return e.Tx.ExecContext(ctx, query, args...)
}

func (e *TxExecutor) QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	return e.Tx.QueryxContext(ctx, query, args...)
}

func (e *TxExecutor) QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row {
	return e.Tx.QueryRowxContext(ctx, query, args...)
}

func (e *TxExecutor) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	return e.Tx.SelectContext(ctx, dest, query, args...)
}

func (e *TxExecutor) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	return e.Tx.GetContext(ctx, dest, query, args...)
}

func (e *TxExecutor) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return sqlx.NamedQueryContext(ctx, e.Tx, query, arg)
}
