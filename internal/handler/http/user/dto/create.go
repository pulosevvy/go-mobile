package dto

type CreateUserDto struct {
	PassportNumber string `json:"passportNumber" binding:"required,passport"`
}
