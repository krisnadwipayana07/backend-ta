package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status int   `json:"status"`
		Error  error `json:"error"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Error = errors.New("")
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, errs error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Error = errs
	return c.JSON(status, response)
}
