package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/gorm"

	"github.com/stretchr/testify/require"

	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_SignUp(t *testing.T) {
	userRepository := &mockUserRepository{}
	userRepository.
		On("create", mock.MatchedBy(func(user *model.User) bool {
			return user.Email == "someone@something.org" && len(user.Password) == 129
		})).
		Return(nil)
	userService := NewService(userRepository)
	tokenService := &mockTokenService{}
	handler := NewHandler(config.Config{}, userService, tokenService)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	signUpRequest := SignUpRequest{Email: "someone@something.org", Password: "passwordpasswordpasswordpassword"}
	c.Request = newPost(t, "", signUpRequest)

	handler.SignUp(c)

	assert.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusCreated, &model.User{Email: "someone@something.org", Password: ""})
	userRepository.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestHandler_SignIn(t *testing.T) {
	userRepository := &mockUserRepository{}
	userService := NewService(userRepository)
	tokenService := &mockTokenService{}
	tokens := &token.Tokens{AccessToken: "access token", TokenType: "token type", RefreshToken: "refresh token", ExpiresIn: uint(123)}
	tokenService.
		On("GetTokens", mock.AnythingOfType("*model.User"), mock.AnythingOfType("string")).
		Return(tokens, nil)
	handler := NewHandler(config.Config{}, userService, tokenService)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")

	handler.SignIn(c)

	require.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusCreated, tokens)
	userRepository.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestHandler_Me(t *testing.T) {
	user := &model.User{Model: gorm.Model{ID: 1}, Email: "someone@something.org", Password: "passwordpasswordpasswordpassword"}
	userRepository := &mockUserRepository{}
	userRepository.
		On("findById", uint(1)).
		Return(user, nil)
	userService := NewService(userRepository)
	tokenService := &mockTokenService{}
	handler := NewHandler(config.Config{}, userService, tokenService)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")

	handler.Me(c)

	require.Empty(t, c.Errors)
	expectedUser := &model.User{Model: gorm.Model{ID: 1}, Email: "someone@something.org", Password: ""}
	assertResponse(t, w, http.StatusOK, expectedUser)
	userRepository.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

/*
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
*/

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

type mockUserRepository struct{ mock.Mock }

func (m *mockUserRepository) create(user *model.User) error {
	called := m.Called(user)
	return called.Error(0)
}

func (m *mockUserRepository) findByEmail(email string) (*model.User, error) {
	panic("implement me")
}

func (m *mockUserRepository) findById(id uint) (*model.User, error) {
	called := m.Called(id)
	return called.Get(0).(*model.User), nil
}

func (m *mockUserRepository) findOrCreate(user *model.User) (*model.User, error) {
	panic("implement me")
}

func newContext(w *httptest.ResponseRecorder, group string) *gin.Context {
	user := &model.User{
		Model: gorm.Model{ID: 1},
		Groups: []model.Group{
			{Name: group},
		},
	}
	c, _ := gin.CreateTestContext(w)
	c.Set("user", user)
	return c
}

func newPost(t *testing.T, path string, jsonBody any) *http.Request {
	body, err := json.Marshal(jsonBody)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "token")

	return req
}

func assertResponse[V any](t *testing.T, rec *httptest.ResponseRecorder, expectedCode int, expectedBody V) {
	require.Equal(t, expectedCode, rec.Code, "HTTP status code does not match")
	assertJSON(t, rec.Body, expectedBody)
}

func assertJSON[V any](t *testing.T, body *bytes.Buffer, expected V) {
	actualBody := new(V)
	err := json.Unmarshal(body.Bytes(), &actualBody)
	require.NoError(t, err)
	require.Equal(t, expected, *actualBody, "HTTP response body does not match")
}
