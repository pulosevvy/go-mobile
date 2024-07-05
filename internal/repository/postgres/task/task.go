package repository

import (
	"go-mobile/package/database/postgres"
)

type taskRepository struct {
	db *postgres.Postgres
}

func NewTaskRepository(db *postgres.Postgres) *taskRepository {
	return &taskRepository{db}
}

func (t taskRepository) GetByUserId(userId string) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) StartTime() {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) EndTime() {
	//TODO implement me
	panic("implement me")
}
