package workers

import (
	"fmt"
	"time"
)

type Entry struct {
	Title   string `json:"Title"`
	Class   string `json:"Class"`
	Content string `json:"Content"`
}

func scrap(query string) Entry {
	var entry Entry

	entry = Entry{Title: "SCP-001", Class: "Euclid", Content: "Testing"}

	fmt.Println(entry)
	return entry
}

func Fetch(query string) string { // has to be uppercase to be public
	fmt.Println("Initializing fetch...")
	fmt.Println(query)
	var data string

	go scrap(query)

	time.Sleep(5 * time.Second)
	return data
}
