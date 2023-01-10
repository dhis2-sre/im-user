package token

import (
	"strings"
	"testing"
	"time"

	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_tokenService_GetTokens_Happy(t *testing.T) {
	id := uint(1)
	email := "someone@something.org"
	password := "passwordpasswordpasswordpassword"

	c, err := config.New()
	assert.NoError(t, err)

	repository := &mockTokenRepository{}
	repository.
		On("setRefreshToken", id, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).
		Return(nil)
	service, err := NewService(c, repository)
	assert.NoError(t, err)

	user := &model.User{
		Model:    gorm.Model{ID: id},
		Email:    email,
		Password: password,
	}

	tokens, err := service.GetTokens(user, "")
	assert.NoError(t, err)

	assert.True(t, strings.HasPrefix(tokens.AccessToken, "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9."))
	// TODO: Assert more
	repository.AssertExpectations(t)
}

func Test_tokenService_ValidateAccessToken(t *testing.T) {
	c, err := config.New()
	assert.NoError(t, err)

	repository := &mockTokenRepository{}
	service, err := NewService(c, repository)
	assert.NoError(t, err)
	accessToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3OTYzODQwNzcsImlhdCI6MTY0MDY5MTQ3NywidXNlciI6eyJJRCI6NCwiQ3JlYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiVXBkYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiRGVsZXRlZEF0IjpudWxsLCJFbWFpbCI6InNvbWVvbmVAc29tZXRoaW5nLm9yZyIsIkdyb3VwcyI6bnVsbCwiQWRtaW5Hcm91cHMiOm51bGx9fQ.FnPIu36kV1T-Jix5Wy-HsZeqxQI6Q_7HQ14C1DWKHETIBSk-vLQ_sCMHVPKA42utEDFI3Xpmf6Gyzv9aPU_Cvg-JDazRprfrZBqn4LzSmT6K3HGoKoQ0b5G8exxz0Ote8NQDB1NBZmYvD1gpVVisCvzaewJTRAvRA3DS0n_O4kU5QENdLNPfWFo0rXOC83sLBsEIe2Ce4TiRrepOCSQKE-_rwQQSA3w30MhFmhAY7Ozcd9i69mtfcvqjORdNJ-zREgiw8B2g9oh7byE1h2oxjvoKC3WRfPeSYoRY6GuMHSSWJdzFKIswlZHdWU1GicPJASBbkKGbP5n5O6FXyeo0bw"

	user, err := service.ValidateAccessToken(accessToken)
	assert.NoError(t, err)

	assert.Equal(t, uint(4), user.ID)
	assert.Equal(t, "someone@something.org", user.Email)
}

func Test_tokenService_ValidateRefreshToken(t *testing.T) {
	id := uint(1)

	c, err := config.New()
	assert.NoError(t, err)

	repository := &mockTokenRepository{}
	service, err := NewService(c, repository)
	assert.NoError(t, err)

	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ3Nzc0MzUsImlhdCI6MTY3MzI0MTQzNSwianRpIjoiODRmYmIyOTMtMDQ5YS00ZjBkLWFkNDYtY2Q3YmJjZDk3MDg0IiwidXNlcklkIjoxfQ.Fm90u_IFAGuHd8ePDWuYo-t9M-CSRdHUNvDjo9EZRTU"
	token, err := service.ValidateRefreshToken(refreshToken)
	assert.NoError(t, err)

	assert.Equal(t, id, token.UserId)
}

func Test_tokenService_SignOut(t *testing.T) {
	id := uint(1)

	c, err := config.New()
	assert.NoError(t, err)

	repository := &mockTokenRepository{}
	repository.
		On("deleteRefreshTokens", id).
		Return(nil)
	service, err := NewService(c, repository)
	assert.NoError(t, err)

	err = service.SignOut(id)
	assert.NoError(t, err)
}

type mockTokenRepository struct{ mock.Mock }

func (m *mockTokenRepository) deleteRefreshToken(userId uint, previousTokenId string) error {
	called := m.Called(userId, previousTokenId)
	return called.Error(0)
}

func (m *mockTokenRepository) setRefreshToken(userId uint, tokenId string, expiresIn time.Duration) error {
	called := m.Called(userId, tokenId, expiresIn)
	return called.Error(0)
}

func (m *mockTokenRepository) deleteRefreshTokens(userId uint) error {
	called := m.Called(userId)
	return called.Error(0)
}
