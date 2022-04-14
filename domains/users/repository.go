package users

type UserRepostory interface{
	CreateUser(user User) error
	GetAllUser() []User
	GetUserById(id int) (User, error)
}