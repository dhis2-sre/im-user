package group

import (
	"github.com/dhis2-sre/im-users/pgk/helper"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (h *Handler) Create(c *gin.Context) {
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
