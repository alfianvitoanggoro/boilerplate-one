package rbac

// Service RBAC untuk pengecekan permission

type RBACService struct{}

// HasPermission mengecek apakah role memiliki permission tertentu
func (s *RBACService) HasPermission(roleID uint, permissionID uint) bool {
	perms, ok := RolePermissions[roleID]
	if !ok {
		return false
	}
	for _, p := range perms {
		if p == permissionID {
			return true
		}
	}
	return false
}
