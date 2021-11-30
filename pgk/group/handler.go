package group

import (
	"fmt"
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/pgk/helper"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/gin-gonic/gin"
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
// @Success 201 {object} map[string]interface{} //model.Group
// @Failure 400 {object} map[string]interface{}
// @Router /groups [post]
// @Param createGroupRequest body CreateGroupRequest true "Create group request"
// @Security OAuth2Password
func (h Handler) Create(c *gin.Context) {
	var request CreateGroupRequest

	if err := helper.DataBinder(c, &request); err != nil {
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
// @Success 201 {none} string
// @Failure 400 {object} map[string]interface{}
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
