// connecting to a mongoDB database
// simple ping test + error handling
// with panics and a deffered close

package main

import (
    "fmt"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

type Person struct {
    Name string
    Phone string
}

func main() {
    session, err := mgo.Dial("127.0.0.1:27017")  //default port
    if err != nil {
        panic("conn Err")
    }

    defer session.Close()


    session.SetMode(mgo.Monotonic, true)
    c := session.DB("test").C("Testing")
    err = c.Insert(&Person{"Ian", "079324211"},
                   &Person{"Alex", "0701234567"})


    if err != nil {
        panic(err)
    }


    result := Person{}
    err = c.Find(bson.M{"name": "Alex"}).One(&result)
    if err != nil {
        panic(err)
    }

    fmt.Println("Phone", result.Phone)
    fmt.Println("Name", result.Name)
}
