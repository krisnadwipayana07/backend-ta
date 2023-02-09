package transaction

import (
	"net/http"
	"snatia/business/transaction"
	"snatia/controllers"
	"snatia/controllers/transaction/request"
	"snatia/controllers/transaction/response"

	"snatia/helper/timeconvert"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase transaction.UseCase
}

func NewTransactionController(transactionUsecase transaction.UseCase) *TransactionController {
	return &TransactionController{
		transactionUsecase: transactionUsecase,
	}
}

func (ctrl TransactionController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := ctrl.transactionUsecase.GetAllTransaction(ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, data)
}
func (ctrl TransactionController) GetTotalByCashier(c echo.Context) error {
	ctx := c.Request().Context()
	startDateTemp := c.QueryParam("dateStart")
	endDateTemp := c.QueryParam("dateEnd")

	startDate := timeconvert.UnixTimestampConvert(startDateTemp)
	endDate := timeconvert.UnixTimestampConvert(endDateTemp)

	label, data, err := ctrl.transactionUsecase.GetTotalByCashier(ctx, startDate, endDate)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, response.FromLabelString(label, data))
}
func (ctrl TransactionController) GetProductSales(c echo.Context) error {
	ctx := c.Request().Context()
	startDateTemp := c.QueryParam("dateStart")
	endDateTemp := c.QueryParam("dateEnd")

	startDate := timeconvert.UnixTimestampConvert(startDateTemp)
	endDate := timeconvert.UnixTimestampConvert(endDateTemp)

	label, data, err := ctrl.transactionUsecase.GetProductSales(ctx, startDate, endDate)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, response.FromLabelString(label, data))
}
func (ctrl TransactionController) AddTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.AddRequest{}
	c.Bind(&req)

	data, err := ctrl.transactionUsecase.AddTransaction(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, data)
}
func (ctrl TransactionController) GetSalesByDay(c echo.Context) error {
	ctx := c.Request().Context()
	startDateTemp := c.QueryParam("dateStart")
	endDateTemp := c.QueryParam("dateEnd")

	startDate := timeconvert.UnixTimestampConvert(startDateTemp)
	endDate := timeconvert.UnixTimestampConvert(endDateTemp)

	label, value, err := ctrl.transactionUsecase.GetSalesByDay(ctx, startDate, endDate)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, response.FromLabelString(label, value))
}
