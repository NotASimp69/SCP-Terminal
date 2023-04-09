package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	workers "github.com/NotASimp69/SCP-Site/backend/modules"
)

func fetchPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var query string = r.URL.Query().Get("query")
	var final_query string

	if query == "" {
		fetched := "Please provide a valid search query"
		jsonResp, _ := json.Marshal(fetched)
		w.Write(jsonResp)
		return

	}

	if query == "random" { // only sanitize query if its not specified as random!
		final_query = query

	} else {
		for _, char := range query {
			var c string = string(char)
			i, err := strconv.Atoi(c)
			if err == nil {
				final_query += strconv.Itoa(i)
			}
		}
	}

	fetched := workers.Fetch(final_query)

	jsonResp, err := json.Marshal(fetched)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/fetch", fetchPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
