package main

import (
	"crud3/controller"
	"crud3/helper/app"
	"crud3/helper/invalid"
	"crud3/repository"
	"crud3/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	db := app.NewDb()
	newRepository := repository.NewRepositoryUser()
	newRepoProduk := repository.NewRepositoryProduk()
	newServiceUser := service.NewServiceUser(newRepository, db, validate)
	newServiceProduk := service.NewServiceProduk(newRepoProduk, db, validate)
	newController := controller.NewController(newServiceUser, newServiceProduk)

	http.HandleFunc("/login", newController.Login)
	http.HandleFunc("/", newController.Index)

	err := http.ListenAndServe(":8080", nil)
	invalid.PanicIfError(err)

}
