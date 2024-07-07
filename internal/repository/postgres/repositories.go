package repository

import (
	"context"
	entity "go-mobile/internal/entitiy"
)

type UserRepository interface {
	GetAll(ctx context.Context)
	GetById(ctx context.Context)
	Create(ctx context.Context, passport, passportSeries, passportNumber string) error
	GetUserByPassport(ctx context.Context, passport string) (*entity.UserToResponse, error)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type TaskRepository interface {
	GetByUserId(ctx context.Context, userId string)
	StartTime(ctx context.Context)
	EndTime(ctx context.Context)
}
