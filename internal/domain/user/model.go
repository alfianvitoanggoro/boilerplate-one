package user

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(100)" json:"name"`
	Email     string `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
	RoleID    uint
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
