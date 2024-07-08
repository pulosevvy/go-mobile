package repository

import (
	"context"
	entity "go-mobile/internal/entitiy"
	taskDto "go-mobile/internal/handler/http/task/dto"
	userDto "go-mobile/internal/handler/http/user/dto"
)

type UserRepository interface {
	GetAll(ctx context.Context, params *userDto.GetAllParams) (*entity.UserListResponse, error)
	Create(ctx context.Context, passport, passportSeries, passportNumber string) (*string, error)
	FindUserByCustomField(ctx context.Context, field, value string) (*entity.UserToResponse, error)
	Update(c context.Context, dto *userDto.UpdateUserDto, userId, series, number string) error
	Delete(ctx context.Context, userId string) error
}

type TaskRepository interface {
	GetByUserId(ctx context.Context, userId string, dto *taskDto.GetByUser) ([]entity.TaskToResponse, error)
	FindTaskByCustomField(ctx context.Context, field, value string) (*entity.TaskToResponse, error)
	CreateTask(ctx context.Context, dto *taskDto.CreateTaskDto) (*string, error)
	StartTime(ctx context.Context, taskId string, dto *taskDto.StartTaskDto) error
	EndTime(ctx context.Context, taskId string, hours float64, dto *taskDto.EndTaskDto) error
}
