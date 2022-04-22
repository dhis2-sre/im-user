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
// swagger:route POST /groups/{groupId}/users/{userId} groupAddUserToGroup
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
	groupIdString := c.Param("groupId")

	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		message := fmt.Sprintf("Failed to parse userId: %s", err)
		badRequest := apperror.NewBadRequest(message)
		_ = c.Error(badRequest)
		return
	}

	groupId, err := strconv.ParseUint(groupIdString, 10, 64)
	if err != nil {
		message := fmt.Sprintf("Failed to parse groupId: %s", err)
		badRequest := apperror.NewBadRequest(message)
		_ = c.Error(badRequest)
		return
	}

	err = h.groupService.AddUser(uint(groupId), uint(userId))
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
// swagger:route POST /groups/{groupId}/cluster-configuration groupAddClusterConfigurationToGroup
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

	groupIdString := c.Param("groupId")
	groupId, err := strconv.ParseUint(groupIdString, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest(err.Error())
		_ = c.Error(badRequest)
		return
	}

	kubernetesConfiguration, err := h.getBytes(request.KubernetesConfiguration)
	if err != nil {
		_ = c.Error(err)
		return
	}

	clusterConfiguration := &model.ClusterConfiguration{
		GroupID:                 uint(groupId),
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

// NameToId group
// swagger:route GET /groups-name-to-id/{name} groupNameToId
//
// Find group id by name
//
// Security:
//  oauth2:
//
// responses:
//   200:
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) NameToId(c *gin.Context) {
	name := c.Param("name")

	/*
		u, err := handler.GetUserFromContext(c)
		if err != nil {
			_ = c.Error(err)
			return
		}

		userWithGroups, err := h.userService.FindById(u.ID)
		if err != nil {
			notFound := apperror.NewNotFound("user", strconv.Itoa(int(u.ID)))
			_ = c.Error(notFound)
			return
		}
	*/
	group, err := h.groupService.FindByName(name)
	if err != nil {
		notFound := apperror.NewNotFound("group", name)
		_ = c.Error(notFound)
		return
	}

	// No authorization checks will be done here, if someone knows the name of the group they can have the id too
	/*
		instance := &model.Instance{GroupID: group.ID}
		canRead := g.instanceAuthorizer.CanRead(userWithGroups, instance)

		if !canRead {
			unauthorized := apperrors.NewUnauthorized("Read access denied")
			handler.HandleError(c, unauthorized)
			return
		}
	*/
	c.JSON(http.StatusOK, group.ID)
}

// FindById group
// swagger:route GET /groups/{id} findGroupById
//
// Return a group by id
//
// responses:
//   200: Group
//   403: Error
//   404: Error
//   415: Error
//
// security:
//   oauth2:
func (h Handler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("Error parsing id")
		_ = c.Error(badRequest)
		return
	}

	group, err := h.groupService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, group)
}
