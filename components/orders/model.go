package orders

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Order struct {
	ID         int       `json:"id"`
	Code       string    `json:"code"`
	ProductId  uuid.UUID `json:"product_id"`
	Name       string    `json:"name"`
	Qty        int       `json:"qty"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
}

//func (o *Order) Save(ctx context.Context) error {
//
//	return errors.New("")
//}

type OrderDeps struct {
	DB *sqlx.DB
}

type CustomValidator struct {
	Validator *validator.Validate
}
