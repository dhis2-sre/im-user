package middleware

import (
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
	id := uint(1)
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"

	userService := &mockUserService{}
	userService.On("SignIn", email, password).Return(&model.User{
		Model:    gorm.Model{ID: id},
		Email:    email,
		Password: password,
	}, nil)

	tokenService := &mockTokenService{}
	m := AuthenticationMiddleware{
		userService:  userService,
		tokenService: tokenService,
	}

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(email, password)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	_, exists := c.Get("user")
	assert.False(t, exists)

	m.BasicAuthentication(c)

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

func TestAuthenticationMiddleware_TokenAuthentication(t *testing.T) {
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
	return called.Get(0).(*model.User), nil
}

func (s *mockUserService) FindById(id uint) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

type mockTokenService struct{ mock.Mock }

func (t *mockTokenService) ValidateAccessToken(tokenString string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
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
