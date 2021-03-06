package db

import (
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

func (da DataAccess) AddUser(profile models.Profile) error {
	c := da.usersCollection()
	err := c.Insert(profile)
	return err
}

func (da DataAccess) DoesUserExist(username string) (bool, error) {

	c := da.usersCollection()

	count, err := c.Find(bson.M{"username": username}).Limit(1).Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (da DataAccess) IsBearerValid(bearer string) (bool, error) {
	c := da.bearersCollection()
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
	c := da.bearersCollection()
	bearerObject := bearerStruct{username, bearer}
	err := c.Insert(bearerObject)
	return err
}

func (da DataAccess) RemoveBearer(bearer string) error {
	c := da.bearersCollection()
	err := c.Remove(bson.M{"bearer": bearer})
	return err
}

func (da DataAccess) HashForUsername(username string) (string, bool, error) {
	profile := models.Profile{}

	c := da.usersCollection()
	query := c.Find(bson.M{"username": username}).Limit(1)
	count, err := query.Count()

	if err != nil {
		return "", false, err
	}

	doesUserExist := count > 0

	if !doesUserExist {
		return "", false, nil
	}

	err = query.One(&profile)

	return profile.Password, true, err
}

func (da DataAccess) GetBearerForUsername(username string) (string, error) {
	bearerObject := bearerStruct{}

	bearers := da.bearersCollection()
	query := bearers.Find(bson.M{"username": username}).Limit(1)

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

func (da DataAccess) session() *mgo.Session {
	return da.dataStore.session()
}

const dbName   = "db"
const usersC   = "users"
const bearersC = "bearers"

func (da DataAccess) usersCollection() *mgo.Collection {
	return da.session().DB(dbName).C(usersC)
}

func (da DataAccess) bearersCollection() *mgo.Collection {
	return da.session().DB(dbName).C(bearersC)
}