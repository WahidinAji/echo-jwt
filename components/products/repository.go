package products

import (
	"context"
	helper "echo-jwt/helpers"
	"fmt"
	"github.com/google/uuid"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindId(ctx context.Context, id uuid.UUID) (*Product, error)
	Update(ctx context.Context, id uuid.UUID, product Product) (*Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Save(ctx context.Context, product Product) (*Product, error)
}

func (d *ProductDependency) FindAll(ctx context.Context) ([]Product, error) {
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
	query := "SELECT id, name, stock, price FROM products order by created_at ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error(), err)
	}
	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Stock, &product.Price)
		if err != nil {
			return nil, fmt.Errorf(helper.ErrScan.Error(), err)
		}
		products = append(products, product)
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(helper.ErrCommit.Error(), err)
	}
	return products, nil
}

func (d *ProductDependency) FindId(ctx context.Context, id uuid.UUID) (*Product, error) {
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
	query := "SELECT EXISTS (SELECT id FROM products WHERE id=$1)"

	var exists bool
	err = tx.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error()+" 0 ", err)
	}

	if !exists {
		return nil, fmt.Errorf(helper.ErrNotExists.Error())
	}

	query = "SELECT id, name, stock, price FROM products WHERE id=$1"
	res, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error()+" 1 ", err)
	}

	var product Product
	if res.Next() {
		err = res.Scan(&product.ID, &product.Name, &product.Stock, &product.Price)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &product, err

}
