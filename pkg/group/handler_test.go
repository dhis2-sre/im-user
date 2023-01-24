package group

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gorm.io/gorm"

	"github.com/dhis2-sre/im-user/internal/errdef"
	"github.com/stretchr/testify/require"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Find(t *testing.T) {
	group := &model.Group{Name: "name"}
	repository := &mockGroupRepository{}
	repository.
		On("find", "name").
		Return(group, nil)
	service := NewService(repository, nil)
	handler := NewHandler(service)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.AddParam("name", "name")

	handler.Find(c)

	assert.Empty(t, c.Errors)
	var body model.Group
	assertResponse(t, recorder, http.StatusOK, &body, group)
	repository.AssertExpectations(t)
}

func assertResponse(t *testing.T, rec *httptest.ResponseRecorder, expectedCode int, bodyType any, expectedBody any) {
	assert.Equal(t, expectedCode, rec.Code, "HTTP status code does not match")
	assertJSON(t, rec.Body, bodyType, expectedBody)
}

func assertJSON(t *testing.T, body *bytes.Buffer, v any, expected any) {
	err := json.Unmarshal(body.Bytes(), v)
	require.NoError(t, err)
	assert.Equal(t, expected, v, "HTTP response body does not match")
}

func TestHandler_Find_NotFound(t *testing.T) {
	repository := &mockGroupRepository{}
	repository.
		On("find", "name").
		Return(nil, errdef.NewNotFound(errors.New("not found")))
	service := NewService(repository, nil)
	handler := NewHandler(service)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.AddParam("name", "name")

	handler.Find(c)

	assert.Equal(t, 1, len(c.Errors))
	assert.Error(t, c.Errors[0].Err, "not found")
	repository.AssertExpectations(t)
}

type mockUserService struct{ mock.Mock }

func (m *mockUserService) FindById(id uint) (*model.User, error) {
	called := m.Called(id)
	user, ok := called.Get(0).(*model.User)
	if ok {
		return user, nil
	}
	return nil, called.Error(1)
}

func TestHandler_AddUserToGroup(t *testing.T) {
	user := &model.User{
		Model: gorm.Model{ID: 1},
	}
	repository := &mockGroupRepository{}
	repository.
		On("find", "name").
		Return(&model.Group{Name: "name"})
	repository.
		On("addUser", &model.Group{Name: "name"}, user).
		Return(nil)
	userService := &mockUserService{}
	userService.
		On("FindById", uint(1)).
		Return(user, nil)
	service := NewService(repository, userService)
	handler := NewHandler(service)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.AddParam("userId", strconv.FormatUint(uint64(1), 10))
	c.AddParam("group", "name")

	handler.AddUserToGroup(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Empty(t, c.Errors)
	repository.AssertExpectations(t)
	userService.AssertExpectations(t)
}

func TestHandler_AddUserToGroup_BadUserId(t *testing.T) {
	repository := &mockGroupRepository{}
	service := NewService(repository, nil)
	handler := NewHandler(service)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.AddParam("userId", "-")
	c.AddParam("group", "name")

	handler.AddUserToGroup(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, 1, len(c.Errors))
	errorMessage := fmt.Sprintf("failed to parse userId: strconv.ParseUint: parsing %q: invalid syntax", "-")
	assert.Error(t, c.Errors[0].Err, errorMessage)
	repository.AssertExpectations(t)
}

func TestHandler_AddUserToGroup_RepositoryError(t *testing.T) {
	group := &model.Group{Name: "name"}
	user := &model.User{
		Model: gorm.Model{ID: 1},
	}
	repository := &mockGroupRepository{}
	repository.
		On("find", "name").
		Return(group)
	repository.
		On("addUser", group, user).
		Return(errdef.NewNotFound(errors.New("some error")))
	userService := &mockUserService{}
	userService.
		On("FindById", uint(1)).
		Return(user, nil)
	service := NewService(repository, userService)
	handler := NewHandler(service)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.AddParam("userId", "1")
	c.AddParam("group", "name")

	handler.AddUserToGroup(c)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, 1, len(c.Errors))
	err := c.Errors[0].Err
	assert.True(t, errdef.IsNotFound(err))
	assert.ErrorContains(t, err, "some error")
	repository.AssertExpectations(t)
}

func TestHandler_Create(t *testing.T) {
	repository := &mockGroupRepository{}
	repository.
		On("create", &model.Group{Name: "name", Hostname: "hostname"}).
		Return(nil)
	service := NewService(repository, nil)
	handler := NewHandler(service)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = newPost(t, "/groups", &CreateGroupRequest{Name: "name", Hostname: "hostname"})

	handler.Create(c)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	repository.AssertExpectations(t)
}

func TestHandler_Create_CanNotCreateGroup(t *testing.T) {
	group := &model.Group{Name: "name", Hostname: "hostname"}
	repository := &mockGroupRepository{}
	repository.
		On("create", group).
		Return(errors.New("some error"))
	service := NewService(repository, nil)
	handler := NewHandler(service)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = newPost(t, "/groups", &CreateGroupRequest{Name: "name", Hostname: "hostname"})

	handler.Create(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, 1, len(c.Errors))
	assert.ErrorContains(t, c.Errors[0].Err, "some error")
	repository.AssertExpectations(t)
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

type mockGroupRepository struct{ mock.Mock }

func (m *mockGroupRepository) create(group *model.Group) error {
	called := m.Called(group)
	return called.Error(0)
}

func (m *mockGroupRepository) addUser(group *model.Group, user *model.User) error {
	called := m.Called(group, user)
	return called.Error(0)
}

func (m *mockGroupRepository) addClusterConfiguration(configuration *model.ClusterConfiguration) error {
	panic("implement me")
}

func (m *mockGroupRepository) getClusterConfiguration(groupName string) (*model.ClusterConfiguration, error) {
	panic("implement me")
}

func (m *mockGroupRepository) find(name string) (*model.Group, error) {
	called := m.Called(name)
	group, ok := called.Get(0).(*model.Group)
	if ok {
		return group, nil
	}
	return nil, called.Error(1)
}

func (m *mockGroupRepository) findOrCreate(group *model.Group) (*model.Group, error) {
	panic("implement me")
}
