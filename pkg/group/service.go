package group

import (
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/user"
)

type service struct {
	groupRepository Repository
	userRepository  user.Repository
}

func NewService(groupRepository Repository, userRepository user.Repository) *service {
	return &service{
		groupRepository,
		userRepository,
	}
}

type Repository interface {
	create(group *model.Group) error
	addUser(group *model.Group, user *model.User) error
	addClusterConfiguration(configuration *model.ClusterConfiguration) error
	getClusterConfiguration(groupName string) (*model.ClusterConfiguration, error)
	find(name string) (*model.Group, error)
	findOrCreate(group *model.Group) (*model.Group, error)
}

func (s *service) Find(name string) (*model.Group, error) {
	return s.groupRepository.find(name)
}

func (s *service) Create(name string, hostname string) (*model.Group, error) {
	group := &model.Group{
		Name:     name,
		Hostname: hostname,
	}

	err := s.groupRepository.create(group)
	if err != nil {
		return nil, err
	}

	return group, err
}

func (s *service) FindOrCreate(name string, hostname string) (*model.Group, error) {
	group := &model.Group{
		Name:     name,
		Hostname: hostname,
	}

	g, err := s.groupRepository.findOrCreate(group)
	if err != nil {
		return nil, err
	}

	return g, err
}

func (s *service) AddUser(groupName string, userId uint) error {
	group, err := s.Find(groupName)
	if err != nil {
		return err
	}

	u, err := s.userRepository.FindById(userId)
	if err != nil {
		return err
	}

	return s.groupRepository.addUser(group, u)
}

func (s *service) AddClusterConfiguration(clusterConfiguration *model.ClusterConfiguration) error {
	return s.groupRepository.addClusterConfiguration(clusterConfiguration)
}

func (s *service) GetClusterConfiguration(groupName string) (*model.ClusterConfiguration, error) {
	return s.groupRepository.getClusterConfiguration(groupName)
}
