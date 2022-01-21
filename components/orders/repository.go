package orders

import (
	"context"
	helper "echo-jwt/helpers"
	"fmt"
)

type Repository interface {
	GetQuery(ctx context.Context) ([]Order, error)
}

func (d *OrderDeps) GetQuery(ctx context.Context) ([]Order, error) {
	db, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrConnFailed.Error(), err)
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrBeginTx.Error(), err)
	}
	defer tx.Rollback()

	query := "SELECT id, name, product_id, code, qty, total_price, status FROM orders order by created_at ASC"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error(), err)
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

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(helper.ErrCommit.Error(), err)
	}
	return orders, nil
}
