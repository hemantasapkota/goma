package goma

import (
	gomadb "github.com/hemantasapkota/goma/gomadb"
	ldb "github.com/hemantasapkota/goma/gomadb/leveldb"
	"testing"
)

type TestObject struct {
	*Object
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (o TestObject) Key() string {
	return "goma.testObject"
}

func TestObjects(t *testing.T) {
	NewLogger(LoggerConfig{Debug: true})

	db, err := ldb.InitDB(".")
	if err != nil {
		t.Log("Could not setup database")
	}
	gomadb.SetLevelDB(db)

	w := TestObject{
		Name: "Hemanta",
		Age:  29,
	}

	w.Save(w)

	w1 := &TestObject{}
	w1.Restore(w1)

	t.Log(w1)

	if w1.Name != "Hemanta" {
		t.Error("Error saving/restore DBObject: Name")
	}

	if w1.Age != 29 {
		t.Error("Error saving/restore DBObject: Age")
	}
}

func TestCache(t *testing.T) {
	cache := GetAppCache()

	w := &TestObject{
		Name: "Hemanta",
		Age:  29,
	}
	cache.Put(w)

	w1 := &TestObject{}
	w1 = cache.Get(w1).(*TestObject)

	if w1.Name != "Hemanta" {
		t.Error("Object retireved from the cache does not match that one that was put in")
	}
}

func TestTimestamps(t *testing.T) {
	ts := Timestamp()
	if ts == "" {
		t.Error("Timestamp cannot be empty.")
	}

	timestamp := "2016-07-04T16:47:09+10:00"

	t1 := ParseDatetime(timestamp)
	if t1.Year() != 2016 {
		t.Error("Error parsing timestamp. Year should be 2016")
	}

	if t1.Month() != 7 {
		t.Error("Error parsing timestamp. Month should be 7")
	}

	if t1.Day() != 4 {
		t.Error("Error parsing timestamp. Day should be 4")
	}

	ts1 := TimestampFrom(t1)

	if ts1 != timestamp {
		t.Error("Timestamps should match: ", timestamp, ts1)
	}
}
