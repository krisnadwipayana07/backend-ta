package base

import (
	"net/http"
	"snatia/business/base"
	"snatia/controllers"
	"snatia/controllers/base/response"
	"snatia/helper/timeconvert"
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
func (ctrl BaseController) GetDataOLTP(c echo.Context) error {
	sId := c.Param("id")
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(sId, 10, 64)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	data, err := ctrl.baseUsecase.GetDataOLTP(ctx, uint(id))

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, data)
}
func (ctrl BaseController) GetDataWithoutConcurrency(c echo.Context) error {
	sId := c.Param("id")
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(sId, 10, 64)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	data, err := ctrl.baseUsecase.GetDataWithoutConcurrency(ctx, uint(id))

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, data)
}

func (ctrl BaseController) GetPageVisitGraph(c echo.Context) error {
	ctx := c.Request().Context()
	startDateTemp := c.QueryParam("dateStart")
	endDateTemp := c.QueryParam("dateEnd")

	startDate := timeconvert.UnixTimestampConvert(startDateTemp)
	endDate := timeconvert.UnixTimestampConvert(endDateTemp)

	label, data, err := ctrl.baseUsecase.GetPageVisitGraph(ctx, startDate, endDate)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, response.FromLabelString(label, data))
}

func (ctrl BaseController) GetPageVisitGraphOLTP(c echo.Context) error {
	ctx := c.Request().Context()
	startDateTemp := c.QueryParam("dateStart")
	endDateTemp := c.QueryParam("dateEnd")

	startDate := timeconvert.UnixTimestampConvert(startDateTemp)
	endDate := timeconvert.UnixTimestampConvert(endDateTemp)

	label, data, err := ctrl.baseUsecase.GetPageVisitGraphOLTP(ctx, startDate, endDate)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, response.FromLabelString(label, data))
}
