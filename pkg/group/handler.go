package group

import (
	"fmt"
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/internal/handler"
	"github.com/dhis2-sre/im-users/pkg/model"
	"github.com/dhis2-sre/im-users/pkg/user"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
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

// Create godoc
// @Summary Create group
// @Description Posting name and hostname...
// @Tags Administrative
// @Accept json
// @Produce json
// @Success 201 {object} dto.Group
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 415 {object} map[string]interface{}
// @Router /groups [post]
// @Param createGroupRequest body CreateGroupRequest true "Create group request"
// @Security OAuth2Password
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

// AddUserToGroup godoc
// @Summary Add user to group
// @Description Add user to group
// @Tags Administrative
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 415 {object} map[string]interface{}
// @Router /users/{userId}/groups/{groupId} [post]
// @Param userId path string true "User id"
// @Param groupId path string true "Group id"
// @Security OAuth2Password
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

type createClusterConfigurationRequest struct {
	KubernetesConfiguration *multipart.FileHeader `form:"kubernetesConfiguration" binding:"required"`
}

// AddClusterConfiguration godoc
// @Summary Add cluster configuration to a group
// @Description Add cluster configuration to a group...
// @Tags Administrative
// @Accept multipart/form-data
// @Produce json
// @Success 201 {object} map[string]interface{} //model.Group
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 415 {object} map[string]interface{}
// @Router /groups/{groupId}/cluster-configuration [post]
// @Param groupId path string true "Group ID"
// @Param kubernetesConfiguration formData file true "SOPS encrypted Kubernetes configuration file"
// @Security OAuth2Password
func (h Handler) AddClusterConfiguration(c *gin.Context) {
	var request createClusterConfigurationRequest
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

// NameToId godoc
// @Summary Group id by group name
// @Description Return group id given group name
// @Tags Restricted
// @Accept json
// @Produce json
// @Success 200 {object} uint
// @Failure 401 {object} string
// @Failure 403 {object} string
// @Failure 404 {object} string
// @Failure 415 {object} string
// @Router /groups-name-to-id/{name} [get]
// @Param name path string true "Group name"
// @Security OAuth2Password
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

// FindById godoc
// swagger:route GET /groups/{id} getGroupById
// Return a group by id
// responses:
//   200: Group
//   403: Error
//   404: Error
//   415: Error
func (h Handler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
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
