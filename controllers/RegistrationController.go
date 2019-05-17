package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"goserver/db"
	"goserver/models"
	"goserver/utils"
	"net/http"
)

type registrationController struct {
	da      db.DataAccessing
	session *mgo.Session
}

func NewRegistrationController(da db.DataAccessing) *registrationController {
	rc := new(registrationController)
	rc.da = da
	return rc
}

func (rc registrationController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p models.Profile
	err := decoder.Decode(&p)

	if err != nil {
		fmt.Printf("\nJson went wrong:%v\n", err)
		// TOOD: return HTTP STATUS
		return
	}

	p.Password = utils.Encryptor{}.Encrypt(p.Password)

	doesUserExist := rc.da.DoesUserExist(p.Username)

	if doesUserExist {
		fmt.Printf("Username already taken\n")
		// TODO: return https status with message
		fmt.Fprint(w, "Username already taken\n")
		return
	}

	
	err = rc.da.CreateUser(p)

	if err != nil {
		fmt.Printf("\nSomething went wrong:%v\n", err)
	} else {
		fmt.Printf("\n Looks like success:%v\n", p)
	}

	fmt.Printf("\n\n%v\n", p.Username)
	fmt.Printf("\n%v\n", p.Password)
	fmt.Printf("\n%v\n\n", p.Email)
}
