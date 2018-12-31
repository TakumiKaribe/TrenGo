package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

const githubURL string = "https://github.com"
const trending string = "/trending"

var titles [50]string
var URLs [50]string
var descriptions [50]string
var languages [50]string
var stars [50]int
var forks [50]int
var rangeStar [50]int

func main() {
	var (
		lang string
		d    bool
		w    bool
		m    bool
	)
	flag.StringVar(&lang, "l", "", "string flag")
	flag.BoolVar(&d, "d", false, "daily search")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")

	flag.Parse()
	if lang != "" {
		lang = "/" + lang
	}

	rangeType := parseRangeType(d, w, m)

	// Request the HTML page.
	res, err := http.Get(githubURL + trending + lang + "?since=" + rangeType.QueryString())
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

	// TODO: parse周りは別ファイルに切り出したい
	trends := doc.Find("body > div.application-main > div.explore-pjax-container.container-lg.p-responsive.clearfix > div > div.col-md-9.float-md-left > div.explore-content > ol")

	trends.Each(func(i int, s *goquery.Selection) {
		s.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Each(func(i int, s *goquery.Selection) {
			titles[i] = strings.TrimSpace(s.Text())
			URLs[i] = strings.TrimSpace(s.AttrOr("href", "default"))
		})
		s.Find("div.py-1").Each(func(i int, s *goquery.Selection) {
			descriptions[i] = strings.TrimSpace(s.Text())
		})
	})

	divf6 := trends.Find("div.f6.text-gray.mt-2")
	divf6.Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, ns *goquery.Selection) {
			text := strings.TrimSpace(ns.Text())
			attr := strings.TrimSpace(ns.Children().AttrOr("class", "default"))
			if attr == "repo-language-color ml-0" && text != "" {
				languages[i] = text
			} else if attr == "octicon octicon-star" && text != "" {
				if strings.TrimSpace(ns.Children().AttrOr("aria-label", "default")) == "star" {
					v, err := strconv.Atoi(strings.Replace(text, ",", "", -1))
					if err == nil {
						stars[i] = v
					}
				}
			} else if attr == "octicon octicon-repo-forked" && text != "" {
				v, err := strconv.Atoi(strings.Replace(text, ",", "", -1))
				if err == nil {
					forks[i] = v
				}
			}
		})
	})

	divf6.Each(func(i int, s *goquery.Selection) {
		ns := s.Find("span.d-inline-block.float-sm-right")
		text := strings.TrimSpace(ns.Text())
		if text == "" {
			rangeStar[i] = 0
		} else {
			splited := strings.Split(text, " ")
			star, err := strconv.Atoi(splited[0])
			if err == nil {
				rangeStar[i] = star
			}
		}
	})

	fmt.Println("[")
	for i := 0; i < divf6.Length(); i++ {
		fmt.Println("{")
		fmt.Printf("    title: %s,\n", "\""+titles[i]+"\"")
		fmt.Printf("    url: %s,\n", "\""+githubURL+URLs[i]+"\"")
		fmt.Printf("    description: %s,\n", "\""+descriptions[i]+"\"")
		fmt.Printf("    language: %s,\n", "\""+languages[i]+"\"")
		fmt.Printf("    sumStars: %d,\n", stars[i])
		fmt.Printf("    forks: %d,\n", forks[i])
		fmt.Printf("    ranged: {\n")
		fmt.Printf("        type: %s,\n", "\""+rangeType.QueryString()+"\"")
		fmt.Printf("        stars: %d\n", rangeStar[i])
		fmt.Printf("    }\n")
		fmt.Print("}")
		if i < len(titles)-1 {
			fmt.Print(",\n")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("]")
}
