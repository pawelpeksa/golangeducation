package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goserver/common"
	"goserver/db"
	"goserver/models"
	"net/http"
)

type registrationController struct {
	da db.DataAccessing
}

func NewRegistrationController(da db.DataAccessing) *registrationController {
	rc := new(registrationController)
	rc.da = da
	return rc
}

func (rc registrationController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("Register")
	decoder := json.NewDecoder(r.Body)
	var p models.Profile
	err := decoder.Decode(&p)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if p.Username == "" || p.Email == "" || p.Password == "" {
		common.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	p.Password = common.Encryptor{}.Encrypt(p.Password)

	doesUserExist := rc.da.DoesUserExist(p.Username)

	if doesUserExist {
		common.RespondError(w, http.StatusConflict, "Username taken")
		return
	}

	err = rc.da.CreateUser(p)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondError(w, http.StatusOK, "Username created")
}
