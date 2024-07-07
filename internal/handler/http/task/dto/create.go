package dto

type CreateTaskDto struct {
	UserId string `json:"user_id" binding:"required,isUuid"`
	Name   string `json:"name" binding:"required"`
}
