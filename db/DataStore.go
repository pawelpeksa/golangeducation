package db

import "gopkg.in/mgo.v2"

type DataStore struct {
	mgoSession *mgo.Session
}

func (ds DataStore) session() *mgo.Session {
	return ds.mgoSession
}
