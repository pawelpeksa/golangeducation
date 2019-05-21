package controllers

import (
	"encoding/json"
	"goserver/common"
	"goserver/db"
	"goserver/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type registrationController struct {
	da db.DataAccessing
}

func NewRegistrationController(da db.DataAccessing) *registrationController {
	rc := new(registrationController)
	rc.da = da
	return rc
}


// 	Required parameters for request
// 	Username string
//	Password string
//	Email    string

func (rc registrationController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p models.Profile
	err := decoder.Decode(&p)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if p.Username == "" || p.Email == "" || p.Password == "" {
		common.RespondError(w, http.StatusBadRequest, "Missing data in json sent to server")
		return
	}

	p.Password = common.Encryptor{}.Encrypt(p.Password)

	doesUserExist, err := rc.da.DoesUserExist(p.Username)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if doesUserExist {
		common.RespondError(w, http.StatusConflict, "Username taken")
		return
	}

	err = rc.da.AddUser(p)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondJSON(w, http.StatusOK, "Username created")
}
