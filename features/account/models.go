package account

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/spaceship/common/mgo"
)

type User struct {
	ID               bson.ObjectId `json:"_id,omitempty"`
	Name             string        `json:"name"`
	Password         string        `json:"password"`
	PasswordHash     []byte        `json:"passwordhash"`
	Email            string        `json:"email"`
	PhoneNumber      string        `json:"phone_number"`
	AvatarURL        string        `json:"avatar_url"`
	IsActive         bool          `json:"is_active"`
	IsVerified       bool          `json:"is_verified"`
	Token            string        `json:"token"`
	TokenExpiryTime  time.Time     `json:"token_expiry_time"`
	VerificationCode string        `json:"verification_code"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

type Datastore struct {
	database *mgo.Database
}

type DatastoreSession struct {
	database *mgo.Database
}

func NewDatastore(initContext context.Context, database *mgo.Database) *Datastore {
	datastore := &Datastore{
		database: database,
	}
	session := datastore.OpenSession(initContext)
	mgo.EnsureOrUpgradeIndexKey(session.users(), "email")

	return datastore
}

func (ds *Datastore) OpenSession(c context.Context) *DatastoreSession {
	db := ds.database
	return &DatastoreSession{
		database: db.FromContext(c),
	}
}

func (datastore *DatastoreSession) users() *mgo.Collection {
	return datastore.database.C("users")
}

func (datastore *DatastoreSession) CreateUser(user User) (*User, error) {
	err := datastore.users().Insert(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (datastore *DatastoreSession) UpdateUser(user *User) error {
	return datastore.users().UpdateId(user.ID, user)
}

func (datastore *DatastoreSession) FindUserByID(id bson.ObjectId) (*User, error) {
	var user User
	err := datastore.users().Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (datastore *DatastoreSession) FindUserByEmail(email string) (*User, error) {
	var user User
	err := datastore.users().Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (datastore *DatastoreSession) RemoveUserByID(id bson.ObjectId) error {
	return datastore.users().Remove(bson.M{"_id": id})
}
