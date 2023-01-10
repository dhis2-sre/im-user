package group

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Create_Happy(t *testing.T) {
	name := "name"
	hostname := "hostname"

	groupService := &mockGroupService{}
	groupService.
		On("Create", name, hostname).
		Return(&model.Group{Name: name, Hostname: hostname}, nil)

	handler := NewHandler(groupService)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = newRequest(t, http.MethodPost, "/groups", &CreateGroupRequest{Name: name, Hostname: hostname})

	handler.Create(c)

	actual := recorder.Code
	assert.Equal(t, http.StatusCreated, actual)

	groupService.AssertExpectations(t)
}

func TestHandler_Create_CanNotCreateGroup(t *testing.T) {
	name := "name"
	hostname := "hostname"
	errorMessage := "some error"

	groupService := &mockGroupService{}
	groupService.
		On("Create", name, hostname).
		Return(nil, errors.New(errorMessage))

	handler := NewHandler(groupService)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = newRequest(t, http.MethodPost, "/groups", &CreateGroupRequest{Name: name, Hostname: hostname})

	handler.Create(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	errs := c.Errors.Errors()
	assert.Equal(t, 1, len(errs))
	assert.True(t, strings.HasPrefix(errs[0], errorMessage))

	groupService.AssertExpectations(t)
}

func newRequest(t *testing.T, method string, path string, request any) *http.Request {
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	req, err := http.NewRequest(method, path, bytes.NewReader(body))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req
}

type mockGroupService struct{ mock.Mock }

func (m *mockGroupService) Create(name string, hostname string) (*model.Group, error) {
	called := m.Called(name, hostname)
	group, ok := called.Get(0).(*model.Group)
	if ok {
		return group, nil
	} else {
		return nil, called.Error(1)
	}
}

func (m *mockGroupService) AddUser(groupName string, userId uint) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockGroupService) AddClusterConfiguration(clusterConfiguration *model.ClusterConfiguration) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockGroupService) GetClusterConfiguration(groupName string) (*model.ClusterConfiguration, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockGroupService) Find(name string) (*model.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockGroupService) FindOrCreate(name string, hostname string) (*model.Group, error) {
	//TODO implement me
	panic("implement me")
}
