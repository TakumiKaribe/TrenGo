package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RangeType int

const (
	today RangeType = iota
	thisWeek
	thisMonth
)

func (rt RangeType) String() string {
	switch rt {
	case today:
		return "daily"
	case thisWeek:
		return "weekly"
	case thisMonth:
		return "monthly"
	default:
		return "unknown"
	}
}

func parse(sel *goquery.Selection, query string, isNested bool) []string {
	ret := []string{}
	if isNested {
		sel.Each(func(i int, s *goquery.Selection) {
			ns := s.Find(query)
			ret = append(ret, strings.TrimSpace(ns.Text()))
		})
	} else {
		sel.Find(query).Each(func(i int, s *goquery.Selection) {
			ret = append(ret, strings.TrimSpace(s.Text()))
		})
	}
	return ret
}

const trendURL string = "https://github.com/trending"

func main() {
	var lang string
	flag.StringVar(&lang, "l", "", "string flag")
	flag.Parse()
	if lang != "" {
		lang = "/" + lang
	}

	// Request the HTML page.
	res, err := http.Get(trendURL + lang + "?since=" + thisWeek.String())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	trends := doc.Find("body > div.application-main > div.explore-pjax-container.container-lg.p-responsive.clearfix > div > div.col-md-9.float-md-left > div.explore-content > ol")

	titles := []string{}
	descriptions := []string{}
	trends.Each(func(i int, s *goquery.Selection) {
		titles = parse(s, "div.d-inline-block.col-9.mb-1 > h3 > a", false)
		descriptions = parse(s, "div.py-1", false)
	})

	languages := []string{}
	stars := []string{}
	forks := []string{}
	for i := 0; i < len(titles); i++ {
		languages = append(languages, "NOT EXIST")
		stars = append(stars, "NOT EXIST")
		forks = append(forks, "NOT EXIST")
	}

	trends.Find("div.f6.text-gray.mt-2").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, ns *goquery.Selection) {
			text := strings.TrimSpace(ns.Text())
			attr := strings.TrimSpace(ns.Children().AttrOr("class", "default"))
			if attr == "repo-language-color ml-0" && text != "" {
				languages[i] = text
			} else if attr == "octicon octicon-star" && text != "" {
				if strings.TrimSpace(ns.Children().AttrOr("aria-label", "default")) == "star" {
					stars[i] = text
				}
			} else if attr == "octicon octicon-repo-forked" && text != "" {
				forks[i] = text
			}
		})
	})

	today := parse(trends.Find("div.f6.text-gray.mt-2"), "span.d-inline-block.float-sm-right", true)

	for i := 0; i < len(titles); i++ {
		fmt.Print("\n------------------------------\n")
		fmt.Printf("[Title] %s\n", titles[i])
		fmt.Printf("[Description] %s\n", descriptions[i])
		fmt.Printf("[Language] %s\n", languages[i])
		fmt.Printf("[Stars] %s\n", stars[i])
		fmt.Printf("[Forks] %s\n", forks[i])
		fmt.Print(today[i])
		fmt.Print("\n")
	}
}
