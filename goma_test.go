package goma

import (
	gomadb "github.com/hemantasapkota/goma/gomadb"
	ldb "github.com/hemantasapkota/goma/gomadb/leveldb"

	"sync"
	"testing"
)

// Container Object
type ContainerItem struct {
	Id int `json:"id"`
}

type Container struct {
	*Object
	sync.Mutex

	Children []ContainerItem
}

func (c *Container) Key() string {
	return "goma.container"
}

func (c *Container) AddChild(child ContainerItem) {
	c.Lock()
	defer c.Unlock()

	c.Children = append(c.Children, child)
}

// Test Object
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

func TestContainerObjectSync(t *testing.T) {
	container := &Container{}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		container.AddChild(ContainerItem{Id: 1})
	}()

	go func() {
		defer wg.Done()

		container.AddChild(ContainerItem{Id: 2})
	}()

	wg.Wait()

	container.Save(container)

}
