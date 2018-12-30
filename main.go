package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

	child := trends.Children()

	titles := []string{}
	descriptions := []string{}
	languages := []string{}
	stars := []string{}
	forks := []string{}
	today := []string{}

	child.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Each(func(i int, s *goquery.Selection) {
		titles = append(titles, strings.TrimSpace(s.Text()))
	})
	child.Find("div.py-1").Each(func(i int, s *goquery.Selection) {
		descriptions = append(descriptions, strings.TrimSpace(s.Text()))
	})
	divf6 := child.Find("div.f6.text-gray.mt-2")
	divf6.Each(func(i int, s *goquery.Selection) {
		t := s.Find("span:nth-child(1)").Text()
		languages = append(languages, strings.TrimSpace(t))
	})
	divf6.Each(func(i int, s *goquery.Selection) {
		t := s.Find("a:nth-child(2)").Text()
		stars = append(stars, strings.TrimSpace(t))
	})
	divf6.Each(func(i int, s *goquery.Selection) {
		t := s.Find("a:nth-child(3)").Text()
		forks = append(forks, strings.TrimSpace(t))
	})
	divf6.Each(func(i int, s *goquery.Selection) {
		t := s.Find("span.d-inline-block.float-sm-right").Text()
		today = append(today, strings.TrimSpace(t))
	})

	for i := 0; i < len(titles); i++ {
		fmt.Print("------------------------------\n")
		fmt.Printf("[Title] %s\n", titles[i])
		fmt.Printf("[Description] %s\n", descriptions[i])
		fmt.Printf("[Language] %s\n", languages[i])
		fmt.Printf("[Stars] %s\n", stars[i])
		fmt.Printf("[Forks] %s\n", forks[i])
		fmt.Print(today[i])
		fmt.Print("\n")
	}
}
