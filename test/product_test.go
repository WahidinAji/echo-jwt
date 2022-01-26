package test

import (
	"context"
	"echo-jwt/components/products"
	conf "echo-jwt/helpers"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
)

const dbName = "echo_jwt_testing"

func TestProducts(t *testing.T) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DbUser, conf.DbPass, dbName)
	dbSqlx, errSqlx := sqlx.Open("postgres", dsn)
	defer dbSqlx.Close()
	if errSqlx != nil {
		log.Fatalf("during opening a postgres client:", fmt.Errorf(conf.ErrConnInv.Error(), errSqlx))
	}

	ctx := context.Background()
	stmt, err := dbSqlx.PrepareContext(ctx, "insert into products(id, name, stock, price) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatalf("prepared error : ", err.Error())
	}
	defer stmt.Close()

	//Get all
	t.Run("Get all products", func(t *testing.T) {
		stmt.ExecContext(ctx, "products 1", 10, 100.99)
		product := products.ProductDependency{DB: dbSqlx}
		fmt.Println(product)
	})
}
