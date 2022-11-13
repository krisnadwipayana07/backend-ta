package base

import (
	"net/http"
	"snatia/business/base"
	"snatia/controllers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BaseController struct {
	baseUsecase base.Usecase
}

func NewBaseController(BaseUsecase base.Usecase) *BaseController {
	return &BaseController{
		baseUsecase: BaseUsecase,
	}
}

func (ctrl BaseController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := ctrl.baseUsecase.GetAllData(ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, data)
}

func (ctrl BaseController) GetData(c echo.Context) error {
	sId := c.Param("id")
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(sId, 10, 64)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	data, err := ctrl.baseUsecase.GetData(ctx, uint(id))

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, data)
}
