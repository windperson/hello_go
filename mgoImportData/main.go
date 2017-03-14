package main

/*
NOTE:
 0. test data is original from:
 https://docs.mongodb.com/getting-started/shell/import-data/#retrieve-the-restaurants-data
 then using mongoimport, mongoexport to get MongoDB extened JSON format file sample.

 1. you need to have a connectable mongoDB which has a "testMongoDemo" collection on localhost,

 2. install mgo & bson package:
 go get -u -v gopkg.in/mgo.v2
 go get -u -v gopkg.in/mgo.v2/bson

*/

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const cName = "testMongoDemo"

func main() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	data, err := loadMongoExportJSON(dir + "/data/mongo_exported.json")

	if err != nil {
		panic(err)
	}

	bulkInsert(session, data)
}

func loadMongoExportJSON(path string) (data []interface{}, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		var lineJSON map[string]interface{}
		err := bson.UnmarshalJSON(fscanner.Bytes(), &lineJSON) // mongodb special json type: https://docs.mongodb.com/manual/reference/mongodb-extended-json/
		if err != nil {
			panic(err)
		}
		//log.Printf("%v\n", lineJSON)
		data = append(data, lineJSON)
	}

	return
}

func bulkInsert(session *mgo.Session, data []interface{}) {
	collection := session.DB(cName).C("bulk_insert_bson_demo")

	collection.RemoveAll(nil) // we need to clear existing data to prevent id collision!

	bulk := collection.Bulk()
	bulk.Unordered() // this is magic part for successful bulk insert!!!
	bulk.Insert(data...)
	bulkResult, bulkErr := bulk.Run()

	if bulkErr != nil {
		panic(bulkErr)
	}

	logPrintObject("bulkResult=%s", bulkResult)

}
func logPrintObject(fmtStr string, o interface{}) {
	objString, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		log.Printf("marshal to json error:%s", err.Error())
	}

	log.Printf(fmtStr, string(objString))
}
