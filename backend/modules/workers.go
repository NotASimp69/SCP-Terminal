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

func Fetch() string { // has to be uppercase to be public
	fmt.Println("Initializing fetch...")
	var data string

	time.Sleep(2 * time.Second)
	return data
}
