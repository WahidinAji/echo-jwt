package products

import (
	"fmt"
	res "github.com/WahidinAji/web-response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (cv CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (d *ProductDependency) GetAll(ctx echo.Context) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}
	return ctx.JSON(http.StatusOK, res.WebResponse(http.StatusOK, "OK", rows))
}

func (d *ProductDependency) GetById(ctx echo.Context) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	userId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	row, err := d.FindId(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, res.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	fmt.Println("row : ", row, " err : ", err)
	return ctx.JSON(http.StatusOK, res.WebResponse(http.StatusOK, "OK", row))
}
