package base

import (
	"snatia/business/base"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Product string
}

type Activity struct {
	Product string
	Date    time.Time
}

func (product Products) ToDomain() base.Domain {
	return base.Domain{
		Product: product.Product,
		Model:   product.Model,
	}
}

func FromDomain(domain base.Domain) Products {
	return Products{
		Product: domain.Product,
		Model:   domain.Model,
	}
}

func ToDomainList(product []Products) []base.Domain {
	list := []base.Domain{}
	for _, v := range product {
		list = append(list, v.ToDomain())
	}
	return list
}
