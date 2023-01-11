package user

import (
	"errors"
	"testing"

	"github.com/dhis2-sre/im-user/internal/errdef"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Test_service_SignUp_Happy(t *testing.T) {
	email := "email"
	password := "password"

	createUserRepository := &mockUserRepository{}
	createUserRepository.
		On("create", mock.AnythingOfType("*model.User")).
		Return(nil, nil)

	s := NewService(createUserRepository)

	user, err := s.SignUp(email, password)
	require.NoError(t, err)

	createUserRepository.AssertExpectations(t)

	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.Password)
}

func Test_service_SignUp_UserExists(t *testing.T) {
	email := "email"
	password := "password"
	errorMessage := "duplicated user"

	repository := &mockUserRepository{}
	repository.
		On("create", mock.AnythingOfType("*model.User")).
		Return(errors.New(errorMessage))

	s := NewService(repository)

	_, err := s.SignUp(email, password)
	require.Error(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, errorMessage, err.Error())
}

func Test_service_SignIn_Happy(t *testing.T) {
	var id uint = 1
	email := "email"
	password := "passwordpasswordpasswordpassword"
	hashedPassword := "c55d1333f8567be7bfcc00fcae72720d30ae465cf62fb31c33303b707a18c2ca.f053011abe74ca660c2d98de8747ad0d6a6f401ccfb513e68127b1a46b42ed19"

	repository := &mockUserRepository{}
	repository.
		On("findByEmail", email).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: hashedPassword,
		}, nil)

	s := NewService(repository)

	user, err := s.SignIn(email, password)
	require.NoError(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, hashedPassword, user.Password)
}

func Test_service_SignIn_BadCredentials(t *testing.T) {
	email := "email"
	password := "password"
	hashedPassword := "c55d1333f8567be7bfcc00fcae72720d30ae465cf62fb31c33303b707a18c2ca.f053011abe74ca660c2d98de8747ad0d6a6f401ccfb513e68127b1a46b42ed19"
	errorMessage := "invalid email and password combination"

	repository := &mockUserRepository{}
	repository.
		On("findByEmail", email).
		Return(&model.User{Password: hashedPassword}, nil)

	s := NewService(repository)

	user, err := s.SignIn(email, password)
	require.Error(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, user == nil, true)
	assert.Equal(t, errorMessage, err.Error())
}

func Test_service_SignIn_NotFound(t *testing.T) {
	email := "email"
	password := "password"
	errorMessage := "invalid email and password combination"

	repository := &mockUserRepository{}
	repository.
		On("findByEmail", email).
		Return(nil, errdef.NewNotFound(errors.New(errorMessage)))

	s := NewService(repository)

	user, err := s.SignIn(email, password)
	require.Error(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, user == nil, true)
	assert.Equal(t, errorMessage, err.Error())
}

func Test_service_FindById_Happy(t *testing.T) {
	var id uint = 1
	email := "email"
	password := "password"

	repository := &mockUserRepository{}
	repository.
		On("findById", id).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		}, nil)

	s := NewService(repository)

	user, err := s.FindById(id)
	require.NoError(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
}

func Test_service_FindById_NotFound(t *testing.T) {
	var id uint = 1
	errorMessage := "not found"

	repository := &mockUserRepository{}
	repository.
		On("findById", id).
		Return(nil, errdef.NewNotFound(errors.New(errorMessage)))

	s := NewService(repository)

	user, err := s.FindById(id)
	require.Error(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, user == nil, true)
	assert.Equal(t, errorMessage, err.Error())
}

func Test_service_FindOrCreate_Happy(t *testing.T) {
	var id uint = 1
	email := "email"
	password := "password"

	repository := &mockUserRepository{}
	repository.
		On("findOrCreate", mock.AnythingOfType("*model.User")).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		}, nil)

	s := NewService(repository)

	user, err := s.FindOrCreate(email, password)
	require.NoError(t, err)

	repository.AssertExpectations(t)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
}

type mockUserRepository struct {
	mock.Mock
}

func (r *mockUserRepository) create(user *model.User) error {
	args := r.Called(user)
	err := args.Error(0)
	return err
}

func (r *mockUserRepository) findByEmail(email string) (*model.User, error) {
	args := r.Called(email)
	user, ok := args.Get(0).(*model.User)
	if ok {
		return user, nil
	} else {
		return nil, errdef.NewNotFound(errors.New(""))
	}
}

func (r *mockUserRepository) findById(id uint) (*model.User, error) {
	args := r.Called(id)
	user, ok := args.Get(0).(*model.User)
	if ok {
		return user, nil
	} else {
		return nil, errdef.NewNotFound(errors.New("not found"))
	}
}

func (r *mockUserRepository) findOrCreate(email *model.User) (*model.User, error) {
	args := r.Called(email)
	user, ok := args.Get(0).(*model.User)
	if ok {
		return user, nil
	} else {
		return nil, errdef.NewNotFound(errors.New(""))
	}
}
