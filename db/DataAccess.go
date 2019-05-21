package db

import (
	"errors"
	"goserver/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbAddress = "10.16.22.198"
//const dbAddress = "localhost"

type DataAccess struct {
	session *mgo.Session
}

type bearerStruct struct {
	Username string
	Bearer   string
}

func NewDataAccess() (DataAccessing, error) {
	da := new(DataAccess)

	session, err := mgo.Dial(dbAddress)
	da.session = session

	return da, err
}

func (da DataAccess) AddUser(profile models.Profile) error {
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

func (da DataAccess) AddBearer(username string, bearer string) error {
	c := da.session.DB("db").C("bearers")
	bearerObject := bearerStruct{username, bearer}
	err := c.Insert(bearerObject)
	return err
}

func (da DataAccess) RemoveBearer(bearer string) error {
	c := da.session.DB("db").C("bearers")
	err := c.Remove(bson.M{"bearer": bearer})
	return err
}

func (da DataAccess) AreCredentaialsOk(username string, encryptedPassword string) (bool, error) {
	return true, nil
	return false, errors.New("TODO")
}

func (da DataAccess) GetBearerForUser(username string) (string, error) {
	bearerObject := bearerStruct{}

	c := da.session.DB("db").C("bearers")
	query := c.Find(bson.M{"username": username}).Limit(1)

	count, err := query.Count()

	if err != nil {
		return "", err
	}

	if count == 0 {
		return "", nil
	}

	err = query.One(&bearerObject)

	return bearerObject.Bearer, err
}
