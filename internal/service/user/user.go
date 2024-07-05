package service

import repository "go-mobile/internal/repository/postgres"

type userService struct {
	repo    repository.UserRepository
	userApi string
}

func NewUserService(repo repository.UserRepository, userApi string) *userService {
	return &userService{
		repo:    repo,
		userApi: userApi,
	}
}

func (us *userService) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (us *userService) GetById() {
	//TODO implement me
	panic("implement me")
}

func (us *userService) Create() {
	//TODO implement me
	panic("implement me")
}

func (us *userService) Delete() {
	//TODO implement me
	panic("implement me")
}

func (us *userService) Update() {
	//TODO implement me
	panic("implement me")
}
