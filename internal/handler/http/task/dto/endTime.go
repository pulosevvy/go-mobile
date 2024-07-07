package dto

type EndTaskDto struct {
	EndTime int64  `json:"end_time" binding:"required"`
	UserID  string `json:"user_id" binding:"required,isUuid"`
}
