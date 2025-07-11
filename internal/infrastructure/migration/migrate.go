package migration

import (
	"boilerplate-one/internal/domain/user"
	"boilerplate-one/pkg/logger"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{}, // tambahkan entity lainnya di sini
	)

	if err != nil {
		logger.Errorf("❌ Migration failed: %v", err)
	}

	logger.Infof("✅ Migration completed successfully.")
}
