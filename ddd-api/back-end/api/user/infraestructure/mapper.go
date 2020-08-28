package users

import (
	"github.com/etrapani/gotraining/back-end/api/user/domain"
)

type userUpdateDto struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type userResponseDto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func toUserDto(source users.User) userResponseDto {
	var result userResponseDto
	result.ID = source.ID
	result.Username = source.Username
	result.Name = source.Name
	result.Lastname = source.Lastname
	return result
}

func toListUserResponseDto(users []users.User) []userResponseDto {
	var result []userResponseDto
	for _, user := range users {
		result = append(result, toUserDto(user))
	}
	return result
}
