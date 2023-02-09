package transaction

import (
	"snatia/business/transaction"
	"snatia/drivers/database/base"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionID uint
	ProductID     uint
	Price         uint
	Qty           uint
	Total         uint

	Product     base.Products `gorm:"foreignKey:ProductID"`
	Transaction Transaction   `gorm:"foreignKey:TransactionID"`
}

type Transaction struct {
	gorm.Model
	CashierName string
	Total       uint
}

func (trx Transaction) ToDomain() transaction.Domain {
	return transaction.Domain{
		Model:       trx.Model,
		CashierName: trx.CashierName,
		Total:       trx.Total,
	}
}
func FromDomain(domain transaction.Domain) Transaction {
	return Transaction{
		Model:       domain.Model,
		CashierName: domain.CashierName,
		Total:       domain.Total,
	}
}
func ToDomainList(trx []Transaction) []transaction.Domain {
	list := []transaction.Domain{}
	for _, v := range trx {
		list = append(list, v.ToDomain())
	}
	return list
}
