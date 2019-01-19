package mgo

import (
	"log"
	"save-api/config"
	"time"

	"github.com/globalsign/mgo"
)

func New(host, name string, config *config.Config) *Database {
	dbSession, err := DialWithInfo(host, "", "")
	if err != nil {
		log.Panicf("Mongo init, err=%v", err)
	}
	return dbSession.DB(name)
}

func DialWithInfo(url, username, password string) (*Session, error) {
	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		return nil, err
	}
	dialInfo.Username = username
	dialInfo.Password = password

	dialInfo.Timeout = 30 * time.Second

	mgoSession, err := mgo.DialWithInfo(dialInfo)
	return NewSession(mgoSession), err
}
