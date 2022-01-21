package main

import (
	orders2 "echo-jwt/components/orders"
	"echo-jwt/components/products"
	"echo-jwt/components/users"
	"echo-jwt/helpers"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	confAppName    = os.Getenv("APP_NAME")
	confServerPort = os.Getenv("SERVER_PORT")
	DbUser         = os.Getenv("DB_USER")
	DbPass         = os.Getenv("DB_PASS")
	DbName         = os.Getenv("DB_NAME")
	DbUrl          = os.Getenv("PG_URL")
)

func init() {
	if confAppName == "" {
		log.Fatal("APP_NAME config is required")
	}
	if confServerPort == "" {
		log.Fatal("SERVER_PORT config is required")
	}
	if DbUrl == "" {
		log.Fatal("PG_URL config is required")
	}
	//if DbUser == "" {
	//	log.Fatal("DB_USER config is required")
	//}
	//if DbPass == "" {
	//	log.Fatal("DB_PASS config is required")
	//}
	//if DbName == "" {
	//	log.Fatal("DB_NAME config is required")
	//}
}

func main() {
	e := echo.New()

	/**Using PostgresSql and Sqlx**/
	//dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUser, DbPass, DbName)
	dsn := DbUrl
	dbSqlx, errSqlx := sqlx.Open("postgres", dsn)
	if errSqlx != nil {
		e.Logger.Fatal("during opening a postgres client:", fmt.Errorf(helpers.ErrConnInv.Error(), errSqlx))
	}
	dbSqlx.SetMaxIdleConns(10)
	dbSqlx.SetMaxOpenConns(100)
	dbSqlx.SetConnMaxIdleTime(5 * time.Minute)
	dbSqlx.SetConnMaxLifetime(60 * time.Minute)
	fmt.Println(dbSqlx.Ping())
	defer dbSqlx.Close()

	e.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, "Hello, World")
	})
	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	//Dependencies
	product := products.ProductDependency{DB: dbSqlx}
	user := users.UserDependency{DB: dbSqlx}

	//without jwt
	e.GET("/products", product.GetAll)
	e.GET("/products/:id", product.GetById)

	api := e.Group("/api")
	api.POST("/login", user.Login)

	orders := orders2.OrderDeps{DB: dbSqlx}
	api.GET("/orders", orders.GetAll)

	//jwt
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		//Claims:     &jwtCustomClaims{},
		Claims:     &users.JwtUserClaims{},
		SigningKey: []byte("secret"),
	}

	////product
	api.Use(middleware.JWTWithConfig(config))
	api.GET("/products", product.GetAll)
	api.GET("/products/:id", product.GetById)

	server := new(http.Server)
	server.Addr = ":" + confServerPort

	if confServerReadTimeout := os.Getenv("SERVER_READ_TIMEOUT_IN_MINUTE"); confServerReadTimeout != "" {
		duration, _ := strconv.Atoi(confServerReadTimeout)
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}

	if confServerWriteTimeout := os.Getenv("SERVER_WRITE_TIMEOUT_IN_MINUTE"); confServerWriteTimeout != "" {
		duration, _ := strconv.Atoi(confServerWriteTimeout)
		server.WriteTimeout = time.Duration(duration) * time.Minute
	}

	e.Logger.Print(confAppName, " is running on http://localhost", server.Addr)
	e.Logger.Fatal(e.StartServer(server))
}
