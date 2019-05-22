package controllers

import (
	"goserver/common"
	"goserver/db"
	"net/http"
	"strings"
	"github.com/julienschmidt/httprouter"
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

func (lc loginController) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	bearer := lc.bearerFromRequest(r)

	if bearer == "" {
		common.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	isBearerValid, err := lc.da.IsBearerValid(bearer)
	
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !isBearerValid {
		common.RespondError(w, http.StatusUnauthorized, "not authorized")
		return
	}

	err = lc.da.RemoveBearer(bearer)

	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondJSON(w, http.StatusOK, "Logged out")
}

func (lc loginController) Authenticate(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		bearer := lc.bearerFromRequest(r)

		if bearer == "" {
			common.RespondError(w, http.StatusBadRequest, "Bad request")
			return
		}

		isBearerValid, err := lc.da.IsBearerValid(bearer)
		
		if err != nil {
			common.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !isBearerValid {
			common.RespondError(w, http.StatusUnauthorized, "not authorized")
			return
		}

		h(w, r, params)
	}
}

func (lc loginController) bearerFromRequest(r *http.Request) string {
	header := r.Header.Get("Authorization")

	parts := strings.Split(header, " ")

	if len(parts) != 2 {
		return ""
	}

	if parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

