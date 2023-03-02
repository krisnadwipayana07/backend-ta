package routes

import (
	"net/http"
	"snatia/controllers/base"
	"snatia/controllers/transaction"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouterControllerList struct {
	BaseController        base.BaseController
	TransactionController transaction.TransactionController
}

func (ctrl RouterControllerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "TA Response return")
	})

	e.GET("/getData", ctrl.BaseController.GetAll)
	e.GET("/getData/:id", ctrl.BaseController.GetData)
	e.GET("/getData-oltp/:id", ctrl.BaseController.GetDataOLTP)
	e.GET("/getDataWithoutConcurrency/:id", ctrl.BaseController.GetDataWithoutConcurrency)

	admin := e.Group("/admin")
	//Graph
	admin.GET("/graph/sales", ctrl.TransactionController.GetTotalByCashier)
	admin.GET("/graph/sales/by-day", ctrl.TransactionController.GetSalesByDay)
	admin.GET("/graph/product", ctrl.TransactionController.GetProductSales)
	admin.GET("/graph/product-visit", ctrl.BaseController.GetPageVisitGraph)
	admin.GET("/graph/product-visit-oltp", ctrl.BaseController.GetPageVisitGraphOLTP)

	admin.GET("/transaction", ctrl.TransactionController.GetAll)
	admin.POST("/transaction/add", ctrl.TransactionController.AddTransaction)

}
