package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type Person struct {
	Name   string        `json:"name"`
	DocId  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Id     string        `json:"userid"`
	Gender string        `json:"gender"`
}

func main() {

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   []string{"localhost:27017"},
		Timeout: time.Second * 60,
	})

	if err != nil {
		fmt.Printf("[ERROR] error while creating session %s\n", err)
	}

	c := session.DB("admin").C("person")

	var p = Person{
		DocId:  bson.ObjectIdHex("5b8d619ca8a8b005487000e3"),
		Name:   "phanikumar",
		Id:     "123",
		Gender: "male",
	}
	update := bson.M{"$set": p}
	info, err := c.UpsertId(p.DocId, update)

	if err != nil {
		fmt.Printf("An Error occured %s\n", err)
	} else {
		fmt.Printf("Upserted id is %v\n", &info.UpsertedId)
		fmt.Printf("If updated %d\n", &info.Updated)
	}
	// c.Insert(&Person{
	// 	Name:   "phanikuh",
	// 	Id:     "123",
	// 	Gender: "male",
	// })

	// result := Person{}
	// query := c.Find(bson.M{"name": "phanik"})
	// err1 := query.One(&result)

	// if err1 != nil {
	// 	fmt.Printf("An error occures %s\n", err1)
	// }

	// fmt.Println(result)
}
