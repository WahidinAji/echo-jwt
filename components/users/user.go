package users

import (
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type Username struct {
	Username string `json:"username"`
}
type UserDependency struct {
	DB *sqlx.DB
}

type JwtUserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
