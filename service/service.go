package service

import (
	"context"
	"crud3/model/web"
	"crud3/repository"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type ServiceUser interface {
	GetSingleUser(ctx context.Context, request web.GetUser) (web.GetUser, map[string]interface{})
}

type ServiceProduk interface {
	GetAllProduk(ctx context.Context, request []web.GetTableProduk) []web.GetTableProduk
}

type ServiceImpl struct {
	repoProduk repository.RepositoryProduk
	repoUser   repository.RepositoryUser
	DB         *sql.DB
	validate   *validator.Validate
}

