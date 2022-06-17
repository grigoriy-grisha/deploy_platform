package main

import (
	initServerApp "github.com/kybasas/cdcd_platform"
	"github.com/kybasas/cdcd_platform/internal/controller"
	"github.com/spf13/viper"
	"log"
)

func main() {
	server := new(initServerApp.Server)

	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if err := server.Run(viper.GetString("port"), controller.InitHandlers()); err != nil {
		log.Fatal(err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
