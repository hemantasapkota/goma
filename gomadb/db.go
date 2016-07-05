package gomadb

import (
	"github.com/hemantasapkota/goma/gomadb/leveldb"
)

type DB interface {
	Put(key string, value string) error
	PutBytes(key string, value []byte) error
	Get(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	Delete(key string) error
}

func SetLevelDB(db *leveldb.GomaDB) {
	leveldb.DB = db
}

func GetDB() DB {
	return leveldb.DB
}
