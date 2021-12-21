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
// swagger:route GET /jwks Jwks
// Return a JWKS containing the public key which can be used to validate the JWT's dispensed at /signin
// responses:
//   200: Jwks
//   415: Error
//   500: Error
func (h *Handler) Jwks(c *gin.Context) {
	jwks, err := helper.CreateJwks(h.publicKey)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, jwks)
}
