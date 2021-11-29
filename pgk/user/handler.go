package user

import (
	"github.com/dhis2-sre/im-users/pgk/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProvideHandler(service Service) Handler {
	return Handler{
		service,
	}
}

type Handler struct {
	service Service
}

// Signup godoc
// @Summary User sign in
// @Description Posting username (email) and password... And user is returned
// @Tags Public
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{} //model.User
// @Failure 400 {object} map[string]interface{}
// @Router /signup [post]
// @Param signupRequest body SignupRequest true "Email and Password json object"
func (h *Handler) Signup(c *gin.Context) {
	type SignupRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=16,lte=128"`
	}

	var request SignupRequest

	err := helper.DataBinder(c, &request)
	if err != nil {
		// TODO: Error handling for the error handler... :-/ ?
		c.Error(err)
		return
	}

	user, err := h.service.Signup(request.Email, request.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
