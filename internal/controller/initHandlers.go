package controller

import (
	"cdcd_platform/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type HandlerRegistration interface {
	Register(c *gin.Engine)
}

func InitHandlers(services service.Service) *gin.Engine {
	router := gin.New()

	registerHandlers(router, []HandlerRegistration{
		NewProjectHandler(services.Project),
	})

	return router
}

func registerHandlers(router *gin.Engine, handlerRegistrations []HandlerRegistration) {
	for _, registration := range handlerRegistrations {
		registration.Register(router)
	}
}
