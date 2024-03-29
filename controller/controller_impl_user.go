package controller

import (
	"crud3/helper/app"
	"crud3/helper/invalid"
	"crud3/model/web"
	"crud3/service"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func NewController(seruser service.ServiceUser, serProduk service.ServiceProduk) Controller {
	return &ControllerImpl{
		serProduk: serProduk,
		serUser:   seruser,
	}
}

func (controller *ControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("views/login.html")
		invalid.PanicIfError(err)
		t.Execute(w, nil)

	case http.MethodPost:
		r.ParseForm()
		input := &web.GetUser{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		web, errorList := controller.serUser.GetSingleUser(r.Context(), *input)
		fmt.Println("erorlist==", errorList)
		if sukses, ok := errorList["sukses"].(bool); !sukses || !ok {
			t, err := template.ParseFiles("views/login.html")
			invalid.PanicIfError(err)
			t.Execute(w, errorList)
			return
		} else {
			ses, _ := app.Store.Get(r, "acc-log")
			ses.Values["username"] = web.Username
			ses.Values["auten"] = true
			ses.Save(r, w)
			time.Sleep(time.Second * 5)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

func (controller *ControllerImpl) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var tl []web.GetTableProduk
		data := make(map[string]interface{})
		produk := controller.serProduk.GetAllProduk(r.Context(), tl)

		list := app.ForRange(produk)
		data["data"] = list

		t, err := template.ParseFiles("views/index.html")
		invalid.PanicIfError(err)
		t.Execute(w, data)
	case http.MethodPost:
		r.ParseForm()
		data := make(map[string]interface{})
		ses, err := app.Store.Get(r, "acc-log")
		if err != nil {
			data["error"] = invalid.ErrNoSesFound
		} else {
			if auten, ok := ses.Values["auten"].(bool); !auten || !ok {
				data["error"] = invalid.ErrNoSesFound
				data["errormsg"] = true
			}
			var tl []web.GetTableProduk
			produk := controller.serProduk.GetAllProduk(r.Context(), tl)
			tabel, err := service.SearchFilter(produk, r.Form.Get("namaproduk"))
			if err != nil {
				data["error"] = err
			}
			list := app.ForRange(tabel)
			data["data"] = list

			t, err := template.ParseFiles("views/index.html")
			invalid.PanicIfError(err)
			t.Execute(w, data)

		}

	}
}
