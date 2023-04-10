package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

type client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewStorage(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		viper.Get("PTOBuilder_storage.username"),
		viper.Get("PTOBuilder_storage.password"),
		viper.Get("PTOBuilder_storage.host"),
		viper.Get("PTOBuilder_storage.port"),
		viper.Get("PTOBuilder_storage.database")))
	return conn, err
}
