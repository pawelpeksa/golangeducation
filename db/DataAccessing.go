package db

import (
	"goserver/models"
)

type DataAccessing interface {
	CreateUser(profile models.Profile) error
	DoesUserExist(username string) (bool, error)
	AreCredentaialsOk(username string, encryptedPassword string) (bool, error)
	IsBearerValid(bearer string) (bool, error)
	AddBearer(bearer string) error
	RemoveBearer(bearer string) error
}
