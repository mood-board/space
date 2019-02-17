package account

import (
	"context"

	"github.com/ofonimefrancis/spaceship/common"
	"github.com/ofonimefrancis/spaceship/common/log"
	"github.com/ofonimefrancis/spaceship/common/mgo"
)

type Handler struct {
	datastore *Datastore
}

type UserDataProvider struct {
	datastore *Datastore
}

func NewUserDataProvider(initContext context.Context, database *mgo.Database) *UserDataProvider {
	datastore := NewDatastore(initContext, database)
	return &UserDataProvider{datastore}
}

func NewHandler(initContext context.Context, database *mgo.Database) *Handler {
	datastore := NewDatastore(initContext, database)
	handler := &Handler{datastore}
	if !common.IsFlagEnabled("test") {
		return handler
	}

	if err := createTestUser(context.Background(), handler, datastore); err != nil {
		log.Warningf("Unable to create test user: %+v\n", err)
	}
	//TODO: Create a test User
	return handler
}

func createTestUser(c context.Context, handler *Handler, datastore *Datastore) error {
	log.Infof("Creating test user account, username=test@spaceship.com, password=phoenix@*01")
	//handler.SetupUser()
	return nil
}

// func (handler Handler) HandleLogin(c context.Context, email, password string) {
// 	//Check If the user's email and password exists
// 	ds := handler.datastore.OpenSession(c)

// 	if !ds.IsValidLoginCredentials(email, password) {
// 		//Invalid  Login Credentials
// 		log.Infof("[LOGIN] Invalid email and password")
// 		return nil, e
// 	}

// }
