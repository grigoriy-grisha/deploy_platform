package controller

import "github.com/gin-gonic/gin"

type HandlerRegistration interface {
	Register(c *gin.Engine)
}

func InitHandlers() *gin.Engine {
	router := gin.New()

	registerHandlers(router, []HandlerRegistration{
		NewProjectHandler(),
	})

	return router
}

func registerHandlers(router *gin.Engine, handlerRegistrations []HandlerRegistration) {
	for _, registration := range handlerRegistrations {
		registration.Register(router)
	}
}
