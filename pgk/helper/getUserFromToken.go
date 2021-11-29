package helper

import (
	"errors"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/model"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"gorm.io/gorm"
	"strings"
)

func GetUserFromToken(config config.Config, c *gin.Context) (*model.User, error) {
	authorizationHeader := c.GetHeader("Authorization")

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		authorizationHeader = strings.TrimPrefix(authorizationHeader, "Bearer ")
	}

	key, err := config.Authentication.Keys.GetPublicKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(
		[]byte(authorizationHeader),
		jwt.WithValidate(true),
		jwt.WithVerify(jwa.RS256, key),
	)
	if err != nil {
		return nil, err
	}

	userData, ok := token.Get("user")
	if !ok {
		return nil, errors.New("user not found in claims")
	}

	userMap, ok := userData.(map[string]interface{})
	if !ok {
		return nil, errors.New("failed to parse user data")
	}

	id := userMap["ID"].(float64)
	email := userMap["Email"].(string)

	return &model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Email: email,
	}, nil
}
