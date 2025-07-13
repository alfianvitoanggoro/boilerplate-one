package rbac

import "errors"

// RoleRepository interface
 type RoleRepository interface {
	GetRoleByID(id uint) (*Role, error)
}

// PermissionRepository interface
 type PermissionRepository interface {
	GetPermissionByID(id uint) (*Permission, error)
}

// Static implementation (for demo)
type StaticRoleRepo struct{}

type StaticPermissionRepo struct{}

func (r *StaticRoleRepo) GetRoleByID(id uint) (*Role, error) {
	name, ok := RoleNames[id]
	if !ok {
		return nil, errors.New("role not found")
	}
	return &Role{ID: id, Name: name}, nil
}

func (p *StaticPermissionRepo) GetPermissionByID(id uint) (*Permission, error) {
	name, ok := PermissionNames[id]
	if !ok {
		return nil, errors.New("permission not found")
	}
	return &Permission{ID: id, Name: name}, nil
}
