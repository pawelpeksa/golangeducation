package db

import (
	"errors"
	"goserver/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataAccess struct {
	dataStore DataStore
}

type bearerStruct struct {
	Username string
	Bearer   string
}

func NewDataAccess(session *mgo.Session) DataAccessing {
	da := new(DataAccess)
	da.dataStore = DataStore{session}
	return da
}

func (da DataAccess) session() *mgo.Session {
	return da.dataStore.session()
}

func (da DataAccess) AddUser(profile models.Profile) error {
	c := da.session().DB("db").C("users")
	err := c.Insert(profile)
	return err
}

func (da DataAccess) DoesUserExist(username string) (bool, error) {

	c := da.session().DB("db").C("users")

	count, err := c.Find(bson.M{"username": username}).Limit(1).Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (da DataAccess) IsBearerValid(bearer string) (bool, error) {
	c := da.session().DB("db").C("bearers")
	query := c.Find(bson.M{"bearer": bearer}).Limit(1)

	count, err := query.Count()

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (da DataAccess) AddBearer(username string, bearer string) error {
	c := da.session().DB("db").C("bearers")
	bearerObject := bearerStruct{username, bearer}
	err := c.Insert(bearerObject)
	return err
}

func (da DataAccess) RemoveBearer(bearer string) error {
	c := da.session().DB("db").C("bearers")
	err := c.Remove(bson.M{"bearer": bearer})
	return err
}

func (da DataAccess) AreCredentaialsOk(username string, encryptedPassword string) (bool, error) {
	return true, nil
	return false, errors.New("TODO")
}

func (da DataAccess) GetBearerForUser(username string) (string, error) {
	bearerObject := bearerStruct{}

	c := da.session().DB("db").C("bearers")
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
