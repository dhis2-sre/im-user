package client

import (
	"context"
	"github.com/dhis2-sre/im-users/swagger/client/client"
	"github.com/dhis2-sre/im-users/swagger/client/client/public"
	"github.com/dhis2-sre/im-users/swagger/client/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	FindUserById(id uint) (*models.DtoUser, error)
	FindGroupById(id uint) (*models.DtoGroup, error)
}

func ProvideClient(host string) Client {
	transport := httptransport.New(host, "", nil)
	userService := client.New(transport, strfmt.Default)
	return &cli{userService}
}

type cli struct {
	userService *client.InstanceManagerUserService
}

func (c cli) FindUserById(id uint) (*models.DtoUser, error) {
	params := &public.GetFindbyidIDParams{ID: int64(id), Context: context.Background()}
	user, err := c.userService.Public.GetFindbyidID(params)
	if err != nil {
		return nil, err
	}
	return user.GetPayload(), nil
}

func (c cli) FindGroupById(id uint) (*models.DtoGroup, error) {
	params := &public.GetGroupsIDParams{ID: int64(id), Context: context.Background()}
	group, err := c.userService.Public.GetGroupsID(params)
	if err != nil {
		return nil, err
	}
	return group.GetPayload(), nil
}
