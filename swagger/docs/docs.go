package docs

// swagger:parameters getUserById getGroupById
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
