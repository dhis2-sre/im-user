package docs

import (
	"github.com/dhis2-sre/im-user/pkg/group"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/lestrrat-go/jwx/jwk"
)

// swagger:parameters findUserById findGroupById
type IdParam struct {
	// in: path
	// required: true
	ID uint `json:"id"`
}

// swagger:response
type Error struct {
	// The error message
	//in: body
	Message string
}

// swagger:response
type Jwks struct {
	//in: body
	Key jwk.Key
}

// swagger:parameters signUp
type _ struct {
	// SignUp request body parameter
	// in: body
	// required: true
	Body user.SignUpRequest
}

// swagger:parameters refreshToken
type _ struct {
	// Refresh token request body parameter
	// in: body
	// required: true
	Body user.RefreshTokenRequest
}

// swagger:parameters groupCreate
type _ struct {
	// Refresh token request body parameter
	// in: body
	// required: true
	Body group.CreateGroupRequest
}

// swagger:parameters groupAddUserToGroup groupAddClusterConfigurationToGroup
type GroupIdParameter struct {
	// in: path
	// required: true
	GroupID uint `json:"groupId"`
}

// swagger:parameters groupAddUserToGroup
type UserIdParameter struct {
	// in: path
	// required: true
	UserID uint `json:"userId"`
}

// swagger:parameters groupAddClusterConfigurationToGroup
type _ struct {
	// SOPS encrypted Kubernetes configuration file
	// in: formData
	// required: true
	// swagger:file
	Body group.CreateClusterConfigurationRequest
}

// swagger:parameters groupNameToId
type NameParameter struct {
	// in: path
	// required: true
	Name uint `json:"name"`
}
