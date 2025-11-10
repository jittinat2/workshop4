package entity

import (
	"time"

	"gorm.io/gorm"
)

// User entity: ID format LBK+6 digits, membership level, names, phone, email, membered_at, point
type User struct {
	ID         string         `gorm:"primaryKey;size:16" json:"id"`
	Level      string         `json:"level"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	Phone      string         `json:"phone"`
	Email      string         `json:"email"`
	MemberedAt *time.Time     `json:"membered_at"`
	Point      int            `json:"point"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
