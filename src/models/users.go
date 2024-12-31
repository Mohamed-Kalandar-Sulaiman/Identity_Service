package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Context  string `gorm:"type:jsonb"` // JSONB field to store user-specific context data
}
