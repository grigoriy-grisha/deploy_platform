package service

import (
	repositories "cdcd_platform/internal/adapters/db/postrges"
	"cdcd_platform/internal/domain/entity"
	"context"
)

type ProjectService struct {
	projectRepository repositories.ProjectRepo
}

func NewProjectService(repo repositories.ProjectRepo) *ProjectService {
	return &ProjectService{projectRepository: repo}
}

func (ps ProjectService) CreateProject(ctx context.Context, project entity.Project) (int, error) {
	return ps.projectRepository.Create(ctx, project)
}
