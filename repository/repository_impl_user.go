package repository

import (
	"context"
	"crud3/helper/invalid"
	"crud3/model/domain"
	"database/sql"
	"errors"
)

func NewRepositoryUser() RepositoryUser {
	return &RepositoryImpl{}
}

func (controller *RepositoryImpl) GetSingleUser(ctx context.Context, tx *sql.Tx, user domain.UserTable) (domain.UserTable, error) {
	script := "select iduser,username,email,password from user where username = ?  limit 1;"
	rows, err := tx.QueryContext(ctx, script, user.Username)
	if err != nil {
		invalid.PanicIfError(err)
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.Username, &user.Email, &user.Password)
		invalid.PanicIfError(err)
		return user, nil

	} else {
		return user, errors.New("password atau username salah")
	}

}
