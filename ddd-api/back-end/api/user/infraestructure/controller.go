package users

import (
	"encoding/json"
	application "github.com/etrapani/gotraining/back-end/api/user/application"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	createUser application.CreateUser
	getUsers   application.GetUsersService
	updateUser application.UpdateUserService
	deleteUser application.DeleteUserService
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

func NewController(createUser application.CreateUser,
	getUsers application.GetUsersService,
	updateUser application.UpdateUserService,
	deleteUserService application.DeleteUserService) Controller {
	return Controller{
		createUser: createUser,
		getUsers:   getUsers,
		updateUser: updateUser,
		deleteUser: deleteUserService,
	}
}

func (c Controller) create(w http.ResponseWriter, r *http.Request) {
	var newUser application.UserCreateRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(w, "The user information is not complete")
		return
	}

	json.Unmarshal(reqBody, &newUser)

	c.createUser.Execute(newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func (c Controller) getOne(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		log.Println(w, "The path varible id cannot convert to int")
		return
	}
	user, err := c.getUsers.GetOneUser(userID)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(toUserDto(user))
}

func (c Controller) getAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.getUsers.GetAllUsers()
	log.Println(w, "The path varible id cannot convert to int")
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(toListUserResponseDto(result))
}

func (c Controller) update(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(w, "The path varible id cannot convert to int")
		return
	}

	var userDto application.UserUpdateRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(w, "The user information is not complete")
		return
	}
	json.Unmarshal(reqBody, &userDto)

	c.updateUser.Execute(userID, userDto)
}

func (c Controller) delete(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(w, "The path varible id cannot convert to int")
		return
	}
	c.deleteUser.Execute(userID)
}
