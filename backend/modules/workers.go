package workers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Entry struct {
	Title       string
	Class       string
	Clearance   int
	Procedures  string
	Description string
	Image       string
}

func scrape(query string) Entry {
	var entry Entry
	//entry = Entry{Title: "SCP-001", Class: "Euclid", Content: "Testing"}

	c := colly.NewCollector(
		colly.AllowedDomains("scp-wiki.wikidot.com"),
	)

	c.OnHTML("div", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		if e.Attr("id") == "page-title" {
			title, err := e.DOM.Html()
			if err == nil {
				title = strings.TrimPrefix(title, "\n")
				title = strings.TrimSuffix(title, "\n")
				title = strings.TrimSpace(title)
				entry.Title = title
			} else {
				entry.Title = "Error"
			}
		} else if e.Attr("id") == "page-content" {
			e.ForEach("p", func(i int, h *colly.HTMLElement) {
				if strings.Contains(h.Text, "Object Class") {
					entry.Class = strings.Replace(h.Text, "Object Class: ", "", -1)
				} else if strings.Contains(h.Text, "Special Containment Procedures") {
					entry.Procedures = strings.Replace(h.Text, "Special Containment Procedures: ", "", -1)
				} else if strings.Contains(h.Text, "Description") {
					entry.Description = strings.Replace(h.Text, "Description: ", "", -1)
				}
			})
			e.ForEach("img", func(i int, h *colly.HTMLElement) {
				fmt.Println(i)
				fmt.Println(h)
				if h.Attr("class") == "image" {
					entry.Image = h.Attr("src")
				}
			})
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalln("Something went wrong while fetching the requested data...")
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://scp-wiki.wikidot.com/" + query)

	return entry
}

func Fetch(query string) Entry { // has to be uppercase to be public
	fmt.Println("Initializing fetch...")

	var data Entry
	c := make(chan Entry)

	go func() {
		c <- scrape(query)
	}()

	data = <-c

	//time.Sleep(5 * time.Second)
	return data
}
