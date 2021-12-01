package dto

import "github.com/dhis2-sre/im-users/pgk/model"

type User struct {
	Id          uint    `json:"id"`
	Email       string  `json:"email"`
	Groups      []Group `json:"groups"`
	AdminGroups []Group `json:"adminGroups"`
}

func ToUser(user *model.User) User {
	dtoGroups := toGroups(user.Groups)
	dtoAdminGroups := toGroups(user.AdminGroups)
	return User{
		Id:          user.ID,
		Email:       user.Email,
		Groups:      dtoGroups,
		AdminGroups: dtoAdminGroups,
	}
}

func toUsers(users []model.User) []User {
	dtoUsers := make([]User, len(users))
	for i, user := range users {
		dtoUsers[i] = ToUser(&user)
	}
	return dtoUsers
}
