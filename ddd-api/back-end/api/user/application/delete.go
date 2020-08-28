package users

import (
	domain "github.com/etrapani/gotraining/back-end/api/user/domain"
	"log"
)

type DeleteUser interface {
	Execute(updateUser UserUpdateRequest) (int, error)
}

type DeleteUserService struct {
	repo domain.Repository
}

func NewDeleteUser(repo domain.Repository) DeleteUserService {
	return DeleteUserService{
		repo: repo,
	}
}

func (s DeleteUserService) Execute(id int) {
	log.Println("Input parameters -> ", id)
	s.repo.Delete(id)
}
