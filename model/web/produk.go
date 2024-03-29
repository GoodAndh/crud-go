package web

type GetTableProduk struct {
	IdUser    int    `json:"user_id"`
	Name      string `json:"name_produk"`
	Deskripsi string `json:"deskripsi_produk"`
	Harga     int    `json:"harga_produk"`
	Quantity  int    `json:"stock_produk"`
	Category  string `json:"category_produk"`
}
