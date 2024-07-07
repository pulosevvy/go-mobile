package repository

import (
	"context"
	"go-mobile/package/database/postgres"
)

type taskRepository struct {
	db *postgres.Postgres
}

func NewTaskRepository(db *postgres.Postgres) *taskRepository {
	return &taskRepository{db}
}

func (t taskRepository) GetByUserId(ctx context.Context, userId string) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) StartTime(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) EndTime(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
