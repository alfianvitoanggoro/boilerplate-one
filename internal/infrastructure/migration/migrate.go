package migration

import (
	"boilerplate-one/internal/domain/rbac"
	"boilerplate-one/internal/domain/user"
	"boilerplate-one/pkg/logger"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{}, // tambahkan entity lainnya di sini
		&rbac.Role{},
		&rbac.Permission{},
		&rbac.RolePermission{},
	)

	if err != nil {
		logger.Errorf("❌ Migration failed: %v", err)
	}

	logger.Infof("✅ Migration completed successfully.")
}
