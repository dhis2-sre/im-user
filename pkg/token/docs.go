package token

import "github.com/lestrrat-go/jwx/jwk"

// swagger:response Error
type _ struct {
	//in: body
	_ string
}

// swagger:response Jwks
type _ struct {
	//in: body
	_ jwk.Key
}
