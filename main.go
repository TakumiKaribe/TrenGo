package main

import (
	"flag"
	"log"

	"github.com/maka-nai/trengo/requester"
)

var (
	lang string
	w    bool
	m    bool
	j    bool
	g    bool
)

func main() {
	flag.StringVar(&lang, "l", "", "language name")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")
	flag.BoolVar(&g, "g", true, "GitHub search")
	flag.BoolVar(&j, "j", false, "json format")

	flag.Parse()
	checkSearchRange()

	var githubResponse requester.GitHubResponse
	if g {
		rangeType := requester.Daily
		if w {
			rangeType = requester.Weekly
		}
		if m {
			rangeType = requester.Monthly
		}

		githubResponse = requester.ParseGitHub(rangeType, lang)
		if j {
			githubResponse.JSONPrint()
		} else {
			githubResponse.CLIPrint()
		}
	}
}

func checkSearchRange() {
	if w && m {
		log.Fatalf("too many parameters.")
	}
}
