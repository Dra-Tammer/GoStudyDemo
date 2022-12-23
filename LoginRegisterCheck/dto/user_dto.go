package dto

import "StudyDemo/LoginRegisterCheck/modules"

// 想给前端返回的可能没有这么多的东西
// 这时候就需要这个函数来处理一下
type UserDto struct {
	Name string `json:"name"`
}

func ToUserDto(user modules.User) UserDto {
	return UserDto{
		Name: user.Name,
	}
}
