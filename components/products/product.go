package products

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID    uuid.UUID `json:"id" query:"id"`
	Name  string    `json:"name" validate:"required"`
	Stock int8      `json:"stock" validate:"required"`
	Price float64   `json:"price" validate:"required"`
}

type ProductDependency struct {
	DB *sqlx.DB
}

type CustomValidator struct {
	Validator *validator.Validate
}
