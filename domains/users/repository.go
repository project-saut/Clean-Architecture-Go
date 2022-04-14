package users

type UserRepository interface {
	CreateUser(user User) error
	GetAllUser() []User
	GetUserById(id int) (User, error)
}
