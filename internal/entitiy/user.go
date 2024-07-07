package entity

type User struct {
	Id             string `json:"id"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	Passport       string `json:"passport"`
	PassportSeries string `json:"passportSeries"`
	PassportNumber string `json:"passportNumber"`
}

type UserToResponse struct {
	Id         *string `json:"id"`
	Surname    *string `json:"surname"`
	Name       *string `json:"name"`
	Patronymic *string `json:"patronymic"`
	Address    *string `json:"address"`
	Passport   *string `json:"passport"`
}

type UserListResponse struct {
	Page       *int              `json:"page"`
	Limit      *int              `json:"limit"`
	TotalCount *int              `json:"totalCount"`
	TotalPage  int               `json:"totalPage"`
	Users      []*UserToResponse `json:"users"`
}

type PeopleApiResponse struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
