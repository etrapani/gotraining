package users

import (
	"log"

	application "github.com/etrapani/gotraining/back-end/api/user/application"
	domain "github.com/etrapani/gotraining/back-end/api/user/domain"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	controller Controller
}

func NewUserRouter() UserRouter {
	var memRepository domain.Repository = NewMemRepository()
	var createUserService = application.NewCreateUserService(memRepository)
	var getUsersService = application.NewGetUsersService(memRepository)
	var updateUserService = application.NewUpdateUserService(memRepository)
	var deleteUserService = application.NewDeleteUser(memRepository)
	return UserRouter{
		controller: NewController(createUserService, getUsersService, updateUserService, deleteUserService),
	}
}

func (ur UserRouter) SetRoutes(s *mux.Router) {
	s.HandleFunc("/", ur.controller.create).Methods("POST")
	s.HandleFunc("/", ur.controller.getAll).Methods("GET")
	s.HandleFunc("/{id}", ur.controller.getOne).Methods("GET")
	s.HandleFunc("/{id}", ur.controller.update).Methods("PATCH")
	s.HandleFunc("/{id}", ur.controller.delete).Methods("DELETE")
	log.Println("add user api")
}
