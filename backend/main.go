package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	workers "github.com/NotASimp69/SCP-Site/backend/modules"
)

func fetchPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var query string = r.URL.Query().Get("query")

	if query == "" {
		fetched := "Please provide a valid search query"
		jsonResp, _ := json.Marshal(fetched)
		w.Write(jsonResp)
		return
	} else {
		fetched := workers.Fetch(query)
		fmt.Println("Fetched ", fetched)

		jsonResp, err := json.Marshal(fetched)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func main() {
	http.HandleFunc("/fetch", fetchPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
