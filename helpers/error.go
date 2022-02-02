package helpers

import "errors"

var (
	ErrConnInv    = errors.New("invalid connection : ")
	ErrNotExists  = errors.New("user id was not found : ")
	ErrExists     = errors.New("email was already exists : ")
	ErrConnFailed = errors.New("connection failed : ")
	ErrQuery      = errors.New("execute query error : ")
	ErrBeginTx    = errors.New("begin transaction error : ")
	ErrScan       = errors.New("scan error : ")
	ErrCommit     = errors.New("commit error : ")
	ErrExec       = errors.New("query exec error : ")

	// ErrNameNotFound user authentication

	ErrNameNotFound      = errors.New("username was not found : ")
	ErrNotMatchUser      = errors.New("wrong password. username and password do not match")
	ErrNameAlreadyExists = errors.New("username was already used")
	ErrHashFailed        = errors.New("hashing password was fail")
)
