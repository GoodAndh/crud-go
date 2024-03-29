package app

import (
	"crud3/model/domain"
	"crud3/model/web"
	"strconv"
)

func ConvertUserToWeb(input domain.UserTable) web.GetUser {
	return web.GetUser{
		Username: input.Username,
		Password: input.Password,
	}
}

func ConvertTableToWeb(input domain.ProdukTable) web.GetTableProduk {
	return web.GetTableProduk{
		IdUser:    input.IdUser,
		Name:      input.Name,
		Deskripsi: input.Deskripsi,
		Harga:     input.Harga,
		Quantity:  input.Quantity,
		Category:  input.Category,
	}
}

func ConvertTableToSlice(input []domain.ProdukTable) []web.GetTableProduk {
	var web []web.GetTableProduk
	for _, v := range input {
		web = append(web, ConvertTableToWeb(v))
	}
	return web
}

func ForRange(input []web.GetTableProduk) map[string]interface{} {
	data := make(map[string]interface{})
	for i, v := range input {
		data1 := map[string]interface{}{
			"nama":      v.Name,
			"deskripsi": v.Deskripsi,
			"harga":     v.Harga,
			"stock":     v.Quantity,
			"category":  v.Category,
		}
		data[strconv.Itoa(i+1)] = data1
	}
	return data

}
