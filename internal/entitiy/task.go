package entity

type TaskToResponse struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Hours     *float64 `json:"hours"`
	StartTask *int64   `json:"start_task"`
	EndTask   *int64   `json:"end_task"`
	UserID    string   `json:"user_id"`
}
