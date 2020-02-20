package gokong

type RbacRequest struct {
	Name      string                 `json:"name" yaml:"name"`
	Comment   string                 `json:"comment,omitempty" yaml:"comment,omitempty"`
	Config    map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	CreatedAt *int                   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
}

// Role
// RbacEndpoint
// RoleEndpointPermission
// RoleEntityPermission
// RolePermission
