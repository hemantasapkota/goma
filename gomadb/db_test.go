package gomadb

import (
	ldb "goma/gomadb/leveldb"
	"testing"
)

var db DB
var err error

func TestInitDB(t *testing.T) {
	db, err = ldb.InitDB(".")
	if err != nil {
		t.Log("DB should not be nil.")
	}
}

func TestPutGetDB(t *testing.T) {
	db.Put("abc", "Hello World")
	val, _ := db.Get("abc")
	if val != "Hello World" {
		t.Error("Inserted value should be the same as put value")
	}
}
