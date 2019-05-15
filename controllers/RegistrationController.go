package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goserver/db"
	"goserver/models"
	"goserver/utils"
	"net/http"
)

type registrationController struct {
	da      db.DataAccess
	session *mgo.Session
}

func NewRegistrationController() *registrationController {
	rc := new(registrationController)
	s, _ := mgo.Dial("localhost")
	rc.session = s
	return rc
}

func (rc registrationController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p models.Profile
	err := decoder.Decode(&p)

	p.Password = utils.Encryptor{}.Encrypt(p.Password)

	if err != nil {
		panic(err)
		// TOOD: return HTTP STATUS
	}

	if err != nil {
		fmt.Printf("\n\nCouldn't open session with mongo with error:%v\n", err)
		// TOOD: return HTTP STATUS
		return
	}

	c := rc.session.DB("db").C("users")

	count, err := c.Find(bson.M{"username": p.Username}).Limit(1).Count()

	if count > 0 {
		fmt.Printf("Username already taken\n")
		// TODO: return https status with message
		fmt.Fprint(w, "Username already taken\n")
		return
	}

	fmt.Printf("\nTEST: %v", count)
	fmt.Printf("\nEND")

	err = c.Insert(p)

	if err != nil {
		fmt.Printf("\nSomething went wrong:%v\n", err)
		// TOOD: return HTTP STATUS
	} else {
		fmt.Printf("\n Looks like success:%v\n", p)
	}

	fmt.Printf("\n\n%v\n", p.Username)
	fmt.Printf("\n%v\n", p.Password)
	fmt.Printf("\n%v\n\n", p.Email)
}
