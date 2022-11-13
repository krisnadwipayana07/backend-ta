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
func (bu baseUsecase) GetAllData(ctx context.Context) ([]Domain, error) {
	res, err := bu.repo.GetAllData(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
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
