package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goserver/models"
)

//const dbAddress = "10.16.22.198"
const dbAddress = "localhost"

type DataAccess struct {
	session *mgo.Session
}

func NewDataAccess() DataAccessing {
	// TODO: handle error while connecting
	session, _ := mgo.Dial(dbAddress)

	da := new(DataAccess)
	da.session = session

	return da
}

func (da DataAccess) CreateUser(profile models.Profile) error {
	c := da.session.DB("db").C("users")
	err := c.Insert(profile)
	return err
}

func (da DataAccess) DoesUserExist(username string) bool {

	c := da.session.DB("db").C("users")

	count, err := c.Find(bson.M{"username": username}).Limit(1).Count()

	if err != nil {
		return false
	}

	return count > 0
}
