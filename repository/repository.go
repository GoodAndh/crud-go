package repository

import (
	"context"
	"crud3/model/domain"
	"database/sql"
)

type RepositoryUser interface {
	GetSingleUser(ctx context.Context, tx *sql.Tx, user domain.UserTable) (domain.UserTable, error)
}
type RepositoryProduk interface {
	GetAllProduk(ctx context.Context,db *sql.DB,user []domain.ProdukTable)[]domain.ProdukTable
}

type RepositoryImpl struct {
}
