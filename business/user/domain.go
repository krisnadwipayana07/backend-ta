package user

import (
	"context"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Username string
	Name     string
	Role     string
	Password string
	IsActive bool
}

type Usecase interface {
	LoginAdmin(ctx context.Context, username string, password string) (Domain, error)
}
type Repository interface {
	LoginAdmin(ctx context.Context, username string, password string) (Domain, error)
}
