package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"

	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type GomaDB struct {
	sync.Mutex
	Store *leveldb.DB
}

var DB *GomaDB

func InitDB(path string) (*GomaDB, error) {
	// If DB has been inited, return
	if DB != nil {
		return DB, nil
	}

	filepath := fmt.Sprintf("%s/goma.db", path)
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		return nil, err
	}

	DB = &GomaDB{
		Store: db,
	}

	return DB, nil
}

func RecoverDB(path string) (*GomaDB, error) {
	// If DB has been inited, return
	if DB != nil {
		return DB, nil
	}

	filepath := fmt.Sprintf("%s/goma.db", path)
	db, err := leveldb.RecoverFile(filepath, nil)
	if err != nil {
		return nil, err
	}

	DB = &GomaDB{
		Store: db,
	}
	return DB, nil
}

func RemoveDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func CloseDB() {
	// Synchronize access
	DB.Lock()
	defer DB.Unlock()

	DB.Store.Close()
	DB.Store = nil
	DB = nil
}

func (db *GomaDB) Put(key string, value string) error {
	db.Lock()
	defer db.Unlock()

	db.Store.Put([]byte(key), []byte(value), nil)
	return nil
}
func (db *GomaDB) PutBytes(key string, data []byte) error {
	db.Lock()
	defer db.Unlock()

	db.Store.Put([]byte(key), data, nil)
	return nil
}
func (db *GomaDB) Get(key string) (string, error) {
	db.Lock()
	defer db.Unlock()

	value, err := db.Store.Get([]byte(key), nil)
	return string(value), err
}
func (db *GomaDB) GetBytes(key string) ([]byte, error) {
	db.Lock()
	defer db.Unlock()
	return db.Store.Get([]byte(key), nil)
}
func (db *GomaDB) Delete(key string) error {
	db.Lock()
	defer db.Unlock()
	return db.Store.Delete([]byte(key), nil)
}
