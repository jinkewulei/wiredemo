package impl

type UserImpl struct {
}

func NewUserImpl() *UserImpl {
	return &UserImpl{}
}

func (u *UserImpl) GetByID(id int32) User {
	return User{}
}
