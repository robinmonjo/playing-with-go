package mongo

import (
	"labix.org/v2/mgo"
	"os"
	"fmt"
	"net/url"
)

//constants
const bubbleSortCollection = "bubble_sort"

//globa
var session *mgo.Session
var dbName string //given by the mongodb URI

func InitSession() {
	if (session != nil) {
		fmt.Println("[WARNING] - Trying to init an already initialised MongoDB session")
		return
	}
	var err error
	session, err = mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if (err != nil) {
		panic(err)
	}
	parsed, err := url.Parse(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		panic(err)
	}
	dbName = parsed.Path[1:]
}

func CloseSession() {
	session.Close()
	fmt.Println("[INFO] - session closed")
}

func InsertBubbleSortResult(jsonResult map[string]interface{}) {
	fmt.Println("[INFO] - inserting", jsonResult)
	c := session.DB(dbName).C(bubbleSortCollection)
	err := c.Insert(jsonResult)
  if err != nil {
		panic(err)
	}
}


