package models

import (
	"gorm.io/gorm"
)

type AuthUser struct {
	gorm.Model

	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
