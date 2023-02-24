package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dhis2-sre/im-user/pkg/model"

	"github.com/dhis2-sre/im-user/internal/errdef"

	"github.com/dhis2-sre/im-user/internal/handler"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewHandler(config config.Config, userService userService, tokenService tokenService) Handler {
	return Handler{
		config,
		userService,
		tokenService,
	}
}

type Handler struct {
	config       config.Config
	userService  userService
	tokenService tokenService
}

type userService interface {
	SignUp(email string, password string) (*model.User, error)
	SignIn(email string, password string) (*model.User, error)
	FindById(id uint) (*model.User, error)
}

type tokenService interface {
	GetTokens(user *model.User, previousTokenId string) (*token.Tokens, error)
	ValidateRefreshToken(tokenString string) (*token.RefreshTokenData, error)
	SignOut(userId uint) error
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=16,lte=128"`
}

// SignUp user
func (h *Handler) SignUp(c *gin.Context) {
	// swagger:route POST /users signUp
	//
	// SignUp user
	//
	// Sign up a user. This endpoint is publicly accessible and therefor anyone can sign up. However, before being able to perform any actions, users needs to be a member of a group. And only administrators can add users to groups.
	//
	// responses:
	//   201: User
	//   400: Error
	//   415: Error
	var request SignUpRequest

	if err := handler.DataBinder(c, &request); err != nil {
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
func (h *Handler) SignIn(c *gin.Context) {
	// swagger:route POST /tokens signIn
	//
	// Sign in
	//
	// Sign in... And get tokens
	//
	// security:
	//   basicAuth:
	//
	// responses:
	//   201: Tokens
	//   401: Error
	//   403: Error
	//   404: Error
	//   415: Error
	user, err := handler.GetUserFromContext(c)
	if err != nil {
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
func (h Handler) RefreshToken(c *gin.Context) {
	// swagger:route POST /refresh refreshToken
	//
	// Refresh tokens
	//
	// Refresh user tokens
	//
	// responses:
	//   201: Tokens
	//   400: Error
	//   415: Error
	var request RefreshTokenRequest

	if err := handler.DataBinder(c, &request); err != nil {
		return
	}

	refreshToken, err := h.tokenService.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	user, err := h.userService.FindById(refreshToken.UserId)
	if err != nil {
		if errdef.IsNotFound(err) {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
		} else {
			_ = c.Error(err)
		}
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
func (h Handler) Me(c *gin.Context) {
	// swagger:route GET /me me
	//
	// User details
	//
	// Current user details
	//
	// security:
	//   oauth2:
	//
	// responses:
	//   200: User
	//   401: Error
	//   403: Error
	//   404: Error
	//   415: Error
	user, err := handler.GetUserFromContext(c)
	if err != nil {
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
func (h Handler) SignOut(c *gin.Context) {
	// swagger:route DELETE /users signOut
	//
	// Sign out
	//
	// Sign out user... The authentication is done using oauth and JWT. A JWT can't easily be invalidated so even after calling this endpoint a user can still sign in assuming the JWT isn't expired. However, the token can't be refreshed using the refresh token supplied upon signin
	//
	// security:
	//	oauth2:
	//
	// responses:
	//	200:
	//	401: Error
	//	415: Error
	user, err := handler.GetUserFromContext(c)
	if err != nil {
		return
	}

	if err := h.tokenService.SignOut(user.ID); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// FindById user
func (h Handler) FindById(c *gin.Context) {
	// swagger:route GET /users/{id} findUserById
	//
	// Find user
	//
	// Find a user by its id
	//
	// security:
	//	oauth2:
	//
	// responses:
	//	200: User
	//	401: Error
	//	403: Error
	//	404: Error
	//	415: Error
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("error parsing id"))
		return
	}

	userWithGroups, err := h.userService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userWithGroups)
}
