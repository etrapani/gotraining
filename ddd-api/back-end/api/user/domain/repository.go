package users

// Repository interface
type Repository interface {
	Save(user User) error

	Exist(id int) bool

	FindOne(id int) (User, error)

	FindAll() []User

	Delete(id int)

	NextId() int
}
