package group

import (
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/user"
)

type Service interface {
	Create(name string, hostname string) (*model.Group, error)
	AddUser(groupName string, userId uint) error
	AddClusterConfiguration(clusterConfiguration *model.ClusterConfiguration) error
	GetClusterConfiguration(groupName string) (*model.ClusterConfiguration, error)
	Find(name string) (*model.Group, error)
	FindOrCreate(name string, hostname string) (*model.Group, error)
}

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

func (s *service) Find(name string) (*model.Group, error) {
	return s.groupRepository.Find(name)
}

func (s *service) Create(name string, hostname string) (*model.Group, error) {
	group := &model.Group{
		Name:     name,
		Hostname: hostname,
	}

	err := s.groupRepository.Create(group)
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

	g, err := s.groupRepository.FindOrCreate(group)
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

	return s.groupRepository.AddUser(group, u)
}

func (s *service) AddClusterConfiguration(clusterConfiguration *model.ClusterConfiguration) error {
	return s.groupRepository.AddClusterConfiguration(clusterConfiguration)
}

func (s *service) GetClusterConfiguration(groupName string) (*model.ClusterConfiguration, error) {
	return s.groupRepository.GetClusterConfiguration(groupName)
}
