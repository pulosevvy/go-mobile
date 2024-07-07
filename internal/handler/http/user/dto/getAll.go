package dto

type GetAllParams struct {
	Limit          int    `form:"limit" binding:"omitempty"`
	Page           int    `form:"page" binding:"omitempty"`
	OrderBy        string `form:"order_by" binding:"omitempty,oneof=name surname patronymic address passport passport_series passport_number created_at"`
	OrderSort      string `form:"order_sort" binding:"omitempty,oneof=asc desc"`
	Name           string `form:"name" binding:"omitempty"`
	Surname        string `form:"surname" binding:"omitempty"`
	Patronymic     string `form:"patronymic" binding:"omitempty"`
	Address        string `form:"address" binding:"omitempty"`
	Passport       string `form:"passport" binding:"omitempty"`
	PassportSeries string `form:"passport_series" binding:"omitempty"`
	PassportNumber string `form:"passport_number" binding:"omitempty"`
}
