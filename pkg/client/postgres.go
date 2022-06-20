package client

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type StorageConfig struct {
	Username string
	Host     string
	Port     string
	Dbname   string
	SSLMode  string
	Password string
}

func NewClient(ctx context.Context, storageConfig StorageConfig) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, getConnectionString(storageConfig))
}

func getConnectionString(sc StorageConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Dbname)
}
