package user

import (
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/internal/handler"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/token"
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
// @Summary User sign in
// @Description Classic basic http auth...
// @Tags Public
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /signin [post]
// @Security BasicAuthentication
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
// @Summary User details
// @Description Show user details
// @Tags Restricted
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /me [get]
// @Security OAuth2Password
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
// @Summary Sign out user
// @Description Delete refresh tokens...
// @Tags Restricted
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /signout [get]
// @Security OAuth2Password
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

	c.JSON(http.StatusOK, "Signed out successfully")
}

// FindById godoc
// @Summary User details by id
// @Description Show user details by id
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /findbyid [get]
// @Param id path string true "User id"
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
