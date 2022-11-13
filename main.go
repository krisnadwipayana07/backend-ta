package main

import (
	"log"
	"net/http"
	"snatia/app/routes"
	BaseUsecase "snatia/business/base"
	BaseController "snatia/controllers/base"
	BaseRepo "snatia/drivers/database/base"
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

	analytics := viper.GetBool("analytics")
	var mong *mongo.Collection

	if analytics {
		configMongo := mongodb.ConfigDB{
			MongoURL: viper.GetString("database.mongodb.url"),
		}
		mong = configMongo.InitialDB()
	}

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

	routesInit := routes.RouterControllerList{
		BaseController: *BaseControllerInterface,
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
