package apiRepository

import entity "go-mobile/internal/entitiy"

type UserApi interface {
	GetPeopleInfo(passportSeries, passportNumber string) (*entity.PeopleApiResponse, error)
}
