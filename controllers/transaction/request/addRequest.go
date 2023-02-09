package request

import "snatia/business/transaction"

type AddRequest struct {
	CashierName string `json:"cashier_name"`
	Total       uint   `json:"total"`
}

type FilterDay struct {
}

func (req AddRequest) ToDomain() transaction.Domain {
	return transaction.Domain{
		CashierName: req.CashierName,
		Total:       req.Total,
	}
}
