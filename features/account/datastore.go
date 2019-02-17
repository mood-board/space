package account

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/spaceship/common/log"
	"github.com/ofonimefrancis/spaceship/common/mgo"
	"github.com/pkg/errors"
)

type User struct {
	ID               bson.ObjectId `json:"_id,omitempty"`
	Name             string        `json:"name"`
	Password         string        `json:"password,omitempty"`
	PasswordHash     PasswordHash  `json:"password_hash"`
	Email            string        `json:"email"`
	PhoneNumber      string        `json:"phone_number"`
	AvatarURL        string        `json:"avatar_url"`
	IsActive         bool          `json:"is_active"`
	IsVerified       bool          `json:"is_verified"`
	Token            string        `json:"token,omitempty"`
	TokenExpiryTime  time.Time     `json:"token_expiry_time"`
	VerificationCode string        `json:"verification_code,omitempty"`
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

func (datastore *DatastoreSession) CreateNewUser(name, email string, passwordHash *PasswordHash) (*User, error) {
	var user User

	//Check if user with the same email already exists
	err := datastore.users().Find(bson.M{"email": email}).One(&user)
	if err != nil {
		err = errors.Wrapf(err, "Error occurred when querying for existing user with email: %s", email)
		return nil, err
	}

	user.ID = bson.NewObjectId()
	user.Email = email
	user.Name = name
	user.PasswordHash = *passwordHash
	user.CreatedAt = bson.Now()
	user.IsActive = false
	user.IsVerified = false

	return datastore.CreateUser(user)
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

//AccountExists Checks if an account with specified email exists
func (datastore *DatastoreSession) AccountExists(email string) bool {
	var user User
	err := datastore.users().Find(bson.M{"email": email}).One(&user)
	return err == nil
}

func (datastore *DatastoreSession) RemoveUserByID(id bson.ObjectId) error {
	return datastore.users().Remove(bson.M{"_id": id})
}

func (datastore *DatastoreSession) IsVerifiedAccount(email string) bool {
	var user User
	err := datastore.users().Find(bson.M{"email": email, "is_verified": true}).One(&user)
	if err != nil {
		log.Info("[isVerifiedAccount Error] No user found with that email")
		return false
	}
	return true
}

//IsValidLoginCredentials Validates if the user's email and password matches a record
func (datastore *DatastoreSession) IsValidLoginCredentials(email, password string) bool {
	var user User
	passwordHash, err := NewPasswordHash(password)
	if err != nil {
		log.Info("Error hashing password...")
		return false
	}
	if err := datastore.users().Find(bson.M{"email": email, "password": *passwordHash}).One(&user); err != nil {
		log.Info("Error finding user with the specified email and password.")
		return false
	}
	return true
}
