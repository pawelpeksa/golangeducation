package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goserver/db"
	"goserver/models"
	"net/http"
)

type RegistrationController struct {
	da db.DataAccess
}

func (rc RegistrationController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p models.Profile
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
		// TOOD: return HTTP STATUS
	}

	session, err := mgo.Dial("localhost")

	if err != nil {
		fmt.Printf("\n\nCouldn't open session with mongo with error:%v\n", err)
		// TOOD: return HTTP STATUS
		return
	}

	c := session.DB("db").C("users")

	count, err := c.Find(bson.M{"username": p.Username}).Limit(1).Count()

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

	session.Close()
}
