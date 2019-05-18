package controllers

import (
	"github.com/julienschmidt/httprouter"
	"goserver/db"
	"net/http"
)

type loginController struct {
	da db.DataAccessing
}

func NewLoginController(da db.DataAccessing) *loginController {
	lc := new(loginController)
	lc.da = da
	return lc
}

func (rc loginController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (rc loginController) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
