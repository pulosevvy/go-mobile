package service

import (
	"context"
	entity "go-mobile/internal/entitiy"
	"go-mobile/internal/handler/http/user/dto"
)

type UserService interface {
	GetAll(ctx context.Context)
	GetById(ctx context.Context)
	Create(ctx context.Context, dto *dto.CreateUserDto) error
	Delete(ctx context.Context)
	Update(ctx context.Context)
	GetUserByPassport(ctx context.Context, passport string) (*entity.UserToResponse, error)
	GetPeopleInfo(passportSeries string, passportNumber string) (*entity.PeopleApiResponse, error)
}

type TaskService interface {
	GetByUser(ctx context.Context)
	StartTime(ctx context.Context)
	EndTime(ctx context.Context)
}
