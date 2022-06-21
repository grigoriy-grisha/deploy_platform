package repositories

import (
	"cdcd_platform/internal/domain/entity"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProjectStorage struct {
	pgx *pgxpool.Pool
}

func NewProjectStorage(pgx *pgxpool.Pool) *ProjectStorage {
	return &ProjectStorage{pgx}
}

func (ps *ProjectStorage) Create(ctx context.Context, project entity.Project) (int, error) {
	tx, err := ps.pgx.Begin(ctx)

	if err != nil {
		return 0, err
	}

	var id int
	row := tx.QueryRow(ctx, SqlCreate(), project.Name, project.Command)

	if err := row.Scan(&id); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, err
		}

		return 0, err
	}

	return id, tx.Commit(ctx)
}

func SqlCreate() string {
	return fmt.Sprintf("INSERT INTO %s (name, command) VALUES ($1, $2) RETURNING id", projectTable)
}

func (ps *ProjectStorage) GetByID(ctx context.Context, id int) (*entity.Project, error) {
	var project []*entity.Project
	err := pgxscan.Select(ctx, ps.pgx, &project, SqlGetById(), id)

	if err != nil {
		return nil, err
	}

	return project[0], nil
}

func SqlGetById() string {
	return fmt.Sprintf("SELECT id, name, command FROM %s WHERE id=$1", projectTable)
}
