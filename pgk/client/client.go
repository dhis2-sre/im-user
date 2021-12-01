package client

import (
	"context"
	"github.com/dhis2-sre/im-users/swagger/client/client/public"
	"github.com/dhis2-sre/im-users/swagger/client/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	FindUserById(id uint) (*models.DtoUser, error)
}

func ProvideClient(host string) Client {
	transport := httptransport.New(host, "", nil)
	c := public.New(transport, strfmt.Default)

	return &client{c}
}

type client struct {
	client public.ClientService
}

func (c client) FindUserById(id uint) (*models.DtoUser, error) {
	// TODO: Why int64?
	params := &public.GetFindbyidIDParams{ID: int64(id), Context: context.Background()}
	user, err := c.client.GetFindbyidID(params)
	if err != nil {
		return nil, err
	}
	return user.GetPayload(), nil
}
