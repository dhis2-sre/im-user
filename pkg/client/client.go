package client

import (
	"context"
	"github.com/dhis2-sre/im-users/swagger/sdk/client/operations"
	"github.com/dhis2-sre/im-users/swagger/sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	FindUserById(id uint) (*models.User, error)
	FindGroupById(id uint) (*models.Group, error)
}

func ProvideClient(host string, basePath string) Client {
	transport := httptransport.New(host, basePath, nil)
	userService := operations.New(transport, strfmt.Default)
	return &cli{userService}
}

type cli struct {
	clientService operations.ClientService
}

func (c cli) FindUserById(id uint) (*models.User, error) {
	params := &operations.GetUserByIDParams{ID: uint64(id), Context: context.Background()}
	userByID, err := c.clientService.GetUserByID(params)
	if err != nil {
		return nil, err
	}
	return userByID.GetPayload(), nil
}

func (c cli) FindGroupById(id uint) (*models.Group, error) {
	params := &operations.GetGroupByIDParams{ID: uint64(id), Context: context.Background()}
	group, err := c.clientService.GetGroupByID(params)
	if err != nil {
		return nil, err
	}
	return group.GetPayload(), nil
}
