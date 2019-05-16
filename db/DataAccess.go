package db

import (
	"goserver/models"
	"gopkg.in/mgo.v2"
	//"goserver/db"
)

type DataAccess struct {
	session *mgo.Session
}

func NewDataAccess() *DataAccessing {
	return nil
}

func (da DataAccess) CreateUser(profile models.Profile) error {
	return nil
}

func (da DataAccess) doesUserExist(username string) bool {
	return false
}
