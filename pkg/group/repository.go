package group

import (
	"github.com/dhis2-sre/im-user/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(group *model.Group) error
	AddUser(group *model.Group, user *model.User) error
	AddClusterConfiguration(configuration *model.ClusterConfiguration) error
	GetClusterConfiguration(groupName string) (*model.ClusterConfiguration, error)
	Find(name string) (*model.Group, error)
	FindOrCreate(group *model.Group) (*model.Group, error)
}

func ProvideRepository(DB *gorm.DB) Repository {
	return &repository{db: DB}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Find(name string) (*model.Group, error) {
	var group *model.Group
	err := r.db.
		//		Preload("ClusterConfiguration").
		Where("name = ?", name).
		First(&group).Error
	return group, err
}

func (r repository) Create(group *model.Group) error {
	return r.db.Create(&group).Error
}

func (r repository) FindOrCreate(group *model.Group) (*model.Group, error) {
	var g *model.Group
	err := r.db.Where(model.Group{Name: group.Name}).Attrs(model.Group{Hostname: group.Hostname}).FirstOrCreate(&g).Error
	return g, err
}

func (r repository) AddUser(group *model.Group, user *model.User) error {
	return r.db.Model(&group).Association("Users").Append([]*model.User{user})
}

func (r repository) AddClusterConfiguration(configuration *model.ClusterConfiguration) error {
	return r.db.Create(&configuration).Error
}

func (r repository) GetClusterConfiguration(groupName string) (*model.ClusterConfiguration, error) {
	var configuration *model.ClusterConfiguration
	err := r.db.
		Where("group_name = ?", groupName).
		First(&configuration).Error
	return configuration, err
}
