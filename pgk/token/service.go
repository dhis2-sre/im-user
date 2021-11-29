package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/model"
	"github.com/dhis2-sre/im-users/pgk/token/helper"
	"github.com/gofrs/uuid"
	"log"
)

type Service interface {
	GetTokens(user *model.User, previousTokenId string) (*Tokens, error)
	ValidateAccessToken(tokenString string) (*model.User, error)
	ValidateRefreshToken(tokenString string) (*RefreshTokenData, error)
	SignOut(userId uint) error
}

func ProvideTokenService(
	c config.Config,
	tokenRepository Repository,
) Service {
	privateKey, err := c.Authentication.Keys.GetPrivateKey()
	// TODO: Return error
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Return error
	publicKey, err := c.Authentication.Keys.GetPublicKey()
	if err != nil {
		log.Fatal(err)
	}

	return &tokenService{
		tokenRepository,
		privateKey,
		publicKey,
		c.Authentication.AccessTokenExpirationSeconds,
		c.Authentication.RefreshTokenSecretKey,
		c.Authentication.RefreshTokenExpirationSeconds,
	}
}

type Tokens struct {
	AccessToken    string    `json:"access_token"`
	TokenType      string    `json:"token_type"`
	RefreshToken   string    `json:"refresh_token"`
	ExpiresIn      uint      `json:"expires_in"`
	RefreshTokenId uuid.UUID `json:"jti"`
}

type RefreshTokenData struct {
	SignedToken string
	ID          uuid.UUID
	UserId      uint
}

type tokenService struct {
	tokenRepository               Repository
	privateKey                    *rsa.PrivateKey
	publicKey                     *rsa.PublicKey
	accessTokenExpirationSeconds  int
	refreshTokenSecretKey         string
	refreshTokenExpirationSeconds int
}

func (t tokenService) GetTokens(user *model.User, previousRefreshTokenId string) (*Tokens, error) {
	if previousRefreshTokenId != "" {
		if err := t.tokenRepository.DeleteRefreshToken(user.ID, previousRefreshTokenId); err != nil {
			message := fmt.Sprintf("Could not delete previous refreshToken for user.Id: %d, tokenId: %s\n", user.ID, previousRefreshTokenId)
			return nil, apperror.NewInternal(message)
		}
	}

	accessToken, err := helper.GenerateAccessToken(user, t.privateKey, t.accessTokenExpirationSeconds)
	if err != nil {
		message := fmt.Sprintf("Error generating accessToken for user: %+v\nError: %s", user, err)
		return nil, apperror.NewInternal(message)
	}

	refreshToken, err := helper.GenerateRefreshToken(user, t.refreshTokenSecretKey, t.refreshTokenExpirationSeconds)
	if err != nil {
		message := fmt.Sprintf("Error generating refreshToken for user: %+v\nError: %s", user, err)
		return nil, apperror.NewInternal(message)
	}

	if err := t.tokenRepository.SetRefreshToken(user.ID, refreshToken.TokenId.String(), refreshToken.ExpiresIn); err != nil {
		message := fmt.Sprintf("Error storing token: %d\nError: %s", user.ID, err)
		return nil, apperror.NewInternal(message)
	}

	return &Tokens{
		AccessToken:    accessToken,
		TokenType:      "bearer",
		RefreshToken:   refreshToken.SignedString,
		ExpiresIn:      uint(t.accessTokenExpirationSeconds),
		RefreshTokenId: refreshToken.TokenId,
	}, nil
}

func (t tokenService) ValidateAccessToken(tokenString string) (*model.User, error) {
	tokenClaims, err := helper.ValidateAccessToken(tokenString, t.publicKey)
	if err != nil {
		log.Printf("Unable to verify token: %s\n", err)
		return nil, apperror.NewUnauthorized("Unable to verify token")
	}

	return tokenClaims.User, nil
}

func (t tokenService) ValidateRefreshToken(tokenString string) (*RefreshTokenData, error) {
	claims, err := helper.ValidateRefreshToken(tokenString, t.refreshTokenSecretKey)

	if err != nil {
		log.Printf("Unable to validate token: %s\n%s\n", tokenString, err)
		return nil, apperror.NewUnauthorized("Unable to verify refresh token")
	}

	tokenId, err := uuid.FromString(claims.ID)
	if err != nil {
		log.Printf("Couldn't parse token id: %s\n%s\n", claims.ID, err)
		return nil, apperror.NewUnauthorized("Unable to verify refresh token")
	}

	return &RefreshTokenData{
		SignedToken: tokenString,
		ID:          tokenId,
		UserId:      claims.UserId,
	}, nil
}

func (t tokenService) SignOut(userId uint) error {
	return t.tokenRepository.DeleteRefreshTokens(userId)
}
