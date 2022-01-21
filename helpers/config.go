package helpers

import "os"

var (
	ConfAppName    = os.Getenv("APP_NAME")
	ConfServerPort = os.Getenv("SERVER_PORT")
	DbUser         = os.Getenv("DB_USER")
	DbPass         = os.Getenv("DB_PASS")
	DbName         = os.Getenv("DB_NAME")
	DbUrl          = os.Getenv("PG_URL")
	JWTSecret      = os.Getenv("JWT_SECRET")
)
