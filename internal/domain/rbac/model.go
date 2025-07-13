package rbac

const (
	RoleAdmin   = 1
	RoleTeacher = 2
	RoleStudent = 3
)

var RoleNames = map[uint]string{
	RoleAdmin:   "admin",
	RoleTeacher: "teacher",
	RoleStudent: "student",
}

// Role struct
type Role struct {
	ID   uint
	Name string
}

// Permission struct
type Permission struct {
	ID   uint
	Name string
}

// RolePermission struct (many-to-many)
type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

// Example permissions
const (
	PermUserCreate = 1
	PermUserRead   = 2
	PermUserUpdate = 3
	PermUserDelete = 4
)

var PermissionNames = map[uint]string{
	PermUserCreate: "user.create",
	PermUserRead:   "user.read",
	PermUserUpdate: "user.update",
	PermUserDelete: "user.delete",
}

// Example role-permission mapping (static, for demo)
var RolePermissions = map[uint][]uint{
	RoleAdmin:   {PermUserCreate, PermUserRead, PermUserUpdate, PermUserDelete},
	RoleTeacher: {PermUserRead},
	RoleStudent: {},
}
