package group

// swagger:parameters groupCreate
type _ struct {
	// Refresh token request body parameter
	// in: body
	// required: true
	_ CreateGroupRequest
}

// swagger:parameters addUserToGroup
type _ struct {
	// in: path
	// required: true
	GroupName string `json:"groupName"`

	// in: path
	// required: true
	UserID uint `json:"userId"`
}

// swagger:parameters addClusterConfigurationToGroup
type _ struct {
	// in: path
	// required: true
	Name string `json:"name"`

	// SOPS encrypted Kubernetes configuration file
	// in: formData
	// required: true
	// swagger:file
	_ CreateClusterConfigurationRequest
}

// swagger:parameters findGroupByName
type _ struct {
	// in: path
	// required: true
	Name string `json:"name"`
}
