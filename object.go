package goma

import (
	"encoding/json"
	"fmt"
	"goma/gomadb"
	"log"
)

type DBObject interface {
	Key() string
}

type Object struct {
}

func (obj *Object) Save(dbObj DBObject) error {
	data, err := json.Marshal(dbObj)
	if err != nil {
		return err
	}
	Log(fmt.Sprintf("Saving: %s : %s", dbObj.Key(), string(data)))
	return gomadb.GetDB().PutBytes(dbObj.Key(), data)
}

func (obj *Object) Restore(dbObj DBObject) error {
	data, err := gomadb.GetDB().GetBytes(dbObj.Key())
	Log(fmt.Sprintf("Restoring %s : %s", dbObj.Key(), string(data)))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &dbObj)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
