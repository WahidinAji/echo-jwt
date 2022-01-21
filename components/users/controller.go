package users

import (
	helper "echo-jwt/helpers"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (d *UserDependency) Login(ctx echo.Context) error {
	//form
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
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

	_token, err := token.SignedString([]byte("secret"))
	if err != nil {
		return fmt.Errorf("token byte error : ", err.Error())
	}
	return ctx.JSON(http.StatusAccepted, helper.JWTResponse{
		MsgDel:   helper.MsgDel{Code: http.StatusAccepted, Status: true},
		Username: user.Username,
		Token:    _token,
	})
}
