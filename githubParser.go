package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var lang string
var rangeType RangeType

var developers [50]string
var names [50]string
var URLs [50]string
var descriptions [50]string
var languages [50]string
var stars [50]int
var forks [50]int
var rangeStar [50]int
var length int

func parseGitHub() {
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

	trends := doc.Find("body > div.application-main > div.explore-pjax-container.container-lg.p-responsive.clearfix > div > div.col-md-9.float-md-left > div.explore-content > ol")

	trends.Each(func(i int, s *goquery.Selection) {
		s.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Each(func(i int, s *goquery.Selection) {
			text := strings.Split(strings.Replace(s.Text(), " ", "", -1), "/")
			developers[i] = strings.TrimSpace(text[0])
			names[i] = strings.TrimSpace(text[1])
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

	length = divf6.Length()
}

func printGitHub(isJSONFormat bool) {
	if isJSONFormat {
		fmt.Println("[")
		for i := 0; i < length; i++ {
			fmt.Println("{")
			fmt.Printf("  developer: %s,\n", "\""+developers[i]+"\"")
			fmt.Printf("  name: %s,\n", "\""+names[i]+"\"")
			fmt.Printf("  url: %s,\n", "\""+githubURL+URLs[i]+"\"")
			fmt.Printf("  description: %s,\n", "\""+descriptions[i]+"\"")
			fmt.Printf("  language: %s,\n", "\""+languages[i]+"\"")
			fmt.Printf("  sumStars: %d,\n", stars[i])
			fmt.Printf("  forks: %d,\n", forks[i])
			fmt.Printf("  ranged: {\n")
			fmt.Printf("    type: %s,\n", "\""+rangeType.QueryString()+"\"")
			fmt.Printf("    stars: %d\n", rangeStar[i])
			fmt.Printf("  }\n")
			fmt.Print("}")
			if i < length-1 {
				fmt.Print(",\n")
			} else {
				fmt.Println()
			}
		}
		fmt.Println("]")

	} else {
		fmt.Printf("========== GitHub Trending ==========\n\n")
		for i := 0; i < length; i++ {
			fmt.Println("★-----★-----★-----★-----★-----★-----★-----★-----★-----★-----★")
			fmt.Printf("  [rank] %d\n", i+1)
			fmt.Printf("  [developer] %s\n", developers[i])
			fmt.Printf("  [name] %s\n", names[i])
			fmt.Printf("  [url] %s\n", githubURL+URLs[i])
			fmt.Printf("  [description] %s\n", descriptions[i])
			fmt.Printf("  [language] %s\n", languages[i])
			fmt.Printf("  [sumStars] %d\n", stars[i])
			fmt.Printf("  [forks] %d\n", forks[i])
			fmt.Printf("  [trend] %d stars %s\n", rangeStar[i], rangeType.QueryString())
			fmt.Println()
		}
	}
}
