package repository

import (
	"go-mobile/package/database/postgres"
)

type userRepository struct {
	db *postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *userRepository {
	return &userRepository{db}
}

func (u userRepository) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetById() {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update() {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
