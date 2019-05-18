package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goserver/models"
)

//const dbAddress = "10.16.22.198"
const dbAddress = "localhost"

type DataAccess struct {
	session *mgo.Session
}

func NewDataAccess() (DataAccessing, error) {
	da := new(DataAccess)

	session, err := mgo.Dial(dbAddress)
	da.session = session

	return da, err
}

func (da DataAccess) CreateUser(profile models.Profile) error {
	c := da.session.DB("db").C("users")
	err := c.Insert(profile)
	return err
}

func (da DataAccess) DoesUserExist(username string) (bool, error) {

	c := da.session.DB("db").C("users")

	count, err := c.Find(bson.M{"username": username}).Limit(1).Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (da DataAccess) IsBearerValid(bearer string) (bool, error) {
	return false, errors.New("TODO")
}

func (da DataAccess) AddBearer(bearer string) error {
	return nil
	return errors.New("TODO")
}

func (da DataAccess) RemoveBearer(bearer string) error {
	return errors.New("TODO")
}

func (da DataAccess) AreCredentaialsOk(username string, encryptedPassword string) (bool, error) {
	return true, nil
	return false, errors.New("TODO")
}
