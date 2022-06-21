package service

import (
	repositories "cdcd_platform/internal/adapters/db/postrges"
	"cdcd_platform/internal/domain/entity"
	"context"
)

type Project interface {
	CreateProject(ctx context.Context, project entity.Project) (int, error)
}

type Service struct {
	Project
}

func NewService(repos *repositories.Repository) Service {
	return Service{
		Project: NewProjectService(repos.ProjectRepo),
	}
}
