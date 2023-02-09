package transaction

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	CashierName string
	Total       uint
}

type DetailDomain struct {
	gorm.Model
	TransactionID uint
	ProductID     uint
	Price         uint
	Qty           uint
	Total         uint
}

type UseCase interface {
	//graph
	GetTotalByCashier(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error)
	GetProductSales(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error)
	GetSalesByDay(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error)
	//transaction
	GetAllTransaction(ctx context.Context) ([]Domain, error)
	AddTransaction(ctx context.Context, domain Domain) (Domain, error)
	//tabel

}
type Repository interface {
	GetSalesByDay(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
	GetTotalByCashier(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error)
	GetProductSales(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error)
	AddTransaction(ctx context.Context, domain Domain) (Domain, error)
}
