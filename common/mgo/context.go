package mgo

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	ContextKey = "mongo_connection"

	sessionTimeout = 30 * time.Minute
)

//FromContext return mongo session from context
func FromContext(c context.Context) *Session {
	if s, ok := c.Value(ContextKey).(*Session); ok {
		return s
	}
	return nil
}

func (d *Database) FromContext(c context.Context) *Database {
	if s, ok := c.Value(ContextKey).(*Session); ok {
		return d.with(s)
	}

	//We always have session inside context during http request
	session := d.Session.Copy()
	c = context.WithValue(c, ContextKey, session)

	stacktrace := debug.Stack()
	go func(c context.Context, session *Session, trace []byte) {
		ticker := time.NewTicker(sessionTimeout)
		count := 1
		sessionID := bson.NewObjectId()
		for {
			select {
			case <-c.Done():
				session.Session.Close()
				return
			case <-ticker.C:
				log.Printf("mongo session %s is opened for %s: %s", sessionID, sessionTimeout*time.Duration(count), string(trace))
				count++
			}
		}
	}(c, session, stacktrace)

	return d.with(session)
}

func (d *Database) with(s *Session) *Database {
	return newDatabase(d.Database.With(s.Session))
}
