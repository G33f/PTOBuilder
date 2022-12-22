package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

func NewStorage(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		viper.Get("PTOBuilder_storage.username"), viper.Get("PTOBuilder_storage.password"), viper.Get("PTOBuilder_storage.host"),
		viper.Get("PTOBuilder_storage.port"), viper.Get("PTOBuilder_storage.database")))
	return conn, err
}
