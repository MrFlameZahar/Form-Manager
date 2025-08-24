package model

import "fmt"

type Permission string

const (
	PermissionRead   Permission = "read"
	PermissionEdit   Permission = "edit"
	PermissionDelete Permission = "delete"
)

func NewPermissions(permission []string) ([]Permission, error) {
	upermissions := make([]Permission, len(permission))
	if len(permission) == 0 {
		return []Permission{}, fmt.Errorf("no permissions provided")
	}
	for _, p := range permission {
		switch p {
		case string(PermissionRead), string(PermissionEdit), string(PermissionDelete):
			continue
		default:
			return []Permission{}, fmt.Errorf("invalid permission: %s", p)
		}
	}
	for i, p := range permission {
		upermissions[i] = Permission(p)
	}
	return upermissions, nil
}

type RoleName string

const (
	RoleAdmin   RoleName = "admin"
	RoleCreator RoleName = "creator"
	RoleGuest   RoleName = "guest"
)

func NewRoleName(roleName string) (RoleName, error) {
	switch roleName {
	case string(RoleAdmin), string(RoleCreator), string(RoleGuest):
		return RoleName(roleName), nil
	default:
		return RoleName(""), fmt.Errorf("invalid role: %v", roleName)
	}
}

type Role struct {
	Name        RoleName
	permissions []Permission
}

func NewRole(roleName string, permissions []string) (Role, error) {
	role, err := NewRoleName(roleName)
	if err != nil {
		return Role{}, err
	}
	perms, err := NewPermissions(permissions)
	if err != nil {
		return Role{}, err
	}
	return Role{Name: role, permissions: perms}, nil
}

func (r *Role) HasPermission(permission Permission) bool {
	for _, p := range r.permissions {
		if p == permission {
			return true
		}
	}
	return false
}
