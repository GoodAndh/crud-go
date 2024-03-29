package controller

import (
	"crud3/service"
	"net/http"
)

type Controller interface {
	Login(w http.ResponseWriter,r *http.Request)
	Index(w http.ResponseWriter,r *http.Request)
}

type ControllerImpl struct {
	serUser service.ServiceUser
	serProduk service.ServiceProduk
}