package client

import (
	"fmt"
	"github.com/dhis2-sre/im-user/internal/di"
	"github.com/dhis2-sre/im-user/internal/server"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFindUserById(t *testing.T) {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, environment.Config.BasePath)

	tokens, err := c.SignIn(environment.Config.AdminUser.Email, environment.Config.AdminUser.Password)
	assert.NoError(t, err)

	u, err := c.FindUserById(tokens.AccessToken, 1)
	assert.NoError(t, err)

	assert.Equal(t, uint64(1), u.ID)
	assert.Equal(t, environment.Config.AdminUser.Email, u.Email)
}

func TestFindGroupByName(t *testing.T) {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, environment.Config.BasePath)

	tokens, err := c.SignIn(environment.Config.AdminUser.Email, environment.Config.AdminUser.Password)
	assert.NoError(t, err)

	g, err := c.FindGroupByName(tokens.AccessToken, model.AdministratorGroupName)
	assert.NoError(t, err)

	assert.Equal(t, model.AdministratorGroupName, g.Name)
}

func TestSignIn(t *testing.T) {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, environment.Config.BasePath)

	tokens, err := c.SignIn(environment.Config.AdminUser.Email, environment.Config.AdminUser.Password)
	assert.NoError(t, err)

	assert.Equal(t, "bearer", tokens.TokenType)
	assert.True(t, tokens.AccessToken != "")
	assert.True(t, tokens.RefreshToken != "")
}

func TestMe(t *testing.T) {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, environment.Config.BasePath)

	tokens, err := c.SignIn(environment.Config.AdminUser.Email, environment.Config.AdminUser.Password)
	assert.NoError(t, err)

	me, err := c.Me(tokens.AccessToken)
	assert.NoError(t, err)

	assert.Equal(t, uint64(1), me.ID)
	assert.Equal(t, environment.Config.AdminUser.Email, me.Email)
}
