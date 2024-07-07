package service

import (
	"context"
	entity "go-mobile/internal/entitiy"
	"go-mobile/internal/handler/http/user/dto"
	apiRepository "go-mobile/internal/repository/api"
	repository "go-mobile/internal/repository/postgres"
	sl "go-mobile/package/logger/slog"
	"log/slog"
	"strings"
)

type userService struct {
	repo    repository.UserRepository
	userApi apiRepository.UserApi
	log     *slog.Logger
}

func NewUserService(repo repository.UserRepository, userApi apiRepository.UserApi, log *slog.Logger) *userService {
	return &userService{repo, userApi, log}
}

func (us *userService) Create(c context.Context, dto *dto.CreateUserDto) error {
	series, number := decomposePassport(dto.PassportNumber)

	err := us.repo.Create(c, dto.PassportNumber, series, number)
	if err != nil {
		us.log.Error("UserService - Create", sl.Err(err))
		return err
	}

	//The terms of reference to do not specify the logic after returning from an external API, so only logging
	_, err = us.GetPeopleInfo(series, number)
	if err != nil {
		us.log.Error("UserService - Create - GetPeopleInfo", sl.Err(err))
		return nil
	}

	return nil
}

func (us *userService) GetAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (us *userService) GetById(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (us *userService) Delete(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (us *userService) Update(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (us *userService) GetUserByPassport(ctx context.Context, passport string) (*entity.UserToResponse, error) {
	user, err := us.repo.GetUserByPassport(ctx, passport)
	if err != nil {
		us.log.Error("UserService - GetUserByPassport", sl.Err(err))
		return nil, err
	}

	return user, nil
}

func (us *userService) GetPeopleInfo(passportSeries, passportNumber string) (*entity.PeopleApiResponse, error) {
	peopleInfo, err := us.userApi.GetPeopleInfo(passportSeries, passportNumber)
	if err != nil {
		us.log.Error("UserService - GetPeopleInfo", sl.Err(err))
		return nil, err
	}
	return peopleInfo, nil
}

func decomposePassport(passportNumber string) (string, string) {
	passport := strings.Split(passportNumber, " ")
	series := passport[0]
	number := passport[1]
	return series, number
}
