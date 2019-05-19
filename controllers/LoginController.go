package controllers

import (
	"github.com/julienschmidt/httprouter"
	"goserver/common"
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

func (lc loginController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username, password, hasAuth := r.BasicAuth()

	if !hasAuth {
		common.RespondError(w, http.StatusBadRequest, "No BasicAuth in request")
		return
	}

	encryptedPassword := common.Encryptor{}.Encrypt(password)

	areCredentaialsOk, err := lc.da.AreCredentaialsOk(username, encryptedPassword)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !areCredentaialsOk {
		common.RespondError(w, http.StatusUnauthorized, "Wrong username or password")
		return
	}

	bearer, err := lc.da.GetBearerForUser(username)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if bearer != "" {
		common.RespondJSON(w, http.StatusOK, map[string]string{"Bearer": bearer})
		return
	}

	uuid, err := common.UUID()
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = lc.da.AddBearer(username, uuid)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondJSON(w, http.StatusOK, map[string]string{"Bearer": uuid})
}

func (rc loginController) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
