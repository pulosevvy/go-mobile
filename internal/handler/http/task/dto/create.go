package dto

type CreateTaskDto struct {
	UserId string `json:"user_id" binding:"isUuid"`
	Name   string `json:"name" binding:"required"`
}
