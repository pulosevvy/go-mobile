package task

import (
	"context"
	repository "go-mobile/internal/repository/postgres"
)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{repo}
}

func (t *taskService) GetByUser(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *taskService) StartTime(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *taskService) EndTime(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
