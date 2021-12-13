package model

import (
	"gorm.io/gorm"
)

const AdministratorGroupName = "administrators"

// Group domain object defining a group
// swagger:model
type Group struct {
	gorm.Model
	Name                 string               `gorm:"unique;"`
	Hostname             string               `gorm:"unique;"`
	Users                []User               `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClusterConfiguration ClusterConfiguration `json:"clusterConfiguration"`
}

type ClusterConfiguration struct {
	gorm.Model
	GroupID                 uint
	KubernetesConfiguration []byte
}
