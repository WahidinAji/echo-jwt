package orders

import (
	res "github.com/WahidinAji/web-response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d *OrderDeps) GetAll(ctx echo.Context) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}
	return ctx.JSON(http.StatusOK, res.WebResponse(http.StatusOK, "OK", rows))
}
