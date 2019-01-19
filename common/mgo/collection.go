package mgo

import "github.com/globalsign/mgo"

type Collection struct {
	*mgo.Collection
}

func newCollection(c *mgo.Collection) *Collection {
	return &Collection{
		Collection: c,
	}
}
