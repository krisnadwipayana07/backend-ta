package transaction

import (
	"context"
	"errors"
	"fmt"
	"snatia/business/transaction"
	"time"

	"gorm.io/gorm"
)

type MySqlTransactionRepository struct {
	DB *gorm.DB
}

func NewMysqlTransactionRepository(db *gorm.DB) transaction.Repository {
	return &MySqlTransactionRepository{
		DB: db,
	}
}

func (rep *MySqlTransactionRepository) GetAllTransaction(ctx context.Context) ([]transaction.Domain, error) {
	var allData []Transaction

	result := rep.DB.Find(&allData)
	if result.Error != nil {
		return []transaction.Domain{}, errors.New("Get All Data Error")
	}
	if result.RowsAffected == 0 {
		return []transaction.Domain{}, errors.New("Transaction Empty")
	}

	res := ToDomainList(allData)
	return res, nil
}
func (rep *MySqlTransactionRepository) GetTotalByCashier(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error) {
	var allData []Transaction

	var label []string
	var value []uint

	comand := fmt.Sprintf("DATE(created_at) BETWEEN DATE('%d-%d-%d') AND DATE('%d-%d-%d')", startDate.Year(), startDate.Month(), startDate.Day(), endDate.Year(), endDate.Month(), endDate.Day())
	// fmt.Println(comand)

	result := rep.DB.Select("cashier_name, SUM(total) as total").Group("cashier_name").Where(comand).Find(&allData)
	if result.Error != nil {
		return []string{}, []uint{}, errors.New("Get All Data Error")
	}
	if result.RowsAffected == 0 {
		return []string{}, []uint{}, errors.New("Transaction Empty")
	}

	for _, item := range allData {
		label = append(label, item.CashierName)
		value = append(value, item.Total)
	}
	return label, value, nil
}
func (rep *MySqlTransactionRepository) GetProductSales(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error) {
	var allData []TransactionDetail

	var label []string
	var value []uint

	comand := fmt.Sprintf("DATE(`Transaction`.`created_at`) BETWEEN DATE('%d-%d-%d') AND DATE('%d-%d-%d')", startDate.Year(), startDate.Month(), startDate.Day(), endDate.Year(), endDate.Month(), endDate.Day())

	result := rep.DB.Select("SUM(qty) as qty").Joins("Product").Joins("Transaction").Group("product_id").Where(comand).Find(&allData)
	if result.Error != nil {
		return []string{}, []uint{}, errors.New("Get All Data Error")
	}
	if result.RowsAffected == 0 {
		return []string{}, []uint{}, errors.New("Transaction Empty")
	}

	for _, item := range allData {
		label = append(label, item.Product.Product)
		value = append(value, item.Qty)
	}
	return label, value, nil
}
func (rep *MySqlTransactionRepository) AddTransaction(ctx context.Context, domain transaction.Domain) (transaction.Domain, error) {
	insertData := FromDomain(domain)

	result := rep.DB.Create(&insertData)
	if result.Error != nil {
		return transaction.Domain{}, errors.New("Inset Data Error")
	}
	data := insertData.ToDomain()
	return data, nil
}
func (rep *MySqlTransactionRepository) GetSalesByDay(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []uint, error) {
	var GraphData []Transaction

	var label []string
	var value []uint

	comand := fmt.Sprintf("DATE(created_at) BETWEEN DATE('%d-%d-%d') AND DATE('%d-%d-%d')", startDate.Year(), startDate.Month(), startDate.Day(), endDate.Year(), endDate.Month(), endDate.Day())

	result := rep.DB.Select("count(id) as id, CAST(created_at as DATE) as cashier_name").Where(comand).Group("CAST(created_at as DATE)").Find(&GraphData)
	if result.Error != nil {
		return label, value, errors.New("Get All Data Error")
	}

	for _, item := range GraphData {
		label = append(label, item.CashierName[0:10])
		value = append(value, item.ID)
	}
	// log.Println(label)
	return label, value, nil
}
