package base

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Product     string
	Pic         string
	Price       uint
	Description string
}

type Usecase interface {
	GetAllData(ctx context.Context) ([]Domain, error)
	GetData(ctx context.Context, id uint) (Domain, error)
	GetDataWithoutConcurrency(ctx context.Context, id uint) (Domain, error)
	GetPageVisitGraph(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []int32, error)
	// BuyProduct(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetAllData(ctx context.Context) ([]Domain, error)
	GetData(ctx context.Context, id uint) (Domain, error)
	GetDataWithoutConcurrency(ctx context.Context, id uint) (Domain, error)
	GetPageVisitGraph(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []int32, error)

	// BuyProduct(ctx context.Context, domain Domain) (Domain, error)
}
