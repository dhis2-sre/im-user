package client

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/dhis2-sre/im-user/internal/server"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/group"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/storage"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindUserById(t *testing.T) {
	cfg := config.New()
	r := engine(t, cfg)

	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, cfg.BasePath)

	tokens, err := c.SignIn(cfg.AdminUser.Email, cfg.AdminUser.Password)
	assert.NoError(t, err)

	u, err := c.FindUserById(tokens.AccessToken, 1)
	assert.NoError(t, err)

	assert.Equal(t, uint64(1), u.ID)
	assert.Equal(t, cfg.AdminUser.Email, u.Email)
}

func TestFindGroupByName(t *testing.T) {
	cfg := config.New()
	r := engine(t, cfg)

	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, cfg.BasePath)

	tokens, err := c.SignIn(cfg.AdminUser.Email, cfg.AdminUser.Password)
	assert.NoError(t, err)

	g, err := c.FindGroupByName(tokens.AccessToken, model.AdministratorGroupName)
	assert.NoError(t, err)

	assert.Equal(t, model.AdministratorGroupName, g.Name)
}

func TestSignIn(t *testing.T) {
	cfg := config.New()
	r := engine(t, cfg)

	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, cfg.BasePath)

	tokens, err := c.SignIn(cfg.AdminUser.Email, cfg.AdminUser.Password)
	assert.NoError(t, err)

	assert.Equal(t, "bearer", tokens.TokenType)
	assert.True(t, tokens.AccessToken != "")
	assert.True(t, tokens.RefreshToken != "")
}

func TestMe(t *testing.T) {
	cfg := config.New()
	r := engine(t, cfg)

	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, cfg.BasePath)

	tokens, err := c.SignIn(cfg.AdminUser.Email, cfg.AdminUser.Password)
	assert.NoError(t, err)

	me, err := c.Me(tokens.AccessToken)
	assert.NoError(t, err)

	assert.Equal(t, uint64(1), me.ID)
	assert.Equal(t, cfg.AdminUser.Email, me.Email)
}

func engine(t *testing.T, cfg config.Config) *gin.Engine {
	client := storage.ProvideRedis(cfg)
	repository := token.ProvideTokenRepository(client)
	tokenSvc := token.ProvideTokenService(cfg, repository)
	tokenHandler := token.ProvideHandler(cfg)

	db, err := storage.ProvideDatabase(cfg)
	require.NoError(t, err)

	usrRepository := user.ProvideRepository(db)
	usrSvc := user.ProvideService(usrRepository)
	usrHandler := user.ProvideHandler(cfg, usrSvc, tokenSvc)

	groupRepository := group.ProvideRepository(db)
	groupSvc := group.ProvideService(groupRepository, usrRepository)
	groupHandler := group.ProvideHandler(groupSvc, usrSvc)

	authenticationMiddleware := middleware.ProvideAuthentication(usrSvc, tokenSvc)
	authorizationMiddleware := middleware.ProvideAuthorization(usrSvc)

	return server.GetEngine(cfg, tokenHandler, usrHandler, groupHandler, authenticationMiddleware, authorizationMiddleware, usrSvc, groupSvc)
}
