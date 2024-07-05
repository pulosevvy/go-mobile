package service

type UserService interface {
	GetAll()
	GetById()
	Create()
	Delete()
	Update()
}

type TaskService interface {
	GetByUser()
	StartTime()
	EndTime()
}
