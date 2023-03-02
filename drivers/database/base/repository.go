package base

import (
	"context"
	"errors"
	"log"
	"snatia/business/base"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
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

func (rep *MysqlBaseRepository) AnalyticsOLAP(ctx context.Context, insertData Activity) {
	_, err := rep.Coll.InsertOne(ctx, insertData)
	if err != nil {
		log.Println(err)
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
		newCtx := context.Background()
		insertData := newProduct.ToActivity()
		go rep.AnalyticsOLAP(newCtx, insertData)
	}

	return newProduct.ToDomain(), nil
}
func (rep *MysqlBaseRepository) GetDataWithoutConcurrency(ctx context.Context, id uint) (base.Domain, error) {
	newProduct := Products{}

	result := rep.DB.Table("products").Where("id = ?", id).Scan(&newProduct)
	if result.RowsAffected == 0 {
		return base.Domain{}, errors.New("Product Not Found")
	}

	analytics := viper.GetBool("analytics")
	if analytics {
		insertData := newProduct.ToActivity()
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
func (rep *MysqlBaseRepository) GetPageVisitGraph(ctx context.Context, startDate time.Time, endDate time.Time) ([]string, []int32, error) {
	// var allData []Activity
	// start := primitive.NewObjectIDFromTimestamp(startDate)
	// end := primitive.NewObjectIDFromTimestamp(endDate)

	// groupStage := bson.D{{"$group", bson.D{
	// 	{"_id", bson.D{{"product", "$product"}, {"date", "$product"}}},
	// 	{"count", bson.D{{"$sum", 1}}},
	// }}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$product"},
		{"count", bson.D{{"$sum", 1}}},
	}}}
	// matchDate := bson.D{{"$match", bson.D{
	// 	{"date", bson.D{
	// 		{"$gte", primitive.NewObjectIDFromTimestamp(startDate)},
	// 		{"$lte", primitive.NewObjectIDFromTimestamp(endDate)},
	// 	}},
	// }}}

	result, err := rep.Coll.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	log.Println(startDate)
	log.Println(endDate)
	log.Println(results)

	var label []string
	var value []int32

	for _, item := range results {
		// product := item["_id"]

		// fmt.Println(item["_id"].(string), item["count"].(int32))
		label = append(label, item["_id"].(string))
		value = append(value, item["count"].(int32))
	}
	return label, value, nil
}

// func (rep *MysqlBaseRepository) BuyProduct(ctx context.Context, domain base.Domain) (base.Domain, error) {

// }
