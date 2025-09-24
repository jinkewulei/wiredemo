package impl

type UserService interface {
	GetByID(id int32) User
}
