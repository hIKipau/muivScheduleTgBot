package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (repo *Repo) UpdateUserCourse(ctx context.Context, userID int64, course string) error {
	_, err := repo.db.Exec(ctx, "INSERT INTO users (id, course) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET course = EXCLUDED.course", userID, course)
	if err != nil {
		return fmt.Errorf("error UpdateUserCourse: %w", err)
	}
	return nil
}
func (repo *Repo) UpdateUserGroup(ctx context.Context, userID int64, group string) error {
	_, err := repo.db.Exec(ctx, "INSERT INTO users (id, groups) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET groups = EXCLUDED.groups", userID, group)
	if err != nil {
		return fmt.Errorf("error UpdateUserGroup: %w", err)
	}
	return nil
}
func (repo *Repo) GetUserInfo(ctx context.Context, userID int64) (string, string, error) {
	var course string
	var group string
	row := repo.db.QueryRow(ctx, "SELECT course, groups FROM users WHERE id = $1", userID)
	err := row.Scan(&course, &group)
	if err != nil {
		return "", "", fmt.Errorf("error when getting user info with row.Scan(): %w", err)
	}
	return course, group, nil
}
