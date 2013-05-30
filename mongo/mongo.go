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
var session *mgo.Session

func InitSession() {
	if (session != nil) {
		fmt.Println("[WARNING] - Trying to init an already initialised MongoDB session")
		return
	}
	var err error
	session, err = mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if (err != nil) {
		panic(err)
	}
	fmt.Println("[INFO] - session opened")
}

func CloseSession() {
	session.Close()
	fmt.Println("[INFO] - session closed")
}

func InsertBubbleSortResult(jsonResult map[string]interface{}) {
	fmt.Println("[INFO] - inserting", jsonResult)
	c := session.DB(dbName).C(bubbleSortCollection)
	err := c.Insert(&jsonResult)
  if err != nil {
		panic(err)
	}
}


