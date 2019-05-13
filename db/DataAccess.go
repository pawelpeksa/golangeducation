package db

import (
	"goserver/models"
)

type DataAccess struct {
	test models.Profile
}

func (da DataAccess) CreateUser(profile models.Profile) error {
	return nil
}

func (da DataAccess) doesUserExist(username string) bool {
	return false
}
