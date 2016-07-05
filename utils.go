package goma

import (
	"encoding/json"
	"time"
)

func Timestamp() string {
	return time.Now().Format(time.RFC3339)
}

func TimestampFrom(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseDatetime(iso8601date string) time.Time {
	t, err := time.Parse(time.RFC3339, iso8601date)
	if err != nil {
		Log(err)
		return time.Now().In(time.Local)
	}
	return t.In(time.Local)
}

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
