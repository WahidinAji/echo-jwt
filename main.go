package main

import (
	orders2 "echo-jwt/components/orders"
	"echo-jwt/components/products"
	"echo-jwt/components/public"
	"echo-jwt/components/users"
	conf "echo-jwt/helpers"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func init() {
	if conf.DbUrl == "" {
		log.Fatal("PG_URL config is required")
	}
	if conf.JWTSecret == "" {
		log.Fatal("JWT_SECRET config is required")
	}
}

func main() {
	e := echo.New()

	/**Using PostgresSql and Sqlx**/
	dsn := conf.DbUrl
	dbSqlx, errSqlx := sqlx.Open("postgres", dsn)
	if errSqlx != nil {
		e.Logger.Fatal("during opening a postgres client:", fmt.Errorf(conf.ErrConnInv.Error(), errSqlx))
	}

	defer dbSqlx.Close()

	t := &public.Template{
		Template: template.Must(template.ParseGlob("*.html")),
	}
	//e.GET("/", func(c echo.Context) (err error) {
	//	return c.JSON(http.StatusOK, "Hello, World")
	//})
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "index")
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

	server := &http.Server{
		Addr: ":" + conf.ConfServerPort,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	e.Logger.Print(conf.ConfAppName, " is running on http://localhost", server.Addr)
	e.Logger.Fatal(e.StartServer(server))
}
