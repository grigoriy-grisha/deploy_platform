package repositories

import (
	"cdcd_platform/internal/domain/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProjectRepo interface {
	Create(ctx context.Context, project entity.Project) (int, error)
	GetByID(ctx context.Context, id int) (*entity.Project, error)
	GetAll(ctx context.Context) ([]*entity.Project, error)
}

type Repository struct {
	ProjectRepo
}

func NewRepository(pgx *pgxpool.Pool) *Repository {
	return &Repository{
		ProjectRepo: NewProjectStorage(pgx),
	}
}
