package service

import (
	"context"
	"crud3/helper/app"
	"crud3/helper/invalid"
	"crud3/model/domain"
	"crud3/model/web"
	"crud3/repository"
	"database/sql"
	"strings"

	"github.com/go-playground/validator/v10"
)

func SearchFilter(data []web.GetTableProduk, input string) ([]web.GetTableProduk, error) {
	var tabel []web.GetTableProduk
	for _, v := range data {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(input)) || strings.Contains(strings.ToLower(v.Category), strings.ToLower(input)) {
			tabel = append(tabel, v)
		}
	}
	if input == "" || len(input) <= 3 || len(tabel) <= 0 {
		return []web.GetTableProduk{}, invalid.ErrNotFound
	}
	return tabel, nil

}

func NewServiceProduk(repoProduk repository.RepositoryProduk, Db *sql.DB, validate *validator.Validate) ServiceProduk {
	return &ServiceImpl{
		repoProduk: repoProduk,
		DB:         Db,
		validate:   validate,
	}
}

func (service *ServiceImpl) GetAllProduk(ctx context.Context, request []web.GetTableProduk) []web.GetTableProduk {

	var t []domain.ProdukTable
	produk := service.repoProduk.GetAllProduk(ctx, service.DB, t)

	return app.ConvertTableToSlice(produk)

}
