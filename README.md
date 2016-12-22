# Goma - ごま #
> Build native mobile apps faster by writing your application logic in GO

With Goma, you can write your mobile app's business logic in GO and reuse them across platforms like IOS and Android. It comes with an embedded database that allows you to save/restore application's objects.

Goma utilizes [gomobile](https://godoc.org/golang.org/x/mobile/cmd/gomobile) to generate platform specific libraries for your apps.

## Requirements

- At least Go 1.5.3 or above
- Xcode
- Android Studio or Eclipse

# Installation

Get or update **Goma** with:

` go get -u github.com/hemantasapkota/goma`

# Concepts & Usage

## goma.Object - Definition & Construction ##

Anything that is persitable can be modeled with a Goma object. It can be your app's **view**, **model** or **both**.

A complete goma object should embed [goma.Object](object.go) and implement the [goma.DBObject](object.go) interface.

```go
type Person struct {
  *goma.Object
   Name string `json:"name"`
   Age int `json:"age"`
}

// Implement goma.DBObject interface's Key() method
func (p *Person) Key() string {
  return "sampleApp.Person"
}
```

## goma.Object - Saving to the database ##

```go

person := &Peron{
 Name: "Jon Panda",
 Age: 20,
}

err := person.Save(person)
if err != nil {
 // do something
}

```

The object is serialized to JSON and stored in DB.

```
{"name":"Jon Panda","age":20}
```

## goma.Object - Restoring from the database ##

```go
person := &Person{}
err := person.Restore(person)

if err != nil {
 // do something
}
```

## goma.AppCache - In-memory cache for your objects ##

Saving and restoring objects are expensive functions, so they must be minimized. Goma comes with an in-memory cache which can be used for storing objects throughout the session of the app.

### Put to cache ###
```go
person := &Person{
 Name: "Panda panda",
 Age: 35,
}
goma.GetAppCache().Put(person)
```

### Get from cache ###
```go
p := goma.GetAppCache().Get(&Person{}).(*Person)

// Do something with p
```

The **Get** method expects an empty struct for the type you're querying for. This struct should also implement [goma.DBObject](object.go) interface.  If the queried object does not exists, **Get** returns the same empty type that was passed in. 

Since **Get** is likely to be called from multiple places in the codebase, this design helps avoid making nil error checks ( `if obj != nil` ) on the returned object. If there's a better design for this, feel free to submit a pull request.

Here's an another example:

```go
type ContainerItem struct {
    Id int `json:"id"`
}

type Container struct {
    *goma.Object
    Children []ContainerItem `json:"children"`
}

func (c *Container) Key() string {
    return "sampleApp.Container"
}

func EmptyConatiner() *Container {
    return &Container{
        Childred: make([]ContainerItem, 0),
    }
}

container := goma.GetAppCache().Get(EmptyContainer()).(*Container)

// If the object does not exist in the cache, the method returns the object returned by EmptyContainer()

```

## App Structure ( TODO )

The following structure is reccomended:

## Marshaling - Return values from Goma ##

[gomobile](https://godoc.org/golang.org/x/mobile/cmd/gomobile) has some restrictions as to what types you can return. They are primitives ( ```int, int64, float float64, string, and bool``` ) and ( byte array ) ```[]byte```. 

Goma leverages Go's excellent JSON support. Complex types like structs, arrays, and maps can be easily marshalled into JSON.

# Wiki
* [Interfacing with Android and IOS](https://github.com/hemantasapkota/goma/wiki/Interfacing-with-Android-and-IOS)

# Examples

* [Meal Tracker](https://github.com/hemantasapkota/goma-examples/tree/master/fitness) - A simple meal tracker that I use personally

See https://github.com/hemantasapkota/goma-examples

# Production Apps

Goma is being used in production for the following apps.

* **OpenLearning IOS** - https://itunes.apple.com/us/app/openlearning/id981790180?ls=1&mt=8
* **OpenLearning Android** - https://play.google.com/store/apps/details?id=openlearning.com.openlearning

## Contribute

We would love for you to contribute to **Goma**, check the ``LICENSE`` file for more info.

## Meta

Hemanta Sapkota – [@ozhemanta](https://twitter.com/ozhemanta) – laex.pearl@gmail.com

Distributed under the MIT license. See ``LICENSE`` for more information.

[https://github.com/hemantasapkota](https://github.com/hemantasapkota/)
