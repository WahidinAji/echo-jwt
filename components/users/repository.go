package users

import (
	"context"
	helper "echo-jwt/helpers"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	LoginUser(ctx context.Context, name, password string) (*Username, error)
	RegisterUser(ctx context.Context, user User) (*Username, error)
}

//authentication

func (d *UserDependency) LoginUser(ctx context.Context, name string, password string) (*Username, error) {
	db, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrConnFailed.Error(), err)
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrBeginTx.Error(), err)
	}
	defer tx.Rollback()

	var authenticate bool

	//check username
	query := "select exists(select name from users where name=$1)"
	err = tx.QueryRowContext(ctx, query, name).Scan(&authenticate)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error()+"1 ", err)
	}
	if !authenticate {
		return nil, fmt.Errorf(helper.ErrNameNotFound.Error(), err)
	}

	//check if username and password doesn't match
	query = "select exists(select name from users where name=$1 and password=$2)"
	err = tx.QueryRowContext(ctx, query, name, password).Scan(&authenticate)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error()+"2 ", err)
	}
	if !authenticate {
		return nil, fmt.Errorf(helper.ErrNotMatchUser.Error())
	}

	var user Username
	query = "SELECT name FROM users WHERE name=$1 and password=$2"

	//row, err := db.QueryContext(ctx, query, name, password)
	row := tx.QueryRowContext(ctx, query, name, password)
	err = row.Scan(&user.Username)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrScan.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(helper.ErrCommit.Error(), err)
	}
	return &user, nil
}

func (d *UserDependency) RegisterUser(ctx context.Context, user User) (*Username, error) {
	db, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrConnFailed.Error(), err)
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrBeginTx.Error(), err)
	}
	defer tx.Rollback()
	//
	var exists bool
	//check username
	query := "select exists(select name from users where name=$1)"
	err = tx.QueryRowContext(ctx, query, user.Name).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error()+"1 ", err)
	}
	if exists {
		return nil, fmt.Errorf(helper.ErrNameAlreadyExists.Error(), err)
	}

	//hash password here
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//pass, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	query = "INSERT INTO users (id, name, password) values($1,$2,$3)"
	_, err = tx.ExecContext(ctx, query, user.ID, user.Name, pass)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrQuery.Error(), err)
	}
	var username Username
	username.Username = user.Name

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(helper.ErrCommit.Error(), err)
	}

	return &username, nil
}
