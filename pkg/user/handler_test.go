package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestHandler_SignUp(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"

	c := config.Config{}

	userService := &mockUserService{}
	userService.
		On("SignUp", email, password).
		Return(&model.User{Model: gorm.Model{ID: id}, Email: email, Password: password})
	tokenService := &mockTokenService{}

	handler := NewHandler(c, userService, tokenService)

	request := SignUpRequest{
		Email:    email,
		Password: password,
	}
	req := newRequest(t, request)

	r := gin.Default()
	r.POST("/users", handler.SignUp)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	body := recorder.Body
	user := &model.User{}
	err := json.Unmarshal(body.Bytes(), user)
	require.NoError(t, err)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, "", user.Password)

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func newRequest(t *testing.T, request any) *http.Request {
	body, err := json.Marshal(request)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req
}

func TestHandler_SignIn_Happy(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"

	accessToken := "access token"
	tokenType := ""
	refreshToken := ""
	expiresIn := uint(0)

	c := config.Config{}

	userService := &mockUserService{}
	userService.
		On("SignIn", email, password).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		})
	tokenService := &mockTokenService{}
	tokenService.
		On("GetTokens", mock.AnythingOfType("*model.User"), mock.AnythingOfType("string")).
		Return(&token.Tokens{
			AccessToken:  accessToken,
			TokenType:    tokenType,
			RefreshToken: refreshToken,
			ExpiresIn:    expiresIn,
		}, nil)

	r := gin.Default()
	authentication := middleware.NewAuthentication(userService, tokenService)
	r.Use(authentication.BasicAuthentication)
	handler := NewHandler(c, userService, tokenService)
	r.POST("/tokens", handler.SignIn)

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/tokens", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(email, password)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	body := recorder.Body
	tokens := &token.Tokens{}
	err = json.Unmarshal(body.Bytes(), tokens)
	require.NoError(t, err)

	assert.Equal(t, accessToken, tokens.AccessToken)
	assert.Equal(t, tokenType, tokens.TokenType)
	assert.Equal(t, refreshToken, tokens.RefreshToken)
	assert.Equal(t, expiresIn, tokens.ExpiresIn)

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestHandler_SignIn_GetTokensError(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"
	errorMessage := "some err"

	c := config.Config{}

	userService := &mockUserService{}
	userService.
		On("SignIn", email, password).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		})
	tokenService := &mockTokenService{}
	tokenService.
		On("GetTokens", mock.AnythingOfType("*model.User"), "").
		Return(nil, errors.New(errorMessage))

	r := gin.Default()
	authentication := middleware.NewAuthentication(userService, tokenService)
	r.Use(middleware.ErrorHandler(), authentication.BasicAuthentication)
	handler := NewHandler(c, userService, tokenService)
	r.POST("/tokens", handler.SignIn)

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/tokens", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(email, password)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	assert.Contains(t, recorder.Body.String(), "something went wrong. We'll look into it if you send us the id")

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestHandler_Me(t *testing.T) {
	bearerToken := "token"
	var id uint = 1
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"

	c := config.Config{}

	userService := &mockUserService{}
	userService.
		On("FindById", id).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		}, nil)
	tokenService := &mockTokenService{}
	tokenService.
		On("ValidateAccessToken", bearerToken).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
		}, nil)

	r := gin.Default()
	authentication := middleware.NewAuthentication(userService, tokenService)
	r.Use(authentication.TokenAuthentication)
	handler := NewHandler(c, userService, tokenService)
	r.GET("/me", handler.Me)

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/me", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", bearerToken)
	r.ServeHTTP(recorder, req)

	body := recorder.Body
	user := &model.User{}
	err = json.Unmarshal(body.Bytes(), user)
	require.NoError(t, err)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, "", user.Password)

	userService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

type mockUserService struct{ mock.Mock }

func (s *mockUserService) FindOrCreate(email string, password string) (*model.User, error) {
	panic("implement me")
}

func (s *mockUserService) SignUp(email string, password string) (*model.User, error) {
	called := s.Called(email, password)
	return called.Get(0).(*model.User), nil
}

func (s *mockUserService) SignIn(email string, password string) (*model.User, error) {
	called := s.Called(email, password)
	return called.Get(0).(*model.User), nil
}

func (s *mockUserService) FindById(id uint) (*model.User, error) {
	called := s.Called(id)
	return called.Get(0).(*model.User), nil
}

type mockTokenService struct{ mock.Mock }

func (t *mockTokenService) ValidateAccessToken(tokenString string) (*model.User, error) {
	called := t.Called(tokenString)
	return called.Get(0).(*model.User), nil
}

func (t *mockTokenService) GetTokens(user *model.User, previousTokenId string) (*token.Tokens, error) {
	called := t.Called(user, previousTokenId)
	tokens, ok := called.Get(0).(*token.Tokens)
	if ok {
		return tokens, nil
	} else {
		return nil, called.Error(1)
	}
}

func (t *mockTokenService) ValidateRefreshToken(tokenString string) (*token.RefreshTokenData, error) {
	panic("implement me")
}

func (t *mockTokenService) SignOut(userId uint) error {
	panic("implement me")
}
