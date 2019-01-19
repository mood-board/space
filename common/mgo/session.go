package mgo

import "github.com/globalsign/mgo"

type Session struct {
	*mgo.Session
}

func NewSession(s *mgo.Session) *Session {
	s.SetSafe(&mgo.Safe{})
	return &Session{Session: s}
}

func (s *Session) DB(name string) *Database {
	return newDatabase(s.Session.DB(name))
}

func (s *Session) Close() {
	panic("think twice before doing it." +
		"if you want it, call session.Sesion.Close()" +
		"if in doubt, contact @opiumated")
}

func (s *Session) Copy() *Session {
	return NewSession(s.Session.Copy())
}

func (s *Session) Clone() *Session {
	return NewSession(s.Session.Clone())
}
