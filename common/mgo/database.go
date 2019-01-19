package mgo

import "github.com/globalsign/mgo"

type Database struct {
	*mgo.Database
	Session *Session
}

func newDatabase(db *mgo.Database) *Database {
	return &Database{
		Database: db,
		Session:  NewSession(db.Session),
	}
}

func (d *Database) C(name string) *Collection {
	return newCollection(d.Database.C(name))
}
