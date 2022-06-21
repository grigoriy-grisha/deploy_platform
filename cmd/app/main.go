package main

import (
	repositories "cdcd_platform/internal/adapters/db/postrges"
	"cdcd_platform/internal/config"
	"cdcd_platform/internal/controller"
	"cdcd_platform/internal/domain/service"
	"cdcd_platform/pkg/client"
	"cdcd_platform/pkg/server"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	if err != nil {
		log.Fatal(err)
	}

	srv := initApp(pool)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initApp(pool *pgxpool.Pool) *server.HttpServer {
	conn, err := pool.Acquire(context.Background())
	conn.Conn()

	if err != nil {
		log.Fatalf("Unable to acquire a database connection: %v\n", err)
	}

	srv := new(server.HttpServer)

	repos := repositories.NewRepository(pool)
	services := service.NewService(repos)

	go func() {
		err := srv.Run(
			viper.GetString("port"),
			controller.InitHandlers(services),
		)

		if err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	return srv
}
