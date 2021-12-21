package docs

import "github.com/lestrrat-go/jwx/jwk"

// swagger:parameters FindUserById FindGroupById
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
