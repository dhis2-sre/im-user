package token

import (
	"crypto/rsa"
	"github.com/dhis2-sre/im-users/pkg/config"
	"github.com/dhis2-sre/im-users/pkg/token/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ProvideHandler(config config.Config) Handler {
	publicKey, err := config.Authentication.Keys.GetPublicKey()
	if err != nil {
		log.Fatalln(err)
	}

	return Handler{
		publicKey,
	}
}

type Handler struct {
	publicKey *rsa.PublicKey
}

// Jwks godoc
// @Summary Return jwks
// @Description Return public used to validate tokens in a jwks
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /signup [post]
func (h *Handler) Jwks(c *gin.Context) {
	jwks, err := helper.CreateJwks(h.publicKey)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, jwks)
}