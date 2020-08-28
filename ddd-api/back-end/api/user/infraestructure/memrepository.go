package users

import (
	"errors"
	"log"

	domain "github.com/etrapani/gotraining/back-end/api/user/domain"
)

type memUser struct {
	ID       int
	Username string
	Name     string
	Lastname string
	Password string
}

type allUsers []memUser

type MemRepository struct {
	users allUsers
	maxId int
}

func NewMemRepository() *MemRepository {
	return &MemRepository{
		users: allUsers{
			memUser{
				ID:       -1,
				Username: "admin@example.com",
				Name:     "admin",
				Lastname: "admin",
				Password: "admin",
			},
		},
		maxId: 0,
	}
}

func (mr *MemRepository) Save(newUser domain.User) error {
	log.Println("Input parameters -> ", newUser)
	var memUser = toUser(newUser)

	for i, singleUser := range mr.users {
		if singleUser.ID == memUser.ID {
			mr.users = append(mr.users[:i], memUser)
			return nil
		}
	}
	mr.users = append(mr.users, memUser)
	mr.maxId = mr.maxId + 1
	log.Println("users -> ", mr.users)
	log.Println("New user created ", newUser)
	return nil
}

func (mr MemRepository) Exist(id int) bool {
	for _, singleUser := range mr.users {
		if singleUser.ID == id {
			return true
		}
	}
	return false
}

func (mr MemRepository) FindOne(id int) (domain.User, error) {
	var result domain.User
	for _, singleUser := range mr.users {
		if singleUser.ID == id {
			result = toMemUser(singleUser)
			return result, nil
		}
	}
	log.Println("User does not exist", id)
	return result, errors.New("the user does not exist")
}

func (mr MemRepository) FindAll() []domain.User {
	var result []domain.User

	for _, singleUser := range mr.users {
		var resultUser = toMemUser(singleUser)
		result = append(result, resultUser)
	}
	log.Println("Output -> ", result)
	return result
}

func (mr *MemRepository) Delete(id int) {
	for i, singleUser := range mr.users {
		if singleUser.ID == id {
			mr.users = append(mr.users[:i], mr.users[i+1:]...)
			log.Println(id, "The user with ID %v has been deleted successfully")
		}
	}
}

func (mr *MemRepository) NextId() int {
	var result = mr.maxId
	log.Println("Output -> ", result)
	return result
}

func toUser(source domain.User) memUser {
	var result memUser
	result.ID = source.ID
	result.Username = source.Username
	result.Name = source.Name
	result.Lastname = source.Lastname
	result.Password = source.Password
	return result
}

func toMemUser(source memUser) domain.User {
	var result domain.User
	result.ID = source.ID
	result.Username = source.Username
	result.Name = source.Name
	result.Lastname = source.Lastname
	result.Password = source.Password
	return result
}
