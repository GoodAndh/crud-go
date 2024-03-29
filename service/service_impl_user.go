package service

import (
	"context"
	"crud3/helper/app"
	"crud3/helper/invalid"
	"crud3/model/domain"
	"crud3/model/web"
	"crud3/repository"
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
)

func NewServiceUser(repoUser repository.RepositoryUser, Db *sql.DB, validate *validator.Validate) ServiceUser {
	return &ServiceImpl{
		DB:       Db,
		validate: validate,
		repoUser: repoUser,
	}
}

func (service *ServiceImpl) GetSingleUser(ctx context.Context, request web.GetUser) (web.GetUser, map[string]interface{}) {
	errorList := make(map[string]interface{})

	tx, err := service.DB.Begin()
	if err != nil {
		invalid.PanicIfError(err)
	}
	defer invalid.CommitOrRollback(tx)

	get := &domain.UserTable{
		Username: request.Username,
		Password: request.Password,
	}
	User, err := service.repoUser.GetSingleUser(ctx, tx, *get)
	if err != nil {
		errorList["error"] = err
		return web.GetUser{}, errorList
	}

	hashedpass := invalid.SetPassword(User.Password)

	if err := invalid.CheckPassword(string(hashedpass), request.Password); err != nil || hashedpass == nil {
		errorList["error"] = errors.New("username")
		return web.GetUser{}, errorList
	} else {
		errorList["sukses"] = true
		return app.ConvertUserToWeb(User), errorList
	}

}
