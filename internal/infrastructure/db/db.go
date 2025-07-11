package db

import (
	"boilerplate-one/internal/config"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func Connect(dbConfig config.DBConfig) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Name,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		sqlDB, err := conn.DB()
		if err != nil {
			log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
		}

		// Pooling settings
		sqlDB.SetMaxIdleConns(dbConfig.SetMaxIdleConns)                                    // idle connections
		sqlDB.SetMaxOpenConns(dbConfig.SetMaxOpenConns)                                    // max open connections
		sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.SetConnMaxLifetime) * time.Minute) // max lifetime per connection

		dbInstance = conn
	})

	return dbInstance
}
