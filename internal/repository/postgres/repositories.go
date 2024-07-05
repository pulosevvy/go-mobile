package repository

type UserRepository interface {
	GetAll()
	GetById()
	Create()
	Update()
	Delete()
}

type TaskRepository interface {
	GetByUserId(userId string)
	StartTime()
	EndTime()
}
