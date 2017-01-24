package goma

import (
	gomadb "github.com/hemantasapkota/goma/gomadb"
	ldb "github.com/hemantasapkota/goma/gomadb/leveldb"

	"sync"
	"testing"
)

type TestObject struct {
	*Object
	sync.Mutex

	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (o *TestObject) Key() string {
	return "goma.testObject"
}

func TestObjects(t *testing.T) {
	NewLogger(LoggerConfig{Debug: true})

	db, err := ldb.InitDB(".")
	if err != nil {
		t.Log("Could not setup database")
	}
	gomadb.SetLevelDB(db)

	w := &TestObject{
		Name: "Hemanta",
		Age:  29,
	}

	w.Save(w)

	// Test raw
	rawData, _ := w.Raw(w)
	if string(rawData) != `{"name":"Hemanta","age":29}` {
		t.Error("Raw data does not match correct JSON")
	}

	w1 := &TestObject{}
	w1.Restore(w1)

	t.Log(w1)

	if w1.Name != "Hemanta" {
		t.Error("Error saving/restore DBObject: Name")
	}

	if w1.Age != 29 {
		t.Error("Error saving/restore DBObject: Age")
	}

	// Test deletion
	w1.Delete(w1)

	w1 = &TestObject{}
	err = w1.Restore(w1)
	if err == nil {
		t.Error("After deleting an obect, it shouldn't exist")
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

	cache.Delete(w1)

	if len(cache.Objects) != 0 {
		t.Log("After deletion, cache should have been empty")
	}
}

func TestCacheSync(t *testing.T) {
	cache := GetAppCache()

	w := &TestObject{
		Name: "Fitmanta",
		Age:  29,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		w.Lock()
		defer w.Unlock()
		defer wg.Done()

		w.Name = "Strictmanta"

		cache.Put(w)
	}()
	go func() {
		w.Lock()
		defer w.Unlock()
		defer wg.Done()

		w.Name = "Litmanta"

		cache.Put(w)
	}()
	wg.Wait()

	w1 := cache.Get(&TestObject{}).(*TestObject)

	if w1.Name != "Litmanta" {
		t.Log("Sync didn't work.")
	}

	// t.Log(w1)

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

func TestConcurrentSave(t *testing.T) {

	w := &TestObject{
		Name: "Hemanta",
		Age:  29,
	}

	t.Log(w.Save(w))

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		w.Lock()
		defer w.Unlock()
		defer wg.Done()

		w.Name = "Fitmanta"
		w.Save(w)
	}()

	go func() {
		w.Lock()
		defer w.Unlock()
		defer wg.Done()

		w.Name = "Yomanta"
		w.Save(w)
	}()

	go func() {
		w.Lock()
		defer w.Unlock()
		defer wg.Done()

		w.Name = "Romanta"
		w.Save(w)
	}()

	wg.Wait()

}
