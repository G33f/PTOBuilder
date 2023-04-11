package storage

import (
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewStorage(ctx context.Context, log *logging.Logger) (pool *pgxpool.Pool, err error) {
	// Get data from the configs and use it to create a connection string for the SQL DataBase
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		viper.Get("PTOBuilder_storage.username"),
		viper.Get("PTOBuilder_storage.password"),
		viper.Get("PTOBuilder_storage.host"),
		viper.Get("PTOBuilder_storage.port"),
		viper.Get("PTOBuilder_storage.database"))
	attempt := viper.GetInt("PTOBuilder_storage.maxAttempt")

	err = utils.DoWhitTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, connectionString)
		if err != nil {
			return err
		}
		err = pool.Ping(ctx)
		if err != nil {
			return err
		}

		return nil
	}, attempt, 5*time.Second)
	if err != nil {
		log.Fatal(err, "(error DoWithTries postgresql)")
	}
	return pool, nil
}
