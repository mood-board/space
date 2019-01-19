package host

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/spaceship/common/log"
)

var id bson.ObjectId

func init() {
	id = bson.NewObjectId()
	log.Infof("HOST ID: %s", id.Hex())
}

func ID() bson.ObjectId {
	return id
}
