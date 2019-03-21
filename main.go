package main

import (
	"flag"
	"log"

	"github.com/TakumiKaribe/trengo/requester"
)

var (
	lang string
	w    bool
	m    bool
	g    bool
)

func main() {
	flag.StringVar(&lang, "l", "", "language name")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")

	flag.Parse()
	if w && m {
		log.Fatalf("too many parameters.")
	}

	var githubResponse requester.GitHubResponse
	rangeType := requester.Daily
	if w {
		rangeType = requester.Weekly
	}
	if m {
		rangeType = requester.Monthly
	}

	githubResponse = requester.ParseGitHub(rangeType, lang)
	githubResponse.CLIPrint()
}
