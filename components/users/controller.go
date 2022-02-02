package users

import (
	helper "echo-jwt/helpers"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (d *UserDependency) Login(ctx echo.Context) error {
	//form
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name can not be null")
	}

	user, err := d.LoginUser(ctx.Request().Context(), username, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Set custom claims
	claims := &JwtUserClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	}

	//create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	_token, err := token.SignedString([]byte(helper.JWTSecret))
	if err != nil {
		return fmt.Errorf("token byte error : ", err.Error())
	}
	return ctx.JSON(http.StatusCreated, helper.AuthResponse{
		Code:    http.StatusCreated,
		Status:  true,
		Message: "Login Success",
		JWTResponse: helper.JWTResponse{
			Username: user.Username,
			Token:    _token,
		},
	})
}

func (d *UserDependency) Register(ctx echo.Context) error {
	//form
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	
	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username and password can not be null")
	}

	addUser := new(User)
	addUser.ID = uuid.New()
	addUser.Name = username
	addUser.Password = password

	user, err := d.RegisterUser(ctx.Request().Context(), *addUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Set custom claims
	claims := &JwtUserClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	}

	//create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	_token, err := token.SignedString([]byte(helper.JWTSecret))
	if err != nil {
		return fmt.Errorf("token byte error : ", err.Error())
	}
	return ctx.JSON(http.StatusCreated, helper.AuthResponse{
		Code:    http.StatusCreated,
		Status:  true,
		Message: "Register Success",
		JWTResponse: helper.JWTResponse{
			Username: user.Username,
			Token:    _token,
		},
	})
}
