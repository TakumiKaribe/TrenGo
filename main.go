package main

import (
	"flag"
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

func main() {
	var (
		lang string
		d    bool
		w    bool
		m    bool
		j    bool
	)
	flag.StringVar(&lang, "l", "", "string flag")
	flag.BoolVar(&d, "d", false, "daily search")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")
	flag.BoolVar(&j, "j", false, "monthly search")

	flag.Parse()
	if lang != "" {
		lang = "/" + lang
	}

	rangeType = parseRangeType(d, w, m)

	parseGitHub()
	printGitHub(j)
}
