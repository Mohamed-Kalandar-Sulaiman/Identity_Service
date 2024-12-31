package models

import "gorm.io/gorm"

type Application struct {
    gorm.Model
    AppName      string `gorm:"not null"`
    PublicKey    string `gorm:"type:text"`
    Context      string `gorm:"type:jsonb"` // JSONB field to store app-specific context data
}
