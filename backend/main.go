package main

import (
	"fmt"
	"log"
	"net/http"

	workers "backend/main/modules"
)

func fetchPage(w http.ResponseWriter, r *http.Request) {
	var query string = r.URL.Query().Get("query")

	var fetched string = workers.Fetch(query)
	fmt.Println(fetched)
}

func main() {
	http.HandleFunc("/fetch", fetchPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
