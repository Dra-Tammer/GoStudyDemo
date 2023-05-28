package dto

import "StudyDemo/LoginRegisterCheck/modules"

type UserDto struct {
	Name string `json:"name"`
}

func ToUserDto(user *modules.User) UserDto {
	return UserDto{
		Name: user.Name,
	}
}
