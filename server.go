package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"playing-with-go/packages/sortAlgorithms"
)

func main() {
  http.HandleFunc("/", rootHandler)
	http.HandleFunc("/sort/", sortHandler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
	  panic(err)
	}
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome")
}

func sortHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(res, "Only allowing POST on /sort/")
		return
	}
	
	//reading body (as []byte)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(res, "[ERROR] parsing body : %s", err)
		return
	}
	
	//conerting json input into slice
	var slice []float64
  err = json.Unmarshal(body, &slice) 
	if err != nil {
		fmt.Fprintf(res, "[ERROR] unmarshalling json body : %s", err)
		return
	}
	
	//calling algo
	sortAlgorithms.BubleSort(slice)
	
	//sending back json
	b, err := json.Marshal(slice)
	if err != nil {
		fmt.Fprintf(res, "[ERROR] marshalling json body : %s", err)
		return
	}
	fmt.Fprintf(res, "%s", string(b))
}


