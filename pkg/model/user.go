package model

import (
	"gorm.io/gorm"
)

// User domain object defining a user
// swagger:model
type User struct {
	gorm.Model
	Email       string  `gorm:"index;unique"`
	Password    string  `json:"-"`
	Groups      []Group `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AdminGroups []Group `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *User) IsMemberOf(groupName string) bool {
	return u.contains(groupName, u.Groups)
}

func (u *User) IsAdminOf(groupName string) bool {
	return u.contains(groupName, u.AdminGroups)
}

func (u *User) contains(groupName string, groups []Group) bool {
	for _, group := range groups {
		if groupName == group.Name {
			return true
		}
	}
	return false
}

func (u *User) IsAdministrator() bool {
	return u.IsMemberOf(AdministratorGroupName)
}
