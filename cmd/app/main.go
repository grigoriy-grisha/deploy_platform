package main

import (
	initServerApp "github.com/kybasas/cdcd_platform"
	"github.com/kybasas/cdcd_platform/internal/controller"
	"log"
)

func main() {
	server := new(initServerApp.Server)

	if err := server.Run("5000", controller.InitHandlers()); err != nil {
		log.Fatal(err.Error())
	}

}
