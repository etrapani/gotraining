package users

type User struct {
	ID       int
	Username string
	Name     string
	Lastname string
	Password string
}

func (u User) fullName() string {
	return u.Name + ", " + u.Lastname
}
