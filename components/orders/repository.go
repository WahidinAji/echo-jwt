package orders

import (
	"context"
	"database/sql"
	helper "echo-jwt/helpers"
	"fmt"
)

type Repository interface {
	GetQuery(ctx context.Context) (*sql.Rows, error)
}

func (d *OrderDeps) GetQuery(ctx context.Context) (*sql.Rows, error) {
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

	rows, err := tx.QueryContext(ctx, query)

	orders := rows

	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error(), err)
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(helper.ErrCommit.Error(), err)
	}
	return orders, nil
}
