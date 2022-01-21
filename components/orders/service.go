package orders

import (
	"context"
	helper "echo-jwt/helpers"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	FindAll(ctx context.Context) ([]Order, error)
}

func (cv CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (d *OrderDeps) FindAll(ctx context.Context) ([]Order, error) {
	rows, err := d.GetQuery(ctx)
	fmt.Println(rows)
	if err != nil {
		return nil, errors.New("Get Query Repo Error : " + err.Error())
	}
	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Name, &order.ProductId, &order.Code, &order.Qty, &order.TotalPrice, &order.Status)
		if err != nil {
			return nil, helper.ErrScan
		}
		orders = append(orders, order)
	}
	return orders, nil
}
