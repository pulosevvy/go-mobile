package dto

type GetByUser struct {
	StartDate string `form:"start_date" binding:"omitempty,dateformat=2006-01-02"`
	EndDate   string `form:"end_date" binding:"omitempty,dateformat=2006-01-02"`

	StartDateUnix int64
	EndDateUnix   int64
}
