package base

import (
	"context"
	"errors"
	"time"
)

type baseUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewBaseUsecase(timeout time.Duration, cr Repository) Usecase {
	return &baseUsecase{
		repo:           cr,
		contextTimeout: timeout,
	}
}

func (bu baseUsecase) GetData(ctx context.Context, id uint) (Domain, error) {
	if id == 0 {
		return Domain{}, errors.New("ID Empty")
	}

	res, err := bu.repo.GetData(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
func (bu baseUsecase) GetDataOLTP(ctx context.Context, id uint) (Domain, error) {
	if id == 0 {
		return Domain{}, errors.New("ID Empty")
	}

	res, err := bu.repo.GetDataOLTP(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
func (bu baseUsecase) GetDataWithoutConcurrency(ctx context.Context, id uint) (Domain, error) {
	if id == 0 {
		return Domain{}, errors.New("ID Empty")
	}

	res, err := bu.repo.GetDataWithoutConcurrency(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
func (bu baseUsecase) GetAllData(ctx context.Context) ([]Domain, error) {
	res, err := bu.repo.GetAllData(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}
func (bu baseUsecase) GetPageVisitGraph(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []int32, error) {
	title, data, err := bu.repo.GetPageVisitGraph(ctx, startDate, endDate)
	if err != nil {
		return []string{}, []int32{}, err
	}
	return title, data, nil
}
func (bu baseUsecase) GetPageVisitGraphOLTP(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []int32, error) {
	title, data, err := bu.repo.GetPageVisitGraphOLTP(ctx, startDate, endDate)
	if err != nil {
		return []string{}, []int32{}, err
	}
	return title, data, nil
}

// func (bu baseUsecase) BuyProduct(ctx context.Context, domain Domain) (Domain, error) {
// 	if domain.Product == "" {
// 		return Domain{}, errors.New("Name Product Empty")
// 	} else if domain.Quantity == 0 {
// 		return Domain{}, errors.New("Quantity Empty")
// 	}

// 	res, err := bu.repo.BuyProduct(ctx, domain)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return res, nil
// }
