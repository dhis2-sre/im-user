package client

import (
	"context"

	"github.com/dhis2-sre/im-user/swagger/sdk/client/operations"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type userClient struct {
	client operations.ClientService
}

func New(host string, basePath string) *userClient {
	transport := httptransport.New(host, basePath, nil)
	return &userClient{
		client: operations.New(transport, strfmt.Default),
	}
}

func (c userClient) FindUserById(token string, id uint) (*models.User, error) {
	params := &operations.FindUserByIDParams{ID: uint64(id), Context: context.Background()}
	clientAuthInfoWriter := httptransport.BearerToken(token)
	userByID, err := c.client.FindUserByID(params, clientAuthInfoWriter)
	if err != nil {
		return nil, err
	}
	return userByID.GetPayload(), nil
}

func (c userClient) FindGroupByName(token string, name string) (*models.Group, error) {
	params := &operations.FindGroupByNameParams{Name: name, Context: context.Background()}
	clientAuthInfoWriter := httptransport.BearerToken(token)
	group, err := c.client.FindGroupByName(params, clientAuthInfoWriter)
	if err != nil {
		return nil, err
	}
	return group.GetPayload(), nil
}

func (c userClient) SignIn(username, password string) (*models.Tokens, error) {
	clientAuthInfoWriter := httptransport.BasicAuth(username, password)
	params := operations.NewSignInParamsWithContext(context.Background())
	response, err := c.client.SignIn(params, clientAuthInfoWriter)
	if err != nil {
		return nil, err
	}
	return response.Payload, err
}

func (c userClient) Me(token string) (*models.User, error) {
	clientAuthInfoWriter := httptransport.BearerToken(token)
	params := operations.NewMeParams().WithDefaults()
	response, err := c.client.Me(params, clientAuthInfoWriter)
	if err != nil {
		return nil, err
	}
	return response.Payload, err
}
