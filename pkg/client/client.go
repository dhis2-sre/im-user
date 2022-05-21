package client

import (
	"context"
	"github.com/dhis2-sre/im-user/swagger/sdk/client/operations"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	FindUserById(token string, id uint) (*models.User, error)
	FindGroupByName(token string, name string) (*models.Group, error)
	SignIn(username, password string) (*models.Tokens, error)
	Me(token string) (*models.User, error)
}

func ProvideClient(host string, basePath string) Client {
	transport := httptransport.New(host, basePath, nil)
	userService := operations.New(transport, strfmt.Default)
	return &cli{userService}
}

type cli struct {
	clientService operations.ClientService
}

func (c cli) FindUserById(token string, id uint) (*models.User, error) {
	params := &operations.FindUserByIDParams{ID: uint64(id), Context: context.Background()}
	clientAuthInfoWriter := httptransport.BearerToken(token)
	userByID, err := c.clientService.FindUserByID(params, clientAuthInfoWriter)
	if err != nil {
		return nil, err
	}
	return userByID.GetPayload(), nil
}

func (c cli) FindGroupByName(token string, name string) (*models.Group, error) {
	params := &operations.FindGroupByNameParams{Name: name, Context: context.Background()}
	clientAuthInfoWriter := httptransport.BearerToken(token)
	group, err := c.clientService.FindGroupByName(params, clientAuthInfoWriter)
	if err != nil {
		return nil, err
	}
	return group.GetPayload(), nil
}

func (c cli) SignIn(username, password string) (*models.Tokens, error) {
	clientAuthInfoWriter := httptransport.BasicAuth(username, password)
	params := operations.NewSignInParamsWithContext(context.Background())
	response, err := c.clientService.SignIn(params, clientAuthInfoWriter)
	return response.Payload, err
}

func (c cli) Me(token string) (*models.User, error) {
	clientAuthInfoWriter := httptransport.BearerToken(token)
	params := operations.NewMeParams().WithDefaults()
	response, err := c.clientService.Me(params, clientAuthInfoWriter)
	return response.Payload, err
}
