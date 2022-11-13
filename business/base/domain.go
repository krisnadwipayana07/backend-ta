package base

import (
	"context"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Product string
}

type Usecase interface {
	GetAllData(ctx context.Context) ([]Domain, error)
	GetData(ctx context.Context, id uint) (Domain, error)
	// BuyProduct(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetAllData(ctx context.Context) ([]Domain, error)
	GetData(ctx context.Context, id uint) (Domain, error)
	// BuyProduct(ctx context.Context, domain Domain) (Domain, error)
}
