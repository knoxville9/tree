package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
)

type Goods struct {
	Title string   `json:"title"`
	Types string   `json:"types"`
	Price string   `json:"price"`
	Img   []string `json:"img"`
	Url   string   `json:"url"`
}

func main() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36"))
	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
	})
	//
	//c.OnRequest(func(request *colly.Request) {
	//	log.Println("start visit: ", request.URL.String())
	//})

	c.OnError(func(response *colly.Response, err error) {
		log.Println(err.Error())
	})

	for i := 1; i < 10; i++ {
		url := "http://bbs.badmintoncn.com/forum.php?mod=forumdisplay&fid=64&page=" + strconv.Itoa(i)
		c.OnHTML(".xst", func(e *colly.HTMLElement) {
			attr := e.Attr("href")

			detail(c, e.Request.AbsoluteURL(attr))
		})
		c.Visit(url)

	}

}

func detail(c *colly.Collector, url string) {

	c = c.Clone()

	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
	})

	c.OnHTML("#postlist", func(e *colly.HTMLElement) {

		title := e.DOM.Find("#thread_subject").Text()
		types := e.DOM.Find("tbody > tr:nth-child(1) > td.plc > div.pct > div > table > tbody > tr:nth-child(1) > td:nth-child(2)").Text()
		price := e.DOM.Find("tbody > tr:nth-child(1) > td.plc > div.pct > div > table > tbody > tr:nth-child(4) > td:nth-child(2)").Text()
		pics := []string{}
		url := e.Request.URL.String()
		e.ForEach("ignore_js_op", func(i int, e *colly.HTMLElement) {
			attr, exists := e.DOM.Find("img").Attr("file")
			if exists {
				pics = append(pics, attr)

			}
		})
		g := &Goods{
			Title: title,
			Types: types,
			Price: price,
			Img:   pics,
			Url:   url,
		}
		marshal, err := json.Marshal(g)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println(string(marshal))

		f, err := os.OpenFile("1.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		if err != nil {
			fmt.Printf("Cannot open file %s!\n", "1.txt")
			return
		}
		defer f.Close()
		f.WriteString(string(marshal))
		f.WriteString("\n")

	})

	c.Visit(url)
}
