package controller

import (
	"cdcd_platform/internal/domain/entity"
	"cdcd_platform/internal/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type projectHandler struct {
	projectService service.Project
}

func NewProjectHandler(projectService service.Project) *projectHandler {
	return &projectHandler{projectService}
}

func (ph projectHandler) Register(router *gin.Engine) {
	router.POST("/project/", ph.Create)
}

func (ph projectHandler) Create(c *gin.Context) {
	var input entity.Project
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ph.projectService.CreateProject(c, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}
