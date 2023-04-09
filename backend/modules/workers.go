package workers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Entry struct {
	Title     string
	Class     string
	Clearance int
	//Procedures  string
	//Description string
	Content  string
	Addendum string
	Image    string
}

func scrape(query string) Entry {
	var entry Entry
	//entry = Entry{Title: "SCP-001", Class: "Euclid", Content: "Testing"}

	c := colly.NewCollector(
		colly.AllowedDomains("scp-wiki.wikidot.com"),
	)

	c.OnHTML("div", func(e *colly.HTMLElement) {
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
				if h.DOM.Parent().Is("blockquote") {
					return
				}

				if strings.Contains(h.Text, "Object Class") {
					entry.Class = strings.Replace(h.Text, "Object Class: ", "", -1)
				} else {
					entry.Content += "\n" + h.Text
				}
				/* else if strings.Contains(h.Text, "Special Containment Procedures") {
					entry.Procedures = strings.Replace(h.Text, "Special Containment Procedures: ", "", -1)
				} else if strings.Contains(h.Text, "Description") {
					entry.Description = strings.Replace(h.Text, "Description: ", "", -1)
				}*/
			})
			e.ForEach("img", func(i int, h *colly.HTMLElement) {
				if h.Attr("class") == "image" {
					entry.Image = h.Attr("src")
				}
			})
			e.ForEach("blockquote", func(i int, h *colly.HTMLElement) {
				h.ForEach("p", func(i int, j *colly.HTMLElement) {
					entry.Addendum += "\n" + j.Text
				})
			})
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong while fetching the requested data...")
	})

	if query == "random" {
		var scpMax int = 7999
		var random string = strconv.Itoa(rand.Intn(scpMax))
		fmt.Println(random)
		c.Visit("https://scp-wiki.wikidot.com/scp-" + random)
	} else {
		c.Visit("https://scp-wiki.wikidot.com/scp-" + query)
	}

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
