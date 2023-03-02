package main

import (
	"log"
	"net/http"
	"snatia/app/routes"
	BaseUsecase "snatia/business/base"
	TransactionUsecase "snatia/business/transaction"
	BaseController "snatia/controllers/base"
	TransactionController "snatia/controllers/transaction"
	BaseRepo "snatia/drivers/database/base"
	TransactionRepo "snatia/drivers/database/transaction"
	UserRepo "snatia/drivers/database/user"
	"snatia/drivers/mongodb"
	"snatia/drivers/mysql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("This Services RUN on DEBUG Mode")
	}
}

type User struct {
	gorm.Model
	Name string
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&BaseRepo.Products{})
	db.AutoMigrate(&BaseRepo.Activitys{})
	db.AutoMigrate(&UserRepo.User{})
	db.AutoMigrate(&TransactionRepo.Transaction{})
	db.AutoMigrate(&TransactionRepo.TransactionDetail{})
}

func startMongo() {

}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString("database.mysql.user"),
		DB_Password: viper.GetString("database.mysql.pass"),
		DB_Host:     viper.GetString("database.mysql.host"),
		DB_Port:     viper.GetString("database.mysql.port"),
		DB_Database: viper.GetString("database.mysql.name"),
	}
	db := configDb.InitialDB()
	dbMigrate(db)

	var mong *mongo.Collection

	configMongo := mongodb.ConfigDB{
		MongoURL: viper.GetString("database.mongodb.url"),
	}
	mong = configMongo.InitialDB()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	//Base
	BaseRepoInterface := BaseRepo.NewMysqlBaseRepository(db, mong)
	BaseUsecaseInterface := BaseUsecase.NewBaseUsecase(timeoutContext, BaseRepoInterface)
	BaseControllerInterface := BaseController.NewBaseController(BaseUsecaseInterface)

	TransactionInterface := TransactionRepo.NewMysqlTransactionRepository(db)
	TransactionUsecaseInterface := TransactionUsecase.NewTransactionUsecase(timeoutContext, TransactionInterface)
	TransactionControllerInterface := TransactionController.NewTransactionController(TransactionUsecaseInterface)

	routesInit := routes.RouterControllerList{
		BaseController:        *BaseControllerInterface,
		TransactionController: *TransactionControllerInterface,
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
