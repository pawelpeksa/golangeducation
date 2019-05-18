package db

import (
	"goserver/models"
)

type DataAccessing interface {
	CreateUser(profile models.Profile) error
	DoesUserExist(username string) (bool, error)
}
