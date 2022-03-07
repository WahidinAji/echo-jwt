package helpers

import "os"

var (
	ConfAppName    = os.Getenv("APP_NAME")
	ConfServerPort = os.Getenv("SERVER_PORT")
	DbUrl          = os.Getenv("PG_URL")
	JWTSecret      = os.Getenv("JWT_SECRET")
	DbTesting      = os.Getenv("DB_TESTING")
)
