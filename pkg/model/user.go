package model

import (
	"sort"

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
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name <= groups[j].Name
	})

	index := sort.Search(len(groups), func(i int) bool {
		return groups[i].Name >= groupName
	})

	return index < len(groups) && groups[index].Name == groupName
}

func (u *User) IsAdministrator() bool {
	return u.IsMemberOf(AdministratorGroupName)
}
