package users

import (
	"context"
	helper "echo-jwt/helpers"
	"fmt"
)

type Repository interface {
	LoginUser(ctx context.Context, name, password string) (bool, *Username, error)
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
