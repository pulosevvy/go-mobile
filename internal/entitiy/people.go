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

type PeopleApiResponse struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
