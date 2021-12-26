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

// Signup godoc
// @Summary User sign in
// @Description Posting username (email) and password... And user is returned
// @Tags Public
// @Accept json
// @Produce json
// @Success 201 {object} dto.User
// @Failure 400 {object} map[string]interface{}
// @Router /signup [post]
// @Param signupRequest body SignupRequest true "Email and Password json object"
func (h *Handler) Signup(c *gin.Context) {
	var request SignupRequest

	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	user, err := h.userService.Signup(request.Email, request.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// SignIn godoc
// swagger:route POST /signin SignIn
// Return user tokens
// responses:
//   201: Tokens
//   403: Error
//   404: Error
//   415: Error
// security:
//   basic:
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

// RefreshToken godoc
// @Summary Refresh tokens
// @Description Post a refresh token and this endpoint will return a fresh set of tokens
// @Tags Public
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /refresh [post]
// @Param refreshTokenRequest body RefreshTokenRequest true "Refresh token request"
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

// Me godoc
// swagger:route GET /me Me
// Return user details
// responses:
//   200: User
//   403: Error
//   404: Error
//   415: Error
// security:
//   oauth2:
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

// SignOut godoc
// swagger:route GET /signout SignOut
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

	c.String(http.StatusOK, "Signed out successfully")
}

// FindById godoc
// swagger:route GET /findbyid/{id} FindUserById
// Return a user by id
// responses:
//   200: User
//   403: Error
//   404: Error
//   415: Error
// security:
//   oauth2:
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
