package group

import (
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/pkg/model"
	"github.com/dhis2-sre/im-users/pkg/user"
	"strconv"
)

type Service interface {
	Create(name string, hostname string) (*model.Group, error)
	FindById(id uint) (*model.Group, error)
	AddUser(groupId uint, userId uint) error
	AddClusterConfiguration(clusterConfiguration *model.ClusterConfiguration) error
	GetClusterConfiguration(groupId uint) (*model.ClusterConfiguration, error)
	FindByName(name string) (*model.Group, error)
}

func ProvideService(groupRepository Repository, userRepository user.Repository) Service {
	return &service{
		groupRepository,
		userRepository,
	}
}

type service struct {
	groupRepository Repository
	userRepository  user.Repository
}

func (s *service) FindByName(name string) (*model.Group, error) {
	return s.groupRepository.FindByName(name)
}

func (s *service) Create(name string, hostname string) (*model.Group, error) {
	group := &model.Group{
		Name:     name,
		Hostname: hostname,
	}

	err := s.groupRepository.Create(group)

	if err != nil {
		return nil, apperror.NewBadRequest(err.Error())
	}

	return group, err
}

func (s *service) FindById(id uint) (*model.Group, error) {
	return s.groupRepository.FindById(id)
}

func (s *service) AddUser(groupId uint, userId uint) error {
	group, err := s.FindById(groupId)
	if err != nil {
		return apperror.NewNotFound("group", strconv.Itoa(int(groupId)))
	}

	u, err := s.userRepository.FindById(userId)
	if err != nil {
		return apperror.NewNotFound("user", strconv.Itoa(int(userId)))
	}

	return s.groupRepository.AddUser(group, u)
}

func (s *service) AddClusterConfiguration(clusterConfiguration *model.ClusterConfiguration) error {
	return s.groupRepository.AddClusterConfiguration(clusterConfiguration)
}

func (s *service) GetClusterConfiguration(groupId uint) (*model.ClusterConfiguration, error) {
	return s.groupRepository.GetClusterConfiguration(groupId)
}
