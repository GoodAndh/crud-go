package repository

import (
	"context"
	"crud3/helper/invalid"
	"crud3/model/domain"
	"database/sql"
)

func NewRepositoryProduk() RepositoryProduk {
	return &RepositoryImpl{}
}

func (repository *RepositoryImpl) GetAllProduk(ctx context.Context, db *sql.DB, user []domain.ProdukTable) []domain.ProdukTable {
	script := "select iditem,nama,deskripsi,category,userid,harga,quantity from produkuser;"
	rows, err := db.QueryContext(ctx, script)
	invalid.PanicIfError(err)
	defer rows.Close()
	for rows.Next() {
		var t domain.ProdukTable
		err := rows.Scan(&t.IdProduk, &t.Name, &t.Deskripsi, &t.Category, &t.IdUser, &t.Harga, &t.Quantity)
		invalid.PanicIfError(err)
		user = append(user, t)
	}
	return user
}
