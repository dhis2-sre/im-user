package user

import (
	"github.com/dhis2-sre/im-user/internal/apperror"
	"github.com/dhis2-sre/im-user/internal/handler"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ProvideHandler(config config.Config, userService Service, tokenService token.Service) Handler {
	return Handler{
		config,
		userService,
		tokenService,
	}
}

type Handler struct {
	config       config.Config
	userService  Service
	tokenService token.Service
}

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=16,lte=128"`
}

// SignUp user
// swagger:route POST /users signUp
//
// SignUp user
//
// responses:
//   201: User
//   400: Error
//   415: Error
func (h *Handler) SignUp(c *gin.Context) {
	var request SignupRequest

	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	user, err := h.userService.SignUp(request.Email, request.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// SignIn user
// swagger:route POST /tokens signIn
//
// Return user tokens
//
// security:
//   basicAuth:
//
// responses:
//   201: Tokens
//   403: Error
//   404: Error
//   415: Error
func (h *Handler) SignIn(c *gin.Context) {
	user, err := handler.GetUserFromContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	tokens, err := h.tokenService.GetTokens(user, "")
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, tokens)
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// RefreshToken user
// swagger:route POST /refresh refreshToken
//
// Refresh user tokens
//
// responses:
//   201: User
//   400: Error
//   415: Error
func (h Handler) RefreshToken(c *gin.Context) {
	var request RefreshTokenRequest

	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	refreshToken, err := h.tokenService.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err := h.userService.FindById(refreshToken.UserId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	tokens, err := h.tokenService.GetTokens(user, refreshToken.ID.String())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, tokens)
}

// Me user
// swagger:route GET /me me
//
// Return user details
//
// security:
//   oauth2:
//
// responses:
//   200: User
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Me(c *gin.Context) {
	user, err := handler.GetUserFromContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	userWithGroups, err := h.userService.FindById(user.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userWithGroups)
}

// SignOut user
// swagger:route DELETE /users signOut
//
// Sign out user
//
// security:
//   oauth2:
//
// responses:
//   200:
//   401: Error
//   415: Error
func (h Handler) SignOut(c *gin.Context) {
	user, err := handler.GetUserFromContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := h.tokenService.SignOut(user.ID); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// FindById user
// swagger:route GET /users/{id} findUserById
//
// Return a user by id
//
// security:
//   oauth2:
//
// responses:
//   200: User
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

	userWithGroups, err := h.userService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userWithGroups)
}
