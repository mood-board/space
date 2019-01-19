package mgo

import (
	"github.com/globalsign/mgo"
	"github.com/ofonimefrancis/spaceship/common/log"
	"github.com/ofonimefrancis/spaceship/common/must"
)

type Index mgo.Index

func EnsureOrUpgrade(collection *Collection, index Index) {
	ensureOrUpgradeIndexF(collection, func() ([]string, string, error) {
		keys := index.Key
		return keys, index.Name, collection.EnsureIndex(
			mgo.Index(index),
		)
	})
}

func EnsureOrUpgradeIndexKey(collection *Collection, keys ...string) {
	ensureOrUpgradeIndexF(collection, func() ([]string, string, error) {
		return keys, "", collection.EnsureIndexKey(keys...)
	})
}

func ensureOrUpgradeIndexF(collection *Collection, ensureFunc func() (keys []string, name string, err error)) {
	keys, name, err := ensureFunc()
	if err == nil {
		return
	}

	if name != "" {
		err = collection.DropIndexName(name)
		if err != nil {
			log.Warningf("Collection: %s Could not upgrade index: %s -> %s", collection.Name, keys, err.Error())
		}
	}

	err = collection.DropIndex(keys...)

	if err != nil {
		log.Warningf("Collection: %s Could not upgrade index: %s -> %s", collection.Name, keys, err.Error())
	}
	_, _, err = ensureFunc()
	must.Do(err)
}
