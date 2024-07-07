package task

import (
	"context"
	entity "go-mobile/internal/entitiy"
	taskDto "go-mobile/internal/handler/http/task/dto"
	repository "go-mobile/internal/repository/postgres"
	sl "go-mobile/package/logger/slog"
	"log/slog"
	"time"
)

type taskService struct {
	repo repository.TaskRepository
	log  *slog.Logger
}

func NewTaskService(repo repository.TaskRepository, log *slog.Logger) *taskService {
	return &taskService{repo, log}
}

func (t *taskService) GetByUserId(ctx context.Context, userId string) ([]entity.TaskToResponse, error) {
	tasks, err := t.repo.GetByUserId(ctx, userId)
	if err != nil {
		t.log.Error("TaskService - CreateTask", sl.Err(err))
		return nil, err
	}

	return tasks, nil
}

func (t *taskService) GetTaskById(ctx context.Context, taskId string) (*entity.TaskToResponse, error) {
	task, err := t.repo.FindTaskByCustomField(ctx, "id", taskId)
	if err != nil {
		t.log.Error("TaskService - CreateTask", sl.Err(err))
		return nil, err
	}

	return task, nil
}

func (t *taskService) CreateTask(ctx context.Context, dto *taskDto.CreateTaskDto) error {
	err := t.repo.CreateTask(ctx, dto)
	if err != nil {
		t.log.Error("TaskService - CreateTask", sl.Err(err))
		return err
	}

	return nil
}

func (t *taskService) StartTime(ctx context.Context, taskId string, dto *taskDto.StartTaskDto) error {
	err := t.repo.StartTime(ctx, taskId, dto)
	if err != nil {
		t.log.Error("TaskService - StartTime", sl.Err(err))
		return err
	}

	return nil
}

func (t *taskService) EndTime(ctx context.Context, task *entity.TaskToResponse, dto *taskDto.EndTaskDto) error {
	if dto.EndTime <= 0 {
		dto.EndTime = time.Now().Unix()
	}

	//duration := time.Unix(dto.EndTime, 0).Sub(task.StartTask)
	//hours := duration.Hours()

	err := t.repo.EndTime(ctx, task.Id, dto)
	if err != nil {
		t.log.Error("TaskService - EndTime", sl.Err(err))
		return err
	}

	return nil
}
