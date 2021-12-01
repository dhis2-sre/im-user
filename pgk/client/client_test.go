package client

import (
	"fmt"
	"github.com/dhis2-sre/im-users/internal/di"
	"github.com/dhis2-sre/im-users/internal/server"
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
	c := ProvideClient(host)

	u, err := c.FindUserById(1)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), u.ID)
	assert.Equal(t, environment.Config.AdminUser.Email, u.Email)
}
