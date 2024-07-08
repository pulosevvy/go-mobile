package task

import (
	"context"
	entity "go-mobile/internal/entitiy"
	taskDto "go-mobile/internal/handler/http/task/dto"
	repository "go-mobile/internal/repository/postgres"
	sl "go-mobile/package/logger/slog"
	"go-mobile/package/methods"
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

func (t *taskService) GetByUserId(ctx context.Context, userId string, dto *taskDto.GetByUser) ([]entity.TaskToResponse, error) {
	layout := "2006-01-02"
	if dto.StartDate != "" {
		startDate, err := time.Parse(layout, dto.StartDate)
		if err != nil {
			t.log.Error("TaskService - GetByUserId", sl.Err(err))
		}
		dto.StartDateUnix = startDate.Unix()
	}
	if dto.EndDate != "" {
		endDate, err := time.Parse(layout, dto.EndDate)
		if err != nil {
			t.log.Error("TaskService - GetByUserId", sl.Err(err))
		}
		dto.EndDateUnix = endDate.Unix()
	}

	tasks, err := t.repo.GetByUserId(ctx, userId, dto)

	if err != nil {
		t.log.Error("TaskService - GetByUserId", sl.Err(err))
		return nil, err
	}

	for i := range tasks {
		if tasks[i].Hours != nil {
			tasks[i].FormattedHours = methods.FloatToHours(*tasks[i].Hours)
		}
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

func (t *taskService) CreateTask(ctx context.Context, dto *taskDto.CreateTaskDto) (*string, error) {
	id, err := t.repo.CreateTask(ctx, dto)
	if err != nil {
		t.log.Error("TaskService - CreateTask", sl.Err(err))
		return nil, err
	}

	return id, nil
}

func (t *taskService) StartTime(ctx context.Context, taskId string, dto *taskDto.StartTaskDto) error {
	if dto.StartTime <= 0 {
		dto.StartTime = time.Now().Unix()
	}

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

	duration := time.Unix(dto.EndTime, 0).Sub(time.Unix(*task.StartTask, 0))
	hours := duration.Hours()

	err := t.repo.EndTime(ctx, task.Id, hours, dto)
	if err != nil {
		t.log.Error("TaskService - EndTime", sl.Err(err))
		return err
	}

	return nil
}
