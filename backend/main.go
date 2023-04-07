package main

import (
	"fmt"
	"log"
	"net/http"

	workers "backend/main/modules"
)

func fetchPage(w http.ResponseWriter, r *http.Request) {

	//var fetched = workers.Entry{Title: "SCP-001", Class: "Euclid", Content: "Testing"}
	var fetched string = workers.Fetch()
	fmt.Println(fetched)
}

func main() {
	http.HandleFunc("/fetch", fetchPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
