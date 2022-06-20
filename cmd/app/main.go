package main

import (
	"cdcd_platform/internal/config"
	"cdcd_platform/pkg/client"
	"context"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err.Error())
	}

	pool, err := client.NewClient(context.Background(), client.StorageConfig{
		Dbname:   viper.GetString("db.dbname"),
		Host:     viper.GetString("db.host"),
		Password: viper.GetString("db.password"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
		Username: viper.GetString("db.username"),
	})

	conn, err := pool.Acquire(context.Background())

	if err != nil {
		log.Fatalf("Unable to acquire a database connection: %v\n", err)
	}

	conn.QueryRow(context.Background(), "INSERT INTO project (name, command) VALUES ($1, $2) RETURNING id", "hello", "hello")
	conn.Conn()
}
