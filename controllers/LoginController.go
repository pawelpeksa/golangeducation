package controllers

import (
	"github.com/julienschmidt/httprouter"
	"goserver/db"
	"net/http"
)

type loginController struct {
	da      db.DataAccess
}


func NewLoginController() *loginController {
	lc := new(loginController)
	return lc
}

func (rc loginController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}


func (rc loginController) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}