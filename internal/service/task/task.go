package task

import repository "go-mobile/internal/repository/postgres"

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{repo}
}

func (t *taskService) GetByUser() {
	//TODO implement me
	panic("implement me")
}

func (t *taskService) StartTime() {
	//TODO implement me
	panic("implement me")
}

func (t *taskService) EndTime() {
	//TODO implement me
	panic("implement me")
}
