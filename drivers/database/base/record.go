package base

import (
	"snatia/business/base"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Product     string
	Price       uint
	Pic         string
	Description string
}

type Activity struct {
	Product string
	Date    time.Time
}

func (product Products) ToDomain() base.Domain {
	return base.Domain{
		Product:     product.Product,
		Model:       product.Model,
		Pic:         product.Pic,
		Price:       product.Price,
		Description: product.Description,
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

func (product Products) ToActivity() Activity {
	return Activity{
		Product: product.Product,
		Date:    time.Now(),
	}
}
