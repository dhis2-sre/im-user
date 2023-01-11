package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAuthenticationMiddleware_BasicAuthentication_Happy(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"

	userService := &mockUserService{}
	userService.On("SignIn", email, password).Return(&model.User{
		Model:    gorm.Model{ID: id},
		Email:    email,
		Password: password,
	}, nil)

	tokenService := &mockTokenService{}
	authentication := NewAuthentication(userService, tokenService)

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(email, password)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	_, exists := c.Get("user")
	assert.False(t, exists)

	authentication.BasicAuthentication(c)

	value, exists := c.Get("user")
	assert.True(t, exists)
	user, ok := value.(*model.User)
	assert.True(t, ok)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestAuthenticationMiddleware_BasicAuthentication_NoCredentials(t *testing.T) {
	userService := &mockUserService{}
	tokenService := &mockTokenService{}
	authentication := NewAuthentication(userService, tokenService)

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	_, exists := c.Get("user")
	assert.False(t, exists)

	authentication.BasicAuthentication(c)

	_, exists = c.Get("user")
	assert.False(t, exists)

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestAuthenticationMiddleware_BasicAuthentication_WrongCredentials(t *testing.T) {
	email := "someone@something.org"
	password := "password"

	userService := &mockUserService{}
	userService.
		On("SignIn", email, password).
		Return(nil, errors.New("wrong credentials"))
	tokenService := &mockTokenService{}
	authentication := NewAuthentication(userService, tokenService)

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(email, password)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	_, exists := c.Get("user")
	assert.False(t, exists)

	authentication.BasicAuthentication(c)

	_, exists = c.Get("user")
	assert.False(t, exists)

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestAuthenticationMiddleware_TokenAuthentication_Happy(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "password"

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer token")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	userService := &mockUserService{}
	tokenService := &mockTokenService{}
	tokenService.
		On("ValidateAccessToken", mock.AnythingOfType("string")).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		}, nil)
	authentication := NewAuthentication(userService, tokenService)

	_, exists := c.Get("user")
	assert.False(t, exists)

	authentication.TokenAuthentication(c)

	value, exists := c.Get("user")
	assert.True(t, exists)
	user, ok := value.(*model.User)
	assert.True(t, ok)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
}

func TestAuthenticationMiddleware_TokenAuthentication_TokenValidationFail(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer token")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	userService := &mockUserService{}
	tokenService := &mockTokenService{}
	errorMessage := "validation fail"
	tokenService.
		On("ValidateAccessToken", mock.AnythingOfType("string")).
		Return(nil, errors.New(errorMessage))
	authentication := NewAuthentication(userService, tokenService)

	_, exists := c.Get("user")
	assert.False(t, exists)

	authentication.TokenAuthentication(c)

	_, exists = c.Get("user")
	assert.False(t, exists)
}

func TestAuthenticationMiddleware_TokenAuthentication_ExternalError(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "password"

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer token")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	userService := &mockUserService{}
	tokenService := &mockTokenService{}
	tokenService.
		On("ValidateAccessToken", mock.AnythingOfType("string")).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		}, nil)
	authentication := NewAuthentication(userService, tokenService)

	_ = c.Error(errors.New("some error which wasn't handled properly"))
	assert.NoError(t, err)

	_, exists := c.Get("user")
	assert.False(t, exists)

	authentication.TokenAuthentication(c)

	_, exists = c.Get("user")
	assert.False(t, exists)
}

type mockUserService struct{ mock.Mock }

func (s *mockUserService) FindOrCreate(email string, password string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *mockUserService) SignUp(email string, password string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *mockUserService) SignIn(email string, password string) (*model.User, error) {
	called := s.Called(email, password)
	user, ok := called.Get(0).(*model.User)
	if ok {
		return user, nil
	} else {
		return nil, called.Error(1)
	}
}

func (s *mockUserService) FindById(id uint) (*model.User, error) {
	called := s.Called(id)
	user, ok := called.Get(0).(*model.User)
	if ok {
		return user, nil
	} else {
		return nil, called.Error(1)
	}
}

type mockTokenService struct{ mock.Mock }

func (t *mockTokenService) ValidateAccessToken(tokenString string) (*model.User, error) {
	called := t.Called(tokenString)
	user, ok := called.Get(0).(*model.User)
	if ok {
		return user, nil
	} else {
		return nil, called.Error(1)
	}
}

func (t *mockTokenService) GetTokens(user *model.User, previousTokenId string) (*token.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (t *mockTokenService) ValidateRefreshToken(tokenString string) (*token.RefreshTokenData, error) {
	//TODO implement me
	panic("implement me")
}

func (t *mockTokenService) SignOut(userId uint) error {
	//TODO implement me
	panic("implement me")
}
