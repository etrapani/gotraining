package users

import (
	domain "github.com/etrapani/gotraining/back-end/api/user/domain"
	"log"
)

type UpdateUser interface {
	Execute(updateUser UserUpdateRequest) (int, error)
}

type UpdateUserService struct {
	repo domain.Repository
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func NewUpdateUserService(repo domain.Repository) UpdateUserService {
	return UpdateUserService{
		repo: repo,
	}
}

func (s UpdateUserService) Execute(id int, userCommand UserUpdateRequest) error {
	log.Println("Input parameters -> ", userCommand)
	user, error := s.repo.FindOne(id)
	if error != nil {
		return error
	}
	var updatedUser = updateUser(user, userCommand)
	return s.repo.Save(updatedUser)
}

func updateUser(target domain.User, source UserUpdateRequest) domain.User {
	target.Name = source.Name
	target.Lastname = source.Lastname
	log.Println("Output -> ", target)
	return target
}
