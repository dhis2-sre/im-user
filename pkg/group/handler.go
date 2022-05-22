package group

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/dhis2-sre/im-user/internal/apperror"
	"github.com/dhis2-sre/im-user/internal/handler"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/gin-gonic/gin"
)

func ProvideHandler(
	groupService Service,
	userService user.Service,
) Handler {
	return Handler{
		groupService,
		userService,
	}
}

type Handler struct {
	groupService Service
	userService  user.Service
}

type CreateGroupRequest struct {
	Name     string `json:"name" binding:"required"`
	Hostname string `json:"hostname" binding:"required"`
}

// Create group
// swagger:route POST /groups groupCreate
//
// Create group
//
// security:
//   oauth2:
//
// responses:
//   201: Group
//   400: Error
//   403: Error
//   415: Error
func (h Handler) Create(c *gin.Context) {
	var request CreateGroupRequest

	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	group, err := h.groupService.Create(request.Name, request.Hostname)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, group)
}

// AddUserToGroup group
// swagger:route POST /groups/{groupName}/users/{userId} addUserToGroup
//
// Add user to group
//
// security:
//   oauth2:
//
// responses:
//   201: Group
//   400: Error
//   403: Error
//   415: Error
func (h Handler) AddUserToGroup(c *gin.Context) {
	userIdString := c.Param("userId")
	groupName := c.Param("groupName")

	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		message := fmt.Sprintf("Failed to parse userId: %s", err)
		badRequest := apperror.NewBadRequest(message)
		_ = c.Error(badRequest)
		return
	}

	err = h.groupService.AddUser(groupName, uint(userId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusCreated)
}

type CreateClusterConfigurationRequest struct {
	KubernetesConfiguration *multipart.FileHeader `form:"kubernetesConfiguration" binding:"required"`
}

// AddClusterConfiguration group
// swagger:route POST /groups/{name}/cluster-configuration addClusterConfigurationToGroup
//
// Add cluster configuration to group
//
// security:
//   oauth2:
//
// responses:
//   201: Group
//   400: Error
//   403: Error
//   415: Error
func (h Handler) AddClusterConfiguration(c *gin.Context) {
	var request CreateClusterConfigurationRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	groupName := c.Param("name")
	if groupName == "" {
		badRequest := apperror.NewBadRequest("group name missing")
		_ = c.Error(badRequest)
		return
	}

	kubernetesConfiguration, err := h.getBytes(request.KubernetesConfiguration)
	if err != nil {
		_ = c.Error(err)
		return
	}

	clusterConfiguration := &model.ClusterConfiguration{
		GroupName:               groupName,
		KubernetesConfiguration: kubernetesConfiguration,
	}

	err = h.groupService.AddClusterConfiguration(clusterConfiguration)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h Handler) getBytes(file *multipart.FileHeader) ([]byte, error) {
	if file == nil {
		return nil, nil
	}

	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(openedFile)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Find group
// swagger:route GET /groups/{name} findGroupByName
//
// Return a group by name
//
// responses:
//   200: Group
//   403: Error
//   404: Error
//   415: Error
//
// security:
//   oauth2:
func (h Handler) Find(c *gin.Context) {
	name := c.Param("name")

	group, err := h.groupService.Find(name)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, group)
}
