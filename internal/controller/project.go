package controller

import (
	"cdcd_platform/internal/domain/entity"
	"cdcd_platform/internal/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type projectHandler struct {
	projectService service.Project
}

func NewProjectHandler(projectService service.Project) *projectHandler {
	return &projectHandler{projectService}
}

func (ph projectHandler) Register(router *gin.Engine) {
	group := router.Group("project")
	{
		group.POST("/", ph.Create)
		group.GET("/", ph.GetAllProjects)
		group.GET("/:id", ph.GetProject)
	}

}

func (ph *projectHandler) Create(c *gin.Context) {
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

func (ph *projectHandler) GetProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	project, err := ph.projectService.GetProject(c, id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, project)
}

func (ph *projectHandler) GetAllProjects(c *gin.Context) {
	projects, err := ph.projectService.GetAllProjects(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, projects)
}
