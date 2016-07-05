package gomadb

import (
	"goma/gomadb/leveldb"
)

type DB interface {
	Put(key string, value string) error
	PutBytes(key string, value []byte) error
	Get(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	Delete(key string) error
}

func GetDB() DB {
	if leveldb.DB == nil {
		leveldb.InitDB(".")
	}
	return leveldb.DB
}
