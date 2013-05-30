package mongo

import (
	"labix.org/v2/mgo"
	"os"
	"fmt"
)

//constants
const dbName = "db_sort"
const bubbleSortCollection = "bubble_sort"

//global session
//var session *mgo.Session


func InsertBubbleSortResult(jsonResult map[string]interface{}) {
	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if (err != nil) {
		panic(err)
	}
	defer session.Close()
	fmt.Println("[INFO] - inserting", jsonResult)
	c := session.DB(dbName).C(bubbleSortCollection)
	fmt.Println("[INFO] - collection", c)
	err = c.Insert(jsonResult)
  if err != nil {
		panic(err)
	}
}


