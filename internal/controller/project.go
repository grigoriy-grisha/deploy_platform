package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type projectHandler struct {
}

func NewProjectHandler() *projectHandler {
	return &projectHandler{}
}

func (ph projectHandler) Register(router *gin.Engine) {
	router.GET("/hello", ph.Handler)
}

func (ph projectHandler) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": "1",
	})
}
