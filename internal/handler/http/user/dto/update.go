package dto

type UpdateUserDto struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Address    *string `json:"address"`
	Passport   string  `json:"passport" binding:"required,passport"`
}
