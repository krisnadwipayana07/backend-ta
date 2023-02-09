package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Name     string
	Role     string
	Password string
	IsActive bool
}
