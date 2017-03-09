package main

/*
NOTE:
 example is from http://www.mrwaggel.be/post/golang-mgo-findandmodify-auto-increment-id/
 to run this example,
 1. you need to have a connectable mongoDB which has a "testMongoDemo" collection on localhost,

 2. run following command to first add data:
 use testMongoDemo
 db.persons.insert({"_id": "counterForPersons", "counterValue": 0})

 3. install mgo & bson package:
 go get -u -v gopkg.in/mgo.v2
 go get -u -v gopkg.in/mgo.v2/bson

*/

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const cName = "testMongoDemo"

func main() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	//select the "sandboxed" collection in mongoDB
	collection := session.DB(cName).C("persons")

	//Create a bson.M interface to store result to
	var result bson.M

	//This is to tell what field and how to modify it
	changeOp := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counterValue": 1}},
		ReturnNew: true,
	}

	//Find the document that keeps the counter value, and apply the modification
	changeInfo, err := collection.Find(bson.M{"_id": "counterForPersons"}).Limit(1).Apply(changeOp, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println("new increment Value is", result["counterValue"])

	changeInfoStr, err := json.MarshalIndent(changeInfo, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("changeInfo=", string(changeInfoStr))

}
