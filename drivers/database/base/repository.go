package base

import (
	"context"
	"errors"
	"snatia/business/base"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type MysqlBaseRepository struct {
	DB   *gorm.DB
	Coll *mongo.Collection
}

func NewMysqlBaseRepository(db *gorm.DB, mong *mongo.Collection) base.Repository {
	return &MysqlBaseRepository{
		DB:   db,
		Coll: mong,
	}
}

func (rep *MysqlBaseRepository) GetData(ctx context.Context, id uint) (base.Domain, error) {
	newProduct := Products{}

	result := rep.DB.Table("products").Where("id = ?", id).Scan(&newProduct)
	if result.RowsAffected == 0 {
		return base.Domain{}, errors.New("Product Not Found")
	}
	analytics := viper.GetBool("analytics")
	if analytics {
		insertData := Activity{Product: newProduct.Product, Date: time.Now()}
		_, err := rep.Coll.InsertOne(ctx, insertData)
		if err != nil {
			return base.Domain{}, err
		}
	}

	return newProduct.ToDomain(), nil
}
func (rep *MysqlBaseRepository) GetAllData(ctx context.Context) ([]base.Domain, error) {
	var allData []Products

	result := rep.DB.Find(&allData)
	if result.Error != nil {
		return []base.Domain{}, errors.New("Get All Data Error")
	}
	if result.RowsAffected == 0 {
		return []base.Domain{}, errors.New("Product Empty")
	}

	res := ToDomainList(allData)
	return res, nil
}

// func (rep *MysqlBaseRepository) BuyProduct(ctx context.Context, domain base.Domain) (base.Domain, error) {

// }
