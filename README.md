# Goma - ごま #

With Goma, you can write your mobile app's business logic in GO and reuse them across platforms like IOS and Android. It comes with an embedded database that allows you to save/restore your application's objects.

Goma utilizes [gomobile](https://godoc.org/golang.org/x/mobile/cmd/gomobile) to generate platform specific libraries for your apps.

# Installation

Get or update **Goma** with:

` go get -u github.com/hemantasapkota/goma`

# Concepts & Usage

## goma.Object ##

Anything that is persitable can be modeled with a Goma object. It can be your app's **view**, **model** or **both**.

A Goma object can be defined by embedding [goma.Object](object.go) into your struct. It should also implement [goma.DBOBject](object.go) interface to make it persistable.

```go
type Person struct {
  *goma.Object
   Name string `json:"name"`
   Age int `json:"age"`
}

// Implement goma.DBObject interface's Key() method
func (p Person) Key() string {
  return "sampleApp.Person"
}
```

# Examples

See https://github.com/hemantasapkota/goma-examples
