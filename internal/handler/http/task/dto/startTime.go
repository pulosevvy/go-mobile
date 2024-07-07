package dto

type StartTaskDto struct {
	StartTime int64  `json:"start_time"`
	UserID    string `json:"user_id" binding:"required,isUuid"`
}
