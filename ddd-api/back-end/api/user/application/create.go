package users

import (
	domain "github.com/etrapani/gotraining/back-end/api/user/domain"
	"log"
)

type CreateUser interface {
	Execute(newUser UserCreateRequest) (int, error)
}

type CreateUserService struct {
	repo domain.Repository
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

func NewCreateUserService(repo domain.Repository) CreateUserService {
	return CreateUserService{
		repo: repo,
	}
}

func (s CreateUserService) Execute(userCommand UserCreateRequest) (int, error) {
	log.Println("Input parameters -> ", userCommand)
	var nextId = s.repo.NextId()
	var newUser = toUser(nextId, userCommand)
	return newUser.ID, s.repo.Save(newUser)
}

func toUser(nextId int, source UserCreateRequest) domain.User {
	var result domain.User
	result.ID = nextId
	result.Username = source.Username
	result.Name = source.Name
	result.Lastname = source.Lastname
	result.Password = source.Password
	return result
}
