package transaction

import (
	"context"
	"time"
)

type transactionUseCase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewTransactionUsecase(timeout time.Duration, cr Repository) UseCase {
	return &transactionUseCase{
		repo:           cr,
		contextTimeout: timeout,
	}
}

func (tu transactionUseCase) GetAllTransaction(ctx context.Context) ([]Domain, error) {
	res, err := tu.repo.GetAllTransaction(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}
func (tu transactionUseCase) GetTotalByCashier(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error) {
	label, value, err := tu.repo.GetTotalByCashier(ctx, startDate, endDate)
	if err != nil {
		return []string{}, []uint{}, err
	}
	return label, value, nil
}
func (tu transactionUseCase) GetProductSales(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error) {
	label, value, err := tu.repo.GetProductSales(ctx, startDate, endDate)
	if err != nil {
		return []string{}, []uint{}, err
	}
	return label, value, nil
}
func (tu transactionUseCase) AddTransaction(ctx context.Context, domain Domain) (Domain, error) {
	data, err := tu.repo.AddTransaction(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}
func (tu transactionUseCase) GetSalesByDay(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error) {
	label, value, err := tu.repo.GetSalesByDay(ctx, startDate, endDate)
	if err != nil {
		return []string{}, []uint{}, err
	}
	return label, value, nil
}
