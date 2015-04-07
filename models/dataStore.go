package models

import "gopkg.in/mgo.v2"

// TODO: move to a config
const (
	DbName         = "jaabo"
	JobsCollection = "jobs"
)

type DataStore struct {
	session *mgo.Session
}

//http://godoc.org/labix.org/v2/mgo#Dial
// Only the first dial is called to connect to db
// further requests are established using New or copy
func NewDataStore() (*DataStore, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &DataStore{session}, nil
}

func (d *DataStore) Close() {
	d.session.Close()
}
