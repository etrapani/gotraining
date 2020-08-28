package users

import (
	domain "github.com/etrapani/gotraining/back-end/api/user/domain"
)

type GetUsers interface {
	Execute(newUser UserCreateRequest) (int, error)
}

type GetUsersService struct {
	repo domain.Repository
}

func NewGetUsersService(repo domain.Repository) GetUsersService {
	return GetUsersService{
		repo: repo,
	}
}

func (s GetUsersService) GetAllUsers() ([]domain.User, error) {
	var result = s.repo.FindAll()
	return result, nil
}

func (s GetUsersService) GetOneUser(id int) (domain.User, error) {
	var result, error = s.repo.FindOne(id)
	return result, error
}
