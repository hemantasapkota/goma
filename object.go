package goma

import (
	"encoding/json"
	"fmt"
	"github.com/hemantasapkota/goma/gomadb"
	"log"
)

type DBObject interface {
	Key() string
}

type Object struct {
}

func (obj *Object) Delete(dbObj DBObject) error {
	Log(fmt.Sprintf("Deleting: %s", dbObj.Key()))
	return gomadb.GetDB().Delete(dbObj.Key())
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

func (obj *Object) Raw(dbObj DBObject) ([]byte, error) {
	var data []byte
	var err error
	data, err = gomadb.GetDB().GetBytes(dbObj.Key())
	if err != nil {
		data, err = json.Marshal(dbObj)
	}
	return data, err
}
