package account

import (
	"context"

	"github.com/ofonimefrancis/spaceship/common"
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

	//TODO: Create a test User
	return handler
}
