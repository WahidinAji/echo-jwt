package main

import (
	orders2 "echo-jwt/components/orders"
	"echo-jwt/components/products"
	"echo-jwt/components/users"
	conf "echo-jwt/helpers"
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

func init() {
	if conf.ConfAppName == "" {
		log.Fatal("APP_NAME config is required")
	}
	if conf.ConfServerPort == "" {
		log.Fatal("SERVER_PORT config is required")
	}
	if conf.DbUrl == "" {
		log.Fatal("PG_URL config is required")
	}
	if conf.DbUser == "" {
		log.Fatal("DB_USER config is required")
	}
	if conf.DbPass == "" {
		log.Fatal("DB_PASS config is required")
	}
	if conf.DbName == "" {
		log.Fatal("DB_NAME config is required")
	}
	if conf.JWTSecret == "" {
		log.Fatal("JWT_SECRET config is required")
	}
}

func main() {
	e := echo.New()

	/**Using PostgresSql and Sqlx**/
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DbUser, conf.DbPass, conf.DbName)
	//dsn := DbUrl
	dbSqlx, errSqlx := sqlx.Open("postgres", dsn)
	if errSqlx != nil {
		e.Logger.Fatal("during opening a postgres client:", fmt.Errorf(conf.ErrConnInv.Error(), errSqlx))
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
	orders := orders2.OrderDeps{DB: dbSqlx}

	api := e.Group("/api")
	api.POST("/login", user.Login)
	api.POST("/register", user.Register)
	//without jwt
	api.GET("/products", product.GetAll)
	api.GET("/products/:id", product.GetById)

	//jwt
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		//Claims:     &jwtCustomClaims{},
		Claims:     &users.JwtUserClaims{},
		SigningKey: []byte(conf.JWTSecret),
	}

	////orders with jwt
	api.Use(middleware.JWTWithConfig(config))
	api.GET("/orders", orders.GetAll)

	server := new(http.Server)
	server.Addr = ":" + conf.ConfServerPort

	if confServerReadTimeout := os.Getenv("SERVER_READ_TIMEOUT_IN_MINUTE"); confServerReadTimeout != "" {
		duration, _ := strconv.Atoi(confServerReadTimeout)
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}

	if confServerWriteTimeout := os.Getenv("SERVER_WRITE_TIMEOUT_IN_MINUTE"); confServerWriteTimeout != "" {
		duration, _ := strconv.Atoi(confServerWriteTimeout)
		server.WriteTimeout = time.Duration(duration) * time.Minute
	}

	e.Logger.Print(conf.ConfAppName, " is running on http://localhost", server.Addr)
	e.Logger.Fatal(e.StartServer(server))
}
