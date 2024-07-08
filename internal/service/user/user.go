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

func (us *userService) Create(c context.Context, dto *dto.CreateUserDto) (*string, error) {
	series, number := decomposePassport(dto.PassportNumber)

	id, err := us.repo.Create(c, dto.PassportNumber, series, number)
	if err != nil {
		us.log.Error("UserService - Create", sl.Err(err))
		return nil, err
	}

	//The terms of reference to do not specify the logic after returning from an external API, so only logging
	_, err = us.GetPeopleInfo(series, number)
	if err != nil {
		us.log.Error("UserService - Create - GetPeopleInfo", sl.Err(err))
	}
	return id, nil
}

func (us *userService) GetAll(ctx context.Context, params *dto.GetAllParams) (*entity.UserListResponse, error) {
	users, err := us.repo.GetAll(ctx, params)
	if err != nil {
		us.log.Error("UserService - GetAll", sl.Err(err))
		return nil, err
	}

	return users, nil
}

func (us *userService) GetUserById(ctx context.Context, id string) (*entity.UserToResponse, error) {
	user, err := us.repo.FindUserByCustomField(ctx, "id", id)
	if err != nil {
		us.log.Error("UserService - GetUserById", sl.Err(err))
		return nil, err
	}

	return user, nil
}

func (us *userService) Delete(c context.Context, userId string) error {
	err := us.repo.Delete(c, userId)
	if err != nil {
		us.log.Error("UserService - Delete", sl.Err(err))
		return err
	}
	return nil
}

func (us *userService) Update(c context.Context, dto *dto.UpdateUserDto, userId string) error {
	series, number := decomposePassport(dto.Passport)
	err := us.repo.Update(c, dto, userId, series, number)
	if err != nil {
		us.log.Error("UserService - Update", sl.Err(err))
		return err
	}
	return nil
}

func (us *userService) GetUserByPassport(ctx context.Context, passport string) (*entity.UserToResponse, error) {
	user, err := us.repo.FindUserByCustomField(ctx, "passport", passport)
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
