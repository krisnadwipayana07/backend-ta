package routes

import (
	"net/http"
	"snatia/controllers/base"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouterControllerList struct {
	BaseController base.BaseController
}

func (ctrl RouterControllerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "SNATIA Response return")
	})

	e.GET("/getData", ctrl.BaseController.GetAll)
	e.GET("/getData/:id", ctrl.BaseController.GetData)
}
