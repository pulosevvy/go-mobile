package service

import (
	"context"
	entity "go-mobile/internal/entitiy"
	taskDto "go-mobile/internal/handler/http/task/dto"
	userDto "go-mobile/internal/handler/http/user/dto"
)

type UserService interface {
	GetAll(ctx context.Context, params *userDto.GetAllParams) (*entity.UserListResponse, error)
	Create(ctx context.Context, dto *userDto.CreateUserDto) (*string, error)
	Delete(ctx context.Context, userId string) error
	Update(c context.Context, dto *userDto.UpdateUserDto, userId string) error
	GetUserByPassport(ctx context.Context, passport string) (*entity.UserToResponse, error)
	GetUserById(ctx context.Context, id string) (*entity.UserToResponse, error)
	GetPeopleInfo(passportSeries string, passportNumber string) (*entity.PeopleApiResponse, error)
}

type TaskService interface {
	GetByUserId(ctx context.Context, userId string, dto *taskDto.GetByUser) ([]entity.TaskToResponse, error)
	GetTaskById(ctx context.Context, taskId string) (*entity.TaskToResponse, error)
	CreateTask(ctx context.Context, dto *taskDto.CreateTaskDto) (*string, error)
	StartTime(ctx context.Context, taskId string, dto *taskDto.StartTaskDto) error
	EndTime(ctx context.Context, task *entity.TaskToResponse, dto *taskDto.EndTaskDto) error
}
