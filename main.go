package main

import (
	"fmt"
	"log"
	"net/http"
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

func main() {
	// Request the HTML page.
	res, err := http.Get("https://github.com/trending")
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
		// repセルの横並び属性が25件
		s.Children().Each(func(j int, ns *goquery.Selection) {
			// 横並び属性を（最大）5分割したもの
			first := ns.First()
			{
				attr, _ := first.Children().Attr("class")
				text := strings.TrimSpace(ns.Children().Last().Text())
				if attr == "repo-language-color ml-0" && text != "" {
					languages[i] = text
				} else if attr == "oction oction-star" && text != "" {
					stars[i] = text
				} else if attr == "oction oction-fork" && text != "" {
					forks[i] = text
				}
				fmt.Printf("%d/%d: %s\n", i, j, text)
			}

			second := first.Next()
			{
				attr, _ := second.Children().Attr("class")
				text := strings.TrimSpace(ns.Children().Last().Text())
				if attr == "repo-language-color ml-0" && text != "" {
					languages[i] = text
				} else if attr == "oction oction-star" && text != "" {
					stars[i] = text
				} else if attr == "oction oction-fork" && text != "" {
					forks[i] = text
				}
				fmt.Printf("%d/%d: %s\n", i, j, text)
			}

			third := second.Next()
			{
				attr, _ := third.Children().Attr("class")
				text := strings.TrimSpace(ns.Children().Last().Text())
				if attr == "repo-language-color ml-0" && text != "" {
					languages[i] = text
				} else if attr == "oction oction-star" && text != "" {
					stars[i] = text
				} else if attr == "oction oction-fork" && text != "" {
					forks[i] = text
				}
				fmt.Printf("%d/%d: %s\n", i, j, text)
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
