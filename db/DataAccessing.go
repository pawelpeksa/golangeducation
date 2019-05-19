package db

import (
	"goserver/models"
)

type DataAccessing interface {
	AddUser(profile models.Profile) error
	DoesUserExist(username string) (bool, error)
	AreCredentaialsOk(username string, encryptedPassword string) (bool, error)
	IsBearerValid(bearer string) (bool, error)
	AddBearer(username string, bearer string) error
	RemoveBearer(bearer string) error
	GetBearerForUser(username string) (string, error)
}
