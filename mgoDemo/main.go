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
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

const cName = "testMongoDemo"

func main() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	upSertDemo(session)
	bulkInsertDemo(session)
}

func upSertDemo(session *mgo.Session){
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

	printObject(changeInfo, "changeInfo=")
}

func bulkInsertDemo(session *mgo.Session){
	len := 1000
	var data = make([]interface{}, len)
	var i uint = 1

	for ; i < uint(len); i++ {

		str := strconv.FormatUint(uint64(i),10)

		title := "title" + str
		name := "name" + str

		data = append(data, struct{
			Title string
			Name string
			Count uint
		}{
			title,
			name,
			i,
		})
	}

	bulkInsert(session, data)

}

func bulkInsert(session *mgo.Session, data []interface{}) {
	collection := session.DB(cName).C("bulk_insert_demo")

	bulk := collection.Bulk()
	bulk.Unordered() // this is magic part for successful bulk insert!!!
	bulk.Insert(data...)
	bulkResult, bulkErr := bulk.Run()
	if bulkErr != nil {
		panic(bulkErr)
	}

	logPrintObject("bulkResult=%s",bulkResult)

}
func logPrintObject(fmtStr string, o interface{}) {
	objString, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		log.Printf("marshal to json error:%s", err.Error())
	}

	log.Printf(fmtStr, string(objString))
}

func printObject(o interface{}, prefix string) {
	objString, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		fmt.Printf("marshal to json error:%s", err.Error())
	}
	fmt.Println(prefix, string(objString))
}
