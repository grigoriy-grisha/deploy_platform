package repositories

import (
	"cdcd_platform/internal/domain/entity"
	"context"
	"fmt"
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
	row := tx.QueryRow(ctx, getCreateSql(), project.Name, project.Command)

	if err := row.Scan(&id); err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	return id, tx.Commit(ctx)
}

func getCreateSql() string {
	return fmt.Sprintf("INSERT INTO %s (name, command) VALUES ($1, $2) RETURNING id", projectTable)
}
